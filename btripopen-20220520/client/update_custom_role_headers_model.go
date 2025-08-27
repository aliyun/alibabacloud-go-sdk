// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCustomRoleHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *UpdateCustomRoleHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripCorpToken(v string) *UpdateCustomRoleHeaders
	GetXAcsBtripCorpToken() *string
}

type UpdateCustomRoleHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// feth00jqwis
	XAcsBtripCorpToken *string `json:"x-acs-btrip-corp-token,omitempty" xml:"x-acs-btrip-corp-token,omitempty"`
}

func (s UpdateCustomRoleHeaders) String() string {
	return dara.Prettify(s)
}

func (s UpdateCustomRoleHeaders) GoString() string {
	return s.String()
}

func (s *UpdateCustomRoleHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *UpdateCustomRoleHeaders) GetXAcsBtripCorpToken() *string {
	return s.XAcsBtripCorpToken
}

func (s *UpdateCustomRoleHeaders) SetCommonHeaders(v map[string]*string) *UpdateCustomRoleHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateCustomRoleHeaders) SetXAcsBtripCorpToken(v string) *UpdateCustomRoleHeaders {
	s.XAcsBtripCorpToken = &v
	return s
}

func (s *UpdateCustomRoleHeaders) Validate() error {
	return dara.Validate(s)
}
