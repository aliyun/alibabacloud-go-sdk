// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFuPointBillSettlementQueryHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *FuPointBillSettlementQueryHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripCorpToken(v string) *FuPointBillSettlementQueryHeaders
	GetXAcsBtripCorpToken() *string
}

type FuPointBillSettlementQueryHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// feth00jqwls
	XAcsBtripCorpToken *string `json:"x-acs-btrip-corp-token,omitempty" xml:"x-acs-btrip-corp-token,omitempty"`
}

func (s FuPointBillSettlementQueryHeaders) String() string {
	return dara.Prettify(s)
}

func (s FuPointBillSettlementQueryHeaders) GoString() string {
	return s.String()
}

func (s *FuPointBillSettlementQueryHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *FuPointBillSettlementQueryHeaders) GetXAcsBtripCorpToken() *string {
	return s.XAcsBtripCorpToken
}

func (s *FuPointBillSettlementQueryHeaders) SetCommonHeaders(v map[string]*string) *FuPointBillSettlementQueryHeaders {
	s.CommonHeaders = v
	return s
}

func (s *FuPointBillSettlementQueryHeaders) SetXAcsBtripCorpToken(v string) *FuPointBillSettlementQueryHeaders {
	s.XAcsBtripCorpToken = &v
	return s
}

func (s *FuPointBillSettlementQueryHeaders) Validate() error {
	return dara.Validate(s)
}
