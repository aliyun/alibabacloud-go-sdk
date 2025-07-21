// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReadFlowVersionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ReadFlowVersionResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *ReadFlowVersionResponseBody
	GetCode() *string
	SetMessage(v string) *ReadFlowVersionResponseBody
	GetMessage() *string
	SetRequestId(v string) *ReadFlowVersionResponseBody
	GetRequestId() *string
	SetResponse(v map[string]interface{}) *ReadFlowVersionResponseBody
	GetResponse() map[string]interface{}
	SetSuccess(v bool) *ReadFlowVersionResponseBody
	GetSuccess() *bool
}

type ReadFlowVersionResponseBody struct {
	// Access denied details; this field is only returned when RAM verification fails.
	//
	// example:
	//
	// 无
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// Status code.
	//
	// example:
	//
	// 示例值示例值示例值
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Error message.
	//
	// example:
	//
	// 示例值示例值
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID.
	//
	// example:
	//
	// 示例值
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Content of the returned data.
	//
	// example:
	//
	// 无
	Response map[string]interface{} `json:"Response,omitempty" xml:"Response,omitempty"`
	// Indicates whether the operation was successful. Values: true: success; false: failure.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ReadFlowVersionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ReadFlowVersionResponseBody) GoString() string {
	return s.String()
}

func (s *ReadFlowVersionResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ReadFlowVersionResponseBody) GetCode() *string {
	return s.Code
}

func (s *ReadFlowVersionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ReadFlowVersionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ReadFlowVersionResponseBody) GetResponse() map[string]interface{} {
	return s.Response
}

func (s *ReadFlowVersionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ReadFlowVersionResponseBody) SetAccessDeniedDetail(v string) *ReadFlowVersionResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ReadFlowVersionResponseBody) SetCode(v string) *ReadFlowVersionResponseBody {
	s.Code = &v
	return s
}

func (s *ReadFlowVersionResponseBody) SetMessage(v string) *ReadFlowVersionResponseBody {
	s.Message = &v
	return s
}

func (s *ReadFlowVersionResponseBody) SetRequestId(v string) *ReadFlowVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReadFlowVersionResponseBody) SetResponse(v map[string]interface{}) *ReadFlowVersionResponseBody {
	s.Response = v
	return s
}

func (s *ReadFlowVersionResponseBody) SetSuccess(v bool) *ReadFlowVersionResponseBody {
	s.Success = &v
	return s
}

func (s *ReadFlowVersionResponseBody) Validate() error {
	return dara.Validate(s)
}
