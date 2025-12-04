// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRetryVccResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *RetryVccResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v int32) *RetryVccResponseBody
	GetCode() *int32
	SetContent(v interface{}) *RetryVccResponseBody
	GetContent() interface{}
	SetMessage(v string) *RetryVccResponseBody
	GetMessage() *string
	SetRequestId(v string) *RetryVccResponseBody
	GetRequestId() *string
}

type RetryVccResponseBody struct {
	// The details about the access denial.
	//
	// example:
	//
	// None
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The response status code.
	//
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	//
	// example:
	//
	// {}
	Content interface{} `json:"Content,omitempty" xml:"Content,omitempty"`
	// The error message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID of the current request
	//
	// example:
	//
	// 7F9082CC-3D94-560F-A575-8E8EF6CE2CB8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RetryVccResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RetryVccResponseBody) GoString() string {
	return s.String()
}

func (s *RetryVccResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *RetryVccResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *RetryVccResponseBody) GetContent() interface{} {
	return s.Content
}

func (s *RetryVccResponseBody) GetMessage() *string {
	return s.Message
}

func (s *RetryVccResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RetryVccResponseBody) SetAccessDeniedDetail(v string) *RetryVccResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *RetryVccResponseBody) SetCode(v int32) *RetryVccResponseBody {
	s.Code = &v
	return s
}

func (s *RetryVccResponseBody) SetContent(v interface{}) *RetryVccResponseBody {
	s.Content = v
	return s
}

func (s *RetryVccResponseBody) SetMessage(v string) *RetryVccResponseBody {
	s.Message = &v
	return s
}

func (s *RetryVccResponseBody) SetRequestId(v string) *RetryVccResponseBody {
	s.RequestId = &v
	return s
}

func (s *RetryVccResponseBody) Validate() error {
	return dara.Validate(s)
}
