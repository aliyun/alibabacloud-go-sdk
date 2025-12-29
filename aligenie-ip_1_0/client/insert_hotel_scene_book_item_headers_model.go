// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInsertHotelSceneBookItemHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *InsertHotelSceneBookItemHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsAligenieAccessToken(v string) *InsertHotelSceneBookItemHeaders
	GetXAcsAligenieAccessToken() *string
	SetAuthorization(v string) *InsertHotelSceneBookItemHeaders
	GetAuthorization() *string
}

type InsertHotelSceneBookItemHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s InsertHotelSceneBookItemHeaders) String() string {
	return dara.Prettify(s)
}

func (s InsertHotelSceneBookItemHeaders) GoString() string {
	return s.String()
}

func (s *InsertHotelSceneBookItemHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *InsertHotelSceneBookItemHeaders) GetXAcsAligenieAccessToken() *string {
	return s.XAcsAligenieAccessToken
}

func (s *InsertHotelSceneBookItemHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *InsertHotelSceneBookItemHeaders) SetCommonHeaders(v map[string]*string) *InsertHotelSceneBookItemHeaders {
	s.CommonHeaders = v
	return s
}

func (s *InsertHotelSceneBookItemHeaders) SetXAcsAligenieAccessToken(v string) *InsertHotelSceneBookItemHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *InsertHotelSceneBookItemHeaders) SetAuthorization(v string) *InsertHotelSceneBookItemHeaders {
	s.Authorization = &v
	return s
}

func (s *InsertHotelSceneBookItemHeaders) Validate() error {
	return dara.Validate(s)
}
