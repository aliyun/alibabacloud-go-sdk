// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveSourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SaveSourceResponseBody
	GetCode() *string
	SetErrorName(v string) *SaveSourceResponseBody
	GetErrorName() *string
	SetHttpCode(v int32) *SaveSourceResponseBody
	GetHttpCode() *int32
	SetMessage(v string) *SaveSourceResponseBody
	GetMessage() *string
	SetRequestId(v string) *SaveSourceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SaveSourceResponseBody
	GetSuccess() *bool
}

type SaveSourceResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	ErrorName *string `json:"ErrorName,omitempty" xml:"ErrorName,omitempty"`
	HttpCode  *int32  `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SaveSourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SaveSourceResponseBody) GoString() string {
	return s.String()
}

func (s *SaveSourceResponseBody) GetCode() *string {
	return s.Code
}

func (s *SaveSourceResponseBody) GetErrorName() *string {
	return s.ErrorName
}

func (s *SaveSourceResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *SaveSourceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SaveSourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SaveSourceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SaveSourceResponseBody) SetCode(v string) *SaveSourceResponseBody {
	s.Code = &v
	return s
}

func (s *SaveSourceResponseBody) SetErrorName(v string) *SaveSourceResponseBody {
	s.ErrorName = &v
	return s
}

func (s *SaveSourceResponseBody) SetHttpCode(v int32) *SaveSourceResponseBody {
	s.HttpCode = &v
	return s
}

func (s *SaveSourceResponseBody) SetMessage(v string) *SaveSourceResponseBody {
	s.Message = &v
	return s
}

func (s *SaveSourceResponseBody) SetRequestId(v string) *SaveSourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveSourceResponseBody) SetSuccess(v bool) *SaveSourceResponseBody {
	s.Success = &v
	return s
}

func (s *SaveSourceResponseBody) Validate() error {
	return dara.Validate(s)
}
