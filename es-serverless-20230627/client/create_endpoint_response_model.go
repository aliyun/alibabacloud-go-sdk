// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEndpointResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateEndpointResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateEndpointResponse
	GetStatusCode() *int32
	SetBody(v *CreateEndpointResponseBody) *CreateEndpointResponse
	GetBody() *CreateEndpointResponseBody
}

type CreateEndpointResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateEndpointResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateEndpointResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateEndpointResponse) GoString() string {
	return s.String()
}

func (s *CreateEndpointResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateEndpointResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateEndpointResponse) GetBody() *CreateEndpointResponseBody {
	return s.Body
}

func (s *CreateEndpointResponse) SetHeaders(v map[string]*string) *CreateEndpointResponse {
	s.Headers = v
	return s
}

func (s *CreateEndpointResponse) SetStatusCode(v int32) *CreateEndpointResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateEndpointResponse) SetBody(v *CreateEndpointResponseBody) *CreateEndpointResponse {
	s.Body = v
	return s
}

func (s *CreateEndpointResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
