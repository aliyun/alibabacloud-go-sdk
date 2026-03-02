// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePdpLaneRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *PdpLaneUpdateCmd) *UpdatePdpLaneRequest
	GetBody() *PdpLaneUpdateCmd
}

type UpdatePdpLaneRequest struct {
	// This parameter is required.
	Body *PdpLaneUpdateCmd `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdatePdpLaneRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdatePdpLaneRequest) GoString() string {
	return s.String()
}

func (s *UpdatePdpLaneRequest) GetBody() *PdpLaneUpdateCmd {
	return s.Body
}

func (s *UpdatePdpLaneRequest) SetBody(v *PdpLaneUpdateCmd) *UpdatePdpLaneRequest {
	s.Body = v
	return s
}

func (s *UpdatePdpLaneRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
