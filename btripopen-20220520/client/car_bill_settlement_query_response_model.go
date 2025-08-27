// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCarBillSettlementQueryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CarBillSettlementQueryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CarBillSettlementQueryResponse
	GetStatusCode() *int32
	SetBody(v *CarBillSettlementQueryResponseBody) *CarBillSettlementQueryResponse
	GetBody() *CarBillSettlementQueryResponseBody
}

type CarBillSettlementQueryResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CarBillSettlementQueryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CarBillSettlementQueryResponse) String() string {
	return dara.Prettify(s)
}

func (s CarBillSettlementQueryResponse) GoString() string {
	return s.String()
}

func (s *CarBillSettlementQueryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CarBillSettlementQueryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CarBillSettlementQueryResponse) GetBody() *CarBillSettlementQueryResponseBody {
	return s.Body
}

func (s *CarBillSettlementQueryResponse) SetHeaders(v map[string]*string) *CarBillSettlementQueryResponse {
	s.Headers = v
	return s
}

func (s *CarBillSettlementQueryResponse) SetStatusCode(v int32) *CarBillSettlementQueryResponse {
	s.StatusCode = &v
	return s
}

func (s *CarBillSettlementQueryResponse) SetBody(v *CarBillSettlementQueryResponseBody) *CarBillSettlementQueryResponse {
	s.Body = v
	return s
}

func (s *CarBillSettlementQueryResponse) Validate() error {
	return dara.Validate(s)
}
