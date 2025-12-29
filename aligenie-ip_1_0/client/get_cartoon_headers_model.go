// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCartoonHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetCartoonHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsAligenieAccessToken(v string) *GetCartoonHeaders
	GetXAcsAligenieAccessToken() *string
	SetAuthorization(v string) *GetCartoonHeaders
	GetAuthorization() *string
}

type GetCartoonHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s GetCartoonHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetCartoonHeaders) GoString() string {
	return s.String()
}

func (s *GetCartoonHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetCartoonHeaders) GetXAcsAligenieAccessToken() *string {
	return s.XAcsAligenieAccessToken
}

func (s *GetCartoonHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *GetCartoonHeaders) SetCommonHeaders(v map[string]*string) *GetCartoonHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetCartoonHeaders) SetXAcsAligenieAccessToken(v string) *GetCartoonHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *GetCartoonHeaders) SetAuthorization(v string) *GetCartoonHeaders {
	s.Authorization = &v
	return s
}

func (s *GetCartoonHeaders) Validate() error {
	return dara.Validate(s)
}
