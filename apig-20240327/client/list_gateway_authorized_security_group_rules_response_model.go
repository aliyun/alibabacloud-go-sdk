// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGatewayAuthorizedSecurityGroupRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListGatewayAuthorizedSecurityGroupRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListGatewayAuthorizedSecurityGroupRulesResponse
	GetStatusCode() *int32
	SetBody(v *ListGatewayAuthorizedSecurityGroupRulesResponseBody) *ListGatewayAuthorizedSecurityGroupRulesResponse
	GetBody() *ListGatewayAuthorizedSecurityGroupRulesResponseBody
}

type ListGatewayAuthorizedSecurityGroupRulesResponse struct {
	Headers    map[string]*string                                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListGatewayAuthorizedSecurityGroupRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListGatewayAuthorizedSecurityGroupRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListGatewayAuthorizedSecurityGroupRulesResponse) GoString() string {
	return s.String()
}

func (s *ListGatewayAuthorizedSecurityGroupRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListGatewayAuthorizedSecurityGroupRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListGatewayAuthorizedSecurityGroupRulesResponse) GetBody() *ListGatewayAuthorizedSecurityGroupRulesResponseBody {
	return s.Body
}

func (s *ListGatewayAuthorizedSecurityGroupRulesResponse) SetHeaders(v map[string]*string) *ListGatewayAuthorizedSecurityGroupRulesResponse {
	s.Headers = v
	return s
}

func (s *ListGatewayAuthorizedSecurityGroupRulesResponse) SetStatusCode(v int32) *ListGatewayAuthorizedSecurityGroupRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListGatewayAuthorizedSecurityGroupRulesResponse) SetBody(v *ListGatewayAuthorizedSecurityGroupRulesResponseBody) *ListGatewayAuthorizedSecurityGroupRulesResponse {
	s.Body = v
	return s
}

func (s *ListGatewayAuthorizedSecurityGroupRulesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
