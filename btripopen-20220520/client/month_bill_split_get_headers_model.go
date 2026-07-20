// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMonthBillSplitGetHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *MonthBillSplitGetHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripSoCorpToken(v string) *MonthBillSplitGetHeaders
	GetXAcsBtripSoCorpToken() *string
}

type MonthBillSplitGetHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// feth****wls
	XAcsBtripSoCorpToken *string `json:"x-acs-btrip-so-corp-token,omitempty" xml:"x-acs-btrip-so-corp-token,omitempty"`
}

func (s MonthBillSplitGetHeaders) String() string {
	return dara.Prettify(s)
}

func (s MonthBillSplitGetHeaders) GoString() string {
	return s.String()
}

func (s *MonthBillSplitGetHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *MonthBillSplitGetHeaders) GetXAcsBtripSoCorpToken() *string {
	return s.XAcsBtripSoCorpToken
}

func (s *MonthBillSplitGetHeaders) SetCommonHeaders(v map[string]*string) *MonthBillSplitGetHeaders {
	s.CommonHeaders = v
	return s
}

func (s *MonthBillSplitGetHeaders) SetXAcsBtripSoCorpToken(v string) *MonthBillSplitGetHeaders {
	s.XAcsBtripSoCorpToken = &v
	return s
}

func (s *MonthBillSplitGetHeaders) Validate() error {
	return dara.Validate(s)
}
