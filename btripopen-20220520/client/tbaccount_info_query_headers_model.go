// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTBAccountInfoQueryHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *TBAccountInfoQueryHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripCorpToken(v string) *TBAccountInfoQueryHeaders
	GetXAcsBtripCorpToken() *string
}

type TBAccountInfoQueryHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// feth00jqwls
	XAcsBtripCorpToken *string `json:"x-acs-btrip-corp-token,omitempty" xml:"x-acs-btrip-corp-token,omitempty"`
}

func (s TBAccountInfoQueryHeaders) String() string {
	return dara.Prettify(s)
}

func (s TBAccountInfoQueryHeaders) GoString() string {
	return s.String()
}

func (s *TBAccountInfoQueryHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *TBAccountInfoQueryHeaders) GetXAcsBtripCorpToken() *string {
	return s.XAcsBtripCorpToken
}

func (s *TBAccountInfoQueryHeaders) SetCommonHeaders(v map[string]*string) *TBAccountInfoQueryHeaders {
	s.CommonHeaders = v
	return s
}

func (s *TBAccountInfoQueryHeaders) SetXAcsBtripCorpToken(v string) *TBAccountInfoQueryHeaders {
	s.XAcsBtripCorpToken = &v
	return s
}

func (s *TBAccountInfoQueryHeaders) Validate() error {
	return dara.Validate(s)
}
