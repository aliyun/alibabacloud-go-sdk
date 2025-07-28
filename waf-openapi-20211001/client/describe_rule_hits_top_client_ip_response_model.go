// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRuleHitsTopClientIpResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeRuleHitsTopClientIpResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeRuleHitsTopClientIpResponse
	GetStatusCode() *int32
	SetBody(v *DescribeRuleHitsTopClientIpResponseBody) *DescribeRuleHitsTopClientIpResponse
	GetBody() *DescribeRuleHitsTopClientIpResponseBody
}

type DescribeRuleHitsTopClientIpResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRuleHitsTopClientIpResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRuleHitsTopClientIpResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeRuleHitsTopClientIpResponse) GoString() string {
	return s.String()
}

func (s *DescribeRuleHitsTopClientIpResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeRuleHitsTopClientIpResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeRuleHitsTopClientIpResponse) GetBody() *DescribeRuleHitsTopClientIpResponseBody {
	return s.Body
}

func (s *DescribeRuleHitsTopClientIpResponse) SetHeaders(v map[string]*string) *DescribeRuleHitsTopClientIpResponse {
	s.Headers = v
	return s
}

func (s *DescribeRuleHitsTopClientIpResponse) SetStatusCode(v int32) *DescribeRuleHitsTopClientIpResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRuleHitsTopClientIpResponse) SetBody(v *DescribeRuleHitsTopClientIpResponseBody) *DescribeRuleHitsTopClientIpResponse {
	s.Body = v
	return s
}

func (s *DescribeRuleHitsTopClientIpResponse) Validate() error {
	return dara.Validate(s)
}
