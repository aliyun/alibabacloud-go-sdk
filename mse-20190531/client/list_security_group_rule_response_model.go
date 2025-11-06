// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSecurityGroupRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListSecurityGroupRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListSecurityGroupRuleResponse
	GetStatusCode() *int32
	SetBody(v *ListSecurityGroupRuleResponseBody) *ListSecurityGroupRuleResponse
	GetBody() *ListSecurityGroupRuleResponseBody
}

type ListSecurityGroupRuleResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSecurityGroupRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSecurityGroupRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s ListSecurityGroupRuleResponse) GoString() string {
	return s.String()
}

func (s *ListSecurityGroupRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListSecurityGroupRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListSecurityGroupRuleResponse) GetBody() *ListSecurityGroupRuleResponseBody {
	return s.Body
}

func (s *ListSecurityGroupRuleResponse) SetHeaders(v map[string]*string) *ListSecurityGroupRuleResponse {
	s.Headers = v
	return s
}

func (s *ListSecurityGroupRuleResponse) SetStatusCode(v int32) *ListSecurityGroupRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSecurityGroupRuleResponse) SetBody(v *ListSecurityGroupRuleResponseBody) *ListSecurityGroupRuleResponse {
	s.Body = v
	return s
}

func (s *ListSecurityGroupRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
