// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCustomEndpointResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateCustomEndpointResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateCustomEndpointResponse
	GetStatusCode() *int32
	SetBody(v *CreateCustomEndpointResponseBody) *CreateCustomEndpointResponse
	GetBody() *CreateCustomEndpointResponseBody
}

type CreateCustomEndpointResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateCustomEndpointResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateCustomEndpointResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomEndpointResponse) GoString() string {
	return s.String()
}

func (s *CreateCustomEndpointResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateCustomEndpointResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateCustomEndpointResponse) GetBody() *CreateCustomEndpointResponseBody {
	return s.Body
}

func (s *CreateCustomEndpointResponse) SetHeaders(v map[string]*string) *CreateCustomEndpointResponse {
	s.Headers = v
	return s
}

func (s *CreateCustomEndpointResponse) SetStatusCode(v int32) *CreateCustomEndpointResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCustomEndpointResponse) SetBody(v *CreateCustomEndpointResponseBody) *CreateCustomEndpointResponse {
	s.Body = v
	return s
}

func (s *CreateCustomEndpointResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
