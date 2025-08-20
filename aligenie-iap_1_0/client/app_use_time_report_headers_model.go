// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAppUseTimeReportHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *AppUseTimeReportHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsAligenieAccessToken(v string) *AppUseTimeReportHeaders
	GetXAcsAligenieAccessToken() *string
	SetAuthorization(v string) *AppUseTimeReportHeaders
	GetAuthorization() *string
}

type AppUseTimeReportHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s AppUseTimeReportHeaders) String() string {
	return dara.Prettify(s)
}

func (s AppUseTimeReportHeaders) GoString() string {
	return s.String()
}

func (s *AppUseTimeReportHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *AppUseTimeReportHeaders) GetXAcsAligenieAccessToken() *string {
	return s.XAcsAligenieAccessToken
}

func (s *AppUseTimeReportHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *AppUseTimeReportHeaders) SetCommonHeaders(v map[string]*string) *AppUseTimeReportHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AppUseTimeReportHeaders) SetXAcsAligenieAccessToken(v string) *AppUseTimeReportHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *AppUseTimeReportHeaders) SetAuthorization(v string) *AppUseTimeReportHeaders {
	s.Authorization = &v
	return s
}

func (s *AppUseTimeReportHeaders) Validate() error {
	return dara.Validate(s)
}
