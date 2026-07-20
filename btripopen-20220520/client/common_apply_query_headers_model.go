// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCommonApplyQueryHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *CommonApplyQueryHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripSoCorpToken(v string) *CommonApplyQueryHeaders
	GetXAcsBtripSoCorpToken() *string
}

type CommonApplyQueryHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// feth00jqwls
	XAcsBtripSoCorpToken *string `json:"x-acs-btrip-so-corp-token,omitempty" xml:"x-acs-btrip-so-corp-token,omitempty"`
}

func (s CommonApplyQueryHeaders) String() string {
	return dara.Prettify(s)
}

func (s CommonApplyQueryHeaders) GoString() string {
	return s.String()
}

func (s *CommonApplyQueryHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *CommonApplyQueryHeaders) GetXAcsBtripSoCorpToken() *string {
	return s.XAcsBtripSoCorpToken
}

func (s *CommonApplyQueryHeaders) SetCommonHeaders(v map[string]*string) *CommonApplyQueryHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CommonApplyQueryHeaders) SetXAcsBtripSoCorpToken(v string) *CommonApplyQueryHeaders {
	s.XAcsBtripSoCorpToken = &v
	return s
}

func (s *CommonApplyQueryHeaders) Validate() error {
	return dara.Validate(s)
}
