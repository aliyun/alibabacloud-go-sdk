// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePbcInstructionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *PbcInstructionUpdateCmd) *UpdatePbcInstructionRequest
	GetBody() *PbcInstructionUpdateCmd
}

type UpdatePbcInstructionRequest struct {
	Body *PbcInstructionUpdateCmd `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdatePbcInstructionRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdatePbcInstructionRequest) GoString() string {
	return s.String()
}

func (s *UpdatePbcInstructionRequest) GetBody() *PbcInstructionUpdateCmd {
	return s.Body
}

func (s *UpdatePbcInstructionRequest) SetBody(v *PbcInstructionUpdateCmd) *UpdatePbcInstructionRequest {
	s.Body = v
	return s
}

func (s *UpdatePbcInstructionRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
