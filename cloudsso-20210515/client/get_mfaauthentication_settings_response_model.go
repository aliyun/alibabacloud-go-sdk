// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMFAAuthenticationSettingsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetMFAAuthenticationSettingsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetMFAAuthenticationSettingsResponse
	GetStatusCode() *int32
	SetBody(v *GetMFAAuthenticationSettingsResponseBody) *GetMFAAuthenticationSettingsResponse
	GetBody() *GetMFAAuthenticationSettingsResponseBody
}

type GetMFAAuthenticationSettingsResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMFAAuthenticationSettingsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMFAAuthenticationSettingsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetMFAAuthenticationSettingsResponse) GoString() string {
	return s.String()
}

func (s *GetMFAAuthenticationSettingsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetMFAAuthenticationSettingsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetMFAAuthenticationSettingsResponse) GetBody() *GetMFAAuthenticationSettingsResponseBody {
	return s.Body
}

func (s *GetMFAAuthenticationSettingsResponse) SetHeaders(v map[string]*string) *GetMFAAuthenticationSettingsResponse {
	s.Headers = v
	return s
}

func (s *GetMFAAuthenticationSettingsResponse) SetStatusCode(v int32) *GetMFAAuthenticationSettingsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMFAAuthenticationSettingsResponse) SetBody(v *GetMFAAuthenticationSettingsResponseBody) *GetMFAAuthenticationSettingsResponse {
	s.Body = v
	return s
}

func (s *GetMFAAuthenticationSettingsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
