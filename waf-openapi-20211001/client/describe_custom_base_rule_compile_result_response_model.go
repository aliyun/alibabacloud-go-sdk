// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCustomBaseRuleCompileResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCustomBaseRuleCompileResultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCustomBaseRuleCompileResultResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCustomBaseRuleCompileResultResponseBody) *DescribeCustomBaseRuleCompileResultResponse
	GetBody() *DescribeCustomBaseRuleCompileResultResponseBody
}

type DescribeCustomBaseRuleCompileResultResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCustomBaseRuleCompileResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCustomBaseRuleCompileResultResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCustomBaseRuleCompileResultResponse) GoString() string {
	return s.String()
}

func (s *DescribeCustomBaseRuleCompileResultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCustomBaseRuleCompileResultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCustomBaseRuleCompileResultResponse) GetBody() *DescribeCustomBaseRuleCompileResultResponseBody {
	return s.Body
}

func (s *DescribeCustomBaseRuleCompileResultResponse) SetHeaders(v map[string]*string) *DescribeCustomBaseRuleCompileResultResponse {
	s.Headers = v
	return s
}

func (s *DescribeCustomBaseRuleCompileResultResponse) SetStatusCode(v int32) *DescribeCustomBaseRuleCompileResultResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCustomBaseRuleCompileResultResponse) SetBody(v *DescribeCustomBaseRuleCompileResultResponseBody) *DescribeCustomBaseRuleCompileResultResponse {
	s.Body = v
	return s
}

func (s *DescribeCustomBaseRuleCompileResultResponse) Validate() error {
	return dara.Validate(s)
}
