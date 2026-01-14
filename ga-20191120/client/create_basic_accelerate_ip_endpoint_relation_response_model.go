// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBasicAccelerateIpEndpointRelationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateBasicAccelerateIpEndpointRelationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateBasicAccelerateIpEndpointRelationResponse
	GetStatusCode() *int32
	SetBody(v *CreateBasicAccelerateIpEndpointRelationResponseBody) *CreateBasicAccelerateIpEndpointRelationResponse
	GetBody() *CreateBasicAccelerateIpEndpointRelationResponseBody
}

type CreateBasicAccelerateIpEndpointRelationResponse struct {
	Headers    map[string]*string                                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateBasicAccelerateIpEndpointRelationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateBasicAccelerateIpEndpointRelationResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateBasicAccelerateIpEndpointRelationResponse) GoString() string {
	return s.String()
}

func (s *CreateBasicAccelerateIpEndpointRelationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateBasicAccelerateIpEndpointRelationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateBasicAccelerateIpEndpointRelationResponse) GetBody() *CreateBasicAccelerateIpEndpointRelationResponseBody {
	return s.Body
}

func (s *CreateBasicAccelerateIpEndpointRelationResponse) SetHeaders(v map[string]*string) *CreateBasicAccelerateIpEndpointRelationResponse {
	s.Headers = v
	return s
}

func (s *CreateBasicAccelerateIpEndpointRelationResponse) SetStatusCode(v int32) *CreateBasicAccelerateIpEndpointRelationResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateBasicAccelerateIpEndpointRelationResponse) SetBody(v *CreateBasicAccelerateIpEndpointRelationResponseBody) *CreateBasicAccelerateIpEndpointRelationResponse {
	s.Body = v
	return s
}

func (s *CreateBasicAccelerateIpEndpointRelationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
