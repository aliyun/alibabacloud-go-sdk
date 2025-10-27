// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePrivateDnsEndpointResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreatePrivateDnsEndpointResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreatePrivateDnsEndpointResponse
	GetStatusCode() *int32
	SetBody(v *CreatePrivateDnsEndpointResponseBody) *CreatePrivateDnsEndpointResponse
	GetBody() *CreatePrivateDnsEndpointResponseBody
}

type CreatePrivateDnsEndpointResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreatePrivateDnsEndpointResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePrivateDnsEndpointResponse) String() string {
	return dara.Prettify(s)
}

func (s CreatePrivateDnsEndpointResponse) GoString() string {
	return s.String()
}

func (s *CreatePrivateDnsEndpointResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreatePrivateDnsEndpointResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreatePrivateDnsEndpointResponse) GetBody() *CreatePrivateDnsEndpointResponseBody {
	return s.Body
}

func (s *CreatePrivateDnsEndpointResponse) SetHeaders(v map[string]*string) *CreatePrivateDnsEndpointResponse {
	s.Headers = v
	return s
}

func (s *CreatePrivateDnsEndpointResponse) SetStatusCode(v int32) *CreatePrivateDnsEndpointResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePrivateDnsEndpointResponse) SetBody(v *CreatePrivateDnsEndpointResponseBody) *CreatePrivateDnsEndpointResponse {
	s.Body = v
	return s
}

func (s *CreatePrivateDnsEndpointResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
