// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApplyQueryHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *ApplyQueryHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripSoCorpToken(v string) *ApplyQueryHeaders
	GetXAcsBtripSoCorpToken() *string
}

type ApplyQueryHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// feth00jqwls
	XAcsBtripSoCorpToken *string `json:"x-acs-btrip-so-corp-token,omitempty" xml:"x-acs-btrip-so-corp-token,omitempty"`
}

func (s ApplyQueryHeaders) String() string {
	return dara.Prettify(s)
}

func (s ApplyQueryHeaders) GoString() string {
	return s.String()
}

func (s *ApplyQueryHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *ApplyQueryHeaders) GetXAcsBtripSoCorpToken() *string {
	return s.XAcsBtripSoCorpToken
}

func (s *ApplyQueryHeaders) SetCommonHeaders(v map[string]*string) *ApplyQueryHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ApplyQueryHeaders) SetXAcsBtripSoCorpToken(v string) *ApplyQueryHeaders {
	s.XAcsBtripSoCorpToken = &v
	return s
}

func (s *ApplyQueryHeaders) Validate() error {
	return dara.Validate(s)
}
