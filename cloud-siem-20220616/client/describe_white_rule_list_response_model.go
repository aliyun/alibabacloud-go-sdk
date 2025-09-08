// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeWhiteRuleListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeWhiteRuleListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeWhiteRuleListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeWhiteRuleListResponseBody) *DescribeWhiteRuleListResponse
	GetBody() *DescribeWhiteRuleListResponseBody
}

type DescribeWhiteRuleListResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeWhiteRuleListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeWhiteRuleListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeWhiteRuleListResponse) GoString() string {
	return s.String()
}

func (s *DescribeWhiteRuleListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeWhiteRuleListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeWhiteRuleListResponse) GetBody() *DescribeWhiteRuleListResponseBody {
	return s.Body
}

func (s *DescribeWhiteRuleListResponse) SetHeaders(v map[string]*string) *DescribeWhiteRuleListResponse {
	s.Headers = v
	return s
}

func (s *DescribeWhiteRuleListResponse) SetStatusCode(v int32) *DescribeWhiteRuleListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeWhiteRuleListResponse) SetBody(v *DescribeWhiteRuleListResponseBody) *DescribeWhiteRuleListResponse {
	s.Body = v
	return s
}

func (s *DescribeWhiteRuleListResponse) Validate() error {
	return dara.Validate(s)
}
