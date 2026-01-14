// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBasicEndpointResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateBasicEndpointResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateBasicEndpointResponse
	GetStatusCode() *int32
	SetBody(v *CreateBasicEndpointResponseBody) *CreateBasicEndpointResponse
	GetBody() *CreateBasicEndpointResponseBody
}

type CreateBasicEndpointResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateBasicEndpointResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateBasicEndpointResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateBasicEndpointResponse) GoString() string {
	return s.String()
}

func (s *CreateBasicEndpointResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateBasicEndpointResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateBasicEndpointResponse) GetBody() *CreateBasicEndpointResponseBody {
	return s.Body
}

func (s *CreateBasicEndpointResponse) SetHeaders(v map[string]*string) *CreateBasicEndpointResponse {
	s.Headers = v
	return s
}

func (s *CreateBasicEndpointResponse) SetStatusCode(v int32) *CreateBasicEndpointResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateBasicEndpointResponse) SetBody(v *CreateBasicEndpointResponseBody) *CreateBasicEndpointResponse {
	s.Body = v
	return s
}

func (s *CreateBasicEndpointResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
