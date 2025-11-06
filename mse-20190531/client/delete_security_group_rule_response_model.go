// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSecurityGroupRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteSecurityGroupRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteSecurityGroupRuleResponse
	GetStatusCode() *int32
	SetBody(v *DeleteSecurityGroupRuleResponseBody) *DeleteSecurityGroupRuleResponse
	GetBody() *DeleteSecurityGroupRuleResponseBody
}

type DeleteSecurityGroupRuleResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteSecurityGroupRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteSecurityGroupRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteSecurityGroupRuleResponse) GoString() string {
	return s.String()
}

func (s *DeleteSecurityGroupRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteSecurityGroupRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteSecurityGroupRuleResponse) GetBody() *DeleteSecurityGroupRuleResponseBody {
	return s.Body
}

func (s *DeleteSecurityGroupRuleResponse) SetHeaders(v map[string]*string) *DeleteSecurityGroupRuleResponse {
	s.Headers = v
	return s
}

func (s *DeleteSecurityGroupRuleResponse) SetStatusCode(v int32) *DeleteSecurityGroupRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSecurityGroupRuleResponse) SetBody(v *DeleteSecurityGroupRuleResponseBody) *DeleteSecurityGroupRuleResponse {
	s.Body = v
	return s
}

func (s *DeleteSecurityGroupRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
