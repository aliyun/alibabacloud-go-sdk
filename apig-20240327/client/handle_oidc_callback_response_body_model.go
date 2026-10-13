// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHandleOidcCallbackResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *HandleOidcCallbackResponseBody
	GetCode() *string
	SetData(v *HandleOidcCallbackResponseBodyData) *HandleOidcCallbackResponseBody
	GetData() *HandleOidcCallbackResponseBodyData
	SetMessage(v string) *HandleOidcCallbackResponseBody
	GetMessage() *string
	SetRequestId(v string) *HandleOidcCallbackResponseBody
	GetRequestId() *string
}

type HandleOidcCallbackResponseBody struct {
	// The response status code.
	//
	// example:
	//
	// Ok
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The response data.
	Data *HandleOidcCallbackResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The response message.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 393E2630-DBE7-5221-AB35-9E740675491A
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s HandleOidcCallbackResponseBody) String() string {
	return dara.Prettify(s)
}

func (s HandleOidcCallbackResponseBody) GoString() string {
	return s.String()
}

func (s *HandleOidcCallbackResponseBody) GetCode() *string {
	return s.Code
}

func (s *HandleOidcCallbackResponseBody) GetData() *HandleOidcCallbackResponseBodyData {
	return s.Data
}

func (s *HandleOidcCallbackResponseBody) GetMessage() *string {
	return s.Message
}

func (s *HandleOidcCallbackResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *HandleOidcCallbackResponseBody) SetCode(v string) *HandleOidcCallbackResponseBody {
	s.Code = &v
	return s
}

func (s *HandleOidcCallbackResponseBody) SetData(v *HandleOidcCallbackResponseBodyData) *HandleOidcCallbackResponseBody {
	s.Data = v
	return s
}

func (s *HandleOidcCallbackResponseBody) SetMessage(v string) *HandleOidcCallbackResponseBody {
	s.Message = &v
	return s
}

func (s *HandleOidcCallbackResponseBody) SetRequestId(v string) *HandleOidcCallbackResponseBody {
	s.RequestId = &v
	return s
}

func (s *HandleOidcCallbackResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type HandleOidcCallbackResponseBodyData struct {
	// The access token returned for the current OIDC logon. The API definition marks this field as sensitive.
	//
	// example:
	//
	// access-token-xxxxxxxx
	AccessToken *string `json:"accessToken,omitempty" xml:"accessToken,omitempty"`
	// The validity period value of the access token. The unit is not specified in the API definition.
	//
	// example:
	//
	// 3600
	ExpiresIn *int64 `json:"expiresIn,omitempty" xml:"expiresIn,omitempty"`
	// The type identifier of the access token. The valid values are not specified in the API definition.
	//
	// example:
	//
	// token-type-xxxxxxxx
	TokenType *string `json:"tokenType,omitempty" xml:"tokenType,omitempty"`
}

func (s HandleOidcCallbackResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s HandleOidcCallbackResponseBodyData) GoString() string {
	return s.String()
}

func (s *HandleOidcCallbackResponseBodyData) GetAccessToken() *string {
	return s.AccessToken
}

func (s *HandleOidcCallbackResponseBodyData) GetExpiresIn() *int64 {
	return s.ExpiresIn
}

func (s *HandleOidcCallbackResponseBodyData) GetTokenType() *string {
	return s.TokenType
}

func (s *HandleOidcCallbackResponseBodyData) SetAccessToken(v string) *HandleOidcCallbackResponseBodyData {
	s.AccessToken = &v
	return s
}

func (s *HandleOidcCallbackResponseBodyData) SetExpiresIn(v int64) *HandleOidcCallbackResponseBodyData {
	s.ExpiresIn = &v
	return s
}

func (s *HandleOidcCallbackResponseBodyData) SetTokenType(v string) *HandleOidcCallbackResponseBodyData {
	s.TokenType = &v
	return s
}

func (s *HandleOidcCallbackResponseBodyData) Validate() error {
	return dara.Validate(s)
}
