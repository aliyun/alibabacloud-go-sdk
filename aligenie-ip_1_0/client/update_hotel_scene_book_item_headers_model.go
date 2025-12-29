// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateHotelSceneBookItemHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *UpdateHotelSceneBookItemHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsAligenieAccessToken(v string) *UpdateHotelSceneBookItemHeaders
	GetXAcsAligenieAccessToken() *string
	SetAuthorization(v string) *UpdateHotelSceneBookItemHeaders
	GetAuthorization() *string
}

type UpdateHotelSceneBookItemHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s UpdateHotelSceneBookItemHeaders) String() string {
	return dara.Prettify(s)
}

func (s UpdateHotelSceneBookItemHeaders) GoString() string {
	return s.String()
}

func (s *UpdateHotelSceneBookItemHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *UpdateHotelSceneBookItemHeaders) GetXAcsAligenieAccessToken() *string {
	return s.XAcsAligenieAccessToken
}

func (s *UpdateHotelSceneBookItemHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *UpdateHotelSceneBookItemHeaders) SetCommonHeaders(v map[string]*string) *UpdateHotelSceneBookItemHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateHotelSceneBookItemHeaders) SetXAcsAligenieAccessToken(v string) *UpdateHotelSceneBookItemHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *UpdateHotelSceneBookItemHeaders) SetAuthorization(v string) *UpdateHotelSceneBookItemHeaders {
	s.Authorization = &v
	return s
}

func (s *UpdateHotelSceneBookItemHeaders) Validate() error {
	return dara.Validate(s)
}
