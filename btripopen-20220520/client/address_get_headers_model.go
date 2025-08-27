// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddressGetHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *AddressGetHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripSoCorpToken(v string) *AddressGetHeaders
	GetXAcsBtripSoCorpToken() *string
}

type AddressGetHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// feth00jqwls
	XAcsBtripSoCorpToken *string `json:"x-acs-btrip-so-corp-token,omitempty" xml:"x-acs-btrip-so-corp-token,omitempty"`
}

func (s AddressGetHeaders) String() string {
	return dara.Prettify(s)
}

func (s AddressGetHeaders) GoString() string {
	return s.String()
}

func (s *AddressGetHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *AddressGetHeaders) GetXAcsBtripSoCorpToken() *string {
	return s.XAcsBtripSoCorpToken
}

func (s *AddressGetHeaders) SetCommonHeaders(v map[string]*string) *AddressGetHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AddressGetHeaders) SetXAcsBtripSoCorpToken(v string) *AddressGetHeaders {
	s.XAcsBtripSoCorpToken = &v
	return s
}

func (s *AddressGetHeaders) Validate() error {
	return dara.Validate(s)
}
