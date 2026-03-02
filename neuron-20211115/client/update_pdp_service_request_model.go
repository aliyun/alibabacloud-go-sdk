// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePdpServiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *PdpServiceUpdateCmd) *UpdatePdpServiceRequest
	GetBody() *PdpServiceUpdateCmd
}

type UpdatePdpServiceRequest struct {
	// This parameter is required.
	Body *PdpServiceUpdateCmd `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdatePdpServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdatePdpServiceRequest) GoString() string {
	return s.String()
}

func (s *UpdatePdpServiceRequest) GetBody() *PdpServiceUpdateCmd {
	return s.Body
}

func (s *UpdatePdpServiceRequest) SetBody(v *PdpServiceUpdateCmd) *UpdatePdpServiceRequest {
	s.Body = v
	return s
}

func (s *UpdatePdpServiceRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
