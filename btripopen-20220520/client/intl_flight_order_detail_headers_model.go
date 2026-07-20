// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIntlFlightOrderDetailHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *IntlFlightOrderDetailHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripCorpToken(v string) *IntlFlightOrderDetailHeaders
	GetXAcsBtripCorpToken() *string
}

type IntlFlightOrderDetailHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// feth00jqwls
	XAcsBtripCorpToken *string `json:"x-acs-btrip-corp-token,omitempty" xml:"x-acs-btrip-corp-token,omitempty"`
}

func (s IntlFlightOrderDetailHeaders) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightOrderDetailHeaders) GoString() string {
	return s.String()
}

func (s *IntlFlightOrderDetailHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *IntlFlightOrderDetailHeaders) GetXAcsBtripCorpToken() *string {
	return s.XAcsBtripCorpToken
}

func (s *IntlFlightOrderDetailHeaders) SetCommonHeaders(v map[string]*string) *IntlFlightOrderDetailHeaders {
	s.CommonHeaders = v
	return s
}

func (s *IntlFlightOrderDetailHeaders) SetXAcsBtripCorpToken(v string) *IntlFlightOrderDetailHeaders {
	s.XAcsBtripCorpToken = &v
	return s
}

func (s *IntlFlightOrderDetailHeaders) Validate() error {
	return dara.Validate(s)
}
