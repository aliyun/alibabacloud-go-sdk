// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTravelStandardListQueryHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *TravelStandardListQueryHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripCorpToken(v string) *TravelStandardListQueryHeaders
	GetXAcsBtripCorpToken() *string
}

type TravelStandardListQueryHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// feth00jqwls
	XAcsBtripCorpToken *string `json:"x-acs-btrip-corp-token,omitempty" xml:"x-acs-btrip-corp-token,omitempty"`
}

func (s TravelStandardListQueryHeaders) String() string {
	return dara.Prettify(s)
}

func (s TravelStandardListQueryHeaders) GoString() string {
	return s.String()
}

func (s *TravelStandardListQueryHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *TravelStandardListQueryHeaders) GetXAcsBtripCorpToken() *string {
	return s.XAcsBtripCorpToken
}

func (s *TravelStandardListQueryHeaders) SetCommonHeaders(v map[string]*string) *TravelStandardListQueryHeaders {
	s.CommonHeaders = v
	return s
}

func (s *TravelStandardListQueryHeaders) SetXAcsBtripCorpToken(v string) *TravelStandardListQueryHeaders {
	s.XAcsBtripCorpToken = &v
	return s
}

func (s *TravelStandardListQueryHeaders) Validate() error {
	return dara.Validate(s)
}
