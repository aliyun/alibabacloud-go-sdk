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
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// 示例值示例值
	Code *string                  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data []map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 示例值示例值示例值
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 示例值示例值
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
