// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEmployeesFromCustomRoleHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *DeleteEmployeesFromCustomRoleHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripCorpToken(v string) *DeleteEmployeesFromCustomRoleHeaders
	GetXAcsBtripCorpToken() *string
}

type DeleteEmployeesFromCustomRoleHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// feth00jqwls
	XAcsBtripCorpToken *string `json:"x-acs-btrip-corp-token,omitempty" xml:"x-acs-btrip-corp-token,omitempty"`
}

func (s DeleteEmployeesFromCustomRoleHeaders) String() string {
	return dara.Prettify(s)
}

func (s DeleteEmployeesFromCustomRoleHeaders) GoString() string {
	return s.String()
}

func (s *DeleteEmployeesFromCustomRoleHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *DeleteEmployeesFromCustomRoleHeaders) GetXAcsBtripCorpToken() *string {
	return s.XAcsBtripCorpToken
}

func (s *DeleteEmployeesFromCustomRoleHeaders) SetCommonHeaders(v map[string]*string) *DeleteEmployeesFromCustomRoleHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteEmployeesFromCustomRoleHeaders) SetXAcsBtripCorpToken(v string) *DeleteEmployeesFromCustomRoleHeaders {
	s.XAcsBtripCorpToken = &v
	return s
}

func (s *DeleteEmployeesFromCustomRoleHeaders) Validate() error {
	return dara.Validate(s)
}
