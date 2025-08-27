// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBtripBillInfoAdjustHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *BtripBillInfoAdjustHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripCorpToken(v string) *BtripBillInfoAdjustHeaders
	GetXAcsBtripCorpToken() *string
}

type BtripBillInfoAdjustHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// feth00jqwls
	XAcsBtripCorpToken *string `json:"x-acs-btrip-corp-token,omitempty" xml:"x-acs-btrip-corp-token,omitempty"`
}

func (s BtripBillInfoAdjustHeaders) String() string {
	return dara.Prettify(s)
}

func (s BtripBillInfoAdjustHeaders) GoString() string {
	return s.String()
}

func (s *BtripBillInfoAdjustHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *BtripBillInfoAdjustHeaders) GetXAcsBtripCorpToken() *string {
	return s.XAcsBtripCorpToken
}

func (s *BtripBillInfoAdjustHeaders) SetCommonHeaders(v map[string]*string) *BtripBillInfoAdjustHeaders {
	s.CommonHeaders = v
	return s
}

func (s *BtripBillInfoAdjustHeaders) SetXAcsBtripCorpToken(v string) *BtripBillInfoAdjustHeaders {
	s.XAcsBtripCorpToken = &v
	return s
}

func (s *BtripBillInfoAdjustHeaders) Validate() error {
	return dara.Validate(s)
}
