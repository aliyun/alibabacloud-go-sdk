// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRuleListByEventCodesListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeRuleListByEventCodesListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeRuleListByEventCodesListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeRuleListByEventCodesListResponseBody) *DescribeRuleListByEventCodesListResponse
	GetBody() *DescribeRuleListByEventCodesListResponseBody
}

type DescribeRuleListByEventCodesListResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRuleListByEventCodesListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRuleListByEventCodesListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeRuleListByEventCodesListResponse) GoString() string {
	return s.String()
}

func (s *DescribeRuleListByEventCodesListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeRuleListByEventCodesListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeRuleListByEventCodesListResponse) GetBody() *DescribeRuleListByEventCodesListResponseBody {
	return s.Body
}

func (s *DescribeRuleListByEventCodesListResponse) SetHeaders(v map[string]*string) *DescribeRuleListByEventCodesListResponse {
	s.Headers = v
	return s
}

func (s *DescribeRuleListByEventCodesListResponse) SetStatusCode(v int32) *DescribeRuleListByEventCodesListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRuleListByEventCodesListResponse) SetBody(v *DescribeRuleListByEventCodesListResponseBody) *DescribeRuleListByEventCodesListResponse {
	s.Body = v
	return s
}

func (s *DescribeRuleListByEventCodesListResponse) Validate() error {
	return dara.Validate(s)
}
