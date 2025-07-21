// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBindingsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListBindingsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListBindingsResponse
	GetStatusCode() *int32
	SetBody(v *ListBindingsResponseBody) *ListBindingsResponse
	GetBody() *ListBindingsResponseBody
}

type ListBindingsResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListBindingsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListBindingsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListBindingsResponse) GoString() string {
	return s.String()
}

func (s *ListBindingsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListBindingsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListBindingsResponse) GetBody() *ListBindingsResponseBody {
	return s.Body
}

func (s *ListBindingsResponse) SetHeaders(v map[string]*string) *ListBindingsResponse {
	s.Headers = v
	return s
}

func (s *ListBindingsResponse) SetStatusCode(v int32) *ListBindingsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListBindingsResponse) SetBody(v *ListBindingsResponseBody) *ListBindingsResponse {
	s.Body = v
	return s
}

func (s *ListBindingsResponse) Validate() error {
	return dara.Validate(s)
}
