// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePlaybookNumberMetricsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePlaybookNumberMetricsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePlaybookNumberMetricsResponse
	GetStatusCode() *int32
	SetBody(v *DescribePlaybookNumberMetricsResponseBody) *DescribePlaybookNumberMetricsResponse
	GetBody() *DescribePlaybookNumberMetricsResponseBody
}

type DescribePlaybookNumberMetricsResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePlaybookNumberMetricsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePlaybookNumberMetricsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePlaybookNumberMetricsResponse) GoString() string {
	return s.String()
}

func (s *DescribePlaybookNumberMetricsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePlaybookNumberMetricsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePlaybookNumberMetricsResponse) GetBody() *DescribePlaybookNumberMetricsResponseBody {
	return s.Body
}

func (s *DescribePlaybookNumberMetricsResponse) SetHeaders(v map[string]*string) *DescribePlaybookNumberMetricsResponse {
	s.Headers = v
	return s
}

func (s *DescribePlaybookNumberMetricsResponse) SetStatusCode(v int32) *DescribePlaybookNumberMetricsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePlaybookNumberMetricsResponse) SetBody(v *DescribePlaybookNumberMetricsResponseBody) *DescribePlaybookNumberMetricsResponse {
	s.Body = v
	return s
}

func (s *DescribePlaybookNumberMetricsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
