// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePublishedRouteEntriesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCenId(v string) *DescribePublishedRouteEntriesRequest
	GetCenId() *string
	SetChildInstanceId(v string) *DescribePublishedRouteEntriesRequest
	GetChildInstanceId() *string
	SetChildInstanceRegionId(v string) *DescribePublishedRouteEntriesRequest
	GetChildInstanceRegionId() *string
	SetChildInstanceRouteTableId(v string) *DescribePublishedRouteEntriesRequest
	GetChildInstanceRouteTableId() *string
	SetChildInstanceType(v string) *DescribePublishedRouteEntriesRequest
	GetChildInstanceType() *string
	SetDestinationCidrBlock(v string) *DescribePublishedRouteEntriesRequest
	GetDestinationCidrBlock() *string
	SetPageNumber(v int32) *DescribePublishedRouteEntriesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribePublishedRouteEntriesRequest
	GetPageSize() *int32
	SetResourceOwnerAccount(v string) *DescribePublishedRouteEntriesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribePublishedRouteEntriesRequest
	GetResourceOwnerId() *int64
}

type DescribePublishedRouteEntriesRequest struct {
	// The ID of the CEN instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cen-7qthudw0ll6jm****
	CenId *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	// The ID of the network instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpc-bp18sth14qii3pnv****
	ChildInstanceId *string `json:"ChildInstanceId,omitempty" xml:"ChildInstanceId,omitempty"`
	// The ID of the region where the network instance is deployed.
	//
	// You can call the [DescribeChildInstanceRegions](https://help.aliyun.com/document_detail/132080.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	ChildInstanceRegionId *string `json:"ChildInstanceRegionId,omitempty" xml:"ChildInstanceRegionId,omitempty"`
	// The ID of the route table of the network instance.
	//
	// example:
	//
	// vtb-bp174d1gje79u1g4****
	ChildInstanceRouteTableId *string `json:"ChildInstanceRouteTableId,omitempty" xml:"ChildInstanceRouteTableId,omitempty"`
	// The type of the network instance. Valid values:
	//
	// 	- **VPC**: VPC
	//
	// 	- **VBR**: VBR
	//
	// 	- **CCN**: Cloud Connect Network (CCN) instance
	//
	// This parameter is required.
	//
	// example:
	//
	// VPC
	ChildInstanceType *string `json:"ChildInstanceType,omitempty" xml:"ChildInstanceType,omitempty"`
	// The destination CIDR block of the route that you want to query.
	//
	// example:
	//
	// 172.16.0.0/16
	DestinationCidrBlock *string `json:"DestinationCidrBlock,omitempty" xml:"DestinationCidrBlock,omitempty"`
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

func (s DescribePublishedRouteEntriesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePublishedRouteEntriesRequest) GoString() string {
	return s.String()
}

func (s *DescribePublishedRouteEntriesRequest) GetCenId() *string {
	return s.CenId
}

func (s *DescribePublishedRouteEntriesRequest) GetChildInstanceId() *string {
	return s.ChildInstanceId
}

func (s *DescribePublishedRouteEntriesRequest) GetChildInstanceRegionId() *string {
	return s.ChildInstanceRegionId
}

func (s *DescribePublishedRouteEntriesRequest) GetChildInstanceRouteTableId() *string {
	return s.ChildInstanceRouteTableId
}

func (s *DescribePublishedRouteEntriesRequest) GetChildInstanceType() *string {
	return s.ChildInstanceType
}

func (s *DescribePublishedRouteEntriesRequest) GetDestinationCidrBlock() *string {
	return s.DestinationCidrBlock
}

func (s *DescribePublishedRouteEntriesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribePublishedRouteEntriesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribePublishedRouteEntriesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribePublishedRouteEntriesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribePublishedRouteEntriesRequest) SetCenId(v string) *DescribePublishedRouteEntriesRequest {
	s.CenId = &v
	return s
}

func (s *DescribePublishedRouteEntriesRequest) SetChildInstanceId(v string) *DescribePublishedRouteEntriesRequest {
	s.ChildInstanceId = &v
	return s
}

func (s *DescribePublishedRouteEntriesRequest) SetChildInstanceRegionId(v string) *DescribePublishedRouteEntriesRequest {
	s.ChildInstanceRegionId = &v
	return s
}

func (s *DescribePublishedRouteEntriesRequest) SetChildInstanceRouteTableId(v string) *DescribePublishedRouteEntriesRequest {
	s.ChildInstanceRouteTableId = &v
	return s
}

func (s *DescribePublishedRouteEntriesRequest) SetChildInstanceType(v string) *DescribePublishedRouteEntriesRequest {
	s.ChildInstanceType = &v
	return s
}

func (s *DescribePublishedRouteEntriesRequest) SetDestinationCidrBlock(v string) *DescribePublishedRouteEntriesRequest {
	s.DestinationCidrBlock = &v
	return s
}

func (s *DescribePublishedRouteEntriesRequest) SetPageNumber(v int32) *DescribePublishedRouteEntriesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribePublishedRouteEntriesRequest) SetPageSize(v int32) *DescribePublishedRouteEntriesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribePublishedRouteEntriesRequest) SetResourceOwnerAccount(v string) *DescribePublishedRouteEntriesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribePublishedRouteEntriesRequest) SetResourceOwnerId(v int64) *DescribePublishedRouteEntriesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribePublishedRouteEntriesRequest) Validate() error {
	return dara.Validate(s)
}
