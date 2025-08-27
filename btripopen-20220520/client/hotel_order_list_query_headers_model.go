// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHotelOrderListQueryHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *HotelOrderListQueryHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripSoCorpToken(v string) *HotelOrderListQueryHeaders
	GetXAcsBtripSoCorpToken() *string
}

type HotelOrderListQueryHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// feth00jqwls
	XAcsBtripSoCorpToken *string `json:"x-acs-btrip-so-corp-token,omitempty" xml:"x-acs-btrip-so-corp-token,omitempty"`
}

func (s HotelOrderListQueryHeaders) String() string {
	return dara.Prettify(s)
}

func (s HotelOrderListQueryHeaders) GoString() string {
	return s.String()
}

func (s *HotelOrderListQueryHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *HotelOrderListQueryHeaders) GetXAcsBtripSoCorpToken() *string {
	return s.XAcsBtripSoCorpToken
}

func (s *HotelOrderListQueryHeaders) SetCommonHeaders(v map[string]*string) *HotelOrderListQueryHeaders {
	s.CommonHeaders = v
	return s
}

func (s *HotelOrderListQueryHeaders) SetXAcsBtripSoCorpToken(v string) *HotelOrderListQueryHeaders {
	s.XAcsBtripSoCorpToken = &v
	return s
}

func (s *HotelOrderListQueryHeaders) Validate() error {
	return dara.Validate(s)
}
