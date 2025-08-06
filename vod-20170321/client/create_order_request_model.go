// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOrderRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *CreateOrderRequest
	GetOwnerAccount() *string
	SetOwnerId(v string) *CreateOrderRequest
	GetOwnerId() *string
	SetResourceOwnerAccount(v string) *CreateOrderRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v string) *CreateOrderRequest
	GetResourceOwnerId() *string
}

type CreateOrderRequest struct {
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *string `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s CreateOrderRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateOrderRequest) GoString() string {
	return s.String()
}

func (s *CreateOrderRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateOrderRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *CreateOrderRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateOrderRequest) GetResourceOwnerId() *string {
	return s.ResourceOwnerId
}

func (s *CreateOrderRequest) SetOwnerAccount(v string) *CreateOrderRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateOrderRequest) SetOwnerId(v string) *CreateOrderRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateOrderRequest) SetResourceOwnerAccount(v string) *CreateOrderRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateOrderRequest) SetResourceOwnerId(v string) *CreateOrderRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateOrderRequest) Validate() error {
	return dara.Validate(s)
}
