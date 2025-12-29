// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteHotelSceneBookItemHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *DeleteHotelSceneBookItemHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsAligenieAccessToken(v string) *DeleteHotelSceneBookItemHeaders
	GetXAcsAligenieAccessToken() *string
	SetAuthorization(v string) *DeleteHotelSceneBookItemHeaders
	GetAuthorization() *string
}

type DeleteHotelSceneBookItemHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s DeleteHotelSceneBookItemHeaders) String() string {
	return dara.Prettify(s)
}

func (s DeleteHotelSceneBookItemHeaders) GoString() string {
	return s.String()
}

func (s *DeleteHotelSceneBookItemHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *DeleteHotelSceneBookItemHeaders) GetXAcsAligenieAccessToken() *string {
	return s.XAcsAligenieAccessToken
}

func (s *DeleteHotelSceneBookItemHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *DeleteHotelSceneBookItemHeaders) SetCommonHeaders(v map[string]*string) *DeleteHotelSceneBookItemHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteHotelSceneBookItemHeaders) SetXAcsAligenieAccessToken(v string) *DeleteHotelSceneBookItemHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *DeleteHotelSceneBookItemHeaders) SetAuthorization(v string) *DeleteHotelSceneBookItemHeaders {
	s.Authorization = &v
	return s
}

func (s *DeleteHotelSceneBookItemHeaders) Validate() error {
	return dara.Validate(s)
}
