// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRuleCountByUserIdResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeRuleCountByUserIdResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeRuleCountByUserIdResponse
	GetStatusCode() *int32
	SetBody(v *DescribeRuleCountByUserIdResponseBody) *DescribeRuleCountByUserIdResponse
	GetBody() *DescribeRuleCountByUserIdResponseBody
}

type DescribeRuleCountByUserIdResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRuleCountByUserIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRuleCountByUserIdResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeRuleCountByUserIdResponse) GoString() string {
	return s.String()
}

func (s *DescribeRuleCountByUserIdResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeRuleCountByUserIdResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeRuleCountByUserIdResponse) GetBody() *DescribeRuleCountByUserIdResponseBody {
	return s.Body
}

func (s *DescribeRuleCountByUserIdResponse) SetHeaders(v map[string]*string) *DescribeRuleCountByUserIdResponse {
	s.Headers = v
	return s
}

func (s *DescribeRuleCountByUserIdResponse) SetStatusCode(v int32) *DescribeRuleCountByUserIdResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRuleCountByUserIdResponse) SetBody(v *DescribeRuleCountByUserIdResponseBody) *DescribeRuleCountByUserIdResponse {
	s.Body = v
	return s
}

func (s *DescribeRuleCountByUserIdResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
