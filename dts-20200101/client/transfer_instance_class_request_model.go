// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTransferInstanceClassRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDtsJobId(v string) *TransferInstanceClassRequest
	GetDtsJobId() *string
	SetInstanceClass(v string) *TransferInstanceClassRequest
	GetInstanceClass() *string
	SetOrderType(v string) *TransferInstanceClassRequest
	GetOrderType() *string
	SetRegionId(v string) *TransferInstanceClassRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *TransferInstanceClassRequest
	GetResourceGroupId() *string
}

type TransferInstanceClassRequest struct {
	// The ID of the data migration or data synchronization task. You can call the [DescribeDtsJobs](https://help.aliyun.com/document_detail/209702.html) operation to query the task ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// r4yr723m199****
	DtsJobId *string `json:"DtsJobId,omitempty" xml:"DtsJobId,omitempty"`
	// The new instance class of the DTS instance. You can call the [DescribeDtsJobDetail](https://help.aliyun.com/document_detail/208925.html) operation to query the original instance class of the DTS instance.
	//
	// 	- DTS supports the following instance classes for a data migration instance: **xxlarge**, **xlarge**, **large**, **medium**, and **small**.
	//
	// 	- DTS supports the following instance classes for a data synchronization instance: **large**, **medium**, **small**, and **micro**.
	//
	// > For more information about the test performance of each instance class, see [Specifications of data migration instances](https://help.aliyun.com/document_detail/26606.html) and [Specifications of data synchronization channels](https://help.aliyun.com/document_detail/26605.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// large
	InstanceClass *string `json:"InstanceClass,omitempty" xml:"InstanceClass,omitempty"`
	// Set the value to **UPGRADE**.
	//
	// This parameter is required.
	//
	// example:
	//
	// UPGRADE
	OrderType *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	// The ID of the region in which the DTS instance resides. For more information, see [Supported regions](https://help.aliyun.com/document_detail/141033.html).
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

func (s TransferInstanceClassRequest) String() string {
	return dara.Prettify(s)
}

func (s TransferInstanceClassRequest) GoString() string {
	return s.String()
}

func (s *TransferInstanceClassRequest) GetDtsJobId() *string {
	return s.DtsJobId
}

func (s *TransferInstanceClassRequest) GetInstanceClass() *string {
	return s.InstanceClass
}

func (s *TransferInstanceClassRequest) GetOrderType() *string {
	return s.OrderType
}

func (s *TransferInstanceClassRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *TransferInstanceClassRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *TransferInstanceClassRequest) SetDtsJobId(v string) *TransferInstanceClassRequest {
	s.DtsJobId = &v
	return s
}

func (s *TransferInstanceClassRequest) SetInstanceClass(v string) *TransferInstanceClassRequest {
	s.InstanceClass = &v
	return s
}

func (s *TransferInstanceClassRequest) SetOrderType(v string) *TransferInstanceClassRequest {
	s.OrderType = &v
	return s
}

func (s *TransferInstanceClassRequest) SetRegionId(v string) *TransferInstanceClassRequest {
	s.RegionId = &v
	return s
}

func (s *TransferInstanceClassRequest) SetResourceGroupId(v string) *TransferInstanceClassRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *TransferInstanceClassRequest) Validate() error {
	return dara.Validate(s)
}
