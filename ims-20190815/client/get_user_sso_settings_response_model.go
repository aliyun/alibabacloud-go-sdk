// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserSsoSettingsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetUserSsoSettingsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetUserSsoSettingsResponse
	GetStatusCode() *int32
	SetBody(v *GetUserSsoSettingsResponseBody) *GetUserSsoSettingsResponse
	GetBody() *GetUserSsoSettingsResponseBody
}

type GetUserSsoSettingsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetUserSsoSettingsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetUserSsoSettingsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetUserSsoSettingsResponse) GoString() string {
	return s.String()
}

func (s *GetUserSsoSettingsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetUserSsoSettingsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetUserSsoSettingsResponse) GetBody() *GetUserSsoSettingsResponseBody {
	return s.Body
}

func (s *GetUserSsoSettingsResponse) SetHeaders(v map[string]*string) *GetUserSsoSettingsResponse {
	s.Headers = v
	return s
}

func (s *GetUserSsoSettingsResponse) SetStatusCode(v int32) *GetUserSsoSettingsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUserSsoSettingsResponse) SetBody(v *GetUserSsoSettingsResponseBody) *GetUserSsoSettingsResponse {
	s.Body = v
	return s
}

func (s *GetUserSsoSettingsResponse) Validate() error {
	return dara.Validate(s)
}
