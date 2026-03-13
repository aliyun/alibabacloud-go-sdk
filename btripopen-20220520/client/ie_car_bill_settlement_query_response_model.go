// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIeCarBillSettlementQueryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *IeCarBillSettlementQueryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *IeCarBillSettlementQueryResponse
	GetStatusCode() *int32
	SetBody(v *IeCarBillSettlementQueryResponseBody) *IeCarBillSettlementQueryResponse
	GetBody() *IeCarBillSettlementQueryResponseBody
}

type IeCarBillSettlementQueryResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *IeCarBillSettlementQueryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s IeCarBillSettlementQueryResponse) String() string {
	return dara.Prettify(s)
}

func (s IeCarBillSettlementQueryResponse) GoString() string {
	return s.String()
}

func (s *IeCarBillSettlementQueryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *IeCarBillSettlementQueryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *IeCarBillSettlementQueryResponse) GetBody() *IeCarBillSettlementQueryResponseBody {
	return s.Body
}

func (s *IeCarBillSettlementQueryResponse) SetHeaders(v map[string]*string) *IeCarBillSettlementQueryResponse {
	s.Headers = v
	return s
}

func (s *IeCarBillSettlementQueryResponse) SetStatusCode(v int32) *IeCarBillSettlementQueryResponse {
	s.StatusCode = &v
	return s
}

func (s *IeCarBillSettlementQueryResponse) SetBody(v *IeCarBillSettlementQueryResponseBody) *IeCarBillSettlementQueryResponse {
	s.Body = v
	return s
}

func (s *IeCarBillSettlementQueryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
