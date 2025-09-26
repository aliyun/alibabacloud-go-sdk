// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAgentRuntimeEndpointsResult interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListAgentRuntimeEndpointsResult
	GetCode() *string
	SetData(v *ListAgentRuntimeEndpointsOutput) *ListAgentRuntimeEndpointsResult
	GetData() *ListAgentRuntimeEndpointsOutput
	SetRequestId(v string) *ListAgentRuntimeEndpointsResult
	GetRequestId() *string
}

type ListAgentRuntimeEndpointsResult struct {
	// SUCCESS 为成功，失败情况返回对应错误类型，比如 ERR_BAD_REQUEST ERR_VALIDATION_FAILED ERR_INTERNAL_SERVER_ERROR
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 智能体运行时端点列表的详细信息
	Data *ListAgentRuntimeEndpointsOutput `json:"data,omitempty" xml:"data,omitempty"`
	// 唯一的请求标识符，用于问题追踪
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListAgentRuntimeEndpointsResult) String() string {
	return dara.Prettify(s)
}

func (s ListAgentRuntimeEndpointsResult) GoString() string {
	return s.String()
}

func (s *ListAgentRuntimeEndpointsResult) GetCode() *string {
	return s.Code
}

func (s *ListAgentRuntimeEndpointsResult) GetData() *ListAgentRuntimeEndpointsOutput {
	return s.Data
}

func (s *ListAgentRuntimeEndpointsResult) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAgentRuntimeEndpointsResult) SetCode(v string) *ListAgentRuntimeEndpointsResult {
	s.Code = &v
	return s
}

func (s *ListAgentRuntimeEndpointsResult) SetData(v *ListAgentRuntimeEndpointsOutput) *ListAgentRuntimeEndpointsResult {
	s.Data = v
	return s
}

func (s *ListAgentRuntimeEndpointsResult) SetRequestId(v string) *ListAgentRuntimeEndpointsResult {
	s.RequestId = &v
	return s
}

func (s *ListAgentRuntimeEndpointsResult) Validate() error {
	return dara.Validate(s)
}
