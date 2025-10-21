// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRegisterUdfFunctionHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *RegisterUdfFunctionHeaders
	GetCommonHeaders() map[string]*string
	SetWorkspace(v string) *RegisterUdfFunctionHeaders
	GetWorkspace() *string
}

type RegisterUdfFunctionHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// The workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// a14bd5d90a****
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s RegisterUdfFunctionHeaders) String() string {
	return dara.Prettify(s)
}

func (s RegisterUdfFunctionHeaders) GoString() string {
	return s.String()
}

func (s *RegisterUdfFunctionHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *RegisterUdfFunctionHeaders) GetWorkspace() *string {
	return s.Workspace
}

func (s *RegisterUdfFunctionHeaders) SetCommonHeaders(v map[string]*string) *RegisterUdfFunctionHeaders {
	s.CommonHeaders = v
	return s
}

func (s *RegisterUdfFunctionHeaders) SetWorkspace(v string) *RegisterUdfFunctionHeaders {
	s.Workspace = &v
	return s
}

func (s *RegisterUdfFunctionHeaders) Validate() error {
	return dara.Validate(s)
}
