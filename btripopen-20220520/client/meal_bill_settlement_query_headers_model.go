// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMealBillSettlementQueryHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *MealBillSettlementQueryHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripSoCorpToken(v string) *MealBillSettlementQueryHeaders
	GetXAcsBtripSoCorpToken() *string
}

type MealBillSettlementQueryHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// feth00jqwls
	XAcsBtripSoCorpToken *string `json:"x-acs-btrip-so-corp-token,omitempty" xml:"x-acs-btrip-so-corp-token,omitempty"`
}

func (s MealBillSettlementQueryHeaders) String() string {
	return dara.Prettify(s)
}

func (s MealBillSettlementQueryHeaders) GoString() string {
	return s.String()
}

func (s *MealBillSettlementQueryHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *MealBillSettlementQueryHeaders) GetXAcsBtripSoCorpToken() *string {
	return s.XAcsBtripSoCorpToken
}

func (s *MealBillSettlementQueryHeaders) SetCommonHeaders(v map[string]*string) *MealBillSettlementQueryHeaders {
	s.CommonHeaders = v
	return s
}

func (s *MealBillSettlementQueryHeaders) SetXAcsBtripSoCorpToken(v string) *MealBillSettlementQueryHeaders {
	s.XAcsBtripSoCorpToken = &v
	return s
}

func (s *MealBillSettlementQueryHeaders) Validate() error {
	return dara.Validate(s)
}
