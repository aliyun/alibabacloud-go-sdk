// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iScanCodeBindHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *ScanCodeBindHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsAligenieAccessToken(v string) *ScanCodeBindHeaders
	GetXAcsAligenieAccessToken() *string
	SetAuthorization(v string) *ScanCodeBindHeaders
	GetAuthorization() *string
}

type ScanCodeBindHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s ScanCodeBindHeaders) String() string {
	return dara.Prettify(s)
}

func (s ScanCodeBindHeaders) GoString() string {
	return s.String()
}

func (s *ScanCodeBindHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *ScanCodeBindHeaders) GetXAcsAligenieAccessToken() *string {
	return s.XAcsAligenieAccessToken
}

func (s *ScanCodeBindHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *ScanCodeBindHeaders) SetCommonHeaders(v map[string]*string) *ScanCodeBindHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ScanCodeBindHeaders) SetXAcsAligenieAccessToken(v string) *ScanCodeBindHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *ScanCodeBindHeaders) SetAuthorization(v string) *ScanCodeBindHeaders {
	s.Authorization = &v
	return s
}

func (s *ScanCodeBindHeaders) Validate() error {
	return dara.Validate(s)
}
