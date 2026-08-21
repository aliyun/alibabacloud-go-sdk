// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTicketCreateOrderResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *TicketCreateOrderResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *TicketCreateOrderResponse
	GetStatusCode() *int32
	SetBody(v *TicketCreateOrderResponseBody) *TicketCreateOrderResponse
	GetBody() *TicketCreateOrderResponseBody
}

type TicketCreateOrderResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TicketCreateOrderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TicketCreateOrderResponse) String() string {
	return dara.Prettify(s)
}

func (s TicketCreateOrderResponse) GoString() string {
	return s.String()
}

func (s *TicketCreateOrderResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *TicketCreateOrderResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *TicketCreateOrderResponse) GetBody() *TicketCreateOrderResponseBody {
	return s.Body
}

func (s *TicketCreateOrderResponse) SetHeaders(v map[string]*string) *TicketCreateOrderResponse {
	s.Headers = v
	return s
}

func (s *TicketCreateOrderResponse) SetStatusCode(v int32) *TicketCreateOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *TicketCreateOrderResponse) SetBody(v *TicketCreateOrderResponseBody) *TicketCreateOrderResponse {
	s.Body = v
	return s
}

func (s *TicketCreateOrderResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
