// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIntlFlightSegmentAvailableCertHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *IntlFlightSegmentAvailableCertHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripCorpToken(v string) *IntlFlightSegmentAvailableCertHeaders
	GetXAcsBtripCorpToken() *string
}

type IntlFlightSegmentAvailableCertHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// feth00jqwls
	XAcsBtripCorpToken *string `json:"x-acs-btrip-corp-token,omitempty" xml:"x-acs-btrip-corp-token,omitempty"`
}

func (s IntlFlightSegmentAvailableCertHeaders) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightSegmentAvailableCertHeaders) GoString() string {
	return s.String()
}

func (s *IntlFlightSegmentAvailableCertHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *IntlFlightSegmentAvailableCertHeaders) GetXAcsBtripCorpToken() *string {
	return s.XAcsBtripCorpToken
}

func (s *IntlFlightSegmentAvailableCertHeaders) SetCommonHeaders(v map[string]*string) *IntlFlightSegmentAvailableCertHeaders {
	s.CommonHeaders = v
	return s
}

func (s *IntlFlightSegmentAvailableCertHeaders) SetXAcsBtripCorpToken(v string) *IntlFlightSegmentAvailableCertHeaders {
	s.XAcsBtripCorpToken = &v
	return s
}

func (s *IntlFlightSegmentAvailableCertHeaders) Validate() error {
	return dara.Validate(s)
}
