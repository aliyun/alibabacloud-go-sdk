// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePdpConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *PdpConfigUpdateCmd) *UpdatePdpConfigRequest
	GetBody() *PdpConfigUpdateCmd
}

type UpdatePdpConfigRequest struct {
	// This parameter is required.
	Body *PdpConfigUpdateCmd `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdatePdpConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdatePdpConfigRequest) GoString() string {
	return s.String()
}

func (s *UpdatePdpConfigRequest) GetBody() *PdpConfigUpdateCmd {
	return s.Body
}

func (s *UpdatePdpConfigRequest) SetBody(v *PdpConfigUpdateCmd) *UpdatePdpConfigRequest {
	s.Body = v
	return s
}

func (s *UpdatePdpConfigRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
