// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMetricRuleCountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeMetricRuleCountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeMetricRuleCountResponse
	GetStatusCode() *int32
	SetBody(v *DescribeMetricRuleCountResponseBody) *DescribeMetricRuleCountResponse
	GetBody() *DescribeMetricRuleCountResponseBody
}

type DescribeMetricRuleCountResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeMetricRuleCountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeMetricRuleCountResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeMetricRuleCountResponse) GoString() string {
	return s.String()
}

func (s *DescribeMetricRuleCountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeMetricRuleCountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeMetricRuleCountResponse) GetBody() *DescribeMetricRuleCountResponseBody {
	return s.Body
}

func (s *DescribeMetricRuleCountResponse) SetHeaders(v map[string]*string) *DescribeMetricRuleCountResponse {
	s.Headers = v
	return s
}

func (s *DescribeMetricRuleCountResponse) SetStatusCode(v int32) *DescribeMetricRuleCountResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeMetricRuleCountResponse) SetBody(v *DescribeMetricRuleCountResponseBody) *DescribeMetricRuleCountResponse {
	s.Body = v
	return s
}

func (s *DescribeMetricRuleCountResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
