// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVpcEndpointServiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteVpcEndpointServiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteVpcEndpointServiceResponse
	GetStatusCode() *int32
	SetBody(v *DeleteVpcEndpointServiceResponseBody) *DeleteVpcEndpointServiceResponse
	GetBody() *DeleteVpcEndpointServiceResponseBody
}

type DeleteVpcEndpointServiceResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteVpcEndpointServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteVpcEndpointServiceResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteVpcEndpointServiceResponse) GoString() string {
	return s.String()
}

func (s *DeleteVpcEndpointServiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteVpcEndpointServiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteVpcEndpointServiceResponse) GetBody() *DeleteVpcEndpointServiceResponseBody {
	return s.Body
}

func (s *DeleteVpcEndpointServiceResponse) SetHeaders(v map[string]*string) *DeleteVpcEndpointServiceResponse {
	s.Headers = v
	return s
}

func (s *DeleteVpcEndpointServiceResponse) SetStatusCode(v int32) *DeleteVpcEndpointServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteVpcEndpointServiceResponse) SetBody(v *DeleteVpcEndpointServiceResponseBody) *DeleteVpcEndpointServiceResponse {
	s.Body = v
	return s
}

func (s *DeleteVpcEndpointServiceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
