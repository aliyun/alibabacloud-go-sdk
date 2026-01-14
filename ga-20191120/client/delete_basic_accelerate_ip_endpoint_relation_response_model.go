// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBasicAccelerateIpEndpointRelationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteBasicAccelerateIpEndpointRelationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteBasicAccelerateIpEndpointRelationResponse
	GetStatusCode() *int32
	SetBody(v *DeleteBasicAccelerateIpEndpointRelationResponseBody) *DeleteBasicAccelerateIpEndpointRelationResponse
	GetBody() *DeleteBasicAccelerateIpEndpointRelationResponseBody
}

type DeleteBasicAccelerateIpEndpointRelationResponse struct {
	Headers    map[string]*string                                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteBasicAccelerateIpEndpointRelationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteBasicAccelerateIpEndpointRelationResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteBasicAccelerateIpEndpointRelationResponse) GoString() string {
	return s.String()
}

func (s *DeleteBasicAccelerateIpEndpointRelationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteBasicAccelerateIpEndpointRelationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteBasicAccelerateIpEndpointRelationResponse) GetBody() *DeleteBasicAccelerateIpEndpointRelationResponseBody {
	return s.Body
}

func (s *DeleteBasicAccelerateIpEndpointRelationResponse) SetHeaders(v map[string]*string) *DeleteBasicAccelerateIpEndpointRelationResponse {
	s.Headers = v
	return s
}

func (s *DeleteBasicAccelerateIpEndpointRelationResponse) SetStatusCode(v int32) *DeleteBasicAccelerateIpEndpointRelationResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteBasicAccelerateIpEndpointRelationResponse) SetBody(v *DeleteBasicAccelerateIpEndpointRelationResponseBody) *DeleteBasicAccelerateIpEndpointRelationResponse {
	s.Body = v
	return s
}

func (s *DeleteBasicAccelerateIpEndpointRelationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
