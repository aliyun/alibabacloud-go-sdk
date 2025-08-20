// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceEndpointResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetInstanceEndpointResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetInstanceEndpointResponse
	GetStatusCode() *int32
	SetBody(v *GetInstanceEndpointResponseBody) *GetInstanceEndpointResponse
	GetBody() *GetInstanceEndpointResponseBody
}

type GetInstanceEndpointResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetInstanceEndpointResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetInstanceEndpointResponse) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceEndpointResponse) GoString() string {
	return s.String()
}

func (s *GetInstanceEndpointResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetInstanceEndpointResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetInstanceEndpointResponse) GetBody() *GetInstanceEndpointResponseBody {
	return s.Body
}

func (s *GetInstanceEndpointResponse) SetHeaders(v map[string]*string) *GetInstanceEndpointResponse {
	s.Headers = v
	return s
}

func (s *GetInstanceEndpointResponse) SetStatusCode(v int32) *GetInstanceEndpointResponse {
	s.StatusCode = &v
	return s
}

func (s *GetInstanceEndpointResponse) SetBody(v *GetInstanceEndpointResponseBody) *GetInstanceEndpointResponse {
	s.Body = v
	return s
}

func (s *GetInstanceEndpointResponse) Validate() error {
	return dara.Validate(s)
}
