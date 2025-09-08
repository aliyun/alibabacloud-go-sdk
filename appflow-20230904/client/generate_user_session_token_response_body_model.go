// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateUserSessionTokenResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GenerateUserSessionTokenResponseBody
	GetRequestId() *string
	SetUserSessionToken(v string) *GenerateUserSessionTokenResponseBody
	GetUserSessionToken() *string
}

type GenerateUserSessionTokenResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// 34C2713A-2270-5EEC-825E-115F9AD3BAC9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Token.
	//
	// example:
	//
	// 960f499au184m7****
	UserSessionToken *string `json:"UserSessionToken,omitempty" xml:"UserSessionToken,omitempty"`
}

func (s GenerateUserSessionTokenResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GenerateUserSessionTokenResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateUserSessionTokenResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GenerateUserSessionTokenResponseBody) GetUserSessionToken() *string {
	return s.UserSessionToken
}

func (s *GenerateUserSessionTokenResponseBody) SetRequestId(v string) *GenerateUserSessionTokenResponseBody {
	s.RequestId = &v
	return s
}

func (s *GenerateUserSessionTokenResponseBody) SetUserSessionToken(v string) *GenerateUserSessionTokenResponseBody {
	s.UserSessionToken = &v
	return s
}

func (s *GenerateUserSessionTokenResponseBody) Validate() error {
	return dara.Validate(s)
}
