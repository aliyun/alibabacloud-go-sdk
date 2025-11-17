// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTicket4CopilotResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateTicket4CopilotResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateTicket4CopilotResponse
	GetStatusCode() *int32
	SetBody(v *CreateTicket4CopilotResponseBody) *CreateTicket4CopilotResponse
	GetBody() *CreateTicket4CopilotResponseBody
}

type CreateTicket4CopilotResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateTicket4CopilotResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateTicket4CopilotResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateTicket4CopilotResponse) GoString() string {
	return s.String()
}

func (s *CreateTicket4CopilotResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateTicket4CopilotResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateTicket4CopilotResponse) GetBody() *CreateTicket4CopilotResponseBody {
	return s.Body
}

func (s *CreateTicket4CopilotResponse) SetHeaders(v map[string]*string) *CreateTicket4CopilotResponse {
	s.Headers = v
	return s
}

func (s *CreateTicket4CopilotResponse) SetStatusCode(v int32) *CreateTicket4CopilotResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateTicket4CopilotResponse) SetBody(v *CreateTicket4CopilotResponseBody) *CreateTicket4CopilotResponse {
	s.Body = v
	return s
}

func (s *CreateTicket4CopilotResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
