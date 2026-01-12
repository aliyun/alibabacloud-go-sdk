// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddWhitelistResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *AddWhitelistResponseBody
	GetCode() *string
	SetErrorName(v string) *AddWhitelistResponseBody
	GetErrorName() *string
	SetHttpCode(v int32) *AddWhitelistResponseBody
	GetHttpCode() *int32
	SetMessage(v string) *AddWhitelistResponseBody
	GetMessage() *string
	SetRequestId(v string) *AddWhitelistResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AddWhitelistResponseBody
	GetSuccess() *bool
}

type AddWhitelistResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	ErrorName *string `json:"ErrorName,omitempty" xml:"ErrorName,omitempty"`
	HttpCode  *int32  `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddWhitelistResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddWhitelistResponseBody) GoString() string {
	return s.String()
}

func (s *AddWhitelistResponseBody) GetCode() *string {
	return s.Code
}

func (s *AddWhitelistResponseBody) GetErrorName() *string {
	return s.ErrorName
}

func (s *AddWhitelistResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *AddWhitelistResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AddWhitelistResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddWhitelistResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AddWhitelistResponseBody) SetCode(v string) *AddWhitelistResponseBody {
	s.Code = &v
	return s
}

func (s *AddWhitelistResponseBody) SetErrorName(v string) *AddWhitelistResponseBody {
	s.ErrorName = &v
	return s
}

func (s *AddWhitelistResponseBody) SetHttpCode(v int32) *AddWhitelistResponseBody {
	s.HttpCode = &v
	return s
}

func (s *AddWhitelistResponseBody) SetMessage(v string) *AddWhitelistResponseBody {
	s.Message = &v
	return s
}

func (s *AddWhitelistResponseBody) SetRequestId(v string) *AddWhitelistResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddWhitelistResponseBody) SetSuccess(v bool) *AddWhitelistResponseBody {
	s.Success = &v
	return s
}

func (s *AddWhitelistResponseBody) Validate() error {
	return dara.Validate(s)
}
