// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDeploymentTargetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *ResourceSpec) *UpdateDeploymentTargetRequest
	GetBody() *ResourceSpec
}

type UpdateDeploymentTargetRequest struct {
	Body *ResourceSpec `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateDeploymentTargetRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateDeploymentTargetRequest) GoString() string {
	return s.String()
}

func (s *UpdateDeploymentTargetRequest) GetBody() *ResourceSpec {
	return s.Body
}

func (s *UpdateDeploymentTargetRequest) SetBody(v *ResourceSpec) *UpdateDeploymentTargetRequest {
	s.Body = v
	return s
}

func (s *UpdateDeploymentTargetRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
