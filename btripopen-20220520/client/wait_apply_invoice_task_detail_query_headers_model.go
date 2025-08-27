// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iWaitApplyInvoiceTaskDetailQueryHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *WaitApplyInvoiceTaskDetailQueryHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripSoCorpToken(v string) *WaitApplyInvoiceTaskDetailQueryHeaders
	GetXAcsBtripSoCorpToken() *string
}

type WaitApplyInvoiceTaskDetailQueryHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// feth00jqwls
	XAcsBtripSoCorpToken *string `json:"x-acs-btrip-so-corp-token,omitempty" xml:"x-acs-btrip-so-corp-token,omitempty"`
}

func (s WaitApplyInvoiceTaskDetailQueryHeaders) String() string {
	return dara.Prettify(s)
}

func (s WaitApplyInvoiceTaskDetailQueryHeaders) GoString() string {
	return s.String()
}

func (s *WaitApplyInvoiceTaskDetailQueryHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *WaitApplyInvoiceTaskDetailQueryHeaders) GetXAcsBtripSoCorpToken() *string {
	return s.XAcsBtripSoCorpToken
}

func (s *WaitApplyInvoiceTaskDetailQueryHeaders) SetCommonHeaders(v map[string]*string) *WaitApplyInvoiceTaskDetailQueryHeaders {
	s.CommonHeaders = v
	return s
}

func (s *WaitApplyInvoiceTaskDetailQueryHeaders) SetXAcsBtripSoCorpToken(v string) *WaitApplyInvoiceTaskDetailQueryHeaders {
	s.XAcsBtripSoCorpToken = &v
	return s
}

func (s *WaitApplyInvoiceTaskDetailQueryHeaders) Validate() error {
	return dara.Validate(s)
}
