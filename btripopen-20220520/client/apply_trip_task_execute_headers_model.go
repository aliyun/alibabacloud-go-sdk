// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApplyTripTaskExecuteHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *ApplyTripTaskExecuteHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripSoCorpToken(v string) *ApplyTripTaskExecuteHeaders
	GetXAcsBtripSoCorpToken() *string
}

type ApplyTripTaskExecuteHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// feth00jqwls
	XAcsBtripSoCorpToken *string `json:"x-acs-btrip-so-corp-token,omitempty" xml:"x-acs-btrip-so-corp-token,omitempty"`
}

func (s ApplyTripTaskExecuteHeaders) String() string {
	return dara.Prettify(s)
}

func (s ApplyTripTaskExecuteHeaders) GoString() string {
	return s.String()
}

func (s *ApplyTripTaskExecuteHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *ApplyTripTaskExecuteHeaders) GetXAcsBtripSoCorpToken() *string {
	return s.XAcsBtripSoCorpToken
}

func (s *ApplyTripTaskExecuteHeaders) SetCommonHeaders(v map[string]*string) *ApplyTripTaskExecuteHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ApplyTripTaskExecuteHeaders) SetXAcsBtripSoCorpToken(v string) *ApplyTripTaskExecuteHeaders {
	s.XAcsBtripSoCorpToken = &v
	return s
}

func (s *ApplyTripTaskExecuteHeaders) Validate() error {
	return dara.Validate(s)
}
