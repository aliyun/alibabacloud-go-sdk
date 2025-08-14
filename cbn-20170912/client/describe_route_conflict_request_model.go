// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRouteConflictRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChildInstanceId(v string) *DescribeRouteConflictRequest
	GetChildInstanceId() *string
	SetChildInstanceRegionId(v string) *DescribeRouteConflictRequest
	GetChildInstanceRegionId() *string
	SetChildInstanceRouteTableId(v string) *DescribeRouteConflictRequest
	GetChildInstanceRouteTableId() *string
	SetChildInstanceType(v string) *DescribeRouteConflictRequest
	GetChildInstanceType() *string
	SetDestinationCidrBlock(v string) *DescribeRouteConflictRequest
	GetDestinationCidrBlock() *string
	SetOwnerAccount(v string) *DescribeRouteConflictRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeRouteConflictRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeRouteConflictRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeRouteConflictRequest
	GetPageSize() *int32
	SetResourceOwnerAccount(v string) *DescribeRouteConflictRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeRouteConflictRequest
	GetResourceOwnerId() *int64
}

type DescribeRouteConflictRequest struct {
	// The ID of the network instance that you want to query.
	//
	// This parameter is required.
	//
	// example:
	//
	// ccn-0q3b7oviikmm9h****
	ChildInstanceId *string `json:"ChildInstanceId,omitempty" xml:"ChildInstanceId,omitempty"`
	// The ID of the region where the network instance is deployed.
	//
	// You can call the [DescribeChildInstanceRegions](https://help.aliyun.com/document_detail/132080.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// ccn-cn-shanghai
	ChildInstanceRegionId *string `json:"ChildInstanceRegionId,omitempty" xml:"ChildInstanceRegionId,omitempty"`
	// The ID of the route table that is configured on the network instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// vtb-bp174d1gje79u1g4t****
	ChildInstanceRouteTableId *string `json:"ChildInstanceRouteTableId,omitempty" xml:"ChildInstanceRouteTableId,omitempty"`
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
	// CCN
	ChildInstanceType *string `json:"ChildInstanceType,omitempty" xml:"ChildInstanceType,omitempty"`
	// The destination CIDR block of the conflicting route.
	//
	// example:
	//
	// 172.16.0.0/16
	DestinationCidrBlock *string `json:"DestinationCidrBlock,omitempty" xml:"DestinationCidrBlock,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The number of the page to return. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Default value: **10**. Valid values: **1*	- to **50**.
	//
	// example:
	//
	// 10
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeRouteConflictRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRouteConflictRequest) GoString() string {
	return s.String()
}

func (s *DescribeRouteConflictRequest) GetChildInstanceId() *string {
	return s.ChildInstanceId
}

func (s *DescribeRouteConflictRequest) GetChildInstanceRegionId() *string {
	return s.ChildInstanceRegionId
}

func (s *DescribeRouteConflictRequest) GetChildInstanceRouteTableId() *string {
	return s.ChildInstanceRouteTableId
}

func (s *DescribeRouteConflictRequest) GetChildInstanceType() *string {
	return s.ChildInstanceType
}

func (s *DescribeRouteConflictRequest) GetDestinationCidrBlock() *string {
	return s.DestinationCidrBlock
}

func (s *DescribeRouteConflictRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeRouteConflictRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeRouteConflictRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeRouteConflictRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeRouteConflictRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeRouteConflictRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeRouteConflictRequest) SetChildInstanceId(v string) *DescribeRouteConflictRequest {
	s.ChildInstanceId = &v
	return s
}

func (s *DescribeRouteConflictRequest) SetChildInstanceRegionId(v string) *DescribeRouteConflictRequest {
	s.ChildInstanceRegionId = &v
	return s
}

func (s *DescribeRouteConflictRequest) SetChildInstanceRouteTableId(v string) *DescribeRouteConflictRequest {
	s.ChildInstanceRouteTableId = &v
	return s
}

func (s *DescribeRouteConflictRequest) SetChildInstanceType(v string) *DescribeRouteConflictRequest {
	s.ChildInstanceType = &v
	return s
}

func (s *DescribeRouteConflictRequest) SetDestinationCidrBlock(v string) *DescribeRouteConflictRequest {
	s.DestinationCidrBlock = &v
	return s
}

func (s *DescribeRouteConflictRequest) SetOwnerAccount(v string) *DescribeRouteConflictRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeRouteConflictRequest) SetOwnerId(v int64) *DescribeRouteConflictRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeRouteConflictRequest) SetPageNumber(v int32) *DescribeRouteConflictRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeRouteConflictRequest) SetPageSize(v int32) *DescribeRouteConflictRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeRouteConflictRequest) SetResourceOwnerAccount(v string) *DescribeRouteConflictRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeRouteConflictRequest) SetResourceOwnerId(v int64) *DescribeRouteConflictRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeRouteConflictRequest) Validate() error {
	return dara.Validate(s)
}
