// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDepartmentSaveHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *DepartmentSaveHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripSoCorpToken(v string) *DepartmentSaveHeaders
	GetXAcsBtripSoCorpToken() *string
}

type DepartmentSaveHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// feth00jqwls
	XAcsBtripSoCorpToken *string `json:"x-acs-btrip-so-corp-token,omitempty" xml:"x-acs-btrip-so-corp-token,omitempty"`
}

func (s DepartmentSaveHeaders) String() string {
	return dara.Prettify(s)
}

func (s DepartmentSaveHeaders) GoString() string {
	return s.String()
}

func (s *DepartmentSaveHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *DepartmentSaveHeaders) GetXAcsBtripSoCorpToken() *string {
	return s.XAcsBtripSoCorpToken
}

func (s *DepartmentSaveHeaders) SetCommonHeaders(v map[string]*string) *DepartmentSaveHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DepartmentSaveHeaders) SetXAcsBtripSoCorpToken(v string) *DepartmentSaveHeaders {
	s.XAcsBtripSoCorpToken = &v
	return s
}

func (s *DepartmentSaveHeaders) Validate() error {
	return dara.Validate(s)
}
