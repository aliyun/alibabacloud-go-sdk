// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFirewallDropStatisticsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeFirewallDropStatisticsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeFirewallDropStatisticsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeFirewallDropStatisticsResponseBody) *DescribeFirewallDropStatisticsResponse
	GetBody() *DescribeFirewallDropStatisticsResponseBody
}

type DescribeFirewallDropStatisticsResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeFirewallDropStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeFirewallDropStatisticsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeFirewallDropStatisticsResponse) GoString() string {
	return s.String()
}

func (s *DescribeFirewallDropStatisticsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeFirewallDropStatisticsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeFirewallDropStatisticsResponse) GetBody() *DescribeFirewallDropStatisticsResponseBody {
	return s.Body
}

func (s *DescribeFirewallDropStatisticsResponse) SetHeaders(v map[string]*string) *DescribeFirewallDropStatisticsResponse {
	s.Headers = v
	return s
}

func (s *DescribeFirewallDropStatisticsResponse) SetStatusCode(v int32) *DescribeFirewallDropStatisticsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeFirewallDropStatisticsResponse) SetBody(v *DescribeFirewallDropStatisticsResponseBody) *DescribeFirewallDropStatisticsResponse {
	s.Body = v
	return s
}

func (s *DescribeFirewallDropStatisticsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
