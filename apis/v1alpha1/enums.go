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

package v1alpha1

type EventCategory string

const (
	EventCategory_insight EventCategory = "insight"
)

type EventDataStoreStatus_SDK string

const (
	EventDataStoreStatus_SDK_CREATED          EventDataStoreStatus_SDK = "CREATED"
	EventDataStoreStatus_SDK_ENABLED          EventDataStoreStatus_SDK = "ENABLED"
	EventDataStoreStatus_SDK_PENDING_DELETION EventDataStoreStatus_SDK = "PENDING_DELETION"
)

type InsightType string

const (
	InsightType_ApiCallRateInsight  InsightType = "ApiCallRateInsight"
	InsightType_ApiErrorRateInsight InsightType = "ApiErrorRateInsight"
)

type LookupAttributeKey string

const (
	LookupAttributeKey_EventId      LookupAttributeKey = "EventId"
	LookupAttributeKey_EventName    LookupAttributeKey = "EventName"
	LookupAttributeKey_ReadOnly     LookupAttributeKey = "ReadOnly"
	LookupAttributeKey_Username     LookupAttributeKey = "Username"
	LookupAttributeKey_ResourceType LookupAttributeKey = "ResourceType"
	LookupAttributeKey_ResourceName LookupAttributeKey = "ResourceName"
	LookupAttributeKey_EventSource  LookupAttributeKey = "EventSource"
	LookupAttributeKey_AccessKeyId  LookupAttributeKey = "AccessKeyId"
)

type QueryStatus string

const (
	QueryStatus_QUEUED    QueryStatus = "QUEUED"
	QueryStatus_RUNNING   QueryStatus = "RUNNING"
	QueryStatus_FINISHED  QueryStatus = "FINISHED"
	QueryStatus_FAILED    QueryStatus = "FAILED"
	QueryStatus_CANCELLED QueryStatus = "CANCELLED"
	QueryStatus_TIMED_OUT QueryStatus = "TIMED_OUT"
)

type ReadWriteType string

const (
	ReadWriteType_ReadOnly  ReadWriteType = "ReadOnly"
	ReadWriteType_WriteOnly ReadWriteType = "WriteOnly"
	ReadWriteType_All       ReadWriteType = "All"
)
