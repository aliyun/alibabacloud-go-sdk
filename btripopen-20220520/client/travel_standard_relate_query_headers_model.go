// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTravelStandardRelateQueryHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *TravelStandardRelateQueryHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripCorpToken(v string) *TravelStandardRelateQueryHeaders
	GetXAcsBtripCorpToken() *string
}

type TravelStandardRelateQueryHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// feth00jqwls
	XAcsBtripCorpToken *string `json:"x-acs-btrip-corp-token,omitempty" xml:"x-acs-btrip-corp-token,omitempty"`
}

func (s TravelStandardRelateQueryHeaders) String() string {
	return dara.Prettify(s)
}

func (s TravelStandardRelateQueryHeaders) GoString() string {
	return s.String()
}

func (s *TravelStandardRelateQueryHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *TravelStandardRelateQueryHeaders) GetXAcsBtripCorpToken() *string {
	return s.XAcsBtripCorpToken
}

func (s *TravelStandardRelateQueryHeaders) SetCommonHeaders(v map[string]*string) *TravelStandardRelateQueryHeaders {
	s.CommonHeaders = v
	return s
}

func (s *TravelStandardRelateQueryHeaders) SetXAcsBtripCorpToken(v string) *TravelStandardRelateQueryHeaders {
	s.XAcsBtripCorpToken = &v
	return s
}

func (s *TravelStandardRelateQueryHeaders) Validate() error {
	return dara.Validate(s)
}
