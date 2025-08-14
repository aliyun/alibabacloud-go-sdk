// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTransitRouterPrefixListAssociationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DeleteTransitRouterPrefixListAssociationRequest
	GetClientToken() *string
	SetDryRun(v bool) *DeleteTransitRouterPrefixListAssociationRequest
	GetDryRun() *bool
	SetNextHop(v string) *DeleteTransitRouterPrefixListAssociationRequest
	GetNextHop() *string
	SetNextHopType(v string) *DeleteTransitRouterPrefixListAssociationRequest
	GetNextHopType() *string
	SetOwnerAccount(v string) *DeleteTransitRouterPrefixListAssociationRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteTransitRouterPrefixListAssociationRequest
	GetOwnerId() *int64
	SetPrefixListId(v string) *DeleteTransitRouterPrefixListAssociationRequest
	GetPrefixListId() *string
	SetRegionId(v string) *DeleteTransitRouterPrefixListAssociationRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DeleteTransitRouterPrefixListAssociationRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteTransitRouterPrefixListAssociationRequest
	GetResourceOwnerId() *int64
	SetTransitRouterId(v string) *DeleteTransitRouterPrefixListAssociationRequest
	GetTransitRouterId() *string
	SetTransitRouterTableId(v string) *DeleteTransitRouterPrefixListAssociationRequest
	GetTransitRouterTableId() *string
}

type DeleteTransitRouterPrefixListAssociationRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among all requests. The token can contain only ASCII characters.
	//
	// >  If you do not set this parameter, **ClientToken*	- is set to the value of **RequestId**. The value of **RequestId*	- for each API request may be different.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-4266****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform a dry run. Valid values:
	//
	// 	- **true**: performs a dry run. The system checks the required parameters, request syntax, and limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// 	- **false*	- (default): performs a dry run and sends the task.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ID of the next hop.
	//
	// > If **NextHopType*	- is set to **BlackHole**, you must set this parameter to **BlackHole**.
	//
	// This parameter is required.
	//
	// example:
	//
	// tr-attach-flbq507rg2ckrj****
	NextHop *string `json:"NextHop,omitempty" xml:"NextHop,omitempty"`
	// The type of the next hop. Valid values:
	//
	// 	- **BlackHole**: All the CIDR blocks in the prefix list are blackhole routes. Packets destined for the CIDR blocks are dropped.
	//
	// 	- **VPC**: The next hop of the CIDR blocks in the prefix list is a VPC connection.
	//
	// 	- **VBR**: The next hop of the CIDR blocks in the prefix list is a VBR connection.
	//
	// 	- **TR**: The next hop of the CIDR blocks in the prefix list is an inter-region connection.
	//
	// example:
	//
	// VPC
	NextHopType  *string `json:"NextHopType,omitempty" xml:"NextHopType,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the prefix list.
	//
	// This parameter is required.
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
	// This parameter is required.
	//
	// example:
	//
	// vtb-6ehgc262hr170qgyc****
	TransitRouterTableId *string `json:"TransitRouterTableId,omitempty" xml:"TransitRouterTableId,omitempty"`
}

func (s DeleteTransitRouterPrefixListAssociationRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteTransitRouterPrefixListAssociationRequest) GoString() string {
	return s.String()
}

func (s *DeleteTransitRouterPrefixListAssociationRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteTransitRouterPrefixListAssociationRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *DeleteTransitRouterPrefixListAssociationRequest) GetNextHop() *string {
	return s.NextHop
}

func (s *DeleteTransitRouterPrefixListAssociationRequest) GetNextHopType() *string {
	return s.NextHopType
}

func (s *DeleteTransitRouterPrefixListAssociationRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteTransitRouterPrefixListAssociationRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteTransitRouterPrefixListAssociationRequest) GetPrefixListId() *string {
	return s.PrefixListId
}

func (s *DeleteTransitRouterPrefixListAssociationRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteTransitRouterPrefixListAssociationRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteTransitRouterPrefixListAssociationRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteTransitRouterPrefixListAssociationRequest) GetTransitRouterId() *string {
	return s.TransitRouterId
}

func (s *DeleteTransitRouterPrefixListAssociationRequest) GetTransitRouterTableId() *string {
	return s.TransitRouterTableId
}

func (s *DeleteTransitRouterPrefixListAssociationRequest) SetClientToken(v string) *DeleteTransitRouterPrefixListAssociationRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteTransitRouterPrefixListAssociationRequest) SetDryRun(v bool) *DeleteTransitRouterPrefixListAssociationRequest {
	s.DryRun = &v
	return s
}

func (s *DeleteTransitRouterPrefixListAssociationRequest) SetNextHop(v string) *DeleteTransitRouterPrefixListAssociationRequest {
	s.NextHop = &v
	return s
}

func (s *DeleteTransitRouterPrefixListAssociationRequest) SetNextHopType(v string) *DeleteTransitRouterPrefixListAssociationRequest {
	s.NextHopType = &v
	return s
}

func (s *DeleteTransitRouterPrefixListAssociationRequest) SetOwnerAccount(v string) *DeleteTransitRouterPrefixListAssociationRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteTransitRouterPrefixListAssociationRequest) SetOwnerId(v int64) *DeleteTransitRouterPrefixListAssociationRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteTransitRouterPrefixListAssociationRequest) SetPrefixListId(v string) *DeleteTransitRouterPrefixListAssociationRequest {
	s.PrefixListId = &v
	return s
}

func (s *DeleteTransitRouterPrefixListAssociationRequest) SetRegionId(v string) *DeleteTransitRouterPrefixListAssociationRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteTransitRouterPrefixListAssociationRequest) SetResourceOwnerAccount(v string) *DeleteTransitRouterPrefixListAssociationRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteTransitRouterPrefixListAssociationRequest) SetResourceOwnerId(v int64) *DeleteTransitRouterPrefixListAssociationRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteTransitRouterPrefixListAssociationRequest) SetTransitRouterId(v string) *DeleteTransitRouterPrefixListAssociationRequest {
	s.TransitRouterId = &v
	return s
}

func (s *DeleteTransitRouterPrefixListAssociationRequest) SetTransitRouterTableId(v string) *DeleteTransitRouterPrefixListAssociationRequest {
	s.TransitRouterTableId = &v
	return s
}

func (s *DeleteTransitRouterPrefixListAssociationRequest) Validate() error {
	return dara.Validate(s)
}
