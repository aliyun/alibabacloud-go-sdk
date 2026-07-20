// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHotelOrderCreateHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *HotelOrderCreateHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripCorpToken(v string) *HotelOrderCreateHeaders
	GetXAcsBtripCorpToken() *string
}

type HotelOrderCreateHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// feth00jqwls
	XAcsBtripCorpToken *string `json:"x-acs-btrip-corp-token,omitempty" xml:"x-acs-btrip-corp-token,omitempty"`
}

func (s HotelOrderCreateHeaders) String() string {
	return dara.Prettify(s)
}

func (s HotelOrderCreateHeaders) GoString() string {
	return s.String()
}

func (s *HotelOrderCreateHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *HotelOrderCreateHeaders) GetXAcsBtripCorpToken() *string {
	return s.XAcsBtripCorpToken
}

func (s *HotelOrderCreateHeaders) SetCommonHeaders(v map[string]*string) *HotelOrderCreateHeaders {
	s.CommonHeaders = v
	return s
}

func (s *HotelOrderCreateHeaders) SetXAcsBtripCorpToken(v string) *HotelOrderCreateHeaders {
	s.XAcsBtripCorpToken = &v
	return s
}

func (s *HotelOrderCreateHeaders) Validate() error {
	return dara.Validate(s)
}
