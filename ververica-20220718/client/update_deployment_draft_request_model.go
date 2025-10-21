// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDeploymentDraftRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *DeploymentDraft) *UpdateDeploymentDraftRequest
	GetBody() *DeploymentDraft
}

type UpdateDeploymentDraftRequest struct {
	// This parameter is required.
	Body *DeploymentDraft `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateDeploymentDraftRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateDeploymentDraftRequest) GoString() string {
	return s.String()
}

func (s *UpdateDeploymentDraftRequest) GetBody() *DeploymentDraft {
	return s.Body
}

func (s *UpdateDeploymentDraftRequest) SetBody(v *DeploymentDraft) *UpdateDeploymentDraftRequest {
	s.Body = v
	return s
}

func (s *UpdateDeploymentDraftRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
