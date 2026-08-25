// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMFAAuthenticationSettingsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateMFAAuthenticationSettingsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateMFAAuthenticationSettingsResponse
	GetStatusCode() *int32
	SetBody(v *UpdateMFAAuthenticationSettingsResponseBody) *UpdateMFAAuthenticationSettingsResponse
	GetBody() *UpdateMFAAuthenticationSettingsResponseBody
}

type UpdateMFAAuthenticationSettingsResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateMFAAuthenticationSettingsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateMFAAuthenticationSettingsResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateMFAAuthenticationSettingsResponse) GoString() string {
	return s.String()
}

func (s *UpdateMFAAuthenticationSettingsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateMFAAuthenticationSettingsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateMFAAuthenticationSettingsResponse) GetBody() *UpdateMFAAuthenticationSettingsResponseBody {
	return s.Body
}

func (s *UpdateMFAAuthenticationSettingsResponse) SetHeaders(v map[string]*string) *UpdateMFAAuthenticationSettingsResponse {
	s.Headers = v
	return s
}

func (s *UpdateMFAAuthenticationSettingsResponse) SetStatusCode(v int32) *UpdateMFAAuthenticationSettingsResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateMFAAuthenticationSettingsResponse) SetBody(v *UpdateMFAAuthenticationSettingsResponseBody) *UpdateMFAAuthenticationSettingsResponse {
	s.Body = v
	return s
}

func (s *UpdateMFAAuthenticationSettingsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
