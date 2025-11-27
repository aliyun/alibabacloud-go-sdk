// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRCMetricListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeRCMetricListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeRCMetricListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeRCMetricListResponseBody) *DescribeRCMetricListResponse
	GetBody() *DescribeRCMetricListResponseBody
}

type DescribeRCMetricListResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRCMetricListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRCMetricListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeRCMetricListResponse) GoString() string {
	return s.String()
}

func (s *DescribeRCMetricListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeRCMetricListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeRCMetricListResponse) GetBody() *DescribeRCMetricListResponseBody {
	return s.Body
}

func (s *DescribeRCMetricListResponse) SetHeaders(v map[string]*string) *DescribeRCMetricListResponse {
	s.Headers = v
	return s
}

func (s *DescribeRCMetricListResponse) SetStatusCode(v int32) *DescribeRCMetricListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRCMetricListResponse) SetBody(v *DescribeRCMetricListResponseBody) *DescribeRCMetricListResponse {
	s.Body = v
	return s
}

func (s *DescribeRCMetricListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
