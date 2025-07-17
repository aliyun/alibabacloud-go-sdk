// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteGatewaySecurityGroupRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteGatewaySecurityGroupRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteGatewaySecurityGroupRuleResponse
	GetStatusCode() *int32
	SetBody(v *DeleteGatewaySecurityGroupRuleResponseBody) *DeleteGatewaySecurityGroupRuleResponse
	GetBody() *DeleteGatewaySecurityGroupRuleResponseBody
}

type DeleteGatewaySecurityGroupRuleResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteGatewaySecurityGroupRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteGatewaySecurityGroupRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteGatewaySecurityGroupRuleResponse) GoString() string {
	return s.String()
}

func (s *DeleteGatewaySecurityGroupRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteGatewaySecurityGroupRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteGatewaySecurityGroupRuleResponse) GetBody() *DeleteGatewaySecurityGroupRuleResponseBody {
	return s.Body
}

func (s *DeleteGatewaySecurityGroupRuleResponse) SetHeaders(v map[string]*string) *DeleteGatewaySecurityGroupRuleResponse {
	s.Headers = v
	return s
}

func (s *DeleteGatewaySecurityGroupRuleResponse) SetStatusCode(v int32) *DeleteGatewaySecurityGroupRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteGatewaySecurityGroupRuleResponse) SetBody(v *DeleteGatewaySecurityGroupRuleResponseBody) *DeleteGatewaySecurityGroupRuleResponse {
	s.Body = v
	return s
}

func (s *DeleteGatewaySecurityGroupRuleResponse) Validate() error {
	return dara.Validate(s)
}
