// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHotelCityCodeListHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *HotelCityCodeListHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripCorpToken(v string) *HotelCityCodeListHeaders
	GetXAcsBtripCorpToken() *string
}

type HotelCityCodeListHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// feth00jqwls
	XAcsBtripCorpToken *string `json:"x-acs-btrip-corp-token,omitempty" xml:"x-acs-btrip-corp-token,omitempty"`
}

func (s HotelCityCodeListHeaders) String() string {
	return dara.Prettify(s)
}

func (s HotelCityCodeListHeaders) GoString() string {
	return s.String()
}

func (s *HotelCityCodeListHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *HotelCityCodeListHeaders) GetXAcsBtripCorpToken() *string {
	return s.XAcsBtripCorpToken
}

func (s *HotelCityCodeListHeaders) SetCommonHeaders(v map[string]*string) *HotelCityCodeListHeaders {
	s.CommonHeaders = v
	return s
}

func (s *HotelCityCodeListHeaders) SetXAcsBtripCorpToken(v string) *HotelCityCodeListHeaders {
	s.XAcsBtripCorpToken = &v
	return s
}

func (s *HotelCityCodeListHeaders) Validate() error {
	return dara.Validate(s)
}
