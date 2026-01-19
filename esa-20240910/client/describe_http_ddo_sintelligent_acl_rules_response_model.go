// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHttpDDoSIntelligentAclRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeHttpDDoSIntelligentAclRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeHttpDDoSIntelligentAclRulesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeHttpDDoSIntelligentAclRulesResponseBody) *DescribeHttpDDoSIntelligentAclRulesResponse
	GetBody() *DescribeHttpDDoSIntelligentAclRulesResponseBody
}

type DescribeHttpDDoSIntelligentAclRulesResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeHttpDDoSIntelligentAclRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeHttpDDoSIntelligentAclRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeHttpDDoSIntelligentAclRulesResponse) GoString() string {
	return s.String()
}

func (s *DescribeHttpDDoSIntelligentAclRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeHttpDDoSIntelligentAclRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeHttpDDoSIntelligentAclRulesResponse) GetBody() *DescribeHttpDDoSIntelligentAclRulesResponseBody {
	return s.Body
}

func (s *DescribeHttpDDoSIntelligentAclRulesResponse) SetHeaders(v map[string]*string) *DescribeHttpDDoSIntelligentAclRulesResponse {
	s.Headers = v
	return s
}

func (s *DescribeHttpDDoSIntelligentAclRulesResponse) SetStatusCode(v int32) *DescribeHttpDDoSIntelligentAclRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeHttpDDoSIntelligentAclRulesResponse) SetBody(v *DescribeHttpDDoSIntelligentAclRulesResponseBody) *DescribeHttpDDoSIntelligentAclRulesResponse {
	s.Body = v
	return s
}

func (s *DescribeHttpDDoSIntelligentAclRulesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
