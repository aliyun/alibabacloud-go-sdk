// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRuleDetailByRuleIdResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeRuleDetailByRuleIdResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeRuleDetailByRuleIdResponse
	GetStatusCode() *int32
	SetBody(v *DescribeRuleDetailByRuleIdResponseBody) *DescribeRuleDetailByRuleIdResponse
	GetBody() *DescribeRuleDetailByRuleIdResponseBody
}

type DescribeRuleDetailByRuleIdResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRuleDetailByRuleIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRuleDetailByRuleIdResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeRuleDetailByRuleIdResponse) GoString() string {
	return s.String()
}

func (s *DescribeRuleDetailByRuleIdResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeRuleDetailByRuleIdResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeRuleDetailByRuleIdResponse) GetBody() *DescribeRuleDetailByRuleIdResponseBody {
	return s.Body
}

func (s *DescribeRuleDetailByRuleIdResponse) SetHeaders(v map[string]*string) *DescribeRuleDetailByRuleIdResponse {
	s.Headers = v
	return s
}

func (s *DescribeRuleDetailByRuleIdResponse) SetStatusCode(v int32) *DescribeRuleDetailByRuleIdResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRuleDetailByRuleIdResponse) SetBody(v *DescribeRuleDetailByRuleIdResponseBody) *DescribeRuleDetailByRuleIdResponse {
	s.Body = v
	return s
}

func (s *DescribeRuleDetailByRuleIdResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
