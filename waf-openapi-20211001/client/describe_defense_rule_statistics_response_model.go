// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDefenseRuleStatisticsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDefenseRuleStatisticsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDefenseRuleStatisticsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDefenseRuleStatisticsResponseBody) *DescribeDefenseRuleStatisticsResponse
	GetBody() *DescribeDefenseRuleStatisticsResponseBody
}

type DescribeDefenseRuleStatisticsResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDefenseRuleStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDefenseRuleStatisticsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDefenseRuleStatisticsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDefenseRuleStatisticsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDefenseRuleStatisticsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDefenseRuleStatisticsResponse) GetBody() *DescribeDefenseRuleStatisticsResponseBody {
	return s.Body
}

func (s *DescribeDefenseRuleStatisticsResponse) SetHeaders(v map[string]*string) *DescribeDefenseRuleStatisticsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDefenseRuleStatisticsResponse) SetStatusCode(v int32) *DescribeDefenseRuleStatisticsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDefenseRuleStatisticsResponse) SetBody(v *DescribeDefenseRuleStatisticsResponseBody) *DescribeDefenseRuleStatisticsResponse {
	s.Body = v
	return s
}

func (s *DescribeDefenseRuleStatisticsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
