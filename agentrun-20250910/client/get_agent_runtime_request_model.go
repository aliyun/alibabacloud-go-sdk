// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAgentRuntimeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentRuntimeVersion(v string) *GetAgentRuntimeRequest
	GetAgentRuntimeVersion() *string
}

type GetAgentRuntimeRequest struct {
	// 指定要获取的智能体运行时版本，如果不指定则返回最新版本
	//
	// example:
	//
	// v1.0.0
	AgentRuntimeVersion *string `json:"agentRuntimeVersion,omitempty" xml:"agentRuntimeVersion,omitempty"`
}

func (s GetAgentRuntimeRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAgentRuntimeRequest) GoString() string {
	return s.String()
}

func (s *GetAgentRuntimeRequest) GetAgentRuntimeVersion() *string {
	return s.AgentRuntimeVersion
}

func (s *GetAgentRuntimeRequest) SetAgentRuntimeVersion(v string) *GetAgentRuntimeRequest {
	s.AgentRuntimeVersion = &v
	return s
}

func (s *GetAgentRuntimeRequest) Validate() error {
	return dara.Validate(s)
}
