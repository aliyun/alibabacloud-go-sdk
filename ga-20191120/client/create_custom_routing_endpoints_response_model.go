// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCustomRoutingEndpointsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateCustomRoutingEndpointsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateCustomRoutingEndpointsResponse
	GetStatusCode() *int32
	SetBody(v *CreateCustomRoutingEndpointsResponseBody) *CreateCustomRoutingEndpointsResponse
	GetBody() *CreateCustomRoutingEndpointsResponseBody
}

type CreateCustomRoutingEndpointsResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateCustomRoutingEndpointsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateCustomRoutingEndpointsResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomRoutingEndpointsResponse) GoString() string {
	return s.String()
}

func (s *CreateCustomRoutingEndpointsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateCustomRoutingEndpointsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateCustomRoutingEndpointsResponse) GetBody() *CreateCustomRoutingEndpointsResponseBody {
	return s.Body
}

func (s *CreateCustomRoutingEndpointsResponse) SetHeaders(v map[string]*string) *CreateCustomRoutingEndpointsResponse {
	s.Headers = v
	return s
}

func (s *CreateCustomRoutingEndpointsResponse) SetStatusCode(v int32) *CreateCustomRoutingEndpointsResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCustomRoutingEndpointsResponse) SetBody(v *CreateCustomRoutingEndpointsResponseBody) *CreateCustomRoutingEndpointsResponse {
	s.Body = v
	return s
}

func (s *CreateCustomRoutingEndpointsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
