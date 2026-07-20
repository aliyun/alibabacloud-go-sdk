// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInsureOrderCancelHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *InsureOrderCancelHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripCorpToken(v string) *InsureOrderCancelHeaders
	GetXAcsBtripCorpToken() *string
}

type InsureOrderCancelHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// feth00jqwls
	XAcsBtripCorpToken *string `json:"x-acs-btrip-corp-token,omitempty" xml:"x-acs-btrip-corp-token,omitempty"`
}

func (s InsureOrderCancelHeaders) String() string {
	return dara.Prettify(s)
}

func (s InsureOrderCancelHeaders) GoString() string {
	return s.String()
}

func (s *InsureOrderCancelHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *InsureOrderCancelHeaders) GetXAcsBtripCorpToken() *string {
	return s.XAcsBtripCorpToken
}

func (s *InsureOrderCancelHeaders) SetCommonHeaders(v map[string]*string) *InsureOrderCancelHeaders {
	s.CommonHeaders = v
	return s
}

func (s *InsureOrderCancelHeaders) SetXAcsBtripCorpToken(v string) *InsureOrderCancelHeaders {
	s.XAcsBtripCorpToken = &v
	return s
}

func (s *InsureOrderCancelHeaders) Validate() error {
	return dara.Validate(s)
}
