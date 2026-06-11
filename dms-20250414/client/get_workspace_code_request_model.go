// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWorkspaceCodeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIac(v string) *GetWorkspaceCodeRequest
	GetIac() *string
	SetPath(v string) *GetWorkspaceCodeRequest
	GetPath() *string
	SetWorkspaceId(v string) *GetWorkspaceCodeRequest
	GetWorkspaceId() *string
}

type GetWorkspaceCodeRequest struct {
	// If the file is a JSON file and Iac is set to true, the returned content is converted from JSON format to YAML format.
	//
	// example:
	//
	// false
	Iac *string `json:"Iac,omitempty" xml:"Iac,omitempty"`
	// The code file path: /Workspace/code/test.py
	//
	// Request path.
	//
	// This parameter is required.
	//
	// example:
	//
	// /Workspace/code/default/test.ipynb
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// The workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12****
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s GetWorkspaceCodeRequest) String() string {
	return dara.Prettify(s)
}

func (s GetWorkspaceCodeRequest) GoString() string {
	return s.String()
}

func (s *GetWorkspaceCodeRequest) GetIac() *string {
	return s.Iac
}

func (s *GetWorkspaceCodeRequest) GetPath() *string {
	return s.Path
}

func (s *GetWorkspaceCodeRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GetWorkspaceCodeRequest) SetIac(v string) *GetWorkspaceCodeRequest {
	s.Iac = &v
	return s
}

func (s *GetWorkspaceCodeRequest) SetPath(v string) *GetWorkspaceCodeRequest {
	s.Path = &v
	return s
}

func (s *GetWorkspaceCodeRequest) SetWorkspaceId(v string) *GetWorkspaceCodeRequest {
	s.WorkspaceId = &v
	return s
}

func (s *GetWorkspaceCodeRequest) Validate() error {
	return dara.Validate(s)
}
