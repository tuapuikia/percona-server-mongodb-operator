// +build !ignore_autogenerated

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupSpec) DeepCopyInto(out *BackupSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupSpec.
func (in *BackupSpec) DeepCopy() *BackupSpec {
	if in == nil {
		return nil
	}
	out := new(BackupSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupStatus) DeepCopyInto(out *BackupStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupStatus.
func (in *BackupStatus) DeepCopy() *BackupStatus {
	if in == nil {
		return nil
	}
	out := new(BackupStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MongodSpec) DeepCopyInto(out *MongodSpec) {
	*out = *in
	if in.Net != nil {
		in, out := &in.Net, &out.Net
		*out = new(MongodSpecNet)
		**out = **in
	}
	if in.AuditLog != nil {
		in, out := &in.AuditLog, &out.AuditLog
		*out = new(MongodSpecAuditLog)
		**out = **in
	}
	if in.OperationProfiling != nil {
		in, out := &in.OperationProfiling, &out.OperationProfiling
		*out = new(MongodSpecOperationProfiling)
		**out = **in
	}
	if in.Replication != nil {
		in, out := &in.Replication, &out.Replication
		*out = new(MongodSpecReplication)
		**out = **in
	}
	if in.Security != nil {
		in, out := &in.Security, &out.Security
		*out = new(MongodSpecSecurity)
		**out = **in
	}
	if in.SetParameter != nil {
		in, out := &in.SetParameter, &out.SetParameter
		*out = new(MongodSpecSetParameter)
		**out = **in
	}
	if in.Storage != nil {
		in, out := &in.Storage, &out.Storage
		*out = new(MongodSpecStorage)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MongodSpec.
func (in *MongodSpec) DeepCopy() *MongodSpec {
	if in == nil {
		return nil
	}
	out := new(MongodSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MongodSpecAuditLog) DeepCopyInto(out *MongodSpecAuditLog) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MongodSpecAuditLog.
func (in *MongodSpecAuditLog) DeepCopy() *MongodSpecAuditLog {
	if in == nil {
		return nil
	}
	out := new(MongodSpecAuditLog)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MongodSpecInMemory) DeepCopyInto(out *MongodSpecInMemory) {
	*out = *in
	if in.EngineConfig != nil {
		in, out := &in.EngineConfig, &out.EngineConfig
		*out = new(MongodSpecInMemoryEngineConfig)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MongodSpecInMemory.
func (in *MongodSpecInMemory) DeepCopy() *MongodSpecInMemory {
	if in == nil {
		return nil
	}
	out := new(MongodSpecInMemory)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MongodSpecInMemoryEngineConfig) DeepCopyInto(out *MongodSpecInMemoryEngineConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MongodSpecInMemoryEngineConfig.
func (in *MongodSpecInMemoryEngineConfig) DeepCopy() *MongodSpecInMemoryEngineConfig {
	if in == nil {
		return nil
	}
	out := new(MongodSpecInMemoryEngineConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MongodSpecMMAPv1) DeepCopyInto(out *MongodSpecMMAPv1) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MongodSpecMMAPv1.
func (in *MongodSpecMMAPv1) DeepCopy() *MongodSpecMMAPv1 {
	if in == nil {
		return nil
	}
	out := new(MongodSpecMMAPv1)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MongodSpecNet) DeepCopyInto(out *MongodSpecNet) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MongodSpecNet.
func (in *MongodSpecNet) DeepCopy() *MongodSpecNet {
	if in == nil {
		return nil
	}
	out := new(MongodSpecNet)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MongodSpecOperationProfiling) DeepCopyInto(out *MongodSpecOperationProfiling) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MongodSpecOperationProfiling.
func (in *MongodSpecOperationProfiling) DeepCopy() *MongodSpecOperationProfiling {
	if in == nil {
		return nil
	}
	out := new(MongodSpecOperationProfiling)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MongodSpecReplication) DeepCopyInto(out *MongodSpecReplication) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MongodSpecReplication.
func (in *MongodSpecReplication) DeepCopy() *MongodSpecReplication {
	if in == nil {
		return nil
	}
	out := new(MongodSpecReplication)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MongodSpecSecurity) DeepCopyInto(out *MongodSpecSecurity) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MongodSpecSecurity.
func (in *MongodSpecSecurity) DeepCopy() *MongodSpecSecurity {
	if in == nil {
		return nil
	}
	out := new(MongodSpecSecurity)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MongodSpecSetParameter) DeepCopyInto(out *MongodSpecSetParameter) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MongodSpecSetParameter.
func (in *MongodSpecSetParameter) DeepCopy() *MongodSpecSetParameter {
	if in == nil {
		return nil
	}
	out := new(MongodSpecSetParameter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MongodSpecStorage) DeepCopyInto(out *MongodSpecStorage) {
	*out = *in
	if in.InMemory != nil {
		in, out := &in.InMemory, &out.InMemory
		*out = new(MongodSpecInMemory)
		(*in).DeepCopyInto(*out)
	}
	if in.MMAPv1 != nil {
		in, out := &in.MMAPv1, &out.MMAPv1
		*out = new(MongodSpecMMAPv1)
		**out = **in
	}
	if in.WiredTiger != nil {
		in, out := &in.WiredTiger, &out.WiredTiger
		*out = new(MongodSpecWiredTiger)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MongodSpecStorage.
func (in *MongodSpecStorage) DeepCopy() *MongodSpecStorage {
	if in == nil {
		return nil
	}
	out := new(MongodSpecStorage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MongodSpecWiredTiger) DeepCopyInto(out *MongodSpecWiredTiger) {
	*out = *in
	if in.CollectionConfig != nil {
		in, out := &in.CollectionConfig, &out.CollectionConfig
		*out = new(MongodSpecWiredTigerCollectionConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.EngineConfig != nil {
		in, out := &in.EngineConfig, &out.EngineConfig
		*out = new(MongodSpecWiredTigerEngineConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.IndexConfig != nil {
		in, out := &in.IndexConfig, &out.IndexConfig
		*out = new(MongodSpecWiredTigerIndexConfig)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MongodSpecWiredTiger.
func (in *MongodSpecWiredTiger) DeepCopy() *MongodSpecWiredTiger {
	if in == nil {
		return nil
	}
	out := new(MongodSpecWiredTiger)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MongodSpecWiredTigerCollectionConfig) DeepCopyInto(out *MongodSpecWiredTigerCollectionConfig) {
	*out = *in
	if in.BlockCompressor != nil {
		in, out := &in.BlockCompressor, &out.BlockCompressor
		*out = new(WiredTigerCompressor)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MongodSpecWiredTigerCollectionConfig.
func (in *MongodSpecWiredTigerCollectionConfig) DeepCopy() *MongodSpecWiredTigerCollectionConfig {
	if in == nil {
		return nil
	}
	out := new(MongodSpecWiredTigerCollectionConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MongodSpecWiredTigerEngineConfig) DeepCopyInto(out *MongodSpecWiredTigerEngineConfig) {
	*out = *in
	if in.JournalCompressor != nil {
		in, out := &in.JournalCompressor, &out.JournalCompressor
		*out = new(WiredTigerCompressor)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MongodSpecWiredTigerEngineConfig.
func (in *MongodSpecWiredTigerEngineConfig) DeepCopy() *MongodSpecWiredTigerEngineConfig {
	if in == nil {
		return nil
	}
	out := new(MongodSpecWiredTigerEngineConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MongodSpecWiredTigerIndexConfig) DeepCopyInto(out *MongodSpecWiredTigerIndexConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MongodSpecWiredTigerIndexConfig.
func (in *MongodSpecWiredTigerIndexConfig) DeepCopy() *MongodSpecWiredTigerIndexConfig {
	if in == nil {
		return nil
	}
	out := new(MongodSpecWiredTigerIndexConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MongosSpec) DeepCopyInto(out *MongosSpec) {
	*out = *in
	if in.ResourcesSpec != nil {
		in, out := &in.ResourcesSpec, &out.ResourcesSpec
		*out = new(ResourcesSpec)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MongosSpec.
func (in *MongosSpec) DeepCopy() *MongosSpec {
	if in == nil {
		return nil
	}
	out := new(MongosSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PerconaServerMongoDB) DeepCopyInto(out *PerconaServerMongoDB) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PerconaServerMongoDB.
func (in *PerconaServerMongoDB) DeepCopy() *PerconaServerMongoDB {
	if in == nil {
		return nil
	}
	out := new(PerconaServerMongoDB)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PerconaServerMongoDB) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PerconaServerMongoDBList) DeepCopyInto(out *PerconaServerMongoDBList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PerconaServerMongoDB, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PerconaServerMongoDBList.
func (in *PerconaServerMongoDBList) DeepCopy() *PerconaServerMongoDBList {
	if in == nil {
		return nil
	}
	out := new(PerconaServerMongoDBList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PerconaServerMongoDBList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PerconaServerMongoDBSpec) DeepCopyInto(out *PerconaServerMongoDBSpec) {
	*out = *in
	if in.Platform != nil {
		in, out := &in.Platform, &out.Platform
		*out = new(Platform)
		**out = **in
	}
	if in.Mongod != nil {
		in, out := &in.Mongod, &out.Mongod
		*out = new(MongodSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Replsets != nil {
		in, out := &in.Replsets, &out.Replsets
		*out = make([]*ReplsetSpec, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(ReplsetSpec)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.Secrets != nil {
		in, out := &in.Secrets, &out.Secrets
		*out = new(SecretsSpec)
		**out = **in
	}
	if in.Backups != nil {
		in, out := &in.Backups, &out.Backups
		*out = make([]*BackupSpec, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(BackupSpec)
				**out = **in
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PerconaServerMongoDBSpec.
func (in *PerconaServerMongoDBSpec) DeepCopy() *PerconaServerMongoDBSpec {
	if in == nil {
		return nil
	}
	out := new(PerconaServerMongoDBSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PerconaServerMongoDBStatus) DeepCopyInto(out *PerconaServerMongoDBStatus) {
	*out = *in
	if in.Backups != nil {
		in, out := &in.Backups, &out.Backups
		*out = make([]*BackupStatus, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(BackupStatus)
				**out = **in
			}
		}
	}
	if in.Replsets != nil {
		in, out := &in.Replsets, &out.Replsets
		*out = make([]*ReplsetStatus, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(ReplsetStatus)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PerconaServerMongoDBStatus.
func (in *PerconaServerMongoDBStatus) DeepCopy() *PerconaServerMongoDBStatus {
	if in == nil {
		return nil
	}
	out := new(PerconaServerMongoDBStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReplsetMemberStatus) DeepCopyInto(out *ReplsetMemberStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReplsetMemberStatus.
func (in *ReplsetMemberStatus) DeepCopy() *ReplsetMemberStatus {
	if in == nil {
		return nil
	}
	out := new(ReplsetMemberStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReplsetSpec) DeepCopyInto(out *ReplsetSpec) {
	*out = *in
	if in.ResourcesSpec != nil {
		in, out := &in.ResourcesSpec, &out.ResourcesSpec
		*out = new(ResourcesSpec)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReplsetSpec.
func (in *ReplsetSpec) DeepCopy() *ReplsetSpec {
	if in == nil {
		return nil
	}
	out := new(ReplsetSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReplsetStatus) DeepCopyInto(out *ReplsetStatus) {
	*out = *in
	if in.Pods != nil {
		in, out := &in.Pods, &out.Pods
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Members != nil {
		in, out := &in.Members, &out.Members
		*out = make([]*ReplsetMemberStatus, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(ReplsetMemberStatus)
				**out = **in
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReplsetStatus.
func (in *ReplsetStatus) DeepCopy() *ReplsetStatus {
	if in == nil {
		return nil
	}
	out := new(ReplsetStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceSpecRequirements) DeepCopyInto(out *ResourceSpecRequirements) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceSpecRequirements.
func (in *ResourceSpecRequirements) DeepCopy() *ResourceSpecRequirements {
	if in == nil {
		return nil
	}
	out := new(ResourceSpecRequirements)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourcesSpec) DeepCopyInto(out *ResourcesSpec) {
	*out = *in
	if in.Limits != nil {
		in, out := &in.Limits, &out.Limits
		*out = new(ResourceSpecRequirements)
		**out = **in
	}
	if in.Requests != nil {
		in, out := &in.Requests, &out.Requests
		*out = new(ResourceSpecRequirements)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourcesSpec.
func (in *ResourcesSpec) DeepCopy() *ResourcesSpec {
	if in == nil {
		return nil
	}
	out := new(ResourcesSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretsSpec) DeepCopyInto(out *SecretsSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretsSpec.
func (in *SecretsSpec) DeepCopy() *SecretsSpec {
	if in == nil {
		return nil
	}
	out := new(SecretsSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServerVersion) DeepCopyInto(out *ServerVersion) {
	*out = *in
	out.Info = in.Info
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServerVersion.
func (in *ServerVersion) DeepCopy() *ServerVersion {
	if in == nil {
		return nil
	}
	out := new(ServerVersion)
	in.DeepCopyInto(out)
	return out
}
