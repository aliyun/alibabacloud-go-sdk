// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCarBillSettlementQueryHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *CarBillSettlementQueryHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripSoCorpToken(v string) *CarBillSettlementQueryHeaders
	GetXAcsBtripSoCorpToken() *string
}

type CarBillSettlementQueryHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// feth00jqwls
	XAcsBtripSoCorpToken *string `json:"x-acs-btrip-so-corp-token,omitempty" xml:"x-acs-btrip-so-corp-token,omitempty"`
}

func (s CarBillSettlementQueryHeaders) String() string {
	return dara.Prettify(s)
}

func (s CarBillSettlementQueryHeaders) GoString() string {
	return s.String()
}

func (s *CarBillSettlementQueryHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *CarBillSettlementQueryHeaders) GetXAcsBtripSoCorpToken() *string {
	return s.XAcsBtripSoCorpToken
}

func (s *CarBillSettlementQueryHeaders) SetCommonHeaders(v map[string]*string) *CarBillSettlementQueryHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CarBillSettlementQueryHeaders) SetXAcsBtripSoCorpToken(v string) *CarBillSettlementQueryHeaders {
	s.XAcsBtripSoCorpToken = &v
	return s
}

func (s *CarBillSettlementQueryHeaders) Validate() error {
	return dara.Validate(s)
}
