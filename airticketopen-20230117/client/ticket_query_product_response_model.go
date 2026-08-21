// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTicketQueryProductResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *TicketQueryProductResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *TicketQueryProductResponse
	GetStatusCode() *int32
	SetBody(v *TicketQueryProductResponseBody) *TicketQueryProductResponse
	GetBody() *TicketQueryProductResponseBody
}

type TicketQueryProductResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TicketQueryProductResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TicketQueryProductResponse) String() string {
	return dara.Prettify(s)
}

func (s TicketQueryProductResponse) GoString() string {
	return s.String()
}

func (s *TicketQueryProductResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *TicketQueryProductResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *TicketQueryProductResponse) GetBody() *TicketQueryProductResponseBody {
	return s.Body
}

func (s *TicketQueryProductResponse) SetHeaders(v map[string]*string) *TicketQueryProductResponse {
	s.Headers = v
	return s
}

func (s *TicketQueryProductResponse) SetStatusCode(v int32) *TicketQueryProductResponse {
	s.StatusCode = &v
	return s
}

func (s *TicketQueryProductResponse) SetBody(v *TicketQueryProductResponseBody) *TicketQueryProductResponse {
	s.Body = v
	return s
}

func (s *TicketQueryProductResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
