// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVerifyFailStatisticsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVerifyFailStatisticsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVerifyFailStatisticsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVerifyFailStatisticsResponseBody) *DescribeVerifyFailStatisticsResponse
	GetBody() *DescribeVerifyFailStatisticsResponseBody
}

type DescribeVerifyFailStatisticsResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVerifyFailStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVerifyFailStatisticsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVerifyFailStatisticsResponse) GoString() string {
	return s.String()
}

func (s *DescribeVerifyFailStatisticsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVerifyFailStatisticsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVerifyFailStatisticsResponse) GetBody() *DescribeVerifyFailStatisticsResponseBody {
	return s.Body
}

func (s *DescribeVerifyFailStatisticsResponse) SetHeaders(v map[string]*string) *DescribeVerifyFailStatisticsResponse {
	s.Headers = v
	return s
}

func (s *DescribeVerifyFailStatisticsResponse) SetStatusCode(v int32) *DescribeVerifyFailStatisticsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVerifyFailStatisticsResponse) SetBody(v *DescribeVerifyFailStatisticsResponseBody) *DescribeVerifyFailStatisticsResponse {
	s.Body = v
	return s
}

func (s *DescribeVerifyFailStatisticsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
