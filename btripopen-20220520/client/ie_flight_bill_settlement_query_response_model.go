// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIeFlightBillSettlementQueryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *IeFlightBillSettlementQueryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *IeFlightBillSettlementQueryResponse
	GetStatusCode() *int32
	SetBody(v *IeFlightBillSettlementQueryResponseBody) *IeFlightBillSettlementQueryResponse
	GetBody() *IeFlightBillSettlementQueryResponseBody
}

type IeFlightBillSettlementQueryResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *IeFlightBillSettlementQueryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s IeFlightBillSettlementQueryResponse) String() string {
	return dara.Prettify(s)
}

func (s IeFlightBillSettlementQueryResponse) GoString() string {
	return s.String()
}

func (s *IeFlightBillSettlementQueryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *IeFlightBillSettlementQueryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *IeFlightBillSettlementQueryResponse) GetBody() *IeFlightBillSettlementQueryResponseBody {
	return s.Body
}

func (s *IeFlightBillSettlementQueryResponse) SetHeaders(v map[string]*string) *IeFlightBillSettlementQueryResponse {
	s.Headers = v
	return s
}

func (s *IeFlightBillSettlementQueryResponse) SetStatusCode(v int32) *IeFlightBillSettlementQueryResponse {
	s.StatusCode = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponse) SetBody(v *IeFlightBillSettlementQueryResponseBody) *IeFlightBillSettlementQueryResponse {
	s.Body = v
	return s
}

func (s *IeFlightBillSettlementQueryResponse) Validate() error {
	return dara.Validate(s)
}
