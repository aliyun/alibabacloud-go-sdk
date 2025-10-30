// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSupabaseProjectResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSupabaseProjectResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSupabaseProjectResponse
	GetStatusCode() *int32
	SetBody(v *GetSupabaseProjectResponseBody) *GetSupabaseProjectResponse
	GetBody() *GetSupabaseProjectResponseBody
}

type GetSupabaseProjectResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSupabaseProjectResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSupabaseProjectResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSupabaseProjectResponse) GoString() string {
	return s.String()
}

func (s *GetSupabaseProjectResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSupabaseProjectResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSupabaseProjectResponse) GetBody() *GetSupabaseProjectResponseBody {
	return s.Body
}

func (s *GetSupabaseProjectResponse) SetHeaders(v map[string]*string) *GetSupabaseProjectResponse {
	s.Headers = v
	return s
}

func (s *GetSupabaseProjectResponse) SetStatusCode(v int32) *GetSupabaseProjectResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSupabaseProjectResponse) SetBody(v *GetSupabaseProjectResponseBody) *GetSupabaseProjectResponse {
	s.Body = v
	return s
}

func (s *GetSupabaseProjectResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
