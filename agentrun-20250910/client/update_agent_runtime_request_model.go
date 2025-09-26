// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAgentRuntimeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *UpdateAgentRuntimeInput) *UpdateAgentRuntimeRequest
	GetBody() *UpdateAgentRuntimeInput
}

type UpdateAgentRuntimeRequest struct {
	// 包含要更新的智能体运行时配置信息的请求体
	//
	// This parameter is required.
	Body *UpdateAgentRuntimeInput `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAgentRuntimeRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAgentRuntimeRequest) GoString() string {
	return s.String()
}

func (s *UpdateAgentRuntimeRequest) GetBody() *UpdateAgentRuntimeInput {
	return s.Body
}

func (s *UpdateAgentRuntimeRequest) SetBody(v *UpdateAgentRuntimeInput) *UpdateAgentRuntimeRequest {
	s.Body = v
	return s
}

func (s *UpdateAgentRuntimeRequest) Validate() error {
	return dara.Validate(s)
}
