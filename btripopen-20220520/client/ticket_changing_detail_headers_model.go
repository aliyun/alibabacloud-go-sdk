// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTicketChangingDetailHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *TicketChangingDetailHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripCorpToken(v string) *TicketChangingDetailHeaders
	GetXAcsBtripCorpToken() *string
}

type TicketChangingDetailHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// feth00jqwis
	XAcsBtripCorpToken *string `json:"x-acs-btrip-corp-token,omitempty" xml:"x-acs-btrip-corp-token,omitempty"`
}

func (s TicketChangingDetailHeaders) String() string {
	return dara.Prettify(s)
}

func (s TicketChangingDetailHeaders) GoString() string {
	return s.String()
}

func (s *TicketChangingDetailHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *TicketChangingDetailHeaders) GetXAcsBtripCorpToken() *string {
	return s.XAcsBtripCorpToken
}

func (s *TicketChangingDetailHeaders) SetCommonHeaders(v map[string]*string) *TicketChangingDetailHeaders {
	s.CommonHeaders = v
	return s
}

func (s *TicketChangingDetailHeaders) SetXAcsBtripCorpToken(v string) *TicketChangingDetailHeaders {
	s.XAcsBtripCorpToken = &v
	return s
}

func (s *TicketChangingDetailHeaders) Validate() error {
	return dara.Validate(s)
}
