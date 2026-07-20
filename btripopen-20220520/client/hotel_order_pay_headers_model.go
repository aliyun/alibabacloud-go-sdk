// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHotelOrderPayHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *HotelOrderPayHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripCorpToken(v string) *HotelOrderPayHeaders
	GetXAcsBtripCorpToken() *string
}

type HotelOrderPayHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// feth00jqwls
	XAcsBtripCorpToken *string `json:"x-acs-btrip-corp-token,omitempty" xml:"x-acs-btrip-corp-token,omitempty"`
}

func (s HotelOrderPayHeaders) String() string {
	return dara.Prettify(s)
}

func (s HotelOrderPayHeaders) GoString() string {
	return s.String()
}

func (s *HotelOrderPayHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *HotelOrderPayHeaders) GetXAcsBtripCorpToken() *string {
	return s.XAcsBtripCorpToken
}

func (s *HotelOrderPayHeaders) SetCommonHeaders(v map[string]*string) *HotelOrderPayHeaders {
	s.CommonHeaders = v
	return s
}

func (s *HotelOrderPayHeaders) SetXAcsBtripCorpToken(v string) *HotelOrderPayHeaders {
	s.XAcsBtripCorpToken = &v
	return s
}

func (s *HotelOrderPayHeaders) Validate() error {
	return dara.Validate(s)
}
