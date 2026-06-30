// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTransitRouterCidrRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DeleteTransitRouterCidrRequest
	GetClientToken() *string
	SetDryRun(v bool) *DeleteTransitRouterCidrRequest
	GetDryRun() *bool
	SetOwnerAccount(v string) *DeleteTransitRouterCidrRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteTransitRouterCidrRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DeleteTransitRouterCidrRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DeleteTransitRouterCidrRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteTransitRouterCidrRequest
	GetResourceOwnerId() *int64
	SetTransitRouterCidrId(v string) *DeleteTransitRouterCidrRequest
	GetTransitRouterCidrId() *string
	SetTransitRouterId(v string) *DeleteTransitRouterCidrRequest
	GetTransitRouterId() *string
}

type DeleteTransitRouterCidrRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// Generate a client-side token that is unique among different requests. The token can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system uses the request ID as the client token. The request ID may be different for each request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform a dry run. Valid values:
	//
	// - **true**: sends a check request without deleting the CIDR block. The system checks whether the required parameters are specified, the request format is valid, and the service limits are met. If the request fails the check, an error message is returned. If the request passes the check, the `DryRunOperation` error code is returned.
	//
	// - **false*	- (default): sends a normal request. After the request passes the check, the CIDR block is deleted.
	//
	// example:
	//
	// false
	DryRun       *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region where the transit router is deployed.
	//
	// Call [DescribeChildInstanceRegions](https://help.aliyun.com/document_detail/132080.html) to obtain the region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the CIDR block of the transit router.
	//
	// Call [ListTransitRouterCidr](https://help.aliyun.com/document_detail/462772.html) to query the ID of the CIDR block.
	//
	// This parameter is required.
	//
	// example:
	//
	// cidr-0zv0q9crqpntzz****
	TransitRouterCidrId *string `json:"TransitRouterCidrId,omitempty" xml:"TransitRouterCidrId,omitempty"`
	// The ID of the transit router.
	//
	// This parameter is required.
	//
	// example:
	//
	// tr-p0w3x8c9em72a40nw****
	TransitRouterId *string `json:"TransitRouterId,omitempty" xml:"TransitRouterId,omitempty"`
}

func (s DeleteTransitRouterCidrRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteTransitRouterCidrRequest) GoString() string {
	return s.String()
}

func (s *DeleteTransitRouterCidrRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteTransitRouterCidrRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *DeleteTransitRouterCidrRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteTransitRouterCidrRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteTransitRouterCidrRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteTransitRouterCidrRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteTransitRouterCidrRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteTransitRouterCidrRequest) GetTransitRouterCidrId() *string {
	return s.TransitRouterCidrId
}

func (s *DeleteTransitRouterCidrRequest) GetTransitRouterId() *string {
	return s.TransitRouterId
}

func (s *DeleteTransitRouterCidrRequest) SetClientToken(v string) *DeleteTransitRouterCidrRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteTransitRouterCidrRequest) SetDryRun(v bool) *DeleteTransitRouterCidrRequest {
	s.DryRun = &v
	return s
}

func (s *DeleteTransitRouterCidrRequest) SetOwnerAccount(v string) *DeleteTransitRouterCidrRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteTransitRouterCidrRequest) SetOwnerId(v int64) *DeleteTransitRouterCidrRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteTransitRouterCidrRequest) SetRegionId(v string) *DeleteTransitRouterCidrRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteTransitRouterCidrRequest) SetResourceOwnerAccount(v string) *DeleteTransitRouterCidrRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteTransitRouterCidrRequest) SetResourceOwnerId(v int64) *DeleteTransitRouterCidrRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteTransitRouterCidrRequest) SetTransitRouterCidrId(v string) *DeleteTransitRouterCidrRequest {
	s.TransitRouterCidrId = &v
	return s
}

func (s *DeleteTransitRouterCidrRequest) SetTransitRouterId(v string) *DeleteTransitRouterCidrRequest {
	s.TransitRouterId = &v
	return s
}

func (s *DeleteTransitRouterCidrRequest) Validate() error {
	return dara.Validate(s)
}
