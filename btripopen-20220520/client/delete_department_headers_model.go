// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDepartmentHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *DeleteDepartmentHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripCorpToken(v string) *DeleteDepartmentHeaders
	GetXAcsBtripCorpToken() *string
}

type DeleteDepartmentHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// feth00jqwls
	XAcsBtripCorpToken *string `json:"x-acs-btrip-corp-token,omitempty" xml:"x-acs-btrip-corp-token,omitempty"`
}

func (s DeleteDepartmentHeaders) String() string {
	return dara.Prettify(s)
}

func (s DeleteDepartmentHeaders) GoString() string {
	return s.String()
}

func (s *DeleteDepartmentHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *DeleteDepartmentHeaders) GetXAcsBtripCorpToken() *string {
	return s.XAcsBtripCorpToken
}

func (s *DeleteDepartmentHeaders) SetCommonHeaders(v map[string]*string) *DeleteDepartmentHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteDepartmentHeaders) SetXAcsBtripCorpToken(v string) *DeleteDepartmentHeaders {
	s.XAcsBtripCorpToken = &v
	return s
}

func (s *DeleteDepartmentHeaders) Validate() error {
	return dara.Validate(s)
}
