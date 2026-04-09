// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDataPermissionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIds(v []*int64) *DeleteDataPermissionsRequest
	GetIds() []*int64
	SetWorkspaceId(v string) *DeleteDataPermissionsRequest
	GetWorkspaceId() *string
}

type DeleteDataPermissionsRequest struct {
	// This parameter is required.
	Ids []*int64 `json:"Ids,omitempty" xml:"Ids,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// llm-xxx
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s DeleteDataPermissionsRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDataPermissionsRequest) GoString() string {
	return s.String()
}

func (s *DeleteDataPermissionsRequest) GetIds() []*int64 {
	return s.Ids
}

func (s *DeleteDataPermissionsRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *DeleteDataPermissionsRequest) SetIds(v []*int64) *DeleteDataPermissionsRequest {
	s.Ids = v
	return s
}

func (s *DeleteDataPermissionsRequest) SetWorkspaceId(v string) *DeleteDataPermissionsRequest {
	s.WorkspaceId = &v
	return s
}

func (s *DeleteDataPermissionsRequest) Validate() error {
	return dara.Validate(s)
}
