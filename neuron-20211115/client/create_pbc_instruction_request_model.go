// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePbcInstructionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *PbcInstructionCreateCmd) *CreatePbcInstructionRequest
	GetBody() *PbcInstructionCreateCmd
}

type CreatePbcInstructionRequest struct {
	Body *PbcInstructionCreateCmd `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePbcInstructionRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePbcInstructionRequest) GoString() string {
	return s.String()
}

func (s *CreatePbcInstructionRequest) GetBody() *PbcInstructionCreateCmd {
	return s.Body
}

func (s *CreatePbcInstructionRequest) SetBody(v *PbcInstructionCreateCmd) *CreatePbcInstructionRequest {
	s.Body = v
	return s
}

func (s *CreatePbcInstructionRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
