// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTravelStandardQueryHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *TravelStandardQueryHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripCorpToken(v string) *TravelStandardQueryHeaders
	GetXAcsBtripCorpToken() *string
}

type TravelStandardQueryHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// A header parameter obtained from the enterprise access credential operation. Add this parameter to the request header. In HTTP mode, you can use corp_token=value in the URL as an alternative.
	//
	// example:
	//
	// feth00jqwls
	XAcsBtripCorpToken *string `json:"x-acs-btrip-corp-token,omitempty" xml:"x-acs-btrip-corp-token,omitempty"`
}

func (s TravelStandardQueryHeaders) String() string {
	return dara.Prettify(s)
}

func (s TravelStandardQueryHeaders) GoString() string {
	return s.String()
}

func (s *TravelStandardQueryHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *TravelStandardQueryHeaders) GetXAcsBtripCorpToken() *string {
	return s.XAcsBtripCorpToken
}

func (s *TravelStandardQueryHeaders) SetCommonHeaders(v map[string]*string) *TravelStandardQueryHeaders {
	s.CommonHeaders = v
	return s
}

func (s *TravelStandardQueryHeaders) SetXAcsBtripCorpToken(v string) *TravelStandardQueryHeaders {
	s.XAcsBtripCorpToken = &v
	return s
}

func (s *TravelStandardQueryHeaders) Validate() error {
	return dara.Validate(s)
}
