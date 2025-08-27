// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIntlFlightListingSearchHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *IntlFlightListingSearchHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripCorpToken(v string) *IntlFlightListingSearchHeaders
	GetXAcsBtripCorpToken() *string
}

type IntlFlightListingSearchHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// feth00jqwls
	XAcsBtripCorpToken *string `json:"x-acs-btrip-corp-token,omitempty" xml:"x-acs-btrip-corp-token,omitempty"`
}

func (s IntlFlightListingSearchHeaders) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightListingSearchHeaders) GoString() string {
	return s.String()
}

func (s *IntlFlightListingSearchHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *IntlFlightListingSearchHeaders) GetXAcsBtripCorpToken() *string {
	return s.XAcsBtripCorpToken
}

func (s *IntlFlightListingSearchHeaders) SetCommonHeaders(v map[string]*string) *IntlFlightListingSearchHeaders {
	s.CommonHeaders = v
	return s
}

func (s *IntlFlightListingSearchHeaders) SetXAcsBtripCorpToken(v string) *IntlFlightListingSearchHeaders {
	s.XAcsBtripCorpToken = &v
	return s
}

func (s *IntlFlightListingSearchHeaders) Validate() error {
	return dara.Validate(s)
}
