// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodUserBillPredictionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVodUserBillPredictionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVodUserBillPredictionResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVodUserBillPredictionResponseBody) *DescribeVodUserBillPredictionResponse
	GetBody() *DescribeVodUserBillPredictionResponseBody
}

type DescribeVodUserBillPredictionResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVodUserBillPredictionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVodUserBillPredictionResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodUserBillPredictionResponse) GoString() string {
	return s.String()
}

func (s *DescribeVodUserBillPredictionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVodUserBillPredictionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVodUserBillPredictionResponse) GetBody() *DescribeVodUserBillPredictionResponseBody {
	return s.Body
}

func (s *DescribeVodUserBillPredictionResponse) SetHeaders(v map[string]*string) *DescribeVodUserBillPredictionResponse {
	s.Headers = v
	return s
}

func (s *DescribeVodUserBillPredictionResponse) SetStatusCode(v int32) *DescribeVodUserBillPredictionResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVodUserBillPredictionResponse) SetBody(v *DescribeVodUserBillPredictionResponseBody) *DescribeVodUserBillPredictionResponse {
	s.Body = v
	return s
}

func (s *DescribeVodUserBillPredictionResponse) Validate() error {
	return dara.Validate(s)
}
