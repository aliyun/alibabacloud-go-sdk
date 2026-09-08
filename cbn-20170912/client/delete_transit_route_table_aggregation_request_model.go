// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTransitRouteTableAggregationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DeleteTransitRouteTableAggregationRequest
	GetClientToken() *string
	SetDryRun(v bool) *DeleteTransitRouteTableAggregationRequest
	GetDryRun() *bool
	SetOwnerAccount(v string) *DeleteTransitRouteTableAggregationRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteTransitRouteTableAggregationRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DeleteTransitRouteTableAggregationRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteTransitRouteTableAggregationRequest
	GetResourceOwnerId() *int64
	SetTransitRouteTableAggregationCidr(v string) *DeleteTransitRouteTableAggregationRequest
	GetTransitRouteTableAggregationCidr() *string
	SetTransitRouteTableId(v string) *DeleteTransitRouteTableAggregationRequest
	GetTransitRouteTableId() *string
}

type DeleteTransitRouteTableAggregationRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The client token can contain only ASCII characters.
	//
	// >If you do not specify this parameter, the system automatically uses the RequestId of the API request as the ClientToken. The RequestId may be different for each API request.
	//
	// example:
	//
	// 02fb3da4-130e-11e9-8e44-001****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform a dry run for this deletion request, including permission and instance status verification. Valid values:
	//
	// - **false*	- (default): Sends a normal request. If the request passes the check, the aggregate route is deleted.
	//
	// - **true**: Sends a check request. Only verification is performed without actually deleting the aggregate route. The check items include whether required parameters are specified and the request format. If the check fails, the corresponding error is returned. If the check passes, the error code `DryRunOperation` is returned.
	//
	// example:
	//
	// false
	DryRun               *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The destination CIDR block of the aggregate route.
	//
	// This parameter is required.
	//
	// example:
	//
	// 192.168.10.0/24
	TransitRouteTableAggregationCidr *string `json:"TransitRouteTableAggregationCidr,omitempty" xml:"TransitRouteTableAggregationCidr,omitempty"`
	// The ID of the Enterprise Edition transit router route table.
	//
	// This parameter is required.
	//
	// example:
	//
	// vtb-6ehgc262hr170qgyc****
	TransitRouteTableId *string `json:"TransitRouteTableId,omitempty" xml:"TransitRouteTableId,omitempty"`
}

func (s DeleteTransitRouteTableAggregationRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteTransitRouteTableAggregationRequest) GoString() string {
	return s.String()
}

func (s *DeleteTransitRouteTableAggregationRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteTransitRouteTableAggregationRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *DeleteTransitRouteTableAggregationRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteTransitRouteTableAggregationRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteTransitRouteTableAggregationRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteTransitRouteTableAggregationRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteTransitRouteTableAggregationRequest) GetTransitRouteTableAggregationCidr() *string {
	return s.TransitRouteTableAggregationCidr
}

func (s *DeleteTransitRouteTableAggregationRequest) GetTransitRouteTableId() *string {
	return s.TransitRouteTableId
}

func (s *DeleteTransitRouteTableAggregationRequest) SetClientToken(v string) *DeleteTransitRouteTableAggregationRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteTransitRouteTableAggregationRequest) SetDryRun(v bool) *DeleteTransitRouteTableAggregationRequest {
	s.DryRun = &v
	return s
}

func (s *DeleteTransitRouteTableAggregationRequest) SetOwnerAccount(v string) *DeleteTransitRouteTableAggregationRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteTransitRouteTableAggregationRequest) SetOwnerId(v int64) *DeleteTransitRouteTableAggregationRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteTransitRouteTableAggregationRequest) SetResourceOwnerAccount(v string) *DeleteTransitRouteTableAggregationRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteTransitRouteTableAggregationRequest) SetResourceOwnerId(v int64) *DeleteTransitRouteTableAggregationRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteTransitRouteTableAggregationRequest) SetTransitRouteTableAggregationCidr(v string) *DeleteTransitRouteTableAggregationRequest {
	s.TransitRouteTableAggregationCidr = &v
	return s
}

func (s *DeleteTransitRouteTableAggregationRequest) SetTransitRouteTableId(v string) *DeleteTransitRouteTableAggregationRequest {
	s.TransitRouteTableId = &v
	return s
}

func (s *DeleteTransitRouteTableAggregationRequest) Validate() error {
	return dara.Validate(s)
}
