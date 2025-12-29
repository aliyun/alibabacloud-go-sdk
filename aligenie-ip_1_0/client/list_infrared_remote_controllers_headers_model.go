// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInfraredRemoteControllersHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *ListInfraredRemoteControllersHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsAligenieAccessToken(v string) *ListInfraredRemoteControllersHeaders
	GetXAcsAligenieAccessToken() *string
	SetAuthorization(v string) *ListInfraredRemoteControllersHeaders
	GetAuthorization() *string
}

type ListInfraredRemoteControllersHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s ListInfraredRemoteControllersHeaders) String() string {
	return dara.Prettify(s)
}

func (s ListInfraredRemoteControllersHeaders) GoString() string {
	return s.String()
}

func (s *ListInfraredRemoteControllersHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *ListInfraredRemoteControllersHeaders) GetXAcsAligenieAccessToken() *string {
	return s.XAcsAligenieAccessToken
}

func (s *ListInfraredRemoteControllersHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *ListInfraredRemoteControllersHeaders) SetCommonHeaders(v map[string]*string) *ListInfraredRemoteControllersHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListInfraredRemoteControllersHeaders) SetXAcsAligenieAccessToken(v string) *ListInfraredRemoteControllersHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *ListInfraredRemoteControllersHeaders) SetAuthorization(v string) *ListInfraredRemoteControllersHeaders {
	s.Authorization = &v
	return s
}

func (s *ListInfraredRemoteControllersHeaders) Validate() error {
	return dara.Validate(s)
}
