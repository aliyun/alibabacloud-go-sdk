// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTriggerDeploymentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *DeploymentTriggerCmd) *TriggerDeploymentRequest
	GetBody() *DeploymentTriggerCmd
}

type TriggerDeploymentRequest struct {
	// This parameter is required.
	Body *DeploymentTriggerCmd `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TriggerDeploymentRequest) String() string {
	return dara.Prettify(s)
}

func (s TriggerDeploymentRequest) GoString() string {
	return s.String()
}

func (s *TriggerDeploymentRequest) GetBody() *DeploymentTriggerCmd {
	return s.Body
}

func (s *TriggerDeploymentRequest) SetBody(v *DeploymentTriggerCmd) *TriggerDeploymentRequest {
	s.Body = v
	return s
}

func (s *TriggerDeploymentRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
