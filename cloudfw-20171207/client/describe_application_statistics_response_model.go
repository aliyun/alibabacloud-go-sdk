// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApplicationStatisticsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeApplicationStatisticsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeApplicationStatisticsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeApplicationStatisticsResponseBody) *DescribeApplicationStatisticsResponse
	GetBody() *DescribeApplicationStatisticsResponseBody
}

type DescribeApplicationStatisticsResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeApplicationStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeApplicationStatisticsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationStatisticsResponse) GoString() string {
	return s.String()
}

func (s *DescribeApplicationStatisticsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeApplicationStatisticsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeApplicationStatisticsResponse) GetBody() *DescribeApplicationStatisticsResponseBody {
	return s.Body
}

func (s *DescribeApplicationStatisticsResponse) SetHeaders(v map[string]*string) *DescribeApplicationStatisticsResponse {
	s.Headers = v
	return s
}

func (s *DescribeApplicationStatisticsResponse) SetStatusCode(v int32) *DescribeApplicationStatisticsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeApplicationStatisticsResponse) SetBody(v *DescribeApplicationStatisticsResponseBody) *DescribeApplicationStatisticsResponse {
	s.Body = v
	return s
}

func (s *DescribeApplicationStatisticsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
