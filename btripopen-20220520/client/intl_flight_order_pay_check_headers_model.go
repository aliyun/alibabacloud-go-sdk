// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIntlFlightOrderPayCheckHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *IntlFlightOrderPayCheckHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripCorpToken(v string) *IntlFlightOrderPayCheckHeaders
	GetXAcsBtripCorpToken() *string
}

type IntlFlightOrderPayCheckHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// feth00jqwis
	XAcsBtripCorpToken *string `json:"x-acs-btrip-corp-token,omitempty" xml:"x-acs-btrip-corp-token,omitempty"`
}

func (s IntlFlightOrderPayCheckHeaders) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightOrderPayCheckHeaders) GoString() string {
	return s.String()
}

func (s *IntlFlightOrderPayCheckHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *IntlFlightOrderPayCheckHeaders) GetXAcsBtripCorpToken() *string {
	return s.XAcsBtripCorpToken
}

func (s *IntlFlightOrderPayCheckHeaders) SetCommonHeaders(v map[string]*string) *IntlFlightOrderPayCheckHeaders {
	s.CommonHeaders = v
	return s
}

func (s *IntlFlightOrderPayCheckHeaders) SetXAcsBtripCorpToken(v string) *IntlFlightOrderPayCheckHeaders {
	s.XAcsBtripCorpToken = &v
	return s
}

func (s *IntlFlightOrderPayCheckHeaders) Validate() error {
	return dara.Validate(s)
}
