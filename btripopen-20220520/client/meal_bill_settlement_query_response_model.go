// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMealBillSettlementQueryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *MealBillSettlementQueryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *MealBillSettlementQueryResponse
	GetStatusCode() *int32
	SetBody(v *MealBillSettlementQueryResponseBody) *MealBillSettlementQueryResponse
	GetBody() *MealBillSettlementQueryResponseBody
}

type MealBillSettlementQueryResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MealBillSettlementQueryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s MealBillSettlementQueryResponse) String() string {
	return dara.Prettify(s)
}

func (s MealBillSettlementQueryResponse) GoString() string {
	return s.String()
}

func (s *MealBillSettlementQueryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *MealBillSettlementQueryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *MealBillSettlementQueryResponse) GetBody() *MealBillSettlementQueryResponseBody {
	return s.Body
}

func (s *MealBillSettlementQueryResponse) SetHeaders(v map[string]*string) *MealBillSettlementQueryResponse {
	s.Headers = v
	return s
}

func (s *MealBillSettlementQueryResponse) SetStatusCode(v int32) *MealBillSettlementQueryResponse {
	s.StatusCode = &v
	return s
}

func (s *MealBillSettlementQueryResponse) SetBody(v *MealBillSettlementQueryResponseBody) *MealBillSettlementQueryResponse {
	s.Body = v
	return s
}

func (s *MealBillSettlementQueryResponse) Validate() error {
	return dara.Validate(s)
}
