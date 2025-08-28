// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTokenRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *GetTokenRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *GetTokenRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *GetTokenRequest
	GetResourceOwnerId() *int64
	SetTokenType(v string) *GetTokenRequest
	GetTokenType() *string
}

type GetTokenRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The token type.
	//
	// example:
	//
	// dyvms
	TokenType *string `json:"TokenType,omitempty" xml:"TokenType,omitempty"`
}

func (s GetTokenRequest) String() string {
	return dara.Prettify(s)
}

func (s GetTokenRequest) GoString() string {
	return s.String()
}

func (s *GetTokenRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *GetTokenRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *GetTokenRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *GetTokenRequest) GetTokenType() *string {
	return s.TokenType
}

func (s *GetTokenRequest) SetOwnerId(v int64) *GetTokenRequest {
	s.OwnerId = &v
	return s
}

func (s *GetTokenRequest) SetResourceOwnerAccount(v string) *GetTokenRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetTokenRequest) SetResourceOwnerId(v int64) *GetTokenRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetTokenRequest) SetTokenType(v string) *GetTokenRequest {
	s.TokenType = &v
	return s
}

func (s *GetTokenRequest) Validate() error {
	return dara.Validate(s)
}
