// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateServiceEndpointResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateServiceEndpointResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateServiceEndpointResponse
	GetStatusCode() *int32
	SetBody(v *CreateServiceEndpointResponseBody) *CreateServiceEndpointResponse
	GetBody() *CreateServiceEndpointResponseBody
}

type CreateServiceEndpointResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateServiceEndpointResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateServiceEndpointResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateServiceEndpointResponse) GoString() string {
	return s.String()
}

func (s *CreateServiceEndpointResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateServiceEndpointResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateServiceEndpointResponse) GetBody() *CreateServiceEndpointResponseBody {
	return s.Body
}

func (s *CreateServiceEndpointResponse) SetHeaders(v map[string]*string) *CreateServiceEndpointResponse {
	s.Headers = v
	return s
}

func (s *CreateServiceEndpointResponse) SetStatusCode(v int32) *CreateServiceEndpointResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateServiceEndpointResponse) SetBody(v *CreateServiceEndpointResponseBody) *CreateServiceEndpointResponse {
	s.Body = v
	return s
}

func (s *CreateServiceEndpointResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
