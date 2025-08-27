// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCooperatorHotelBillSettlementQueryHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *CooperatorHotelBillSettlementQueryHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripCorpToken(v string) *CooperatorHotelBillSettlementQueryHeaders
	GetXAcsBtripCorpToken() *string
}

type CooperatorHotelBillSettlementQueryHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// feth00jqwls
	XAcsBtripCorpToken *string `json:"x-acs-btrip-corp-token,omitempty" xml:"x-acs-btrip-corp-token,omitempty"`
}

func (s CooperatorHotelBillSettlementQueryHeaders) String() string {
	return dara.Prettify(s)
}

func (s CooperatorHotelBillSettlementQueryHeaders) GoString() string {
	return s.String()
}

func (s *CooperatorHotelBillSettlementQueryHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *CooperatorHotelBillSettlementQueryHeaders) GetXAcsBtripCorpToken() *string {
	return s.XAcsBtripCorpToken
}

func (s *CooperatorHotelBillSettlementQueryHeaders) SetCommonHeaders(v map[string]*string) *CooperatorHotelBillSettlementQueryHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CooperatorHotelBillSettlementQueryHeaders) SetXAcsBtripCorpToken(v string) *CooperatorHotelBillSettlementQueryHeaders {
	s.XAcsBtripCorpToken = &v
	return s
}

func (s *CooperatorHotelBillSettlementQueryHeaders) Validate() error {
	return dara.Validate(s)
}
