// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteInstructionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteInstructionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteInstructionResponse
	GetStatusCode() *int32
	SetBody(v *DeleteInstructionResponseBody) *DeleteInstructionResponse
	GetBody() *DeleteInstructionResponseBody
}

type DeleteInstructionResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteInstructionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteInstructionResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteInstructionResponse) GoString() string {
	return s.String()
}

func (s *DeleteInstructionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteInstructionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteInstructionResponse) GetBody() *DeleteInstructionResponseBody {
	return s.Body
}

func (s *DeleteInstructionResponse) SetHeaders(v map[string]*string) *DeleteInstructionResponse {
	s.Headers = v
	return s
}

func (s *DeleteInstructionResponse) SetStatusCode(v int32) *DeleteInstructionResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteInstructionResponse) SetBody(v *DeleteInstructionResponseBody) *DeleteInstructionResponse {
	s.Body = v
	return s
}

func (s *DeleteInstructionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
