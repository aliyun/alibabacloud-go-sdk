// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHotelAskingPriceHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *HotelAskingPriceHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripCorpToken(v string) *HotelAskingPriceHeaders
	GetXAcsBtripCorpToken() *string
}

type HotelAskingPriceHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// feth00jqwls
	XAcsBtripCorpToken *string `json:"x-acs-btrip-corp-token,omitempty" xml:"x-acs-btrip-corp-token,omitempty"`
}

func (s HotelAskingPriceHeaders) String() string {
	return dara.Prettify(s)
}

func (s HotelAskingPriceHeaders) GoString() string {
	return s.String()
}

func (s *HotelAskingPriceHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *HotelAskingPriceHeaders) GetXAcsBtripCorpToken() *string {
	return s.XAcsBtripCorpToken
}

func (s *HotelAskingPriceHeaders) SetCommonHeaders(v map[string]*string) *HotelAskingPriceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *HotelAskingPriceHeaders) SetXAcsBtripCorpToken(v string) *HotelAskingPriceHeaders {
	s.XAcsBtripCorpToken = &v
	return s
}

func (s *HotelAskingPriceHeaders) Validate() error {
	return dara.Validate(s)
}
