// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iProjectDeleteHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *ProjectDeleteHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripSoCorpToken(v string) *ProjectDeleteHeaders
	GetXAcsBtripSoCorpToken() *string
}

type ProjectDeleteHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// feth00jqwls
	XAcsBtripSoCorpToken *string `json:"x-acs-btrip-so-corp-token,omitempty" xml:"x-acs-btrip-so-corp-token,omitempty"`
}

func (s ProjectDeleteHeaders) String() string {
	return dara.Prettify(s)
}

func (s ProjectDeleteHeaders) GoString() string {
	return s.String()
}

func (s *ProjectDeleteHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *ProjectDeleteHeaders) GetXAcsBtripSoCorpToken() *string {
	return s.XAcsBtripSoCorpToken
}

func (s *ProjectDeleteHeaders) SetCommonHeaders(v map[string]*string) *ProjectDeleteHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ProjectDeleteHeaders) SetXAcsBtripSoCorpToken(v string) *ProjectDeleteHeaders {
	s.XAcsBtripSoCorpToken = &v
	return s
}

func (s *ProjectDeleteHeaders) Validate() error {
	return dara.Validate(s)
}
