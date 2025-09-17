// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableAccessTokenRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessTokenId(v string) *DisableAccessTokenRequest
	GetAccessTokenId() *string
	SetOwnerId(v int64) *DisableAccessTokenRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DisableAccessTokenRequest
	GetResourceOwnerAccount() *string
}

type DisableAccessTokenRequest struct {
	// The ID of the activation code.
	//
	// This parameter is required.
	//
	// example:
	//
	// at-bp12g5gwup0yzmce****
	AccessTokenId        *string `json:"AccessTokenId,omitempty" xml:"AccessTokenId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
}

func (s DisableAccessTokenRequest) String() string {
	return dara.Prettify(s)
}

func (s DisableAccessTokenRequest) GoString() string {
	return s.String()
}

func (s *DisableAccessTokenRequest) GetAccessTokenId() *string {
	return s.AccessTokenId
}

func (s *DisableAccessTokenRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DisableAccessTokenRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DisableAccessTokenRequest) SetAccessTokenId(v string) *DisableAccessTokenRequest {
	s.AccessTokenId = &v
	return s
}

func (s *DisableAccessTokenRequest) SetOwnerId(v int64) *DisableAccessTokenRequest {
	s.OwnerId = &v
	return s
}

func (s *DisableAccessTokenRequest) SetResourceOwnerAccount(v string) *DisableAccessTokenRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DisableAccessTokenRequest) Validate() error {
	return dara.Validate(s)
}
