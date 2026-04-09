// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDataPermissionsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIdsShrink(v string) *DeleteDataPermissionsShrinkRequest
	GetIdsShrink() *string
	SetWorkspaceId(v string) *DeleteDataPermissionsShrinkRequest
	GetWorkspaceId() *string
}

type DeleteDataPermissionsShrinkRequest struct {
	// This parameter is required.
	IdsShrink *string `json:"Ids,omitempty" xml:"Ids,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// llm-xxx
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s DeleteDataPermissionsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDataPermissionsShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeleteDataPermissionsShrinkRequest) GetIdsShrink() *string {
	return s.IdsShrink
}

func (s *DeleteDataPermissionsShrinkRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *DeleteDataPermissionsShrinkRequest) SetIdsShrink(v string) *DeleteDataPermissionsShrinkRequest {
	s.IdsShrink = &v
	return s
}

func (s *DeleteDataPermissionsShrinkRequest) SetWorkspaceId(v string) *DeleteDataPermissionsShrinkRequest {
	s.WorkspaceId = &v
	return s
}

func (s *DeleteDataPermissionsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
