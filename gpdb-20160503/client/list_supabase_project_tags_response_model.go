// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSupabaseProjectTagsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListSupabaseProjectTagsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListSupabaseProjectTagsResponse
	GetStatusCode() *int32
	SetBody(v *ListSupabaseProjectTagsResponseBody) *ListSupabaseProjectTagsResponse
	GetBody() *ListSupabaseProjectTagsResponseBody
}

type ListSupabaseProjectTagsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSupabaseProjectTagsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSupabaseProjectTagsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListSupabaseProjectTagsResponse) GoString() string {
	return s.String()
}

func (s *ListSupabaseProjectTagsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListSupabaseProjectTagsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListSupabaseProjectTagsResponse) GetBody() *ListSupabaseProjectTagsResponseBody {
	return s.Body
}

func (s *ListSupabaseProjectTagsResponse) SetHeaders(v map[string]*string) *ListSupabaseProjectTagsResponse {
	s.Headers = v
	return s
}

func (s *ListSupabaseProjectTagsResponse) SetStatusCode(v int32) *ListSupabaseProjectTagsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSupabaseProjectTagsResponse) SetBody(v *ListSupabaseProjectTagsResponseBody) *ListSupabaseProjectTagsResponse {
	s.Body = v
	return s
}

func (s *ListSupabaseProjectTagsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
