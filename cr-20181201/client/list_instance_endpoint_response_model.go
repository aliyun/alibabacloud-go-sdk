// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstanceEndpointResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListInstanceEndpointResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListInstanceEndpointResponse
	GetStatusCode() *int32
	SetBody(v *ListInstanceEndpointResponseBody) *ListInstanceEndpointResponse
	GetBody() *ListInstanceEndpointResponseBody
}

type ListInstanceEndpointResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListInstanceEndpointResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListInstanceEndpointResponse) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceEndpointResponse) GoString() string {
	return s.String()
}

func (s *ListInstanceEndpointResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListInstanceEndpointResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListInstanceEndpointResponse) GetBody() *ListInstanceEndpointResponseBody {
	return s.Body
}

func (s *ListInstanceEndpointResponse) SetHeaders(v map[string]*string) *ListInstanceEndpointResponse {
	s.Headers = v
	return s
}

func (s *ListInstanceEndpointResponse) SetStatusCode(v int32) *ListInstanceEndpointResponse {
	s.StatusCode = &v
	return s
}

func (s *ListInstanceEndpointResponse) SetBody(v *ListInstanceEndpointResponseBody) *ListInstanceEndpointResponse {
	s.Body = v
	return s
}

func (s *ListInstanceEndpointResponse) Validate() error {
	return dara.Validate(s)
}
