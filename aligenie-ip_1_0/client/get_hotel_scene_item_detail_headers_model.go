// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHotelSceneItemDetailHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetHotelSceneItemDetailHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsAligenieAccessToken(v string) *GetHotelSceneItemDetailHeaders
	GetXAcsAligenieAccessToken() *string
	SetAuthorization(v string) *GetHotelSceneItemDetailHeaders
	GetAuthorization() *string
}

type GetHotelSceneItemDetailHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s GetHotelSceneItemDetailHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetHotelSceneItemDetailHeaders) GoString() string {
	return s.String()
}

func (s *GetHotelSceneItemDetailHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetHotelSceneItemDetailHeaders) GetXAcsAligenieAccessToken() *string {
	return s.XAcsAligenieAccessToken
}

func (s *GetHotelSceneItemDetailHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *GetHotelSceneItemDetailHeaders) SetCommonHeaders(v map[string]*string) *GetHotelSceneItemDetailHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetHotelSceneItemDetailHeaders) SetXAcsAligenieAccessToken(v string) *GetHotelSceneItemDetailHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *GetHotelSceneItemDetailHeaders) SetAuthorization(v string) *GetHotelSceneItemDetailHeaders {
	s.Authorization = &v
	return s
}

func (s *GetHotelSceneItemDetailHeaders) Validate() error {
	return dara.Validate(s)
}
