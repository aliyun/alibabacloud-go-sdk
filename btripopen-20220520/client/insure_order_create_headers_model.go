// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInsureOrderCreateHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *InsureOrderCreateHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripCorpToken(v string) *InsureOrderCreateHeaders
	GetXAcsBtripCorpToken() *string
}

type InsureOrderCreateHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// feth00jqwls
	XAcsBtripCorpToken *string `json:"x-acs-btrip-corp-token,omitempty" xml:"x-acs-btrip-corp-token,omitempty"`
}

func (s InsureOrderCreateHeaders) String() string {
	return dara.Prettify(s)
}

func (s InsureOrderCreateHeaders) GoString() string {
	return s.String()
}

func (s *InsureOrderCreateHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *InsureOrderCreateHeaders) GetXAcsBtripCorpToken() *string {
	return s.XAcsBtripCorpToken
}

func (s *InsureOrderCreateHeaders) SetCommonHeaders(v map[string]*string) *InsureOrderCreateHeaders {
	s.CommonHeaders = v
	return s
}

func (s *InsureOrderCreateHeaders) SetXAcsBtripCorpToken(v string) *InsureOrderCreateHeaders {
	s.XAcsBtripCorpToken = &v
	return s
}

func (s *InsureOrderCreateHeaders) Validate() error {
	return dara.Validate(s)
}
