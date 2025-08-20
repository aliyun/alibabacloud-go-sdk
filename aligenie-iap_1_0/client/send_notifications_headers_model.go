// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendNotificationsHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *SendNotificationsHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsAligenieAccessToken(v string) *SendNotificationsHeaders
	GetXAcsAligenieAccessToken() *string
	SetAuthorization(v string) *SendNotificationsHeaders
	GetAuthorization() *string
}

type SendNotificationsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s SendNotificationsHeaders) String() string {
	return dara.Prettify(s)
}

func (s SendNotificationsHeaders) GoString() string {
	return s.String()
}

func (s *SendNotificationsHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *SendNotificationsHeaders) GetXAcsAligenieAccessToken() *string {
	return s.XAcsAligenieAccessToken
}

func (s *SendNotificationsHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *SendNotificationsHeaders) SetCommonHeaders(v map[string]*string) *SendNotificationsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SendNotificationsHeaders) SetXAcsAligenieAccessToken(v string) *SendNotificationsHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *SendNotificationsHeaders) SetAuthorization(v string) *SendNotificationsHeaders {
	s.Authorization = &v
	return s
}

func (s *SendNotificationsHeaders) Validate() error {
	return dara.Validate(s)
}
