// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMetricListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeMetricListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeMetricListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeMetricListResponseBody) *DescribeMetricListResponse
	GetBody() *DescribeMetricListResponseBody
}

type DescribeMetricListResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeMetricListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeMetricListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeMetricListResponse) GoString() string {
	return s.String()
}

func (s *DescribeMetricListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeMetricListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeMetricListResponse) GetBody() *DescribeMetricListResponseBody {
	return s.Body
}

func (s *DescribeMetricListResponse) SetHeaders(v map[string]*string) *DescribeMetricListResponse {
	s.Headers = v
	return s
}

func (s *DescribeMetricListResponse) SetStatusCode(v int32) *DescribeMetricListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeMetricListResponse) SetBody(v *DescribeMetricListResponseBody) *DescribeMetricListResponse {
	s.Body = v
	return s
}

func (s *DescribeMetricListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
