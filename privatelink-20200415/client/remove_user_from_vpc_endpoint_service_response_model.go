// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveUserFromVpcEndpointServiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RemoveUserFromVpcEndpointServiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RemoveUserFromVpcEndpointServiceResponse
	GetStatusCode() *int32
	SetBody(v *RemoveUserFromVpcEndpointServiceResponseBody) *RemoveUserFromVpcEndpointServiceResponse
	GetBody() *RemoveUserFromVpcEndpointServiceResponseBody
}

type RemoveUserFromVpcEndpointServiceResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveUserFromVpcEndpointServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveUserFromVpcEndpointServiceResponse) String() string {
	return dara.Prettify(s)
}

func (s RemoveUserFromVpcEndpointServiceResponse) GoString() string {
	return s.String()
}

func (s *RemoveUserFromVpcEndpointServiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RemoveUserFromVpcEndpointServiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RemoveUserFromVpcEndpointServiceResponse) GetBody() *RemoveUserFromVpcEndpointServiceResponseBody {
	return s.Body
}

func (s *RemoveUserFromVpcEndpointServiceResponse) SetHeaders(v map[string]*string) *RemoveUserFromVpcEndpointServiceResponse {
	s.Headers = v
	return s
}

func (s *RemoveUserFromVpcEndpointServiceResponse) SetStatusCode(v int32) *RemoveUserFromVpcEndpointServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveUserFromVpcEndpointServiceResponse) SetBody(v *RemoveUserFromVpcEndpointServiceResponseBody) *RemoveUserFromVpcEndpointServiceResponse {
	s.Body = v
	return s
}

func (s *RemoveUserFromVpcEndpointServiceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
