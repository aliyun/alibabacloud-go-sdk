// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHotelOrderChangeApplyHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *HotelOrderChangeApplyHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripCorpToken(v string) *HotelOrderChangeApplyHeaders
	GetXAcsBtripCorpToken() *string
}

type HotelOrderChangeApplyHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// feth00jqwls
	XAcsBtripCorpToken *string `json:"x-acs-btrip-corp-token,omitempty" xml:"x-acs-btrip-corp-token,omitempty"`
}

func (s HotelOrderChangeApplyHeaders) String() string {
	return dara.Prettify(s)
}

func (s HotelOrderChangeApplyHeaders) GoString() string {
	return s.String()
}

func (s *HotelOrderChangeApplyHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *HotelOrderChangeApplyHeaders) GetXAcsBtripCorpToken() *string {
	return s.XAcsBtripCorpToken
}

func (s *HotelOrderChangeApplyHeaders) SetCommonHeaders(v map[string]*string) *HotelOrderChangeApplyHeaders {
	s.CommonHeaders = v
	return s
}

func (s *HotelOrderChangeApplyHeaders) SetXAcsBtripCorpToken(v string) *HotelOrderChangeApplyHeaders {
	s.XAcsBtripCorpToken = &v
	return s
}

func (s *HotelOrderChangeApplyHeaders) Validate() error {
	return dara.Validate(s)
}
