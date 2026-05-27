// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUntagSupabaseProjectResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UntagSupabaseProjectResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UntagSupabaseProjectResponse
	GetStatusCode() *int32
	SetBody(v *UntagSupabaseProjectResponseBody) *UntagSupabaseProjectResponse
	GetBody() *UntagSupabaseProjectResponseBody
}

type UntagSupabaseProjectResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UntagSupabaseProjectResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UntagSupabaseProjectResponse) String() string {
	return dara.Prettify(s)
}

func (s UntagSupabaseProjectResponse) GoString() string {
	return s.String()
}

func (s *UntagSupabaseProjectResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UntagSupabaseProjectResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UntagSupabaseProjectResponse) GetBody() *UntagSupabaseProjectResponseBody {
	return s.Body
}

func (s *UntagSupabaseProjectResponse) SetHeaders(v map[string]*string) *UntagSupabaseProjectResponse {
	s.Headers = v
	return s
}

func (s *UntagSupabaseProjectResponse) SetStatusCode(v int32) *UntagSupabaseProjectResponse {
	s.StatusCode = &v
	return s
}

func (s *UntagSupabaseProjectResponse) SetBody(v *UntagSupabaseProjectResponseBody) *UntagSupabaseProjectResponse {
	s.Body = v
	return s
}

func (s *UntagSupabaseProjectResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
