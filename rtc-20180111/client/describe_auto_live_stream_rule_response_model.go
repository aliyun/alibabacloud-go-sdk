// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAutoLiveStreamRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAutoLiveStreamRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAutoLiveStreamRuleResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAutoLiveStreamRuleResponseBody) *DescribeAutoLiveStreamRuleResponse
	GetBody() *DescribeAutoLiveStreamRuleResponseBody
}

type DescribeAutoLiveStreamRuleResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAutoLiveStreamRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAutoLiveStreamRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAutoLiveStreamRuleResponse) GoString() string {
	return s.String()
}

func (s *DescribeAutoLiveStreamRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAutoLiveStreamRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAutoLiveStreamRuleResponse) GetBody() *DescribeAutoLiveStreamRuleResponseBody {
	return s.Body
}

func (s *DescribeAutoLiveStreamRuleResponse) SetHeaders(v map[string]*string) *DescribeAutoLiveStreamRuleResponse {
	s.Headers = v
	return s
}

func (s *DescribeAutoLiveStreamRuleResponse) SetStatusCode(v int32) *DescribeAutoLiveStreamRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAutoLiveStreamRuleResponse) SetBody(v *DescribeAutoLiveStreamRuleResponseBody) *DescribeAutoLiveStreamRuleResponse {
	s.Body = v
	return s
}

func (s *DescribeAutoLiveStreamRuleResponse) Validate() error {
	return dara.Validate(s)
}
