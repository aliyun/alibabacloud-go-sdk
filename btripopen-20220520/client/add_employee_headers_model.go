// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddEmployeeHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *AddEmployeeHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripCorpToken(v string) *AddEmployeeHeaders
	GetXAcsBtripCorpToken() *string
}

type AddEmployeeHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// feth00jqwls
	XAcsBtripCorpToken *string `json:"x-acs-btrip-corp-token,omitempty" xml:"x-acs-btrip-corp-token,omitempty"`
}

func (s AddEmployeeHeaders) String() string {
	return dara.Prettify(s)
}

func (s AddEmployeeHeaders) GoString() string {
	return s.String()
}

func (s *AddEmployeeHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *AddEmployeeHeaders) GetXAcsBtripCorpToken() *string {
	return s.XAcsBtripCorpToken
}

func (s *AddEmployeeHeaders) SetCommonHeaders(v map[string]*string) *AddEmployeeHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AddEmployeeHeaders) SetXAcsBtripCorpToken(v string) *AddEmployeeHeaders {
	s.XAcsBtripCorpToken = &v
	return s
}

func (s *AddEmployeeHeaders) Validate() error {
	return dara.Validate(s)
}
