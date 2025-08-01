// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddSecurityGroupRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddSecurityGroupRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddSecurityGroupRuleResponse
	GetStatusCode() *int32
	SetBody(v *AddSecurityGroupRuleResponseBody) *AddSecurityGroupRuleResponse
	GetBody() *AddSecurityGroupRuleResponseBody
}

type AddSecurityGroupRuleResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddSecurityGroupRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddSecurityGroupRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s AddSecurityGroupRuleResponse) GoString() string {
	return s.String()
}

func (s *AddSecurityGroupRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddSecurityGroupRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddSecurityGroupRuleResponse) GetBody() *AddSecurityGroupRuleResponseBody {
	return s.Body
}

func (s *AddSecurityGroupRuleResponse) SetHeaders(v map[string]*string) *AddSecurityGroupRuleResponse {
	s.Headers = v
	return s
}

func (s *AddSecurityGroupRuleResponse) SetStatusCode(v int32) *AddSecurityGroupRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *AddSecurityGroupRuleResponse) SetBody(v *AddSecurityGroupRuleResponseBody) *AddSecurityGroupRuleResponse {
	s.Body = v
	return s
}

func (s *AddSecurityGroupRuleResponse) Validate() error {
	return dara.Validate(s)
}
