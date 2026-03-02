// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePdpLaneRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *PdpLaneCreateCmd) *CreatePdpLaneRequest
	GetBody() *PdpLaneCreateCmd
}

type CreatePdpLaneRequest struct {
	// This parameter is required.
	Body *PdpLaneCreateCmd `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePdpLaneRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePdpLaneRequest) GoString() string {
	return s.String()
}

func (s *CreatePdpLaneRequest) GetBody() *PdpLaneCreateCmd {
	return s.Body
}

func (s *CreatePdpLaneRequest) SetBody(v *PdpLaneCreateCmd) *CreatePdpLaneRequest {
	s.Body = v
	return s
}

func (s *CreatePdpLaneRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
