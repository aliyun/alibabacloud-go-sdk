// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceCreateAndDeleteStatisticsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeInstanceCreateAndDeleteStatisticsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeInstanceCreateAndDeleteStatisticsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeInstanceCreateAndDeleteStatisticsResponseBody) *DescribeInstanceCreateAndDeleteStatisticsResponse
	GetBody() *DescribeInstanceCreateAndDeleteStatisticsResponseBody
}

type DescribeInstanceCreateAndDeleteStatisticsResponse struct {
	Headers    map[string]*string                                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeInstanceCreateAndDeleteStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeInstanceCreateAndDeleteStatisticsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceCreateAndDeleteStatisticsResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstanceCreateAndDeleteStatisticsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeInstanceCreateAndDeleteStatisticsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeInstanceCreateAndDeleteStatisticsResponse) GetBody() *DescribeInstanceCreateAndDeleteStatisticsResponseBody {
	return s.Body
}

func (s *DescribeInstanceCreateAndDeleteStatisticsResponse) SetHeaders(v map[string]*string) *DescribeInstanceCreateAndDeleteStatisticsResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstanceCreateAndDeleteStatisticsResponse) SetStatusCode(v int32) *DescribeInstanceCreateAndDeleteStatisticsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInstanceCreateAndDeleteStatisticsResponse) SetBody(v *DescribeInstanceCreateAndDeleteStatisticsResponseBody) *DescribeInstanceCreateAndDeleteStatisticsResponse {
	s.Body = v
	return s
}

func (s *DescribeInstanceCreateAndDeleteStatisticsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
