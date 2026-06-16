// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSandboxRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *CreateSandboxInput) *CreateSandboxRequest
	GetBody() *CreateSandboxInput
}

type CreateSandboxRequest struct {
	// The configuration for the sandbox.
	//
	// This parameter is required.
	Body *CreateSandboxInput `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSandboxRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSandboxRequest) GoString() string {
	return s.String()
}

func (s *CreateSandboxRequest) GetBody() *CreateSandboxInput {
	return s.Body
}

func (s *CreateSandboxRequest) SetBody(v *CreateSandboxInput) *CreateSandboxRequest {
	s.Body = v
	return s
}

func (s *CreateSandboxRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
