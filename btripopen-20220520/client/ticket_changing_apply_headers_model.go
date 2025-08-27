// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTicketChangingApplyHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *TicketChangingApplyHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripCorpToken(v string) *TicketChangingApplyHeaders
	GetXAcsBtripCorpToken() *string
}

type TicketChangingApplyHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// feth00jqwls
	XAcsBtripCorpToken *string `json:"x-acs-btrip-corp-token,omitempty" xml:"x-acs-btrip-corp-token,omitempty"`
}

func (s TicketChangingApplyHeaders) String() string {
	return dara.Prettify(s)
}

func (s TicketChangingApplyHeaders) GoString() string {
	return s.String()
}

func (s *TicketChangingApplyHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *TicketChangingApplyHeaders) GetXAcsBtripCorpToken() *string {
	return s.XAcsBtripCorpToken
}

func (s *TicketChangingApplyHeaders) SetCommonHeaders(v map[string]*string) *TicketChangingApplyHeaders {
	s.CommonHeaders = v
	return s
}

func (s *TicketChangingApplyHeaders) SetXAcsBtripCorpToken(v string) *TicketChangingApplyHeaders {
	s.XAcsBtripCorpToken = &v
	return s
}

func (s *TicketChangingApplyHeaders) Validate() error {
	return dara.Validate(s)
}
