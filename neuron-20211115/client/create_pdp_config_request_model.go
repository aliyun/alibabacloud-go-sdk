// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePdpConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *PdpConfig) *CreatePdpConfigRequest
	GetBody() *PdpConfig
}

type CreatePdpConfigRequest struct {
	// This parameter is required.
	Body *PdpConfig `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePdpConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePdpConfigRequest) GoString() string {
	return s.String()
}

func (s *CreatePdpConfigRequest) GetBody() *PdpConfig {
	return s.Body
}

func (s *CreatePdpConfigRequest) SetBody(v *PdpConfig) *CreatePdpConfigRequest {
	s.Body = v
	return s
}

func (s *CreatePdpConfigRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
