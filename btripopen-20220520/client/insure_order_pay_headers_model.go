// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInsureOrderPayHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *InsureOrderPayHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripCorpToken(v string) *InsureOrderPayHeaders
	GetXAcsBtripCorpToken() *string
}

type InsureOrderPayHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// feth00jqwls
	XAcsBtripCorpToken *string `json:"x-acs-btrip-corp-token,omitempty" xml:"x-acs-btrip-corp-token,omitempty"`
}

func (s InsureOrderPayHeaders) String() string {
	return dara.Prettify(s)
}

func (s InsureOrderPayHeaders) GoString() string {
	return s.String()
}

func (s *InsureOrderPayHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *InsureOrderPayHeaders) GetXAcsBtripCorpToken() *string {
	return s.XAcsBtripCorpToken
}

func (s *InsureOrderPayHeaders) SetCommonHeaders(v map[string]*string) *InsureOrderPayHeaders {
	s.CommonHeaders = v
	return s
}

func (s *InsureOrderPayHeaders) SetXAcsBtripCorpToken(v string) *InsureOrderPayHeaders {
	s.XAcsBtripCorpToken = &v
	return s
}

func (s *InsureOrderPayHeaders) Validate() error {
	return dara.Validate(s)
}
