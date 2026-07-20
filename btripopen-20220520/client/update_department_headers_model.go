// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDepartmentHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *UpdateDepartmentHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripCorpToken(v string) *UpdateDepartmentHeaders
	GetXAcsBtripCorpToken() *string
}

type UpdateDepartmentHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// feth00jqwis
	XAcsBtripCorpToken *string `json:"x-acs-btrip-corp-token,omitempty" xml:"x-acs-btrip-corp-token,omitempty"`
}

func (s UpdateDepartmentHeaders) String() string {
	return dara.Prettify(s)
}

func (s UpdateDepartmentHeaders) GoString() string {
	return s.String()
}

func (s *UpdateDepartmentHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *UpdateDepartmentHeaders) GetXAcsBtripCorpToken() *string {
	return s.XAcsBtripCorpToken
}

func (s *UpdateDepartmentHeaders) SetCommonHeaders(v map[string]*string) *UpdateDepartmentHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateDepartmentHeaders) SetXAcsBtripCorpToken(v string) *UpdateDepartmentHeaders {
	s.XAcsBtripCorpToken = &v
	return s
}

func (s *UpdateDepartmentHeaders) Validate() error {
	return dara.Validate(s)
}
