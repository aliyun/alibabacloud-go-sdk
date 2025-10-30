// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachSecurityGroupFromVpcEndpointResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DetachSecurityGroupFromVpcEndpointResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DetachSecurityGroupFromVpcEndpointResponse
	GetStatusCode() *int32
	SetBody(v *DetachSecurityGroupFromVpcEndpointResponseBody) *DetachSecurityGroupFromVpcEndpointResponse
	GetBody() *DetachSecurityGroupFromVpcEndpointResponseBody
}

type DetachSecurityGroupFromVpcEndpointResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DetachSecurityGroupFromVpcEndpointResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DetachSecurityGroupFromVpcEndpointResponse) String() string {
	return dara.Prettify(s)
}

func (s DetachSecurityGroupFromVpcEndpointResponse) GoString() string {
	return s.String()
}

func (s *DetachSecurityGroupFromVpcEndpointResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DetachSecurityGroupFromVpcEndpointResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DetachSecurityGroupFromVpcEndpointResponse) GetBody() *DetachSecurityGroupFromVpcEndpointResponseBody {
	return s.Body
}

func (s *DetachSecurityGroupFromVpcEndpointResponse) SetHeaders(v map[string]*string) *DetachSecurityGroupFromVpcEndpointResponse {
	s.Headers = v
	return s
}

func (s *DetachSecurityGroupFromVpcEndpointResponse) SetStatusCode(v int32) *DetachSecurityGroupFromVpcEndpointResponse {
	s.StatusCode = &v
	return s
}

func (s *DetachSecurityGroupFromVpcEndpointResponse) SetBody(v *DetachSecurityGroupFromVpcEndpointResponseBody) *DetachSecurityGroupFromVpcEndpointResponse {
	s.Body = v
	return s
}

func (s *DetachSecurityGroupFromVpcEndpointResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
