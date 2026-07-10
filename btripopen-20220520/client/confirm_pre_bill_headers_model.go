// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfirmPreBillHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *ConfirmPreBillHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripSoCorpToken(v string) *ConfirmPreBillHeaders
	GetXAcsBtripSoCorpToken() *string
}

type ConfirmPreBillHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// A header parameter obtained from the enterprise access credential operation. Add this parameter to the header. In HTTP mode, you can use so_corp_token=value or dtb_corp_token=value in the URL as an alternative.
	//
	// example:
	//
	// feth00jqwls
	XAcsBtripSoCorpToken *string `json:"x-acs-btrip-so-corp-token,omitempty" xml:"x-acs-btrip-so-corp-token,omitempty"`
}

func (s ConfirmPreBillHeaders) String() string {
	return dara.Prettify(s)
}

func (s ConfirmPreBillHeaders) GoString() string {
	return s.String()
}

func (s *ConfirmPreBillHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *ConfirmPreBillHeaders) GetXAcsBtripSoCorpToken() *string {
	return s.XAcsBtripSoCorpToken
}

func (s *ConfirmPreBillHeaders) SetCommonHeaders(v map[string]*string) *ConfirmPreBillHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ConfirmPreBillHeaders) SetXAcsBtripSoCorpToken(v string) *ConfirmPreBillHeaders {
	s.XAcsBtripSoCorpToken = &v
	return s
}

func (s *ConfirmPreBillHeaders) Validate() error {
	return dara.Validate(s)
}
