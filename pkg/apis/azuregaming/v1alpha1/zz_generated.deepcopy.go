// +build !ignore_autogenerated

/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DedicatedGameServer) DeepCopyInto(out *DedicatedGameServer) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DedicatedGameServer.
func (in *DedicatedGameServer) DeepCopy() *DedicatedGameServer {
	if in == nil {
		return nil
	}
	out := new(DedicatedGameServer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DedicatedGameServer) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DedicatedGameServerCollection) DeepCopyInto(out *DedicatedGameServerCollection) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DedicatedGameServerCollection.
func (in *DedicatedGameServerCollection) DeepCopy() *DedicatedGameServerCollection {
	if in == nil {
		return nil
	}
	out := new(DedicatedGameServerCollection)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DedicatedGameServerCollection) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DedicatedGameServerCollectionList) DeepCopyInto(out *DedicatedGameServerCollectionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DedicatedGameServerCollection, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DedicatedGameServerCollectionList.
func (in *DedicatedGameServerCollectionList) DeepCopy() *DedicatedGameServerCollectionList {
	if in == nil {
		return nil
	}
	out := new(DedicatedGameServerCollectionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DedicatedGameServerCollectionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DedicatedGameServerCollectionSpec) DeepCopyInto(out *DedicatedGameServerCollectionSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DedicatedGameServerCollectionSpec.
func (in *DedicatedGameServerCollectionSpec) DeepCopy() *DedicatedGameServerCollectionSpec {
	if in == nil {
		return nil
	}
	out := new(DedicatedGameServerCollectionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DedicatedGameServerCollectionStatus) DeepCopyInto(out *DedicatedGameServerCollectionStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DedicatedGameServerCollectionStatus.
func (in *DedicatedGameServerCollectionStatus) DeepCopy() *DedicatedGameServerCollectionStatus {
	if in == nil {
		return nil
	}
	out := new(DedicatedGameServerCollectionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DedicatedGameServerList) DeepCopyInto(out *DedicatedGameServerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DedicatedGameServer, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DedicatedGameServerList.
func (in *DedicatedGameServerList) DeepCopy() *DedicatedGameServerList {
	if in == nil {
		return nil
	}
	out := new(DedicatedGameServerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DedicatedGameServerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DedicatedGameServerSpec) DeepCopyInto(out *DedicatedGameServerSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DedicatedGameServerSpec.
func (in *DedicatedGameServerSpec) DeepCopy() *DedicatedGameServerSpec {
	if in == nil {
		return nil
	}
	out := new(DedicatedGameServerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DedicatedGameServerStatus) DeepCopyInto(out *DedicatedGameServerStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DedicatedGameServerStatus.
func (in *DedicatedGameServerStatus) DeepCopy() *DedicatedGameServerStatus {
	if in == nil {
		return nil
	}
	out := new(DedicatedGameServerStatus)
	in.DeepCopyInto(out)
	return out
}