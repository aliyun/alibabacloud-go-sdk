// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHotelBillSettlementQueryHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *HotelBillSettlementQueryHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripSoCorpToken(v string) *HotelBillSettlementQueryHeaders
	GetXAcsBtripSoCorpToken() *string
}

type HotelBillSettlementQueryHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// feth00jqwls
	XAcsBtripSoCorpToken *string `json:"x-acs-btrip-so-corp-token,omitempty" xml:"x-acs-btrip-so-corp-token,omitempty"`
}

func (s HotelBillSettlementQueryHeaders) String() string {
	return dara.Prettify(s)
}

func (s HotelBillSettlementQueryHeaders) GoString() string {
	return s.String()
}

func (s *HotelBillSettlementQueryHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *HotelBillSettlementQueryHeaders) GetXAcsBtripSoCorpToken() *string {
	return s.XAcsBtripSoCorpToken
}

func (s *HotelBillSettlementQueryHeaders) SetCommonHeaders(v map[string]*string) *HotelBillSettlementQueryHeaders {
	s.CommonHeaders = v
	return s
}

func (s *HotelBillSettlementQueryHeaders) SetXAcsBtripSoCorpToken(v string) *HotelBillSettlementQueryHeaders {
	s.XAcsBtripSoCorpToken = &v
	return s
}

func (s *HotelBillSettlementQueryHeaders) Validate() error {
	return dara.Validate(s)
}
