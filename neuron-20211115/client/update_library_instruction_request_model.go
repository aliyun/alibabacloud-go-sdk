// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLibraryInstructionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *LibraryInstruction) *UpdateLibraryInstructionRequest
	GetBody() *LibraryInstruction
}

type UpdateLibraryInstructionRequest struct {
	Body *LibraryInstruction `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateLibraryInstructionRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateLibraryInstructionRequest) GoString() string {
	return s.String()
}

func (s *UpdateLibraryInstructionRequest) GetBody() *LibraryInstruction {
	return s.Body
}

func (s *UpdateLibraryInstructionRequest) SetBody(v *LibraryInstruction) *UpdateLibraryInstructionRequest {
	s.Body = v
	return s
}

func (s *UpdateLibraryInstructionRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
