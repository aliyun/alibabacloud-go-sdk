// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTicketChangingCancelHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *TicketChangingCancelHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripCorpToken(v string) *TicketChangingCancelHeaders
	GetXAcsBtripCorpToken() *string
}

type TicketChangingCancelHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// feth00jqwls
	XAcsBtripCorpToken *string `json:"x-acs-btrip-corp-token,omitempty" xml:"x-acs-btrip-corp-token,omitempty"`
}

func (s TicketChangingCancelHeaders) String() string {
	return dara.Prettify(s)
}

func (s TicketChangingCancelHeaders) GoString() string {
	return s.String()
}

func (s *TicketChangingCancelHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *TicketChangingCancelHeaders) GetXAcsBtripCorpToken() *string {
	return s.XAcsBtripCorpToken
}

func (s *TicketChangingCancelHeaders) SetCommonHeaders(v map[string]*string) *TicketChangingCancelHeaders {
	s.CommonHeaders = v
	return s
}

func (s *TicketChangingCancelHeaders) SetXAcsBtripCorpToken(v string) *TicketChangingCancelHeaders {
	s.XAcsBtripCorpToken = &v
	return s
}

func (s *TicketChangingCancelHeaders) Validate() error {
	return dara.Validate(s)
}
