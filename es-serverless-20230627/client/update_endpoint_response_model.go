// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateEndpointResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateEndpointResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateEndpointResponse
	GetStatusCode() *int32
	SetBody(v *UpdateEndpointResponseBody) *UpdateEndpointResponse
	GetBody() *UpdateEndpointResponseBody
}

type UpdateEndpointResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateEndpointResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateEndpointResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateEndpointResponse) GoString() string {
	return s.String()
}

func (s *UpdateEndpointResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateEndpointResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateEndpointResponse) GetBody() *UpdateEndpointResponseBody {
	return s.Body
}

func (s *UpdateEndpointResponse) SetHeaders(v map[string]*string) *UpdateEndpointResponse {
	s.Headers = v
	return s
}

func (s *UpdateEndpointResponse) SetStatusCode(v int32) *UpdateEndpointResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateEndpointResponse) SetBody(v *UpdateEndpointResponseBody) *UpdateEndpointResponse {
	s.Body = v
	return s
}

func (s *UpdateEndpointResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
