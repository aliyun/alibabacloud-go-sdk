// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenTransitRouterServiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *OpenTransitRouterServiceRequest
	GetClientToken() *string
	SetOwnerAccount(v string) *OpenTransitRouterServiceRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *OpenTransitRouterServiceRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *OpenTransitRouterServiceRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *OpenTransitRouterServiceRequest
	GetResourceOwnerId() *int64
}

type OpenTransitRouterServiceRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the value, but you must make sure that it is unique among different requests. The client token can contain only ASCII characters.
	//
	// example:
	//
	// 02fb3da4-130e-11e9-8e44-001****
	ClientToken          *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s OpenTransitRouterServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s OpenTransitRouterServiceRequest) GoString() string {
	return s.String()
}

func (s *OpenTransitRouterServiceRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *OpenTransitRouterServiceRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *OpenTransitRouterServiceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *OpenTransitRouterServiceRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *OpenTransitRouterServiceRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *OpenTransitRouterServiceRequest) SetClientToken(v string) *OpenTransitRouterServiceRequest {
	s.ClientToken = &v
	return s
}

func (s *OpenTransitRouterServiceRequest) SetOwnerAccount(v string) *OpenTransitRouterServiceRequest {
	s.OwnerAccount = &v
	return s
}

func (s *OpenTransitRouterServiceRequest) SetOwnerId(v int64) *OpenTransitRouterServiceRequest {
	s.OwnerId = &v
	return s
}

func (s *OpenTransitRouterServiceRequest) SetResourceOwnerAccount(v string) *OpenTransitRouterServiceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *OpenTransitRouterServiceRequest) SetResourceOwnerId(v int64) *OpenTransitRouterServiceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *OpenTransitRouterServiceRequest) Validate() error {
	return dara.Validate(s)
}
