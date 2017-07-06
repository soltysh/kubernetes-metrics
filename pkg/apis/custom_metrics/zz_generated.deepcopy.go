// +build !ignore_autogenerated

/*
Copyright 2017 The Kubernetes Authors.

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

// This file was autogenerated by deepcopy-gen. Do not edit it manually!

package custom_metrics

import (
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	reflect "reflect"
)

// Deprecated: register deep-copy functions.
func init() {
	SchemeBuilder.Register(RegisterDeepCopies)
}

// Deprecated: RegisterDeepCopies adds deep-copy functions to the given scheme. Public
// to allow building arbitrary schemes.
func RegisterDeepCopies(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedDeepCopyFuncs(
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*MetricValue).DeepCopyInto(out.(*MetricValue))
			return nil
		}, InType: reflect.TypeOf(&MetricValue{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*MetricValueList).DeepCopyInto(out.(*MetricValueList))
			return nil
		}, InType: reflect.TypeOf(&MetricValueList{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*ObjectReference).DeepCopyInto(out.(*ObjectReference))
			return nil
		}, InType: reflect.TypeOf(&ObjectReference{})},
	)
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetricValue) DeepCopyInto(out *MetricValue) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.DescribedObject = in.DescribedObject
	in.Timestamp.DeepCopyInto(&out.Timestamp)
	if in.WindowSeconds != nil {
		in, out := &in.WindowSeconds, &out.WindowSeconds
		if *in == nil {
			*out = nil
		} else {
			*out = new(int64)
			**out = **in
		}
	}
	out.Value = in.Value.DeepCopy()
	return
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, creating a new MetricValue.
func (x *MetricValue) DeepCopy() *MetricValue {
	if x == nil {
		return nil
	}
	out := new(MetricValue)
	x.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (x *MetricValue) DeepCopyObject() runtime.Object {
	if c := x.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetricValueList) DeepCopyInto(out *MetricValueList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MetricValue, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, creating a new MetricValueList.
func (x *MetricValueList) DeepCopy() *MetricValueList {
	if x == nil {
		return nil
	}
	out := new(MetricValueList)
	x.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (x *MetricValueList) DeepCopyObject() runtime.Object {
	if c := x.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObjectReference) DeepCopyInto(out *ObjectReference) {
	*out = *in
	return
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, creating a new ObjectReference.
func (x *ObjectReference) DeepCopy() *ObjectReference {
	if x == nil {
		return nil
	}
	out := new(ObjectReference)
	x.DeepCopyInto(out)
	return out
}
