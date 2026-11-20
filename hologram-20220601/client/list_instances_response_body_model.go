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
	// The error code.
	//
	// example:
	//
	// null
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// null
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *string `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The list of instances.
	InstanceList []*ListInstancesResponseBodyInstanceList `json:"InstanceList,omitempty" xml:"InstanceList,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// D1303CD4-AA70-5998-8025-F55B22C50840
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the API call is successful.
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
	if s.InstanceList != nil {
		for _, item := range s.InstanceList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListInstancesResponseBodyInstanceList struct {
	// The commodity code, which is consistent with the product detail code in the User Center.
	//
	// example:
	//
	// hologram_postpay_public_cn
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	// The creation time. The time is displayed in UTC and formatted according to ISO 8601.
	//
	// example:
	//
	// 2022-12-16T02:24:05Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// Specifies whether data lakehouse acceleration is enabled.
	//
	// example:
	//
	// true
	EnableHiveAccess *string `json:"EnableHiveAccess,omitempty" xml:"EnableHiveAccess,omitempty"`
	// Specifies whether SSL is enabled.
	//
	// example:
	//
	// true
	EnableSSL *string `json:"EnableSSL,omitempty" xml:"EnableSSL,omitempty"`
	// The list of endpoints.
	Endpoints []*ListInstancesResponseBodyInstanceListEndpoints `json:"Endpoints,omitempty" xml:"Endpoints,omitempty" type:"Repeated"`
	// The expiration time. The time is displayed in UTC and formatted according to ISO 8601.
	//
	// example:
	//
	// 2023-05-04T16:00:00.000Z
	ExpirationTime *string `json:"ExpirationTime,omitempty" xml:"ExpirationTime,omitempty"`
	// The billing method of the instance.
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
	// The instance name.
	//
	// example:
	//
	// test_instance
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The instance status.
	//
	// example:
	//
	// Running
	InstanceStatus *string `json:"InstanceStatus,omitempty" xml:"InstanceStatus,omitempty"`
	// The instance type.
	//
	// example:
	//
	// Standard
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The leader instance ID.
	//
	// example:
	//
	// hgprecn-cn-2r42sqvxm006
	LeaderInstanceId *string `json:"LeaderInstanceId,omitempty" xml:"LeaderInstanceId,omitempty"`
	// The region ID.
	//
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
	// The storage type. Valid values:
	//
	// 	- redundant: 3AZ
	//
	// 	- local: 1AZ
	//
	// example:
	//
	// redundant
	StorageType *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	// The reason for suspension.
	//
	// example:
	//
	// Manual
	SuspendReason *string `json:"SuspendReason,omitempty" xml:"SuspendReason,omitempty"`
	// The list of tags.
	Tags []*ListInstancesResponseBodyInstanceListTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The instance version.
	//
	// example:
	//
	// 1.3.37
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
	// The zone ID.
	//
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

func (s *ListInstancesResponseBodyInstanceList) GetEnableSSL() *string {
	return s.EnableSSL
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

func (s *ListInstancesResponseBodyInstanceList) SetEnableSSL(v string) *ListInstancesResponseBodyInstanceList {
	s.EnableSSL = &v
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
	if s.Endpoints != nil {
		for _, item := range s.Endpoints {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListInstancesResponseBodyInstanceListEndpoints struct {
	// Specifies whether the endpoint is enabled.
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
	// The ID of the VPC-connected instance.
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
