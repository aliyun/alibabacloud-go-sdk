// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDeploymentDraftRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *DeploymentDraft) *CreateDeploymentDraftRequest
	GetBody() *DeploymentDraft
}

type CreateDeploymentDraftRequest struct {
	// This parameter is required.
	Body *DeploymentDraft `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDeploymentDraftRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDeploymentDraftRequest) GoString() string {
	return s.String()
}

func (s *CreateDeploymentDraftRequest) GetBody() *DeploymentDraft {
	return s.Body
}

func (s *CreateDeploymentDraftRequest) SetBody(v *DeploymentDraft) *CreateDeploymentDraftRequest {
	s.Body = v
	return s
}

func (s *CreateDeploymentDraftRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
