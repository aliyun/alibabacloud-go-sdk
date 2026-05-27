// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSumDailyBillsByItemResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SumDailyBillsByItemResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SumDailyBillsByItemResponse
	GetStatusCode() *int32
	SetBody(v *SumDailyBillsByItemResponseBody) *SumDailyBillsByItemResponse
	GetBody() *SumDailyBillsByItemResponseBody
}

type SumDailyBillsByItemResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SumDailyBillsByItemResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SumDailyBillsByItemResponse) String() string {
	return dara.Prettify(s)
}

func (s SumDailyBillsByItemResponse) GoString() string {
	return s.String()
}

func (s *SumDailyBillsByItemResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SumDailyBillsByItemResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SumDailyBillsByItemResponse) GetBody() *SumDailyBillsByItemResponseBody {
	return s.Body
}

func (s *SumDailyBillsByItemResponse) SetHeaders(v map[string]*string) *SumDailyBillsByItemResponse {
	s.Headers = v
	return s
}

func (s *SumDailyBillsByItemResponse) SetStatusCode(v int32) *SumDailyBillsByItemResponse {
	s.StatusCode = &v
	return s
}

func (s *SumDailyBillsByItemResponse) SetBody(v *SumDailyBillsByItemResponseBody) *SumDailyBillsByItemResponse {
	s.Body = v
	return s
}

func (s *SumDailyBillsByItemResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
