// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateEmployeeLeaveStatusHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *UpdateEmployeeLeaveStatusHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripCorpToken(v string) *UpdateEmployeeLeaveStatusHeaders
	GetXAcsBtripCorpToken() *string
}

type UpdateEmployeeLeaveStatusHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// feth00jqwls
	XAcsBtripCorpToken *string `json:"x-acs-btrip-corp-token,omitempty" xml:"x-acs-btrip-corp-token,omitempty"`
}

func (s UpdateEmployeeLeaveStatusHeaders) String() string {
	return dara.Prettify(s)
}

func (s UpdateEmployeeLeaveStatusHeaders) GoString() string {
	return s.String()
}

func (s *UpdateEmployeeLeaveStatusHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *UpdateEmployeeLeaveStatusHeaders) GetXAcsBtripCorpToken() *string {
	return s.XAcsBtripCorpToken
}

func (s *UpdateEmployeeLeaveStatusHeaders) SetCommonHeaders(v map[string]*string) *UpdateEmployeeLeaveStatusHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateEmployeeLeaveStatusHeaders) SetXAcsBtripCorpToken(v string) *UpdateEmployeeLeaveStatusHeaders {
	s.XAcsBtripCorpToken = &v
	return s
}

func (s *UpdateEmployeeLeaveStatusHeaders) Validate() error {
	return dara.Validate(s)
}
