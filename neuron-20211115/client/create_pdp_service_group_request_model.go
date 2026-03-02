// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePdpServiceGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *PdpServiceGroupCreateCmd) *CreatePdpServiceGroupRequest
	GetBody() *PdpServiceGroupCreateCmd
}

type CreatePdpServiceGroupRequest struct {
	// This parameter is required.
	Body *PdpServiceGroupCreateCmd `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePdpServiceGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePdpServiceGroupRequest) GoString() string {
	return s.String()
}

func (s *CreatePdpServiceGroupRequest) GetBody() *PdpServiceGroupCreateCmd {
	return s.Body
}

func (s *CreatePdpServiceGroupRequest) SetBody(v *PdpServiceGroupCreateCmd) *CreatePdpServiceGroupRequest {
	s.Body = v
	return s
}

func (s *CreatePdpServiceGroupRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
