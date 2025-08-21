// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePlayingListHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *CreatePlayingListHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsAligenieAccessToken(v string) *CreatePlayingListHeaders
	GetXAcsAligenieAccessToken() *string
	SetAuthorization(v string) *CreatePlayingListHeaders
	GetAuthorization() *string
}

type CreatePlayingListHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s CreatePlayingListHeaders) String() string {
	return dara.Prettify(s)
}

func (s CreatePlayingListHeaders) GoString() string {
	return s.String()
}

func (s *CreatePlayingListHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *CreatePlayingListHeaders) GetXAcsAligenieAccessToken() *string {
	return s.XAcsAligenieAccessToken
}

func (s *CreatePlayingListHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *CreatePlayingListHeaders) SetCommonHeaders(v map[string]*string) *CreatePlayingListHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreatePlayingListHeaders) SetXAcsAligenieAccessToken(v string) *CreatePlayingListHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *CreatePlayingListHeaders) SetAuthorization(v string) *CreatePlayingListHeaders {
	s.Authorization = &v
	return s
}

func (s *CreatePlayingListHeaders) Validate() error {
	return dara.Validate(s)
}
