//go:build !ignore_autogenerated

// Code generated by controller-gen. DO NOT EDIT.

package v1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RBDPVCBackup) DeepCopyInto(out *RBDPVCBackup) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RBDPVCBackup.
func (in *RBDPVCBackup) DeepCopy() *RBDPVCBackup {
	if in == nil {
		return nil
	}
	out := new(RBDPVCBackup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RBDPVCBackup) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RBDPVCBackupList) DeepCopyInto(out *RBDPVCBackupList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]RBDPVCBackup, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RBDPVCBackupList.
func (in *RBDPVCBackupList) DeepCopy() *RBDPVCBackupList {
	if in == nil {
		return nil
	}
	out := new(RBDPVCBackupList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RBDPVCBackupList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RBDPVCBackupSpec) DeepCopyInto(out *RBDPVCBackupSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RBDPVCBackupSpec.
func (in *RBDPVCBackupSpec) DeepCopy() *RBDPVCBackupSpec {
	if in == nil {
		return nil
	}
	out := new(RBDPVCBackupSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RBDPVCBackupStatus) DeepCopyInto(out *RBDPVCBackupStatus) {
	*out = *in
	in.CreatedAt.DeepCopyInto(&out.CreatedAt)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RBDPVCBackupStatus.
func (in *RBDPVCBackupStatus) DeepCopy() *RBDPVCBackupStatus {
	if in == nil {
		return nil
	}
	out := new(RBDPVCBackupStatus)
	in.DeepCopyInto(out)
	return out
}
