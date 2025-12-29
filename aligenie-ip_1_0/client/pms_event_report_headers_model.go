// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPmsEventReportHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *PmsEventReportHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsAligenieAccessToken(v string) *PmsEventReportHeaders
	GetXAcsAligenieAccessToken() *string
	SetAuthorization(v string) *PmsEventReportHeaders
	GetAuthorization() *string
}

type PmsEventReportHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s PmsEventReportHeaders) String() string {
	return dara.Prettify(s)
}

func (s PmsEventReportHeaders) GoString() string {
	return s.String()
}

func (s *PmsEventReportHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *PmsEventReportHeaders) GetXAcsAligenieAccessToken() *string {
	return s.XAcsAligenieAccessToken
}

func (s *PmsEventReportHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *PmsEventReportHeaders) SetCommonHeaders(v map[string]*string) *PmsEventReportHeaders {
	s.CommonHeaders = v
	return s
}

func (s *PmsEventReportHeaders) SetXAcsAligenieAccessToken(v string) *PmsEventReportHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *PmsEventReportHeaders) SetAuthorization(v string) *PmsEventReportHeaders {
	s.Authorization = &v
	return s
}

func (s *PmsEventReportHeaders) Validate() error {
	return dara.Validate(s)
}
