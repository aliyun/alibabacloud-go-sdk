// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTrainBillSettlementQueryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *TrainBillSettlementQueryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *TrainBillSettlementQueryResponse
	GetStatusCode() *int32
	SetBody(v *TrainBillSettlementQueryResponseBody) *TrainBillSettlementQueryResponse
	GetBody() *TrainBillSettlementQueryResponseBody
}

type TrainBillSettlementQueryResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TrainBillSettlementQueryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TrainBillSettlementQueryResponse) String() string {
	return dara.Prettify(s)
}

func (s TrainBillSettlementQueryResponse) GoString() string {
	return s.String()
}

func (s *TrainBillSettlementQueryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *TrainBillSettlementQueryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *TrainBillSettlementQueryResponse) GetBody() *TrainBillSettlementQueryResponseBody {
	return s.Body
}

func (s *TrainBillSettlementQueryResponse) SetHeaders(v map[string]*string) *TrainBillSettlementQueryResponse {
	s.Headers = v
	return s
}

func (s *TrainBillSettlementQueryResponse) SetStatusCode(v int32) *TrainBillSettlementQueryResponse {
	s.StatusCode = &v
	return s
}

func (s *TrainBillSettlementQueryResponse) SetBody(v *TrainBillSettlementQueryResponseBody) *TrainBillSettlementQueryResponse {
	s.Body = v
	return s
}

func (s *TrainBillSettlementQueryResponse) Validate() error {
	return dara.Validate(s)
}
