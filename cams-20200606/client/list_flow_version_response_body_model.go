// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFlowVersionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ListFlowVersionResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *ListFlowVersionResponseBody
	GetCode() *string
	SetMessage(v string) *ListFlowVersionResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListFlowVersionResponseBody
	GetRequestId() *string
	SetResponse(v map[string]interface{}) *ListFlowVersionResponseBody
	GetResponse() map[string]interface{}
	SetSuccess(v bool) *ListFlowVersionResponseBody
	GetSuccess() *bool
}

type ListFlowVersionResponseBody struct {
	// Details of access denied.
	//
	// example:
	//
	// 无
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// Status code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Error message.
	//
	// example:
	//
	// 无
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID.
	//
	// example:
	//
	// 90E63D28-E31D-1EB2-8939-A9486641****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Content of the returned data.
	//
	// example:
	//
	// 无
	Response map[string]interface{} `json:"Response,omitempty" xml:"Response,omitempty"`
	// Whether the operation was successful. Values: true: success; false: failure.
	//
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListFlowVersionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListFlowVersionResponseBody) GoString() string {
	return s.String()
}

func (s *ListFlowVersionResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ListFlowVersionResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListFlowVersionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListFlowVersionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListFlowVersionResponseBody) GetResponse() map[string]interface{} {
	return s.Response
}

func (s *ListFlowVersionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListFlowVersionResponseBody) SetAccessDeniedDetail(v string) *ListFlowVersionResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ListFlowVersionResponseBody) SetCode(v string) *ListFlowVersionResponseBody {
	s.Code = &v
	return s
}

func (s *ListFlowVersionResponseBody) SetMessage(v string) *ListFlowVersionResponseBody {
	s.Message = &v
	return s
}

func (s *ListFlowVersionResponseBody) SetRequestId(v string) *ListFlowVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListFlowVersionResponseBody) SetResponse(v map[string]interface{}) *ListFlowVersionResponseBody {
	s.Response = v
	return s
}

func (s *ListFlowVersionResponseBody) SetSuccess(v bool) *ListFlowVersionResponseBody {
	s.Success = &v
	return s
}

func (s *ListFlowVersionResponseBody) Validate() error {
	return dara.Validate(s)
}
