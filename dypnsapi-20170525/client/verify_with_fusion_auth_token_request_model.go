// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVerifyWithFusionAuthTokenRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *VerifyWithFusionAuthTokenRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *VerifyWithFusionAuthTokenRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *VerifyWithFusionAuthTokenRequest
	GetResourceOwnerId() *int64
	SetVerifyToken(v string) *VerifyWithFusionAuthTokenRequest
	GetVerifyToken() *string
}

type VerifyWithFusionAuthTokenRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The unified verification token that is returned by the client SDKs.
	//
	// This parameter is required.
	//
	// example:
	//
	// LD108enNdlsl*******sFLKCks1==
	VerifyToken *string `json:"VerifyToken,omitempty" xml:"VerifyToken,omitempty"`
}

func (s VerifyWithFusionAuthTokenRequest) String() string {
	return dara.Prettify(s)
}

func (s VerifyWithFusionAuthTokenRequest) GoString() string {
	return s.String()
}

func (s *VerifyWithFusionAuthTokenRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *VerifyWithFusionAuthTokenRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *VerifyWithFusionAuthTokenRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *VerifyWithFusionAuthTokenRequest) GetVerifyToken() *string {
	return s.VerifyToken
}

func (s *VerifyWithFusionAuthTokenRequest) SetOwnerId(v int64) *VerifyWithFusionAuthTokenRequest {
	s.OwnerId = &v
	return s
}

func (s *VerifyWithFusionAuthTokenRequest) SetResourceOwnerAccount(v string) *VerifyWithFusionAuthTokenRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *VerifyWithFusionAuthTokenRequest) SetResourceOwnerId(v int64) *VerifyWithFusionAuthTokenRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *VerifyWithFusionAuthTokenRequest) SetVerifyToken(v string) *VerifyWithFusionAuthTokenRequest {
	s.VerifyToken = &v
	return s
}

func (s *VerifyWithFusionAuthTokenRequest) Validate() error {
	return dara.Validate(s)
}
