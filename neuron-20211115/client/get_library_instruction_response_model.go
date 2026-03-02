// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLibraryInstructionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetLibraryInstructionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetLibraryInstructionResponse
	GetStatusCode() *int32
	SetBody(v *LibraryInstruction) *GetLibraryInstructionResponse
	GetBody() *LibraryInstruction
}

type GetLibraryInstructionResponse struct {
	Headers    map[string]*string  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *LibraryInstruction `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetLibraryInstructionResponse) String() string {
	return dara.Prettify(s)
}

func (s GetLibraryInstructionResponse) GoString() string {
	return s.String()
}

func (s *GetLibraryInstructionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetLibraryInstructionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetLibraryInstructionResponse) GetBody() *LibraryInstruction {
	return s.Body
}

func (s *GetLibraryInstructionResponse) SetHeaders(v map[string]*string) *GetLibraryInstructionResponse {
	s.Headers = v
	return s
}

func (s *GetLibraryInstructionResponse) SetStatusCode(v int32) *GetLibraryInstructionResponse {
	s.StatusCode = &v
	return s
}

func (s *GetLibraryInstructionResponse) SetBody(v *LibraryInstruction) *GetLibraryInstructionResponse {
	s.Body = v
	return s
}

func (s *GetLibraryInstructionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
