// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iWakeUpAppHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *WakeUpAppHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsAligenieAccessToken(v string) *WakeUpAppHeaders
	GetXAcsAligenieAccessToken() *string
	SetAuthorization(v string) *WakeUpAppHeaders
	GetAuthorization() *string
}

type WakeUpAppHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s WakeUpAppHeaders) String() string {
	return dara.Prettify(s)
}

func (s WakeUpAppHeaders) GoString() string {
	return s.String()
}

func (s *WakeUpAppHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *WakeUpAppHeaders) GetXAcsAligenieAccessToken() *string {
	return s.XAcsAligenieAccessToken
}

func (s *WakeUpAppHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *WakeUpAppHeaders) SetCommonHeaders(v map[string]*string) *WakeUpAppHeaders {
	s.CommonHeaders = v
	return s
}

func (s *WakeUpAppHeaders) SetXAcsAligenieAccessToken(v string) *WakeUpAppHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *WakeUpAppHeaders) SetAuthorization(v string) *WakeUpAppHeaders {
	s.Authorization = &v
	return s
}

func (s *WakeUpAppHeaders) Validate() error {
	return dara.Validate(s)
}
