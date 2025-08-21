// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAuthLoginWithAligenieUserInfoGeneratedByPhoneNumberResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *AuthLoginWithAligenieUserInfoGeneratedByPhoneNumberResponseBody
	GetCode() *int32
	SetMessage(v string) *AuthLoginWithAligenieUserInfoGeneratedByPhoneNumberResponseBody
	GetMessage() *string
	SetRequestId(v string) *AuthLoginWithAligenieUserInfoGeneratedByPhoneNumberResponseBody
	GetRequestId() *string
	SetResult(v *AuthLoginWithAligenieUserInfoGeneratedByPhoneNumberResponseBodyResult) *AuthLoginWithAligenieUserInfoGeneratedByPhoneNumberResponseBody
	GetResult() *AuthLoginWithAligenieUserInfoGeneratedByPhoneNumberResponseBodyResult
	SetSuccess(v bool) *AuthLoginWithAligenieUserInfoGeneratedByPhoneNumberResponseBody
	GetSuccess() *bool
}

type AuthLoginWithAligenieUserInfoGeneratedByPhoneNumberResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 73C67BD9-175A-1324-8202-9FAABBB3E6FA
	RequestId *string                                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *AuthLoginWithAligenieUserInfoGeneratedByPhoneNumberResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AuthLoginWithAligenieUserInfoGeneratedByPhoneNumberResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AuthLoginWithAligenieUserInfoGeneratedByPhoneNumberResponseBody) GoString() string {
	return s.String()
}

func (s *AuthLoginWithAligenieUserInfoGeneratedByPhoneNumberResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *AuthLoginWithAligenieUserInfoGeneratedByPhoneNumberResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AuthLoginWithAligenieUserInfoGeneratedByPhoneNumberResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AuthLoginWithAligenieUserInfoGeneratedByPhoneNumberResponseBody) GetResult() *AuthLoginWithAligenieUserInfoGeneratedByPhoneNumberResponseBodyResult {
	return s.Result
}

func (s *AuthLoginWithAligenieUserInfoGeneratedByPhoneNumberResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AuthLoginWithAligenieUserInfoGeneratedByPhoneNumberResponseBody) SetCode(v int32) *AuthLoginWithAligenieUserInfoGeneratedByPhoneNumberResponseBody {
	s.Code = &v
	return s
}

func (s *AuthLoginWithAligenieUserInfoGeneratedByPhoneNumberResponseBody) SetMessage(v string) *AuthLoginWithAligenieUserInfoGeneratedByPhoneNumberResponseBody {
	s.Message = &v
	return s
}

func (s *AuthLoginWithAligenieUserInfoGeneratedByPhoneNumberResponseBody) SetRequestId(v string) *AuthLoginWithAligenieUserInfoGeneratedByPhoneNumberResponseBody {
	s.RequestId = &v
	return s
}

func (s *AuthLoginWithAligenieUserInfoGeneratedByPhoneNumberResponseBody) SetResult(v *AuthLoginWithAligenieUserInfoGeneratedByPhoneNumberResponseBodyResult) *AuthLoginWithAligenieUserInfoGeneratedByPhoneNumberResponseBody {
	s.Result = v
	return s
}

func (s *AuthLoginWithAligenieUserInfoGeneratedByPhoneNumberResponseBody) SetSuccess(v bool) *AuthLoginWithAligenieUserInfoGeneratedByPhoneNumberResponseBody {
	s.Success = &v
	return s
}

func (s *AuthLoginWithAligenieUserInfoGeneratedByPhoneNumberResponseBody) Validate() error {
	return dara.Validate(s)
}

type AuthLoginWithAligenieUserInfoGeneratedByPhoneNumberResponseBodyResult struct {
	// example:
	//
	// 1659506854230
	ExpiredTimeLong *int64 `json:"ExpiredTimeLong,omitempty" xml:"ExpiredTimeLong,omitempty"`
	// example:
	//
	// d15aa92de679d0d225aa845268be19ee
	LoginStateAccessToken *string `json:"LoginStateAccessToken,omitempty" xml:"LoginStateAccessToken,omitempty"`
}

func (s AuthLoginWithAligenieUserInfoGeneratedByPhoneNumberResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s AuthLoginWithAligenieUserInfoGeneratedByPhoneNumberResponseBodyResult) GoString() string {
	return s.String()
}

func (s *AuthLoginWithAligenieUserInfoGeneratedByPhoneNumberResponseBodyResult) GetExpiredTimeLong() *int64 {
	return s.ExpiredTimeLong
}

func (s *AuthLoginWithAligenieUserInfoGeneratedByPhoneNumberResponseBodyResult) GetLoginStateAccessToken() *string {
	return s.LoginStateAccessToken
}

func (s *AuthLoginWithAligenieUserInfoGeneratedByPhoneNumberResponseBodyResult) SetExpiredTimeLong(v int64) *AuthLoginWithAligenieUserInfoGeneratedByPhoneNumberResponseBodyResult {
	s.ExpiredTimeLong = &v
	return s
}

func (s *AuthLoginWithAligenieUserInfoGeneratedByPhoneNumberResponseBodyResult) SetLoginStateAccessToken(v string) *AuthLoginWithAligenieUserInfoGeneratedByPhoneNumberResponseBodyResult {
	s.LoginStateAccessToken = &v
	return s
}

func (s *AuthLoginWithAligenieUserInfoGeneratedByPhoneNumberResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
