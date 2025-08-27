// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryCorpDetailInfoHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *QueryCorpDetailInfoHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripCorpToken(v string) *QueryCorpDetailInfoHeaders
	GetXAcsBtripCorpToken() *string
}

type QueryCorpDetailInfoHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// aqfr****21
	XAcsBtripCorpToken *string `json:"x-acs-btrip-corp-token,omitempty" xml:"x-acs-btrip-corp-token,omitempty"`
}

func (s QueryCorpDetailInfoHeaders) String() string {
	return dara.Prettify(s)
}

func (s QueryCorpDetailInfoHeaders) GoString() string {
	return s.String()
}

func (s *QueryCorpDetailInfoHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *QueryCorpDetailInfoHeaders) GetXAcsBtripCorpToken() *string {
	return s.XAcsBtripCorpToken
}

func (s *QueryCorpDetailInfoHeaders) SetCommonHeaders(v map[string]*string) *QueryCorpDetailInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryCorpDetailInfoHeaders) SetXAcsBtripCorpToken(v string) *QueryCorpDetailInfoHeaders {
	s.XAcsBtripCorpToken = &v
	return s
}

func (s *QueryCorpDetailInfoHeaders) Validate() error {
	return dara.Validate(s)
}
