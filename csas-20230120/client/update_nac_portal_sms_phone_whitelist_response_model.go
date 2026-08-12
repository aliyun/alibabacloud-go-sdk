// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateNacPortalSmsPhoneWhitelistResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateNacPortalSmsPhoneWhitelistResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateNacPortalSmsPhoneWhitelistResponse
	GetStatusCode() *int32
	SetBody(v *UpdateNacPortalSmsPhoneWhitelistResponseBody) *UpdateNacPortalSmsPhoneWhitelistResponse
	GetBody() *UpdateNacPortalSmsPhoneWhitelistResponseBody
}

type UpdateNacPortalSmsPhoneWhitelistResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateNacPortalSmsPhoneWhitelistResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateNacPortalSmsPhoneWhitelistResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateNacPortalSmsPhoneWhitelistResponse) GoString() string {
	return s.String()
}

func (s *UpdateNacPortalSmsPhoneWhitelistResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateNacPortalSmsPhoneWhitelistResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateNacPortalSmsPhoneWhitelistResponse) GetBody() *UpdateNacPortalSmsPhoneWhitelistResponseBody {
	return s.Body
}

func (s *UpdateNacPortalSmsPhoneWhitelistResponse) SetHeaders(v map[string]*string) *UpdateNacPortalSmsPhoneWhitelistResponse {
	s.Headers = v
	return s
}

func (s *UpdateNacPortalSmsPhoneWhitelistResponse) SetStatusCode(v int32) *UpdateNacPortalSmsPhoneWhitelistResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateNacPortalSmsPhoneWhitelistResponse) SetBody(v *UpdateNacPortalSmsPhoneWhitelistResponseBody) *UpdateNacPortalSmsPhoneWhitelistResponse {
	s.Body = v
	return s
}

func (s *UpdateNacPortalSmsPhoneWhitelistResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
