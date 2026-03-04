// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWorkspaceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *CreateWorkspaceInput) *CreateWorkspaceRequest
	GetBody() *CreateWorkspaceInput
}

type CreateWorkspaceRequest struct {
	Body *CreateWorkspaceInput `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateWorkspaceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateWorkspaceRequest) GoString() string {
	return s.String()
}

func (s *CreateWorkspaceRequest) GetBody() *CreateWorkspaceInput {
	return s.Body
}

func (s *CreateWorkspaceRequest) SetBody(v *CreateWorkspaceInput) *CreateWorkspaceRequest {
	s.Body = v
	return s
}

func (s *CreateWorkspaceRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
