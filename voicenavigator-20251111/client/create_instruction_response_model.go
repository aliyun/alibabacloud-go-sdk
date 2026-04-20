// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateInstructionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateInstructionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateInstructionResponse
	GetStatusCode() *int32
	SetBody(v *CreateInstructionResponseBody) *CreateInstructionResponse
	GetBody() *CreateInstructionResponseBody
}

type CreateInstructionResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateInstructionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateInstructionResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateInstructionResponse) GoString() string {
	return s.String()
}

func (s *CreateInstructionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateInstructionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateInstructionResponse) GetBody() *CreateInstructionResponseBody {
	return s.Body
}

func (s *CreateInstructionResponse) SetHeaders(v map[string]*string) *CreateInstructionResponse {
	s.Headers = v
	return s
}

func (s *CreateInstructionResponse) SetStatusCode(v int32) *CreateInstructionResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateInstructionResponse) SetBody(v *CreateInstructionResponseBody) *CreateInstructionResponse {
	s.Body = v
	return s
}

func (s *CreateInstructionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
