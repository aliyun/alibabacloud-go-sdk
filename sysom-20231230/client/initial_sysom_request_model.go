// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInitialSysomRequest interface {
	dara.Model
	String() string
	GoString() string
	SetXDebugId(v string) *InitialSysomRequest
	GetXDebugId() *string
	SetCheckOnly(v bool) *InitialSysomRequest
	GetCheckOnly() *bool
	SetSource(v string) *InitialSysomRequest
	GetSource() *string
	SetXSysomInvokeSource(v string) *InitialSysomRequest
	GetXSysomInvokeSource() *string
}

type InitialSysomRequest struct {
	XDebugId *string `json:"X-Debug-Id,omitempty" xml:"X-Debug-Id,omitempty"`
	// Specifies whether to only check if the service-linked role exists.
	CheckOnly *bool `json:"check_only,omitempty" xml:"check_only,omitempty"`
	// The source. Set this parameter to console.
	//
	// example:
	//
	// console
	Source             *string `json:"source,omitempty" xml:"source,omitempty"`
	XSysomInvokeSource *string `json:"x-sysom-invoke-source,omitempty" xml:"x-sysom-invoke-source,omitempty"`
}

func (s InitialSysomRequest) String() string {
	return dara.Prettify(s)
}

func (s InitialSysomRequest) GoString() string {
	return s.String()
}

func (s *InitialSysomRequest) GetXDebugId() *string {
	return s.XDebugId
}

func (s *InitialSysomRequest) GetCheckOnly() *bool {
	return s.CheckOnly
}

func (s *InitialSysomRequest) GetSource() *string {
	return s.Source
}

func (s *InitialSysomRequest) GetXSysomInvokeSource() *string {
	return s.XSysomInvokeSource
}

func (s *InitialSysomRequest) SetXDebugId(v string) *InitialSysomRequest {
	s.XDebugId = &v
	return s
}

func (s *InitialSysomRequest) SetCheckOnly(v bool) *InitialSysomRequest {
	s.CheckOnly = &v
	return s
}

func (s *InitialSysomRequest) SetSource(v string) *InitialSysomRequest {
	s.Source = &v
	return s
}

func (s *InitialSysomRequest) SetXSysomInvokeSource(v string) *InitialSysomRequest {
	s.XSysomInvokeSource = &v
	return s
}

func (s *InitialSysomRequest) Validate() error {
	return dara.Validate(s)
}
