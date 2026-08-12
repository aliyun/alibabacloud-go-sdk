// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNacPortalSmsPhoneWhitelistResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetNacPortalSmsPhoneWhitelistResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetNacPortalSmsPhoneWhitelistResponse
	GetStatusCode() *int32
	SetBody(v *GetNacPortalSmsPhoneWhitelistResponseBody) *GetNacPortalSmsPhoneWhitelistResponse
	GetBody() *GetNacPortalSmsPhoneWhitelistResponseBody
}

type GetNacPortalSmsPhoneWhitelistResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetNacPortalSmsPhoneWhitelistResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetNacPortalSmsPhoneWhitelistResponse) String() string {
	return dara.Prettify(s)
}

func (s GetNacPortalSmsPhoneWhitelistResponse) GoString() string {
	return s.String()
}

func (s *GetNacPortalSmsPhoneWhitelistResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetNacPortalSmsPhoneWhitelistResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetNacPortalSmsPhoneWhitelistResponse) GetBody() *GetNacPortalSmsPhoneWhitelistResponseBody {
	return s.Body
}

func (s *GetNacPortalSmsPhoneWhitelistResponse) SetHeaders(v map[string]*string) *GetNacPortalSmsPhoneWhitelistResponse {
	s.Headers = v
	return s
}

func (s *GetNacPortalSmsPhoneWhitelistResponse) SetStatusCode(v int32) *GetNacPortalSmsPhoneWhitelistResponse {
	s.StatusCode = &v
	return s
}

func (s *GetNacPortalSmsPhoneWhitelistResponse) SetBody(v *GetNacPortalSmsPhoneWhitelistResponseBody) *GetNacPortalSmsPhoneWhitelistResponse {
	s.Body = v
	return s
}

func (s *GetNacPortalSmsPhoneWhitelistResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
