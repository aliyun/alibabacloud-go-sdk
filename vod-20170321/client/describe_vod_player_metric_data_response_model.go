// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodPlayerMetricDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVodPlayerMetricDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVodPlayerMetricDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVodPlayerMetricDataResponseBody) *DescribeVodPlayerMetricDataResponse
	GetBody() *DescribeVodPlayerMetricDataResponseBody
}

type DescribeVodPlayerMetricDataResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVodPlayerMetricDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVodPlayerMetricDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodPlayerMetricDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeVodPlayerMetricDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVodPlayerMetricDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVodPlayerMetricDataResponse) GetBody() *DescribeVodPlayerMetricDataResponseBody {
	return s.Body
}

func (s *DescribeVodPlayerMetricDataResponse) SetHeaders(v map[string]*string) *DescribeVodPlayerMetricDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeVodPlayerMetricDataResponse) SetStatusCode(v int32) *DescribeVodPlayerMetricDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVodPlayerMetricDataResponse) SetBody(v *DescribeVodPlayerMetricDataResponseBody) *DescribeVodPlayerMetricDataResponse {
	s.Body = v
	return s
}

func (s *DescribeVodPlayerMetricDataResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
