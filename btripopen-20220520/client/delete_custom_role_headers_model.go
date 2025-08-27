// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCustomRoleHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *DeleteCustomRoleHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripCorpToken(v string) *DeleteCustomRoleHeaders
	GetXAcsBtripCorpToken() *string
}

type DeleteCustomRoleHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// feth00jqwls
	XAcsBtripCorpToken *string `json:"x-acs-btrip-corp-token,omitempty" xml:"x-acs-btrip-corp-token,omitempty"`
}

func (s DeleteCustomRoleHeaders) String() string {
	return dara.Prettify(s)
}

func (s DeleteCustomRoleHeaders) GoString() string {
	return s.String()
}

func (s *DeleteCustomRoleHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *DeleteCustomRoleHeaders) GetXAcsBtripCorpToken() *string {
	return s.XAcsBtripCorpToken
}

func (s *DeleteCustomRoleHeaders) SetCommonHeaders(v map[string]*string) *DeleteCustomRoleHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteCustomRoleHeaders) SetXAcsBtripCorpToken(v string) *DeleteCustomRoleHeaders {
	s.XAcsBtripCorpToken = &v
	return s
}

func (s *DeleteCustomRoleHeaders) Validate() error {
	return dara.Validate(s)
}
