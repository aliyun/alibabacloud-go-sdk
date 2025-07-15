// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHighDefinitionMonitorLogAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeHighDefinitionMonitorLogAttributeRequest
	GetInstanceId() *string
	SetInstanceType(v string) *DescribeHighDefinitionMonitorLogAttributeRequest
	GetInstanceType() *string
	SetOwnerAccount(v string) *DescribeHighDefinitionMonitorLogAttributeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeHighDefinitionMonitorLogAttributeRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeHighDefinitionMonitorLogAttributeRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeHighDefinitionMonitorLogAttributeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeHighDefinitionMonitorLogAttributeRequest
	GetResourceOwnerId() *int64
}

type DescribeHighDefinitionMonitorLogAttributeRequest struct {
	// The ID of the instance whose fine-grained monitoring configurations you want to query.
	//
	// This parameter is required.
	//
	// example:
	//
	// eip-wz9fi6qboho9fwgx7****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The type of the instance. Set the value to **EIP**.
	//
	// example:
	//
	// EIP
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the instance.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeHighDefinitionMonitorLogAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeHighDefinitionMonitorLogAttributeRequest) GoString() string {
	return s.String()
}

func (s *DescribeHighDefinitionMonitorLogAttributeRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeHighDefinitionMonitorLogAttributeRequest) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DescribeHighDefinitionMonitorLogAttributeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeHighDefinitionMonitorLogAttributeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeHighDefinitionMonitorLogAttributeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeHighDefinitionMonitorLogAttributeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeHighDefinitionMonitorLogAttributeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeHighDefinitionMonitorLogAttributeRequest) SetInstanceId(v string) *DescribeHighDefinitionMonitorLogAttributeRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeHighDefinitionMonitorLogAttributeRequest) SetInstanceType(v string) *DescribeHighDefinitionMonitorLogAttributeRequest {
	s.InstanceType = &v
	return s
}

func (s *DescribeHighDefinitionMonitorLogAttributeRequest) SetOwnerAccount(v string) *DescribeHighDefinitionMonitorLogAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeHighDefinitionMonitorLogAttributeRequest) SetOwnerId(v int64) *DescribeHighDefinitionMonitorLogAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeHighDefinitionMonitorLogAttributeRequest) SetRegionId(v string) *DescribeHighDefinitionMonitorLogAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeHighDefinitionMonitorLogAttributeRequest) SetResourceOwnerAccount(v string) *DescribeHighDefinitionMonitorLogAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeHighDefinitionMonitorLogAttributeRequest) SetResourceOwnerId(v int64) *DescribeHighDefinitionMonitorLogAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeHighDefinitionMonitorLogAttributeRequest) Validate() error {
	return dara.Validate(s)
}
