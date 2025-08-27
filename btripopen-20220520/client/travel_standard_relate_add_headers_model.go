// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTravelStandardRelateAddHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *TravelStandardRelateAddHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripCorpToken(v string) *TravelStandardRelateAddHeaders
	GetXAcsBtripCorpToken() *string
}

type TravelStandardRelateAddHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// feth00jqwls
	XAcsBtripCorpToken *string `json:"x-acs-btrip-corp-token,omitempty" xml:"x-acs-btrip-corp-token,omitempty"`
}

func (s TravelStandardRelateAddHeaders) String() string {
	return dara.Prettify(s)
}

func (s TravelStandardRelateAddHeaders) GoString() string {
	return s.String()
}

func (s *TravelStandardRelateAddHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *TravelStandardRelateAddHeaders) GetXAcsBtripCorpToken() *string {
	return s.XAcsBtripCorpToken
}

func (s *TravelStandardRelateAddHeaders) SetCommonHeaders(v map[string]*string) *TravelStandardRelateAddHeaders {
	s.CommonHeaders = v
	return s
}

func (s *TravelStandardRelateAddHeaders) SetXAcsBtripCorpToken(v string) *TravelStandardRelateAddHeaders {
	s.XAcsBtripCorpToken = &v
	return s
}

func (s *TravelStandardRelateAddHeaders) Validate() error {
	return dara.Validate(s)
}
