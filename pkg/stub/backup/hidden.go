package backup

import (
	"github.com/Percona-Lab/percona-server-mongodb-operator/internal/mongod"
	"github.com/Percona-Lab/percona-server-mongodb-operator/internal/util"
	"github.com/Percona-Lab/percona-server-mongodb-operator/pkg/apis/psmdb/v1alpha1"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	BackupDataVolClaimName = "mongod-backup"
	statefulSetSuffix      = "backup"
)

var (
	secretFileMode int32 = 0060
)

// NewBackupStatefulSet returns a Hidden-node stateful set for backups
func (c *Controller) NewBackupStatefulSet(m *v1alpha1.PerconaServerMongoDB, replset *v1alpha1.ReplsetSpec, mongodResources, bkpResources corev1.ResourceRequirements, runUID *int64) *appsv1.StatefulSet {
	ls := util.LabelsForPerconaServerMongoDBReplset(m, replset)
	set := util.NewStatefulSet(m, m.Name+"-"+replset.Name+"-"+statefulSetSuffix)
	set.Spec = appsv1.StatefulSetSpec{
		ServiceName: m.Name + "-" + replset.Name,
		Selector: &metav1.LabelSelector{
			MatchLabels: ls,
		},
		Template: corev1.PodTemplateSpec{
			ObjectMeta: metav1.ObjectMeta{
				Labels: ls,
			},
			Spec: corev1.PodSpec{
				Affinity: mongod.NewPodAffinity(ls),
				Containers: []corev1.Container{
					mongod.NewBackupContainer(m, replset, mongodResources, runUID),
					c.NewAgentContainer(replset),
				},
				SecurityContext: &corev1.PodSecurityContext{
					FSGroup: runUID,
				},
				Volumes: []corev1.Volume{
					{
						Name: m.Spec.Secrets.Key,
						VolumeSource: corev1.VolumeSource{
							Secret: &corev1.SecretVolumeSource{
								DefaultMode: &secretFileMode,
								SecretName:  m.Spec.Secrets.Key,
								Optional:    &util.FalseVar,
							},
						},
					},
				},
			},
		},
		VolumeClaimTemplates: []corev1.PersistentVolumeClaim{
			util.NewPersistentVolumeClaim(m, mongodResources, mongod.MongodDataVolClaimName, replset.StorageClass),
			util.NewPersistentVolumeClaim(m, bkpResources, BackupDataVolClaimName, ""),
		},
	}
	util.AddOwnerRefToObject(set, util.AsOwner(m))
	return set
}
