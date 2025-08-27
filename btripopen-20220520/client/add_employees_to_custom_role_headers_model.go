// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddEmployeesToCustomRoleHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *AddEmployeesToCustomRoleHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripCorpToken(v string) *AddEmployeesToCustomRoleHeaders
	GetXAcsBtripCorpToken() *string
}

type AddEmployeesToCustomRoleHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// feth00jqwls
	XAcsBtripCorpToken *string `json:"x-acs-btrip-corp-token,omitempty" xml:"x-acs-btrip-corp-token,omitempty"`
}

func (s AddEmployeesToCustomRoleHeaders) String() string {
	return dara.Prettify(s)
}

func (s AddEmployeesToCustomRoleHeaders) GoString() string {
	return s.String()
}

func (s *AddEmployeesToCustomRoleHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *AddEmployeesToCustomRoleHeaders) GetXAcsBtripCorpToken() *string {
	return s.XAcsBtripCorpToken
}

func (s *AddEmployeesToCustomRoleHeaders) SetCommonHeaders(v map[string]*string) *AddEmployeesToCustomRoleHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AddEmployeesToCustomRoleHeaders) SetXAcsBtripCorpToken(v string) *AddEmployeesToCustomRoleHeaders {
	s.XAcsBtripCorpToken = &v
	return s
}

func (s *AddEmployeesToCustomRoleHeaders) Validate() error {
	return dara.Validate(s)
}
