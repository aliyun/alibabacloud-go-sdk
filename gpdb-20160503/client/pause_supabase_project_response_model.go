// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPauseSupabaseProjectResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PauseSupabaseProjectResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PauseSupabaseProjectResponse
	GetStatusCode() *int32
	SetBody(v *PauseSupabaseProjectResponseBody) *PauseSupabaseProjectResponse
	GetBody() *PauseSupabaseProjectResponseBody
}

type PauseSupabaseProjectResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PauseSupabaseProjectResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PauseSupabaseProjectResponse) String() string {
	return dara.Prettify(s)
}

func (s PauseSupabaseProjectResponse) GoString() string {
	return s.String()
}

func (s *PauseSupabaseProjectResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PauseSupabaseProjectResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PauseSupabaseProjectResponse) GetBody() *PauseSupabaseProjectResponseBody {
	return s.Body
}

func (s *PauseSupabaseProjectResponse) SetHeaders(v map[string]*string) *PauseSupabaseProjectResponse {
	s.Headers = v
	return s
}

func (s *PauseSupabaseProjectResponse) SetStatusCode(v int32) *PauseSupabaseProjectResponse {
	s.StatusCode = &v
	return s
}

func (s *PauseSupabaseProjectResponse) SetBody(v *PauseSupabaseProjectResponseBody) *PauseSupabaseProjectResponse {
	s.Body = v
	return s
}

func (s *PauseSupabaseProjectResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
