// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVariableHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *DeleteVariableHeaders
	GetCommonHeaders() map[string]*string
	SetWorkspace(v string) *DeleteVariableHeaders
	GetWorkspace() *string
}

type DeleteVariableHeaders struct {
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

func (s DeleteVariableHeaders) String() string {
	return dara.Prettify(s)
}

func (s DeleteVariableHeaders) GoString() string {
	return s.String()
}

func (s *DeleteVariableHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *DeleteVariableHeaders) GetWorkspace() *string {
	return s.Workspace
}

func (s *DeleteVariableHeaders) SetCommonHeaders(v map[string]*string) *DeleteVariableHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteVariableHeaders) SetWorkspace(v string) *DeleteVariableHeaders {
	s.Workspace = &v
	return s
}

func (s *DeleteVariableHeaders) Validate() error {
	return dara.Validate(s)
}
