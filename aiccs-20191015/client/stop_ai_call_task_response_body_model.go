// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopAiCallTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *StopAiCallTaskResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *StopAiCallTaskResponseBody
	GetCode() *string
	SetData(v bool) *StopAiCallTaskResponseBody
	GetData() *bool
	SetMessage(v string) *StopAiCallTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *StopAiCallTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *StopAiCallTaskResponseBody
	GetSuccess() *bool
}

type StopAiCallTaskResponseBody struct {
	// The detailed reason for the access denial.
	//
	// example:
	//
	// None
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The request status code. A value of `OK` indicates that the request was successful.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the task was stopped successfully. Valid values:
	//
	// - **true**: The operation was successful.
	//
	// - **false**: The operation failed.
	//
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The description of the status code.
	//
	// example:
	//
	// 成功
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 7124F972-BAB8-5D1E-90FC-01CB10713B29
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the API call was successful. Valid values:
	//
	// - **true**: The call was successful.
	//
	// - **false**: The call failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s StopAiCallTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StopAiCallTaskResponseBody) GoString() string {
	return s.String()
}

func (s *StopAiCallTaskResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *StopAiCallTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *StopAiCallTaskResponseBody) GetData() *bool {
	return s.Data
}

func (s *StopAiCallTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *StopAiCallTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StopAiCallTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *StopAiCallTaskResponseBody) SetAccessDeniedDetail(v string) *StopAiCallTaskResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *StopAiCallTaskResponseBody) SetCode(v string) *StopAiCallTaskResponseBody {
	s.Code = &v
	return s
}

func (s *StopAiCallTaskResponseBody) SetData(v bool) *StopAiCallTaskResponseBody {
	s.Data = &v
	return s
}

func (s *StopAiCallTaskResponseBody) SetMessage(v string) *StopAiCallTaskResponseBody {
	s.Message = &v
	return s
}

func (s *StopAiCallTaskResponseBody) SetRequestId(v string) *StopAiCallTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopAiCallTaskResponseBody) SetSuccess(v bool) *StopAiCallTaskResponseBody {
	s.Success = &v
	return s
}

func (s *StopAiCallTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
