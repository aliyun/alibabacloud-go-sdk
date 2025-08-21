// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAuthLoginWithTaobaoUserInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEncryptedTaobaoUserIdentifier(v string) *AuthLoginWithTaobaoUserInfoRequest
	GetEncryptedTaobaoUserIdentifier() *string
	SetSessionId(v string) *AuthLoginWithTaobaoUserInfoRequest
	GetSessionId() *string
}

type AuthLoginWithTaobaoUserInfoRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// KsVgypxAipf+xNECMZV2ONMcheqiIoEGFvgx+T8s1oV6/euTK9+ImYvLVPsSqFDh
	EncryptedTaobaoUserIdentifier *string `json:"EncryptedTaobaoUserIdentifier,omitempty" xml:"EncryptedTaobaoUserIdentifier,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// dbe2eb4458302b9246c6da17fbc95f4b
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
}

func (s AuthLoginWithTaobaoUserInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s AuthLoginWithTaobaoUserInfoRequest) GoString() string {
	return s.String()
}

func (s *AuthLoginWithTaobaoUserInfoRequest) GetEncryptedTaobaoUserIdentifier() *string {
	return s.EncryptedTaobaoUserIdentifier
}

func (s *AuthLoginWithTaobaoUserInfoRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *AuthLoginWithTaobaoUserInfoRequest) SetEncryptedTaobaoUserIdentifier(v string) *AuthLoginWithTaobaoUserInfoRequest {
	s.EncryptedTaobaoUserIdentifier = &v
	return s
}

func (s *AuthLoginWithTaobaoUserInfoRequest) SetSessionId(v string) *AuthLoginWithTaobaoUserInfoRequest {
	s.SessionId = &v
	return s
}

func (s *AuthLoginWithTaobaoUserInfoRequest) Validate() error {
	return dara.Validate(s)
}
