// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRefreshTransitRouteTableAggregationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *RefreshTransitRouteTableAggregationRequest
	GetClientToken() *string
	SetOwnerAccount(v string) *RefreshTransitRouteTableAggregationRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *RefreshTransitRouteTableAggregationRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *RefreshTransitRouteTableAggregationRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *RefreshTransitRouteTableAggregationRequest
	GetResourceOwnerId() *int64
	SetTransitRouteTableAggregationCidr(v string) *RefreshTransitRouteTableAggregationRequest
	GetTransitRouteTableAggregationCidr() *string
	SetTransitRouteTableId(v string) *RefreshTransitRouteTableAggregationRequest
	GetTransitRouteTableId() *string
}

type RefreshTransitRouteTableAggregationRequest struct {
	// example:
	//
	// 02fb3da4-130e-11e9-8e44-001****
	ClientToken          *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 192.168.53.0/24
	TransitRouteTableAggregationCidr *string `json:"TransitRouteTableAggregationCidr,omitempty" xml:"TransitRouteTableAggregationCidr,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// vtb-iq8qgruq1ry8jc7vt****
	TransitRouteTableId *string `json:"TransitRouteTableId,omitempty" xml:"TransitRouteTableId,omitempty"`
}

func (s RefreshTransitRouteTableAggregationRequest) String() string {
	return dara.Prettify(s)
}

func (s RefreshTransitRouteTableAggregationRequest) GoString() string {
	return s.String()
}

func (s *RefreshTransitRouteTableAggregationRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *RefreshTransitRouteTableAggregationRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *RefreshTransitRouteTableAggregationRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *RefreshTransitRouteTableAggregationRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *RefreshTransitRouteTableAggregationRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *RefreshTransitRouteTableAggregationRequest) GetTransitRouteTableAggregationCidr() *string {
	return s.TransitRouteTableAggregationCidr
}

func (s *RefreshTransitRouteTableAggregationRequest) GetTransitRouteTableId() *string {
	return s.TransitRouteTableId
}

func (s *RefreshTransitRouteTableAggregationRequest) SetClientToken(v string) *RefreshTransitRouteTableAggregationRequest {
	s.ClientToken = &v
	return s
}

func (s *RefreshTransitRouteTableAggregationRequest) SetOwnerAccount(v string) *RefreshTransitRouteTableAggregationRequest {
	s.OwnerAccount = &v
	return s
}

func (s *RefreshTransitRouteTableAggregationRequest) SetOwnerId(v int64) *RefreshTransitRouteTableAggregationRequest {
	s.OwnerId = &v
	return s
}

func (s *RefreshTransitRouteTableAggregationRequest) SetResourceOwnerAccount(v string) *RefreshTransitRouteTableAggregationRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *RefreshTransitRouteTableAggregationRequest) SetResourceOwnerId(v int64) *RefreshTransitRouteTableAggregationRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *RefreshTransitRouteTableAggregationRequest) SetTransitRouteTableAggregationCidr(v string) *RefreshTransitRouteTableAggregationRequest {
	s.TransitRouteTableAggregationCidr = &v
	return s
}

func (s *RefreshTransitRouteTableAggregationRequest) SetTransitRouteTableId(v string) *RefreshTransitRouteTableAggregationRequest {
	s.TransitRouteTableId = &v
	return s
}

func (s *RefreshTransitRouteTableAggregationRequest) Validate() error {
	return dara.Validate(s)
}
