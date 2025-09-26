// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAgentRuntimeVersionsResult interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListAgentRuntimeVersionsResult
	GetCode() *string
	SetData(v *ListAgentRuntimeVersionsOutput) *ListAgentRuntimeVersionsResult
	GetData() *ListAgentRuntimeVersionsOutput
	SetRequestId(v string) *ListAgentRuntimeVersionsResult
	GetRequestId() *string
}

type ListAgentRuntimeVersionsResult struct {
	// SUCCESS 为成功，失败情况返回对应错误类型，比如 ERR_BAD_REQUEST ERR_VALIDATION_FAILED ERR_INTERNAL_SERVER_ERROR
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 智能体运行时版本列表的详细信息
	Data *ListAgentRuntimeVersionsOutput `json:"data,omitempty" xml:"data,omitempty"`
	// 唯一的请求标识符，用于问题追踪
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListAgentRuntimeVersionsResult) String() string {
	return dara.Prettify(s)
}

func (s ListAgentRuntimeVersionsResult) GoString() string {
	return s.String()
}

func (s *ListAgentRuntimeVersionsResult) GetCode() *string {
	return s.Code
}

func (s *ListAgentRuntimeVersionsResult) GetData() *ListAgentRuntimeVersionsOutput {
	return s.Data
}

func (s *ListAgentRuntimeVersionsResult) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAgentRuntimeVersionsResult) SetCode(v string) *ListAgentRuntimeVersionsResult {
	s.Code = &v
	return s
}

func (s *ListAgentRuntimeVersionsResult) SetData(v *ListAgentRuntimeVersionsOutput) *ListAgentRuntimeVersionsResult {
	s.Data = v
	return s
}

func (s *ListAgentRuntimeVersionsResult) SetRequestId(v string) *ListAgentRuntimeVersionsResult {
	s.RequestId = &v
	return s
}

func (s *ListAgentRuntimeVersionsResult) Validate() error {
	return dara.Validate(s)
}
