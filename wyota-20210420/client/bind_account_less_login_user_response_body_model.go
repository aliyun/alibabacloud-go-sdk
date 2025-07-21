// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindAccountLessLoginUserResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *BindAccountLessLoginUserResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *BindAccountLessLoginUserResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *BindAccountLessLoginUserResponseBody
	GetMessage() *string
	SetRequestId(v string) *BindAccountLessLoginUserResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *BindAccountLessLoginUserResponseBody
	GetSuccess() *bool
}

type BindAccountLessLoginUserResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s BindAccountLessLoginUserResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BindAccountLessLoginUserResponseBody) GoString() string {
	return s.String()
}

func (s *BindAccountLessLoginUserResponseBody) GetCode() *string {
	return s.Code
}

func (s *BindAccountLessLoginUserResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *BindAccountLessLoginUserResponseBody) GetMessage() *string {
	return s.Message
}

func (s *BindAccountLessLoginUserResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BindAccountLessLoginUserResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *BindAccountLessLoginUserResponseBody) SetCode(v string) *BindAccountLessLoginUserResponseBody {
	s.Code = &v
	return s
}

func (s *BindAccountLessLoginUserResponseBody) SetHttpStatusCode(v int32) *BindAccountLessLoginUserResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *BindAccountLessLoginUserResponseBody) SetMessage(v string) *BindAccountLessLoginUserResponseBody {
	s.Message = &v
	return s
}

func (s *BindAccountLessLoginUserResponseBody) SetRequestId(v string) *BindAccountLessLoginUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *BindAccountLessLoginUserResponseBody) SetSuccess(v bool) *BindAccountLessLoginUserResponseBody {
	s.Success = &v
	return s
}

func (s *BindAccountLessLoginUserResponseBody) Validate() error {
	return dara.Validate(s)
}
