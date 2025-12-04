// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteErResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *DeleteErResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v int32) *DeleteErResponseBody
	GetCode() *int32
	SetContent(v interface{}) *DeleteErResponseBody
	GetContent() interface{}
	SetMessage(v string) *DeleteErResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteErResponseBody
	GetRequestId() *string
}

type DeleteErResponseBody struct {
	// The details about the access denial.
	//
	// example:
	//
	// {}
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The response status code.
	//
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// Response body
	//
	// example:
	//
	// {}
	Content interface{} `json:"Content,omitempty" xml:"Content,omitempty"`
	// The error message. (If the instance is in the Exception state, the exception cause is prompted.)
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID of the current request
	//
	// example:
	//
	// 9C50C9CD-E799-54DA-BA7A-1FAF3DF80857
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteErResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteErResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteErResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *DeleteErResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DeleteErResponseBody) GetContent() interface{} {
	return s.Content
}

func (s *DeleteErResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteErResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteErResponseBody) SetAccessDeniedDetail(v string) *DeleteErResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *DeleteErResponseBody) SetCode(v int32) *DeleteErResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteErResponseBody) SetContent(v interface{}) *DeleteErResponseBody {
	s.Content = v
	return s
}

func (s *DeleteErResponseBody) SetMessage(v string) *DeleteErResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteErResponseBody) SetRequestId(v string) *DeleteErResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteErResponseBody) Validate() error {
	return dara.Validate(s)
}
