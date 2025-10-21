// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateVariableHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *UpdateVariableHeaders
	GetCommonHeaders() map[string]*string
	SetWorkspace(v string) *UpdateVariableHeaders
	GetWorkspace() *string
}

type UpdateVariableHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// a14bd5d90a****
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s UpdateVariableHeaders) String() string {
	return dara.Prettify(s)
}

func (s UpdateVariableHeaders) GoString() string {
	return s.String()
}

func (s *UpdateVariableHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *UpdateVariableHeaders) GetWorkspace() *string {
	return s.Workspace
}

func (s *UpdateVariableHeaders) SetCommonHeaders(v map[string]*string) *UpdateVariableHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateVariableHeaders) SetWorkspace(v string) *UpdateVariableHeaders {
	s.Workspace = &v
	return s
}

func (s *UpdateVariableHeaders) Validate() error {
	return dara.Validate(s)
}
