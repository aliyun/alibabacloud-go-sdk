// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAgentRuntimeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *CreateAgentRuntimeInput) *CreateAgentRuntimeRequest
	GetBody() *CreateAgentRuntimeInput
}

type CreateAgentRuntimeRequest struct {
	// Request parameters for creating an agent runtime
	//
	// This parameter is required.
	Body *CreateAgentRuntimeInput `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAgentRuntimeRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAgentRuntimeRequest) GoString() string {
	return s.String()
}

func (s *CreateAgentRuntimeRequest) GetBody() *CreateAgentRuntimeInput {
	return s.Body
}

func (s *CreateAgentRuntimeRequest) SetBody(v *CreateAgentRuntimeInput) *CreateAgentRuntimeRequest {
	s.Body = v
	return s
}

func (s *CreateAgentRuntimeRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
