// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAligenieUserInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *GetAligenieUserInfoResponseBody
	GetCode() *int32
	SetMessage(v string) *GetAligenieUserInfoResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetAligenieUserInfoResponseBody
	GetRequestId() *string
	SetResult(v *GetAligenieUserInfoResponseBodyResult) *GetAligenieUserInfoResponseBody
	GetResult() *GetAligenieUserInfoResponseBodyResult
	SetSuccess(v bool) *GetAligenieUserInfoResponseBody
	GetSuccess() *bool
}

type GetAligenieUserInfoResponseBody struct {
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
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *GetAligenieUserInfoResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	Success   *bool                                  `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetAligenieUserInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAligenieUserInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetAligenieUserInfoResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetAligenieUserInfoResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetAligenieUserInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAligenieUserInfoResponseBody) GetResult() *GetAligenieUserInfoResponseBodyResult {
	return s.Result
}

func (s *GetAligenieUserInfoResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetAligenieUserInfoResponseBody) SetCode(v int32) *GetAligenieUserInfoResponseBody {
	s.Code = &v
	return s
}

func (s *GetAligenieUserInfoResponseBody) SetMessage(v string) *GetAligenieUserInfoResponseBody {
	s.Message = &v
	return s
}

func (s *GetAligenieUserInfoResponseBody) SetRequestId(v string) *GetAligenieUserInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAligenieUserInfoResponseBody) SetResult(v *GetAligenieUserInfoResponseBodyResult) *GetAligenieUserInfoResponseBody {
	s.Result = v
	return s
}

func (s *GetAligenieUserInfoResponseBody) SetSuccess(v bool) *GetAligenieUserInfoResponseBody {
	s.Success = &v
	return s
}

func (s *GetAligenieUserInfoResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetAligenieUserInfoResponseBodyResult struct {
	// example:
	//
	// XXX
	AligenieNickname *string `json:"AligenieNickname,omitempty" xml:"AligenieNickname,omitempty"`
	// example:
	//
	// http://img.alicdn.com/xxx.jpg
	Avatar    *string `json:"Avatar,omitempty" xml:"Avatar,omitempty"`
	Deletable *bool   `json:"Deletable,omitempty" xml:"Deletable,omitempty"`
}

func (s GetAligenieUserInfoResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s GetAligenieUserInfoResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetAligenieUserInfoResponseBodyResult) GetAligenieNickname() *string {
	return s.AligenieNickname
}

func (s *GetAligenieUserInfoResponseBodyResult) GetAvatar() *string {
	return s.Avatar
}

func (s *GetAligenieUserInfoResponseBodyResult) GetDeletable() *bool {
	return s.Deletable
}

func (s *GetAligenieUserInfoResponseBodyResult) SetAligenieNickname(v string) *GetAligenieUserInfoResponseBodyResult {
	s.AligenieNickname = &v
	return s
}

func (s *GetAligenieUserInfoResponseBodyResult) SetAvatar(v string) *GetAligenieUserInfoResponseBodyResult {
	s.Avatar = &v
	return s
}

func (s *GetAligenieUserInfoResponseBodyResult) SetDeletable(v bool) *GetAligenieUserInfoResponseBodyResult {
	s.Deletable = &v
	return s
}

func (s *GetAligenieUserInfoResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
