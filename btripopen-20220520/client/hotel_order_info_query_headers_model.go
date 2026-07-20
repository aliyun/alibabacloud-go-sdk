// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHotelOrderInfoQueryHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *HotelOrderInfoQueryHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripCorpToken(v string) *HotelOrderInfoQueryHeaders
	GetXAcsBtripCorpToken() *string
}

type HotelOrderInfoQueryHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// feth00jqwls
	XAcsBtripCorpToken *string `json:"x-acs-btrip-corp-token,omitempty" xml:"x-acs-btrip-corp-token,omitempty"`
}

func (s HotelOrderInfoQueryHeaders) String() string {
	return dara.Prettify(s)
}

func (s HotelOrderInfoQueryHeaders) GoString() string {
	return s.String()
}

func (s *HotelOrderInfoQueryHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *HotelOrderInfoQueryHeaders) GetXAcsBtripCorpToken() *string {
	return s.XAcsBtripCorpToken
}

func (s *HotelOrderInfoQueryHeaders) SetCommonHeaders(v map[string]*string) *HotelOrderInfoQueryHeaders {
	s.CommonHeaders = v
	return s
}

func (s *HotelOrderInfoQueryHeaders) SetXAcsBtripCorpToken(v string) *HotelOrderInfoQueryHeaders {
	s.XAcsBtripCorpToken = &v
	return s
}

func (s *HotelOrderInfoQueryHeaders) Validate() error {
	return dara.Validate(s)
}
