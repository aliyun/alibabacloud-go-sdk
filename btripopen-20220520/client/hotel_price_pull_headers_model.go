// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHotelPricePullHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *HotelPricePullHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripCorpToken(v string) *HotelPricePullHeaders
	GetXAcsBtripCorpToken() *string
}

type HotelPricePullHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// feth00jqwls
	XAcsBtripCorpToken *string `json:"x-acs-btrip-corp-token,omitempty" xml:"x-acs-btrip-corp-token,omitempty"`
}

func (s HotelPricePullHeaders) String() string {
	return dara.Prettify(s)
}

func (s HotelPricePullHeaders) GoString() string {
	return s.String()
}

func (s *HotelPricePullHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *HotelPricePullHeaders) GetXAcsBtripCorpToken() *string {
	return s.XAcsBtripCorpToken
}

func (s *HotelPricePullHeaders) SetCommonHeaders(v map[string]*string) *HotelPricePullHeaders {
	s.CommonHeaders = v
	return s
}

func (s *HotelPricePullHeaders) SetXAcsBtripCorpToken(v string) *HotelPricePullHeaders {
	s.XAcsBtripCorpToken = &v
	return s
}

func (s *HotelPricePullHeaders) Validate() error {
	return dara.Validate(s)
}
