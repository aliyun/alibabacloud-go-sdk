// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRcuSceneHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *UpdateRcuSceneHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsAligenieAccessToken(v string) *UpdateRcuSceneHeaders
	GetXAcsAligenieAccessToken() *string
	SetAuthorization(v string) *UpdateRcuSceneHeaders
	GetAuthorization() *string
}

type UpdateRcuSceneHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s UpdateRcuSceneHeaders) String() string {
	return dara.Prettify(s)
}

func (s UpdateRcuSceneHeaders) GoString() string {
	return s.String()
}

func (s *UpdateRcuSceneHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *UpdateRcuSceneHeaders) GetXAcsAligenieAccessToken() *string {
	return s.XAcsAligenieAccessToken
}

func (s *UpdateRcuSceneHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *UpdateRcuSceneHeaders) SetCommonHeaders(v map[string]*string) *UpdateRcuSceneHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateRcuSceneHeaders) SetXAcsAligenieAccessToken(v string) *UpdateRcuSceneHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *UpdateRcuSceneHeaders) SetAuthorization(v string) *UpdateRcuSceneHeaders {
	s.Authorization = &v
	return s
}

func (s *UpdateRcuSceneHeaders) Validate() error {
	return dara.Validate(s)
}
