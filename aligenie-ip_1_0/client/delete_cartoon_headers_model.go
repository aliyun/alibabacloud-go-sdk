// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCartoonHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *DeleteCartoonHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsAligenieAccessToken(v string) *DeleteCartoonHeaders
	GetXAcsAligenieAccessToken() *string
	SetAuthorization(v string) *DeleteCartoonHeaders
	GetAuthorization() *string
}

type DeleteCartoonHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s DeleteCartoonHeaders) String() string {
	return dara.Prettify(s)
}

func (s DeleteCartoonHeaders) GoString() string {
	return s.String()
}

func (s *DeleteCartoonHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *DeleteCartoonHeaders) GetXAcsAligenieAccessToken() *string {
	return s.XAcsAligenieAccessToken
}

func (s *DeleteCartoonHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *DeleteCartoonHeaders) SetCommonHeaders(v map[string]*string) *DeleteCartoonHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteCartoonHeaders) SetXAcsAligenieAccessToken(v string) *DeleteCartoonHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *DeleteCartoonHeaders) SetAuthorization(v string) *DeleteCartoonHeaders {
	s.Authorization = &v
	return s
}

func (s *DeleteCartoonHeaders) Validate() error {
	return dara.Validate(s)
}
