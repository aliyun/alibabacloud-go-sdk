// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestartSupabaseProjectResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RestartSupabaseProjectResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RestartSupabaseProjectResponse
	GetStatusCode() *int32
	SetBody(v *RestartSupabaseProjectResponseBody) *RestartSupabaseProjectResponse
	GetBody() *RestartSupabaseProjectResponseBody
}

type RestartSupabaseProjectResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RestartSupabaseProjectResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RestartSupabaseProjectResponse) String() string {
	return dara.Prettify(s)
}

func (s RestartSupabaseProjectResponse) GoString() string {
	return s.String()
}

func (s *RestartSupabaseProjectResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RestartSupabaseProjectResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RestartSupabaseProjectResponse) GetBody() *RestartSupabaseProjectResponseBody {
	return s.Body
}

func (s *RestartSupabaseProjectResponse) SetHeaders(v map[string]*string) *RestartSupabaseProjectResponse {
	s.Headers = v
	return s
}

func (s *RestartSupabaseProjectResponse) SetStatusCode(v int32) *RestartSupabaseProjectResponse {
	s.StatusCode = &v
	return s
}

func (s *RestartSupabaseProjectResponse) SetBody(v *RestartSupabaseProjectResponseBody) *RestartSupabaseProjectResponse {
	s.Body = v
	return s
}

func (s *RestartSupabaseProjectResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
