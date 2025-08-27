// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIeHotelBillSettlementQueryHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *IeHotelBillSettlementQueryHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripSoCorpToken(v string) *IeHotelBillSettlementQueryHeaders
	GetXAcsBtripSoCorpToken() *string
}

type IeHotelBillSettlementQueryHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// feth00jqwls
	XAcsBtripSoCorpToken *string `json:"x-acs-btrip-so-corp-token,omitempty" xml:"x-acs-btrip-so-corp-token,omitempty"`
}

func (s IeHotelBillSettlementQueryHeaders) String() string {
	return dara.Prettify(s)
}

func (s IeHotelBillSettlementQueryHeaders) GoString() string {
	return s.String()
}

func (s *IeHotelBillSettlementQueryHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *IeHotelBillSettlementQueryHeaders) GetXAcsBtripSoCorpToken() *string {
	return s.XAcsBtripSoCorpToken
}

func (s *IeHotelBillSettlementQueryHeaders) SetCommonHeaders(v map[string]*string) *IeHotelBillSettlementQueryHeaders {
	s.CommonHeaders = v
	return s
}

func (s *IeHotelBillSettlementQueryHeaders) SetXAcsBtripSoCorpToken(v string) *IeHotelBillSettlementQueryHeaders {
	s.XAcsBtripSoCorpToken = &v
	return s
}

func (s *IeHotelBillSettlementQueryHeaders) Validate() error {
	return dara.Validate(s)
}
