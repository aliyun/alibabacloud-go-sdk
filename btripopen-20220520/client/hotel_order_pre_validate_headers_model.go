// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHotelOrderPreValidateHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *HotelOrderPreValidateHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripCorpToken(v string) *HotelOrderPreValidateHeaders
	GetXAcsBtripCorpToken() *string
}

type HotelOrderPreValidateHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// feth00jqwls
	XAcsBtripCorpToken *string `json:"x-acs-btrip-corp-token,omitempty" xml:"x-acs-btrip-corp-token,omitempty"`
}

func (s HotelOrderPreValidateHeaders) String() string {
	return dara.Prettify(s)
}

func (s HotelOrderPreValidateHeaders) GoString() string {
	return s.String()
}

func (s *HotelOrderPreValidateHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *HotelOrderPreValidateHeaders) GetXAcsBtripCorpToken() *string {
	return s.XAcsBtripCorpToken
}

func (s *HotelOrderPreValidateHeaders) SetCommonHeaders(v map[string]*string) *HotelOrderPreValidateHeaders {
	s.CommonHeaders = v
	return s
}

func (s *HotelOrderPreValidateHeaders) SetXAcsBtripCorpToken(v string) *HotelOrderPreValidateHeaders {
	s.XAcsBtripCorpToken = &v
	return s
}

func (s *HotelOrderPreValidateHeaders) Validate() error {
	return dara.Validate(s)
}
