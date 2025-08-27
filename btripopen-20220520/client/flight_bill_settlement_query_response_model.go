// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFlightBillSettlementQueryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *FlightBillSettlementQueryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *FlightBillSettlementQueryResponse
	GetStatusCode() *int32
	SetBody(v *FlightBillSettlementQueryResponseBody) *FlightBillSettlementQueryResponse
	GetBody() *FlightBillSettlementQueryResponseBody
}

type FlightBillSettlementQueryResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *FlightBillSettlementQueryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s FlightBillSettlementQueryResponse) String() string {
	return dara.Prettify(s)
}

func (s FlightBillSettlementQueryResponse) GoString() string {
	return s.String()
}

func (s *FlightBillSettlementQueryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *FlightBillSettlementQueryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *FlightBillSettlementQueryResponse) GetBody() *FlightBillSettlementQueryResponseBody {
	return s.Body
}

func (s *FlightBillSettlementQueryResponse) SetHeaders(v map[string]*string) *FlightBillSettlementQueryResponse {
	s.Headers = v
	return s
}

func (s *FlightBillSettlementQueryResponse) SetStatusCode(v int32) *FlightBillSettlementQueryResponse {
	s.StatusCode = &v
	return s
}

func (s *FlightBillSettlementQueryResponse) SetBody(v *FlightBillSettlementQueryResponseBody) *FlightBillSettlementQueryResponse {
	s.Body = v
	return s
}

func (s *FlightBillSettlementQueryResponse) Validate() error {
	return dara.Validate(s)
}
