//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by ack-generate. DO NOT EDIT.

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	corev1alpha1 "github.com/aws-controllers-k8s/runtime/apis/core/v1alpha1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AdvancedEventSelector) DeepCopyInto(out *AdvancedEventSelector) {
	*out = *in
	if in.FieldSelectors != nil {
		in, out := &in.FieldSelectors, &out.FieldSelectors
		*out = make([]*AdvancedFieldSelector, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(AdvancedFieldSelector)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AdvancedEventSelector.
func (in *AdvancedEventSelector) DeepCopy() *AdvancedEventSelector {
	if in == nil {
		return nil
	}
	out := new(AdvancedEventSelector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AdvancedFieldSelector) DeepCopyInto(out *AdvancedFieldSelector) {
	*out = *in
	if in.EndsWith != nil {
		in, out := &in.EndsWith, &out.EndsWith
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Equals != nil {
		in, out := &in.Equals, &out.Equals
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Field != nil {
		in, out := &in.Field, &out.Field
		*out = new(string)
		**out = **in
	}
	if in.NotEndsWith != nil {
		in, out := &in.NotEndsWith, &out.NotEndsWith
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.NotEquals != nil {
		in, out := &in.NotEquals, &out.NotEquals
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.NotStartsWith != nil {
		in, out := &in.NotStartsWith, &out.NotStartsWith
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.StartsWith != nil {
		in, out := &in.StartsWith, &out.StartsWith
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AdvancedFieldSelector.
func (in *AdvancedFieldSelector) DeepCopy() *AdvancedFieldSelector {
	if in == nil {
		return nil
	}
	out := new(AdvancedFieldSelector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataResource) DeepCopyInto(out *DataResource) {
	*out = *in
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataResource.
func (in *DataResource) DeepCopy() *DataResource {
	if in == nil {
		return nil
	}
	out := new(DataResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Event) DeepCopyInto(out *Event) {
	*out = *in
	if in.AccessKeyID != nil {
		in, out := &in.AccessKeyID, &out.AccessKeyID
		*out = new(string)
		**out = **in
	}
	if in.CloudTrailEvent != nil {
		in, out := &in.CloudTrailEvent, &out.CloudTrailEvent
		*out = new(string)
		**out = **in
	}
	if in.EventID != nil {
		in, out := &in.EventID, &out.EventID
		*out = new(string)
		**out = **in
	}
	if in.EventName != nil {
		in, out := &in.EventName, &out.EventName
		*out = new(string)
		**out = **in
	}
	if in.EventSource != nil {
		in, out := &in.EventSource, &out.EventSource
		*out = new(string)
		**out = **in
	}
	if in.EventTime != nil {
		in, out := &in.EventTime, &out.EventTime
		*out = (*in).DeepCopy()
	}
	if in.ReadOnly != nil {
		in, out := &in.ReadOnly, &out.ReadOnly
		*out = new(string)
		**out = **in
	}
	if in.Username != nil {
		in, out := &in.Username, &out.Username
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Event.
func (in *Event) DeepCopy() *Event {
	if in == nil {
		return nil
	}
	out := new(Event)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventDataStore) DeepCopyInto(out *EventDataStore) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventDataStore.
func (in *EventDataStore) DeepCopy() *EventDataStore {
	if in == nil {
		return nil
	}
	out := new(EventDataStore)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EventDataStore) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventDataStoreList) DeepCopyInto(out *EventDataStoreList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]EventDataStore, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventDataStoreList.
func (in *EventDataStoreList) DeepCopy() *EventDataStoreList {
	if in == nil {
		return nil
	}
	out := new(EventDataStoreList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EventDataStoreList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventDataStoreSpec) DeepCopyInto(out *EventDataStoreSpec) {
	*out = *in
	if in.AdvancedEventSelectors != nil {
		in, out := &in.AdvancedEventSelectors, &out.AdvancedEventSelectors
		*out = make([]*AdvancedEventSelector, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(AdvancedEventSelector)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.MultiRegionEnabled != nil {
		in, out := &in.MultiRegionEnabled, &out.MultiRegionEnabled
		*out = new(bool)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.OrganizationEnabled != nil {
		in, out := &in.OrganizationEnabled, &out.OrganizationEnabled
		*out = new(bool)
		**out = **in
	}
	if in.RetentionPeriod != nil {
		in, out := &in.RetentionPeriod, &out.RetentionPeriod
		*out = new(int64)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]*Tag, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(Tag)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.TerminationProtectionEnabled != nil {
		in, out := &in.TerminationProtectionEnabled, &out.TerminationProtectionEnabled
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventDataStoreSpec.
func (in *EventDataStoreSpec) DeepCopy() *EventDataStoreSpec {
	if in == nil {
		return nil
	}
	out := new(EventDataStoreSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventDataStoreStatus) DeepCopyInto(out *EventDataStoreStatus) {
	*out = *in
	if in.ACKResourceMetadata != nil {
		in, out := &in.ACKResourceMetadata, &out.ACKResourceMetadata
		*out = new(corev1alpha1.ResourceMetadata)
		(*in).DeepCopyInto(*out)
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]*corev1alpha1.Condition, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(corev1alpha1.Condition)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.CreatedTimestamp != nil {
		in, out := &in.CreatedTimestamp, &out.CreatedTimestamp
		*out = (*in).DeepCopy()
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(string)
		**out = **in
	}
	if in.UpdatedTimestamp != nil {
		in, out := &in.UpdatedTimestamp, &out.UpdatedTimestamp
		*out = (*in).DeepCopy()
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventDataStoreStatus.
func (in *EventDataStoreStatus) DeepCopy() *EventDataStoreStatus {
	if in == nil {
		return nil
	}
	out := new(EventDataStoreStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventDataStore_SDK) DeepCopyInto(out *EventDataStore_SDK) {
	*out = *in
	if in.AdvancedEventSelectors != nil {
		in, out := &in.AdvancedEventSelectors, &out.AdvancedEventSelectors
		*out = make([]*AdvancedEventSelector, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(AdvancedEventSelector)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.CreatedTimestamp != nil {
		in, out := &in.CreatedTimestamp, &out.CreatedTimestamp
		*out = (*in).DeepCopy()
	}
	if in.EventDataStoreARN != nil {
		in, out := &in.EventDataStoreARN, &out.EventDataStoreARN
		*out = new(string)
		**out = **in
	}
	if in.MultiRegionEnabled != nil {
		in, out := &in.MultiRegionEnabled, &out.MultiRegionEnabled
		*out = new(bool)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.OrganizationEnabled != nil {
		in, out := &in.OrganizationEnabled, &out.OrganizationEnabled
		*out = new(bool)
		**out = **in
	}
	if in.RetentionPeriod != nil {
		in, out := &in.RetentionPeriod, &out.RetentionPeriod
		*out = new(int64)
		**out = **in
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(string)
		**out = **in
	}
	if in.TerminationProtectionEnabled != nil {
		in, out := &in.TerminationProtectionEnabled, &out.TerminationProtectionEnabled
		*out = new(bool)
		**out = **in
	}
	if in.UpdatedTimestamp != nil {
		in, out := &in.UpdatedTimestamp, &out.UpdatedTimestamp
		*out = (*in).DeepCopy()
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventDataStore_SDK.
func (in *EventDataStore_SDK) DeepCopy() *EventDataStore_SDK {
	if in == nil {
		return nil
	}
	out := new(EventDataStore_SDK)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventSelector) DeepCopyInto(out *EventSelector) {
	*out = *in
	if in.IncludeManagementEvents != nil {
		in, out := &in.IncludeManagementEvents, &out.IncludeManagementEvents
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventSelector.
func (in *EventSelector) DeepCopy() *EventSelector {
	if in == nil {
		return nil
	}
	out := new(EventSelector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LookupAttribute) DeepCopyInto(out *LookupAttribute) {
	*out = *in
	if in.AttributeValue != nil {
		in, out := &in.AttributeValue, &out.AttributeValue
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LookupAttribute.
func (in *LookupAttribute) DeepCopy() *LookupAttribute {
	if in == nil {
		return nil
	}
	out := new(LookupAttribute)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PublicKey) DeepCopyInto(out *PublicKey) {
	*out = *in
	if in.Fingerprint != nil {
		in, out := &in.Fingerprint, &out.Fingerprint
		*out = new(string)
		**out = **in
	}
	if in.ValidityEndTime != nil {
		in, out := &in.ValidityEndTime, &out.ValidityEndTime
		*out = (*in).DeepCopy()
	}
	if in.ValidityStartTime != nil {
		in, out := &in.ValidityStartTime, &out.ValidityStartTime
		*out = (*in).DeepCopy()
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PublicKey.
func (in *PublicKey) DeepCopy() *PublicKey {
	if in == nil {
		return nil
	}
	out := new(PublicKey)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Query) DeepCopyInto(out *Query) {
	*out = *in
	if in.CreationTime != nil {
		in, out := &in.CreationTime, &out.CreationTime
		*out = (*in).DeepCopy()
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Query.
func (in *Query) DeepCopy() *Query {
	if in == nil {
		return nil
	}
	out := new(Query)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QueryStatisticsForDescribeQuery) DeepCopyInto(out *QueryStatisticsForDescribeQuery) {
	*out = *in
	if in.CreationTime != nil {
		in, out := &in.CreationTime, &out.CreationTime
		*out = (*in).DeepCopy()
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QueryStatisticsForDescribeQuery.
func (in *QueryStatisticsForDescribeQuery) DeepCopy() *QueryStatisticsForDescribeQuery {
	if in == nil {
		return nil
	}
	out := new(QueryStatisticsForDescribeQuery)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Resource) DeepCopyInto(out *Resource) {
	*out = *in
	if in.ResourceName != nil {
		in, out := &in.ResourceName, &out.ResourceName
		*out = new(string)
		**out = **in
	}
	if in.ResourceType != nil {
		in, out := &in.ResourceType, &out.ResourceType
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Resource.
func (in *Resource) DeepCopy() *Resource {
	if in == nil {
		return nil
	}
	out := new(Resource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceTag) DeepCopyInto(out *ResourceTag) {
	*out = *in
	if in.ResourceID != nil {
		in, out := &in.ResourceID, &out.ResourceID
		*out = new(string)
		**out = **in
	}
	if in.TagsList != nil {
		in, out := &in.TagsList, &out.TagsList
		*out = make([]*Tag, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(Tag)
				(*in).DeepCopyInto(*out)
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceTag.
func (in *ResourceTag) DeepCopy() *ResourceTag {
	if in == nil {
		return nil
	}
	out := new(ResourceTag)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Tag) DeepCopyInto(out *Tag) {
	*out = *in
	if in.Key != nil {
		in, out := &in.Key, &out.Key
		*out = new(string)
		**out = **in
	}
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Tag.
func (in *Tag) DeepCopy() *Tag {
	if in == nil {
		return nil
	}
	out := new(Tag)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Trail) DeepCopyInto(out *Trail) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Trail.
func (in *Trail) DeepCopy() *Trail {
	if in == nil {
		return nil
	}
	out := new(Trail)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Trail) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TrailInfo) DeepCopyInto(out *TrailInfo) {
	*out = *in
	if in.HomeRegion != nil {
		in, out := &in.HomeRegion, &out.HomeRegion
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.TrailARN != nil {
		in, out := &in.TrailARN, &out.TrailARN
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TrailInfo.
func (in *TrailInfo) DeepCopy() *TrailInfo {
	if in == nil {
		return nil
	}
	out := new(TrailInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TrailList) DeepCopyInto(out *TrailList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Trail, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TrailList.
func (in *TrailList) DeepCopy() *TrailList {
	if in == nil {
		return nil
	}
	out := new(TrailList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TrailList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TrailSpec) DeepCopyInto(out *TrailSpec) {
	*out = *in
	if in.CloudWatchLogsLogGroupARN != nil {
		in, out := &in.CloudWatchLogsLogGroupARN, &out.CloudWatchLogsLogGroupARN
		*out = new(string)
		**out = **in
	}
	if in.CloudWatchLogsRoleARN != nil {
		in, out := &in.CloudWatchLogsRoleARN, &out.CloudWatchLogsRoleARN
		*out = new(string)
		**out = **in
	}
	if in.EnableLogFileValidation != nil {
		in, out := &in.EnableLogFileValidation, &out.EnableLogFileValidation
		*out = new(bool)
		**out = **in
	}
	if in.IncludeGlobalServiceEvents != nil {
		in, out := &in.IncludeGlobalServiceEvents, &out.IncludeGlobalServiceEvents
		*out = new(bool)
		**out = **in
	}
	if in.IsMultiRegionTrail != nil {
		in, out := &in.IsMultiRegionTrail, &out.IsMultiRegionTrail
		*out = new(bool)
		**out = **in
	}
	if in.IsOrganizationTrail != nil {
		in, out := &in.IsOrganizationTrail, &out.IsOrganizationTrail
		*out = new(bool)
		**out = **in
	}
	if in.KMSKeyID != nil {
		in, out := &in.KMSKeyID, &out.KMSKeyID
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.S3BucketName != nil {
		in, out := &in.S3BucketName, &out.S3BucketName
		*out = new(string)
		**out = **in
	}
	if in.S3KeyPrefix != nil {
		in, out := &in.S3KeyPrefix, &out.S3KeyPrefix
		*out = new(string)
		**out = **in
	}
	if in.SNSTopicName != nil {
		in, out := &in.SNSTopicName, &out.SNSTopicName
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]*Tag, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(Tag)
				(*in).DeepCopyInto(*out)
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TrailSpec.
func (in *TrailSpec) DeepCopy() *TrailSpec {
	if in == nil {
		return nil
	}
	out := new(TrailSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TrailStatus) DeepCopyInto(out *TrailStatus) {
	*out = *in
	if in.ACKResourceMetadata != nil {
		in, out := &in.ACKResourceMetadata, &out.ACKResourceMetadata
		*out = new(corev1alpha1.ResourceMetadata)
		(*in).DeepCopyInto(*out)
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]*corev1alpha1.Condition, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(corev1alpha1.Condition)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.LogFileValidationEnabled != nil {
		in, out := &in.LogFileValidationEnabled, &out.LogFileValidationEnabled
		*out = new(bool)
		**out = **in
	}
	if in.SNSTopicARN != nil {
		in, out := &in.SNSTopicARN, &out.SNSTopicARN
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TrailStatus.
func (in *TrailStatus) DeepCopy() *TrailStatus {
	if in == nil {
		return nil
	}
	out := new(TrailStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Trail_SDK) DeepCopyInto(out *Trail_SDK) {
	*out = *in
	if in.CloudWatchLogsLogGroupARN != nil {
		in, out := &in.CloudWatchLogsLogGroupARN, &out.CloudWatchLogsLogGroupARN
		*out = new(string)
		**out = **in
	}
	if in.CloudWatchLogsRoleARN != nil {
		in, out := &in.CloudWatchLogsRoleARN, &out.CloudWatchLogsRoleARN
		*out = new(string)
		**out = **in
	}
	if in.HasCustomEventSelectors != nil {
		in, out := &in.HasCustomEventSelectors, &out.HasCustomEventSelectors
		*out = new(bool)
		**out = **in
	}
	if in.HasInsightSelectors != nil {
		in, out := &in.HasInsightSelectors, &out.HasInsightSelectors
		*out = new(bool)
		**out = **in
	}
	if in.HomeRegion != nil {
		in, out := &in.HomeRegion, &out.HomeRegion
		*out = new(string)
		**out = **in
	}
	if in.IncludeGlobalServiceEvents != nil {
		in, out := &in.IncludeGlobalServiceEvents, &out.IncludeGlobalServiceEvents
		*out = new(bool)
		**out = **in
	}
	if in.IsMultiRegionTrail != nil {
		in, out := &in.IsMultiRegionTrail, &out.IsMultiRegionTrail
		*out = new(bool)
		**out = **in
	}
	if in.IsOrganizationTrail != nil {
		in, out := &in.IsOrganizationTrail, &out.IsOrganizationTrail
		*out = new(bool)
		**out = **in
	}
	if in.KMSKeyID != nil {
		in, out := &in.KMSKeyID, &out.KMSKeyID
		*out = new(string)
		**out = **in
	}
	if in.LogFileValidationEnabled != nil {
		in, out := &in.LogFileValidationEnabled, &out.LogFileValidationEnabled
		*out = new(bool)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.S3BucketName != nil {
		in, out := &in.S3BucketName, &out.S3BucketName
		*out = new(string)
		**out = **in
	}
	if in.S3KeyPrefix != nil {
		in, out := &in.S3KeyPrefix, &out.S3KeyPrefix
		*out = new(string)
		**out = **in
	}
	if in.SNSTopicARN != nil {
		in, out := &in.SNSTopicARN, &out.SNSTopicARN
		*out = new(string)
		**out = **in
	}
	if in.SNSTopicName != nil {
		in, out := &in.SNSTopicName, &out.SNSTopicName
		*out = new(string)
		**out = **in
	}
	if in.TrailARN != nil {
		in, out := &in.TrailARN, &out.TrailARN
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Trail_SDK.
func (in *Trail_SDK) DeepCopy() *Trail_SDK {
	if in == nil {
		return nil
	}
	out := new(Trail_SDK)
	in.DeepCopyInto(out)
	return out
}
