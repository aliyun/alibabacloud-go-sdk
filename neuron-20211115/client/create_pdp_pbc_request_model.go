// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePdpPbcRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *PdpPbc) *CreatePdpPbcRequest
	GetBody() *PdpPbc
}

type CreatePdpPbcRequest struct {
	// This parameter is required.
	Body *PdpPbc `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePdpPbcRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePdpPbcRequest) GoString() string {
	return s.String()
}

func (s *CreatePdpPbcRequest) GetBody() *PdpPbc {
	return s.Body
}

func (s *CreatePdpPbcRequest) SetBody(v *PdpPbc) *CreatePdpPbcRequest {
	s.Body = v
	return s
}

func (s *CreatePdpPbcRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
