// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTBAccountUnbindHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *TBAccountUnbindHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripCorpToken(v string) *TBAccountUnbindHeaders
	GetXAcsBtripCorpToken() *string
}

type TBAccountUnbindHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// feth00jqwls
	XAcsBtripCorpToken *string `json:"x-acs-btrip-corp-token,omitempty" xml:"x-acs-btrip-corp-token,omitempty"`
}

func (s TBAccountUnbindHeaders) String() string {
	return dara.Prettify(s)
}

func (s TBAccountUnbindHeaders) GoString() string {
	return s.String()
}

func (s *TBAccountUnbindHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *TBAccountUnbindHeaders) GetXAcsBtripCorpToken() *string {
	return s.XAcsBtripCorpToken
}

func (s *TBAccountUnbindHeaders) SetCommonHeaders(v map[string]*string) *TBAccountUnbindHeaders {
	s.CommonHeaders = v
	return s
}

func (s *TBAccountUnbindHeaders) SetXAcsBtripCorpToken(v string) *TBAccountUnbindHeaders {
	s.XAcsBtripCorpToken = &v
	return s
}

func (s *TBAccountUnbindHeaders) Validate() error {
	return dara.Validate(s)
}
