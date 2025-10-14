// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMetricRuleTemplateListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeMetricRuleTemplateListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeMetricRuleTemplateListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeMetricRuleTemplateListResponseBody) *DescribeMetricRuleTemplateListResponse
	GetBody() *DescribeMetricRuleTemplateListResponseBody
}

type DescribeMetricRuleTemplateListResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeMetricRuleTemplateListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeMetricRuleTemplateListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeMetricRuleTemplateListResponse) GoString() string {
	return s.String()
}

func (s *DescribeMetricRuleTemplateListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeMetricRuleTemplateListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeMetricRuleTemplateListResponse) GetBody() *DescribeMetricRuleTemplateListResponseBody {
	return s.Body
}

func (s *DescribeMetricRuleTemplateListResponse) SetHeaders(v map[string]*string) *DescribeMetricRuleTemplateListResponse {
	s.Headers = v
	return s
}

func (s *DescribeMetricRuleTemplateListResponse) SetStatusCode(v int32) *DescribeMetricRuleTemplateListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeMetricRuleTemplateListResponse) SetBody(v *DescribeMetricRuleTemplateListResponseBody) *DescribeMetricRuleTemplateListResponse {
	s.Body = v
	return s
}

func (s *DescribeMetricRuleTemplateListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
