// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHotelOrderQueryHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *HotelOrderQueryHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripSoCorpToken(v string) *HotelOrderQueryHeaders
	GetXAcsBtripSoCorpToken() *string
}

type HotelOrderQueryHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// feth00jqwls
	XAcsBtripSoCorpToken *string `json:"x-acs-btrip-so-corp-token,omitempty" xml:"x-acs-btrip-so-corp-token,omitempty"`
}

func (s HotelOrderQueryHeaders) String() string {
	return dara.Prettify(s)
}

func (s HotelOrderQueryHeaders) GoString() string {
	return s.String()
}

func (s *HotelOrderQueryHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *HotelOrderQueryHeaders) GetXAcsBtripSoCorpToken() *string {
	return s.XAcsBtripSoCorpToken
}

func (s *HotelOrderQueryHeaders) SetCommonHeaders(v map[string]*string) *HotelOrderQueryHeaders {
	s.CommonHeaders = v
	return s
}

func (s *HotelOrderQueryHeaders) SetXAcsBtripSoCorpToken(v string) *HotelOrderQueryHeaders {
	s.XAcsBtripSoCorpToken = &v
	return s
}

func (s *HotelOrderQueryHeaders) Validate() error {
	return dara.Validate(s)
}
