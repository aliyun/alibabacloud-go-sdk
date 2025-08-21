// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetJiangSuTelecomDataHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetJiangSuTelecomDataHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsAligenieAccessToken(v string) *GetJiangSuTelecomDataHeaders
	GetXAcsAligenieAccessToken() *string
	SetAuthorization(v string) *GetJiangSuTelecomDataHeaders
	GetAuthorization() *string
}

type GetJiangSuTelecomDataHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s GetJiangSuTelecomDataHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetJiangSuTelecomDataHeaders) GoString() string {
	return s.String()
}

func (s *GetJiangSuTelecomDataHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetJiangSuTelecomDataHeaders) GetXAcsAligenieAccessToken() *string {
	return s.XAcsAligenieAccessToken
}

func (s *GetJiangSuTelecomDataHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *GetJiangSuTelecomDataHeaders) SetCommonHeaders(v map[string]*string) *GetJiangSuTelecomDataHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetJiangSuTelecomDataHeaders) SetXAcsAligenieAccessToken(v string) *GetJiangSuTelecomDataHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *GetJiangSuTelecomDataHeaders) SetAuthorization(v string) *GetJiangSuTelecomDataHeaders {
	s.Authorization = &v
	return s
}

func (s *GetJiangSuTelecomDataHeaders) Validate() error {
	return dara.Validate(s)
}
