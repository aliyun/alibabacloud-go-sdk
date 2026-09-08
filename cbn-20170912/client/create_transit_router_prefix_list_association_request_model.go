// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTransitRouterPrefixListAssociationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateTransitRouterPrefixListAssociationRequest
	GetClientToken() *string
	SetDryRun(v bool) *CreateTransitRouterPrefixListAssociationRequest
	GetDryRun() *bool
	SetNextHop(v string) *CreateTransitRouterPrefixListAssociationRequest
	GetNextHop() *string
	SetNextHopType(v string) *CreateTransitRouterPrefixListAssociationRequest
	GetNextHopType() *string
	SetOwnerAccount(v string) *CreateTransitRouterPrefixListAssociationRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateTransitRouterPrefixListAssociationRequest
	GetOwnerId() *int64
	SetOwnerUid(v int64) *CreateTransitRouterPrefixListAssociationRequest
	GetOwnerUid() *int64
	SetPrefixListId(v string) *CreateTransitRouterPrefixListAssociationRequest
	GetPrefixListId() *string
	SetRegionId(v string) *CreateTransitRouterPrefixListAssociationRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *CreateTransitRouterPrefixListAssociationRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateTransitRouterPrefixListAssociationRequest
	GetResourceOwnerId() *int64
	SetTransitRouterId(v string) *CreateTransitRouterPrefixListAssociationRequest
	GetTransitRouterId() *string
	SetTransitRouterTableId(v string) *CreateTransitRouterPrefixListAssociationRequest
	GetTransitRouterTableId() *string
}

type CreateTransitRouterPrefixListAssociationRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The client token can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **RequestId*	- of the API request as the **ClientToken**. The **RequestId*	- may be different for each API request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-4266****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform a dry run. Valid values:
	//
	// - **true**: performs a dry run without creating the association. The system checks the required parameters, request syntax, and business restrictions. If the check fails, the corresponding error is returned. If the check succeeds, the error code `DryRunOperation` is returned.
	//
	// - **false*	- (default): performs a dry run and creates the association after the check is passed.
	//
	// > This parameter is currently not available.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ID of the next hop connection.
	//
	// To configure all CIDR blocks in the prefix list as blackhole routes, set this parameter to **BlackHole**.
	//
	// This parameter is required.
	//
	// example:
	//
	// tr-attach-flbq507rg2ckrj****
	NextHop *string `json:"NextHop,omitempty" xml:"NextHop,omitempty"`
	// The next hop type. Valid values:
	//
	// - **BlackHole**: All CIDR blocks in the prefix list are blackhole routes. All traffic destined for the CIDR blocks in the prefix list is dropped.
	//
	// - **VPC**: The next hop of the CIDR blocks in the prefix list is a Virtual Private Cloud (VPC) connection.
	//
	// - **VBR**: The next hop of the CIDR blocks in the prefix list is a Virtual Border Router (VBR) connection.
	//
	// - **TR**: The next hop of the CIDR blocks in the prefix list is an inter-region connection.
	//
	// - **ECR**: The next hop of the CIDR blocks in the prefix list is an Express Connect Router (ECR) instance.
	//
	// example:
	//
	// VPC
	NextHopType  *string `json:"NextHopType,omitempty" xml:"NextHopType,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the Alibaba Cloud account that owns the prefix list.
	//
	// example:
	//
	// 1210123456123456
	OwnerUid *int64 `json:"OwnerUid,omitempty" xml:"OwnerUid,omitempty"`
	// The ID of the prefix list.
	//
	// This parameter is required.
	//
	// example:
	//
	// pl-6ehtn5kqxgeyy08fi****
	PrefixListId *string `json:"PrefixListId,omitempty" xml:"PrefixListId,omitempty"`
	// The region ID of the transit router instance.
	//
	// You can call the [DescribeChildInstanceRegions](https://help.aliyun.com/document_detail/132080.html) operation to query the region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the transit router instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// tr-6ehx7q2jze8ch5ji0****
	TransitRouterId *string `json:"TransitRouterId,omitempty" xml:"TransitRouterId,omitempty"`
	// The ID of the transit router route table.
	//
	// This parameter is required.
	//
	// example:
	//
	// vtb-6ehgc262hr170qgyc****
	TransitRouterTableId *string `json:"TransitRouterTableId,omitempty" xml:"TransitRouterTableId,omitempty"`
}

func (s CreateTransitRouterPrefixListAssociationRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateTransitRouterPrefixListAssociationRequest) GoString() string {
	return s.String()
}

func (s *CreateTransitRouterPrefixListAssociationRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateTransitRouterPrefixListAssociationRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *CreateTransitRouterPrefixListAssociationRequest) GetNextHop() *string {
	return s.NextHop
}

func (s *CreateTransitRouterPrefixListAssociationRequest) GetNextHopType() *string {
	return s.NextHopType
}

func (s *CreateTransitRouterPrefixListAssociationRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateTransitRouterPrefixListAssociationRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateTransitRouterPrefixListAssociationRequest) GetOwnerUid() *int64 {
	return s.OwnerUid
}

func (s *CreateTransitRouterPrefixListAssociationRequest) GetPrefixListId() *string {
	return s.PrefixListId
}

func (s *CreateTransitRouterPrefixListAssociationRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateTransitRouterPrefixListAssociationRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateTransitRouterPrefixListAssociationRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateTransitRouterPrefixListAssociationRequest) GetTransitRouterId() *string {
	return s.TransitRouterId
}

func (s *CreateTransitRouterPrefixListAssociationRequest) GetTransitRouterTableId() *string {
	return s.TransitRouterTableId
}

func (s *CreateTransitRouterPrefixListAssociationRequest) SetClientToken(v string) *CreateTransitRouterPrefixListAssociationRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateTransitRouterPrefixListAssociationRequest) SetDryRun(v bool) *CreateTransitRouterPrefixListAssociationRequest {
	s.DryRun = &v
	return s
}

func (s *CreateTransitRouterPrefixListAssociationRequest) SetNextHop(v string) *CreateTransitRouterPrefixListAssociationRequest {
	s.NextHop = &v
	return s
}

func (s *CreateTransitRouterPrefixListAssociationRequest) SetNextHopType(v string) *CreateTransitRouterPrefixListAssociationRequest {
	s.NextHopType = &v
	return s
}

func (s *CreateTransitRouterPrefixListAssociationRequest) SetOwnerAccount(v string) *CreateTransitRouterPrefixListAssociationRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateTransitRouterPrefixListAssociationRequest) SetOwnerId(v int64) *CreateTransitRouterPrefixListAssociationRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateTransitRouterPrefixListAssociationRequest) SetOwnerUid(v int64) *CreateTransitRouterPrefixListAssociationRequest {
	s.OwnerUid = &v
	return s
}

func (s *CreateTransitRouterPrefixListAssociationRequest) SetPrefixListId(v string) *CreateTransitRouterPrefixListAssociationRequest {
	s.PrefixListId = &v
	return s
}

func (s *CreateTransitRouterPrefixListAssociationRequest) SetRegionId(v string) *CreateTransitRouterPrefixListAssociationRequest {
	s.RegionId = &v
	return s
}

func (s *CreateTransitRouterPrefixListAssociationRequest) SetResourceOwnerAccount(v string) *CreateTransitRouterPrefixListAssociationRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateTransitRouterPrefixListAssociationRequest) SetResourceOwnerId(v int64) *CreateTransitRouterPrefixListAssociationRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateTransitRouterPrefixListAssociationRequest) SetTransitRouterId(v string) *CreateTransitRouterPrefixListAssociationRequest {
	s.TransitRouterId = &v
	return s
}

func (s *CreateTransitRouterPrefixListAssociationRequest) SetTransitRouterTableId(v string) *CreateTransitRouterPrefixListAssociationRequest {
	s.TransitRouterTableId = &v
	return s
}

func (s *CreateTransitRouterPrefixListAssociationRequest) Validate() error {
	return dara.Validate(s)
}
