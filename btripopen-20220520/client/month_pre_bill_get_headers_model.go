// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMonthPreBillGetHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *MonthPreBillGetHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripSoCorpToken(v string) *MonthPreBillGetHeaders
	GetXAcsBtripSoCorpToken() *string
}

type MonthPreBillGetHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsBtripSoCorpToken *string            `json:"x-acs-btrip-so-corp-token,omitempty" xml:"x-acs-btrip-so-corp-token,omitempty"`
}

func (s MonthPreBillGetHeaders) String() string {
	return dara.Prettify(s)
}

func (s MonthPreBillGetHeaders) GoString() string {
	return s.String()
}

func (s *MonthPreBillGetHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *MonthPreBillGetHeaders) GetXAcsBtripSoCorpToken() *string {
	return s.XAcsBtripSoCorpToken
}

func (s *MonthPreBillGetHeaders) SetCommonHeaders(v map[string]*string) *MonthPreBillGetHeaders {
	s.CommonHeaders = v
	return s
}

func (s *MonthPreBillGetHeaders) SetXAcsBtripSoCorpToken(v string) *MonthPreBillGetHeaders {
	s.XAcsBtripSoCorpToken = &v
	return s
}

func (s *MonthPreBillGetHeaders) Validate() error {
	return dara.Validate(s)
}
