// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateEmployeeHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *UpdateEmployeeHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripCorpToken(v string) *UpdateEmployeeHeaders
	GetXAcsBtripCorpToken() *string
}

type UpdateEmployeeHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// feth00jqwls
	XAcsBtripCorpToken *string `json:"x-acs-btrip-corp-token,omitempty" xml:"x-acs-btrip-corp-token,omitempty"`
}

func (s UpdateEmployeeHeaders) String() string {
	return dara.Prettify(s)
}

func (s UpdateEmployeeHeaders) GoString() string {
	return s.String()
}

func (s *UpdateEmployeeHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *UpdateEmployeeHeaders) GetXAcsBtripCorpToken() *string {
	return s.XAcsBtripCorpToken
}

func (s *UpdateEmployeeHeaders) SetCommonHeaders(v map[string]*string) *UpdateEmployeeHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateEmployeeHeaders) SetXAcsBtripCorpToken(v string) *UpdateEmployeeHeaders {
	s.XAcsBtripCorpToken = &v
	return s
}

func (s *UpdateEmployeeHeaders) Validate() error {
	return dara.Validate(s)
}
