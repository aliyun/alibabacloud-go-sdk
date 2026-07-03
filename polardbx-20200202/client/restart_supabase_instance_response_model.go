// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestartSupabaseInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RestartSupabaseInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RestartSupabaseInstanceResponse
	GetStatusCode() *int32
	SetBody(v *RestartSupabaseInstanceResponseBody) *RestartSupabaseInstanceResponse
	GetBody() *RestartSupabaseInstanceResponseBody
}

type RestartSupabaseInstanceResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RestartSupabaseInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RestartSupabaseInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s RestartSupabaseInstanceResponse) GoString() string {
	return s.String()
}

func (s *RestartSupabaseInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RestartSupabaseInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RestartSupabaseInstanceResponse) GetBody() *RestartSupabaseInstanceResponseBody {
	return s.Body
}

func (s *RestartSupabaseInstanceResponse) SetHeaders(v map[string]*string) *RestartSupabaseInstanceResponse {
	s.Headers = v
	return s
}

func (s *RestartSupabaseInstanceResponse) SetStatusCode(v int32) *RestartSupabaseInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *RestartSupabaseInstanceResponse) SetBody(v *RestartSupabaseInstanceResponseBody) *RestartSupabaseInstanceResponse {
	s.Body = v
	return s
}

func (s *RestartSupabaseInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
