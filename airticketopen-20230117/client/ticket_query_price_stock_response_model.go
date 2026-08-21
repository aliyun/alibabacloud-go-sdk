// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTicketQueryPriceStockResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *TicketQueryPriceStockResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *TicketQueryPriceStockResponse
	GetStatusCode() *int32
	SetBody(v *TicketQueryPriceStockResponseBody) *TicketQueryPriceStockResponse
	GetBody() *TicketQueryPriceStockResponseBody
}

type TicketQueryPriceStockResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TicketQueryPriceStockResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TicketQueryPriceStockResponse) String() string {
	return dara.Prettify(s)
}

func (s TicketQueryPriceStockResponse) GoString() string {
	return s.String()
}

func (s *TicketQueryPriceStockResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *TicketQueryPriceStockResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *TicketQueryPriceStockResponse) GetBody() *TicketQueryPriceStockResponseBody {
	return s.Body
}

func (s *TicketQueryPriceStockResponse) SetHeaders(v map[string]*string) *TicketQueryPriceStockResponse {
	s.Headers = v
	return s
}

func (s *TicketQueryPriceStockResponse) SetStatusCode(v int32) *TicketQueryPriceStockResponse {
	s.StatusCode = &v
	return s
}

func (s *TicketQueryPriceStockResponse) SetBody(v *TicketQueryPriceStockResponseBody) *TicketQueryPriceStockResponse {
	s.Body = v
	return s
}

func (s *TicketQueryPriceStockResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
