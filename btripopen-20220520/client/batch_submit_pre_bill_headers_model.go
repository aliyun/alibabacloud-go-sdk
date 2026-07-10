// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchSubmitPreBillHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *BatchSubmitPreBillHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripSoCorpToken(v string) *BatchSubmitPreBillHeaders
	GetXAcsBtripSoCorpToken() *string
}

type BatchSubmitPreBillHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// A header parameter obtained from the "enterprise access credential" operation (added in the header). In HTTP mode, you can use so_corp_token=value or dtb_corp_token=value in the URL as a substitute.
	//
	// example:
	//
	// feth00jqwls
	XAcsBtripSoCorpToken *string `json:"x-acs-btrip-so-corp-token,omitempty" xml:"x-acs-btrip-so-corp-token,omitempty"`
}

func (s BatchSubmitPreBillHeaders) String() string {
	return dara.Prettify(s)
}

func (s BatchSubmitPreBillHeaders) GoString() string {
	return s.String()
}

func (s *BatchSubmitPreBillHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *BatchSubmitPreBillHeaders) GetXAcsBtripSoCorpToken() *string {
	return s.XAcsBtripSoCorpToken
}

func (s *BatchSubmitPreBillHeaders) SetCommonHeaders(v map[string]*string) *BatchSubmitPreBillHeaders {
	s.CommonHeaders = v
	return s
}

func (s *BatchSubmitPreBillHeaders) SetXAcsBtripSoCorpToken(v string) *BatchSubmitPreBillHeaders {
	s.XAcsBtripSoCorpToken = &v
	return s
}

func (s *BatchSubmitPreBillHeaders) Validate() error {
	return dara.Validate(s)
}
