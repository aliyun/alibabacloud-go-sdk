// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTicket4AiProResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateTicket4AiProResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateTicket4AiProResponse
	GetStatusCode() *int32
	SetBody(v *CreateTicket4AiProResponseBody) *CreateTicket4AiProResponse
	GetBody() *CreateTicket4AiProResponseBody
}

type CreateTicket4AiProResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateTicket4AiProResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateTicket4AiProResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateTicket4AiProResponse) GoString() string {
	return s.String()
}

func (s *CreateTicket4AiProResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateTicket4AiProResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateTicket4AiProResponse) GetBody() *CreateTicket4AiProResponseBody {
	return s.Body
}

func (s *CreateTicket4AiProResponse) SetHeaders(v map[string]*string) *CreateTicket4AiProResponse {
	s.Headers = v
	return s
}

func (s *CreateTicket4AiProResponse) SetStatusCode(v int32) *CreateTicket4AiProResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateTicket4AiProResponse) SetBody(v *CreateTicket4AiProResponseBody) *CreateTicket4AiProResponse {
	s.Body = v
	return s
}

func (s *CreateTicket4AiProResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
