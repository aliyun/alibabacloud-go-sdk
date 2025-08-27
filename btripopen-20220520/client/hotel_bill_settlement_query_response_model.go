// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHotelBillSettlementQueryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *HotelBillSettlementQueryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *HotelBillSettlementQueryResponse
	GetStatusCode() *int32
	SetBody(v *HotelBillSettlementQueryResponseBody) *HotelBillSettlementQueryResponse
	GetBody() *HotelBillSettlementQueryResponseBody
}

type HotelBillSettlementQueryResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *HotelBillSettlementQueryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s HotelBillSettlementQueryResponse) String() string {
	return dara.Prettify(s)
}

func (s HotelBillSettlementQueryResponse) GoString() string {
	return s.String()
}

func (s *HotelBillSettlementQueryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *HotelBillSettlementQueryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *HotelBillSettlementQueryResponse) GetBody() *HotelBillSettlementQueryResponseBody {
	return s.Body
}

func (s *HotelBillSettlementQueryResponse) SetHeaders(v map[string]*string) *HotelBillSettlementQueryResponse {
	s.Headers = v
	return s
}

func (s *HotelBillSettlementQueryResponse) SetStatusCode(v int32) *HotelBillSettlementQueryResponse {
	s.StatusCode = &v
	return s
}

func (s *HotelBillSettlementQueryResponse) SetBody(v *HotelBillSettlementQueryResponseBody) *HotelBillSettlementQueryResponse {
	s.Body = v
	return s
}

func (s *HotelBillSettlementQueryResponse) Validate() error {
	return dara.Validate(s)
}
