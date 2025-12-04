// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVccGrantRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *DeleteVccGrantRuleResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v int32) *DeleteVccGrantRuleResponseBody
	GetCode() *int32
	SetContent(v interface{}) *DeleteVccGrantRuleResponseBody
	GetContent() interface{}
	SetMessage(v string) *DeleteVccGrantRuleResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteVccGrantRuleResponseBody
	GetRequestId() *string
}

type DeleteVccGrantRuleResponseBody struct {
	// The details about the access denial.
	//
	// >  This parameter is returned only if Resource Access Management (RAM) permission verification failed.
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
	// BDBCC783-84CA-5733-8EEA-645C88B9009C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteVccGrantRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteVccGrantRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteVccGrantRuleResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *DeleteVccGrantRuleResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DeleteVccGrantRuleResponseBody) GetContent() interface{} {
	return s.Content
}

func (s *DeleteVccGrantRuleResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteVccGrantRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteVccGrantRuleResponseBody) SetAccessDeniedDetail(v string) *DeleteVccGrantRuleResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *DeleteVccGrantRuleResponseBody) SetCode(v int32) *DeleteVccGrantRuleResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteVccGrantRuleResponseBody) SetContent(v interface{}) *DeleteVccGrantRuleResponseBody {
	s.Content = v
	return s
}

func (s *DeleteVccGrantRuleResponseBody) SetMessage(v string) *DeleteVccGrantRuleResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteVccGrantRuleResponseBody) SetRequestId(v string) *DeleteVccGrantRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteVccGrantRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
