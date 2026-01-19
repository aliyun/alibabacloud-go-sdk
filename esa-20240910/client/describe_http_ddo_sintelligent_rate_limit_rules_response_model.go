// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHttpDDoSIntelligentRateLimitRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeHttpDDoSIntelligentRateLimitRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeHttpDDoSIntelligentRateLimitRulesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeHttpDDoSIntelligentRateLimitRulesResponseBody) *DescribeHttpDDoSIntelligentRateLimitRulesResponse
	GetBody() *DescribeHttpDDoSIntelligentRateLimitRulesResponseBody
}

type DescribeHttpDDoSIntelligentRateLimitRulesResponse struct {
	Headers    map[string]*string                                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeHttpDDoSIntelligentRateLimitRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeHttpDDoSIntelligentRateLimitRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeHttpDDoSIntelligentRateLimitRulesResponse) GoString() string {
	return s.String()
}

func (s *DescribeHttpDDoSIntelligentRateLimitRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeHttpDDoSIntelligentRateLimitRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeHttpDDoSIntelligentRateLimitRulesResponse) GetBody() *DescribeHttpDDoSIntelligentRateLimitRulesResponseBody {
	return s.Body
}

func (s *DescribeHttpDDoSIntelligentRateLimitRulesResponse) SetHeaders(v map[string]*string) *DescribeHttpDDoSIntelligentRateLimitRulesResponse {
	s.Headers = v
	return s
}

func (s *DescribeHttpDDoSIntelligentRateLimitRulesResponse) SetStatusCode(v int32) *DescribeHttpDDoSIntelligentRateLimitRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeHttpDDoSIntelligentRateLimitRulesResponse) SetBody(v *DescribeHttpDDoSIntelligentRateLimitRulesResponseBody) *DescribeHttpDDoSIntelligentRateLimitRulesResponse {
	s.Body = v
	return s
}

func (s *DescribeHttpDDoSIntelligentRateLimitRulesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
