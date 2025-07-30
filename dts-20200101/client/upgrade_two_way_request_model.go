// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeTwoWayRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceClass(v string) *UpgradeTwoWayRequest
	GetInstanceClass() *string
	SetInstanceId(v string) *UpgradeTwoWayRequest
	GetInstanceId() *string
	SetRegionId(v string) *UpgradeTwoWayRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *UpgradeTwoWayRequest
	GetResourceGroupId() *string
}

type UpgradeTwoWayRequest struct {
	// The instance class of the two-way synchronization task. Valid values: **large**, **medium**, **micro**, and **small**.
	//
	// >  For more information, see [Specifications of data synchronization instances](https://help.aliyun.com/document_detail/26605.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// large
	InstanceClass *string `json:"InstanceClass,omitempty" xml:"InstanceClass,omitempty"`
	// The ID of the data synchronization instance. You can call the [DescribeDtsJobs](https://help.aliyun.com/document_detail/209702.html) operation to query the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// dtsh77p49x4k28****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the DTS instance. For more information, see [List of supported regions](https://help.aliyun.com/document_detail/141033.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Resource group ID.
	//
	// example:
	//
	// rg-acfmzawhxxc****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s UpgradeTwoWayRequest) String() string {
	return dara.Prettify(s)
}

func (s UpgradeTwoWayRequest) GoString() string {
	return s.String()
}

func (s *UpgradeTwoWayRequest) GetInstanceClass() *string {
	return s.InstanceClass
}

func (s *UpgradeTwoWayRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpgradeTwoWayRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpgradeTwoWayRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *UpgradeTwoWayRequest) SetInstanceClass(v string) *UpgradeTwoWayRequest {
	s.InstanceClass = &v
	return s
}

func (s *UpgradeTwoWayRequest) SetInstanceId(v string) *UpgradeTwoWayRequest {
	s.InstanceId = &v
	return s
}

func (s *UpgradeTwoWayRequest) SetRegionId(v string) *UpgradeTwoWayRequest {
	s.RegionId = &v
	return s
}

func (s *UpgradeTwoWayRequest) SetResourceGroupId(v string) *UpgradeTwoWayRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *UpgradeTwoWayRequest) Validate() error {
	return dara.Validate(s)
}
