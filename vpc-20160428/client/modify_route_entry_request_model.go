// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyRouteEntryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *ModifyRouteEntryRequest
	GetDescription() *string
	SetDestinationCidrBlock(v string) *ModifyRouteEntryRequest
	GetDestinationCidrBlock() *string
	SetDryRun(v bool) *ModifyRouteEntryRequest
	GetDryRun() *bool
	SetNewNextHopId(v string) *ModifyRouteEntryRequest
	GetNewNextHopId() *string
	SetNewNextHopType(v string) *ModifyRouteEntryRequest
	GetNewNextHopType() *string
	SetOwnerAccount(v string) *ModifyRouteEntryRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyRouteEntryRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ModifyRouteEntryRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifyRouteEntryRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyRouteEntryRequest
	GetResourceOwnerId() *int64
	SetRouteEntryId(v string) *ModifyRouteEntryRequest
	GetRouteEntryId() *string
	SetRouteEntryName(v string) *ModifyRouteEntryRequest
	GetRouteEntryName() *string
	SetRouteTableId(v string) *ModifyRouteEntryRequest
	GetRouteTableId() *string
}

type ModifyRouteEntryRequest struct {
	// The description of the route entry.
	//
	// The description must be 1 to 256 characters in length, and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// EntryDescription
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The destination CIDR block of the route entry. Only IPv4 CIDR blocks, IPv6 CIDR blocks, and prefix lists are supported.
	//
	// example:
	//
	// 192.168.0.0/24
	DestinationCidrBlock *string `json:"DestinationCidrBlock,omitempty" xml:"DestinationCidrBlock,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// 	- **true**: performs a dry run. The system checks the request for potential issues, including the AccessKey pair, the permissions of the RAM user, and the required parameters. If the request fails the dry run, an error message is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// 	- **false*	- (default): performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ID of the new next hop instance.
	//
	// example:
	//
	// eni-bp17y37ytsenqyim****
	NewNextHopId *string `json:"NewNextHopId,omitempty" xml:"NewNextHopId,omitempty"`
	// The new next hop type of the route.
	//
	// example:
	//
	// NetworkInterface
	NewNextHopType *string `json:"NewNextHopType,omitempty" xml:"NewNextHopType,omitempty"`
	OwnerAccount   *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId        *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region to which the route belongs.
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
	// The ID of the custom route entry.
	//
	// example:
	//
	// rte-acfvgfsghfdd****
	RouteEntryId *string `json:"RouteEntryId,omitempty" xml:"RouteEntryId,omitempty"`
	// The name of the route entry.
	//
	// The name must be 1 to 128 characters in length, and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// EntryName
	RouteEntryName *string `json:"RouteEntryName,omitempty" xml:"RouteEntryName,omitempty"`
	// The ID of the route table to which the route entry belongs.
	//
	// example:
	//
	// vtb-bp1nk7zk65du3pni8z9td
	RouteTableId *string `json:"RouteTableId,omitempty" xml:"RouteTableId,omitempty"`
}

func (s ModifyRouteEntryRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyRouteEntryRequest) GoString() string {
	return s.String()
}

func (s *ModifyRouteEntryRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifyRouteEntryRequest) GetDestinationCidrBlock() *string {
	return s.DestinationCidrBlock
}

func (s *ModifyRouteEntryRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *ModifyRouteEntryRequest) GetNewNextHopId() *string {
	return s.NewNextHopId
}

func (s *ModifyRouteEntryRequest) GetNewNextHopType() *string {
	return s.NewNextHopType
}

func (s *ModifyRouteEntryRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyRouteEntryRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyRouteEntryRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyRouteEntryRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyRouteEntryRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyRouteEntryRequest) GetRouteEntryId() *string {
	return s.RouteEntryId
}

func (s *ModifyRouteEntryRequest) GetRouteEntryName() *string {
	return s.RouteEntryName
}

func (s *ModifyRouteEntryRequest) GetRouteTableId() *string {
	return s.RouteTableId
}

func (s *ModifyRouteEntryRequest) SetDescription(v string) *ModifyRouteEntryRequest {
	s.Description = &v
	return s
}

func (s *ModifyRouteEntryRequest) SetDestinationCidrBlock(v string) *ModifyRouteEntryRequest {
	s.DestinationCidrBlock = &v
	return s
}

func (s *ModifyRouteEntryRequest) SetDryRun(v bool) *ModifyRouteEntryRequest {
	s.DryRun = &v
	return s
}

func (s *ModifyRouteEntryRequest) SetNewNextHopId(v string) *ModifyRouteEntryRequest {
	s.NewNextHopId = &v
	return s
}

func (s *ModifyRouteEntryRequest) SetNewNextHopType(v string) *ModifyRouteEntryRequest {
	s.NewNextHopType = &v
	return s
}

func (s *ModifyRouteEntryRequest) SetOwnerAccount(v string) *ModifyRouteEntryRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyRouteEntryRequest) SetOwnerId(v int64) *ModifyRouteEntryRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyRouteEntryRequest) SetRegionId(v string) *ModifyRouteEntryRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyRouteEntryRequest) SetResourceOwnerAccount(v string) *ModifyRouteEntryRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyRouteEntryRequest) SetResourceOwnerId(v int64) *ModifyRouteEntryRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyRouteEntryRequest) SetRouteEntryId(v string) *ModifyRouteEntryRequest {
	s.RouteEntryId = &v
	return s
}

func (s *ModifyRouteEntryRequest) SetRouteEntryName(v string) *ModifyRouteEntryRequest {
	s.RouteEntryName = &v
	return s
}

func (s *ModifyRouteEntryRequest) SetRouteTableId(v string) *ModifyRouteEntryRequest {
	s.RouteTableId = &v
	return s
}

func (s *ModifyRouteEntryRequest) Validate() error {
	return dara.Validate(s)
}
