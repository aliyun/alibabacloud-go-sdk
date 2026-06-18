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
	// The reason why access was denied.
	//
	// example:
	//
	// None
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The status code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the task was updated successfully. Valid values:
	//
	// - **true**: The update is successful.
	//
	// - **false**: The update failed.
	//
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The error message that is returned when an error occurs.
	//
	// example:
	//
	// 参数不合法
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// FCD94A7F-316D-54D1-9BFC-814006CB1C34
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the API call was successful. Valid values:
	//
	// - **true**: Successful.
	//
	// - **false**: Failed.
	//
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
