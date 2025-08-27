// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCarApplyAddHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *CarApplyAddHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripSoCorpToken(v string) *CarApplyAddHeaders
	GetXAcsBtripSoCorpToken() *string
}

type CarApplyAddHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// feth00jqwls
	XAcsBtripSoCorpToken *string `json:"x-acs-btrip-so-corp-token,omitempty" xml:"x-acs-btrip-so-corp-token,omitempty"`
}

func (s CarApplyAddHeaders) String() string {
	return dara.Prettify(s)
}

func (s CarApplyAddHeaders) GoString() string {
	return s.String()
}

func (s *CarApplyAddHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *CarApplyAddHeaders) GetXAcsBtripSoCorpToken() *string {
	return s.XAcsBtripSoCorpToken
}

func (s *CarApplyAddHeaders) SetCommonHeaders(v map[string]*string) *CarApplyAddHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CarApplyAddHeaders) SetXAcsBtripSoCorpToken(v string) *CarApplyAddHeaders {
	s.XAcsBtripSoCorpToken = &v
	return s
}

func (s *CarApplyAddHeaders) Validate() error {
	return dara.Validate(s)
}
