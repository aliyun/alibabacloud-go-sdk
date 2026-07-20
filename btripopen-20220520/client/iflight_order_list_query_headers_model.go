// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIFlightOrderListQueryHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *IFlightOrderListQueryHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripCorpToken(v string) *IFlightOrderListQueryHeaders
	GetXAcsBtripCorpToken() *string
}

type IFlightOrderListQueryHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// feth00jqwls
	XAcsBtripCorpToken *string `json:"x-acs-btrip-corp-token,omitempty" xml:"x-acs-btrip-corp-token,omitempty"`
}

func (s IFlightOrderListQueryHeaders) String() string {
	return dara.Prettify(s)
}

func (s IFlightOrderListQueryHeaders) GoString() string {
	return s.String()
}

func (s *IFlightOrderListQueryHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *IFlightOrderListQueryHeaders) GetXAcsBtripCorpToken() *string {
	return s.XAcsBtripCorpToken
}

func (s *IFlightOrderListQueryHeaders) SetCommonHeaders(v map[string]*string) *IFlightOrderListQueryHeaders {
	s.CommonHeaders = v
	return s
}

func (s *IFlightOrderListQueryHeaders) SetXAcsBtripCorpToken(v string) *IFlightOrderListQueryHeaders {
	s.XAcsBtripCorpToken = &v
	return s
}

func (s *IFlightOrderListQueryHeaders) Validate() error {
	return dara.Validate(s)
}
