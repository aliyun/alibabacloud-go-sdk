// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAgentRuntimeResult interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *AgentRuntimeResult
	GetCode() *string
	SetData(v *AgentRuntime) *AgentRuntimeResult
	GetData() *AgentRuntime
	SetRequestId(v string) *AgentRuntimeResult
	GetRequestId() *string
}

type AgentRuntimeResult struct {
	// SUCCESS 为成功，失败情况返回对应错误类型，比如 ERR_BAD_REQUEST ERR_VALIDATION_FAILED ERR_INTERNAL_SERVER_ERROR
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 智能体运行时的详细信息
	Data *AgentRuntime `json:"data,omitempty" xml:"data,omitempty"`
	// 唯一的请求标识符，用于问题追踪
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s AgentRuntimeResult) String() string {
	return dara.Prettify(s)
}

func (s AgentRuntimeResult) GoString() string {
	return s.String()
}

func (s *AgentRuntimeResult) GetCode() *string {
	return s.Code
}

func (s *AgentRuntimeResult) GetData() *AgentRuntime {
	return s.Data
}

func (s *AgentRuntimeResult) GetRequestId() *string {
	return s.RequestId
}

func (s *AgentRuntimeResult) SetCode(v string) *AgentRuntimeResult {
	s.Code = &v
	return s
}

func (s *AgentRuntimeResult) SetData(v *AgentRuntime) *AgentRuntimeResult {
	s.Data = v
	return s
}

func (s *AgentRuntimeResult) SetRequestId(v string) *AgentRuntimeResult {
	s.RequestId = &v
	return s
}

func (s *AgentRuntimeResult) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
