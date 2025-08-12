// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCustomMetricListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCustomMetricListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCustomMetricListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCustomMetricListResponseBody) *DescribeCustomMetricListResponse
	GetBody() *DescribeCustomMetricListResponseBody
}

type DescribeCustomMetricListResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCustomMetricListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCustomMetricListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCustomMetricListResponse) GoString() string {
	return s.String()
}

func (s *DescribeCustomMetricListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCustomMetricListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCustomMetricListResponse) GetBody() *DescribeCustomMetricListResponseBody {
	return s.Body
}

func (s *DescribeCustomMetricListResponse) SetHeaders(v map[string]*string) *DescribeCustomMetricListResponse {
	s.Headers = v
	return s
}

func (s *DescribeCustomMetricListResponse) SetStatusCode(v int32) *DescribeCustomMetricListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCustomMetricListResponse) SetBody(v *DescribeCustomMetricListResponseBody) *DescribeCustomMetricListResponse {
	s.Body = v
	return s
}

func (s *DescribeCustomMetricListResponse) Validate() error {
	return dara.Validate(s)
}
