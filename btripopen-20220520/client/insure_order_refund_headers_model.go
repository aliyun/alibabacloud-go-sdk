// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInsureOrderRefundHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *InsureOrderRefundHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripCorpToken(v string) *InsureOrderRefundHeaders
	GetXAcsBtripCorpToken() *string
}

type InsureOrderRefundHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// feth00jqwls
	XAcsBtripCorpToken *string `json:"x-acs-btrip-corp-token,omitempty" xml:"x-acs-btrip-corp-token,omitempty"`
}

func (s InsureOrderRefundHeaders) String() string {
	return dara.Prettify(s)
}

func (s InsureOrderRefundHeaders) GoString() string {
	return s.String()
}

func (s *InsureOrderRefundHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *InsureOrderRefundHeaders) GetXAcsBtripCorpToken() *string {
	return s.XAcsBtripCorpToken
}

func (s *InsureOrderRefundHeaders) SetCommonHeaders(v map[string]*string) *InsureOrderRefundHeaders {
	s.CommonHeaders = v
	return s
}

func (s *InsureOrderRefundHeaders) SetXAcsBtripCorpToken(v string) *InsureOrderRefundHeaders {
	s.XAcsBtripCorpToken = &v
	return s
}

func (s *InsureOrderRefundHeaders) Validate() error {
	return dara.Validate(s)
}
