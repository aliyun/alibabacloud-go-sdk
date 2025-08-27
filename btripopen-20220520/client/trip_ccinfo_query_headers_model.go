// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTripCCInfoQueryHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *TripCCInfoQueryHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripCorpToken(v string) *TripCCInfoQueryHeaders
	GetXAcsBtripCorpToken() *string
}

type TripCCInfoQueryHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// feth00jqwls
	XAcsBtripCorpToken *string `json:"x-acs-btrip-corp-token,omitempty" xml:"x-acs-btrip-corp-token,omitempty"`
}

func (s TripCCInfoQueryHeaders) String() string {
	return dara.Prettify(s)
}

func (s TripCCInfoQueryHeaders) GoString() string {
	return s.String()
}

func (s *TripCCInfoQueryHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *TripCCInfoQueryHeaders) GetXAcsBtripCorpToken() *string {
	return s.XAcsBtripCorpToken
}

func (s *TripCCInfoQueryHeaders) SetCommonHeaders(v map[string]*string) *TripCCInfoQueryHeaders {
	s.CommonHeaders = v
	return s
}

func (s *TripCCInfoQueryHeaders) SetXAcsBtripCorpToken(v string) *TripCCInfoQueryHeaders {
	s.XAcsBtripCorpToken = &v
	return s
}

func (s *TripCCInfoQueryHeaders) Validate() error {
	return dara.Validate(s)
}
