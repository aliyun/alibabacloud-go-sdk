// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iScgSearchHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *ScgSearchHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsAligenieAccessToken(v string) *ScgSearchHeaders
	GetXAcsAligenieAccessToken() *string
	SetAuthorization(v string) *ScgSearchHeaders
	GetAuthorization() *string
}

type ScgSearchHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s ScgSearchHeaders) String() string {
	return dara.Prettify(s)
}

func (s ScgSearchHeaders) GoString() string {
	return s.String()
}

func (s *ScgSearchHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *ScgSearchHeaders) GetXAcsAligenieAccessToken() *string {
	return s.XAcsAligenieAccessToken
}

func (s *ScgSearchHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *ScgSearchHeaders) SetCommonHeaders(v map[string]*string) *ScgSearchHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ScgSearchHeaders) SetXAcsAligenieAccessToken(v string) *ScgSearchHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *ScgSearchHeaders) SetAuthorization(v string) *ScgSearchHeaders {
	s.Authorization = &v
	return s
}

func (s *ScgSearchHeaders) Validate() error {
	return dara.Validate(s)
}
