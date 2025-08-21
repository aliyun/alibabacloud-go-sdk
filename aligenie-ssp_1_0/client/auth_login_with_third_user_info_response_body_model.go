// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAuthLoginWithThirdUserInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *AuthLoginWithThirdUserInfoResponseBody
	GetCode() *int32
	SetDataObj(v *AuthLoginWithThirdUserInfoResponseBodyDataObj) *AuthLoginWithThirdUserInfoResponseBody
	GetDataObj() *AuthLoginWithThirdUserInfoResponseBodyDataObj
	SetMessage(v string) *AuthLoginWithThirdUserInfoResponseBody
	GetMessage() *string
	SetRequestId(v string) *AuthLoginWithThirdUserInfoResponseBody
	GetRequestId() *string
	SetResult(v *AuthLoginWithThirdUserInfoResponseBodyResult) *AuthLoginWithThirdUserInfoResponseBody
	GetResult() *AuthLoginWithThirdUserInfoResponseBodyResult
	SetSuccess(v bool) *AuthLoginWithThirdUserInfoResponseBody
	GetSuccess() *bool
}

type AuthLoginWithThirdUserInfoResponseBody struct {
	// example:
	//
	// 200
	Code    *int32                                         `json:"Code,omitempty" xml:"Code,omitempty"`
	DataObj *AuthLoginWithThirdUserInfoResponseBodyDataObj `json:"DataObj,omitempty" xml:"DataObj,omitempty" type:"Struct"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 73C67BD9-175A-1324-8202-9FAABBB3E6FA
	RequestId *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *AuthLoginWithThirdUserInfoResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AuthLoginWithThirdUserInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AuthLoginWithThirdUserInfoResponseBody) GoString() string {
	return s.String()
}

func (s *AuthLoginWithThirdUserInfoResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *AuthLoginWithThirdUserInfoResponseBody) GetDataObj() *AuthLoginWithThirdUserInfoResponseBodyDataObj {
	return s.DataObj
}

func (s *AuthLoginWithThirdUserInfoResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AuthLoginWithThirdUserInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AuthLoginWithThirdUserInfoResponseBody) GetResult() *AuthLoginWithThirdUserInfoResponseBodyResult {
	return s.Result
}

func (s *AuthLoginWithThirdUserInfoResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AuthLoginWithThirdUserInfoResponseBody) SetCode(v int32) *AuthLoginWithThirdUserInfoResponseBody {
	s.Code = &v
	return s
}

func (s *AuthLoginWithThirdUserInfoResponseBody) SetDataObj(v *AuthLoginWithThirdUserInfoResponseBodyDataObj) *AuthLoginWithThirdUserInfoResponseBody {
	s.DataObj = v
	return s
}

func (s *AuthLoginWithThirdUserInfoResponseBody) SetMessage(v string) *AuthLoginWithThirdUserInfoResponseBody {
	s.Message = &v
	return s
}

func (s *AuthLoginWithThirdUserInfoResponseBody) SetRequestId(v string) *AuthLoginWithThirdUserInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *AuthLoginWithThirdUserInfoResponseBody) SetResult(v *AuthLoginWithThirdUserInfoResponseBodyResult) *AuthLoginWithThirdUserInfoResponseBody {
	s.Result = v
	return s
}

func (s *AuthLoginWithThirdUserInfoResponseBody) SetSuccess(v bool) *AuthLoginWithThirdUserInfoResponseBody {
	s.Success = &v
	return s
}

func (s *AuthLoginWithThirdUserInfoResponseBody) Validate() error {
	return dara.Validate(s)
}

type AuthLoginWithThirdUserInfoResponseBodyDataObj struct {
	// example:
	//
	// dbe2eb4458302b9246c6da17fbc95f4b
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
}

func (s AuthLoginWithThirdUserInfoResponseBodyDataObj) String() string {
	return dara.Prettify(s)
}

func (s AuthLoginWithThirdUserInfoResponseBodyDataObj) GoString() string {
	return s.String()
}

func (s *AuthLoginWithThirdUserInfoResponseBodyDataObj) GetSessionId() *string {
	return s.SessionId
}

func (s *AuthLoginWithThirdUserInfoResponseBodyDataObj) SetSessionId(v string) *AuthLoginWithThirdUserInfoResponseBodyDataObj {
	s.SessionId = &v
	return s
}

func (s *AuthLoginWithThirdUserInfoResponseBodyDataObj) Validate() error {
	return dara.Validate(s)
}

type AuthLoginWithThirdUserInfoResponseBodyResult struct {
	// example:
	//
	// 1659428051452
	ExpiredTimeLong *int64 `json:"ExpiredTimeLong,omitempty" xml:"ExpiredTimeLong,omitempty"`
	// example:
	//
	// bd9ccdb121ee950ddead51e943e081fe
	LoginStateAccessToken *string `json:"LoginStateAccessToken,omitempty" xml:"LoginStateAccessToken,omitempty"`
}

func (s AuthLoginWithThirdUserInfoResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s AuthLoginWithThirdUserInfoResponseBodyResult) GoString() string {
	return s.String()
}

func (s *AuthLoginWithThirdUserInfoResponseBodyResult) GetExpiredTimeLong() *int64 {
	return s.ExpiredTimeLong
}

func (s *AuthLoginWithThirdUserInfoResponseBodyResult) GetLoginStateAccessToken() *string {
	return s.LoginStateAccessToken
}

func (s *AuthLoginWithThirdUserInfoResponseBodyResult) SetExpiredTimeLong(v int64) *AuthLoginWithThirdUserInfoResponseBodyResult {
	s.ExpiredTimeLong = &v
	return s
}

func (s *AuthLoginWithThirdUserInfoResponseBodyResult) SetLoginStateAccessToken(v string) *AuthLoginWithThirdUserInfoResponseBodyResult {
	s.LoginStateAccessToken = &v
	return s
}

func (s *AuthLoginWithThirdUserInfoResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
