// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCommonApplySyncHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *CommonApplySyncHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripSoCorpToken(v string) *CommonApplySyncHeaders
	GetXAcsBtripSoCorpToken() *string
}

type CommonApplySyncHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// feth00jqwls
	XAcsBtripSoCorpToken *string `json:"x-acs-btrip-so-corp-token,omitempty" xml:"x-acs-btrip-so-corp-token,omitempty"`
}

func (s CommonApplySyncHeaders) String() string {
	return dara.Prettify(s)
}

func (s CommonApplySyncHeaders) GoString() string {
	return s.String()
}

func (s *CommonApplySyncHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *CommonApplySyncHeaders) GetXAcsBtripSoCorpToken() *string {
	return s.XAcsBtripSoCorpToken
}

func (s *CommonApplySyncHeaders) SetCommonHeaders(v map[string]*string) *CommonApplySyncHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CommonApplySyncHeaders) SetXAcsBtripSoCorpToken(v string) *CommonApplySyncHeaders {
	s.XAcsBtripSoCorpToken = &v
	return s
}

func (s *CommonApplySyncHeaders) Validate() error {
	return dara.Validate(s)
}
