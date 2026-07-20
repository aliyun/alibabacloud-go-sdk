// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOrderRefundDetailQueryHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *OrderRefundDetailQueryHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripCorpToken(v string) *OrderRefundDetailQueryHeaders
	GetXAcsBtripCorpToken() *string
}

type OrderRefundDetailQueryHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// feth00jqwls
	XAcsBtripCorpToken *string `json:"x-acs-btrip-corp-token,omitempty" xml:"x-acs-btrip-corp-token,omitempty"`
}

func (s OrderRefundDetailQueryHeaders) String() string {
	return dara.Prettify(s)
}

func (s OrderRefundDetailQueryHeaders) GoString() string {
	return s.String()
}

func (s *OrderRefundDetailQueryHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *OrderRefundDetailQueryHeaders) GetXAcsBtripCorpToken() *string {
	return s.XAcsBtripCorpToken
}

func (s *OrderRefundDetailQueryHeaders) SetCommonHeaders(v map[string]*string) *OrderRefundDetailQueryHeaders {
	s.CommonHeaders = v
	return s
}

func (s *OrderRefundDetailQueryHeaders) SetXAcsBtripCorpToken(v string) *OrderRefundDetailQueryHeaders {
	s.XAcsBtripCorpToken = &v
	return s
}

func (s *OrderRefundDetailQueryHeaders) Validate() error {
	return dara.Validate(s)
}
