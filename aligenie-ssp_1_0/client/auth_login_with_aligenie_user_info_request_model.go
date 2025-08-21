// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAuthLoginWithAligenieUserInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEncryptedAligenieUserIdentifier(v string) *AuthLoginWithAligenieUserInfoRequest
	GetEncryptedAligenieUserIdentifier() *string
	SetSessionId(v string) *AuthLoginWithAligenieUserInfoRequest
	GetSessionId() *string
}

type AuthLoginWithAligenieUserInfoRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// UYugfm/3Nb9q24AyES2rYmC5tIglSoDX3Mbna/vrldcjGPtC8VzFwo+CU5c4CHLjrK7ekskG2WVaevM5Zi9f0w==
	EncryptedAligenieUserIdentifier *string `json:"EncryptedAligenieUserIdentifier,omitempty" xml:"EncryptedAligenieUserIdentifier,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// dbe2eb4458302b9246c6da17fbc95f4b
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
}

func (s AuthLoginWithAligenieUserInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s AuthLoginWithAligenieUserInfoRequest) GoString() string {
	return s.String()
}

func (s *AuthLoginWithAligenieUserInfoRequest) GetEncryptedAligenieUserIdentifier() *string {
	return s.EncryptedAligenieUserIdentifier
}

func (s *AuthLoginWithAligenieUserInfoRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *AuthLoginWithAligenieUserInfoRequest) SetEncryptedAligenieUserIdentifier(v string) *AuthLoginWithAligenieUserInfoRequest {
	s.EncryptedAligenieUserIdentifier = &v
	return s
}

func (s *AuthLoginWithAligenieUserInfoRequest) SetSessionId(v string) *AuthLoginWithAligenieUserInfoRequest {
	s.SessionId = &v
	return s
}

func (s *AuthLoginWithAligenieUserInfoRequest) Validate() error {
	return dara.Validate(s)
}
