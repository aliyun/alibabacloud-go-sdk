// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBindingRelationsForFlowVersionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ListBindingRelationsForFlowVersionResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *ListBindingRelationsForFlowVersionResponseBody
	GetCode() *string
	SetData(v []map[string]interface{}) *ListBindingRelationsForFlowVersionResponseBody
	GetData() []map[string]interface{}
	SetMessage(v string) *ListBindingRelationsForFlowVersionResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListBindingRelationsForFlowVersionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListBindingRelationsForFlowVersionResponseBody
	GetSuccess() *bool
}

type ListBindingRelationsForFlowVersionResponseBody struct {
	// Access denied details.
	//
	// example:
	//
	// None
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// Error code. For more information, see [Error Codes](https://help.aliyun.com/document_detail/196974.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Returned data list.
	Data []map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
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
	// 90E63D28-E31D-1EB2-8939-A94866411B2O
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the operation was successful. Values:
	//
	// - true: Success.
	//
	// - false: Failure.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListBindingRelationsForFlowVersionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListBindingRelationsForFlowVersionResponseBody) GoString() string {
	return s.String()
}

func (s *ListBindingRelationsForFlowVersionResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ListBindingRelationsForFlowVersionResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListBindingRelationsForFlowVersionResponseBody) GetData() []map[string]interface{} {
	return s.Data
}

func (s *ListBindingRelationsForFlowVersionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListBindingRelationsForFlowVersionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListBindingRelationsForFlowVersionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListBindingRelationsForFlowVersionResponseBody) SetAccessDeniedDetail(v string) *ListBindingRelationsForFlowVersionResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ListBindingRelationsForFlowVersionResponseBody) SetCode(v string) *ListBindingRelationsForFlowVersionResponseBody {
	s.Code = &v
	return s
}

func (s *ListBindingRelationsForFlowVersionResponseBody) SetData(v []map[string]interface{}) *ListBindingRelationsForFlowVersionResponseBody {
	s.Data = v
	return s
}

func (s *ListBindingRelationsForFlowVersionResponseBody) SetMessage(v string) *ListBindingRelationsForFlowVersionResponseBody {
	s.Message = &v
	return s
}

func (s *ListBindingRelationsForFlowVersionResponseBody) SetRequestId(v string) *ListBindingRelationsForFlowVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListBindingRelationsForFlowVersionResponseBody) SetSuccess(v bool) *ListBindingRelationsForFlowVersionResponseBody {
	s.Success = &v
	return s
}

func (s *ListBindingRelationsForFlowVersionResponseBody) Validate() error {
	return dara.Validate(s)
}
