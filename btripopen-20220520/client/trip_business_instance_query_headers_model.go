// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTripBusinessInstanceQueryHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *TripBusinessInstanceQueryHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripSoCorpToken(v string) *TripBusinessInstanceQueryHeaders
	GetXAcsBtripSoCorpToken() *string
}

type TripBusinessInstanceQueryHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// feth00jqwls
	XAcsBtripSoCorpToken *string `json:"x-acs-btrip-so-corp-token,omitempty" xml:"x-acs-btrip-so-corp-token,omitempty"`
}

func (s TripBusinessInstanceQueryHeaders) String() string {
	return dara.Prettify(s)
}

func (s TripBusinessInstanceQueryHeaders) GoString() string {
	return s.String()
}

func (s *TripBusinessInstanceQueryHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *TripBusinessInstanceQueryHeaders) GetXAcsBtripSoCorpToken() *string {
	return s.XAcsBtripSoCorpToken
}

func (s *TripBusinessInstanceQueryHeaders) SetCommonHeaders(v map[string]*string) *TripBusinessInstanceQueryHeaders {
	s.CommonHeaders = v
	return s
}

func (s *TripBusinessInstanceQueryHeaders) SetXAcsBtripSoCorpToken(v string) *TripBusinessInstanceQueryHeaders {
	s.XAcsBtripSoCorpToken = &v
	return s
}

func (s *TripBusinessInstanceQueryHeaders) Validate() error {
	return dara.Validate(s)
}
