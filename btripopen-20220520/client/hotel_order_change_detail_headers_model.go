// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHotelOrderChangeDetailHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *HotelOrderChangeDetailHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripCorpToken(v string) *HotelOrderChangeDetailHeaders
	GetXAcsBtripCorpToken() *string
}

type HotelOrderChangeDetailHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// feth00jqwls
	XAcsBtripCorpToken *string `json:"x-acs-btrip-corp-token,omitempty" xml:"x-acs-btrip-corp-token,omitempty"`
}

func (s HotelOrderChangeDetailHeaders) String() string {
	return dara.Prettify(s)
}

func (s HotelOrderChangeDetailHeaders) GoString() string {
	return s.String()
}

func (s *HotelOrderChangeDetailHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *HotelOrderChangeDetailHeaders) GetXAcsBtripCorpToken() *string {
	return s.XAcsBtripCorpToken
}

func (s *HotelOrderChangeDetailHeaders) SetCommonHeaders(v map[string]*string) *HotelOrderChangeDetailHeaders {
	s.CommonHeaders = v
	return s
}

func (s *HotelOrderChangeDetailHeaders) SetXAcsBtripCorpToken(v string) *HotelOrderChangeDetailHeaders {
	s.XAcsBtripCorpToken = &v
	return s
}

func (s *HotelOrderChangeDetailHeaders) Validate() error {
	return dara.Validate(s)
}
