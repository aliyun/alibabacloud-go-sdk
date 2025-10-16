// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAgentRuntimeEndpointRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *CreateAgentRuntimeEndpointInput) *CreateAgentRuntimeEndpointRequest
	GetBody() *CreateAgentRuntimeEndpointInput
}

type CreateAgentRuntimeEndpointRequest struct {
	// 包含要创建的智能体运行时端点配置信息的请求体
	//
	// This parameter is required.
	Body *CreateAgentRuntimeEndpointInput `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAgentRuntimeEndpointRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAgentRuntimeEndpointRequest) GoString() string {
	return s.String()
}

func (s *CreateAgentRuntimeEndpointRequest) GetBody() *CreateAgentRuntimeEndpointInput {
	return s.Body
}

func (s *CreateAgentRuntimeEndpointRequest) SetBody(v *CreateAgentRuntimeEndpointInput) *CreateAgentRuntimeEndpointRequest {
	s.Body = v
	return s
}

func (s *CreateAgentRuntimeEndpointRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
