// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGroupCorpTokenHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GroupCorpTokenHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripAccessToken(v string) *GroupCorpTokenHeaders
	GetXAcsBtripAccessToken() *string
}

type GroupCorpTokenHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// feth00jqwls
	XAcsBtripAccessToken *string `json:"x-acs-btrip-access-token,omitempty" xml:"x-acs-btrip-access-token,omitempty"`
}

func (s GroupCorpTokenHeaders) String() string {
	return dara.Prettify(s)
}

func (s GroupCorpTokenHeaders) GoString() string {
	return s.String()
}

func (s *GroupCorpTokenHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GroupCorpTokenHeaders) GetXAcsBtripAccessToken() *string {
	return s.XAcsBtripAccessToken
}

func (s *GroupCorpTokenHeaders) SetCommonHeaders(v map[string]*string) *GroupCorpTokenHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GroupCorpTokenHeaders) SetXAcsBtripAccessToken(v string) *GroupCorpTokenHeaders {
	s.XAcsBtripAccessToken = &v
	return s
}

func (s *GroupCorpTokenHeaders) Validate() error {
	return dara.Validate(s)
}
