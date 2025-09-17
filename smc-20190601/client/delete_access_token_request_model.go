// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAccessTokenRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessTokenId(v string) *DeleteAccessTokenRequest
	GetAccessTokenId() *string
	SetOwnerId(v int64) *DeleteAccessTokenRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DeleteAccessTokenRequest
	GetResourceOwnerAccount() *string
}

type DeleteAccessTokenRequest struct {
	// The ID of the activation code.
	//
	// This parameter is required.
	//
	// example:
	//
	// at-bp1akz2zp67r0k6r****
	AccessTokenId        *string `json:"AccessTokenId,omitempty" xml:"AccessTokenId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
}

func (s DeleteAccessTokenRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAccessTokenRequest) GoString() string {
	return s.String()
}

func (s *DeleteAccessTokenRequest) GetAccessTokenId() *string {
	return s.AccessTokenId
}

func (s *DeleteAccessTokenRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteAccessTokenRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteAccessTokenRequest) SetAccessTokenId(v string) *DeleteAccessTokenRequest {
	s.AccessTokenId = &v
	return s
}

func (s *DeleteAccessTokenRequest) SetOwnerId(v int64) *DeleteAccessTokenRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteAccessTokenRequest) SetResourceOwnerAccount(v string) *DeleteAccessTokenRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteAccessTokenRequest) Validate() error {
	return dara.Validate(s)
}
