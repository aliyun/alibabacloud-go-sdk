// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTagSupabaseProjectResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *TagSupabaseProjectResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *TagSupabaseProjectResponse
	GetStatusCode() *int32
	SetBody(v *TagSupabaseProjectResponseBody) *TagSupabaseProjectResponse
	GetBody() *TagSupabaseProjectResponseBody
}

type TagSupabaseProjectResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TagSupabaseProjectResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TagSupabaseProjectResponse) String() string {
	return dara.Prettify(s)
}

func (s TagSupabaseProjectResponse) GoString() string {
	return s.String()
}

func (s *TagSupabaseProjectResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *TagSupabaseProjectResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *TagSupabaseProjectResponse) GetBody() *TagSupabaseProjectResponseBody {
	return s.Body
}

func (s *TagSupabaseProjectResponse) SetHeaders(v map[string]*string) *TagSupabaseProjectResponse {
	s.Headers = v
	return s
}

func (s *TagSupabaseProjectResponse) SetStatusCode(v int32) *TagSupabaseProjectResponse {
	s.StatusCode = &v
	return s
}

func (s *TagSupabaseProjectResponse) SetBody(v *TagSupabaseProjectResponseBody) *TagSupabaseProjectResponse {
	s.Body = v
	return s
}

func (s *TagSupabaseProjectResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
