// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUserQueryHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *UserQueryHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripSoCorpToken(v string) *UserQueryHeaders
	GetXAcsBtripSoCorpToken() *string
}

type UserQueryHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// feth00jqwls
	XAcsBtripSoCorpToken *string `json:"x-acs-btrip-so-corp-token,omitempty" xml:"x-acs-btrip-so-corp-token,omitempty"`
}

func (s UserQueryHeaders) String() string {
	return dara.Prettify(s)
}

func (s UserQueryHeaders) GoString() string {
	return s.String()
}

func (s *UserQueryHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *UserQueryHeaders) GetXAcsBtripSoCorpToken() *string {
	return s.XAcsBtripSoCorpToken
}

func (s *UserQueryHeaders) SetCommonHeaders(v map[string]*string) *UserQueryHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UserQueryHeaders) SetXAcsBtripSoCorpToken(v string) *UserQueryHeaders {
	s.XAcsBtripSoCorpToken = &v
	return s
}

func (s *UserQueryHeaders) Validate() error {
	return dara.Validate(s)
}
