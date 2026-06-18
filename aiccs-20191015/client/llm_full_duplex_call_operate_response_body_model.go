// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLlmFullDuplexCallOperateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *LlmFullDuplexCallOperateResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *LlmFullDuplexCallOperateResponseBody
	GetCode() *string
	SetData(v bool) *LlmFullDuplexCallOperateResponseBody
	GetData() *bool
	SetMessage(v string) *LlmFullDuplexCallOperateResponseBody
	GetMessage() *string
	SetRequestId(v string) *LlmFullDuplexCallOperateResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *LlmFullDuplexCallOperateResponseBody
	GetSuccess() *bool
}

type LlmFullDuplexCallOperateResponseBody struct {
	// The details about the access denial. This parameter is returned only when RAM authentication fails.
	//
	// example:
	//
	// None
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The status code. A value of Success indicates that the request was successful.
	//
	// example:
	//
	// 示例值示例值
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The result of the command execution. Valid values:
	//
	// - **true**: The command was executed.
	//
	// - **false**: The command failed to be executed.
	//
	// example:
	//
	// false
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The description of the status code.
	//
	// example:
	//
	// 示例值示例值示例值
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 示例值示例值示例值
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful. Valid values:
	//
	// - **true**: The call was successful.
	//
	// - **false**: The call failed.
	//
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s LlmFullDuplexCallOperateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s LlmFullDuplexCallOperateResponseBody) GoString() string {
	return s.String()
}

func (s *LlmFullDuplexCallOperateResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *LlmFullDuplexCallOperateResponseBody) GetCode() *string {
	return s.Code
}

func (s *LlmFullDuplexCallOperateResponseBody) GetData() *bool {
	return s.Data
}

func (s *LlmFullDuplexCallOperateResponseBody) GetMessage() *string {
	return s.Message
}

func (s *LlmFullDuplexCallOperateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *LlmFullDuplexCallOperateResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *LlmFullDuplexCallOperateResponseBody) SetAccessDeniedDetail(v string) *LlmFullDuplexCallOperateResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *LlmFullDuplexCallOperateResponseBody) SetCode(v string) *LlmFullDuplexCallOperateResponseBody {
	s.Code = &v
	return s
}

func (s *LlmFullDuplexCallOperateResponseBody) SetData(v bool) *LlmFullDuplexCallOperateResponseBody {
	s.Data = &v
	return s
}

func (s *LlmFullDuplexCallOperateResponseBody) SetMessage(v string) *LlmFullDuplexCallOperateResponseBody {
	s.Message = &v
	return s
}

func (s *LlmFullDuplexCallOperateResponseBody) SetRequestId(v string) *LlmFullDuplexCallOperateResponseBody {
	s.RequestId = &v
	return s
}

func (s *LlmFullDuplexCallOperateResponseBody) SetSuccess(v bool) *LlmFullDuplexCallOperateResponseBody {
	s.Success = &v
	return s
}

func (s *LlmFullDuplexCallOperateResponseBody) Validate() error {
	return dara.Validate(s)
}
