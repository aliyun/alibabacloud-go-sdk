// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeResolverRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeResolverRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeResolverRulesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeResolverRulesResponseBody) *DescribeResolverRulesResponse
	GetBody() *DescribeResolverRulesResponseBody
}

type DescribeResolverRulesResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeResolverRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeResolverRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeResolverRulesResponse) GoString() string {
	return s.String()
}

func (s *DescribeResolverRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeResolverRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeResolverRulesResponse) GetBody() *DescribeResolverRulesResponseBody {
	return s.Body
}

func (s *DescribeResolverRulesResponse) SetHeaders(v map[string]*string) *DescribeResolverRulesResponse {
	s.Headers = v
	return s
}

func (s *DescribeResolverRulesResponse) SetStatusCode(v int32) *DescribeResolverRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeResolverRulesResponse) SetBody(v *DescribeResolverRulesResponseBody) *DescribeResolverRulesResponse {
	s.Body = v
	return s
}

func (s *DescribeResolverRulesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
