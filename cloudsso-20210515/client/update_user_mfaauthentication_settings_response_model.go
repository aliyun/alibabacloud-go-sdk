// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateUserMFAAuthenticationSettingsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateUserMFAAuthenticationSettingsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateUserMFAAuthenticationSettingsResponse
	GetStatusCode() *int32
	SetBody(v *UpdateUserMFAAuthenticationSettingsResponseBody) *UpdateUserMFAAuthenticationSettingsResponse
	GetBody() *UpdateUserMFAAuthenticationSettingsResponseBody
}

type UpdateUserMFAAuthenticationSettingsResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateUserMFAAuthenticationSettingsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateUserMFAAuthenticationSettingsResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateUserMFAAuthenticationSettingsResponse) GoString() string {
	return s.String()
}

func (s *UpdateUserMFAAuthenticationSettingsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateUserMFAAuthenticationSettingsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateUserMFAAuthenticationSettingsResponse) GetBody() *UpdateUserMFAAuthenticationSettingsResponseBody {
	return s.Body
}

func (s *UpdateUserMFAAuthenticationSettingsResponse) SetHeaders(v map[string]*string) *UpdateUserMFAAuthenticationSettingsResponse {
	s.Headers = v
	return s
}

func (s *UpdateUserMFAAuthenticationSettingsResponse) SetStatusCode(v int32) *UpdateUserMFAAuthenticationSettingsResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateUserMFAAuthenticationSettingsResponse) SetBody(v *UpdateUserMFAAuthenticationSettingsResponseBody) *UpdateUserMFAAuthenticationSettingsResponse {
	s.Body = v
	return s
}

func (s *UpdateUserMFAAuthenticationSettingsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
