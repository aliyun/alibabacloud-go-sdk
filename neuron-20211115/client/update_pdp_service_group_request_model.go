// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePdpServiceGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *PdpServiceGroupUpdateCmd) *UpdatePdpServiceGroupRequest
	GetBody() *PdpServiceGroupUpdateCmd
}

type UpdatePdpServiceGroupRequest struct {
	// This parameter is required.
	Body *PdpServiceGroupUpdateCmd `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdatePdpServiceGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdatePdpServiceGroupRequest) GoString() string {
	return s.String()
}

func (s *UpdatePdpServiceGroupRequest) GetBody() *PdpServiceGroupUpdateCmd {
	return s.Body
}

func (s *UpdatePdpServiceGroupRequest) SetBody(v *PdpServiceGroupUpdateCmd) *UpdatePdpServiceGroupRequest {
	s.Body = v
	return s
}

func (s *UpdatePdpServiceGroupRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
