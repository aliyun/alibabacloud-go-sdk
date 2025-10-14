// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTicketingResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *TicketingResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *TicketingResponse
	GetStatusCode() *int32
	SetBody(v *TicketingResponseBody) *TicketingResponse
	GetBody() *TicketingResponseBody
}

type TicketingResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TicketingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TicketingResponse) String() string {
	return dara.Prettify(s)
}

func (s TicketingResponse) GoString() string {
	return s.String()
}

func (s *TicketingResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *TicketingResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *TicketingResponse) GetBody() *TicketingResponseBody {
	return s.Body
}

func (s *TicketingResponse) SetHeaders(v map[string]*string) *TicketingResponse {
	s.Headers = v
	return s
}

func (s *TicketingResponse) SetStatusCode(v int32) *TicketingResponse {
	s.StatusCode = &v
	return s
}

func (s *TicketingResponse) SetBody(v *TicketingResponseBody) *TicketingResponse {
	s.Body = v
	return s
}

func (s *TicketingResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
