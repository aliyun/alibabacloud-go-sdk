// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserMFAAuthenticationSettingsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetUserMFAAuthenticationSettingsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetUserMFAAuthenticationSettingsResponse
	GetStatusCode() *int32
	SetBody(v *GetUserMFAAuthenticationSettingsResponseBody) *GetUserMFAAuthenticationSettingsResponse
	GetBody() *GetUserMFAAuthenticationSettingsResponseBody
}

type GetUserMFAAuthenticationSettingsResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetUserMFAAuthenticationSettingsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetUserMFAAuthenticationSettingsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetUserMFAAuthenticationSettingsResponse) GoString() string {
	return s.String()
}

func (s *GetUserMFAAuthenticationSettingsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetUserMFAAuthenticationSettingsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetUserMFAAuthenticationSettingsResponse) GetBody() *GetUserMFAAuthenticationSettingsResponseBody {
	return s.Body
}

func (s *GetUserMFAAuthenticationSettingsResponse) SetHeaders(v map[string]*string) *GetUserMFAAuthenticationSettingsResponse {
	s.Headers = v
	return s
}

func (s *GetUserMFAAuthenticationSettingsResponse) SetStatusCode(v int32) *GetUserMFAAuthenticationSettingsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUserMFAAuthenticationSettingsResponse) SetBody(v *GetUserMFAAuthenticationSettingsResponseBody) *GetUserMFAAuthenticationSettingsResponse {
	s.Body = v
	return s
}

func (s *GetUserMFAAuthenticationSettingsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
