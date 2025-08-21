// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSubHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *DeleteSubHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsAligenieAccessToken(v string) *DeleteSubHeaders
	GetXAcsAligenieAccessToken() *string
	SetAuthorization(v string) *DeleteSubHeaders
	GetAuthorization() *string
}

type DeleteSubHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s DeleteSubHeaders) String() string {
	return dara.Prettify(s)
}

func (s DeleteSubHeaders) GoString() string {
	return s.String()
}

func (s *DeleteSubHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *DeleteSubHeaders) GetXAcsAligenieAccessToken() *string {
	return s.XAcsAligenieAccessToken
}

func (s *DeleteSubHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *DeleteSubHeaders) SetCommonHeaders(v map[string]*string) *DeleteSubHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteSubHeaders) SetXAcsAligenieAccessToken(v string) *DeleteSubHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *DeleteSubHeaders) SetAuthorization(v string) *DeleteSubHeaders {
	s.Authorization = &v
	return s
}

func (s *DeleteSubHeaders) Validate() error {
	return dara.Validate(s)
}
