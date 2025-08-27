// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCorpTokenHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *CorpTokenHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripAccessToken(v string) *CorpTokenHeaders
	GetXAcsBtripAccessToken() *string
}

type CorpTokenHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// vf9_bvla0qs
	XAcsBtripAccessToken *string `json:"x-acs-btrip-access-token,omitempty" xml:"x-acs-btrip-access-token,omitempty"`
}

func (s CorpTokenHeaders) String() string {
	return dara.Prettify(s)
}

func (s CorpTokenHeaders) GoString() string {
	return s.String()
}

func (s *CorpTokenHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *CorpTokenHeaders) GetXAcsBtripAccessToken() *string {
	return s.XAcsBtripAccessToken
}

func (s *CorpTokenHeaders) SetCommonHeaders(v map[string]*string) *CorpTokenHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CorpTokenHeaders) SetXAcsBtripAccessToken(v string) *CorpTokenHeaders {
	s.XAcsBtripAccessToken = &v
	return s
}

func (s *CorpTokenHeaders) Validate() error {
	return dara.Validate(s)
}
