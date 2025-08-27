// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApplyListQueryHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *ApplyListQueryHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripSoCorpToken(v string) *ApplyListQueryHeaders
	GetXAcsBtripSoCorpToken() *string
}

type ApplyListQueryHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// feth00jqwls
	XAcsBtripSoCorpToken *string `json:"x-acs-btrip-so-corp-token,omitempty" xml:"x-acs-btrip-so-corp-token,omitempty"`
}

func (s ApplyListQueryHeaders) String() string {
	return dara.Prettify(s)
}

func (s ApplyListQueryHeaders) GoString() string {
	return s.String()
}

func (s *ApplyListQueryHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *ApplyListQueryHeaders) GetXAcsBtripSoCorpToken() *string {
	return s.XAcsBtripSoCorpToken
}

func (s *ApplyListQueryHeaders) SetCommonHeaders(v map[string]*string) *ApplyListQueryHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ApplyListQueryHeaders) SetXAcsBtripSoCorpToken(v string) *ApplyListQueryHeaders {
	s.XAcsBtripSoCorpToken = &v
	return s
}

func (s *ApplyListQueryHeaders) Validate() error {
	return dara.Validate(s)
}
