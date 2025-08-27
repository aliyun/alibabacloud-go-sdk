// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCarApplyModifyHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *CarApplyModifyHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripSoCorpToken(v string) *CarApplyModifyHeaders
	GetXAcsBtripSoCorpToken() *string
}

type CarApplyModifyHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// feth00jqwls
	XAcsBtripSoCorpToken *string `json:"x-acs-btrip-so-corp-token,omitempty" xml:"x-acs-btrip-so-corp-token,omitempty"`
}

func (s CarApplyModifyHeaders) String() string {
	return dara.Prettify(s)
}

func (s CarApplyModifyHeaders) GoString() string {
	return s.String()
}

func (s *CarApplyModifyHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *CarApplyModifyHeaders) GetXAcsBtripSoCorpToken() *string {
	return s.XAcsBtripSoCorpToken
}

func (s *CarApplyModifyHeaders) SetCommonHeaders(v map[string]*string) *CarApplyModifyHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CarApplyModifyHeaders) SetXAcsBtripSoCorpToken(v string) *CarApplyModifyHeaders {
	s.XAcsBtripSoCorpToken = &v
	return s
}

func (s *CarApplyModifyHeaders) Validate() error {
	return dara.Validate(s)
}
