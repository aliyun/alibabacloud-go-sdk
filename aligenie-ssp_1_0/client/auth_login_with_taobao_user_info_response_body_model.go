// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAuthLoginWithTaobaoUserInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *AuthLoginWithTaobaoUserInfoResponseBody
	GetCode() *int32
	SetMessage(v string) *AuthLoginWithTaobaoUserInfoResponseBody
	GetMessage() *string
	SetRequestId(v string) *AuthLoginWithTaobaoUserInfoResponseBody
	GetRequestId() *string
	SetResult(v *AuthLoginWithTaobaoUserInfoResponseBodyResult) *AuthLoginWithTaobaoUserInfoResponseBody
	GetResult() *AuthLoginWithTaobaoUserInfoResponseBodyResult
	SetSuccess(v bool) *AuthLoginWithTaobaoUserInfoResponseBody
	GetSuccess() *bool
}

type AuthLoginWithTaobaoUserInfoResponseBody struct {
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
	RequestId *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *AuthLoginWithTaobaoUserInfoResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	Success   *bool                                          `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AuthLoginWithTaobaoUserInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AuthLoginWithTaobaoUserInfoResponseBody) GoString() string {
	return s.String()
}

func (s *AuthLoginWithTaobaoUserInfoResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *AuthLoginWithTaobaoUserInfoResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AuthLoginWithTaobaoUserInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AuthLoginWithTaobaoUserInfoResponseBody) GetResult() *AuthLoginWithTaobaoUserInfoResponseBodyResult {
	return s.Result
}

func (s *AuthLoginWithTaobaoUserInfoResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AuthLoginWithTaobaoUserInfoResponseBody) SetCode(v int32) *AuthLoginWithTaobaoUserInfoResponseBody {
	s.Code = &v
	return s
}

func (s *AuthLoginWithTaobaoUserInfoResponseBody) SetMessage(v string) *AuthLoginWithTaobaoUserInfoResponseBody {
	s.Message = &v
	return s
}

func (s *AuthLoginWithTaobaoUserInfoResponseBody) SetRequestId(v string) *AuthLoginWithTaobaoUserInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *AuthLoginWithTaobaoUserInfoResponseBody) SetResult(v *AuthLoginWithTaobaoUserInfoResponseBodyResult) *AuthLoginWithTaobaoUserInfoResponseBody {
	s.Result = v
	return s
}

func (s *AuthLoginWithTaobaoUserInfoResponseBody) SetSuccess(v bool) *AuthLoginWithTaobaoUserInfoResponseBody {
	s.Success = &v
	return s
}

func (s *AuthLoginWithTaobaoUserInfoResponseBody) Validate() error {
	return dara.Validate(s)
}

type AuthLoginWithTaobaoUserInfoResponseBodyResult struct {
	// example:
	//
	// 1659506854230
	ExpiredTimeLong *int64 `json:"ExpiredTimeLong,omitempty" xml:"ExpiredTimeLong,omitempty"`
	// example:
	//
	// d15aa92de679d0d225aa845268be19ee
	LoginStateAccessToken *string `json:"LoginStateAccessToken,omitempty" xml:"LoginStateAccessToken,omitempty"`
}

func (s AuthLoginWithTaobaoUserInfoResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s AuthLoginWithTaobaoUserInfoResponseBodyResult) GoString() string {
	return s.String()
}

func (s *AuthLoginWithTaobaoUserInfoResponseBodyResult) GetExpiredTimeLong() *int64 {
	return s.ExpiredTimeLong
}

func (s *AuthLoginWithTaobaoUserInfoResponseBodyResult) GetLoginStateAccessToken() *string {
	return s.LoginStateAccessToken
}

func (s *AuthLoginWithTaobaoUserInfoResponseBodyResult) SetExpiredTimeLong(v int64) *AuthLoginWithTaobaoUserInfoResponseBodyResult {
	s.ExpiredTimeLong = &v
	return s
}

func (s *AuthLoginWithTaobaoUserInfoResponseBodyResult) SetLoginStateAccessToken(v string) *AuthLoginWithTaobaoUserInfoResponseBodyResult {
	s.LoginStateAccessToken = &v
	return s
}

func (s *AuthLoginWithTaobaoUserInfoResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
