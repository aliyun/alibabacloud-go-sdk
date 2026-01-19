// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHttpDDoSAttackRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeHttpDDoSAttackRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeHttpDDoSAttackRulesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeHttpDDoSAttackRulesResponseBody) *DescribeHttpDDoSAttackRulesResponse
	GetBody() *DescribeHttpDDoSAttackRulesResponseBody
}

type DescribeHttpDDoSAttackRulesResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeHttpDDoSAttackRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeHttpDDoSAttackRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeHttpDDoSAttackRulesResponse) GoString() string {
	return s.String()
}

func (s *DescribeHttpDDoSAttackRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeHttpDDoSAttackRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeHttpDDoSAttackRulesResponse) GetBody() *DescribeHttpDDoSAttackRulesResponseBody {
	return s.Body
}

func (s *DescribeHttpDDoSAttackRulesResponse) SetHeaders(v map[string]*string) *DescribeHttpDDoSAttackRulesResponse {
	s.Headers = v
	return s
}

func (s *DescribeHttpDDoSAttackRulesResponse) SetStatusCode(v int32) *DescribeHttpDDoSAttackRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeHttpDDoSAttackRulesResponse) SetBody(v *DescribeHttpDDoSAttackRulesResponseBody) *DescribeHttpDDoSAttackRulesResponse {
	s.Body = v
	return s
}

func (s *DescribeHttpDDoSAttackRulesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
