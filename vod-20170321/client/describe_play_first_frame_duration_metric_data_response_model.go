// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePlayFirstFrameDurationMetricDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePlayFirstFrameDurationMetricDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePlayFirstFrameDurationMetricDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribePlayFirstFrameDurationMetricDataResponseBody) *DescribePlayFirstFrameDurationMetricDataResponse
	GetBody() *DescribePlayFirstFrameDurationMetricDataResponseBody
}

type DescribePlayFirstFrameDurationMetricDataResponse struct {
	Headers    map[string]*string                                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePlayFirstFrameDurationMetricDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePlayFirstFrameDurationMetricDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePlayFirstFrameDurationMetricDataResponse) GoString() string {
	return s.String()
}

func (s *DescribePlayFirstFrameDurationMetricDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePlayFirstFrameDurationMetricDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePlayFirstFrameDurationMetricDataResponse) GetBody() *DescribePlayFirstFrameDurationMetricDataResponseBody {
	return s.Body
}

func (s *DescribePlayFirstFrameDurationMetricDataResponse) SetHeaders(v map[string]*string) *DescribePlayFirstFrameDurationMetricDataResponse {
	s.Headers = v
	return s
}

func (s *DescribePlayFirstFrameDurationMetricDataResponse) SetStatusCode(v int32) *DescribePlayFirstFrameDurationMetricDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePlayFirstFrameDurationMetricDataResponse) SetBody(v *DescribePlayFirstFrameDurationMetricDataResponseBody) *DescribePlayFirstFrameDurationMetricDataResponse {
	s.Body = v
	return s
}

func (s *DescribePlayFirstFrameDurationMetricDataResponse) Validate() error {
	return dara.Validate(s)
}
