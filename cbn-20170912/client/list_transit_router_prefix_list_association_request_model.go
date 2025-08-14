// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTransitRouterPrefixListAssociationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNextHop(v string) *ListTransitRouterPrefixListAssociationRequest
	GetNextHop() *string
	SetNextHopInstanceId(v string) *ListTransitRouterPrefixListAssociationRequest
	GetNextHopInstanceId() *string
	SetNextHopType(v string) *ListTransitRouterPrefixListAssociationRequest
	GetNextHopType() *string
	SetOwnerAccount(v string) *ListTransitRouterPrefixListAssociationRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ListTransitRouterPrefixListAssociationRequest
	GetOwnerId() *int64
	SetOwnerUid(v int64) *ListTransitRouterPrefixListAssociationRequest
	GetOwnerUid() *int64
	SetPageNumber(v int32) *ListTransitRouterPrefixListAssociationRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListTransitRouterPrefixListAssociationRequest
	GetPageSize() *int32
	SetPrefixListId(v string) *ListTransitRouterPrefixListAssociationRequest
	GetPrefixListId() *string
	SetRegionId(v string) *ListTransitRouterPrefixListAssociationRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ListTransitRouterPrefixListAssociationRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ListTransitRouterPrefixListAssociationRequest
	GetResourceOwnerId() *int64
	SetStatus(v string) *ListTransitRouterPrefixListAssociationRequest
	GetStatus() *string
	SetTransitRouterId(v string) *ListTransitRouterPrefixListAssociationRequest
	GetTransitRouterId() *string
	SetTransitRouterTableId(v string) *ListTransitRouterPrefixListAssociationRequest
	GetTransitRouterTableId() *string
}

type ListTransitRouterPrefixListAssociationRequest struct {
	// The ID of the next hop.
	//
	// > Set the value to **BlackHole*	- if you want to query the prefix list that generates blackhole routes.
	//
	// example:
	//
	// tr-attach-flbq507rg2ckrj****
	NextHop *string `json:"NextHop,omitempty" xml:"NextHop,omitempty"`
	// The ID of the network instance associated with the next hop connection.
	//
	// example:
	//
	// vpc-6eh7fp9hdqa2wv85t****
	NextHopInstanceId *string `json:"NextHopInstanceId,omitempty" xml:"NextHopInstanceId,omitempty"`
	// The type of the next hop. Valid values:
	//
	// 	- **BlackHole**: The prefix list that generates blackhole routes.
	//
	// 	- **VPC**: The prefix list whose next hop is a virtual private cloud (VPC) connection.
	//
	// 	- **VBR**: The prefix list whose next hop is a virtual border router (VBR) connection.
	//
	// 	- **TR**: The prefix list whose next hop is an inter-region connection on the transit router.
	//
	// example:
	//
	// VPC
	NextHopType  *string `json:"NextHopType,omitempty" xml:"NextHopType,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the Alibaba Cloud account to which the prefix list belongs.
	//
	// example:
	//
	// 1210123456123456
	OwnerUid *int64 `json:"OwnerUid,omitempty" xml:"OwnerUid,omitempty"`
	// The number of the page to return. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Maximum value: **100**. Default value: **10**.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the prefix list.
	//
	// example:
	//
	// pl-6ehtn5kqxgeyy08fi****
	PrefixListId *string `json:"PrefixListId,omitempty" xml:"PrefixListId,omitempty"`
	// The ID of the region where the transit router is deployed.
	//
	// You can call the [DescribeChildInstanceRegions](https://help.aliyun.com/document_detail/132080.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The status of the prefix list. Valid values:
	//
	// 	- **Active**
	//
	// 	- **Updating**
	//
	// example:
	//
	// Active
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the transit router.
	//
	// This parameter is required.
	//
	// example:
	//
	// tr-6ehx7q2jze8ch5ji0****
	TransitRouterId *string `json:"TransitRouterId,omitempty" xml:"TransitRouterId,omitempty"`
	// The ID of the route table of the transit router.
	//
	// example:
	//
	// vtb-6ehgc262hr170qgyc****
	TransitRouterTableId *string `json:"TransitRouterTableId,omitempty" xml:"TransitRouterTableId,omitempty"`
}

func (s ListTransitRouterPrefixListAssociationRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTransitRouterPrefixListAssociationRequest) GoString() string {
	return s.String()
}

func (s *ListTransitRouterPrefixListAssociationRequest) GetNextHop() *string {
	return s.NextHop
}

func (s *ListTransitRouterPrefixListAssociationRequest) GetNextHopInstanceId() *string {
	return s.NextHopInstanceId
}

func (s *ListTransitRouterPrefixListAssociationRequest) GetNextHopType() *string {
	return s.NextHopType
}

func (s *ListTransitRouterPrefixListAssociationRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ListTransitRouterPrefixListAssociationRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListTransitRouterPrefixListAssociationRequest) GetOwnerUid() *int64 {
	return s.OwnerUid
}

func (s *ListTransitRouterPrefixListAssociationRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListTransitRouterPrefixListAssociationRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListTransitRouterPrefixListAssociationRequest) GetPrefixListId() *string {
	return s.PrefixListId
}

func (s *ListTransitRouterPrefixListAssociationRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListTransitRouterPrefixListAssociationRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ListTransitRouterPrefixListAssociationRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ListTransitRouterPrefixListAssociationRequest) GetStatus() *string {
	return s.Status
}

func (s *ListTransitRouterPrefixListAssociationRequest) GetTransitRouterId() *string {
	return s.TransitRouterId
}

func (s *ListTransitRouterPrefixListAssociationRequest) GetTransitRouterTableId() *string {
	return s.TransitRouterTableId
}

func (s *ListTransitRouterPrefixListAssociationRequest) SetNextHop(v string) *ListTransitRouterPrefixListAssociationRequest {
	s.NextHop = &v
	return s
}

func (s *ListTransitRouterPrefixListAssociationRequest) SetNextHopInstanceId(v string) *ListTransitRouterPrefixListAssociationRequest {
	s.NextHopInstanceId = &v
	return s
}

func (s *ListTransitRouterPrefixListAssociationRequest) SetNextHopType(v string) *ListTransitRouterPrefixListAssociationRequest {
	s.NextHopType = &v
	return s
}

func (s *ListTransitRouterPrefixListAssociationRequest) SetOwnerAccount(v string) *ListTransitRouterPrefixListAssociationRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ListTransitRouterPrefixListAssociationRequest) SetOwnerId(v int64) *ListTransitRouterPrefixListAssociationRequest {
	s.OwnerId = &v
	return s
}

func (s *ListTransitRouterPrefixListAssociationRequest) SetOwnerUid(v int64) *ListTransitRouterPrefixListAssociationRequest {
	s.OwnerUid = &v
	return s
}

func (s *ListTransitRouterPrefixListAssociationRequest) SetPageNumber(v int32) *ListTransitRouterPrefixListAssociationRequest {
	s.PageNumber = &v
	return s
}

func (s *ListTransitRouterPrefixListAssociationRequest) SetPageSize(v int32) *ListTransitRouterPrefixListAssociationRequest {
	s.PageSize = &v
	return s
}

func (s *ListTransitRouterPrefixListAssociationRequest) SetPrefixListId(v string) *ListTransitRouterPrefixListAssociationRequest {
	s.PrefixListId = &v
	return s
}

func (s *ListTransitRouterPrefixListAssociationRequest) SetRegionId(v string) *ListTransitRouterPrefixListAssociationRequest {
	s.RegionId = &v
	return s
}

func (s *ListTransitRouterPrefixListAssociationRequest) SetResourceOwnerAccount(v string) *ListTransitRouterPrefixListAssociationRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListTransitRouterPrefixListAssociationRequest) SetResourceOwnerId(v int64) *ListTransitRouterPrefixListAssociationRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListTransitRouterPrefixListAssociationRequest) SetStatus(v string) *ListTransitRouterPrefixListAssociationRequest {
	s.Status = &v
	return s
}

func (s *ListTransitRouterPrefixListAssociationRequest) SetTransitRouterId(v string) *ListTransitRouterPrefixListAssociationRequest {
	s.TransitRouterId = &v
	return s
}

func (s *ListTransitRouterPrefixListAssociationRequest) SetTransitRouterTableId(v string) *ListTransitRouterPrefixListAssociationRequest {
	s.TransitRouterTableId = &v
	return s
}

func (s *ListTransitRouterPrefixListAssociationRequest) Validate() error {
	return dara.Validate(s)
}
