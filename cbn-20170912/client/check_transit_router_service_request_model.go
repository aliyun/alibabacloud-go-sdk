// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckTransitRouterServiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CheckTransitRouterServiceRequest
	GetClientToken() *string
	SetOwnerAccount(v string) *CheckTransitRouterServiceRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CheckTransitRouterServiceRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *CheckTransitRouterServiceRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CheckTransitRouterServiceRequest
	GetResourceOwnerId() *int64
}

type CheckTransitRouterServiceRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate a token, but you must make sure that it is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
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

func (s CheckTransitRouterServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s CheckTransitRouterServiceRequest) GoString() string {
	return s.String()
}

func (s *CheckTransitRouterServiceRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CheckTransitRouterServiceRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CheckTransitRouterServiceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CheckTransitRouterServiceRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CheckTransitRouterServiceRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CheckTransitRouterServiceRequest) SetClientToken(v string) *CheckTransitRouterServiceRequest {
	s.ClientToken = &v
	return s
}

func (s *CheckTransitRouterServiceRequest) SetOwnerAccount(v string) *CheckTransitRouterServiceRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CheckTransitRouterServiceRequest) SetOwnerId(v int64) *CheckTransitRouterServiceRequest {
	s.OwnerId = &v
	return s
}

func (s *CheckTransitRouterServiceRequest) SetResourceOwnerAccount(v string) *CheckTransitRouterServiceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CheckTransitRouterServiceRequest) SetResourceOwnerId(v int64) *CheckTransitRouterServiceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CheckTransitRouterServiceRequest) Validate() error {
	return dara.Validate(s)
}
