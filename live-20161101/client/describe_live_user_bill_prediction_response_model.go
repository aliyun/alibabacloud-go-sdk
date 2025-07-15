// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveUserBillPredictionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeLiveUserBillPredictionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeLiveUserBillPredictionResponse
	GetStatusCode() *int32
	SetBody(v *DescribeLiveUserBillPredictionResponseBody) *DescribeLiveUserBillPredictionResponse
	GetBody() *DescribeLiveUserBillPredictionResponseBody
}

type DescribeLiveUserBillPredictionResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLiveUserBillPredictionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLiveUserBillPredictionResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveUserBillPredictionResponse) GoString() string {
	return s.String()
}

func (s *DescribeLiveUserBillPredictionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeLiveUserBillPredictionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeLiveUserBillPredictionResponse) GetBody() *DescribeLiveUserBillPredictionResponseBody {
	return s.Body
}

func (s *DescribeLiveUserBillPredictionResponse) SetHeaders(v map[string]*string) *DescribeLiveUserBillPredictionResponse {
	s.Headers = v
	return s
}

func (s *DescribeLiveUserBillPredictionResponse) SetStatusCode(v int32) *DescribeLiveUserBillPredictionResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLiveUserBillPredictionResponse) SetBody(v *DescribeLiveUserBillPredictionResponseBody) *DescribeLiveUserBillPredictionResponse {
	s.Body = v
	return s
}

func (s *DescribeLiveUserBillPredictionResponse) Validate() error {
	return dara.Validate(s)
}
