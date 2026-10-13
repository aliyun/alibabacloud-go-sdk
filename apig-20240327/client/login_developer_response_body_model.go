// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLoginDeveloperResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *LoginDeveloperResponseBody
	GetCode() *string
	SetData(v *LoginDeveloperResponseBodyData) *LoginDeveloperResponseBody
	GetData() *LoginDeveloperResponseBodyData
	SetMessage(v string) *LoginDeveloperResponseBody
	GetMessage() *string
	SetRequestId(v string) *LoginDeveloperResponseBody
	GetRequestId() *string
}

type LoginDeveloperResponseBody struct {
	// The response status code.
	//
	// example:
	//
	// Ok
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The response data.
	Data *LoginDeveloperResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
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

func (s LoginDeveloperResponseBody) String() string {
	return dara.Prettify(s)
}

func (s LoginDeveloperResponseBody) GoString() string {
	return s.String()
}

func (s *LoginDeveloperResponseBody) GetCode() *string {
	return s.Code
}

func (s *LoginDeveloperResponseBody) GetData() *LoginDeveloperResponseBodyData {
	return s.Data
}

func (s *LoginDeveloperResponseBody) GetMessage() *string {
	return s.Message
}

func (s *LoginDeveloperResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *LoginDeveloperResponseBody) SetCode(v string) *LoginDeveloperResponseBody {
	s.Code = &v
	return s
}

func (s *LoginDeveloperResponseBody) SetData(v *LoginDeveloperResponseBodyData) *LoginDeveloperResponseBody {
	s.Data = v
	return s
}

func (s *LoginDeveloperResponseBody) SetMessage(v string) *LoginDeveloperResponseBody {
	s.Message = &v
	return s
}

func (s *LoginDeveloperResponseBody) SetRequestId(v string) *LoginDeveloperResponseBody {
	s.RequestId = &v
	return s
}

func (s *LoginDeveloperResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type LoginDeveloperResponseBodyData struct {
	// The access token returned for the current developer logon. This field is marked as sensitive in the API definition.
	//
	// example:
	//
	// access-token-xxxxxxxx
	AccessToken *string `json:"accessToken,omitempty" xml:"accessToken,omitempty"`
	// The int64 numeric field associated with the access token. The unit and exact meaning are not declared in the API definition.
	//
	// example:
	//
	// 3600
	ExpiresIn *int64 `json:"expiresIn,omitempty" xml:"expiresIn,omitempty"`
}

func (s LoginDeveloperResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s LoginDeveloperResponseBodyData) GoString() string {
	return s.String()
}

func (s *LoginDeveloperResponseBodyData) GetAccessToken() *string {
	return s.AccessToken
}

func (s *LoginDeveloperResponseBodyData) GetExpiresIn() *int64 {
	return s.ExpiresIn
}

func (s *LoginDeveloperResponseBodyData) SetAccessToken(v string) *LoginDeveloperResponseBodyData {
	s.AccessToken = &v
	return s
}

func (s *LoginDeveloperResponseBodyData) SetExpiresIn(v int64) *LoginDeveloperResponseBodyData {
	s.ExpiresIn = &v
	return s
}

func (s *LoginDeveloperResponseBodyData) Validate() error {
	return dara.Validate(s)
}
