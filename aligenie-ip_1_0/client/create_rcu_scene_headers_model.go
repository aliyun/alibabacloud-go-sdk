// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRcuSceneHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *CreateRcuSceneHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsAligenieAccessToken(v string) *CreateRcuSceneHeaders
	GetXAcsAligenieAccessToken() *string
	SetAuthorization(v string) *CreateRcuSceneHeaders
	GetAuthorization() *string
}

type CreateRcuSceneHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s CreateRcuSceneHeaders) String() string {
	return dara.Prettify(s)
}

func (s CreateRcuSceneHeaders) GoString() string {
	return s.String()
}

func (s *CreateRcuSceneHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *CreateRcuSceneHeaders) GetXAcsAligenieAccessToken() *string {
	return s.XAcsAligenieAccessToken
}

func (s *CreateRcuSceneHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *CreateRcuSceneHeaders) SetCommonHeaders(v map[string]*string) *CreateRcuSceneHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateRcuSceneHeaders) SetXAcsAligenieAccessToken(v string) *CreateRcuSceneHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *CreateRcuSceneHeaders) SetAuthorization(v string) *CreateRcuSceneHeaders {
	s.Authorization = &v
	return s
}

func (s *CreateRcuSceneHeaders) Validate() error {
	return dara.Validate(s)
}
