// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeployDeploymentDraftAsyncRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *DraftDeployParams) *DeployDeploymentDraftAsyncRequest
	GetBody() *DraftDeployParams
}

type DeployDeploymentDraftAsyncRequest struct {
	// This parameter is required.
	Body *DraftDeployParams `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeployDeploymentDraftAsyncRequest) String() string {
	return dara.Prettify(s)
}

func (s DeployDeploymentDraftAsyncRequest) GoString() string {
	return s.String()
}

func (s *DeployDeploymentDraftAsyncRequest) GetBody() *DraftDeployParams {
	return s.Body
}

func (s *DeployDeploymentDraftAsyncRequest) SetBody(v *DraftDeployParams) *DeployDeploymentDraftAsyncRequest {
	s.Body = v
	return s
}

func (s *DeployDeploymentDraftAsyncRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
