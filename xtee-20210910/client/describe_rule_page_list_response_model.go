// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRulePageListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeRulePageListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeRulePageListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeRulePageListResponseBody) *DescribeRulePageListResponse
	GetBody() *DescribeRulePageListResponseBody
}

type DescribeRulePageListResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRulePageListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRulePageListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeRulePageListResponse) GoString() string {
	return s.String()
}

func (s *DescribeRulePageListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeRulePageListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeRulePageListResponse) GetBody() *DescribeRulePageListResponseBody {
	return s.Body
}

func (s *DescribeRulePageListResponse) SetHeaders(v map[string]*string) *DescribeRulePageListResponse {
	s.Headers = v
	return s
}

func (s *DescribeRulePageListResponse) SetStatusCode(v int32) *DescribeRulePageListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRulePageListResponse) SetBody(v *DescribeRulePageListResponseBody) *DescribeRulePageListResponse {
	s.Body = v
	return s
}

func (s *DescribeRulePageListResponse) Validate() error {
	return dara.Validate(s)
}
