// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVariableHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *CreateVariableHeaders
	GetCommonHeaders() map[string]*string
	SetWorkspace(v string) *CreateVariableHeaders
	GetWorkspace() *string
}

type CreateVariableHeaders struct {
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

func (s CreateVariableHeaders) String() string {
	return dara.Prettify(s)
}

func (s CreateVariableHeaders) GoString() string {
	return s.String()
}

func (s *CreateVariableHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *CreateVariableHeaders) GetWorkspace() *string {
	return s.Workspace
}

func (s *CreateVariableHeaders) SetCommonHeaders(v map[string]*string) *CreateVariableHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateVariableHeaders) SetWorkspace(v string) *CreateVariableHeaders {
	s.Workspace = &v
	return s
}

func (s *CreateVariableHeaders) Validate() error {
	return dara.Validate(s)
}
