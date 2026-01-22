// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetUserSsoSettingsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetUserSsoSettingsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetUserSsoSettingsResponse
	GetStatusCode() *int32
	SetBody(v *SetUserSsoSettingsResponseBody) *SetUserSsoSettingsResponse
	GetBody() *SetUserSsoSettingsResponseBody
}

type SetUserSsoSettingsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetUserSsoSettingsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetUserSsoSettingsResponse) String() string {
	return dara.Prettify(s)
}

func (s SetUserSsoSettingsResponse) GoString() string {
	return s.String()
}

func (s *SetUserSsoSettingsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetUserSsoSettingsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetUserSsoSettingsResponse) GetBody() *SetUserSsoSettingsResponseBody {
	return s.Body
}

func (s *SetUserSsoSettingsResponse) SetHeaders(v map[string]*string) *SetUserSsoSettingsResponse {
	s.Headers = v
	return s
}

func (s *SetUserSsoSettingsResponse) SetStatusCode(v int32) *SetUserSsoSettingsResponse {
	s.StatusCode = &v
	return s
}

func (s *SetUserSsoSettingsResponse) SetBody(v *SetUserSsoSettingsResponseBody) *SetUserSsoSettingsResponse {
	s.Body = v
	return s
}

func (s *SetUserSsoSettingsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
