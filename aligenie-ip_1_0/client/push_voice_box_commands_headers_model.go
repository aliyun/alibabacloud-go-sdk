// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPushVoiceBoxCommandsHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *PushVoiceBoxCommandsHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsAligenieAccessToken(v string) *PushVoiceBoxCommandsHeaders
	GetXAcsAligenieAccessToken() *string
	SetAuthorization(v string) *PushVoiceBoxCommandsHeaders
	GetAuthorization() *string
}

type PushVoiceBoxCommandsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s PushVoiceBoxCommandsHeaders) String() string {
	return dara.Prettify(s)
}

func (s PushVoiceBoxCommandsHeaders) GoString() string {
	return s.String()
}

func (s *PushVoiceBoxCommandsHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *PushVoiceBoxCommandsHeaders) GetXAcsAligenieAccessToken() *string {
	return s.XAcsAligenieAccessToken
}

func (s *PushVoiceBoxCommandsHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *PushVoiceBoxCommandsHeaders) SetCommonHeaders(v map[string]*string) *PushVoiceBoxCommandsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *PushVoiceBoxCommandsHeaders) SetXAcsAligenieAccessToken(v string) *PushVoiceBoxCommandsHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *PushVoiceBoxCommandsHeaders) SetAuthorization(v string) *PushVoiceBoxCommandsHeaders {
	s.Authorization = &v
	return s
}

func (s *PushVoiceBoxCommandsHeaders) Validate() error {
	return dara.Validate(s)
}
