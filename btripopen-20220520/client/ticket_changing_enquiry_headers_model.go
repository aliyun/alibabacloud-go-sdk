// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTicketChangingEnquiryHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *TicketChangingEnquiryHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripCorpToken(v string) *TicketChangingEnquiryHeaders
	GetXAcsBtripCorpToken() *string
}

type TicketChangingEnquiryHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// feth00jqwls
	XAcsBtripCorpToken *string `json:"x-acs-btrip-corp-token,omitempty" xml:"x-acs-btrip-corp-token,omitempty"`
}

func (s TicketChangingEnquiryHeaders) String() string {
	return dara.Prettify(s)
}

func (s TicketChangingEnquiryHeaders) GoString() string {
	return s.String()
}

func (s *TicketChangingEnquiryHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *TicketChangingEnquiryHeaders) GetXAcsBtripCorpToken() *string {
	return s.XAcsBtripCorpToken
}

func (s *TicketChangingEnquiryHeaders) SetCommonHeaders(v map[string]*string) *TicketChangingEnquiryHeaders {
	s.CommonHeaders = v
	return s
}

func (s *TicketChangingEnquiryHeaders) SetXAcsBtripCorpToken(v string) *TicketChangingEnquiryHeaders {
	s.XAcsBtripCorpToken = &v
	return s
}

func (s *TicketChangingEnquiryHeaders) Validate() error {
	return dara.Validate(s)
}
