// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWorkspaceCodeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPath(v string) *ListWorkspaceCodeRequest
	GetPath() *string
	SetWorkspaceId(v string) *ListWorkspaceCodeRequest
	GetWorkspaceId() *string
}

type ListWorkspaceCodeRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// /Workspace/code/default/test.ipynb
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 56kv1pvl9uvt9**********bb
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListWorkspaceCodeRequest) String() string {
	return dara.Prettify(s)
}

func (s ListWorkspaceCodeRequest) GoString() string {
	return s.String()
}

func (s *ListWorkspaceCodeRequest) GetPath() *string {
	return s.Path
}

func (s *ListWorkspaceCodeRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListWorkspaceCodeRequest) SetPath(v string) *ListWorkspaceCodeRequest {
	s.Path = &v
	return s
}

func (s *ListWorkspaceCodeRequest) SetWorkspaceId(v string) *ListWorkspaceCodeRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ListWorkspaceCodeRequest) Validate() error {
	return dara.Validate(s)
}
