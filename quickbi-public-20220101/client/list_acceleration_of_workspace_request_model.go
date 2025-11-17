// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAccelerationOfWorkspaceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCreatorId(v string) *ListAccelerationOfWorkspaceRequest
	GetCreatorId() *string
	SetCubeName(v string) *ListAccelerationOfWorkspaceRequest
	GetCubeName() *string
	SetPageNo(v int32) *ListAccelerationOfWorkspaceRequest
	GetPageNo() *int32
	SetPageSize(v int32) *ListAccelerationOfWorkspaceRequest
	GetPageSize() *int32
	SetWorkspaceId(v string) *ListAccelerationOfWorkspaceRequest
	GetWorkspaceId() *string
}

type ListAccelerationOfWorkspaceRequest struct {
	// example:
	//
	// 1***************139
	CreatorId *string `json:"CreatorId,omitempty" xml:"CreatorId,omitempty"`
	// example:
	//
	// test
	CubeName *string `json:"CubeName,omitempty" xml:"CubeName,omitempty"`
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 95296e95-ca89-4c7d-8af9-dedf0ad0****
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListAccelerationOfWorkspaceRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAccelerationOfWorkspaceRequest) GoString() string {
	return s.String()
}

func (s *ListAccelerationOfWorkspaceRequest) GetCreatorId() *string {
	return s.CreatorId
}

func (s *ListAccelerationOfWorkspaceRequest) GetCubeName() *string {
	return s.CubeName
}

func (s *ListAccelerationOfWorkspaceRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *ListAccelerationOfWorkspaceRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAccelerationOfWorkspaceRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListAccelerationOfWorkspaceRequest) SetCreatorId(v string) *ListAccelerationOfWorkspaceRequest {
	s.CreatorId = &v
	return s
}

func (s *ListAccelerationOfWorkspaceRequest) SetCubeName(v string) *ListAccelerationOfWorkspaceRequest {
	s.CubeName = &v
	return s
}

func (s *ListAccelerationOfWorkspaceRequest) SetPageNo(v int32) *ListAccelerationOfWorkspaceRequest {
	s.PageNo = &v
	return s
}

func (s *ListAccelerationOfWorkspaceRequest) SetPageSize(v int32) *ListAccelerationOfWorkspaceRequest {
	s.PageSize = &v
	return s
}

func (s *ListAccelerationOfWorkspaceRequest) SetWorkspaceId(v string) *ListAccelerationOfWorkspaceRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ListAccelerationOfWorkspaceRequest) Validate() error {
	return dara.Validate(s)
}
