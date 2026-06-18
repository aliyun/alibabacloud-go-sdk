// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartAiCallTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *StartAiCallTaskResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *StartAiCallTaskResponseBody
	GetCode() *string
	SetData(v bool) *StartAiCallTaskResponseBody
	GetData() *bool
	SetMessage(v string) *StartAiCallTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *StartAiCallTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *StartAiCallTaskResponseBody
	GetSuccess() *bool
}

type StartAiCallTaskResponseBody struct {
	// The detailed reason why access is denied.
	//
	// example:
	//
	// None
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The status code. OK indicates that the request was successful.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the task was successfully started. Valid values:
	//
	// - **true**: Success.
	//
	// - **false**: Failure.
	//
	// example:
	//
	// false
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The status message.
	//
	// example:
	//
	// 成功
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 58A32FFF-86EE-5CF0-B365-97E8C574C7F4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the API call was successful. Valid values:
	//
	// - **true**: Success.
	//
	// - **false**: Failure.
	//
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s StartAiCallTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StartAiCallTaskResponseBody) GoString() string {
	return s.String()
}

func (s *StartAiCallTaskResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *StartAiCallTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *StartAiCallTaskResponseBody) GetData() *bool {
	return s.Data
}

func (s *StartAiCallTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *StartAiCallTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StartAiCallTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *StartAiCallTaskResponseBody) SetAccessDeniedDetail(v string) *StartAiCallTaskResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *StartAiCallTaskResponseBody) SetCode(v string) *StartAiCallTaskResponseBody {
	s.Code = &v
	return s
}

func (s *StartAiCallTaskResponseBody) SetData(v bool) *StartAiCallTaskResponseBody {
	s.Data = &v
	return s
}

func (s *StartAiCallTaskResponseBody) SetMessage(v string) *StartAiCallTaskResponseBody {
	s.Message = &v
	return s
}

func (s *StartAiCallTaskResponseBody) SetRequestId(v string) *StartAiCallTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartAiCallTaskResponseBody) SetSuccess(v bool) *StartAiCallTaskResponseBody {
	s.Success = &v
	return s
}

func (s *StartAiCallTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
