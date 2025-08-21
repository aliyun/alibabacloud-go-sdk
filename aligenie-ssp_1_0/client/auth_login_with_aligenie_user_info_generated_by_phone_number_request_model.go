// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAuthLoginWithAligenieUserInfoGeneratedByPhoneNumberRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSessionId(v string) *AuthLoginWithAligenieUserInfoGeneratedByPhoneNumberRequest
	GetSessionId() *string
}

type AuthLoginWithAligenieUserInfoGeneratedByPhoneNumberRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// dbe2eb4458302b9246c6da17fbc95f4b
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
}

func (s AuthLoginWithAligenieUserInfoGeneratedByPhoneNumberRequest) String() string {
	return dara.Prettify(s)
}

func (s AuthLoginWithAligenieUserInfoGeneratedByPhoneNumberRequest) GoString() string {
	return s.String()
}

func (s *AuthLoginWithAligenieUserInfoGeneratedByPhoneNumberRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *AuthLoginWithAligenieUserInfoGeneratedByPhoneNumberRequest) SetSessionId(v string) *AuthLoginWithAligenieUserInfoGeneratedByPhoneNumberRequest {
	s.SessionId = &v
	return s
}

func (s *AuthLoginWithAligenieUserInfoGeneratedByPhoneNumberRequest) Validate() error {
	return dara.Validate(s)
}
