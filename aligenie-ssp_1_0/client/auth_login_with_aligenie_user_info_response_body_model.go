// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAuthLoginWithAligenieUserInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *AuthLoginWithAligenieUserInfoResponseBody
	GetCode() *int32
	SetMessage(v string) *AuthLoginWithAligenieUserInfoResponseBody
	GetMessage() *string
	SetRequestId(v string) *AuthLoginWithAligenieUserInfoResponseBody
	GetRequestId() *string
	SetResult(v *AuthLoginWithAligenieUserInfoResponseBodyResult) *AuthLoginWithAligenieUserInfoResponseBody
	GetResult() *AuthLoginWithAligenieUserInfoResponseBodyResult
	SetSuccess(v bool) *AuthLoginWithAligenieUserInfoResponseBody
	GetSuccess() *bool
}

type AuthLoginWithAligenieUserInfoResponseBody struct {
	// Response code
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// Response message
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID
	//
	// example:
	//
	// 73C67BD9-175A-1324-8202-9FAABBB3E6FA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Response Result
	Result *AuthLoginWithAligenieUserInfoResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// Flag indicating whether the invocation succeeded
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AuthLoginWithAligenieUserInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AuthLoginWithAligenieUserInfoResponseBody) GoString() string {
	return s.String()
}

func (s *AuthLoginWithAligenieUserInfoResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *AuthLoginWithAligenieUserInfoResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AuthLoginWithAligenieUserInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AuthLoginWithAligenieUserInfoResponseBody) GetResult() *AuthLoginWithAligenieUserInfoResponseBodyResult {
	return s.Result
}

func (s *AuthLoginWithAligenieUserInfoResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AuthLoginWithAligenieUserInfoResponseBody) SetCode(v int32) *AuthLoginWithAligenieUserInfoResponseBody {
	s.Code = &v
	return s
}

func (s *AuthLoginWithAligenieUserInfoResponseBody) SetMessage(v string) *AuthLoginWithAligenieUserInfoResponseBody {
	s.Message = &v
	return s
}

func (s *AuthLoginWithAligenieUserInfoResponseBody) SetRequestId(v string) *AuthLoginWithAligenieUserInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *AuthLoginWithAligenieUserInfoResponseBody) SetResult(v *AuthLoginWithAligenieUserInfoResponseBodyResult) *AuthLoginWithAligenieUserInfoResponseBody {
	s.Result = v
	return s
}

func (s *AuthLoginWithAligenieUserInfoResponseBody) SetSuccess(v bool) *AuthLoginWithAligenieUserInfoResponseBody {
	s.Success = &v
	return s
}

func (s *AuthLoginWithAligenieUserInfoResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AuthLoginWithAligenieUserInfoResponseBodyResult struct {
	// Expiration time of the login state access token (long integer)
	//
	// example:
	//
	// 1659506854230
	ExpiredTimeLong *int64 `json:"ExpiredTimeLong,omitempty" xml:"ExpiredTimeLong,omitempty"`
	// Login state access token
	//
	// example:
	//
	// d15aa92de679d0d225aa845268be19ee
	LoginStateAccessToken *string `json:"LoginStateAccessToken,omitempty" xml:"LoginStateAccessToken,omitempty"`
}

func (s AuthLoginWithAligenieUserInfoResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s AuthLoginWithAligenieUserInfoResponseBodyResult) GoString() string {
	return s.String()
}

func (s *AuthLoginWithAligenieUserInfoResponseBodyResult) GetExpiredTimeLong() *int64 {
	return s.ExpiredTimeLong
}

func (s *AuthLoginWithAligenieUserInfoResponseBodyResult) GetLoginStateAccessToken() *string {
	return s.LoginStateAccessToken
}

func (s *AuthLoginWithAligenieUserInfoResponseBodyResult) SetExpiredTimeLong(v int64) *AuthLoginWithAligenieUserInfoResponseBodyResult {
	s.ExpiredTimeLong = &v
	return s
}

func (s *AuthLoginWithAligenieUserInfoResponseBodyResult) SetLoginStateAccessToken(v string) *AuthLoginWithAligenieUserInfoResponseBodyResult {
	s.LoginStateAccessToken = &v
	return s
}

func (s *AuthLoginWithAligenieUserInfoResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
