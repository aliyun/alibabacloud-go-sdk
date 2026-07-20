// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGroupUserSaveHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GroupUserSaveHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripCorpToken(v string) *GroupUserSaveHeaders
	GetXAcsBtripCorpToken() *string
}

type GroupUserSaveHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// feth00jqwls
	XAcsBtripCorpToken *string `json:"x-acs-btrip-corp-token,omitempty" xml:"x-acs-btrip-corp-token,omitempty"`
}

func (s GroupUserSaveHeaders) String() string {
	return dara.Prettify(s)
}

func (s GroupUserSaveHeaders) GoString() string {
	return s.String()
}

func (s *GroupUserSaveHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GroupUserSaveHeaders) GetXAcsBtripCorpToken() *string {
	return s.XAcsBtripCorpToken
}

func (s *GroupUserSaveHeaders) SetCommonHeaders(v map[string]*string) *GroupUserSaveHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GroupUserSaveHeaders) SetXAcsBtripCorpToken(v string) *GroupUserSaveHeaders {
	s.XAcsBtripCorpToken = &v
	return s
}

func (s *GroupUserSaveHeaders) Validate() error {
	return dara.Validate(s)
}
