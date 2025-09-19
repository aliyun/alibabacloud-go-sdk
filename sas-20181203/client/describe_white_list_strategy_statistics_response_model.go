// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeWhiteListStrategyStatisticsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeWhiteListStrategyStatisticsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeWhiteListStrategyStatisticsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeWhiteListStrategyStatisticsResponseBody) *DescribeWhiteListStrategyStatisticsResponse
	GetBody() *DescribeWhiteListStrategyStatisticsResponseBody
}

type DescribeWhiteListStrategyStatisticsResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeWhiteListStrategyStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeWhiteListStrategyStatisticsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeWhiteListStrategyStatisticsResponse) GoString() string {
	return s.String()
}

func (s *DescribeWhiteListStrategyStatisticsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeWhiteListStrategyStatisticsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeWhiteListStrategyStatisticsResponse) GetBody() *DescribeWhiteListStrategyStatisticsResponseBody {
	return s.Body
}

func (s *DescribeWhiteListStrategyStatisticsResponse) SetHeaders(v map[string]*string) *DescribeWhiteListStrategyStatisticsResponse {
	s.Headers = v
	return s
}

func (s *DescribeWhiteListStrategyStatisticsResponse) SetStatusCode(v int32) *DescribeWhiteListStrategyStatisticsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeWhiteListStrategyStatisticsResponse) SetBody(v *DescribeWhiteListStrategyStatisticsResponseBody) *DescribeWhiteListStrategyStatisticsResponse {
	s.Body = v
	return s
}

func (s *DescribeWhiteListStrategyStatisticsResponse) Validate() error {
	return dara.Validate(s)
}
