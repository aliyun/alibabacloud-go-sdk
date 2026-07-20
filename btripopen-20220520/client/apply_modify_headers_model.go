// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApplyModifyHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *ApplyModifyHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripSoCorpToken(v string) *ApplyModifyHeaders
	GetXAcsBtripSoCorpToken() *string
}

type ApplyModifyHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// feth00jqwls
	XAcsBtripSoCorpToken *string `json:"x-acs-btrip-so-corp-token,omitempty" xml:"x-acs-btrip-so-corp-token,omitempty"`
}

func (s ApplyModifyHeaders) String() string {
	return dara.Prettify(s)
}

func (s ApplyModifyHeaders) GoString() string {
	return s.String()
}

func (s *ApplyModifyHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *ApplyModifyHeaders) GetXAcsBtripSoCorpToken() *string {
	return s.XAcsBtripSoCorpToken
}

func (s *ApplyModifyHeaders) SetCommonHeaders(v map[string]*string) *ApplyModifyHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ApplyModifyHeaders) SetXAcsBtripSoCorpToken(v string) *ApplyModifyHeaders {
	s.XAcsBtripSoCorpToken = &v
	return s
}

func (s *ApplyModifyHeaders) Validate() error {
	return dara.Validate(s)
}
