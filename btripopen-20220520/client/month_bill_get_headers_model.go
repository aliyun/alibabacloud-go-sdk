// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMonthBillGetHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *MonthBillGetHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripSoCorpToken(v string) *MonthBillGetHeaders
	GetXAcsBtripSoCorpToken() *string
}

type MonthBillGetHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// feth00jqwls
	XAcsBtripSoCorpToken *string `json:"x-acs-btrip-so-corp-token,omitempty" xml:"x-acs-btrip-so-corp-token,omitempty"`
}

func (s MonthBillGetHeaders) String() string {
	return dara.Prettify(s)
}

func (s MonthBillGetHeaders) GoString() string {
	return s.String()
}

func (s *MonthBillGetHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *MonthBillGetHeaders) GetXAcsBtripSoCorpToken() *string {
	return s.XAcsBtripSoCorpToken
}

func (s *MonthBillGetHeaders) SetCommonHeaders(v map[string]*string) *MonthBillGetHeaders {
	s.CommonHeaders = v
	return s
}

func (s *MonthBillGetHeaders) SetXAcsBtripSoCorpToken(v string) *MonthBillGetHeaders {
	s.XAcsBtripSoCorpToken = &v
	return s
}

func (s *MonthBillGetHeaders) Validate() error {
	return dara.Validate(s)
}
