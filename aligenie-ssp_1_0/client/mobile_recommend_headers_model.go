// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMobileRecommendHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *MobileRecommendHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsAligenieAccessToken(v string) *MobileRecommendHeaders
	GetXAcsAligenieAccessToken() *string
	SetAuthorization(v string) *MobileRecommendHeaders
	GetAuthorization() *string
}

type MobileRecommendHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s MobileRecommendHeaders) String() string {
	return dara.Prettify(s)
}

func (s MobileRecommendHeaders) GoString() string {
	return s.String()
}

func (s *MobileRecommendHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *MobileRecommendHeaders) GetXAcsAligenieAccessToken() *string {
	return s.XAcsAligenieAccessToken
}

func (s *MobileRecommendHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *MobileRecommendHeaders) SetCommonHeaders(v map[string]*string) *MobileRecommendHeaders {
	s.CommonHeaders = v
	return s
}

func (s *MobileRecommendHeaders) SetXAcsAligenieAccessToken(v string) *MobileRecommendHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *MobileRecommendHeaders) SetAuthorization(v string) *MobileRecommendHeaders {
	s.Authorization = &v
	return s
}

func (s *MobileRecommendHeaders) Validate() error {
	return dara.Validate(s)
}
