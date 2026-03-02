// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevertPdpServiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *DeploymentRevertCmd) *RevertPdpServiceRequest
	GetBody() *DeploymentRevertCmd
}

type RevertPdpServiceRequest struct {
	Body *DeploymentRevertCmd `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RevertPdpServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s RevertPdpServiceRequest) GoString() string {
	return s.String()
}

func (s *RevertPdpServiceRequest) GetBody() *DeploymentRevertCmd {
	return s.Body
}

func (s *RevertPdpServiceRequest) SetBody(v *DeploymentRevertCmd) *RevertPdpServiceRequest {
	s.Body = v
	return s
}

func (s *RevertPdpServiceRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
