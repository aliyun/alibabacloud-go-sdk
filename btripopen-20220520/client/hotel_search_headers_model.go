// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHotelSearchHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *HotelSearchHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripCorpToken(v string) *HotelSearchHeaders
	GetXAcsBtripCorpToken() *string
}

type HotelSearchHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// feth00jqwls
	XAcsBtripCorpToken *string `json:"x-acs-btrip-corp-token,omitempty" xml:"x-acs-btrip-corp-token,omitempty"`
}

func (s HotelSearchHeaders) String() string {
	return dara.Prettify(s)
}

func (s HotelSearchHeaders) GoString() string {
	return s.String()
}

func (s *HotelSearchHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *HotelSearchHeaders) GetXAcsBtripCorpToken() *string {
	return s.XAcsBtripCorpToken
}

func (s *HotelSearchHeaders) SetCommonHeaders(v map[string]*string) *HotelSearchHeaders {
	s.CommonHeaders = v
	return s
}

func (s *HotelSearchHeaders) SetXAcsBtripCorpToken(v string) *HotelSearchHeaders {
	s.XAcsBtripCorpToken = &v
	return s
}

func (s *HotelSearchHeaders) Validate() error {
	return dara.Validate(s)
}
