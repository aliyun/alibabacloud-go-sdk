// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIeCarBillSettlementQueryHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *IeCarBillSettlementQueryHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripSoCorpToken(v string) *IeCarBillSettlementQueryHeaders
	GetXAcsBtripSoCorpToken() *string
}

type IeCarBillSettlementQueryHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// feth00jqwls
	XAcsBtripSoCorpToken *string `json:"x-acs-btrip-so-corp-token,omitempty" xml:"x-acs-btrip-so-corp-token,omitempty"`
}

func (s IeCarBillSettlementQueryHeaders) String() string {
	return dara.Prettify(s)
}

func (s IeCarBillSettlementQueryHeaders) GoString() string {
	return s.String()
}

func (s *IeCarBillSettlementQueryHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *IeCarBillSettlementQueryHeaders) GetXAcsBtripSoCorpToken() *string {
	return s.XAcsBtripSoCorpToken
}

func (s *IeCarBillSettlementQueryHeaders) SetCommonHeaders(v map[string]*string) *IeCarBillSettlementQueryHeaders {
	s.CommonHeaders = v
	return s
}

func (s *IeCarBillSettlementQueryHeaders) SetXAcsBtripSoCorpToken(v string) *IeCarBillSettlementQueryHeaders {
	s.XAcsBtripSoCorpToken = &v
	return s
}

func (s *IeCarBillSettlementQueryHeaders) Validate() error {
	return dara.Validate(s)
}
