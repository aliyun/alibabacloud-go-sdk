// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryReimbursementOrderHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *QueryReimbursementOrderHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripCorpToken(v string) *QueryReimbursementOrderHeaders
	GetXAcsBtripCorpToken() *string
}

type QueryReimbursementOrderHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// feth00jqwls
	XAcsBtripCorpToken *string `json:"x-acs-btrip-corp-token,omitempty" xml:"x-acs-btrip-corp-token,omitempty"`
}

func (s QueryReimbursementOrderHeaders) String() string {
	return dara.Prettify(s)
}

func (s QueryReimbursementOrderHeaders) GoString() string {
	return s.String()
}

func (s *QueryReimbursementOrderHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *QueryReimbursementOrderHeaders) GetXAcsBtripCorpToken() *string {
	return s.XAcsBtripCorpToken
}

func (s *QueryReimbursementOrderHeaders) SetCommonHeaders(v map[string]*string) *QueryReimbursementOrderHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryReimbursementOrderHeaders) SetXAcsBtripCorpToken(v string) *QueryReimbursementOrderHeaders {
	s.XAcsBtripCorpToken = &v
	return s
}

func (s *QueryReimbursementOrderHeaders) Validate() error {
	return dara.Validate(s)
}
