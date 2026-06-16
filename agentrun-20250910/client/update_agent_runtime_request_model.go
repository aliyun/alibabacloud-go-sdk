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
	// The input parameters for updating an agent runtime.
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
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
