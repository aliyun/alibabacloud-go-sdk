// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRollbackPdpServiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *DeploymentRollbackCmd) *RollbackPdpServiceRequest
	GetBody() *DeploymentRollbackCmd
}

type RollbackPdpServiceRequest struct {
	// This parameter is required.
	Body *DeploymentRollbackCmd `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RollbackPdpServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s RollbackPdpServiceRequest) GoString() string {
	return s.String()
}

func (s *RollbackPdpServiceRequest) GetBody() *DeploymentRollbackCmd {
	return s.Body
}

func (s *RollbackPdpServiceRequest) SetBody(v *DeploymentRollbackCmd) *RollbackPdpServiceRequest {
	s.Body = v
	return s
}

func (s *RollbackPdpServiceRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
