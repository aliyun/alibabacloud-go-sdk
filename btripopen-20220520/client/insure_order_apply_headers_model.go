// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInsureOrderApplyHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *InsureOrderApplyHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripCorpToken(v string) *InsureOrderApplyHeaders
	GetXAcsBtripCorpToken() *string
}

type InsureOrderApplyHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// feth00jqwls
	XAcsBtripCorpToken *string `json:"x-acs-btrip-corp-token,omitempty" xml:"x-acs-btrip-corp-token,omitempty"`
}

func (s InsureOrderApplyHeaders) String() string {
	return dara.Prettify(s)
}

func (s InsureOrderApplyHeaders) GoString() string {
	return s.String()
}

func (s *InsureOrderApplyHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *InsureOrderApplyHeaders) GetXAcsBtripCorpToken() *string {
	return s.XAcsBtripCorpToken
}

func (s *InsureOrderApplyHeaders) SetCommonHeaders(v map[string]*string) *InsureOrderApplyHeaders {
	s.CommonHeaders = v
	return s
}

func (s *InsureOrderApplyHeaders) SetXAcsBtripCorpToken(v string) *InsureOrderApplyHeaders {
	s.XAcsBtripCorpToken = &v
	return s
}

func (s *InsureOrderApplyHeaders) Validate() error {
	return dara.Validate(s)
}
