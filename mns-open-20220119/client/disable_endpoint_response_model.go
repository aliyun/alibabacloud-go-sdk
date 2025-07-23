// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableEndpointResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DisableEndpointResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DisableEndpointResponse
	GetStatusCode() *int32
	SetBody(v *DisableEndpointResponseBody) *DisableEndpointResponse
	GetBody() *DisableEndpointResponseBody
}

type DisableEndpointResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisableEndpointResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisableEndpointResponse) String() string {
	return dara.Prettify(s)
}

func (s DisableEndpointResponse) GoString() string {
	return s.String()
}

func (s *DisableEndpointResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DisableEndpointResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DisableEndpointResponse) GetBody() *DisableEndpointResponseBody {
	return s.Body
}

func (s *DisableEndpointResponse) SetHeaders(v map[string]*string) *DisableEndpointResponse {
	s.Headers = v
	return s
}

func (s *DisableEndpointResponse) SetStatusCode(v int32) *DisableEndpointResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableEndpointResponse) SetBody(v *DisableEndpointResponseBody) *DisableEndpointResponse {
	s.Body = v
	return s
}

func (s *DisableEndpointResponse) Validate() error {
	return dara.Validate(s)
}
