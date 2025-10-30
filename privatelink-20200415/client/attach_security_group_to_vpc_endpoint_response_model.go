// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachSecurityGroupToVpcEndpointResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AttachSecurityGroupToVpcEndpointResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AttachSecurityGroupToVpcEndpointResponse
	GetStatusCode() *int32
	SetBody(v *AttachSecurityGroupToVpcEndpointResponseBody) *AttachSecurityGroupToVpcEndpointResponse
	GetBody() *AttachSecurityGroupToVpcEndpointResponseBody
}

type AttachSecurityGroupToVpcEndpointResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AttachSecurityGroupToVpcEndpointResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AttachSecurityGroupToVpcEndpointResponse) String() string {
	return dara.Prettify(s)
}

func (s AttachSecurityGroupToVpcEndpointResponse) GoString() string {
	return s.String()
}

func (s *AttachSecurityGroupToVpcEndpointResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AttachSecurityGroupToVpcEndpointResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AttachSecurityGroupToVpcEndpointResponse) GetBody() *AttachSecurityGroupToVpcEndpointResponseBody {
	return s.Body
}

func (s *AttachSecurityGroupToVpcEndpointResponse) SetHeaders(v map[string]*string) *AttachSecurityGroupToVpcEndpointResponse {
	s.Headers = v
	return s
}

func (s *AttachSecurityGroupToVpcEndpointResponse) SetStatusCode(v int32) *AttachSecurityGroupToVpcEndpointResponse {
	s.StatusCode = &v
	return s
}

func (s *AttachSecurityGroupToVpcEndpointResponse) SetBody(v *AttachSecurityGroupToVpcEndpointResponseBody) *AttachSecurityGroupToVpcEndpointResponse {
	s.Body = v
	return s
}

func (s *AttachSecurityGroupToVpcEndpointResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
