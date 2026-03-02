// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePdpServiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *PdpService) *CreatePdpServiceRequest
	GetBody() *PdpService
}

type CreatePdpServiceRequest struct {
	// This parameter is required.
	Body *PdpService `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePdpServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePdpServiceRequest) GoString() string {
	return s.String()
}

func (s *CreatePdpServiceRequest) GetBody() *PdpService {
	return s.Body
}

func (s *CreatePdpServiceRequest) SetBody(v *PdpService) *CreatePdpServiceRequest {
	s.Body = v
	return s
}

func (s *CreatePdpServiceRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
