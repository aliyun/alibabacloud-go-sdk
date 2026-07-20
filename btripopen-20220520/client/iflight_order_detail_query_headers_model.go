// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIFlightOrderDetailQueryHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *IFlightOrderDetailQueryHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripCorpToken(v string) *IFlightOrderDetailQueryHeaders
	GetXAcsBtripCorpToken() *string
}

type IFlightOrderDetailQueryHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// feth00jqwls
	XAcsBtripCorpToken *string `json:"x-acs-btrip-corp-token,omitempty" xml:"x-acs-btrip-corp-token,omitempty"`
}

func (s IFlightOrderDetailQueryHeaders) String() string {
	return dara.Prettify(s)
}

func (s IFlightOrderDetailQueryHeaders) GoString() string {
	return s.String()
}

func (s *IFlightOrderDetailQueryHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *IFlightOrderDetailQueryHeaders) GetXAcsBtripCorpToken() *string {
	return s.XAcsBtripCorpToken
}

func (s *IFlightOrderDetailQueryHeaders) SetCommonHeaders(v map[string]*string) *IFlightOrderDetailQueryHeaders {
	s.CommonHeaders = v
	return s
}

func (s *IFlightOrderDetailQueryHeaders) SetXAcsBtripCorpToken(v string) *IFlightOrderDetailQueryHeaders {
	s.XAcsBtripCorpToken = &v
	return s
}

func (s *IFlightOrderDetailQueryHeaders) Validate() error {
	return dara.Validate(s)
}
