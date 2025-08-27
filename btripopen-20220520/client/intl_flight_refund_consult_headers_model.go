// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIntlFlightRefundConsultHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *IntlFlightRefundConsultHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripCorpToken(v string) *IntlFlightRefundConsultHeaders
	GetXAcsBtripCorpToken() *string
}

type IntlFlightRefundConsultHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// feth00jqwls
	XAcsBtripCorpToken *string `json:"x-acs-btrip-corp-token,omitempty" xml:"x-acs-btrip-corp-token,omitempty"`
}

func (s IntlFlightRefundConsultHeaders) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightRefundConsultHeaders) GoString() string {
	return s.String()
}

func (s *IntlFlightRefundConsultHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *IntlFlightRefundConsultHeaders) GetXAcsBtripCorpToken() *string {
	return s.XAcsBtripCorpToken
}

func (s *IntlFlightRefundConsultHeaders) SetCommonHeaders(v map[string]*string) *IntlFlightRefundConsultHeaders {
	s.CommonHeaders = v
	return s
}

func (s *IntlFlightRefundConsultHeaders) SetXAcsBtripCorpToken(v string) *IntlFlightRefundConsultHeaders {
	s.XAcsBtripCorpToken = &v
	return s
}

func (s *IntlFlightRefundConsultHeaders) Validate() error {
	return dara.Validate(s)
}
