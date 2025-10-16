// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAgentRuntimeEndpointRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *UpdateAgentRuntimeEndpointInput) *UpdateAgentRuntimeEndpointRequest
	GetBody() *UpdateAgentRuntimeEndpointInput
}

type UpdateAgentRuntimeEndpointRequest struct {
	// 包含要更新的智能体运行时端点配置信息的请求体
	//
	// This parameter is required.
	Body *UpdateAgentRuntimeEndpointInput `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAgentRuntimeEndpointRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAgentRuntimeEndpointRequest) GoString() string {
	return s.String()
}

func (s *UpdateAgentRuntimeEndpointRequest) GetBody() *UpdateAgentRuntimeEndpointInput {
	return s.Body
}

func (s *UpdateAgentRuntimeEndpointRequest) SetBody(v *UpdateAgentRuntimeEndpointInput) *UpdateAgentRuntimeEndpointRequest {
	s.Body = v
	return s
}

func (s *UpdateAgentRuntimeEndpointRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
