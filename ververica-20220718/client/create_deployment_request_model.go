// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDeploymentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *Deployment) *CreateDeploymentRequest
	GetBody() *Deployment
}

type CreateDeploymentRequest struct {
	// The content of the deployment.
	//
	// This parameter is required.
	Body *Deployment `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDeploymentRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDeploymentRequest) GoString() string {
	return s.String()
}

func (s *CreateDeploymentRequest) GetBody() *Deployment {
	return s.Body
}

func (s *CreateDeploymentRequest) SetBody(v *Deployment) *CreateDeploymentRequest {
	s.Body = v
	return s
}

func (s *CreateDeploymentRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
