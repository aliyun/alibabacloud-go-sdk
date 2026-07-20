// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInsureOrderUrlDetailHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *InsureOrderUrlDetailHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripCorpToken(v string) *InsureOrderUrlDetailHeaders
	GetXAcsBtripCorpToken() *string
}

type InsureOrderUrlDetailHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// feth00jqwls
	XAcsBtripCorpToken *string `json:"x-acs-btrip-corp-token,omitempty" xml:"x-acs-btrip-corp-token,omitempty"`
}

func (s InsureOrderUrlDetailHeaders) String() string {
	return dara.Prettify(s)
}

func (s InsureOrderUrlDetailHeaders) GoString() string {
	return s.String()
}

func (s *InsureOrderUrlDetailHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *InsureOrderUrlDetailHeaders) GetXAcsBtripCorpToken() *string {
	return s.XAcsBtripCorpToken
}

func (s *InsureOrderUrlDetailHeaders) SetCommonHeaders(v map[string]*string) *InsureOrderUrlDetailHeaders {
	s.CommonHeaders = v
	return s
}

func (s *InsureOrderUrlDetailHeaders) SetXAcsBtripCorpToken(v string) *InsureOrderUrlDetailHeaders {
	s.XAcsBtripCorpToken = &v
	return s
}

func (s *InsureOrderUrlDetailHeaders) Validate() error {
	return dara.Validate(s)
}
