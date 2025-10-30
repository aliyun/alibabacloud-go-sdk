// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVpcEndpointZonesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListVpcEndpointZonesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListVpcEndpointZonesResponse
	GetStatusCode() *int32
	SetBody(v *ListVpcEndpointZonesResponseBody) *ListVpcEndpointZonesResponse
	GetBody() *ListVpcEndpointZonesResponseBody
}

type ListVpcEndpointZonesResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListVpcEndpointZonesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListVpcEndpointZonesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListVpcEndpointZonesResponse) GoString() string {
	return s.String()
}

func (s *ListVpcEndpointZonesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListVpcEndpointZonesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListVpcEndpointZonesResponse) GetBody() *ListVpcEndpointZonesResponseBody {
	return s.Body
}

func (s *ListVpcEndpointZonesResponse) SetHeaders(v map[string]*string) *ListVpcEndpointZonesResponse {
	s.Headers = v
	return s
}

func (s *ListVpcEndpointZonesResponse) SetStatusCode(v int32) *ListVpcEndpointZonesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListVpcEndpointZonesResponse) SetBody(v *ListVpcEndpointZonesResponseBody) *ListVpcEndpointZonesResponse {
	s.Body = v
	return s
}

func (s *ListVpcEndpointZonesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
