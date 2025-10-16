// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAgentRuntimeEndpointResult interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *AgentRuntimeEndpointResult
	GetCode() *string
	SetData(v *AgentRuntimeEndpoint) *AgentRuntimeEndpointResult
	GetData() *AgentRuntimeEndpoint
	SetRequestId(v string) *AgentRuntimeEndpointResult
	GetRequestId() *string
}

type AgentRuntimeEndpointResult struct {
	// SUCCESS 为成功，失败情况返回对应错误类型，比如 ERR_BAD_REQUEST ERR_VALIDATION_FAILED ERR_INTERNAL_SERVER_ERROR
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 智能体运行时端点的详细信息
	Data *AgentRuntimeEndpoint `json:"data,omitempty" xml:"data,omitempty"`
	// 唯一的请求标识符，用于问题追踪
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s AgentRuntimeEndpointResult) String() string {
	return dara.Prettify(s)
}

func (s AgentRuntimeEndpointResult) GoString() string {
	return s.String()
}

func (s *AgentRuntimeEndpointResult) GetCode() *string {
	return s.Code
}

func (s *AgentRuntimeEndpointResult) GetData() *AgentRuntimeEndpoint {
	return s.Data
}

func (s *AgentRuntimeEndpointResult) GetRequestId() *string {
	return s.RequestId
}

func (s *AgentRuntimeEndpointResult) SetCode(v string) *AgentRuntimeEndpointResult {
	s.Code = &v
	return s
}

func (s *AgentRuntimeEndpointResult) SetData(v *AgentRuntimeEndpoint) *AgentRuntimeEndpointResult {
	s.Data = v
	return s
}

func (s *AgentRuntimeEndpointResult) SetRequestId(v string) *AgentRuntimeEndpointResult {
	s.RequestId = &v
	return s
}

func (s *AgentRuntimeEndpointResult) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
