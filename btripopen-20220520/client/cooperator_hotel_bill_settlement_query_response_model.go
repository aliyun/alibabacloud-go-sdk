// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCooperatorHotelBillSettlementQueryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CooperatorHotelBillSettlementQueryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CooperatorHotelBillSettlementQueryResponse
	GetStatusCode() *int32
	SetBody(v *CooperatorHotelBillSettlementQueryResponseBody) *CooperatorHotelBillSettlementQueryResponse
	GetBody() *CooperatorHotelBillSettlementQueryResponseBody
}

type CooperatorHotelBillSettlementQueryResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CooperatorHotelBillSettlementQueryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CooperatorHotelBillSettlementQueryResponse) String() string {
	return dara.Prettify(s)
}

func (s CooperatorHotelBillSettlementQueryResponse) GoString() string {
	return s.String()
}

func (s *CooperatorHotelBillSettlementQueryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CooperatorHotelBillSettlementQueryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CooperatorHotelBillSettlementQueryResponse) GetBody() *CooperatorHotelBillSettlementQueryResponseBody {
	return s.Body
}

func (s *CooperatorHotelBillSettlementQueryResponse) SetHeaders(v map[string]*string) *CooperatorHotelBillSettlementQueryResponse {
	s.Headers = v
	return s
}

func (s *CooperatorHotelBillSettlementQueryResponse) SetStatusCode(v int32) *CooperatorHotelBillSettlementQueryResponse {
	s.StatusCode = &v
	return s
}

func (s *CooperatorHotelBillSettlementQueryResponse) SetBody(v *CooperatorHotelBillSettlementQueryResponseBody) *CooperatorHotelBillSettlementQueryResponse {
	s.Body = v
	return s
}

func (s *CooperatorHotelBillSettlementQueryResponse) Validate() error {
	return dara.Validate(s)
}
