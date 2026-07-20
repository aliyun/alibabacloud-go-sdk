// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInsureOrderDetailHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *InsureOrderDetailHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripCorpToken(v string) *InsureOrderDetailHeaders
	GetXAcsBtripCorpToken() *string
}

type InsureOrderDetailHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// feth00jqwls
	XAcsBtripCorpToken *string `json:"x-acs-btrip-corp-token,omitempty" xml:"x-acs-btrip-corp-token,omitempty"`
}

func (s InsureOrderDetailHeaders) String() string {
	return dara.Prettify(s)
}

func (s InsureOrderDetailHeaders) GoString() string {
	return s.String()
}

func (s *InsureOrderDetailHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *InsureOrderDetailHeaders) GetXAcsBtripCorpToken() *string {
	return s.XAcsBtripCorpToken
}

func (s *InsureOrderDetailHeaders) SetCommonHeaders(v map[string]*string) *InsureOrderDetailHeaders {
	s.CommonHeaders = v
	return s
}

func (s *InsureOrderDetailHeaders) SetXAcsBtripCorpToken(v string) *InsureOrderDetailHeaders {
	s.XAcsBtripCorpToken = &v
	return s
}

func (s *InsureOrderDetailHeaders) Validate() error {
	return dara.Validate(s)
}
