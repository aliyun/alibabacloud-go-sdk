// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMessengerPageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *CreateMessengerPageResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *CreateMessengerPageResponseBody
	GetCode() *string
	SetMessage(v string) *CreateMessengerPageResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateMessengerPageResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateMessengerPageResponseBody
	GetSuccess() *bool
}

type CreateMessengerPageResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 示例值示例值
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 39***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateMessengerPageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateMessengerPageResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMessengerPageResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *CreateMessengerPageResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateMessengerPageResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateMessengerPageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateMessengerPageResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateMessengerPageResponseBody) SetAccessDeniedDetail(v string) *CreateMessengerPageResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CreateMessengerPageResponseBody) SetCode(v string) *CreateMessengerPageResponseBody {
	s.Code = &v
	return s
}

func (s *CreateMessengerPageResponseBody) SetMessage(v string) *CreateMessengerPageResponseBody {
	s.Message = &v
	return s
}

func (s *CreateMessengerPageResponseBody) SetRequestId(v string) *CreateMessengerPageResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateMessengerPageResponseBody) SetSuccess(v bool) *CreateMessengerPageResponseBody {
	s.Success = &v
	return s
}

func (s *CreateMessengerPageResponseBody) Validate() error {
	return dara.Validate(s)
}
