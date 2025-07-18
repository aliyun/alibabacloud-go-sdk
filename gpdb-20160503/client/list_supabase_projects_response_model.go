// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSupabaseProjectsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListSupabaseProjectsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListSupabaseProjectsResponse
	GetStatusCode() *int32
	SetBody(v *ListSupabaseProjectsResponseBody) *ListSupabaseProjectsResponse
	GetBody() *ListSupabaseProjectsResponseBody
}

type ListSupabaseProjectsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSupabaseProjectsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSupabaseProjectsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListSupabaseProjectsResponse) GoString() string {
	return s.String()
}

func (s *ListSupabaseProjectsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListSupabaseProjectsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListSupabaseProjectsResponse) GetBody() *ListSupabaseProjectsResponseBody {
	return s.Body
}

func (s *ListSupabaseProjectsResponse) SetHeaders(v map[string]*string) *ListSupabaseProjectsResponse {
	s.Headers = v
	return s
}

func (s *ListSupabaseProjectsResponse) SetStatusCode(v int32) *ListSupabaseProjectsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSupabaseProjectsResponse) SetBody(v *ListSupabaseProjectsResponseBody) *ListSupabaseProjectsResponse {
	s.Body = v
	return s
}

func (s *ListSupabaseProjectsResponse) Validate() error {
	return dara.Validate(s)
}
