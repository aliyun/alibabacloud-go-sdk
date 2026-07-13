// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateServiceEndpointResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateServiceEndpointResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateServiceEndpointResponse
	GetStatusCode() *int32
	SetBody(v *UpdateServiceEndpointResponseBody) *UpdateServiceEndpointResponse
	GetBody() *UpdateServiceEndpointResponseBody
}

type UpdateServiceEndpointResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateServiceEndpointResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateServiceEndpointResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateServiceEndpointResponse) GoString() string {
	return s.String()
}

func (s *UpdateServiceEndpointResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateServiceEndpointResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateServiceEndpointResponse) GetBody() *UpdateServiceEndpointResponseBody {
	return s.Body
}

func (s *UpdateServiceEndpointResponse) SetHeaders(v map[string]*string) *UpdateServiceEndpointResponse {
	s.Headers = v
	return s
}

func (s *UpdateServiceEndpointResponse) SetStatusCode(v int32) *UpdateServiceEndpointResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateServiceEndpointResponse) SetBody(v *UpdateServiceEndpointResponseBody) *UpdateServiceEndpointResponse {
	s.Body = v
	return s
}

func (s *UpdateServiceEndpointResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
