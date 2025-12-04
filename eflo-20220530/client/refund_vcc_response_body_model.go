// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRefundVccResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *RefundVccResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v int32) *RefundVccResponseBody
	GetCode() *int32
	SetContent(v interface{}) *RefundVccResponseBody
	GetContent() interface{}
	SetMessage(v string) *RefundVccResponseBody
	GetMessage() *string
	SetRequestId(v string) *RefundVccResponseBody
	GetRequestId() *string
}

type RefundVccResponseBody struct {
	// The details about the access denial.
	//
	// >  This parameter is returned only if Resource Access Management (RAM) permission verification failed.
	//
	// example:
	//
	// None
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The response status code
	//
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// Response content
	//
	// example:
	//
	// {}
	Content interface{} `json:"Content,omitempty" xml:"Content,omitempty"`
	// Response message, which is \\"success\\" if the request succeeds
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID of the current request
	//
	// example:
	//
	// AC8C713A-A9F4-5984-A5E1-76496DF35153
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RefundVccResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RefundVccResponseBody) GoString() string {
	return s.String()
}

func (s *RefundVccResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *RefundVccResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *RefundVccResponseBody) GetContent() interface{} {
	return s.Content
}

func (s *RefundVccResponseBody) GetMessage() *string {
	return s.Message
}

func (s *RefundVccResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RefundVccResponseBody) SetAccessDeniedDetail(v string) *RefundVccResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *RefundVccResponseBody) SetCode(v int32) *RefundVccResponseBody {
	s.Code = &v
	return s
}

func (s *RefundVccResponseBody) SetContent(v interface{}) *RefundVccResponseBody {
	s.Content = v
	return s
}

func (s *RefundVccResponseBody) SetMessage(v string) *RefundVccResponseBody {
	s.Message = &v
	return s
}

func (s *RefundVccResponseBody) SetRequestId(v string) *RefundVccResponseBody {
	s.RequestId = &v
	return s
}

func (s *RefundVccResponseBody) Validate() error {
	return dara.Validate(s)
}
