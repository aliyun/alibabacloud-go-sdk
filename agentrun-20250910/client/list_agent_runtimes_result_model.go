// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAgentRuntimesResult interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListAgentRuntimesResult
	GetCode() *string
	SetData(v *ListAgentRuntimesOutput) *ListAgentRuntimesResult
	GetData() *ListAgentRuntimesOutput
	SetRequestId(v string) *ListAgentRuntimesResult
	GetRequestId() *string
}

type ListAgentRuntimesResult struct {
	// SUCCESS 为成功，失败情况返回对应错误类型，比如 ERR_BAD_REQUEST ERR_VALIDATION_FAILED ERR_INTERNAL_SERVER_ERROR
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 智能体运行时列表的详细信息
	Data *ListAgentRuntimesOutput `json:"data,omitempty" xml:"data,omitempty"`
	// 唯一的请求标识符，用于问题追踪
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListAgentRuntimesResult) String() string {
	return dara.Prettify(s)
}

func (s ListAgentRuntimesResult) GoString() string {
	return s.String()
}

func (s *ListAgentRuntimesResult) GetCode() *string {
	return s.Code
}

func (s *ListAgentRuntimesResult) GetData() *ListAgentRuntimesOutput {
	return s.Data
}

func (s *ListAgentRuntimesResult) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAgentRuntimesResult) SetCode(v string) *ListAgentRuntimesResult {
	s.Code = &v
	return s
}

func (s *ListAgentRuntimesResult) SetData(v *ListAgentRuntimesOutput) *ListAgentRuntimesResult {
	s.Data = v
	return s
}

func (s *ListAgentRuntimesResult) SetRequestId(v string) *ListAgentRuntimesResult {
	s.RequestId = &v
	return s
}

func (s *ListAgentRuntimesResult) Validate() error {
	return dara.Validate(s)
}
