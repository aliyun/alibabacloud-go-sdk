// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTicketChangingFlightListHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *TicketChangingFlightListHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripCorpToken(v string) *TicketChangingFlightListHeaders
	GetXAcsBtripCorpToken() *string
}

type TicketChangingFlightListHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// feth00jqwls
	XAcsBtripCorpToken *string `json:"x-acs-btrip-corp-token,omitempty" xml:"x-acs-btrip-corp-token,omitempty"`
}

func (s TicketChangingFlightListHeaders) String() string {
	return dara.Prettify(s)
}

func (s TicketChangingFlightListHeaders) GoString() string {
	return s.String()
}

func (s *TicketChangingFlightListHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *TicketChangingFlightListHeaders) GetXAcsBtripCorpToken() *string {
	return s.XAcsBtripCorpToken
}

func (s *TicketChangingFlightListHeaders) SetCommonHeaders(v map[string]*string) *TicketChangingFlightListHeaders {
	s.CommonHeaders = v
	return s
}

func (s *TicketChangingFlightListHeaders) SetXAcsBtripCorpToken(v string) *TicketChangingFlightListHeaders {
	s.XAcsBtripCorpToken = &v
	return s
}

func (s *TicketChangingFlightListHeaders) Validate() error {
	return dara.Validate(s)
}
