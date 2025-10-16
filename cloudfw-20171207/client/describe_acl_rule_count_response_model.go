// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAclRuleCountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAclRuleCountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAclRuleCountResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAclRuleCountResponseBody) *DescribeAclRuleCountResponse
	GetBody() *DescribeAclRuleCountResponseBody
}

type DescribeAclRuleCountResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAclRuleCountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAclRuleCountResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAclRuleCountResponse) GoString() string {
	return s.String()
}

func (s *DescribeAclRuleCountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAclRuleCountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAclRuleCountResponse) GetBody() *DescribeAclRuleCountResponseBody {
	return s.Body
}

func (s *DescribeAclRuleCountResponse) SetHeaders(v map[string]*string) *DescribeAclRuleCountResponse {
	s.Headers = v
	return s
}

func (s *DescribeAclRuleCountResponse) SetStatusCode(v int32) *DescribeAclRuleCountResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAclRuleCountResponse) SetBody(v *DescribeAclRuleCountResponseBody) *DescribeAclRuleCountResponse {
	s.Body = v
	return s
}

func (s *DescribeAclRuleCountResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
