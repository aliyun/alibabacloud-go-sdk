// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLibraryInstructionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateLibraryInstructionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateLibraryInstructionResponse
	GetStatusCode() *int32
	SetBody(v *LibraryInstruction) *CreateLibraryInstructionResponse
	GetBody() *LibraryInstruction
}

type CreateLibraryInstructionResponse struct {
	Headers    map[string]*string  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *LibraryInstruction `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateLibraryInstructionResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateLibraryInstructionResponse) GoString() string {
	return s.String()
}

func (s *CreateLibraryInstructionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateLibraryInstructionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateLibraryInstructionResponse) GetBody() *LibraryInstruction {
	return s.Body
}

func (s *CreateLibraryInstructionResponse) SetHeaders(v map[string]*string) *CreateLibraryInstructionResponse {
	s.Headers = v
	return s
}

func (s *CreateLibraryInstructionResponse) SetStatusCode(v int32) *CreateLibraryInstructionResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateLibraryInstructionResponse) SetBody(v *LibraryInstruction) *CreateLibraryInstructionResponse {
	s.Body = v
	return s
}

func (s *CreateLibraryInstructionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
