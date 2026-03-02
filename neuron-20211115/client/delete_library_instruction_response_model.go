// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLibraryInstructionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteLibraryInstructionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteLibraryInstructionResponse
	GetStatusCode() *int32
}

type DeleteLibraryInstructionResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s DeleteLibraryInstructionResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteLibraryInstructionResponse) GoString() string {
	return s.String()
}

func (s *DeleteLibraryInstructionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteLibraryInstructionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteLibraryInstructionResponse) SetHeaders(v map[string]*string) *DeleteLibraryInstructionResponse {
	s.Headers = v
	return s
}

func (s *DeleteLibraryInstructionResponse) SetStatusCode(v int32) *DeleteLibraryInstructionResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteLibraryInstructionResponse) Validate() error {
	return dara.Validate(s)
}
