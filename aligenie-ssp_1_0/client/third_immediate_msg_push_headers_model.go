// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iThirdImmediateMsgPushHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *ThirdImmediateMsgPushHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsAligenieAccessToken(v string) *ThirdImmediateMsgPushHeaders
	GetXAcsAligenieAccessToken() *string
	SetAuthorization(v string) *ThirdImmediateMsgPushHeaders
	GetAuthorization() *string
}

type ThirdImmediateMsgPushHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s ThirdImmediateMsgPushHeaders) String() string {
	return dara.Prettify(s)
}

func (s ThirdImmediateMsgPushHeaders) GoString() string {
	return s.String()
}

func (s *ThirdImmediateMsgPushHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *ThirdImmediateMsgPushHeaders) GetXAcsAligenieAccessToken() *string {
	return s.XAcsAligenieAccessToken
}

func (s *ThirdImmediateMsgPushHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *ThirdImmediateMsgPushHeaders) SetCommonHeaders(v map[string]*string) *ThirdImmediateMsgPushHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ThirdImmediateMsgPushHeaders) SetXAcsAligenieAccessToken(v string) *ThirdImmediateMsgPushHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *ThirdImmediateMsgPushHeaders) SetAuthorization(v string) *ThirdImmediateMsgPushHeaders {
	s.Authorization = &v
	return s
}

func (s *ThirdImmediateMsgPushHeaders) Validate() error {
	return dara.Validate(s)
}
