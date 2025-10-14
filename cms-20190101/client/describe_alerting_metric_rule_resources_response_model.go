// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAlertingMetricRuleResourcesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAlertingMetricRuleResourcesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAlertingMetricRuleResourcesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAlertingMetricRuleResourcesResponseBody) *DescribeAlertingMetricRuleResourcesResponse
	GetBody() *DescribeAlertingMetricRuleResourcesResponseBody
}

type DescribeAlertingMetricRuleResourcesResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAlertingMetricRuleResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAlertingMetricRuleResourcesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAlertingMetricRuleResourcesResponse) GoString() string {
	return s.String()
}

func (s *DescribeAlertingMetricRuleResourcesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAlertingMetricRuleResourcesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAlertingMetricRuleResourcesResponse) GetBody() *DescribeAlertingMetricRuleResourcesResponseBody {
	return s.Body
}

func (s *DescribeAlertingMetricRuleResourcesResponse) SetHeaders(v map[string]*string) *DescribeAlertingMetricRuleResourcesResponse {
	s.Headers = v
	return s
}

func (s *DescribeAlertingMetricRuleResourcesResponse) SetStatusCode(v int32) *DescribeAlertingMetricRuleResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAlertingMetricRuleResourcesResponse) SetBody(v *DescribeAlertingMetricRuleResourcesResponseBody) *DescribeAlertingMetricRuleResourcesResponse {
	s.Body = v
	return s
}

func (s *DescribeAlertingMetricRuleResourcesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
