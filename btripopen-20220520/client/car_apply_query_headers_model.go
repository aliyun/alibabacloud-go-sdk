// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCarApplyQueryHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *CarApplyQueryHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripSoCorpToken(v string) *CarApplyQueryHeaders
	GetXAcsBtripSoCorpToken() *string
}

type CarApplyQueryHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// feth00jqwls
	XAcsBtripSoCorpToken *string `json:"x-acs-btrip-so-corp-token,omitempty" xml:"x-acs-btrip-so-corp-token,omitempty"`
}

func (s CarApplyQueryHeaders) String() string {
	return dara.Prettify(s)
}

func (s CarApplyQueryHeaders) GoString() string {
	return s.String()
}

func (s *CarApplyQueryHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *CarApplyQueryHeaders) GetXAcsBtripSoCorpToken() *string {
	return s.XAcsBtripSoCorpToken
}

func (s *CarApplyQueryHeaders) SetCommonHeaders(v map[string]*string) *CarApplyQueryHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CarApplyQueryHeaders) SetXAcsBtripSoCorpToken(v string) *CarApplyQueryHeaders {
	s.XAcsBtripSoCorpToken = &v
	return s
}

func (s *CarApplyQueryHeaders) Validate() error {
	return dara.Validate(s)
}
