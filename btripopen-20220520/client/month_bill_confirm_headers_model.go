// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMonthBillConfirmHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *MonthBillConfirmHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripCorpToken(v string) *MonthBillConfirmHeaders
	GetXAcsBtripCorpToken() *string
}

type MonthBillConfirmHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// feth00jqwis
	XAcsBtripCorpToken *string `json:"x-acs-btrip-corp-token,omitempty" xml:"x-acs-btrip-corp-token,omitempty"`
}

func (s MonthBillConfirmHeaders) String() string {
	return dara.Prettify(s)
}

func (s MonthBillConfirmHeaders) GoString() string {
	return s.String()
}

func (s *MonthBillConfirmHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *MonthBillConfirmHeaders) GetXAcsBtripCorpToken() *string {
	return s.XAcsBtripCorpToken
}

func (s *MonthBillConfirmHeaders) SetCommonHeaders(v map[string]*string) *MonthBillConfirmHeaders {
	s.CommonHeaders = v
	return s
}

func (s *MonthBillConfirmHeaders) SetXAcsBtripCorpToken(v string) *MonthBillConfirmHeaders {
	s.XAcsBtripCorpToken = &v
	return s
}

func (s *MonthBillConfirmHeaders) Validate() error {
	return dara.Validate(s)
}
