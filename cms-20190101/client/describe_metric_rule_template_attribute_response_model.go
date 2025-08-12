// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMetricRuleTemplateAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeMetricRuleTemplateAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeMetricRuleTemplateAttributeResponse
	GetStatusCode() *int32
	SetBody(v *DescribeMetricRuleTemplateAttributeResponseBody) *DescribeMetricRuleTemplateAttributeResponse
	GetBody() *DescribeMetricRuleTemplateAttributeResponseBody
}

type DescribeMetricRuleTemplateAttributeResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeMetricRuleTemplateAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeMetricRuleTemplateAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeMetricRuleTemplateAttributeResponse) GoString() string {
	return s.String()
}

func (s *DescribeMetricRuleTemplateAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeMetricRuleTemplateAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeMetricRuleTemplateAttributeResponse) GetBody() *DescribeMetricRuleTemplateAttributeResponseBody {
	return s.Body
}

func (s *DescribeMetricRuleTemplateAttributeResponse) SetHeaders(v map[string]*string) *DescribeMetricRuleTemplateAttributeResponse {
	s.Headers = v
	return s
}

func (s *DescribeMetricRuleTemplateAttributeResponse) SetStatusCode(v int32) *DescribeMetricRuleTemplateAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeMetricRuleTemplateAttributeResponse) SetBody(v *DescribeMetricRuleTemplateAttributeResponseBody) *DescribeMetricRuleTemplateAttributeResponse {
	s.Body = v
	return s
}

func (s *DescribeMetricRuleTemplateAttributeResponse) Validate() error {
	return dara.Validate(s)
}
