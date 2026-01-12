// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAuthUserResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *AuthUserResponseBody
	GetCode() *string
	SetData(v *AuthUserResponseBodyData) *AuthUserResponseBody
	GetData() *AuthUserResponseBodyData
	SetErrorName(v string) *AuthUserResponseBody
	GetErrorName() *string
	SetHttpCode(v int32) *AuthUserResponseBody
	GetHttpCode() *int32
	SetMessage(v string) *AuthUserResponseBody
	GetMessage() *string
	SetRequestId(v string) *AuthUserResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AuthUserResponseBody
	GetSuccess() *bool
}

type AuthUserResponseBody struct {
	Code      *string                   `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *AuthUserResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorName *string                   `json:"ErrorName,omitempty" xml:"ErrorName,omitempty"`
	HttpCode  *int32                    `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	Message   *string                   `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                     `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AuthUserResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AuthUserResponseBody) GoString() string {
	return s.String()
}

func (s *AuthUserResponseBody) GetCode() *string {
	return s.Code
}

func (s *AuthUserResponseBody) GetData() *AuthUserResponseBodyData {
	return s.Data
}

func (s *AuthUserResponseBody) GetErrorName() *string {
	return s.ErrorName
}

func (s *AuthUserResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *AuthUserResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AuthUserResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AuthUserResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AuthUserResponseBody) SetCode(v string) *AuthUserResponseBody {
	s.Code = &v
	return s
}

func (s *AuthUserResponseBody) SetData(v *AuthUserResponseBodyData) *AuthUserResponseBody {
	s.Data = v
	return s
}

func (s *AuthUserResponseBody) SetErrorName(v string) *AuthUserResponseBody {
	s.ErrorName = &v
	return s
}

func (s *AuthUserResponseBody) SetHttpCode(v int32) *AuthUserResponseBody {
	s.HttpCode = &v
	return s
}

func (s *AuthUserResponseBody) SetMessage(v string) *AuthUserResponseBody {
	s.Message = &v
	return s
}

func (s *AuthUserResponseBody) SetRequestId(v string) *AuthUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *AuthUserResponseBody) SetSuccess(v bool) *AuthUserResponseBody {
	s.Success = &v
	return s
}

func (s *AuthUserResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AuthUserResponseBodyData struct {
	JwtToken *string `json:"JwtToken,omitempty" xml:"JwtToken,omitempty"`
	Type     *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s AuthUserResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s AuthUserResponseBodyData) GoString() string {
	return s.String()
}

func (s *AuthUserResponseBodyData) GetJwtToken() *string {
	return s.JwtToken
}

func (s *AuthUserResponseBodyData) GetType() *string {
	return s.Type
}

func (s *AuthUserResponseBodyData) SetJwtToken(v string) *AuthUserResponseBodyData {
	s.JwtToken = &v
	return s
}

func (s *AuthUserResponseBodyData) SetType(v string) *AuthUserResponseBodyData {
	s.Type = &v
	return s
}

func (s *AuthUserResponseBodyData) Validate() error {
	return dara.Validate(s)
}
