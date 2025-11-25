// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePrivateDnsStatisticsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePrivateDnsStatisticsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePrivateDnsStatisticsResponse
	GetStatusCode() *int32
	SetBody(v *DescribePrivateDnsStatisticsResponseBody) *DescribePrivateDnsStatisticsResponse
	GetBody() *DescribePrivateDnsStatisticsResponseBody
}

type DescribePrivateDnsStatisticsResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePrivateDnsStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePrivateDnsStatisticsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePrivateDnsStatisticsResponse) GoString() string {
	return s.String()
}

func (s *DescribePrivateDnsStatisticsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePrivateDnsStatisticsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePrivateDnsStatisticsResponse) GetBody() *DescribePrivateDnsStatisticsResponseBody {
	return s.Body
}

func (s *DescribePrivateDnsStatisticsResponse) SetHeaders(v map[string]*string) *DescribePrivateDnsStatisticsResponse {
	s.Headers = v
	return s
}

func (s *DescribePrivateDnsStatisticsResponse) SetStatusCode(v int32) *DescribePrivateDnsStatisticsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePrivateDnsStatisticsResponse) SetBody(v *DescribePrivateDnsStatisticsResponseBody) *DescribePrivateDnsStatisticsResponse {
	s.Body = v
	return s
}

func (s *DescribePrivateDnsStatisticsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
