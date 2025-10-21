// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVariablesHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *ListVariablesHeaders
	GetCommonHeaders() map[string]*string
	SetWorkspace(v string) *ListVariablesHeaders
	GetWorkspace() *string
}

type ListVariablesHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// The workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// a14bda1c4a****
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s ListVariablesHeaders) String() string {
	return dara.Prettify(s)
}

func (s ListVariablesHeaders) GoString() string {
	return s.String()
}

func (s *ListVariablesHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *ListVariablesHeaders) GetWorkspace() *string {
	return s.Workspace
}

func (s *ListVariablesHeaders) SetCommonHeaders(v map[string]*string) *ListVariablesHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListVariablesHeaders) SetWorkspace(v string) *ListVariablesHeaders {
	s.Workspace = &v
	return s
}

func (s *ListVariablesHeaders) Validate() error {
	return dara.Validate(s)
}
