// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iProgressControlHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *ProgressControlHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsAligenieAccessToken(v string) *ProgressControlHeaders
	GetXAcsAligenieAccessToken() *string
	SetAuthorization(v string) *ProgressControlHeaders
	GetAuthorization() *string
}

type ProgressControlHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s ProgressControlHeaders) String() string {
	return dara.Prettify(s)
}

func (s ProgressControlHeaders) GoString() string {
	return s.String()
}

func (s *ProgressControlHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *ProgressControlHeaders) GetXAcsAligenieAccessToken() *string {
	return s.XAcsAligenieAccessToken
}

func (s *ProgressControlHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *ProgressControlHeaders) SetCommonHeaders(v map[string]*string) *ProgressControlHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ProgressControlHeaders) SetXAcsAligenieAccessToken(v string) *ProgressControlHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *ProgressControlHeaders) SetAuthorization(v string) *ProgressControlHeaders {
	s.Authorization = &v
	return s
}

func (s *ProgressControlHeaders) Validate() error {
	return dara.Validate(s)
}
