// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLoginResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *LoginResponseBody
	GetCode() *string
	SetData(v *LoginResponseBodyData) *LoginResponseBody
	GetData() *LoginResponseBodyData
	SetErrorName(v string) *LoginResponseBody
	GetErrorName() *string
	SetHttpCode(v int32) *LoginResponseBody
	GetHttpCode() *int32
	SetMessage(v string) *LoginResponseBody
	GetMessage() *string
	SetRequestId(v string) *LoginResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *LoginResponseBody
	GetSuccess() *bool
}

type LoginResponseBody struct {
	Code      *string                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *LoginResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorName *string                `json:"ErrorName,omitempty" xml:"ErrorName,omitempty"`
	HttpCode  *int32                 `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	Message   *string                `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                  `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s LoginResponseBody) String() string {
	return dara.Prettify(s)
}

func (s LoginResponseBody) GoString() string {
	return s.String()
}

func (s *LoginResponseBody) GetCode() *string {
	return s.Code
}

func (s *LoginResponseBody) GetData() *LoginResponseBodyData {
	return s.Data
}

func (s *LoginResponseBody) GetErrorName() *string {
	return s.ErrorName
}

func (s *LoginResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *LoginResponseBody) GetMessage() *string {
	return s.Message
}

func (s *LoginResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *LoginResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *LoginResponseBody) SetCode(v string) *LoginResponseBody {
	s.Code = &v
	return s
}

func (s *LoginResponseBody) SetData(v *LoginResponseBodyData) *LoginResponseBody {
	s.Data = v
	return s
}

func (s *LoginResponseBody) SetErrorName(v string) *LoginResponseBody {
	s.ErrorName = &v
	return s
}

func (s *LoginResponseBody) SetHttpCode(v int32) *LoginResponseBody {
	s.HttpCode = &v
	return s
}

func (s *LoginResponseBody) SetMessage(v string) *LoginResponseBody {
	s.Message = &v
	return s
}

func (s *LoginResponseBody) SetRequestId(v string) *LoginResponseBody {
	s.RequestId = &v
	return s
}

func (s *LoginResponseBody) SetSuccess(v bool) *LoginResponseBody {
	s.Success = &v
	return s
}

func (s *LoginResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type LoginResponseBodyData struct {
	JwtToken *string `json:"JwtToken,omitempty" xml:"JwtToken,omitempty"`
	Nickname *string `json:"Nickname,omitempty" xml:"Nickname,omitempty"`
	Uid      *string `json:"Uid,omitempty" xml:"Uid,omitempty"`
}

func (s LoginResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s LoginResponseBodyData) GoString() string {
	return s.String()
}

func (s *LoginResponseBodyData) GetJwtToken() *string {
	return s.JwtToken
}

func (s *LoginResponseBodyData) GetNickname() *string {
	return s.Nickname
}

func (s *LoginResponseBodyData) GetUid() *string {
	return s.Uid
}

func (s *LoginResponseBodyData) SetJwtToken(v string) *LoginResponseBodyData {
	s.JwtToken = &v
	return s
}

func (s *LoginResponseBodyData) SetNickname(v string) *LoginResponseBodyData {
	s.Nickname = &v
	return s
}

func (s *LoginResponseBodyData) SetUid(v string) *LoginResponseBodyData {
	s.Uid = &v
	return s
}

func (s *LoginResponseBodyData) Validate() error {
	return dara.Validate(s)
}
