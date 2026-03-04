// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWorkspaceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *UpdateWorkspaceInput) *UpdateWorkspaceRequest
	GetBody() *UpdateWorkspaceInput
}

type UpdateWorkspaceRequest struct {
	Body *UpdateWorkspaceInput `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateWorkspaceRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateWorkspaceRequest) GoString() string {
	return s.String()
}

func (s *UpdateWorkspaceRequest) GetBody() *UpdateWorkspaceInput {
	return s.Body
}

func (s *UpdateWorkspaceRequest) SetBody(v *UpdateWorkspaceInput) *UpdateWorkspaceRequest {
	s.Body = v
	return s
}

func (s *UpdateWorkspaceRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
