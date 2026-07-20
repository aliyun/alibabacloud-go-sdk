// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHotelExceedApplyQueryHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *HotelExceedApplyQueryHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripSoCorpToken(v string) *HotelExceedApplyQueryHeaders
	GetXAcsBtripSoCorpToken() *string
}

type HotelExceedApplyQueryHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// feth00jqwls
	XAcsBtripSoCorpToken *string `json:"x-acs-btrip-so-corp-token,omitempty" xml:"x-acs-btrip-so-corp-token,omitempty"`
}

func (s HotelExceedApplyQueryHeaders) String() string {
	return dara.Prettify(s)
}

func (s HotelExceedApplyQueryHeaders) GoString() string {
	return s.String()
}

func (s *HotelExceedApplyQueryHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *HotelExceedApplyQueryHeaders) GetXAcsBtripSoCorpToken() *string {
	return s.XAcsBtripSoCorpToken
}

func (s *HotelExceedApplyQueryHeaders) SetCommonHeaders(v map[string]*string) *HotelExceedApplyQueryHeaders {
	s.CommonHeaders = v
	return s
}

func (s *HotelExceedApplyQueryHeaders) SetXAcsBtripSoCorpToken(v string) *HotelExceedApplyQueryHeaders {
	s.XAcsBtripSoCorpToken = &v
	return s
}

func (s *HotelExceedApplyQueryHeaders) Validate() error {
	return dara.Validate(s)
}
