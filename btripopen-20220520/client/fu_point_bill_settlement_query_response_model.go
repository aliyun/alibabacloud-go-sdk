// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFuPointBillSettlementQueryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *FuPointBillSettlementQueryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *FuPointBillSettlementQueryResponse
	GetStatusCode() *int32
	SetBody(v *FuPointBillSettlementQueryResponseBody) *FuPointBillSettlementQueryResponse
	GetBody() *FuPointBillSettlementQueryResponseBody
}

type FuPointBillSettlementQueryResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *FuPointBillSettlementQueryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s FuPointBillSettlementQueryResponse) String() string {
	return dara.Prettify(s)
}

func (s FuPointBillSettlementQueryResponse) GoString() string {
	return s.String()
}

func (s *FuPointBillSettlementQueryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *FuPointBillSettlementQueryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *FuPointBillSettlementQueryResponse) GetBody() *FuPointBillSettlementQueryResponseBody {
	return s.Body
}

func (s *FuPointBillSettlementQueryResponse) SetHeaders(v map[string]*string) *FuPointBillSettlementQueryResponse {
	s.Headers = v
	return s
}

func (s *FuPointBillSettlementQueryResponse) SetStatusCode(v int32) *FuPointBillSettlementQueryResponse {
	s.StatusCode = &v
	return s
}

func (s *FuPointBillSettlementQueryResponse) SetBody(v *FuPointBillSettlementQueryResponseBody) *FuPointBillSettlementQueryResponse {
	s.Body = v
	return s
}

func (s *FuPointBillSettlementQueryResponse) Validate() error {
	return dara.Validate(s)
}
