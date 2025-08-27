// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVasBillSettlementQueryHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *VasBillSettlementQueryHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripCorpToken(v string) *VasBillSettlementQueryHeaders
	GetXAcsBtripCorpToken() *string
}

type VasBillSettlementQueryHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// feth00jqwls
	XAcsBtripCorpToken *string `json:"x-acs-btrip-corp-token,omitempty" xml:"x-acs-btrip-corp-token,omitempty"`
}

func (s VasBillSettlementQueryHeaders) String() string {
	return dara.Prettify(s)
}

func (s VasBillSettlementQueryHeaders) GoString() string {
	return s.String()
}

func (s *VasBillSettlementQueryHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *VasBillSettlementQueryHeaders) GetXAcsBtripCorpToken() *string {
	return s.XAcsBtripCorpToken
}

func (s *VasBillSettlementQueryHeaders) SetCommonHeaders(v map[string]*string) *VasBillSettlementQueryHeaders {
	s.CommonHeaders = v
	return s
}

func (s *VasBillSettlementQueryHeaders) SetXAcsBtripCorpToken(v string) *VasBillSettlementQueryHeaders {
	s.XAcsBtripCorpToken = &v
	return s
}

func (s *VasBillSettlementQueryHeaders) Validate() error {
	return dara.Validate(s)
}
