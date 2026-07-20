// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHotelStaticInfoHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *HotelStaticInfoHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripCorpToken(v string) *HotelStaticInfoHeaders
	GetXAcsBtripCorpToken() *string
}

type HotelStaticInfoHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// feth00jqwls
	XAcsBtripCorpToken *string `json:"x-acs-btrip-corp-token,omitempty" xml:"x-acs-btrip-corp-token,omitempty"`
}

func (s HotelStaticInfoHeaders) String() string {
	return dara.Prettify(s)
}

func (s HotelStaticInfoHeaders) GoString() string {
	return s.String()
}

func (s *HotelStaticInfoHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *HotelStaticInfoHeaders) GetXAcsBtripCorpToken() *string {
	return s.XAcsBtripCorpToken
}

func (s *HotelStaticInfoHeaders) SetCommonHeaders(v map[string]*string) *HotelStaticInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *HotelStaticInfoHeaders) SetXAcsBtripCorpToken(v string) *HotelStaticInfoHeaders {
	s.XAcsBtripCorpToken = &v
	return s
}

func (s *HotelStaticInfoHeaders) Validate() error {
	return dara.Validate(s)
}
