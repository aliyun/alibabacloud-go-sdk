// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSmsAuthorizationLetterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *CreateSmsAuthorizationLetterResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *CreateSmsAuthorizationLetterResponseBody
	GetCode() *string
	SetData(v string) *CreateSmsAuthorizationLetterResponseBody
	GetData() *string
	SetMessage(v string) *CreateSmsAuthorizationLetterResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateSmsAuthorizationLetterResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateSmsAuthorizationLetterResponseBody
	GetSuccess() *bool
}

type CreateSmsAuthorizationLetterResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 10000****
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// F655A8D5-B967-440B-8683-DAD6FF8DE990
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateSmsAuthorizationLetterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateSmsAuthorizationLetterResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSmsAuthorizationLetterResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *CreateSmsAuthorizationLetterResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateSmsAuthorizationLetterResponseBody) GetData() *string {
	return s.Data
}

func (s *CreateSmsAuthorizationLetterResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateSmsAuthorizationLetterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateSmsAuthorizationLetterResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateSmsAuthorizationLetterResponseBody) SetAccessDeniedDetail(v string) *CreateSmsAuthorizationLetterResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CreateSmsAuthorizationLetterResponseBody) SetCode(v string) *CreateSmsAuthorizationLetterResponseBody {
	s.Code = &v
	return s
}

func (s *CreateSmsAuthorizationLetterResponseBody) SetData(v string) *CreateSmsAuthorizationLetterResponseBody {
	s.Data = &v
	return s
}

func (s *CreateSmsAuthorizationLetterResponseBody) SetMessage(v string) *CreateSmsAuthorizationLetterResponseBody {
	s.Message = &v
	return s
}

func (s *CreateSmsAuthorizationLetterResponseBody) SetRequestId(v string) *CreateSmsAuthorizationLetterResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSmsAuthorizationLetterResponseBody) SetSuccess(v bool) *CreateSmsAuthorizationLetterResponseBody {
	s.Success = &v
	return s
}

func (s *CreateSmsAuthorizationLetterResponseBody) Validate() error {
	return dara.Validate(s)
}
