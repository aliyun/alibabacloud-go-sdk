// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCooperatorHotelEventPushHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *CooperatorHotelEventPushHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripCorpToken(v string) *CooperatorHotelEventPushHeaders
	GetXAcsBtripCorpToken() *string
}

type CooperatorHotelEventPushHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// feth00jqwls
	XAcsBtripCorpToken *string `json:"x-acs-btrip-corp-token,omitempty" xml:"x-acs-btrip-corp-token,omitempty"`
}

func (s CooperatorHotelEventPushHeaders) String() string {
	return dara.Prettify(s)
}

func (s CooperatorHotelEventPushHeaders) GoString() string {
	return s.String()
}

func (s *CooperatorHotelEventPushHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *CooperatorHotelEventPushHeaders) GetXAcsBtripCorpToken() *string {
	return s.XAcsBtripCorpToken
}

func (s *CooperatorHotelEventPushHeaders) SetCommonHeaders(v map[string]*string) *CooperatorHotelEventPushHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CooperatorHotelEventPushHeaders) SetXAcsBtripCorpToken(v string) *CooperatorHotelEventPushHeaders {
	s.XAcsBtripCorpToken = &v
	return s
}

func (s *CooperatorHotelEventPushHeaders) Validate() error {
	return dara.Validate(s)
}
