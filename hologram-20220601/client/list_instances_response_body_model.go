// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *ListInstancesResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListInstancesResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v string) *ListInstancesResponseBody
	GetHttpStatusCode() *string
	SetInstanceList(v []*ListInstancesResponseBodyInstanceList) *ListInstancesResponseBody
	GetInstanceList() []*ListInstancesResponseBodyInstanceList
	SetRequestId(v string) *ListInstancesResponseBody
	GetRequestId() *string
	SetSuccess(v string) *ListInstancesResponseBody
	GetSuccess() *string
}

type ListInstancesResponseBody struct {
	// The error code returned if the request failed.
	//
	// example:
	//
	// 404
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned if the request failed.
	//
	// example:
	//
	// Internal server error.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *string `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The instances.
	InstanceList []*ListInstancesResponseBodyInstanceList `json:"InstanceList,omitempty" xml:"InstanceList,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// D1303CD4-AA70-5998-8025-F55B22C50840
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListInstancesResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListInstancesResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *ListInstancesResponseBody) GetInstanceList() []*ListInstancesResponseBodyInstanceList {
	return s.InstanceList
}

func (s *ListInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListInstancesResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *ListInstancesResponseBody) SetErrorCode(v string) *ListInstancesResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListInstancesResponseBody) SetErrorMessage(v string) *ListInstancesResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListInstancesResponseBody) SetHttpStatusCode(v string) *ListInstancesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListInstancesResponseBody) SetInstanceList(v []*ListInstancesResponseBodyInstanceList) *ListInstancesResponseBody {
	s.InstanceList = v
	return s
}

func (s *ListInstancesResponseBody) SetRequestId(v string) *ListInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInstancesResponseBody) SetSuccess(v string) *ListInstancesResponseBody {
	s.Success = &v
	return s
}

func (s *ListInstancesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListInstancesResponseBodyInstanceList struct {
	// The commodity code, which is the same as that on the Billing Management page.
	//
	// example:
	//
	// hologram_postpay_public_cn
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	// The time when the cluster was created.
	//
	// example:
	//
	// 2022-12-16T02:24:05Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// Indicates whether lakehouse acceleration is enabled.
	//
	// Valid values:
	//
	// 	- true
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- false
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// true
	EnableHiveAccess *string `json:"EnableHiveAccess,omitempty" xml:"EnableHiveAccess,omitempty"`
	// The list of endpoints.
	Endpoints []*ListInstancesResponseBodyInstanceListEndpoints `json:"Endpoints,omitempty" xml:"Endpoints,omitempty" type:"Repeated"`
	// The time when the cluster expires.
	//
	// example:
	//
	// 2023-05-04T16:00:00.000Z
	ExpirationTime *string `json:"ExpirationTime,omitempty" xml:"ExpirationTime,omitempty"`
	// The billing method of the instance. Valid values:
	//
	// Valid values:
	//
	// 	- PostPaid
	//
	//     <!-- -->
	//
	//     :
	//
	//     <!-- -->
	//
	//     pay-as-you-go
	//
	//     <!-- -->
	//
	//     .
	//
	// 	- PrePaid
	//
	//     <!-- -->
	//
	//     :
	//
	//     <!-- -->
	//
	//     subscription
	//
	//     <!-- -->
	//
	//     .
	//
	// example:
	//
	// PrePaid
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" xml:"InstanceChargeType,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// hgpostcn-cn-aaab9ad2d8fb
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the instance.
	//
	// example:
	//
	// test_instance
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The status of the instance.
	//
	// Valid values:
	//
	// 	- Creating
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- Running
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- Suspended
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- Allocating
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// Running
	InstanceStatus *string `json:"InstanceStatus,omitempty" xml:"InstanceStatus,omitempty"`
	// The type of the instance.
	//
	// Valid values:
	//
	// 	- Follower
	//
	//     <!-- -->
	//
	//     :
	//
	//     <!-- -->
	//
	//     read-only secondary instance
	//
	//     <!-- -->
	//
	//     .
	//
	// 	- Standard
	//
	//     <!-- -->
	//
	//     :
	//
	//     <!-- -->
	//
	//     normal instance
	//
	//     <!-- -->
	//
	//     .
	//
	// example:
	//
	// Standard
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The ID of the primary instance.
	//
	// example:
	//
	// hgprecn-cn-2r42sqvxm006
	LeaderInstanceId *string `json:"LeaderInstanceId,omitempty" xml:"LeaderInstanceId,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmvscak73zmby
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	StorageType     *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	// The reason for the suspension.
	//
	// example:
	//
	// Manual
	SuspendReason *string `json:"SuspendReason,omitempty" xml:"SuspendReason,omitempty"`
	// The tags that are added to the resource.
	Tags []*ListInstancesResponseBodyInstanceListTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The version of the cluster.
	//
	// example:
	//
	// 1.3.37
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
	// example:
	//
	// cn-hangzhou-h
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s ListInstancesResponseBodyInstanceList) String() string {
	return dara.Prettify(s)
}

func (s ListInstancesResponseBodyInstanceList) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBodyInstanceList) GetCommodityCode() *string {
	return s.CommodityCode
}

func (s *ListInstancesResponseBodyInstanceList) GetCreationTime() *string {
	return s.CreationTime
}

func (s *ListInstancesResponseBodyInstanceList) GetEnableHiveAccess() *string {
	return s.EnableHiveAccess
}

func (s *ListInstancesResponseBodyInstanceList) GetEndpoints() []*ListInstancesResponseBodyInstanceListEndpoints {
	return s.Endpoints
}

func (s *ListInstancesResponseBodyInstanceList) GetExpirationTime() *string {
	return s.ExpirationTime
}

func (s *ListInstancesResponseBodyInstanceList) GetInstanceChargeType() *string {
	return s.InstanceChargeType
}

func (s *ListInstancesResponseBodyInstanceList) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListInstancesResponseBodyInstanceList) GetInstanceName() *string {
	return s.InstanceName
}

func (s *ListInstancesResponseBodyInstanceList) GetInstanceStatus() *string {
	return s.InstanceStatus
}

func (s *ListInstancesResponseBodyInstanceList) GetInstanceType() *string {
	return s.InstanceType
}

func (s *ListInstancesResponseBodyInstanceList) GetLeaderInstanceId() *string {
	return s.LeaderInstanceId
}

func (s *ListInstancesResponseBodyInstanceList) GetRegionId() *string {
	return s.RegionId
}

func (s *ListInstancesResponseBodyInstanceList) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListInstancesResponseBodyInstanceList) GetStorageType() *string {
	return s.StorageType
}

func (s *ListInstancesResponseBodyInstanceList) GetSuspendReason() *string {
	return s.SuspendReason
}

func (s *ListInstancesResponseBodyInstanceList) GetTags() []*ListInstancesResponseBodyInstanceListTags {
	return s.Tags
}

func (s *ListInstancesResponseBodyInstanceList) GetVersion() *string {
	return s.Version
}

func (s *ListInstancesResponseBodyInstanceList) GetZoneId() *string {
	return s.ZoneId
}

func (s *ListInstancesResponseBodyInstanceList) SetCommodityCode(v string) *ListInstancesResponseBodyInstanceList {
	s.CommodityCode = &v
	return s
}

func (s *ListInstancesResponseBodyInstanceList) SetCreationTime(v string) *ListInstancesResponseBodyInstanceList {
	s.CreationTime = &v
	return s
}

func (s *ListInstancesResponseBodyInstanceList) SetEnableHiveAccess(v string) *ListInstancesResponseBodyInstanceList {
	s.EnableHiveAccess = &v
	return s
}

func (s *ListInstancesResponseBodyInstanceList) SetEndpoints(v []*ListInstancesResponseBodyInstanceListEndpoints) *ListInstancesResponseBodyInstanceList {
	s.Endpoints = v
	return s
}

func (s *ListInstancesResponseBodyInstanceList) SetExpirationTime(v string) *ListInstancesResponseBodyInstanceList {
	s.ExpirationTime = &v
	return s
}

func (s *ListInstancesResponseBodyInstanceList) SetInstanceChargeType(v string) *ListInstancesResponseBodyInstanceList {
	s.InstanceChargeType = &v
	return s
}

func (s *ListInstancesResponseBodyInstanceList) SetInstanceId(v string) *ListInstancesResponseBodyInstanceList {
	s.InstanceId = &v
	return s
}

func (s *ListInstancesResponseBodyInstanceList) SetInstanceName(v string) *ListInstancesResponseBodyInstanceList {
	s.InstanceName = &v
	return s
}

func (s *ListInstancesResponseBodyInstanceList) SetInstanceStatus(v string) *ListInstancesResponseBodyInstanceList {
	s.InstanceStatus = &v
	return s
}

func (s *ListInstancesResponseBodyInstanceList) SetInstanceType(v string) *ListInstancesResponseBodyInstanceList {
	s.InstanceType = &v
	return s
}

func (s *ListInstancesResponseBodyInstanceList) SetLeaderInstanceId(v string) *ListInstancesResponseBodyInstanceList {
	s.LeaderInstanceId = &v
	return s
}

func (s *ListInstancesResponseBodyInstanceList) SetRegionId(v string) *ListInstancesResponseBodyInstanceList {
	s.RegionId = &v
	return s
}

func (s *ListInstancesResponseBodyInstanceList) SetResourceGroupId(v string) *ListInstancesResponseBodyInstanceList {
	s.ResourceGroupId = &v
	return s
}

func (s *ListInstancesResponseBodyInstanceList) SetStorageType(v string) *ListInstancesResponseBodyInstanceList {
	s.StorageType = &v
	return s
}

func (s *ListInstancesResponseBodyInstanceList) SetSuspendReason(v string) *ListInstancesResponseBodyInstanceList {
	s.SuspendReason = &v
	return s
}

func (s *ListInstancesResponseBodyInstanceList) SetTags(v []*ListInstancesResponseBodyInstanceListTags) *ListInstancesResponseBodyInstanceList {
	s.Tags = v
	return s
}

func (s *ListInstancesResponseBodyInstanceList) SetVersion(v string) *ListInstancesResponseBodyInstanceList {
	s.Version = &v
	return s
}

func (s *ListInstancesResponseBodyInstanceList) SetZoneId(v string) *ListInstancesResponseBodyInstanceList {
	s.ZoneId = &v
	return s
}

func (s *ListInstancesResponseBodyInstanceList) Validate() error {
	return dara.Validate(s)
}

type ListInstancesResponseBodyInstanceListEndpoints struct {
	// Indicates whether the endpoint is enabled.
	//
	// Valid values:
	//
	// 	- true
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- false
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// true
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The endpoint.
	//
	// example:
	//
	// hgpostcn-cn-aaab9ad2d8fb-cn-hangzhou-internal.hologres.aliyuncs.com:80
	Endpoint *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	// The network type.
	//
	// Valid values:
	//
	// 	- VPCSingleTunnel
	//
	//     <!-- -->
	//
	//     :
	//
	//     <!-- -->
	//
	//     virtual private cloud (VPC)
	//
	//     <!-- -->
	//
	//     .
	//
	// 	- Intranet
	//
	//     <!-- -->
	//
	//     :
	//
	//     <!-- -->
	//
	//     internal network
	//
	//     <!-- -->
	//
	// 	- VPCAnyTunnel
	//
	//     <!-- -->
	//
	//     : This value is not supported by new instances
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     .
	//
	// 	- Internet
	//
	//     <!-- -->
	//
	//     :
	//
	//     <!-- -->
	//
	//     Internet
	//
	//     <!-- -->
	//
	//     .
	//
	// example:
	//
	// Internet
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The vSwitch ID.
	//
	// example:
	//
	// vsw-wz9oap28raidjevhuszg4
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The VPC ID.
	//
	// example:
	//
	// vpc-uf6mrahzyu7uorlqqpz5f
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The ID of the VPC to which the instance belongs.
	//
	// example:
	//
	// hgpostcn-cn-wwo3665tx004-frontend-st
	VpcInstanceId *string `json:"VpcInstanceId,omitempty" xml:"VpcInstanceId,omitempty"`
}

func (s ListInstancesResponseBodyInstanceListEndpoints) String() string {
	return dara.Prettify(s)
}

func (s ListInstancesResponseBodyInstanceListEndpoints) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBodyInstanceListEndpoints) GetEnabled() *bool {
	return s.Enabled
}

func (s *ListInstancesResponseBodyInstanceListEndpoints) GetEndpoint() *string {
	return s.Endpoint
}

func (s *ListInstancesResponseBodyInstanceListEndpoints) GetType() *string {
	return s.Type
}

func (s *ListInstancesResponseBodyInstanceListEndpoints) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *ListInstancesResponseBodyInstanceListEndpoints) GetVpcId() *string {
	return s.VpcId
}

func (s *ListInstancesResponseBodyInstanceListEndpoints) GetVpcInstanceId() *string {
	return s.VpcInstanceId
}

func (s *ListInstancesResponseBodyInstanceListEndpoints) SetEnabled(v bool) *ListInstancesResponseBodyInstanceListEndpoints {
	s.Enabled = &v
	return s
}

func (s *ListInstancesResponseBodyInstanceListEndpoints) SetEndpoint(v string) *ListInstancesResponseBodyInstanceListEndpoints {
	s.Endpoint = &v
	return s
}

func (s *ListInstancesResponseBodyInstanceListEndpoints) SetType(v string) *ListInstancesResponseBodyInstanceListEndpoints {
	s.Type = &v
	return s
}

func (s *ListInstancesResponseBodyInstanceListEndpoints) SetVSwitchId(v string) *ListInstancesResponseBodyInstanceListEndpoints {
	s.VSwitchId = &v
	return s
}

func (s *ListInstancesResponseBodyInstanceListEndpoints) SetVpcId(v string) *ListInstancesResponseBodyInstanceListEndpoints {
	s.VpcId = &v
	return s
}

func (s *ListInstancesResponseBodyInstanceListEndpoints) SetVpcInstanceId(v string) *ListInstancesResponseBodyInstanceListEndpoints {
	s.VpcInstanceId = &v
	return s
}

func (s *ListInstancesResponseBodyInstanceListEndpoints) Validate() error {
	return dara.Validate(s)
}

type ListInstancesResponseBodyInstanceListTags struct {
	// The tag key.
	//
	// example:
	//
	// tag
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// value
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListInstancesResponseBodyInstanceListTags) String() string {
	return dara.Prettify(s)
}

func (s ListInstancesResponseBodyInstanceListTags) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBodyInstanceListTags) GetKey() *string {
	return s.Key
}

func (s *ListInstancesResponseBodyInstanceListTags) GetValue() *string {
	return s.Value
}

func (s *ListInstancesResponseBodyInstanceListTags) SetKey(v string) *ListInstancesResponseBodyInstanceListTags {
	s.Key = &v
	return s
}

func (s *ListInstancesResponseBodyInstanceListTags) SetValue(v string) *ListInstancesResponseBodyInstanceListTags {
	s.Value = &v
	return s
}

func (s *ListInstancesResponseBodyInstanceListTags) Validate() error {
	return dara.Validate(s)
}
