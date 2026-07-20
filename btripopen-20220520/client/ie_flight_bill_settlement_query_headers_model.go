// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIeFlightBillSettlementQueryHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *IeFlightBillSettlementQueryHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripSoCorpToken(v string) *IeFlightBillSettlementQueryHeaders
	GetXAcsBtripSoCorpToken() *string
}

type IeFlightBillSettlementQueryHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// feth00jqwls
	XAcsBtripSoCorpToken *string `json:"x-acs-btrip-so-corp-token,omitempty" xml:"x-acs-btrip-so-corp-token,omitempty"`
}

func (s IeFlightBillSettlementQueryHeaders) String() string {
	return dara.Prettify(s)
}

func (s IeFlightBillSettlementQueryHeaders) GoString() string {
	return s.String()
}

func (s *IeFlightBillSettlementQueryHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *IeFlightBillSettlementQueryHeaders) GetXAcsBtripSoCorpToken() *string {
	return s.XAcsBtripSoCorpToken
}

func (s *IeFlightBillSettlementQueryHeaders) SetCommonHeaders(v map[string]*string) *IeFlightBillSettlementQueryHeaders {
	s.CommonHeaders = v
	return s
}

func (s *IeFlightBillSettlementQueryHeaders) SetXAcsBtripSoCorpToken(v string) *IeFlightBillSettlementQueryHeaders {
	s.XAcsBtripSoCorpToken = &v
	return s
}

func (s *IeFlightBillSettlementQueryHeaders) Validate() error {
	return dara.Validate(s)
}
