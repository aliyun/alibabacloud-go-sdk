// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBasicAccelerateIpEndpointRelationsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListBasicAccelerateIpEndpointRelationsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListBasicAccelerateIpEndpointRelationsResponse
	GetStatusCode() *int32
	SetBody(v *ListBasicAccelerateIpEndpointRelationsResponseBody) *ListBasicAccelerateIpEndpointRelationsResponse
	GetBody() *ListBasicAccelerateIpEndpointRelationsResponseBody
}

type ListBasicAccelerateIpEndpointRelationsResponse struct {
	Headers    map[string]*string                                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListBasicAccelerateIpEndpointRelationsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListBasicAccelerateIpEndpointRelationsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListBasicAccelerateIpEndpointRelationsResponse) GoString() string {
	return s.String()
}

func (s *ListBasicAccelerateIpEndpointRelationsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListBasicAccelerateIpEndpointRelationsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListBasicAccelerateIpEndpointRelationsResponse) GetBody() *ListBasicAccelerateIpEndpointRelationsResponseBody {
	return s.Body
}

func (s *ListBasicAccelerateIpEndpointRelationsResponse) SetHeaders(v map[string]*string) *ListBasicAccelerateIpEndpointRelationsResponse {
	s.Headers = v
	return s
}

func (s *ListBasicAccelerateIpEndpointRelationsResponse) SetStatusCode(v int32) *ListBasicAccelerateIpEndpointRelationsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListBasicAccelerateIpEndpointRelationsResponse) SetBody(v *ListBasicAccelerateIpEndpointRelationsResponseBody) *ListBasicAccelerateIpEndpointRelationsResponse {
	s.Body = v
	return s
}

func (s *ListBasicAccelerateIpEndpointRelationsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
