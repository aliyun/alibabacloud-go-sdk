// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLoginAppResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *LoginAppResponseBody
	GetCode() *string
	SetData(v *LoginAppResponseBodyData) *LoginAppResponseBody
	GetData() *LoginAppResponseBodyData
	SetErrorName(v string) *LoginAppResponseBody
	GetErrorName() *string
	SetHttpCode(v int32) *LoginAppResponseBody
	GetHttpCode() *int32
	SetMessage(v string) *LoginAppResponseBody
	GetMessage() *string
	SetRequestId(v string) *LoginAppResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *LoginAppResponseBody
	GetSuccess() *bool
}

type LoginAppResponseBody struct {
	Code      *string                   `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *LoginAppResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorName *string                   `json:"ErrorName,omitempty" xml:"ErrorName,omitempty"`
	HttpCode  *int32                    `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	Message   *string                   `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                     `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s LoginAppResponseBody) String() string {
	return dara.Prettify(s)
}

func (s LoginAppResponseBody) GoString() string {
	return s.String()
}

func (s *LoginAppResponseBody) GetCode() *string {
	return s.Code
}

func (s *LoginAppResponseBody) GetData() *LoginAppResponseBodyData {
	return s.Data
}

func (s *LoginAppResponseBody) GetErrorName() *string {
	return s.ErrorName
}

func (s *LoginAppResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *LoginAppResponseBody) GetMessage() *string {
	return s.Message
}

func (s *LoginAppResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *LoginAppResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *LoginAppResponseBody) SetCode(v string) *LoginAppResponseBody {
	s.Code = &v
	return s
}

func (s *LoginAppResponseBody) SetData(v *LoginAppResponseBodyData) *LoginAppResponseBody {
	s.Data = v
	return s
}

func (s *LoginAppResponseBody) SetErrorName(v string) *LoginAppResponseBody {
	s.ErrorName = &v
	return s
}

func (s *LoginAppResponseBody) SetHttpCode(v int32) *LoginAppResponseBody {
	s.HttpCode = &v
	return s
}

func (s *LoginAppResponseBody) SetMessage(v string) *LoginAppResponseBody {
	s.Message = &v
	return s
}

func (s *LoginAppResponseBody) SetRequestId(v string) *LoginAppResponseBody {
	s.RequestId = &v
	return s
}

func (s *LoginAppResponseBody) SetSuccess(v bool) *LoginAppResponseBody {
	s.Success = &v
	return s
}

func (s *LoginAppResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type LoginAppResponseBodyData struct {
	JwtToken *string `json:"JwtToken,omitempty" xml:"JwtToken,omitempty"`
	Nickname *string `json:"Nickname,omitempty" xml:"Nickname,omitempty"`
	Uid      *string `json:"Uid,omitempty" xml:"Uid,omitempty"`
}

func (s LoginAppResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s LoginAppResponseBodyData) GoString() string {
	return s.String()
}

func (s *LoginAppResponseBodyData) GetJwtToken() *string {
	return s.JwtToken
}

func (s *LoginAppResponseBodyData) GetNickname() *string {
	return s.Nickname
}

func (s *LoginAppResponseBodyData) GetUid() *string {
	return s.Uid
}

func (s *LoginAppResponseBodyData) SetJwtToken(v string) *LoginAppResponseBodyData {
	s.JwtToken = &v
	return s
}

func (s *LoginAppResponseBodyData) SetNickname(v string) *LoginAppResponseBodyData {
	s.Nickname = &v
	return s
}

func (s *LoginAppResponseBodyData) SetUid(v string) *LoginAppResponseBodyData {
	s.Uid = &v
	return s
}

func (s *LoginAppResponseBodyData) Validate() error {
	return dara.Validate(s)
}
