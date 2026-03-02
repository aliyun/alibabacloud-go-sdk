// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateGreyPdpServiceGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *PdpGreyServiceGroupCreateCmd) *CreateGreyPdpServiceGroupRequest
	GetBody() *PdpGreyServiceGroupCreateCmd
}

type CreateGreyPdpServiceGroupRequest struct {
	// This parameter is required.
	Body *PdpGreyServiceGroupCreateCmd `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateGreyPdpServiceGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateGreyPdpServiceGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateGreyPdpServiceGroupRequest) GetBody() *PdpGreyServiceGroupCreateCmd {
	return s.Body
}

func (s *CreateGreyPdpServiceGroupRequest) SetBody(v *PdpGreyServiceGroupCreateCmd) *CreateGreyPdpServiceGroupRequest {
	s.Body = v
	return s
}

func (s *CreateGreyPdpServiceGroupRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
