// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateHotelSceneItemHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *UpdateHotelSceneItemHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsAligenieAccessToken(v string) *UpdateHotelSceneItemHeaders
	GetXAcsAligenieAccessToken() *string
	SetAuthorization(v string) *UpdateHotelSceneItemHeaders
	GetAuthorization() *string
}

type UpdateHotelSceneItemHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s UpdateHotelSceneItemHeaders) String() string {
	return dara.Prettify(s)
}

func (s UpdateHotelSceneItemHeaders) GoString() string {
	return s.String()
}

func (s *UpdateHotelSceneItemHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *UpdateHotelSceneItemHeaders) GetXAcsAligenieAccessToken() *string {
	return s.XAcsAligenieAccessToken
}

func (s *UpdateHotelSceneItemHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *UpdateHotelSceneItemHeaders) SetCommonHeaders(v map[string]*string) *UpdateHotelSceneItemHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateHotelSceneItemHeaders) SetXAcsAligenieAccessToken(v string) *UpdateHotelSceneItemHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *UpdateHotelSceneItemHeaders) SetAuthorization(v string) *UpdateHotelSceneItemHeaders {
	s.Authorization = &v
	return s
}

func (s *UpdateHotelSceneItemHeaders) Validate() error {
	return dara.Validate(s)
}
