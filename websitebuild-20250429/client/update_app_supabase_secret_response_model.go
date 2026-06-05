// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAppSupabaseSecretResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateAppSupabaseSecretResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateAppSupabaseSecretResponse
	GetStatusCode() *int32
	SetBody(v *UpdateAppSupabaseSecretResponseBody) *UpdateAppSupabaseSecretResponse
	GetBody() *UpdateAppSupabaseSecretResponseBody
}

type UpdateAppSupabaseSecretResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAppSupabaseSecretResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAppSupabaseSecretResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateAppSupabaseSecretResponse) GoString() string {
	return s.String()
}

func (s *UpdateAppSupabaseSecretResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateAppSupabaseSecretResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateAppSupabaseSecretResponse) GetBody() *UpdateAppSupabaseSecretResponseBody {
	return s.Body
}

func (s *UpdateAppSupabaseSecretResponse) SetHeaders(v map[string]*string) *UpdateAppSupabaseSecretResponse {
	s.Headers = v
	return s
}

func (s *UpdateAppSupabaseSecretResponse) SetStatusCode(v int32) *UpdateAppSupabaseSecretResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAppSupabaseSecretResponse) SetBody(v *UpdateAppSupabaseSecretResponseBody) *UpdateAppSupabaseSecretResponse {
	s.Body = v
	return s
}

func (s *UpdateAppSupabaseSecretResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
