// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCenAttachedChildInstanceAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCenId(v string) *DescribeCenAttachedChildInstanceAttributeRequest
	GetCenId() *string
	SetChildInstanceId(v string) *DescribeCenAttachedChildInstanceAttributeRequest
	GetChildInstanceId() *string
	SetChildInstanceRegionId(v string) *DescribeCenAttachedChildInstanceAttributeRequest
	GetChildInstanceRegionId() *string
	SetChildInstanceType(v string) *DescribeCenAttachedChildInstanceAttributeRequest
	GetChildInstanceType() *string
	SetOwnerAccount(v string) *DescribeCenAttachedChildInstanceAttributeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeCenAttachedChildInstanceAttributeRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DescribeCenAttachedChildInstanceAttributeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeCenAttachedChildInstanceAttributeRequest
	GetResourceOwnerId() *int64
}

type DescribeCenAttachedChildInstanceAttributeRequest struct {
	// The ID of the CEN instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cen-5mv960yjhja0dh****
	CenId *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	// The ID of the network instance that is attached to the CEN instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpc-2zebdboka7d7t37vo****
	ChildInstanceId *string `json:"ChildInstanceId,omitempty" xml:"ChildInstanceId,omitempty"`
	// The region ID of the network instance.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	ChildInstanceRegionId *string `json:"ChildInstanceRegionId,omitempty" xml:"ChildInstanceRegionId,omitempty"`
	// The type of the network instance. Valid values:
	//
	// 	- **VPC**: virtual private cloud (VPC)
	//
	// 	- **VBR**: virtual border router (VBR)
	//
	// 	- **CCN**: Cloud Connect Network (CCN) instance
	//
	// This parameter is required.
	//
	// example:
	//
	// VPC
	ChildInstanceType    *string `json:"ChildInstanceType,omitempty" xml:"ChildInstanceType,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeCenAttachedChildInstanceAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCenAttachedChildInstanceAttributeRequest) GoString() string {
	return s.String()
}

func (s *DescribeCenAttachedChildInstanceAttributeRequest) GetCenId() *string {
	return s.CenId
}

func (s *DescribeCenAttachedChildInstanceAttributeRequest) GetChildInstanceId() *string {
	return s.ChildInstanceId
}

func (s *DescribeCenAttachedChildInstanceAttributeRequest) GetChildInstanceRegionId() *string {
	return s.ChildInstanceRegionId
}

func (s *DescribeCenAttachedChildInstanceAttributeRequest) GetChildInstanceType() *string {
	return s.ChildInstanceType
}

func (s *DescribeCenAttachedChildInstanceAttributeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeCenAttachedChildInstanceAttributeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeCenAttachedChildInstanceAttributeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeCenAttachedChildInstanceAttributeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeCenAttachedChildInstanceAttributeRequest) SetCenId(v string) *DescribeCenAttachedChildInstanceAttributeRequest {
	s.CenId = &v
	return s
}

func (s *DescribeCenAttachedChildInstanceAttributeRequest) SetChildInstanceId(v string) *DescribeCenAttachedChildInstanceAttributeRequest {
	s.ChildInstanceId = &v
	return s
}

func (s *DescribeCenAttachedChildInstanceAttributeRequest) SetChildInstanceRegionId(v string) *DescribeCenAttachedChildInstanceAttributeRequest {
	s.ChildInstanceRegionId = &v
	return s
}

func (s *DescribeCenAttachedChildInstanceAttributeRequest) SetChildInstanceType(v string) *DescribeCenAttachedChildInstanceAttributeRequest {
	s.ChildInstanceType = &v
	return s
}

func (s *DescribeCenAttachedChildInstanceAttributeRequest) SetOwnerAccount(v string) *DescribeCenAttachedChildInstanceAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeCenAttachedChildInstanceAttributeRequest) SetOwnerId(v int64) *DescribeCenAttachedChildInstanceAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeCenAttachedChildInstanceAttributeRequest) SetResourceOwnerAccount(v string) *DescribeCenAttachedChildInstanceAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeCenAttachedChildInstanceAttributeRequest) SetResourceOwnerId(v int64) *DescribeCenAttachedChildInstanceAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeCenAttachedChildInstanceAttributeRequest) Validate() error {
	return dara.Validate(s)
}
