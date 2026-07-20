// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHotelIndexInfoHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *HotelIndexInfoHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripCorpToken(v string) *HotelIndexInfoHeaders
	GetXAcsBtripCorpToken() *string
}

type HotelIndexInfoHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// feth00jqwls
	XAcsBtripCorpToken *string `json:"x-acs-btrip-corp-token,omitempty" xml:"x-acs-btrip-corp-token,omitempty"`
}

func (s HotelIndexInfoHeaders) String() string {
	return dara.Prettify(s)
}

func (s HotelIndexInfoHeaders) GoString() string {
	return s.String()
}

func (s *HotelIndexInfoHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *HotelIndexInfoHeaders) GetXAcsBtripCorpToken() *string {
	return s.XAcsBtripCorpToken
}

func (s *HotelIndexInfoHeaders) SetCommonHeaders(v map[string]*string) *HotelIndexInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *HotelIndexInfoHeaders) SetXAcsBtripCorpToken(v string) *HotelIndexInfoHeaders {
	s.XAcsBtripCorpToken = &v
	return s
}

func (s *HotelIndexInfoHeaders) Validate() error {
	return dara.Validate(s)
}
