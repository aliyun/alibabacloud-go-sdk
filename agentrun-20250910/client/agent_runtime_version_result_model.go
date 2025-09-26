// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAgentRuntimeVersionResult interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *AgentRuntimeVersionResult
	GetCode() *string
	SetData(v *AgentRuntimeVersion) *AgentRuntimeVersionResult
	GetData() *AgentRuntimeVersion
	SetRequestId(v string) *AgentRuntimeVersionResult
	GetRequestId() *string
}

type AgentRuntimeVersionResult struct {
	// SUCCESS 为成功，失败情况返回对应错误类型，比如 ERR_BAD_REQUEST ERR_VALIDATION_FAILED ERR_INTERNAL_SERVER_ERROR
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 智能体运行时版本的详细信息
	Data *AgentRuntimeVersion `json:"data,omitempty" xml:"data,omitempty"`
	// 唯一的请求标识符，用于问题追踪
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s AgentRuntimeVersionResult) String() string {
	return dara.Prettify(s)
}

func (s AgentRuntimeVersionResult) GoString() string {
	return s.String()
}

func (s *AgentRuntimeVersionResult) GetCode() *string {
	return s.Code
}

func (s *AgentRuntimeVersionResult) GetData() *AgentRuntimeVersion {
	return s.Data
}

func (s *AgentRuntimeVersionResult) GetRequestId() *string {
	return s.RequestId
}

func (s *AgentRuntimeVersionResult) SetCode(v string) *AgentRuntimeVersionResult {
	s.Code = &v
	return s
}

func (s *AgentRuntimeVersionResult) SetData(v *AgentRuntimeVersion) *AgentRuntimeVersionResult {
	s.Data = v
	return s
}

func (s *AgentRuntimeVersionResult) SetRequestId(v string) *AgentRuntimeVersionResult {
	s.RequestId = &v
	return s
}

func (s *AgentRuntimeVersionResult) Validate() error {
	return dara.Validate(s)
}
