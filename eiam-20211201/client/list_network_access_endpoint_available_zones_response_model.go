// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNetworkAccessEndpointAvailableZonesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListNetworkAccessEndpointAvailableZonesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListNetworkAccessEndpointAvailableZonesResponse
	GetStatusCode() *int32
	SetBody(v *ListNetworkAccessEndpointAvailableZonesResponseBody) *ListNetworkAccessEndpointAvailableZonesResponse
	GetBody() *ListNetworkAccessEndpointAvailableZonesResponseBody
}

type ListNetworkAccessEndpointAvailableZonesResponse struct {
	Headers    map[string]*string                                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListNetworkAccessEndpointAvailableZonesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListNetworkAccessEndpointAvailableZonesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListNetworkAccessEndpointAvailableZonesResponse) GoString() string {
	return s.String()
}

func (s *ListNetworkAccessEndpointAvailableZonesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListNetworkAccessEndpointAvailableZonesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListNetworkAccessEndpointAvailableZonesResponse) GetBody() *ListNetworkAccessEndpointAvailableZonesResponseBody {
	return s.Body
}

func (s *ListNetworkAccessEndpointAvailableZonesResponse) SetHeaders(v map[string]*string) *ListNetworkAccessEndpointAvailableZonesResponse {
	s.Headers = v
	return s
}

func (s *ListNetworkAccessEndpointAvailableZonesResponse) SetStatusCode(v int32) *ListNetworkAccessEndpointAvailableZonesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListNetworkAccessEndpointAvailableZonesResponse) SetBody(v *ListNetworkAccessEndpointAvailableZonesResponseBody) *ListNetworkAccessEndpointAvailableZonesResponse {
	s.Body = v
	return s
}

func (s *ListNetworkAccessEndpointAvailableZonesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
