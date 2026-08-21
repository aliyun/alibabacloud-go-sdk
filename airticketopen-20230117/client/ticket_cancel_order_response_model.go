// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTicketCancelOrderResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *TicketCancelOrderResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *TicketCancelOrderResponse
	GetStatusCode() *int32
	SetBody(v *TicketCancelOrderResponseBody) *TicketCancelOrderResponse
	GetBody() *TicketCancelOrderResponseBody
}

type TicketCancelOrderResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TicketCancelOrderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TicketCancelOrderResponse) String() string {
	return dara.Prettify(s)
}

func (s TicketCancelOrderResponse) GoString() string {
	return s.String()
}

func (s *TicketCancelOrderResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *TicketCancelOrderResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *TicketCancelOrderResponse) GetBody() *TicketCancelOrderResponseBody {
	return s.Body
}

func (s *TicketCancelOrderResponse) SetHeaders(v map[string]*string) *TicketCancelOrderResponse {
	s.Headers = v
	return s
}

func (s *TicketCancelOrderResponse) SetStatusCode(v int32) *TicketCancelOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *TicketCancelOrderResponse) SetBody(v *TicketCancelOrderResponseBody) *TicketCancelOrderResponse {
	s.Body = v
	return s
}

func (s *TicketCancelOrderResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
