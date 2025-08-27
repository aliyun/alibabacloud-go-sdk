// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHotelOrderCancelHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *HotelOrderCancelHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripCorpToken(v string) *HotelOrderCancelHeaders
	GetXAcsBtripCorpToken() *string
}

type HotelOrderCancelHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// feth00jqwls
	XAcsBtripCorpToken *string `json:"x-acs-btrip-corp-token,omitempty" xml:"x-acs-btrip-corp-token,omitempty"`
}

func (s HotelOrderCancelHeaders) String() string {
	return dara.Prettify(s)
}

func (s HotelOrderCancelHeaders) GoString() string {
	return s.String()
}

func (s *HotelOrderCancelHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *HotelOrderCancelHeaders) GetXAcsBtripCorpToken() *string {
	return s.XAcsBtripCorpToken
}

func (s *HotelOrderCancelHeaders) SetCommonHeaders(v map[string]*string) *HotelOrderCancelHeaders {
	s.CommonHeaders = v
	return s
}

func (s *HotelOrderCancelHeaders) SetXAcsBtripCorpToken(v string) *HotelOrderCancelHeaders {
	s.XAcsBtripCorpToken = &v
	return s
}

func (s *HotelOrderCancelHeaders) Validate() error {
	return dara.Validate(s)
}
