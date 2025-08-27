// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFlightBillSettlementQueryHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *FlightBillSettlementQueryHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripSoCorpToken(v string) *FlightBillSettlementQueryHeaders
	GetXAcsBtripSoCorpToken() *string
}

type FlightBillSettlementQueryHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// feth00jqwls
	XAcsBtripSoCorpToken *string `json:"x-acs-btrip-so-corp-token,omitempty" xml:"x-acs-btrip-so-corp-token,omitempty"`
}

func (s FlightBillSettlementQueryHeaders) String() string {
	return dara.Prettify(s)
}

func (s FlightBillSettlementQueryHeaders) GoString() string {
	return s.String()
}

func (s *FlightBillSettlementQueryHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *FlightBillSettlementQueryHeaders) GetXAcsBtripSoCorpToken() *string {
	return s.XAcsBtripSoCorpToken
}

func (s *FlightBillSettlementQueryHeaders) SetCommonHeaders(v map[string]*string) *FlightBillSettlementQueryHeaders {
	s.CommonHeaders = v
	return s
}

func (s *FlightBillSettlementQueryHeaders) SetXAcsBtripSoCorpToken(v string) *FlightBillSettlementQueryHeaders {
	s.XAcsBtripSoCorpToken = &v
	return s
}

func (s *FlightBillSettlementQueryHeaders) Validate() error {
	return dara.Validate(s)
}
