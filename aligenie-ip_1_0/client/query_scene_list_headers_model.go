// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySceneListHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *QuerySceneListHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsAligenieAccessToken(v string) *QuerySceneListHeaders
	GetXAcsAligenieAccessToken() *string
	SetAuthorization(v string) *QuerySceneListHeaders
	GetAuthorization() *string
}

type QuerySceneListHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s QuerySceneListHeaders) String() string {
	return dara.Prettify(s)
}

func (s QuerySceneListHeaders) GoString() string {
	return s.String()
}

func (s *QuerySceneListHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *QuerySceneListHeaders) GetXAcsAligenieAccessToken() *string {
	return s.XAcsAligenieAccessToken
}

func (s *QuerySceneListHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *QuerySceneListHeaders) SetCommonHeaders(v map[string]*string) *QuerySceneListHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QuerySceneListHeaders) SetXAcsAligenieAccessToken(v string) *QuerySceneListHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *QuerySceneListHeaders) SetAuthorization(v string) *QuerySceneListHeaders {
	s.Authorization = &v
	return s
}

func (s *QuerySceneListHeaders) Validate() error {
	return dara.Validate(s)
}
