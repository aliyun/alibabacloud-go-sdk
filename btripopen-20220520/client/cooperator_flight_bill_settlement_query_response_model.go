// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCooperatorFlightBillSettlementQueryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CooperatorFlightBillSettlementQueryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CooperatorFlightBillSettlementQueryResponse
	GetStatusCode() *int32
	SetBody(v *CooperatorFlightBillSettlementQueryResponseBody) *CooperatorFlightBillSettlementQueryResponse
	GetBody() *CooperatorFlightBillSettlementQueryResponseBody
}

type CooperatorFlightBillSettlementQueryResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CooperatorFlightBillSettlementQueryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CooperatorFlightBillSettlementQueryResponse) String() string {
	return dara.Prettify(s)
}

func (s CooperatorFlightBillSettlementQueryResponse) GoString() string {
	return s.String()
}

func (s *CooperatorFlightBillSettlementQueryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CooperatorFlightBillSettlementQueryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CooperatorFlightBillSettlementQueryResponse) GetBody() *CooperatorFlightBillSettlementQueryResponseBody {
	return s.Body
}

func (s *CooperatorFlightBillSettlementQueryResponse) SetHeaders(v map[string]*string) *CooperatorFlightBillSettlementQueryResponse {
	s.Headers = v
	return s
}

func (s *CooperatorFlightBillSettlementQueryResponse) SetStatusCode(v int32) *CooperatorFlightBillSettlementQueryResponse {
	s.StatusCode = &v
	return s
}

func (s *CooperatorFlightBillSettlementQueryResponse) SetBody(v *CooperatorFlightBillSettlementQueryResponseBody) *CooperatorFlightBillSettlementQueryResponse {
	s.Body = v
	return s
}

func (s *CooperatorFlightBillSettlementQueryResponse) Validate() error {
	return dara.Validate(s)
}
