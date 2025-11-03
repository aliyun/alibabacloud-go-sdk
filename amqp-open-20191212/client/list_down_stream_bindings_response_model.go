// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDownStreamBindingsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDownStreamBindingsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDownStreamBindingsResponse
	GetStatusCode() *int32
	SetBody(v *ListDownStreamBindingsResponseBody) *ListDownStreamBindingsResponse
	GetBody() *ListDownStreamBindingsResponseBody
}

type ListDownStreamBindingsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDownStreamBindingsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDownStreamBindingsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDownStreamBindingsResponse) GoString() string {
	return s.String()
}

func (s *ListDownStreamBindingsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDownStreamBindingsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDownStreamBindingsResponse) GetBody() *ListDownStreamBindingsResponseBody {
	return s.Body
}

func (s *ListDownStreamBindingsResponse) SetHeaders(v map[string]*string) *ListDownStreamBindingsResponse {
	s.Headers = v
	return s
}

func (s *ListDownStreamBindingsResponse) SetStatusCode(v int32) *ListDownStreamBindingsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDownStreamBindingsResponse) SetBody(v *ListDownStreamBindingsResponseBody) *ListDownStreamBindingsResponse {
	s.Body = v
	return s
}

func (s *ListDownStreamBindingsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
