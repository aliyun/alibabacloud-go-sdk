// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachResourceToVpcEndpointServiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AttachResourceToVpcEndpointServiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AttachResourceToVpcEndpointServiceResponse
	GetStatusCode() *int32
	SetBody(v *AttachResourceToVpcEndpointServiceResponseBody) *AttachResourceToVpcEndpointServiceResponse
	GetBody() *AttachResourceToVpcEndpointServiceResponseBody
}

type AttachResourceToVpcEndpointServiceResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AttachResourceToVpcEndpointServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AttachResourceToVpcEndpointServiceResponse) String() string {
	return dara.Prettify(s)
}

func (s AttachResourceToVpcEndpointServiceResponse) GoString() string {
	return s.String()
}

func (s *AttachResourceToVpcEndpointServiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AttachResourceToVpcEndpointServiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AttachResourceToVpcEndpointServiceResponse) GetBody() *AttachResourceToVpcEndpointServiceResponseBody {
	return s.Body
}

func (s *AttachResourceToVpcEndpointServiceResponse) SetHeaders(v map[string]*string) *AttachResourceToVpcEndpointServiceResponse {
	s.Headers = v
	return s
}

func (s *AttachResourceToVpcEndpointServiceResponse) SetStatusCode(v int32) *AttachResourceToVpcEndpointServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *AttachResourceToVpcEndpointServiceResponse) SetBody(v *AttachResourceToVpcEndpointServiceResponseBody) *AttachResourceToVpcEndpointServiceResponse {
	s.Body = v
	return s
}

func (s *AttachResourceToVpcEndpointServiceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
