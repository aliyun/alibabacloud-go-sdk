// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCenChildInstanceRouteEntriesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCenId(v string) *DescribeCenChildInstanceRouteEntriesRequest
	GetCenId() *string
	SetChildInstanceId(v string) *DescribeCenChildInstanceRouteEntriesRequest
	GetChildInstanceId() *string
	SetChildInstanceRegionId(v string) *DescribeCenChildInstanceRouteEntriesRequest
	GetChildInstanceRegionId() *string
	SetChildInstanceRouteTableId(v string) *DescribeCenChildInstanceRouteEntriesRequest
	GetChildInstanceRouteTableId() *string
	SetChildInstanceType(v string) *DescribeCenChildInstanceRouteEntriesRequest
	GetChildInstanceType() *string
	SetOwnerAccount(v string) *DescribeCenChildInstanceRouteEntriesRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeCenChildInstanceRouteEntriesRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeCenChildInstanceRouteEntriesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeCenChildInstanceRouteEntriesRequest
	GetPageSize() *int32
	SetResourceOwnerAccount(v string) *DescribeCenChildInstanceRouteEntriesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeCenChildInstanceRouteEntriesRequest
	GetResourceOwnerId() *int64
	SetStatus(v string) *DescribeCenChildInstanceRouteEntriesRequest
	GetStatus() *string
}

type DescribeCenChildInstanceRouteEntriesRequest struct {
	// The ID of the CEN instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cen-7qthudw0ll6jmc****
	CenId *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	// The ID of the network instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpc-bp18sth14qii3pnvo****
	ChildInstanceId *string `json:"ChildInstanceId,omitempty" xml:"ChildInstanceId,omitempty"`
	// The ID of the region where the network instance is deployed.
	//
	// You can call the [DescribeChildInstanceRegions](https://help.aliyun.com/document_detail/132080.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
	ChildInstanceRegionId *string `json:"ChildInstanceRegionId,omitempty" xml:"ChildInstanceRegionId,omitempty"`
	// The ID of the route table of the network instance. If you do not specify a route table ID, the system queries the routes in the system route tables of the VPCs by default.
	//
	// example:
	//
	// vtb-p0wxx3apzgn6uqp3r****
	ChildInstanceRouteTableId *string `json:"ChildInstanceRouteTableId,omitempty" xml:"ChildInstanceRouteTableId,omitempty"`
	// The type of the network instance. Valid values:
	//
	// 	- **VPC**: virtual private cloud (VPC)
	//
	// 	- **VBR**: virtual border router (VBR)
	//
	// 	- **CCN**: Cloud Connect Network (CCN) instance
	//
	// 	- **ECR**: Express Connect Router (ECR)
	//
	// This parameter is required.
	//
	// example:
	//
	// VPC
	ChildInstanceType *string `json:"ChildInstanceType,omitempty" xml:"ChildInstanceType,omitempty"`
	OwnerAccount      *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId           *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The number of the page to return. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values: **1*	- to **500**. Default value: **10**.
	//
	// example:
	//
	// 10
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The status of the route. Valid values:
	//
	// 	- **Active**: available
	//
	// 	- **Candidate**: standby
	//
	// 	- **Rejected**: rejected
	//
	// 	- **Prohibited**: prohibited
	//
	// 	- **All*	- (default): all routes
	//
	// example:
	//
	// Active
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeCenChildInstanceRouteEntriesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCenChildInstanceRouteEntriesRequest) GoString() string {
	return s.String()
}

func (s *DescribeCenChildInstanceRouteEntriesRequest) GetCenId() *string {
	return s.CenId
}

func (s *DescribeCenChildInstanceRouteEntriesRequest) GetChildInstanceId() *string {
	return s.ChildInstanceId
}

func (s *DescribeCenChildInstanceRouteEntriesRequest) GetChildInstanceRegionId() *string {
	return s.ChildInstanceRegionId
}

func (s *DescribeCenChildInstanceRouteEntriesRequest) GetChildInstanceRouteTableId() *string {
	return s.ChildInstanceRouteTableId
}

func (s *DescribeCenChildInstanceRouteEntriesRequest) GetChildInstanceType() *string {
	return s.ChildInstanceType
}

func (s *DescribeCenChildInstanceRouteEntriesRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeCenChildInstanceRouteEntriesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeCenChildInstanceRouteEntriesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeCenChildInstanceRouteEntriesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeCenChildInstanceRouteEntriesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeCenChildInstanceRouteEntriesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeCenChildInstanceRouteEntriesRequest) GetStatus() *string {
	return s.Status
}

func (s *DescribeCenChildInstanceRouteEntriesRequest) SetCenId(v string) *DescribeCenChildInstanceRouteEntriesRequest {
	s.CenId = &v
	return s
}

func (s *DescribeCenChildInstanceRouteEntriesRequest) SetChildInstanceId(v string) *DescribeCenChildInstanceRouteEntriesRequest {
	s.ChildInstanceId = &v
	return s
}

func (s *DescribeCenChildInstanceRouteEntriesRequest) SetChildInstanceRegionId(v string) *DescribeCenChildInstanceRouteEntriesRequest {
	s.ChildInstanceRegionId = &v
	return s
}

func (s *DescribeCenChildInstanceRouteEntriesRequest) SetChildInstanceRouteTableId(v string) *DescribeCenChildInstanceRouteEntriesRequest {
	s.ChildInstanceRouteTableId = &v
	return s
}

func (s *DescribeCenChildInstanceRouteEntriesRequest) SetChildInstanceType(v string) *DescribeCenChildInstanceRouteEntriesRequest {
	s.ChildInstanceType = &v
	return s
}

func (s *DescribeCenChildInstanceRouteEntriesRequest) SetOwnerAccount(v string) *DescribeCenChildInstanceRouteEntriesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeCenChildInstanceRouteEntriesRequest) SetOwnerId(v int64) *DescribeCenChildInstanceRouteEntriesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeCenChildInstanceRouteEntriesRequest) SetPageNumber(v int32) *DescribeCenChildInstanceRouteEntriesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeCenChildInstanceRouteEntriesRequest) SetPageSize(v int32) *DescribeCenChildInstanceRouteEntriesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeCenChildInstanceRouteEntriesRequest) SetResourceOwnerAccount(v string) *DescribeCenChildInstanceRouteEntriesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeCenChildInstanceRouteEntriesRequest) SetResourceOwnerId(v int64) *DescribeCenChildInstanceRouteEntriesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeCenChildInstanceRouteEntriesRequest) SetStatus(v string) *DescribeCenChildInstanceRouteEntriesRequest {
	s.Status = &v
	return s
}

func (s *DescribeCenChildInstanceRouteEntriesRequest) Validate() error {
	return dara.Validate(s)
}
