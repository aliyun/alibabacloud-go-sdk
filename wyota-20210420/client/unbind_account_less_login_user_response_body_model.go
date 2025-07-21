// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnbindAccountLessLoginUserResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UnbindAccountLessLoginUserResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *UnbindAccountLessLoginUserResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UnbindAccountLessLoginUserResponseBody
	GetMessage() *string
	SetRequestId(v string) *UnbindAccountLessLoginUserResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UnbindAccountLessLoginUserResponseBody
	GetSuccess() *bool
}

type UnbindAccountLessLoginUserResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UnbindAccountLessLoginUserResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UnbindAccountLessLoginUserResponseBody) GoString() string {
	return s.String()
}

func (s *UnbindAccountLessLoginUserResponseBody) GetCode() *string {
	return s.Code
}

func (s *UnbindAccountLessLoginUserResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UnbindAccountLessLoginUserResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UnbindAccountLessLoginUserResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UnbindAccountLessLoginUserResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UnbindAccountLessLoginUserResponseBody) SetCode(v string) *UnbindAccountLessLoginUserResponseBody {
	s.Code = &v
	return s
}

func (s *UnbindAccountLessLoginUserResponseBody) SetHttpStatusCode(v int32) *UnbindAccountLessLoginUserResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UnbindAccountLessLoginUserResponseBody) SetMessage(v string) *UnbindAccountLessLoginUserResponseBody {
	s.Message = &v
	return s
}

func (s *UnbindAccountLessLoginUserResponseBody) SetRequestId(v string) *UnbindAccountLessLoginUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *UnbindAccountLessLoginUserResponseBody) SetSuccess(v bool) *UnbindAccountLessLoginUserResponseBody {
	s.Success = &v
	return s
}

func (s *UnbindAccountLessLoginUserResponseBody) Validate() error {
	return dara.Validate(s)
}
