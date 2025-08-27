// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApplyApproveHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *ApplyApproveHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripSoCorpToken(v string) *ApplyApproveHeaders
	GetXAcsBtripSoCorpToken() *string
}

type ApplyApproveHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// feth00jqwls
	XAcsBtripSoCorpToken *string `json:"x-acs-btrip-so-corp-token,omitempty" xml:"x-acs-btrip-so-corp-token,omitempty"`
}

func (s ApplyApproveHeaders) String() string {
	return dara.Prettify(s)
}

func (s ApplyApproveHeaders) GoString() string {
	return s.String()
}

func (s *ApplyApproveHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *ApplyApproveHeaders) GetXAcsBtripSoCorpToken() *string {
	return s.XAcsBtripSoCorpToken
}

func (s *ApplyApproveHeaders) SetCommonHeaders(v map[string]*string) *ApplyApproveHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ApplyApproveHeaders) SetXAcsBtripSoCorpToken(v string) *ApplyApproveHeaders {
	s.XAcsBtripSoCorpToken = &v
	return s
}

func (s *ApplyApproveHeaders) Validate() error {
	return dara.Validate(s)
}
