// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCustomRoutingEndpointsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateCustomRoutingEndpointsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateCustomRoutingEndpointsResponse
	GetStatusCode() *int32
	SetBody(v *UpdateCustomRoutingEndpointsResponseBody) *UpdateCustomRoutingEndpointsResponse
	GetBody() *UpdateCustomRoutingEndpointsResponseBody
}

type UpdateCustomRoutingEndpointsResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateCustomRoutingEndpointsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateCustomRoutingEndpointsResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateCustomRoutingEndpointsResponse) GoString() string {
	return s.String()
}

func (s *UpdateCustomRoutingEndpointsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateCustomRoutingEndpointsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateCustomRoutingEndpointsResponse) GetBody() *UpdateCustomRoutingEndpointsResponseBody {
	return s.Body
}

func (s *UpdateCustomRoutingEndpointsResponse) SetHeaders(v map[string]*string) *UpdateCustomRoutingEndpointsResponse {
	s.Headers = v
	return s
}

func (s *UpdateCustomRoutingEndpointsResponse) SetStatusCode(v int32) *UpdateCustomRoutingEndpointsResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateCustomRoutingEndpointsResponse) SetBody(v *UpdateCustomRoutingEndpointsResponseBody) *UpdateCustomRoutingEndpointsResponse {
	s.Body = v
	return s
}

func (s *UpdateCustomRoutingEndpointsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
