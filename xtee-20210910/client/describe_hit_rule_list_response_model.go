// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHitRuleListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeHitRuleListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeHitRuleListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeHitRuleListResponseBody) *DescribeHitRuleListResponse
	GetBody() *DescribeHitRuleListResponseBody
}

type DescribeHitRuleListResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeHitRuleListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeHitRuleListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeHitRuleListResponse) GoString() string {
	return s.String()
}

func (s *DescribeHitRuleListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeHitRuleListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeHitRuleListResponse) GetBody() *DescribeHitRuleListResponseBody {
	return s.Body
}

func (s *DescribeHitRuleListResponse) SetHeaders(v map[string]*string) *DescribeHitRuleListResponse {
	s.Headers = v
	return s
}

func (s *DescribeHitRuleListResponse) SetStatusCode(v int32) *DescribeHitRuleListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeHitRuleListResponse) SetBody(v *DescribeHitRuleListResponseBody) *DescribeHitRuleListResponse {
	s.Body = v
	return s
}

func (s *DescribeHitRuleListResponse) Validate() error {
	return dara.Validate(s)
}
