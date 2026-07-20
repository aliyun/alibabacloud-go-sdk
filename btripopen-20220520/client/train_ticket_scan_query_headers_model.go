// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTrainTicketScanQueryHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *TrainTicketScanQueryHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripSoCorpToken(v string) *TrainTicketScanQueryHeaders
	GetXAcsBtripSoCorpToken() *string
}

type TrainTicketScanQueryHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// feth00jqwls
	XAcsBtripSoCorpToken *string `json:"x-acs-btrip-so-corp-token,omitempty" xml:"x-acs-btrip-so-corp-token,omitempty"`
}

func (s TrainTicketScanQueryHeaders) String() string {
	return dara.Prettify(s)
}

func (s TrainTicketScanQueryHeaders) GoString() string {
	return s.String()
}

func (s *TrainTicketScanQueryHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *TrainTicketScanQueryHeaders) GetXAcsBtripSoCorpToken() *string {
	return s.XAcsBtripSoCorpToken
}

func (s *TrainTicketScanQueryHeaders) SetCommonHeaders(v map[string]*string) *TrainTicketScanQueryHeaders {
	s.CommonHeaders = v
	return s
}

func (s *TrainTicketScanQueryHeaders) SetXAcsBtripSoCorpToken(v string) *TrainTicketScanQueryHeaders {
	s.XAcsBtripSoCorpToken = &v
	return s
}

func (s *TrainTicketScanQueryHeaders) Validate() error {
	return dara.Validate(s)
}
