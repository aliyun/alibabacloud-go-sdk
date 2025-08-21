// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudPlayerHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *CloudPlayerHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsAligenieAccessToken(v string) *CloudPlayerHeaders
	GetXAcsAligenieAccessToken() *string
	SetAuthorization(v string) *CloudPlayerHeaders
	GetAuthorization() *string
}

type CloudPlayerHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s CloudPlayerHeaders) String() string {
	return dara.Prettify(s)
}

func (s CloudPlayerHeaders) GoString() string {
	return s.String()
}

func (s *CloudPlayerHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *CloudPlayerHeaders) GetXAcsAligenieAccessToken() *string {
	return s.XAcsAligenieAccessToken
}

func (s *CloudPlayerHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *CloudPlayerHeaders) SetCommonHeaders(v map[string]*string) *CloudPlayerHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CloudPlayerHeaders) SetXAcsAligenieAccessToken(v string) *CloudPlayerHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *CloudPlayerHeaders) SetAuthorization(v string) *CloudPlayerHeaders {
	s.Authorization = &v
	return s
}

func (s *CloudPlayerHeaders) Validate() error {
	return dara.Validate(s)
}
