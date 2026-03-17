// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserFlowStatisticsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeUserFlowStatisticsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeUserFlowStatisticsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeUserFlowStatisticsResponseBody) *DescribeUserFlowStatisticsResponse
	GetBody() *DescribeUserFlowStatisticsResponseBody
}

type DescribeUserFlowStatisticsResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeUserFlowStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeUserFlowStatisticsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserFlowStatisticsResponse) GoString() string {
	return s.String()
}

func (s *DescribeUserFlowStatisticsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeUserFlowStatisticsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeUserFlowStatisticsResponse) GetBody() *DescribeUserFlowStatisticsResponseBody {
	return s.Body
}

func (s *DescribeUserFlowStatisticsResponse) SetHeaders(v map[string]*string) *DescribeUserFlowStatisticsResponse {
	s.Headers = v
	return s
}

func (s *DescribeUserFlowStatisticsResponse) SetStatusCode(v int32) *DescribeUserFlowStatisticsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeUserFlowStatisticsResponse) SetBody(v *DescribeUserFlowStatisticsResponseBody) *DescribeUserFlowStatisticsResponse {
	s.Body = v
	return s
}

func (s *DescribeUserFlowStatisticsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
