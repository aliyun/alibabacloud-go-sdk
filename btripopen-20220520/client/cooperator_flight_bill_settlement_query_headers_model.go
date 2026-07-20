// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCooperatorFlightBillSettlementQueryHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *CooperatorFlightBillSettlementQueryHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripCorpToken(v string) *CooperatorFlightBillSettlementQueryHeaders
	GetXAcsBtripCorpToken() *string
}

type CooperatorFlightBillSettlementQueryHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// feth00jqwls
	XAcsBtripCorpToken *string `json:"x-acs-btrip-corp-token,omitempty" xml:"x-acs-btrip-corp-token,omitempty"`
}

func (s CooperatorFlightBillSettlementQueryHeaders) String() string {
	return dara.Prettify(s)
}

func (s CooperatorFlightBillSettlementQueryHeaders) GoString() string {
	return s.String()
}

func (s *CooperatorFlightBillSettlementQueryHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *CooperatorFlightBillSettlementQueryHeaders) GetXAcsBtripCorpToken() *string {
	return s.XAcsBtripCorpToken
}

func (s *CooperatorFlightBillSettlementQueryHeaders) SetCommonHeaders(v map[string]*string) *CooperatorFlightBillSettlementQueryHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CooperatorFlightBillSettlementQueryHeaders) SetXAcsBtripCorpToken(v string) *CooperatorFlightBillSettlementQueryHeaders {
	s.XAcsBtripCorpToken = &v
	return s
}

func (s *CooperatorFlightBillSettlementQueryHeaders) Validate() error {
	return dara.Validate(s)
}
