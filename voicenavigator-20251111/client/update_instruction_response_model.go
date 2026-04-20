// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateInstructionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateInstructionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateInstructionResponse
	GetStatusCode() *int32
	SetBody(v *UpdateInstructionResponseBody) *UpdateInstructionResponse
	GetBody() *UpdateInstructionResponseBody
}

type UpdateInstructionResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateInstructionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateInstructionResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateInstructionResponse) GoString() string {
	return s.String()
}

func (s *UpdateInstructionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateInstructionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateInstructionResponse) GetBody() *UpdateInstructionResponseBody {
	return s.Body
}

func (s *UpdateInstructionResponse) SetHeaders(v map[string]*string) *UpdateInstructionResponse {
	s.Headers = v
	return s
}

func (s *UpdateInstructionResponse) SetStatusCode(v int32) *UpdateInstructionResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateInstructionResponse) SetBody(v *UpdateInstructionResponseBody) *UpdateInstructionResponse {
	s.Body = v
	return s
}

func (s *UpdateInstructionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
