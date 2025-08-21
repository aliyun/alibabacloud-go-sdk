// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetContentHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetContentHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsAligenieAccessToken(v string) *GetContentHeaders
	GetXAcsAligenieAccessToken() *string
	SetAuthorization(v string) *GetContentHeaders
	GetAuthorization() *string
}

type GetContentHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s GetContentHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetContentHeaders) GoString() string {
	return s.String()
}

func (s *GetContentHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetContentHeaders) GetXAcsAligenieAccessToken() *string {
	return s.XAcsAligenieAccessToken
}

func (s *GetContentHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *GetContentHeaders) SetCommonHeaders(v map[string]*string) *GetContentHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetContentHeaders) SetXAcsAligenieAccessToken(v string) *GetContentHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *GetContentHeaders) SetAuthorization(v string) *GetContentHeaders {
	s.Authorization = &v
	return s
}

func (s *GetContentHeaders) Validate() error {
	return dara.Validate(s)
}
