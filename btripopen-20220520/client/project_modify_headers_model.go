// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iProjectModifyHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *ProjectModifyHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripSoCorpToken(v string) *ProjectModifyHeaders
	GetXAcsBtripSoCorpToken() *string
}

type ProjectModifyHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// feth00jqwls
	XAcsBtripSoCorpToken *string `json:"x-acs-btrip-so-corp-token,omitempty" xml:"x-acs-btrip-so-corp-token,omitempty"`
}

func (s ProjectModifyHeaders) String() string {
	return dara.Prettify(s)
}

func (s ProjectModifyHeaders) GoString() string {
	return s.String()
}

func (s *ProjectModifyHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *ProjectModifyHeaders) GetXAcsBtripSoCorpToken() *string {
	return s.XAcsBtripSoCorpToken
}

func (s *ProjectModifyHeaders) SetCommonHeaders(v map[string]*string) *ProjectModifyHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ProjectModifyHeaders) SetXAcsBtripSoCorpToken(v string) *ProjectModifyHeaders {
	s.XAcsBtripSoCorpToken = &v
	return s
}

func (s *ProjectModifyHeaders) Validate() error {
	return dara.Validate(s)
}
