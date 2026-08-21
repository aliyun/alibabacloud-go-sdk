// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTicketQueryOrderResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *TicketQueryOrderResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *TicketQueryOrderResponse
	GetStatusCode() *int32
	SetBody(v *TicketQueryOrderResponseBody) *TicketQueryOrderResponse
	GetBody() *TicketQueryOrderResponseBody
}

type TicketQueryOrderResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TicketQueryOrderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TicketQueryOrderResponse) String() string {
	return dara.Prettify(s)
}

func (s TicketQueryOrderResponse) GoString() string {
	return s.String()
}

func (s *TicketQueryOrderResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *TicketQueryOrderResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *TicketQueryOrderResponse) GetBody() *TicketQueryOrderResponseBody {
	return s.Body
}

func (s *TicketQueryOrderResponse) SetHeaders(v map[string]*string) *TicketQueryOrderResponse {
	s.Headers = v
	return s
}

func (s *TicketQueryOrderResponse) SetStatusCode(v int32) *TicketQueryOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *TicketQueryOrderResponse) SetBody(v *TicketQueryOrderResponseBody) *TicketQueryOrderResponse {
	s.Body = v
	return s
}

func (s *TicketQueryOrderResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
