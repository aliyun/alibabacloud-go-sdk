// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSupabaseProjectResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateSupabaseProjectResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateSupabaseProjectResponse
	GetStatusCode() *int32
	SetBody(v *CreateSupabaseProjectResponseBody) *CreateSupabaseProjectResponse
	GetBody() *CreateSupabaseProjectResponseBody
}

type CreateSupabaseProjectResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateSupabaseProjectResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSupabaseProjectResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateSupabaseProjectResponse) GoString() string {
	return s.String()
}

func (s *CreateSupabaseProjectResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateSupabaseProjectResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateSupabaseProjectResponse) GetBody() *CreateSupabaseProjectResponseBody {
	return s.Body
}

func (s *CreateSupabaseProjectResponse) SetHeaders(v map[string]*string) *CreateSupabaseProjectResponse {
	s.Headers = v
	return s
}

func (s *CreateSupabaseProjectResponse) SetStatusCode(v int32) *CreateSupabaseProjectResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSupabaseProjectResponse) SetBody(v *CreateSupabaseProjectResponseBody) *CreateSupabaseProjectResponse {
	s.Body = v
	return s
}

func (s *CreateSupabaseProjectResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
