// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddGatewaySecurityGroupRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddGatewaySecurityGroupRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddGatewaySecurityGroupRuleResponse
	GetStatusCode() *int32
	SetBody(v *AddGatewaySecurityGroupRuleResponseBody) *AddGatewaySecurityGroupRuleResponse
	GetBody() *AddGatewaySecurityGroupRuleResponseBody
}

type AddGatewaySecurityGroupRuleResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddGatewaySecurityGroupRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddGatewaySecurityGroupRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s AddGatewaySecurityGroupRuleResponse) GoString() string {
	return s.String()
}

func (s *AddGatewaySecurityGroupRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddGatewaySecurityGroupRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddGatewaySecurityGroupRuleResponse) GetBody() *AddGatewaySecurityGroupRuleResponseBody {
	return s.Body
}

func (s *AddGatewaySecurityGroupRuleResponse) SetHeaders(v map[string]*string) *AddGatewaySecurityGroupRuleResponse {
	s.Headers = v
	return s
}

func (s *AddGatewaySecurityGroupRuleResponse) SetStatusCode(v int32) *AddGatewaySecurityGroupRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *AddGatewaySecurityGroupRuleResponse) SetBody(v *AddGatewaySecurityGroupRuleResponseBody) *AddGatewaySecurityGroupRuleResponse {
	s.Body = v
	return s
}

func (s *AddGatewaySecurityGroupRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
