// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iContinueDeploymentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *DeploymentContinueCmd) *ContinueDeploymentRequest
	GetBody() *DeploymentContinueCmd
}

type ContinueDeploymentRequest struct {
	// This parameter is required.
	Body *DeploymentContinueCmd `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ContinueDeploymentRequest) String() string {
	return dara.Prettify(s)
}

func (s ContinueDeploymentRequest) GoString() string {
	return s.String()
}

func (s *ContinueDeploymentRequest) GetBody() *DeploymentContinueCmd {
	return s.Body
}

func (s *ContinueDeploymentRequest) SetBody(v *DeploymentContinueCmd) *ContinueDeploymentRequest {
	s.Body = v
	return s
}

func (s *ContinueDeploymentRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
