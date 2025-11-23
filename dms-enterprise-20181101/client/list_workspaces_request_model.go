// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWorkspacesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlreadyJoined(v bool) *ListWorkspacesRequest
	GetAlreadyJoined() *bool
	SetOwnerId(v int64) *ListWorkspacesRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *ListWorkspacesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListWorkspacesRequest
	GetPageSize() *int32
	SetRegion(v string) *ListWorkspacesRequest
	GetRegion() *string
	SetSearchKey(v string) *ListWorkspacesRequest
	GetSearchKey() *string
	SetServiceAccountId(v int64) *ListWorkspacesRequest
	GetServiceAccountId() *int64
	SetVpcId(v string) *ListWorkspacesRequest
	GetVpcId() *string
	SetWorkspaceId(v int64) *ListWorkspacesRequest
	GetWorkspaceId() *int64
}

type ListWorkspacesRequest struct {
	// Specifies whether the current user has joined the workspace.
	//
	// example:
	//
	// true
	AlreadyJoined *bool  `json:"AlreadyJoined,omitempty" xml:"AlreadyJoined,omitempty"`
	OwnerId       *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region in which the bucket is located.
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The search keyword. Fuzzy match is supported.
	//
	// example:
	//
	// poc_test
	SearchKey *string `json:"SearchKey,omitempty" xml:"SearchKey,omitempty"`
	// The service account ID.
	//
	// example:
	//
	// 12345
	ServiceAccountId *int64 `json:"ServiceAccountId,omitempty" xml:"ServiceAccountId,omitempty"`
	// The virtual private cloud (VPC) ID.
	//
	// > This parameter cannot be used as a filter.
	//
	// example:
	//
	// vpc-bp10wnlcmor****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The workspace ID.
	//
	// example:
	//
	// 12****
	WorkspaceId *int64 `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListWorkspacesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListWorkspacesRequest) GoString() string {
	return s.String()
}

func (s *ListWorkspacesRequest) GetAlreadyJoined() *bool {
	return s.AlreadyJoined
}

func (s *ListWorkspacesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListWorkspacesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListWorkspacesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListWorkspacesRequest) GetRegion() *string {
	return s.Region
}

func (s *ListWorkspacesRequest) GetSearchKey() *string {
	return s.SearchKey
}

func (s *ListWorkspacesRequest) GetServiceAccountId() *int64 {
	return s.ServiceAccountId
}

func (s *ListWorkspacesRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *ListWorkspacesRequest) GetWorkspaceId() *int64 {
	return s.WorkspaceId
}

func (s *ListWorkspacesRequest) SetAlreadyJoined(v bool) *ListWorkspacesRequest {
	s.AlreadyJoined = &v
	return s
}

func (s *ListWorkspacesRequest) SetOwnerId(v int64) *ListWorkspacesRequest {
	s.OwnerId = &v
	return s
}

func (s *ListWorkspacesRequest) SetPageNumber(v int32) *ListWorkspacesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListWorkspacesRequest) SetPageSize(v int32) *ListWorkspacesRequest {
	s.PageSize = &v
	return s
}

func (s *ListWorkspacesRequest) SetRegion(v string) *ListWorkspacesRequest {
	s.Region = &v
	return s
}

func (s *ListWorkspacesRequest) SetSearchKey(v string) *ListWorkspacesRequest {
	s.SearchKey = &v
	return s
}

func (s *ListWorkspacesRequest) SetServiceAccountId(v int64) *ListWorkspacesRequest {
	s.ServiceAccountId = &v
	return s
}

func (s *ListWorkspacesRequest) SetVpcId(v string) *ListWorkspacesRequest {
	s.VpcId = &v
	return s
}

func (s *ListWorkspacesRequest) SetWorkspaceId(v int64) *ListWorkspacesRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ListWorkspacesRequest) Validate() error {
	return dara.Validate(s)
}
