// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMetricRuleTargetsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeMetricRuleTargetsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeMetricRuleTargetsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeMetricRuleTargetsResponseBody) *DescribeMetricRuleTargetsResponse
	GetBody() *DescribeMetricRuleTargetsResponseBody
}

type DescribeMetricRuleTargetsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeMetricRuleTargetsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeMetricRuleTargetsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeMetricRuleTargetsResponse) GoString() string {
	return s.String()
}

func (s *DescribeMetricRuleTargetsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeMetricRuleTargetsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeMetricRuleTargetsResponse) GetBody() *DescribeMetricRuleTargetsResponseBody {
	return s.Body
}

func (s *DescribeMetricRuleTargetsResponse) SetHeaders(v map[string]*string) *DescribeMetricRuleTargetsResponse {
	s.Headers = v
	return s
}

func (s *DescribeMetricRuleTargetsResponse) SetStatusCode(v int32) *DescribeMetricRuleTargetsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeMetricRuleTargetsResponse) SetBody(v *DescribeMetricRuleTargetsResponseBody) *DescribeMetricRuleTargetsResponse {
	s.Body = v
	return s
}

func (s *DescribeMetricRuleTargetsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
