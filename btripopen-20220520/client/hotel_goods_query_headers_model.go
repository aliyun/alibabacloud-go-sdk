// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHotelGoodsQueryHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *HotelGoodsQueryHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripCorpToken(v string) *HotelGoodsQueryHeaders
	GetXAcsBtripCorpToken() *string
}

type HotelGoodsQueryHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// feth00jqwls
	XAcsBtripCorpToken *string `json:"x-acs-btrip-corp-token,omitempty" xml:"x-acs-btrip-corp-token,omitempty"`
}

func (s HotelGoodsQueryHeaders) String() string {
	return dara.Prettify(s)
}

func (s HotelGoodsQueryHeaders) GoString() string {
	return s.String()
}

func (s *HotelGoodsQueryHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *HotelGoodsQueryHeaders) GetXAcsBtripCorpToken() *string {
	return s.XAcsBtripCorpToken
}

func (s *HotelGoodsQueryHeaders) SetCommonHeaders(v map[string]*string) *HotelGoodsQueryHeaders {
	s.CommonHeaders = v
	return s
}

func (s *HotelGoodsQueryHeaders) SetXAcsBtripCorpToken(v string) *HotelGoodsQueryHeaders {
	s.XAcsBtripCorpToken = &v
	return s
}

func (s *HotelGoodsQueryHeaders) Validate() error {
	return dara.Validate(s)
}
