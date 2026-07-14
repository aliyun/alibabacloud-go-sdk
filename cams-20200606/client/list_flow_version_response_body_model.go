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
	// Details about the access denial.
	//
	// example:
	//
	// None
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The error code. For more information, see [Error codes](https://help.aliyun.com/document_detail/196974.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message.
	//
	// example:
	//
	// None
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 90E63D28-E31D-1EB2-8939-A94866411B2O
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The response data.
	Response map[string]interface{} `json:"Response,omitempty" xml:"Response,omitempty"`
	// Indicates whether the operation was successful. Valid values:
	//
	// - true: The operation was successful.
	//
	// - false: The operation failed.
	//
	// example:
	//
	// true
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
