// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTrainBillSettlementQueryHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *TrainBillSettlementQueryHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripSoCorpToken(v string) *TrainBillSettlementQueryHeaders
	GetXAcsBtripSoCorpToken() *string
}

type TrainBillSettlementQueryHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// feth00jqwls
	XAcsBtripSoCorpToken *string `json:"x-acs-btrip-so-corp-token,omitempty" xml:"x-acs-btrip-so-corp-token,omitempty"`
}

func (s TrainBillSettlementQueryHeaders) String() string {
	return dara.Prettify(s)
}

func (s TrainBillSettlementQueryHeaders) GoString() string {
	return s.String()
}

func (s *TrainBillSettlementQueryHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *TrainBillSettlementQueryHeaders) GetXAcsBtripSoCorpToken() *string {
	return s.XAcsBtripSoCorpToken
}

func (s *TrainBillSettlementQueryHeaders) SetCommonHeaders(v map[string]*string) *TrainBillSettlementQueryHeaders {
	s.CommonHeaders = v
	return s
}

func (s *TrainBillSettlementQueryHeaders) SetXAcsBtripSoCorpToken(v string) *TrainBillSettlementQueryHeaders {
	s.XAcsBtripSoCorpToken = &v
	return s
}

func (s *TrainBillSettlementQueryHeaders) Validate() error {
	return dara.Validate(s)
}
