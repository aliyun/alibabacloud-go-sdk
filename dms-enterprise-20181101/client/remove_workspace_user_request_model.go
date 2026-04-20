// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveWorkspaceUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDmsUserIds(v string) *RemoveWorkspaceUserRequest
	GetDmsUserIds() *string
	SetWorkspaceId(v string) *RemoveWorkspaceUserRequest
	GetWorkspaceId() *string
}

type RemoveWorkspaceUserRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 123***
	DmsUserIds *string `json:"DmsUserIds,omitempty" xml:"DmsUserIds,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 12****
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s RemoveWorkspaceUserRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveWorkspaceUserRequest) GoString() string {
	return s.String()
}

func (s *RemoveWorkspaceUserRequest) GetDmsUserIds() *string {
	return s.DmsUserIds
}

func (s *RemoveWorkspaceUserRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *RemoveWorkspaceUserRequest) SetDmsUserIds(v string) *RemoveWorkspaceUserRequest {
	s.DmsUserIds = &v
	return s
}

func (s *RemoveWorkspaceUserRequest) SetWorkspaceId(v string) *RemoveWorkspaceUserRequest {
	s.WorkspaceId = &v
	return s
}

func (s *RemoveWorkspaceUserRequest) Validate() error {
	return dara.Validate(s)
}
