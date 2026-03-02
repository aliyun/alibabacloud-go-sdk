// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLibraryInstructionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateLibraryInstructionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateLibraryInstructionResponse
	GetStatusCode() *int32
	SetBody(v *LibraryInstruction) *UpdateLibraryInstructionResponse
	GetBody() *LibraryInstruction
}

type UpdateLibraryInstructionResponse struct {
	Headers    map[string]*string  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *LibraryInstruction `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateLibraryInstructionResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateLibraryInstructionResponse) GoString() string {
	return s.String()
}

func (s *UpdateLibraryInstructionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateLibraryInstructionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateLibraryInstructionResponse) GetBody() *LibraryInstruction {
	return s.Body
}

func (s *UpdateLibraryInstructionResponse) SetHeaders(v map[string]*string) *UpdateLibraryInstructionResponse {
	s.Headers = v
	return s
}

func (s *UpdateLibraryInstructionResponse) SetStatusCode(v int32) *UpdateLibraryInstructionResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateLibraryInstructionResponse) SetBody(v *LibraryInstruction) *UpdateLibraryInstructionResponse {
	s.Body = v
	return s
}

func (s *UpdateLibraryInstructionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
