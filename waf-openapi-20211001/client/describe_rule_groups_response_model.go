// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRuleGroupsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeRuleGroupsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeRuleGroupsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeRuleGroupsResponseBody) *DescribeRuleGroupsResponse
	GetBody() *DescribeRuleGroupsResponseBody
}

type DescribeRuleGroupsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRuleGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRuleGroupsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeRuleGroupsResponse) GoString() string {
	return s.String()
}

func (s *DescribeRuleGroupsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeRuleGroupsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeRuleGroupsResponse) GetBody() *DescribeRuleGroupsResponseBody {
	return s.Body
}

func (s *DescribeRuleGroupsResponse) SetHeaders(v map[string]*string) *DescribeRuleGroupsResponse {
	s.Headers = v
	return s
}

func (s *DescribeRuleGroupsResponse) SetStatusCode(v int32) *DescribeRuleGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRuleGroupsResponse) SetBody(v *DescribeRuleGroupsResponseBody) *DescribeRuleGroupsResponse {
	s.Body = v
	return s
}

func (s *DescribeRuleGroupsResponse) Validate() error {
	return dara.Validate(s)
}
