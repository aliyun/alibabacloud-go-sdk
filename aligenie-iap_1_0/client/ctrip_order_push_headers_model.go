// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCtripOrderPushHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *CtripOrderPushHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsAligenieAccessToken(v string) *CtripOrderPushHeaders
	GetXAcsAligenieAccessToken() *string
	SetAuthorization(v string) *CtripOrderPushHeaders
	GetAuthorization() *string
}

type CtripOrderPushHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s CtripOrderPushHeaders) String() string {
	return dara.Prettify(s)
}

func (s CtripOrderPushHeaders) GoString() string {
	return s.String()
}

func (s *CtripOrderPushHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *CtripOrderPushHeaders) GetXAcsAligenieAccessToken() *string {
	return s.XAcsAligenieAccessToken
}

func (s *CtripOrderPushHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *CtripOrderPushHeaders) SetCommonHeaders(v map[string]*string) *CtripOrderPushHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CtripOrderPushHeaders) SetXAcsAligenieAccessToken(v string) *CtripOrderPushHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *CtripOrderPushHeaders) SetAuthorization(v string) *CtripOrderPushHeaders {
	s.Authorization = &v
	return s
}

func (s *CtripOrderPushHeaders) Validate() error {
	return dara.Validate(s)
}
