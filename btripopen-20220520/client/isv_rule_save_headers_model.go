// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIsvRuleSaveHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *IsvRuleSaveHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripSoCorpToken(v string) *IsvRuleSaveHeaders
	GetXAcsBtripSoCorpToken() *string
}

type IsvRuleSaveHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// feth00jqwls
	XAcsBtripSoCorpToken *string `json:"x-acs-btrip-so-corp-token,omitempty" xml:"x-acs-btrip-so-corp-token,omitempty"`
}

func (s IsvRuleSaveHeaders) String() string {
	return dara.Prettify(s)
}

func (s IsvRuleSaveHeaders) GoString() string {
	return s.String()
}

func (s *IsvRuleSaveHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *IsvRuleSaveHeaders) GetXAcsBtripSoCorpToken() *string {
	return s.XAcsBtripSoCorpToken
}

func (s *IsvRuleSaveHeaders) SetCommonHeaders(v map[string]*string) *IsvRuleSaveHeaders {
	s.CommonHeaders = v
	return s
}

func (s *IsvRuleSaveHeaders) SetXAcsBtripSoCorpToken(v string) *IsvRuleSaveHeaders {
	s.XAcsBtripSoCorpToken = &v
	return s
}

func (s *IsvRuleSaveHeaders) Validate() error {
	return dara.Validate(s)
}
