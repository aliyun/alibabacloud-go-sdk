// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUnionIdHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetUnionIdHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsAligenieAccessToken(v string) *GetUnionIdHeaders
	GetXAcsAligenieAccessToken() *string
	SetAuthorization(v string) *GetUnionIdHeaders
	GetAuthorization() *string
}

type GetUnionIdHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s GetUnionIdHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetUnionIdHeaders) GoString() string {
	return s.String()
}

func (s *GetUnionIdHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetUnionIdHeaders) GetXAcsAligenieAccessToken() *string {
	return s.XAcsAligenieAccessToken
}

func (s *GetUnionIdHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *GetUnionIdHeaders) SetCommonHeaders(v map[string]*string) *GetUnionIdHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetUnionIdHeaders) SetXAcsAligenieAccessToken(v string) *GetUnionIdHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *GetUnionIdHeaders) SetAuthorization(v string) *GetUnionIdHeaders {
	s.Authorization = &v
	return s
}

func (s *GetUnionIdHeaders) Validate() error {
	return dara.Validate(s)
}
