// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRcuSceneHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *DeleteRcuSceneHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsAligenieAccessToken(v string) *DeleteRcuSceneHeaders
	GetXAcsAligenieAccessToken() *string
	SetAuthorization(v string) *DeleteRcuSceneHeaders
	GetAuthorization() *string
}

type DeleteRcuSceneHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s DeleteRcuSceneHeaders) String() string {
	return dara.Prettify(s)
}

func (s DeleteRcuSceneHeaders) GoString() string {
	return s.String()
}

func (s *DeleteRcuSceneHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *DeleteRcuSceneHeaders) GetXAcsAligenieAccessToken() *string {
	return s.XAcsAligenieAccessToken
}

func (s *DeleteRcuSceneHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *DeleteRcuSceneHeaders) SetCommonHeaders(v map[string]*string) *DeleteRcuSceneHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteRcuSceneHeaders) SetXAcsAligenieAccessToken(v string) *DeleteRcuSceneHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *DeleteRcuSceneHeaders) SetAuthorization(v string) *DeleteRcuSceneHeaders {
	s.Authorization = &v
	return s
}

func (s *DeleteRcuSceneHeaders) Validate() error {
	return dara.Validate(s)
}
