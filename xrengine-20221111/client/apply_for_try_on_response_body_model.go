// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApplyForTryOnResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ApplyForTryOnResponseBody
	GetCode() *string
	SetErrorName(v string) *ApplyForTryOnResponseBody
	GetErrorName() *string
	SetHttpCode(v int32) *ApplyForTryOnResponseBody
	GetHttpCode() *int32
	SetMessage(v string) *ApplyForTryOnResponseBody
	GetMessage() *string
	SetRequestId(v string) *ApplyForTryOnResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ApplyForTryOnResponseBody
	GetSuccess() *bool
}

type ApplyForTryOnResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	ErrorName *string `json:"ErrorName,omitempty" xml:"ErrorName,omitempty"`
	HttpCode  *int32  `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ApplyForTryOnResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ApplyForTryOnResponseBody) GoString() string {
	return s.String()
}

func (s *ApplyForTryOnResponseBody) GetCode() *string {
	return s.Code
}

func (s *ApplyForTryOnResponseBody) GetErrorName() *string {
	return s.ErrorName
}

func (s *ApplyForTryOnResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *ApplyForTryOnResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ApplyForTryOnResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ApplyForTryOnResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ApplyForTryOnResponseBody) SetCode(v string) *ApplyForTryOnResponseBody {
	s.Code = &v
	return s
}

func (s *ApplyForTryOnResponseBody) SetErrorName(v string) *ApplyForTryOnResponseBody {
	s.ErrorName = &v
	return s
}

func (s *ApplyForTryOnResponseBody) SetHttpCode(v int32) *ApplyForTryOnResponseBody {
	s.HttpCode = &v
	return s
}

func (s *ApplyForTryOnResponseBody) SetMessage(v string) *ApplyForTryOnResponseBody {
	s.Message = &v
	return s
}

func (s *ApplyForTryOnResponseBody) SetRequestId(v string) *ApplyForTryOnResponseBody {
	s.RequestId = &v
	return s
}

func (s *ApplyForTryOnResponseBody) SetSuccess(v bool) *ApplyForTryOnResponseBody {
	s.Success = &v
	return s
}

func (s *ApplyForTryOnResponseBody) Validate() error {
	return dara.Validate(s)
}
