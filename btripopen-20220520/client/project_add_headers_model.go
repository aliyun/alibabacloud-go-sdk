// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iProjectAddHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *ProjectAddHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripSoCorpToken(v string) *ProjectAddHeaders
	GetXAcsBtripSoCorpToken() *string
}

type ProjectAddHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// feth00jqwls
	XAcsBtripSoCorpToken *string `json:"x-acs-btrip-so-corp-token,omitempty" xml:"x-acs-btrip-so-corp-token,omitempty"`
}

func (s ProjectAddHeaders) String() string {
	return dara.Prettify(s)
}

func (s ProjectAddHeaders) GoString() string {
	return s.String()
}

func (s *ProjectAddHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *ProjectAddHeaders) GetXAcsBtripSoCorpToken() *string {
	return s.XAcsBtripSoCorpToken
}

func (s *ProjectAddHeaders) SetCommonHeaders(v map[string]*string) *ProjectAddHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ProjectAddHeaders) SetXAcsBtripSoCorpToken(v string) *ProjectAddHeaders {
	s.XAcsBtripSoCorpToken = &v
	return s
}

func (s *ProjectAddHeaders) Validate() error {
	return dara.Validate(s)
}
