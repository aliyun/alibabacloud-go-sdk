// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateBasicEndpointResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateBasicEndpointResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateBasicEndpointResponse
	GetStatusCode() *int32
	SetBody(v *UpdateBasicEndpointResponseBody) *UpdateBasicEndpointResponse
	GetBody() *UpdateBasicEndpointResponseBody
}

type UpdateBasicEndpointResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateBasicEndpointResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateBasicEndpointResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateBasicEndpointResponse) GoString() string {
	return s.String()
}

func (s *UpdateBasicEndpointResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateBasicEndpointResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateBasicEndpointResponse) GetBody() *UpdateBasicEndpointResponseBody {
	return s.Body
}

func (s *UpdateBasicEndpointResponse) SetHeaders(v map[string]*string) *UpdateBasicEndpointResponse {
	s.Headers = v
	return s
}

func (s *UpdateBasicEndpointResponse) SetStatusCode(v int32) *UpdateBasicEndpointResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateBasicEndpointResponse) SetBody(v *UpdateBasicEndpointResponseBody) *UpdateBasicEndpointResponse {
	s.Body = v
	return s
}

func (s *UpdateBasicEndpointResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
