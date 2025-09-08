// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCustomizeRuleTestResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCustomizeRuleTestResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCustomizeRuleTestResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCustomizeRuleTestResponseBody) *DescribeCustomizeRuleTestResponse
	GetBody() *DescribeCustomizeRuleTestResponseBody
}

type DescribeCustomizeRuleTestResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCustomizeRuleTestResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCustomizeRuleTestResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCustomizeRuleTestResponse) GoString() string {
	return s.String()
}

func (s *DescribeCustomizeRuleTestResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCustomizeRuleTestResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCustomizeRuleTestResponse) GetBody() *DescribeCustomizeRuleTestResponseBody {
	return s.Body
}

func (s *DescribeCustomizeRuleTestResponse) SetHeaders(v map[string]*string) *DescribeCustomizeRuleTestResponse {
	s.Headers = v
	return s
}

func (s *DescribeCustomizeRuleTestResponse) SetStatusCode(v int32) *DescribeCustomizeRuleTestResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCustomizeRuleTestResponse) SetBody(v *DescribeCustomizeRuleTestResponseBody) *DescribeCustomizeRuleTestResponse {
	s.Body = v
	return s
}

func (s *DescribeCustomizeRuleTestResponse) Validate() error {
	return dara.Validate(s)
}
