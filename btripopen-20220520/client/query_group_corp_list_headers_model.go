// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryGroupCorpListHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *QueryGroupCorpListHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripCorpToken(v string) *QueryGroupCorpListHeaders
	GetXAcsBtripCorpToken() *string
}

type QueryGroupCorpListHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// feth00jqwls
	XAcsBtripCorpToken *string `json:"x-acs-btrip-corp-token,omitempty" xml:"x-acs-btrip-corp-token,omitempty"`
}

func (s QueryGroupCorpListHeaders) String() string {
	return dara.Prettify(s)
}

func (s QueryGroupCorpListHeaders) GoString() string {
	return s.String()
}

func (s *QueryGroupCorpListHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *QueryGroupCorpListHeaders) GetXAcsBtripCorpToken() *string {
	return s.XAcsBtripCorpToken
}

func (s *QueryGroupCorpListHeaders) SetCommonHeaders(v map[string]*string) *QueryGroupCorpListHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryGroupCorpListHeaders) SetXAcsBtripCorpToken(v string) *QueryGroupCorpListHeaders {
	s.XAcsBtripCorpToken = &v
	return s
}

func (s *QueryGroupCorpListHeaders) Validate() error {
	return dara.Validate(s)
}
