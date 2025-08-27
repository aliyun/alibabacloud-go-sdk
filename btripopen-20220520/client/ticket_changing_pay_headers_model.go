// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTicketChangingPayHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *TicketChangingPayHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripCorpToken(v string) *TicketChangingPayHeaders
	GetXAcsBtripCorpToken() *string
}

type TicketChangingPayHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// feth00jqwls
	XAcsBtripCorpToken *string `json:"x-acs-btrip-corp-token,omitempty" xml:"x-acs-btrip-corp-token,omitempty"`
}

func (s TicketChangingPayHeaders) String() string {
	return dara.Prettify(s)
}

func (s TicketChangingPayHeaders) GoString() string {
	return s.String()
}

func (s *TicketChangingPayHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *TicketChangingPayHeaders) GetXAcsBtripCorpToken() *string {
	return s.XAcsBtripCorpToken
}

func (s *TicketChangingPayHeaders) SetCommonHeaders(v map[string]*string) *TicketChangingPayHeaders {
	s.CommonHeaders = v
	return s
}

func (s *TicketChangingPayHeaders) SetXAcsBtripCorpToken(v string) *TicketChangingPayHeaders {
	s.XAcsBtripCorpToken = &v
	return s
}

func (s *TicketChangingPayHeaders) Validate() error {
	return dara.Validate(s)
}
