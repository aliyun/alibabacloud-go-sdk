// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePlaybookMetricsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePlaybookMetricsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePlaybookMetricsResponse
	GetStatusCode() *int32
	SetBody(v *DescribePlaybookMetricsResponseBody) *DescribePlaybookMetricsResponse
	GetBody() *DescribePlaybookMetricsResponseBody
}

type DescribePlaybookMetricsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePlaybookMetricsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePlaybookMetricsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePlaybookMetricsResponse) GoString() string {
	return s.String()
}

func (s *DescribePlaybookMetricsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePlaybookMetricsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePlaybookMetricsResponse) GetBody() *DescribePlaybookMetricsResponseBody {
	return s.Body
}

func (s *DescribePlaybookMetricsResponse) SetHeaders(v map[string]*string) *DescribePlaybookMetricsResponse {
	s.Headers = v
	return s
}

func (s *DescribePlaybookMetricsResponse) SetStatusCode(v int32) *DescribePlaybookMetricsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePlaybookMetricsResponse) SetBody(v *DescribePlaybookMetricsResponseBody) *DescribePlaybookMetricsResponse {
	s.Body = v
	return s
}

func (s *DescribePlaybookMetricsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
