// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePdpPbcRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *PdpPbcUpdateCmd) *UpdatePdpPbcRequest
	GetBody() *PdpPbcUpdateCmd
}

type UpdatePdpPbcRequest struct {
	Body *PdpPbcUpdateCmd `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdatePdpPbcRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdatePdpPbcRequest) GoString() string {
	return s.String()
}

func (s *UpdatePdpPbcRequest) GetBody() *PdpPbcUpdateCmd {
	return s.Body
}

func (s *UpdatePdpPbcRequest) SetBody(v *PdpPbcUpdateCmd) *UpdatePdpPbcRequest {
	s.Body = v
	return s
}

func (s *UpdatePdpPbcRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
