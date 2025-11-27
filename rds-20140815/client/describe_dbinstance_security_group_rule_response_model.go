// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBInstanceSecurityGroupRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDBInstanceSecurityGroupRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDBInstanceSecurityGroupRuleResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDBInstanceSecurityGroupRuleResponseBody) *DescribeDBInstanceSecurityGroupRuleResponse
	GetBody() *DescribeDBInstanceSecurityGroupRuleResponseBody
}

type DescribeDBInstanceSecurityGroupRuleResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDBInstanceSecurityGroupRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDBInstanceSecurityGroupRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceSecurityGroupRuleResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceSecurityGroupRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDBInstanceSecurityGroupRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDBInstanceSecurityGroupRuleResponse) GetBody() *DescribeDBInstanceSecurityGroupRuleResponseBody {
	return s.Body
}

func (s *DescribeDBInstanceSecurityGroupRuleResponse) SetHeaders(v map[string]*string) *DescribeDBInstanceSecurityGroupRuleResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBInstanceSecurityGroupRuleResponse) SetStatusCode(v int32) *DescribeDBInstanceSecurityGroupRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDBInstanceSecurityGroupRuleResponse) SetBody(v *DescribeDBInstanceSecurityGroupRuleResponseBody) *DescribeDBInstanceSecurityGroupRuleResponse {
	s.Body = v
	return s
}

func (s *DescribeDBInstanceSecurityGroupRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
