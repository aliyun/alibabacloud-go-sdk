// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInsureRefundDetailHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *InsureRefundDetailHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripCorpToken(v string) *InsureRefundDetailHeaders
	GetXAcsBtripCorpToken() *string
}

type InsureRefundDetailHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// feth00jqwls
	XAcsBtripCorpToken *string `json:"x-acs-btrip-corp-token,omitempty" xml:"x-acs-btrip-corp-token,omitempty"`
}

func (s InsureRefundDetailHeaders) String() string {
	return dara.Prettify(s)
}

func (s InsureRefundDetailHeaders) GoString() string {
	return s.String()
}

func (s *InsureRefundDetailHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *InsureRefundDetailHeaders) GetXAcsBtripCorpToken() *string {
	return s.XAcsBtripCorpToken
}

func (s *InsureRefundDetailHeaders) SetCommonHeaders(v map[string]*string) *InsureRefundDetailHeaders {
	s.CommonHeaders = v
	return s
}

func (s *InsureRefundDetailHeaders) SetXAcsBtripCorpToken(v string) *InsureRefundDetailHeaders {
	s.XAcsBtripCorpToken = &v
	return s
}

func (s *InsureRefundDetailHeaders) Validate() error {
	return dara.Validate(s)
}
