// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWabaMmlStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *UpdateWabaMmlStatusResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *UpdateWabaMmlStatusResponseBody
	GetCode() *string
	SetMessage(v string) *UpdateWabaMmlStatusResponseBody
	GetMessage() *string
	SetModel(v map[string]interface{}) *UpdateWabaMmlStatusResponseBody
	GetModel() map[string]interface{}
	SetRequestId(v string) *UpdateWabaMmlStatusResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateWabaMmlStatusResponseBody
	GetSuccess() *bool
}

type UpdateWabaMmlStatusResponseBody struct {
	// example:
	//
	// None
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// 示例值示例值
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 示例值示例值示例值
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 1
	Model map[string]interface{} `json:"Model,omitempty" xml:"Model,omitempty"`
	// example:
	//
	// 示例值示例值
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateWabaMmlStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateWabaMmlStatusResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateWabaMmlStatusResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *UpdateWabaMmlStatusResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateWabaMmlStatusResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateWabaMmlStatusResponseBody) GetModel() map[string]interface{} {
	return s.Model
}

func (s *UpdateWabaMmlStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateWabaMmlStatusResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateWabaMmlStatusResponseBody) SetAccessDeniedDetail(v string) *UpdateWabaMmlStatusResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *UpdateWabaMmlStatusResponseBody) SetCode(v string) *UpdateWabaMmlStatusResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateWabaMmlStatusResponseBody) SetMessage(v string) *UpdateWabaMmlStatusResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateWabaMmlStatusResponseBody) SetModel(v map[string]interface{}) *UpdateWabaMmlStatusResponseBody {
	s.Model = v
	return s
}

func (s *UpdateWabaMmlStatusResponseBody) SetRequestId(v string) *UpdateWabaMmlStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateWabaMmlStatusResponseBody) SetSuccess(v bool) *UpdateWabaMmlStatusResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateWabaMmlStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
