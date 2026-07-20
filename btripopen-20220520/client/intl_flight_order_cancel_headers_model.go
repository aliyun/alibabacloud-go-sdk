// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIntlFlightOrderCancelHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *IntlFlightOrderCancelHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripCorpToken(v string) *IntlFlightOrderCancelHeaders
	GetXAcsBtripCorpToken() *string
}

type IntlFlightOrderCancelHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// feth00jqwls
	XAcsBtripCorpToken *string `json:"x-acs-btrip-corp-token,omitempty" xml:"x-acs-btrip-corp-token,omitempty"`
}

func (s IntlFlightOrderCancelHeaders) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightOrderCancelHeaders) GoString() string {
	return s.String()
}

func (s *IntlFlightOrderCancelHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *IntlFlightOrderCancelHeaders) GetXAcsBtripCorpToken() *string {
	return s.XAcsBtripCorpToken
}

func (s *IntlFlightOrderCancelHeaders) SetCommonHeaders(v map[string]*string) *IntlFlightOrderCancelHeaders {
	s.CommonHeaders = v
	return s
}

func (s *IntlFlightOrderCancelHeaders) SetXAcsBtripCorpToken(v string) *IntlFlightOrderCancelHeaders {
	s.XAcsBtripCorpToken = &v
	return s
}

func (s *IntlFlightOrderCancelHeaders) Validate() error {
	return dara.Validate(s)
}
