// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFirewallTemplateApplyResultsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeFirewallTemplateApplyResultsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeFirewallTemplateApplyResultsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeFirewallTemplateApplyResultsResponseBody) *DescribeFirewallTemplateApplyResultsResponse
	GetBody() *DescribeFirewallTemplateApplyResultsResponseBody
}

type DescribeFirewallTemplateApplyResultsResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeFirewallTemplateApplyResultsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeFirewallTemplateApplyResultsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeFirewallTemplateApplyResultsResponse) GoString() string {
	return s.String()
}

func (s *DescribeFirewallTemplateApplyResultsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeFirewallTemplateApplyResultsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeFirewallTemplateApplyResultsResponse) GetBody() *DescribeFirewallTemplateApplyResultsResponseBody {
	return s.Body
}

func (s *DescribeFirewallTemplateApplyResultsResponse) SetHeaders(v map[string]*string) *DescribeFirewallTemplateApplyResultsResponse {
	s.Headers = v
	return s
}

func (s *DescribeFirewallTemplateApplyResultsResponse) SetStatusCode(v int32) *DescribeFirewallTemplateApplyResultsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeFirewallTemplateApplyResultsResponse) SetBody(v *DescribeFirewallTemplateApplyResultsResponseBody) *DescribeFirewallTemplateApplyResultsResponse {
	s.Body = v
	return s
}

func (s *DescribeFirewallTemplateApplyResultsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
