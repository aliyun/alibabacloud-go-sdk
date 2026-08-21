// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTicketQueryRefundOrderResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *TicketQueryRefundOrderResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *TicketQueryRefundOrderResponse
	GetStatusCode() *int32
	SetBody(v *TicketQueryRefundOrderResponseBody) *TicketQueryRefundOrderResponse
	GetBody() *TicketQueryRefundOrderResponseBody
}

type TicketQueryRefundOrderResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TicketQueryRefundOrderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TicketQueryRefundOrderResponse) String() string {
	return dara.Prettify(s)
}

func (s TicketQueryRefundOrderResponse) GoString() string {
	return s.String()
}

func (s *TicketQueryRefundOrderResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *TicketQueryRefundOrderResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *TicketQueryRefundOrderResponse) GetBody() *TicketQueryRefundOrderResponseBody {
	return s.Body
}

func (s *TicketQueryRefundOrderResponse) SetHeaders(v map[string]*string) *TicketQueryRefundOrderResponse {
	s.Headers = v
	return s
}

func (s *TicketQueryRefundOrderResponse) SetStatusCode(v int32) *TicketQueryRefundOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *TicketQueryRefundOrderResponse) SetBody(v *TicketQueryRefundOrderResponseBody) *TicketQueryRefundOrderResponse {
	s.Body = v
	return s
}

func (s *TicketQueryRefundOrderResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
