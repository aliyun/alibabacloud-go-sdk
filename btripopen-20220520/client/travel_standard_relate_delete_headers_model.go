// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTravelStandardRelateDeleteHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *TravelStandardRelateDeleteHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripCorpToken(v string) *TravelStandardRelateDeleteHeaders
	GetXAcsBtripCorpToken() *string
}

type TravelStandardRelateDeleteHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// feth00jqwis
	XAcsBtripCorpToken *string `json:"x-acs-btrip-corp-token,omitempty" xml:"x-acs-btrip-corp-token,omitempty"`
}

func (s TravelStandardRelateDeleteHeaders) String() string {
	return dara.Prettify(s)
}

func (s TravelStandardRelateDeleteHeaders) GoString() string {
	return s.String()
}

func (s *TravelStandardRelateDeleteHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *TravelStandardRelateDeleteHeaders) GetXAcsBtripCorpToken() *string {
	return s.XAcsBtripCorpToken
}

func (s *TravelStandardRelateDeleteHeaders) SetCommonHeaders(v map[string]*string) *TravelStandardRelateDeleteHeaders {
	s.CommonHeaders = v
	return s
}

func (s *TravelStandardRelateDeleteHeaders) SetXAcsBtripCorpToken(v string) *TravelStandardRelateDeleteHeaders {
	s.XAcsBtripCorpToken = &v
	return s
}

func (s *TravelStandardRelateDeleteHeaders) Validate() error {
	return dara.Validate(s)
}
