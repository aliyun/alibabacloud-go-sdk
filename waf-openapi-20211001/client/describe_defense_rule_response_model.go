// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDefenseRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDefenseRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDefenseRuleResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDefenseRuleResponseBody) *DescribeDefenseRuleResponse
	GetBody() *DescribeDefenseRuleResponseBody
}

type DescribeDefenseRuleResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDefenseRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDefenseRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDefenseRuleResponse) GoString() string {
	return s.String()
}

func (s *DescribeDefenseRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDefenseRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDefenseRuleResponse) GetBody() *DescribeDefenseRuleResponseBody {
	return s.Body
}

func (s *DescribeDefenseRuleResponse) SetHeaders(v map[string]*string) *DescribeDefenseRuleResponse {
	s.Headers = v
	return s
}

func (s *DescribeDefenseRuleResponse) SetStatusCode(v int32) *DescribeDefenseRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDefenseRuleResponse) SetBody(v *DescribeDefenseRuleResponseBody) *DescribeDefenseRuleResponse {
	s.Body = v
	return s
}

func (s *DescribeDefenseRuleResponse) Validate() error {
	return dara.Validate(s)
}
