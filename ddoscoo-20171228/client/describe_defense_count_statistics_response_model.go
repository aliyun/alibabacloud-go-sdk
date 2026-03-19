// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDefenseCountStatisticsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDefenseCountStatisticsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDefenseCountStatisticsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDefenseCountStatisticsResponseBody) *DescribeDefenseCountStatisticsResponse
	GetBody() *DescribeDefenseCountStatisticsResponseBody
}

type DescribeDefenseCountStatisticsResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDefenseCountStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDefenseCountStatisticsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDefenseCountStatisticsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDefenseCountStatisticsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDefenseCountStatisticsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDefenseCountStatisticsResponse) GetBody() *DescribeDefenseCountStatisticsResponseBody {
	return s.Body
}

func (s *DescribeDefenseCountStatisticsResponse) SetHeaders(v map[string]*string) *DescribeDefenseCountStatisticsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDefenseCountStatisticsResponse) SetStatusCode(v int32) *DescribeDefenseCountStatisticsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDefenseCountStatisticsResponse) SetBody(v *DescribeDefenseCountStatisticsResponseBody) *DescribeDefenseCountStatisticsResponse {
	s.Body = v
	return s
}

func (s *DescribeDefenseCountStatisticsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
