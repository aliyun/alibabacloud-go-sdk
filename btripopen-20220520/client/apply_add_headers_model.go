// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApplyAddHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *ApplyAddHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripSoCorpToken(v string) *ApplyAddHeaders
	GetXAcsBtripSoCorpToken() *string
}

type ApplyAddHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// feth00jqwls
	XAcsBtripSoCorpToken *string `json:"x-acs-btrip-so-corp-token,omitempty" xml:"x-acs-btrip-so-corp-token,omitempty"`
}

func (s ApplyAddHeaders) String() string {
	return dara.Prettify(s)
}

func (s ApplyAddHeaders) GoString() string {
	return s.String()
}

func (s *ApplyAddHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *ApplyAddHeaders) GetXAcsBtripSoCorpToken() *string {
	return s.XAcsBtripSoCorpToken
}

func (s *ApplyAddHeaders) SetCommonHeaders(v map[string]*string) *ApplyAddHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ApplyAddHeaders) SetXAcsBtripSoCorpToken(v string) *ApplyAddHeaders {
	s.XAcsBtripSoCorpToken = &v
	return s
}

func (s *ApplyAddHeaders) Validate() error {
	return dara.Validate(s)
}
