// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTicketQueryShelfResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *TicketQueryShelfResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *TicketQueryShelfResponse
	GetStatusCode() *int32
	SetBody(v *TicketQueryShelfResponseBody) *TicketQueryShelfResponse
	GetBody() *TicketQueryShelfResponseBody
}

type TicketQueryShelfResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TicketQueryShelfResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TicketQueryShelfResponse) String() string {
	return dara.Prettify(s)
}

func (s TicketQueryShelfResponse) GoString() string {
	return s.String()
}

func (s *TicketQueryShelfResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *TicketQueryShelfResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *TicketQueryShelfResponse) GetBody() *TicketQueryShelfResponseBody {
	return s.Body
}

func (s *TicketQueryShelfResponse) SetHeaders(v map[string]*string) *TicketQueryShelfResponse {
	s.Headers = v
	return s
}

func (s *TicketQueryShelfResponse) SetStatusCode(v int32) *TicketQueryShelfResponse {
	s.StatusCode = &v
	return s
}

func (s *TicketQueryShelfResponse) SetBody(v *TicketQueryShelfResponseBody) *TicketQueryShelfResponse {
	s.Body = v
	return s
}

func (s *TicketQueryShelfResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
