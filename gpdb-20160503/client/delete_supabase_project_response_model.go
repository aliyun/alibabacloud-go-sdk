// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSupabaseProjectResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteSupabaseProjectResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteSupabaseProjectResponse
	GetStatusCode() *int32
	SetBody(v *DeleteSupabaseProjectResponseBody) *DeleteSupabaseProjectResponse
	GetBody() *DeleteSupabaseProjectResponseBody
}

type DeleteSupabaseProjectResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteSupabaseProjectResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteSupabaseProjectResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteSupabaseProjectResponse) GoString() string {
	return s.String()
}

func (s *DeleteSupabaseProjectResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteSupabaseProjectResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteSupabaseProjectResponse) GetBody() *DeleteSupabaseProjectResponseBody {
	return s.Body
}

func (s *DeleteSupabaseProjectResponse) SetHeaders(v map[string]*string) *DeleteSupabaseProjectResponse {
	s.Headers = v
	return s
}

func (s *DeleteSupabaseProjectResponse) SetStatusCode(v int32) *DeleteSupabaseProjectResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSupabaseProjectResponse) SetBody(v *DeleteSupabaseProjectResponseBody) *DeleteSupabaseProjectResponse {
	s.Body = v
	return s
}

func (s *DeleteSupabaseProjectResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
