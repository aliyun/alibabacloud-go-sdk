// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAliasName(v string) *GetInstanceResponseBody
	GetAliasName() *string
	SetCreateTime(v string) *GetInstanceResponseBody
	GetCreateTime() *string
	SetElasticVCUUpperLimit(v float32) *GetInstanceResponseBody
	GetElasticVCUUpperLimit() *float32
	SetInstanceDescription(v string) *GetInstanceResponseBody
	GetInstanceDescription() *string
	SetInstanceName(v string) *GetInstanceResponseBody
	GetInstanceName() *string
	SetInstanceSpecification(v string) *GetInstanceResponseBody
	GetInstanceSpecification() *string
	SetInstanceStatus(v string) *GetInstanceResponseBody
	GetInstanceStatus() *string
	SetIsMultiAZ(v bool) *GetInstanceResponseBody
	GetIsMultiAZ() *bool
	SetNetwork(v string) *GetInstanceResponseBody
	GetNetwork() *string
	SetNetworkSourceACL(v []*string) *GetInstanceResponseBody
	GetNetworkSourceACL() []*string
	SetNetworkTypeACL(v []*string) *GetInstanceResponseBody
	GetNetworkTypeACL() []*string
	SetPaymentType(v string) *GetInstanceResponseBody
	GetPaymentType() *string
	SetPolicy(v string) *GetInstanceResponseBody
	GetPolicy() *string
	SetPolicyVersion(v int64) *GetInstanceResponseBody
	GetPolicyVersion() *int64
	SetRegionId(v string) *GetInstanceResponseBody
	GetRegionId() *string
	SetRequestId(v string) *GetInstanceResponseBody
	GetRequestId() *string
	SetResourceGroupId(v string) *GetInstanceResponseBody
	GetResourceGroupId() *string
	SetSPInstanceId(v string) *GetInstanceResponseBody
	GetSPInstanceId() *string
	SetStorageType(v string) *GetInstanceResponseBody
	GetStorageType() *string
	SetTableQuota(v int32) *GetInstanceResponseBody
	GetTableQuota() *int32
	SetTags(v []*GetInstanceResponseBodyTags) *GetInstanceResponseBody
	GetTags() []*GetInstanceResponseBodyTags
	SetUserId(v string) *GetInstanceResponseBody
	GetUserId() *string
	SetVCUQuota(v int32) *GetInstanceResponseBody
	GetVCUQuota() *int32
}

type GetInstanceResponseBody struct {
	// The instance alias.
	//
	// example:
	//
	// instance-test
	AliasName *string `json:"AliasName,omitempty" xml:"AliasName,omitempty"`
	// The time when the instance was created.
	//
	// example:
	//
	// 2019-12-23T07:24:33Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The upper limit for the VCUs of the instance.
	//
	// example:
	//
	// 6
	ElasticVCUUpperLimit *float32 `json:"ElasticVCUUpperLimit,omitempty" xml:"ElasticVCUUpperLimit,omitempty"`
	// The description of the instance.
	//
	// example:
	//
	// Description of the test instance.
	InstanceDescription *string `json:"InstanceDescription,omitempty" xml:"InstanceDescription,omitempty"`
	// The name of the instance.
	//
	// example:
	//
	// instance-test
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The type of the instance.
	//
	// 	- SSD: high-performance instance
	//
	// 	- HYBRID: capacity instance
	//
	// example:
	//
	// SSD
	InstanceSpecification *string `json:"InstanceSpecification,omitempty" xml:"InstanceSpecification,omitempty"`
	// The status of the instance.
	//
	// 	- normal: The instance works as expected.
	//
	// 	- forbidden: The instance is disabled.
	//
	// 	- Deleting: The instance is being deleted.
	//
	// example:
	//
	// normal
	InstanceStatus *string `json:"InstanceStatus,omitempty" xml:"InstanceStatus,omitempty"`
	// Indicates whether zone-redundant storage (ZRS) is enabled for the instance.
	//
	// 	- true: ZRS is enabled for the instance.
	//
	// 	- false: Locally redundant storage (LRS) is enabled for the instance.
	//
	// example:
	//
	// true
	IsMultiAZ *bool `json:"IsMultiAZ,omitempty" xml:"IsMultiAZ,omitempty"`
	// The network type of the instance. Valid values:
	//
	// 	- VPC: The instance can be accessed only over the bound virtual private clouds (VPCs).
	//
	// 	- VPC_CONSOLE: The instance can be accessed from the Tablestore console or over the bound VPCs.
	//
	// 	- NORMAL: The instance can be accessed from networks of the custom types.
	//
	// example:
	//
	// VPC
	Network *string `json:"Network,omitempty" xml:"Network,omitempty"`
	// The sources of the network from which access is allowed. Valid value:
	//
	// TRUST_PROXY: console
	NetworkSourceACL []*string `json:"NetworkSourceACL,omitempty" xml:"NetworkSourceACL,omitempty" type:"Repeated"`
	// The types of the network from which access is allowed. Valid values:
	//
	// 	- CLASSIC: the classic network
	//
	// 	- INTERNET: the Internet
	//
	// 	- VPC: VPCs
	NetworkTypeACL []*string `json:"NetworkTypeACL,omitempty" xml:"NetworkTypeACL,omitempty" type:"Repeated"`
	// The billing method. Valid values:
	//
	// 	- Subscription: subscription
	//
	// 	- PayAsYouGo: pay-as-you-go
	//
	// example:
	//
	// Subscription
	PaymentType *string `json:"PaymentType,omitempty" xml:"PaymentType,omitempty"`
	// The instance policy.
	//
	// example:
	//
	// {
	//
	//     "Version": "1",
	//
	//     "Statement": [
	//
	//         {
	//
	//             "Action": [
	//
	//                 "ots:*"
	//
	//             ],
	//
	//             "Resource": [
	//
	//                 "acs:ots:*:13791xxxxxxxxxxx:instance/myinstance*"
	//
	//             ],
	//
	//             "Principal": [
	//
	//                 "*"
	//
	//             ],
	//
	//             "Effect": "Allow",
	//
	//             "Condition": {
	//
	//                 "StringEquals": {
	//
	//                     "ots:TLSVersion": [
	//
	//                         "1.2"
	//
	//                     ]
	//
	//                 },
	//
	//                 "IpAddress": {
	//
	//                     "acs:SourceIp": [
	//
	//                         "192.168.0.1",
	//
	//                         "198.51.100.1"
	//
	//                     ]
	//
	//                 }
	//
	//             }
	//
	//         }
	//
	//     ]
	//
	// }
	Policy *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// The version of the instance policy.
	//
	// example:
	//
	// 1
	PolicyVersion *int64 `json:"PolicyVersion,omitempty" xml:"PolicyVersion,omitempty"`
	// The region ID of the instance.
	//
	// example:
	//
	// cn-chengdu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The request ID, which can be used to troubleshoot issues.
	//
	// example:
	//
	// 757E172A-F94B-5E78-8A23-D9068E42F2E9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the resource group to which the instance belongs.
	//
	// example:
	//
	// rg-acfmxh4em5jncda
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// ots_standard_public_cn-9lb34u7u5001
	SPInstanceId *string `json:"SPInstanceId,omitempty" xml:"SPInstanceId,omitempty"`
	// The storage type.
	//
	// 	- SSD: high-performance
	//
	// 	- HYBRID: capacity
	//
	// example:
	//
	// HYBRID
	StorageType *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	// The total number of tables in the instance.
	//
	// example:
	//
	// 100
	TableQuota *int32 `json:"TableQuota,omitempty" xml:"TableQuota,omitempty"`
	// The tags of the instance.
	Tags []*GetInstanceResponseBodyTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The user ID.
	//
	// example:
	//
	// 16542312566
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The VCU quota.
	//
	// example:
	//
	// 20
	VCUQuota *int32 `json:"VCUQuota,omitempty" xml:"VCUQuota,omitempty"`
}

func (s GetInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *GetInstanceResponseBody) GetAliasName() *string {
	return s.AliasName
}

func (s *GetInstanceResponseBody) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetInstanceResponseBody) GetElasticVCUUpperLimit() *float32 {
	return s.ElasticVCUUpperLimit
}

func (s *GetInstanceResponseBody) GetInstanceDescription() *string {
	return s.InstanceDescription
}

func (s *GetInstanceResponseBody) GetInstanceName() *string {
	return s.InstanceName
}

func (s *GetInstanceResponseBody) GetInstanceSpecification() *string {
	return s.InstanceSpecification
}

func (s *GetInstanceResponseBody) GetInstanceStatus() *string {
	return s.InstanceStatus
}

func (s *GetInstanceResponseBody) GetIsMultiAZ() *bool {
	return s.IsMultiAZ
}

func (s *GetInstanceResponseBody) GetNetwork() *string {
	return s.Network
}

func (s *GetInstanceResponseBody) GetNetworkSourceACL() []*string {
	return s.NetworkSourceACL
}

func (s *GetInstanceResponseBody) GetNetworkTypeACL() []*string {
	return s.NetworkTypeACL
}

func (s *GetInstanceResponseBody) GetPaymentType() *string {
	return s.PaymentType
}

func (s *GetInstanceResponseBody) GetPolicy() *string {
	return s.Policy
}

func (s *GetInstanceResponseBody) GetPolicyVersion() *int64 {
	return s.PolicyVersion
}

func (s *GetInstanceResponseBody) GetRegionId() *string {
	return s.RegionId
}

func (s *GetInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetInstanceResponseBody) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *GetInstanceResponseBody) GetSPInstanceId() *string {
	return s.SPInstanceId
}

func (s *GetInstanceResponseBody) GetStorageType() *string {
	return s.StorageType
}

func (s *GetInstanceResponseBody) GetTableQuota() *int32 {
	return s.TableQuota
}

func (s *GetInstanceResponseBody) GetTags() []*GetInstanceResponseBodyTags {
	return s.Tags
}

func (s *GetInstanceResponseBody) GetUserId() *string {
	return s.UserId
}

func (s *GetInstanceResponseBody) GetVCUQuota() *int32 {
	return s.VCUQuota
}

func (s *GetInstanceResponseBody) SetAliasName(v string) *GetInstanceResponseBody {
	s.AliasName = &v
	return s
}

func (s *GetInstanceResponseBody) SetCreateTime(v string) *GetInstanceResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetInstanceResponseBody) SetElasticVCUUpperLimit(v float32) *GetInstanceResponseBody {
	s.ElasticVCUUpperLimit = &v
	return s
}

func (s *GetInstanceResponseBody) SetInstanceDescription(v string) *GetInstanceResponseBody {
	s.InstanceDescription = &v
	return s
}

func (s *GetInstanceResponseBody) SetInstanceName(v string) *GetInstanceResponseBody {
	s.InstanceName = &v
	return s
}

func (s *GetInstanceResponseBody) SetInstanceSpecification(v string) *GetInstanceResponseBody {
	s.InstanceSpecification = &v
	return s
}

func (s *GetInstanceResponseBody) SetInstanceStatus(v string) *GetInstanceResponseBody {
	s.InstanceStatus = &v
	return s
}

func (s *GetInstanceResponseBody) SetIsMultiAZ(v bool) *GetInstanceResponseBody {
	s.IsMultiAZ = &v
	return s
}

func (s *GetInstanceResponseBody) SetNetwork(v string) *GetInstanceResponseBody {
	s.Network = &v
	return s
}

func (s *GetInstanceResponseBody) SetNetworkSourceACL(v []*string) *GetInstanceResponseBody {
	s.NetworkSourceACL = v
	return s
}

func (s *GetInstanceResponseBody) SetNetworkTypeACL(v []*string) *GetInstanceResponseBody {
	s.NetworkTypeACL = v
	return s
}

func (s *GetInstanceResponseBody) SetPaymentType(v string) *GetInstanceResponseBody {
	s.PaymentType = &v
	return s
}

func (s *GetInstanceResponseBody) SetPolicy(v string) *GetInstanceResponseBody {
	s.Policy = &v
	return s
}

func (s *GetInstanceResponseBody) SetPolicyVersion(v int64) *GetInstanceResponseBody {
	s.PolicyVersion = &v
	return s
}

func (s *GetInstanceResponseBody) SetRegionId(v string) *GetInstanceResponseBody {
	s.RegionId = &v
	return s
}

func (s *GetInstanceResponseBody) SetRequestId(v string) *GetInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetInstanceResponseBody) SetResourceGroupId(v string) *GetInstanceResponseBody {
	s.ResourceGroupId = &v
	return s
}

func (s *GetInstanceResponseBody) SetSPInstanceId(v string) *GetInstanceResponseBody {
	s.SPInstanceId = &v
	return s
}

func (s *GetInstanceResponseBody) SetStorageType(v string) *GetInstanceResponseBody {
	s.StorageType = &v
	return s
}

func (s *GetInstanceResponseBody) SetTableQuota(v int32) *GetInstanceResponseBody {
	s.TableQuota = &v
	return s
}

func (s *GetInstanceResponseBody) SetTags(v []*GetInstanceResponseBodyTags) *GetInstanceResponseBody {
	s.Tags = v
	return s
}

func (s *GetInstanceResponseBody) SetUserId(v string) *GetInstanceResponseBody {
	s.UserId = &v
	return s
}

func (s *GetInstanceResponseBody) SetVCUQuota(v int32) *GetInstanceResponseBody {
	s.VCUQuota = &v
	return s
}

func (s *GetInstanceResponseBody) Validate() error {
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

type GetInstanceResponseBodyTags struct {
	// The tag key.
	//
	// example:
	//
	// tag
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// (Deprecated) The tag key.
	//
	// example:
	//
	// keyTestA
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// (Deprecated) The tag value.
	//
	// example:
	//
	// 00004a20240452b0
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
	// The tag value.
	//
	// example:
	//
	// 333
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetInstanceResponseBodyTags) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceResponseBodyTags) GoString() string {
	return s.String()
}

func (s *GetInstanceResponseBodyTags) GetKey() *string {
	return s.Key
}

func (s *GetInstanceResponseBodyTags) GetTagKey() *string {
	return s.TagKey
}

func (s *GetInstanceResponseBodyTags) GetTagValue() *string {
	return s.TagValue
}

func (s *GetInstanceResponseBodyTags) GetValue() *string {
	return s.Value
}

func (s *GetInstanceResponseBodyTags) SetKey(v string) *GetInstanceResponseBodyTags {
	s.Key = &v
	return s
}

func (s *GetInstanceResponseBodyTags) SetTagKey(v string) *GetInstanceResponseBodyTags {
	s.TagKey = &v
	return s
}

func (s *GetInstanceResponseBodyTags) SetTagValue(v string) *GetInstanceResponseBodyTags {
	s.TagValue = &v
	return s
}

func (s *GetInstanceResponseBodyTags) SetValue(v string) *GetInstanceResponseBodyTags {
	s.Value = &v
	return s
}

func (s *GetInstanceResponseBodyTags) Validate() error {
	return dara.Validate(s)
}
