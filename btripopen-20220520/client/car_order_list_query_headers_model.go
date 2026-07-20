// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCarOrderListQueryHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *CarOrderListQueryHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripSoCorpToken(v string) *CarOrderListQueryHeaders
	GetXAcsBtripSoCorpToken() *string
}

type CarOrderListQueryHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// feth00jqwls
	XAcsBtripSoCorpToken *string `json:"x-acs-btrip-so-corp-token,omitempty" xml:"x-acs-btrip-so-corp-token,omitempty"`
}

func (s CarOrderListQueryHeaders) String() string {
	return dara.Prettify(s)
}

func (s CarOrderListQueryHeaders) GoString() string {
	return s.String()
}

func (s *CarOrderListQueryHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *CarOrderListQueryHeaders) GetXAcsBtripSoCorpToken() *string {
	return s.XAcsBtripSoCorpToken
}

func (s *CarOrderListQueryHeaders) SetCommonHeaders(v map[string]*string) *CarOrderListQueryHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CarOrderListQueryHeaders) SetXAcsBtripSoCorpToken(v string) *CarOrderListQueryHeaders {
	s.XAcsBtripSoCorpToken = &v
	return s
}

func (s *CarOrderListQueryHeaders) Validate() error {
	return dara.Validate(s)
}
