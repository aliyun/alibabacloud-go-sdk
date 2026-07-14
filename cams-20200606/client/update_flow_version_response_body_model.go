// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateFlowVersionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *UpdateFlowVersionResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *UpdateFlowVersionResponseBody
	GetCode() *string
	SetMessage(v string) *UpdateFlowVersionResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateFlowVersionResponseBody
	GetRequestId() *string
	SetResponse(v map[string]interface{}) *UpdateFlowVersionResponseBody
	GetResponse() map[string]interface{}
	SetSuccess(v bool) *UpdateFlowVersionResponseBody
	GetSuccess() *bool
}

type UpdateFlowVersionResponseBody struct {
	// The details about the access denial.
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
	// The content of the returned data.
	//
	// example:
	//
	// 436489
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

func (s UpdateFlowVersionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateFlowVersionResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateFlowVersionResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *UpdateFlowVersionResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateFlowVersionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateFlowVersionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateFlowVersionResponseBody) GetResponse() map[string]interface{} {
	return s.Response
}

func (s *UpdateFlowVersionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateFlowVersionResponseBody) SetAccessDeniedDetail(v string) *UpdateFlowVersionResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *UpdateFlowVersionResponseBody) SetCode(v string) *UpdateFlowVersionResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateFlowVersionResponseBody) SetMessage(v string) *UpdateFlowVersionResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateFlowVersionResponseBody) SetRequestId(v string) *UpdateFlowVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateFlowVersionResponseBody) SetResponse(v map[string]interface{}) *UpdateFlowVersionResponseBody {
	s.Response = v
	return s
}

func (s *UpdateFlowVersionResponseBody) SetSuccess(v bool) *UpdateFlowVersionResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateFlowVersionResponseBody) Validate() error {
	return dara.Validate(s)
}
