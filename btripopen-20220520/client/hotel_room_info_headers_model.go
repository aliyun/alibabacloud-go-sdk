// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHotelRoomInfoHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *HotelRoomInfoHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripCorpToken(v string) *HotelRoomInfoHeaders
	GetXAcsBtripCorpToken() *string
}

type HotelRoomInfoHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// feth00jqwls
	XAcsBtripCorpToken *string `json:"x-acs-btrip-corp-token,omitempty" xml:"x-acs-btrip-corp-token,omitempty"`
}

func (s HotelRoomInfoHeaders) String() string {
	return dara.Prettify(s)
}

func (s HotelRoomInfoHeaders) GoString() string {
	return s.String()
}

func (s *HotelRoomInfoHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *HotelRoomInfoHeaders) GetXAcsBtripCorpToken() *string {
	return s.XAcsBtripCorpToken
}

func (s *HotelRoomInfoHeaders) SetCommonHeaders(v map[string]*string) *HotelRoomInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *HotelRoomInfoHeaders) SetXAcsBtripCorpToken(v string) *HotelRoomInfoHeaders {
	s.XAcsBtripCorpToken = &v
	return s
}

func (s *HotelRoomInfoHeaders) Validate() error {
	return dara.Validate(s)
}
