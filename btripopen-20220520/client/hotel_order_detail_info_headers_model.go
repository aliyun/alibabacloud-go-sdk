// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHotelOrderDetailInfoHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *HotelOrderDetailInfoHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripCorpToken(v string) *HotelOrderDetailInfoHeaders
	GetXAcsBtripCorpToken() *string
}

type HotelOrderDetailInfoHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// feth00jqwls
	XAcsBtripCorpToken *string `json:"x-acs-btrip-corp-token,omitempty" xml:"x-acs-btrip-corp-token,omitempty"`
}

func (s HotelOrderDetailInfoHeaders) String() string {
	return dara.Prettify(s)
}

func (s HotelOrderDetailInfoHeaders) GoString() string {
	return s.String()
}

func (s *HotelOrderDetailInfoHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *HotelOrderDetailInfoHeaders) GetXAcsBtripCorpToken() *string {
	return s.XAcsBtripCorpToken
}

func (s *HotelOrderDetailInfoHeaders) SetCommonHeaders(v map[string]*string) *HotelOrderDetailInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *HotelOrderDetailInfoHeaders) SetXAcsBtripCorpToken(v string) *HotelOrderDetailInfoHeaders {
	s.XAcsBtripCorpToken = &v
	return s
}

func (s *HotelOrderDetailInfoHeaders) Validate() error {
	return dara.Validate(s)
}
