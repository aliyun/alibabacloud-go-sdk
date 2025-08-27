// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHotelSuggestV2Headers interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *HotelSuggestV2Headers
	GetCommonHeaders() map[string]*string
	SetXAcsBtripCorpToken(v string) *HotelSuggestV2Headers
	GetXAcsBtripCorpToken() *string
}

type HotelSuggestV2Headers struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// feth00jqwls
	XAcsBtripCorpToken *string `json:"x-acs-btrip-corp-token,omitempty" xml:"x-acs-btrip-corp-token,omitempty"`
}

func (s HotelSuggestV2Headers) String() string {
	return dara.Prettify(s)
}

func (s HotelSuggestV2Headers) GoString() string {
	return s.String()
}

func (s *HotelSuggestV2Headers) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *HotelSuggestV2Headers) GetXAcsBtripCorpToken() *string {
	return s.XAcsBtripCorpToken
}

func (s *HotelSuggestV2Headers) SetCommonHeaders(v map[string]*string) *HotelSuggestV2Headers {
	s.CommonHeaders = v
	return s
}

func (s *HotelSuggestV2Headers) SetXAcsBtripCorpToken(v string) *HotelSuggestV2Headers {
	s.XAcsBtripCorpToken = &v
	return s
}

func (s *HotelSuggestV2Headers) Validate() error {
	return dara.Validate(s)
}
