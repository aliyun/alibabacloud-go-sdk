// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTransitRouterCidrRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *ListTransitRouterCidrRequest
	GetClientToken() *string
	SetDryRun(v bool) *ListTransitRouterCidrRequest
	GetDryRun() *bool
	SetOwnerAccount(v string) *ListTransitRouterCidrRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ListTransitRouterCidrRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ListTransitRouterCidrRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ListTransitRouterCidrRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ListTransitRouterCidrRequest
	GetResourceOwnerId() *int64
	SetTransitRouterCidrId(v string) *ListTransitRouterCidrRequest
	GetTransitRouterCidrId() *string
	SetTransitRouterId(v string) *ListTransitRouterCidrRequest
	GetTransitRouterId() *string
}

type ListTransitRouterCidrRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the request ID as the client token. The request ID may be different for each request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether only to precheck the API request. Valid values:
	//
	// 	- **true**: prechecks the request but does not query the CIDR block. The system checks the required parameters, the request format, and the service limits. If the request fails the check, an error message is returned. If the request passes the check, the `DryRunOperation` error code is returned.
	//
	// 	- **false*	- (default): sends the request. After the request passes the check, the operation is performed.
	//
	// example:
	//
	// false
	DryRun       *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the transit router.
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
	// The ID of the CIDR block.
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

func (s ListTransitRouterCidrRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTransitRouterCidrRequest) GoString() string {
	return s.String()
}

func (s *ListTransitRouterCidrRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ListTransitRouterCidrRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *ListTransitRouterCidrRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ListTransitRouterCidrRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListTransitRouterCidrRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListTransitRouterCidrRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ListTransitRouterCidrRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ListTransitRouterCidrRequest) GetTransitRouterCidrId() *string {
	return s.TransitRouterCidrId
}

func (s *ListTransitRouterCidrRequest) GetTransitRouterId() *string {
	return s.TransitRouterId
}

func (s *ListTransitRouterCidrRequest) SetClientToken(v string) *ListTransitRouterCidrRequest {
	s.ClientToken = &v
	return s
}

func (s *ListTransitRouterCidrRequest) SetDryRun(v bool) *ListTransitRouterCidrRequest {
	s.DryRun = &v
	return s
}

func (s *ListTransitRouterCidrRequest) SetOwnerAccount(v string) *ListTransitRouterCidrRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ListTransitRouterCidrRequest) SetOwnerId(v int64) *ListTransitRouterCidrRequest {
	s.OwnerId = &v
	return s
}

func (s *ListTransitRouterCidrRequest) SetRegionId(v string) *ListTransitRouterCidrRequest {
	s.RegionId = &v
	return s
}

func (s *ListTransitRouterCidrRequest) SetResourceOwnerAccount(v string) *ListTransitRouterCidrRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListTransitRouterCidrRequest) SetResourceOwnerId(v int64) *ListTransitRouterCidrRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListTransitRouterCidrRequest) SetTransitRouterCidrId(v string) *ListTransitRouterCidrRequest {
	s.TransitRouterCidrId = &v
	return s
}

func (s *ListTransitRouterCidrRequest) SetTransitRouterId(v string) *ListTransitRouterCidrRequest {
	s.TransitRouterId = &v
	return s
}

func (s *ListTransitRouterCidrRequest) Validate() error {
	return dara.Validate(s)
}
