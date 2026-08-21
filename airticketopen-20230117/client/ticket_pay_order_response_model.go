// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTicketPayOrderResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *TicketPayOrderResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *TicketPayOrderResponse
	GetStatusCode() *int32
	SetBody(v *TicketPayOrderResponseBody) *TicketPayOrderResponse
	GetBody() *TicketPayOrderResponseBody
}

type TicketPayOrderResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TicketPayOrderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TicketPayOrderResponse) String() string {
	return dara.Prettify(s)
}

func (s TicketPayOrderResponse) GoString() string {
	return s.String()
}

func (s *TicketPayOrderResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *TicketPayOrderResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *TicketPayOrderResponse) GetBody() *TicketPayOrderResponseBody {
	return s.Body
}

func (s *TicketPayOrderResponse) SetHeaders(v map[string]*string) *TicketPayOrderResponse {
	s.Headers = v
	return s
}

func (s *TicketPayOrderResponse) SetStatusCode(v int32) *TicketPayOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *TicketPayOrderResponse) SetBody(v *TicketPayOrderResponseBody) *TicketPayOrderResponse {
	s.Body = v
	return s
}

func (s *TicketPayOrderResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
