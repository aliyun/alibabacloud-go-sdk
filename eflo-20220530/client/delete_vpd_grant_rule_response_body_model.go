// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVpdGrantRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *DeleteVpdGrantRuleResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v int32) *DeleteVpdGrantRuleResponseBody
	GetCode() *int32
	SetContent(v interface{}) *DeleteVpdGrantRuleResponseBody
	GetContent() interface{}
	SetMessage(v string) *DeleteVpdGrantRuleResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteVpdGrantRuleResponseBody
	GetRequestId() *string
}

type DeleteVpdGrantRuleResponseBody struct {
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
	// 0901F411-28FA-5B9C-BAEE-7776463FF0DC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteVpdGrantRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteVpdGrantRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteVpdGrantRuleResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *DeleteVpdGrantRuleResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DeleteVpdGrantRuleResponseBody) GetContent() interface{} {
	return s.Content
}

func (s *DeleteVpdGrantRuleResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteVpdGrantRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteVpdGrantRuleResponseBody) SetAccessDeniedDetail(v string) *DeleteVpdGrantRuleResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *DeleteVpdGrantRuleResponseBody) SetCode(v int32) *DeleteVpdGrantRuleResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteVpdGrantRuleResponseBody) SetContent(v interface{}) *DeleteVpdGrantRuleResponseBody {
	s.Content = v
	return s
}

func (s *DeleteVpdGrantRuleResponseBody) SetMessage(v string) *DeleteVpdGrantRuleResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteVpdGrantRuleResponseBody) SetRequestId(v string) *DeleteVpdGrantRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteVpdGrantRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
