// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachResourceFromVpcEndpointServiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DetachResourceFromVpcEndpointServiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DetachResourceFromVpcEndpointServiceResponse
	GetStatusCode() *int32
	SetBody(v *DetachResourceFromVpcEndpointServiceResponseBody) *DetachResourceFromVpcEndpointServiceResponse
	GetBody() *DetachResourceFromVpcEndpointServiceResponseBody
}

type DetachResourceFromVpcEndpointServiceResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DetachResourceFromVpcEndpointServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DetachResourceFromVpcEndpointServiceResponse) String() string {
	return dara.Prettify(s)
}

func (s DetachResourceFromVpcEndpointServiceResponse) GoString() string {
	return s.String()
}

func (s *DetachResourceFromVpcEndpointServiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DetachResourceFromVpcEndpointServiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DetachResourceFromVpcEndpointServiceResponse) GetBody() *DetachResourceFromVpcEndpointServiceResponseBody {
	return s.Body
}

func (s *DetachResourceFromVpcEndpointServiceResponse) SetHeaders(v map[string]*string) *DetachResourceFromVpcEndpointServiceResponse {
	s.Headers = v
	return s
}

func (s *DetachResourceFromVpcEndpointServiceResponse) SetStatusCode(v int32) *DetachResourceFromVpcEndpointServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *DetachResourceFromVpcEndpointServiceResponse) SetBody(v *DetachResourceFromVpcEndpointServiceResponseBody) *DetachResourceFromVpcEndpointServiceResponse {
	s.Body = v
	return s
}

func (s *DetachResourceFromVpcEndpointServiceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
