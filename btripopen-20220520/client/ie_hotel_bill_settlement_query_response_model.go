// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIeHotelBillSettlementQueryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *IeHotelBillSettlementQueryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *IeHotelBillSettlementQueryResponse
	GetStatusCode() *int32
	SetBody(v *IeHotelBillSettlementQueryResponseBody) *IeHotelBillSettlementQueryResponse
	GetBody() *IeHotelBillSettlementQueryResponseBody
}

type IeHotelBillSettlementQueryResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *IeHotelBillSettlementQueryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s IeHotelBillSettlementQueryResponse) String() string {
	return dara.Prettify(s)
}

func (s IeHotelBillSettlementQueryResponse) GoString() string {
	return s.String()
}

func (s *IeHotelBillSettlementQueryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *IeHotelBillSettlementQueryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *IeHotelBillSettlementQueryResponse) GetBody() *IeHotelBillSettlementQueryResponseBody {
	return s.Body
}

func (s *IeHotelBillSettlementQueryResponse) SetHeaders(v map[string]*string) *IeHotelBillSettlementQueryResponse {
	s.Headers = v
	return s
}

func (s *IeHotelBillSettlementQueryResponse) SetStatusCode(v int32) *IeHotelBillSettlementQueryResponse {
	s.StatusCode = &v
	return s
}

func (s *IeHotelBillSettlementQueryResponse) SetBody(v *IeHotelBillSettlementQueryResponseBody) *IeHotelBillSettlementQueryResponse {
	s.Body = v
	return s
}

func (s *IeHotelBillSettlementQueryResponse) Validate() error {
	return dara.Validate(s)
}
