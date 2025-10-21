// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDeploymentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *Deployment) *UpdateDeploymentRequest
	GetBody() *Deployment
}

type UpdateDeploymentRequest struct {
	// The information about the deployment that you want to update.
	//
	// This parameter is required.
	Body *Deployment `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateDeploymentRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateDeploymentRequest) GoString() string {
	return s.String()
}

func (s *UpdateDeploymentRequest) GetBody() *Deployment {
	return s.Body
}

func (s *UpdateDeploymentRequest) SetBody(v *Deployment) *UpdateDeploymentRequest {
	s.Body = v
	return s
}

func (s *UpdateDeploymentRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
