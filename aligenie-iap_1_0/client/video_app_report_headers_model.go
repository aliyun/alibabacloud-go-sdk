// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVideoAppReportHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *VideoAppReportHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsAligenieAccessToken(v string) *VideoAppReportHeaders
	GetXAcsAligenieAccessToken() *string
	SetAuthorization(v string) *VideoAppReportHeaders
	GetAuthorization() *string
}

type VideoAppReportHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s VideoAppReportHeaders) String() string {
	return dara.Prettify(s)
}

func (s VideoAppReportHeaders) GoString() string {
	return s.String()
}

func (s *VideoAppReportHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *VideoAppReportHeaders) GetXAcsAligenieAccessToken() *string {
	return s.XAcsAligenieAccessToken
}

func (s *VideoAppReportHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *VideoAppReportHeaders) SetCommonHeaders(v map[string]*string) *VideoAppReportHeaders {
	s.CommonHeaders = v
	return s
}

func (s *VideoAppReportHeaders) SetXAcsAligenieAccessToken(v string) *VideoAppReportHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *VideoAppReportHeaders) SetAuthorization(v string) *VideoAppReportHeaders {
	s.Authorization = &v
	return s
}

func (s *VideoAppReportHeaders) Validate() error {
	return dara.Validate(s)
}
