// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPushNotificationsHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *PushNotificationsHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsAligenieAccessToken(v string) *PushNotificationsHeaders
	GetXAcsAligenieAccessToken() *string
	SetAuthorization(v string) *PushNotificationsHeaders
	GetAuthorization() *string
}

type PushNotificationsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s PushNotificationsHeaders) String() string {
	return dara.Prettify(s)
}

func (s PushNotificationsHeaders) GoString() string {
	return s.String()
}

func (s *PushNotificationsHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *PushNotificationsHeaders) GetXAcsAligenieAccessToken() *string {
	return s.XAcsAligenieAccessToken
}

func (s *PushNotificationsHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *PushNotificationsHeaders) SetCommonHeaders(v map[string]*string) *PushNotificationsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *PushNotificationsHeaders) SetXAcsAligenieAccessToken(v string) *PushNotificationsHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *PushNotificationsHeaders) SetAuthorization(v string) *PushNotificationsHeaders {
	s.Authorization = &v
	return s
}

func (s *PushNotificationsHeaders) Validate() error {
	return dara.Validate(s)
}
