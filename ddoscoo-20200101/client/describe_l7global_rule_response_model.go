// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeL7GlobalRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeL7GlobalRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeL7GlobalRuleResponse
	GetStatusCode() *int32
	SetBody(v *DescribeL7GlobalRuleResponseBody) *DescribeL7GlobalRuleResponse
	GetBody() *DescribeL7GlobalRuleResponseBody
}

type DescribeL7GlobalRuleResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeL7GlobalRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeL7GlobalRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeL7GlobalRuleResponse) GoString() string {
	return s.String()
}

func (s *DescribeL7GlobalRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeL7GlobalRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeL7GlobalRuleResponse) GetBody() *DescribeL7GlobalRuleResponseBody {
	return s.Body
}

func (s *DescribeL7GlobalRuleResponse) SetHeaders(v map[string]*string) *DescribeL7GlobalRuleResponse {
	s.Headers = v
	return s
}

func (s *DescribeL7GlobalRuleResponse) SetStatusCode(v int32) *DescribeL7GlobalRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeL7GlobalRuleResponse) SetBody(v *DescribeL7GlobalRuleResponseBody) *DescribeL7GlobalRuleResponse {
	s.Body = v
	return s
}

func (s *DescribeL7GlobalRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
