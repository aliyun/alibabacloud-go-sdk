// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApplyExternalNodeStatusUpdateHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *ApplyExternalNodeStatusUpdateHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripCorpToken(v string) *ApplyExternalNodeStatusUpdateHeaders
	GetXAcsBtripCorpToken() *string
}

type ApplyExternalNodeStatusUpdateHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// feth00jqwls
	XAcsBtripCorpToken *string `json:"x-acs-btrip-corp-token,omitempty" xml:"x-acs-btrip-corp-token,omitempty"`
}

func (s ApplyExternalNodeStatusUpdateHeaders) String() string {
	return dara.Prettify(s)
}

func (s ApplyExternalNodeStatusUpdateHeaders) GoString() string {
	return s.String()
}

func (s *ApplyExternalNodeStatusUpdateHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *ApplyExternalNodeStatusUpdateHeaders) GetXAcsBtripCorpToken() *string {
	return s.XAcsBtripCorpToken
}

func (s *ApplyExternalNodeStatusUpdateHeaders) SetCommonHeaders(v map[string]*string) *ApplyExternalNodeStatusUpdateHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ApplyExternalNodeStatusUpdateHeaders) SetXAcsBtripCorpToken(v string) *ApplyExternalNodeStatusUpdateHeaders {
	s.XAcsBtripCorpToken = &v
	return s
}

func (s *ApplyExternalNodeStatusUpdateHeaders) Validate() error {
	return dara.Validate(s)
}
