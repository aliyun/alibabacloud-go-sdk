// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGatewayAuthorizableSecurityGroupsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListGatewayAuthorizableSecurityGroupsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListGatewayAuthorizableSecurityGroupsResponse
	GetStatusCode() *int32
	SetBody(v *ListGatewayAuthorizableSecurityGroupsResponseBody) *ListGatewayAuthorizableSecurityGroupsResponse
	GetBody() *ListGatewayAuthorizableSecurityGroupsResponseBody
}

type ListGatewayAuthorizableSecurityGroupsResponse struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListGatewayAuthorizableSecurityGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListGatewayAuthorizableSecurityGroupsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListGatewayAuthorizableSecurityGroupsResponse) GoString() string {
	return s.String()
}

func (s *ListGatewayAuthorizableSecurityGroupsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListGatewayAuthorizableSecurityGroupsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListGatewayAuthorizableSecurityGroupsResponse) GetBody() *ListGatewayAuthorizableSecurityGroupsResponseBody {
	return s.Body
}

func (s *ListGatewayAuthorizableSecurityGroupsResponse) SetHeaders(v map[string]*string) *ListGatewayAuthorizableSecurityGroupsResponse {
	s.Headers = v
	return s
}

func (s *ListGatewayAuthorizableSecurityGroupsResponse) SetStatusCode(v int32) *ListGatewayAuthorizableSecurityGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListGatewayAuthorizableSecurityGroupsResponse) SetBody(v *ListGatewayAuthorizableSecurityGroupsResponseBody) *ListGatewayAuthorizableSecurityGroupsResponse {
	s.Body = v
	return s
}

func (s *ListGatewayAuthorizableSecurityGroupsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
