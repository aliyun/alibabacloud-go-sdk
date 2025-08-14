// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHasRuleNameByEventCodeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeHasRuleNameByEventCodeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeHasRuleNameByEventCodeResponse
	GetStatusCode() *int32
	SetBody(v *DescribeHasRuleNameByEventCodeResponseBody) *DescribeHasRuleNameByEventCodeResponse
	GetBody() *DescribeHasRuleNameByEventCodeResponseBody
}

type DescribeHasRuleNameByEventCodeResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeHasRuleNameByEventCodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeHasRuleNameByEventCodeResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeHasRuleNameByEventCodeResponse) GoString() string {
	return s.String()
}

func (s *DescribeHasRuleNameByEventCodeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeHasRuleNameByEventCodeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeHasRuleNameByEventCodeResponse) GetBody() *DescribeHasRuleNameByEventCodeResponseBody {
	return s.Body
}

func (s *DescribeHasRuleNameByEventCodeResponse) SetHeaders(v map[string]*string) *DescribeHasRuleNameByEventCodeResponse {
	s.Headers = v
	return s
}

func (s *DescribeHasRuleNameByEventCodeResponse) SetStatusCode(v int32) *DescribeHasRuleNameByEventCodeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeHasRuleNameByEventCodeResponse) SetBody(v *DescribeHasRuleNameByEventCodeResponseBody) *DescribeHasRuleNameByEventCodeResponse {
	s.Body = v
	return s
}

func (s *DescribeHasRuleNameByEventCodeResponse) Validate() error {
	return dara.Validate(s)
}
