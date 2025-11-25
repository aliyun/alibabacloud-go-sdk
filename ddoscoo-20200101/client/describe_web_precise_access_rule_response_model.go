// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeWebPreciseAccessRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeWebPreciseAccessRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeWebPreciseAccessRuleResponse
	GetStatusCode() *int32
	SetBody(v *DescribeWebPreciseAccessRuleResponseBody) *DescribeWebPreciseAccessRuleResponse
	GetBody() *DescribeWebPreciseAccessRuleResponseBody
}

type DescribeWebPreciseAccessRuleResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeWebPreciseAccessRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeWebPreciseAccessRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeWebPreciseAccessRuleResponse) GoString() string {
	return s.String()
}

func (s *DescribeWebPreciseAccessRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeWebPreciseAccessRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeWebPreciseAccessRuleResponse) GetBody() *DescribeWebPreciseAccessRuleResponseBody {
	return s.Body
}

func (s *DescribeWebPreciseAccessRuleResponse) SetHeaders(v map[string]*string) *DescribeWebPreciseAccessRuleResponse {
	s.Headers = v
	return s
}

func (s *DescribeWebPreciseAccessRuleResponse) SetStatusCode(v int32) *DescribeWebPreciseAccessRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeWebPreciseAccessRuleResponse) SetBody(v *DescribeWebPreciseAccessRuleResponseBody) *DescribeWebPreciseAccessRuleResponse {
	s.Body = v
	return s
}

func (s *DescribeWebPreciseAccessRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
