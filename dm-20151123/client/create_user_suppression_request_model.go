// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateUserSuppressionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddress(v string) *CreateUserSuppressionRequest
	GetAddress() *string
	SetOwnerId(v int64) *CreateUserSuppressionRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *CreateUserSuppressionRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateUserSuppressionRequest
	GetResourceOwnerId() *int64
}

type CreateUserSuppressionRequest struct {
	// Email address or domain name
	//
	// example:
	//
	// test@example.net
	Address              *string `json:"Address,omitempty" xml:"Address,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s CreateUserSuppressionRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateUserSuppressionRequest) GoString() string {
	return s.String()
}

func (s *CreateUserSuppressionRequest) GetAddress() *string {
	return s.Address
}

func (s *CreateUserSuppressionRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateUserSuppressionRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateUserSuppressionRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateUserSuppressionRequest) SetAddress(v string) *CreateUserSuppressionRequest {
	s.Address = &v
	return s
}

func (s *CreateUserSuppressionRequest) SetOwnerId(v int64) *CreateUserSuppressionRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateUserSuppressionRequest) SetResourceOwnerAccount(v string) *CreateUserSuppressionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateUserSuppressionRequest) SetResourceOwnerId(v int64) *CreateUserSuppressionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateUserSuppressionRequest) Validate() error {
	return dara.Validate(s)
}
