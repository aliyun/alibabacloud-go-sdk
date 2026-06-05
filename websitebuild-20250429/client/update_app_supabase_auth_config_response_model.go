// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAppSupabaseAuthConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateAppSupabaseAuthConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateAppSupabaseAuthConfigResponse
	GetStatusCode() *int32
	SetBody(v *UpdateAppSupabaseAuthConfigResponseBody) *UpdateAppSupabaseAuthConfigResponse
	GetBody() *UpdateAppSupabaseAuthConfigResponseBody
}

type UpdateAppSupabaseAuthConfigResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAppSupabaseAuthConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAppSupabaseAuthConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateAppSupabaseAuthConfigResponse) GoString() string {
	return s.String()
}

func (s *UpdateAppSupabaseAuthConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateAppSupabaseAuthConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateAppSupabaseAuthConfigResponse) GetBody() *UpdateAppSupabaseAuthConfigResponseBody {
	return s.Body
}

func (s *UpdateAppSupabaseAuthConfigResponse) SetHeaders(v map[string]*string) *UpdateAppSupabaseAuthConfigResponse {
	s.Headers = v
	return s
}

func (s *UpdateAppSupabaseAuthConfigResponse) SetStatusCode(v int32) *UpdateAppSupabaseAuthConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAppSupabaseAuthConfigResponse) SetBody(v *UpdateAppSupabaseAuthConfigResponseBody) *UpdateAppSupabaseAuthConfigResponse {
	s.Body = v
	return s
}

func (s *UpdateAppSupabaseAuthConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
