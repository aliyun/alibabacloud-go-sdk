// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBasicAccelerateIpEndpointRelationsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateBasicAccelerateIpEndpointRelationsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateBasicAccelerateIpEndpointRelationsResponse
	GetStatusCode() *int32
	SetBody(v *CreateBasicAccelerateIpEndpointRelationsResponseBody) *CreateBasicAccelerateIpEndpointRelationsResponse
	GetBody() *CreateBasicAccelerateIpEndpointRelationsResponseBody
}

type CreateBasicAccelerateIpEndpointRelationsResponse struct {
	Headers    map[string]*string                                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateBasicAccelerateIpEndpointRelationsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateBasicAccelerateIpEndpointRelationsResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateBasicAccelerateIpEndpointRelationsResponse) GoString() string {
	return s.String()
}

func (s *CreateBasicAccelerateIpEndpointRelationsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateBasicAccelerateIpEndpointRelationsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateBasicAccelerateIpEndpointRelationsResponse) GetBody() *CreateBasicAccelerateIpEndpointRelationsResponseBody {
	return s.Body
}

func (s *CreateBasicAccelerateIpEndpointRelationsResponse) SetHeaders(v map[string]*string) *CreateBasicAccelerateIpEndpointRelationsResponse {
	s.Headers = v
	return s
}

func (s *CreateBasicAccelerateIpEndpointRelationsResponse) SetStatusCode(v int32) *CreateBasicAccelerateIpEndpointRelationsResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateBasicAccelerateIpEndpointRelationsResponse) SetBody(v *CreateBasicAccelerateIpEndpointRelationsResponseBody) *CreateBasicAccelerateIpEndpointRelationsResponse {
	s.Body = v
	return s
}

func (s *CreateBasicAccelerateIpEndpointRelationsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
