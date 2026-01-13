// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstances(v []*ListInstancesResponseBodyInstances) *ListInstancesResponseBody
	GetInstances() []*ListInstancesResponseBodyInstances
	SetNextToken(v string) *ListInstancesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListInstancesResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListInstancesResponseBody
	GetTotalCount() *int64
}

type ListInstancesResponseBody struct {
	// The details about the instances.
	Instances []*ListInstancesResponseBodyInstances `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Repeated"`
	// The token that determines the start position of the next query. If this parameter is empty, all instances that you want to query are returned.
	//
	// example:
	//
	// CAESCG15aC1xxxxx
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID, which can be used to troubleshoot issues.
	//
	// example:
	//
	// E734979F-5A44-5993-9CE5-C23103576923
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of instances returned.
	//
	// example:
	//
	// 2
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBody) GetInstances() []*ListInstancesResponseBodyInstances {
	return s.Instances
}

func (s *ListInstancesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListInstancesResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListInstancesResponseBody) SetInstances(v []*ListInstancesResponseBodyInstances) *ListInstancesResponseBody {
	s.Instances = v
	return s
}

func (s *ListInstancesResponseBody) SetNextToken(v string) *ListInstancesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListInstancesResponseBody) SetRequestId(v string) *ListInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInstancesResponseBody) SetTotalCount(v int64) *ListInstancesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListInstancesResponseBody) Validate() error {
	if s.Instances != nil {
		for _, item := range s.Instances {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListInstancesResponseBodyInstances struct {
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
	// 2019-04-07T09:19:21Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The instance description.
	//
	// example:
	//
	// Description of the test instance.
	InstanceDescription *string `json:"InstanceDescription,omitempty" xml:"InstanceDescription,omitempty"`
	// The name of the instance, which is used to uniquely identify the instance.
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
	// HYBRID
	InstanceSpecification *string `json:"InstanceSpecification,omitempty" xml:"InstanceSpecification,omitempty"`
	// The status of the instance.
	//
	// 	- normal: The instance runs as expected.
	//
	// 	- forbidden: The instance is disabled.
	//
	// 	- Deleting: The instance is being released.
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
	// The billing method.
	//
	// 	- Subscription: subscription
	//
	// 	- PayAsYouGo: pay as you go
	//
	// example:
	//
	// Subscription
	PaymentType *string `json:"PaymentType,omitempty" xml:"PaymentType,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-huhehaote
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmxh4em5jnbcd
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// ots_standard_public_cn-g4t3igqjj002
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
	// The user ID.
	//
	// example:
	//
	// 13542356466
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The VCU quota.
	//
	// example:
	//
	// 3
	VCUQuota *int32 `json:"VCUQuota,omitempty" xml:"VCUQuota,omitempty"`
}

func (s ListInstancesResponseBodyInstances) String() string {
	return dara.Prettify(s)
}

func (s ListInstancesResponseBodyInstances) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBodyInstances) GetAliasName() *string {
	return s.AliasName
}

func (s *ListInstancesResponseBodyInstances) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListInstancesResponseBodyInstances) GetInstanceDescription() *string {
	return s.InstanceDescription
}

func (s *ListInstancesResponseBodyInstances) GetInstanceName() *string {
	return s.InstanceName
}

func (s *ListInstancesResponseBodyInstances) GetInstanceSpecification() *string {
	return s.InstanceSpecification
}

func (s *ListInstancesResponseBodyInstances) GetInstanceStatus() *string {
	return s.InstanceStatus
}

func (s *ListInstancesResponseBodyInstances) GetIsMultiAZ() *bool {
	return s.IsMultiAZ
}

func (s *ListInstancesResponseBodyInstances) GetPaymentType() *string {
	return s.PaymentType
}

func (s *ListInstancesResponseBodyInstances) GetRegionId() *string {
	return s.RegionId
}

func (s *ListInstancesResponseBodyInstances) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListInstancesResponseBodyInstances) GetSPInstanceId() *string {
	return s.SPInstanceId
}

func (s *ListInstancesResponseBodyInstances) GetStorageType() *string {
	return s.StorageType
}

func (s *ListInstancesResponseBodyInstances) GetUserId() *string {
	return s.UserId
}

func (s *ListInstancesResponseBodyInstances) GetVCUQuota() *int32 {
	return s.VCUQuota
}

func (s *ListInstancesResponseBodyInstances) SetAliasName(v string) *ListInstancesResponseBodyInstances {
	s.AliasName = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetCreateTime(v string) *ListInstancesResponseBodyInstances {
	s.CreateTime = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetInstanceDescription(v string) *ListInstancesResponseBodyInstances {
	s.InstanceDescription = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetInstanceName(v string) *ListInstancesResponseBodyInstances {
	s.InstanceName = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetInstanceSpecification(v string) *ListInstancesResponseBodyInstances {
	s.InstanceSpecification = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetInstanceStatus(v string) *ListInstancesResponseBodyInstances {
	s.InstanceStatus = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetIsMultiAZ(v bool) *ListInstancesResponseBodyInstances {
	s.IsMultiAZ = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetPaymentType(v string) *ListInstancesResponseBodyInstances {
	s.PaymentType = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetRegionId(v string) *ListInstancesResponseBodyInstances {
	s.RegionId = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetResourceGroupId(v string) *ListInstancesResponseBodyInstances {
	s.ResourceGroupId = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetSPInstanceId(v string) *ListInstancesResponseBodyInstances {
	s.SPInstanceId = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetStorageType(v string) *ListInstancesResponseBodyInstances {
	s.StorageType = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetUserId(v string) *ListInstancesResponseBodyInstances {
	s.UserId = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetVCUQuota(v int32) *ListInstancesResponseBodyInstances {
	s.VCUQuota = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) Validate() error {
	return dara.Validate(s)
}
