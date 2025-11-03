// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateInstanceEndpointStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateInstanceEndpointStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateInstanceEndpointStatusResponse
	GetStatusCode() *int32
	SetBody(v *UpdateInstanceEndpointStatusResponseBody) *UpdateInstanceEndpointStatusResponse
	GetBody() *UpdateInstanceEndpointStatusResponseBody
}

type UpdateInstanceEndpointStatusResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateInstanceEndpointStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateInstanceEndpointStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateInstanceEndpointStatusResponse) GoString() string {
	return s.String()
}

func (s *UpdateInstanceEndpointStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateInstanceEndpointStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateInstanceEndpointStatusResponse) GetBody() *UpdateInstanceEndpointStatusResponseBody {
	return s.Body
}

func (s *UpdateInstanceEndpointStatusResponse) SetHeaders(v map[string]*string) *UpdateInstanceEndpointStatusResponse {
	s.Headers = v
	return s
}

func (s *UpdateInstanceEndpointStatusResponse) SetStatusCode(v int32) *UpdateInstanceEndpointStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateInstanceEndpointStatusResponse) SetBody(v *UpdateInstanceEndpointStatusResponseBody) *UpdateInstanceEndpointStatusResponse {
	s.Body = v
	return s
}

func (s *UpdateInstanceEndpointStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
