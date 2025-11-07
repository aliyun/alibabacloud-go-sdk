// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVerifyStatisticsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVerifyStatisticsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVerifyStatisticsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVerifyStatisticsResponseBody) *DescribeVerifyStatisticsResponse
	GetBody() *DescribeVerifyStatisticsResponseBody
}

type DescribeVerifyStatisticsResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVerifyStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVerifyStatisticsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVerifyStatisticsResponse) GoString() string {
	return s.String()
}

func (s *DescribeVerifyStatisticsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVerifyStatisticsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVerifyStatisticsResponse) GetBody() *DescribeVerifyStatisticsResponseBody {
	return s.Body
}

func (s *DescribeVerifyStatisticsResponse) SetHeaders(v map[string]*string) *DescribeVerifyStatisticsResponse {
	s.Headers = v
	return s
}

func (s *DescribeVerifyStatisticsResponse) SetStatusCode(v int32) *DescribeVerifyStatisticsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVerifyStatisticsResponse) SetBody(v *DescribeVerifyStatisticsResponseBody) *DescribeVerifyStatisticsResponse {
	s.Body = v
	return s
}

func (s *DescribeVerifyStatisticsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
