// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTicketingCheckResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *TicketingCheckResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *TicketingCheckResponse
	GetStatusCode() *int32
	SetBody(v *TicketingCheckResponseBody) *TicketingCheckResponse
	GetBody() *TicketingCheckResponseBody
}

type TicketingCheckResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TicketingCheckResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TicketingCheckResponse) String() string {
	return dara.Prettify(s)
}

func (s TicketingCheckResponse) GoString() string {
	return s.String()
}

func (s *TicketingCheckResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *TicketingCheckResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *TicketingCheckResponse) GetBody() *TicketingCheckResponseBody {
	return s.Body
}

func (s *TicketingCheckResponse) SetHeaders(v map[string]*string) *TicketingCheckResponse {
	s.Headers = v
	return s
}

func (s *TicketingCheckResponse) SetStatusCode(v int32) *TicketingCheckResponse {
	s.StatusCode = &v
	return s
}

func (s *TicketingCheckResponse) SetBody(v *TicketingCheckResponseBody) *TicketingCheckResponse {
	s.Body = v
	return s
}

func (s *TicketingCheckResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
