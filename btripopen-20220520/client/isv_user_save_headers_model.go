// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIsvUserSaveHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *IsvUserSaveHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripSoCorpToken(v string) *IsvUserSaveHeaders
	GetXAcsBtripSoCorpToken() *string
}

type IsvUserSaveHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// feth00jqwls
	XAcsBtripSoCorpToken *string `json:"x-acs-btrip-so-corp-token,omitempty" xml:"x-acs-btrip-so-corp-token,omitempty"`
}

func (s IsvUserSaveHeaders) String() string {
	return dara.Prettify(s)
}

func (s IsvUserSaveHeaders) GoString() string {
	return s.String()
}

func (s *IsvUserSaveHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *IsvUserSaveHeaders) GetXAcsBtripSoCorpToken() *string {
	return s.XAcsBtripSoCorpToken
}

func (s *IsvUserSaveHeaders) SetCommonHeaders(v map[string]*string) *IsvUserSaveHeaders {
	s.CommonHeaders = v
	return s
}

func (s *IsvUserSaveHeaders) SetXAcsBtripSoCorpToken(v string) *IsvUserSaveHeaders {
	s.XAcsBtripSoCorpToken = &v
	return s
}

func (s *IsvUserSaveHeaders) Validate() error {
	return dara.Validate(s)
}
