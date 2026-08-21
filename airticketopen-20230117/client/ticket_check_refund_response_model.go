// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTicketCheckRefundResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *TicketCheckRefundResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *TicketCheckRefundResponse
	GetStatusCode() *int32
	SetBody(v *TicketCheckRefundResponseBody) *TicketCheckRefundResponse
	GetBody() *TicketCheckRefundResponseBody
}

type TicketCheckRefundResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TicketCheckRefundResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TicketCheckRefundResponse) String() string {
	return dara.Prettify(s)
}

func (s TicketCheckRefundResponse) GoString() string {
	return s.String()
}

func (s *TicketCheckRefundResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *TicketCheckRefundResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *TicketCheckRefundResponse) GetBody() *TicketCheckRefundResponseBody {
	return s.Body
}

func (s *TicketCheckRefundResponse) SetHeaders(v map[string]*string) *TicketCheckRefundResponse {
	s.Headers = v
	return s
}

func (s *TicketCheckRefundResponse) SetStatusCode(v int32) *TicketCheckRefundResponse {
	s.StatusCode = &v
	return s
}

func (s *TicketCheckRefundResponse) SetBody(v *TicketCheckRefundResponseBody) *TicketCheckRefundResponse {
	s.Body = v
	return s
}

func (s *TicketCheckRefundResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
