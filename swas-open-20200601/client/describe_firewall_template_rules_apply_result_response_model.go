// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFirewallTemplateRulesApplyResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeFirewallTemplateRulesApplyResultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeFirewallTemplateRulesApplyResultResponse
	GetStatusCode() *int32
	SetBody(v *DescribeFirewallTemplateRulesApplyResultResponseBody) *DescribeFirewallTemplateRulesApplyResultResponse
	GetBody() *DescribeFirewallTemplateRulesApplyResultResponseBody
}

type DescribeFirewallTemplateRulesApplyResultResponse struct {
	Headers    map[string]*string                                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeFirewallTemplateRulesApplyResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeFirewallTemplateRulesApplyResultResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeFirewallTemplateRulesApplyResultResponse) GoString() string {
	return s.String()
}

func (s *DescribeFirewallTemplateRulesApplyResultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeFirewallTemplateRulesApplyResultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeFirewallTemplateRulesApplyResultResponse) GetBody() *DescribeFirewallTemplateRulesApplyResultResponseBody {
	return s.Body
}

func (s *DescribeFirewallTemplateRulesApplyResultResponse) SetHeaders(v map[string]*string) *DescribeFirewallTemplateRulesApplyResultResponse {
	s.Headers = v
	return s
}

func (s *DescribeFirewallTemplateRulesApplyResultResponse) SetStatusCode(v int32) *DescribeFirewallTemplateRulesApplyResultResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeFirewallTemplateRulesApplyResultResponse) SetBody(v *DescribeFirewallTemplateRulesApplyResultResponseBody) *DescribeFirewallTemplateRulesApplyResultResponse {
	s.Body = v
	return s
}

func (s *DescribeFirewallTemplateRulesApplyResultResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
