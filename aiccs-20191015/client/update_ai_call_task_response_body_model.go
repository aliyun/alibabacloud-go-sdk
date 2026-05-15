// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAiCallTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *UpdateAiCallTaskResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *UpdateAiCallTaskResponseBody
	GetCode() *string
	SetData(v bool) *UpdateAiCallTaskResponseBody
	GetData() *bool
	SetMessage(v string) *UpdateAiCallTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateAiCallTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateAiCallTaskResponseBody
	GetSuccess() *bool
}

type UpdateAiCallTaskResponseBody struct {
	// example:
	//
	// None
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 示例值示例值示例值
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// FCD94A7F-316D-54D1-9BFC-814006CB1C34
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateAiCallTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateAiCallTaskResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAiCallTaskResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *UpdateAiCallTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateAiCallTaskResponseBody) GetData() *bool {
	return s.Data
}

func (s *UpdateAiCallTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateAiCallTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateAiCallTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateAiCallTaskResponseBody) SetAccessDeniedDetail(v string) *UpdateAiCallTaskResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *UpdateAiCallTaskResponseBody) SetCode(v string) *UpdateAiCallTaskResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateAiCallTaskResponseBody) SetData(v bool) *UpdateAiCallTaskResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateAiCallTaskResponseBody) SetMessage(v string) *UpdateAiCallTaskResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateAiCallTaskResponseBody) SetRequestId(v string) *UpdateAiCallTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAiCallTaskResponseBody) SetSuccess(v bool) *UpdateAiCallTaskResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateAiCallTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
