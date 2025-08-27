// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddDepartmentHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *AddDepartmentHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripCorpToken(v string) *AddDepartmentHeaders
	GetXAcsBtripCorpToken() *string
}

type AddDepartmentHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// feth00jqwls
	XAcsBtripCorpToken *string `json:"x-acs-btrip-corp-token,omitempty" xml:"x-acs-btrip-corp-token,omitempty"`
}

func (s AddDepartmentHeaders) String() string {
	return dara.Prettify(s)
}

func (s AddDepartmentHeaders) GoString() string {
	return s.String()
}

func (s *AddDepartmentHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *AddDepartmentHeaders) GetXAcsBtripCorpToken() *string {
	return s.XAcsBtripCorpToken
}

func (s *AddDepartmentHeaders) SetCommonHeaders(v map[string]*string) *AddDepartmentHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AddDepartmentHeaders) SetXAcsBtripCorpToken(v string) *AddDepartmentHeaders {
	s.XAcsBtripCorpToken = &v
	return s
}

func (s *AddDepartmentHeaders) Validate() error {
	return dara.Validate(s)
}
