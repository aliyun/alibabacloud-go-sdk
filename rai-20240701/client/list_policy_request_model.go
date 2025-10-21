// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIsSidecarPolicy(v int32) *ListPolicyRequest
	GetIsSidecarPolicy() *int32
	SetOrder(v string) *ListPolicyRequest
	GetOrder() *string
	SetPageNumber(v int32) *ListPolicyRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListPolicyRequest
	GetPageSize() *int32
	SetPolicyIdentifier(v string) *ListPolicyRequest
	GetPolicyIdentifier() *string
	SetPolicyName(v string) *ListPolicyRequest
	GetPolicyName() *string
	SetRegionId(v string) *ListPolicyRequest
	GetRegionId() *string
	SetSceneType(v int32) *ListPolicyRequest
	GetSceneType() *int32
	SetSortBy(v string) *ListPolicyRequest
	GetSortBy() *string
	SetWorkspaceId(v int64) *ListPolicyRequest
	GetWorkspaceId() *int64
}

type ListPolicyRequest struct {
	IsSidecarPolicy *int32 `json:"IsSidecarPolicy,omitempty" xml:"IsSidecarPolicy,omitempty"`
	// Sort order. The value range is as follows:
	//
	// - asc (default): ascending
	//
	// - desc: descending
	//
	// example:
	//
	// asc
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// Page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// Page size, the maximum number of results returned per page.
	//
	// Maximum limit: 100.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Policy identifier.
	//
	// example:
	//
	// x1bc5xgs4uhx
	PolicyIdentifier *string `json:"PolicyIdentifier,omitempty" xml:"PolicyIdentifier,omitempty"`
	// Detection policy name.
	//
	// example:
	//
	// testPolicy
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	// Region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SceneType *int32  `json:"SceneType,omitempty" xml:"SceneType,omitempty"`
	// Sort field.
	//
	// example:
	//
	// GmtModified
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// Workspace ID.
	//
	// example:
	//
	// 62****
	WorkspaceId *int64 `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPolicyRequest) GoString() string {
	return s.String()
}

func (s *ListPolicyRequest) GetIsSidecarPolicy() *int32 {
	return s.IsSidecarPolicy
}

func (s *ListPolicyRequest) GetOrder() *string {
	return s.Order
}

func (s *ListPolicyRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListPolicyRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListPolicyRequest) GetPolicyIdentifier() *string {
	return s.PolicyIdentifier
}

func (s *ListPolicyRequest) GetPolicyName() *string {
	return s.PolicyName
}

func (s *ListPolicyRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListPolicyRequest) GetSceneType() *int32 {
	return s.SceneType
}

func (s *ListPolicyRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListPolicyRequest) GetWorkspaceId() *int64 {
	return s.WorkspaceId
}

func (s *ListPolicyRequest) SetIsSidecarPolicy(v int32) *ListPolicyRequest {
	s.IsSidecarPolicy = &v
	return s
}

func (s *ListPolicyRequest) SetOrder(v string) *ListPolicyRequest {
	s.Order = &v
	return s
}

func (s *ListPolicyRequest) SetPageNumber(v int32) *ListPolicyRequest {
	s.PageNumber = &v
	return s
}

func (s *ListPolicyRequest) SetPageSize(v int32) *ListPolicyRequest {
	s.PageSize = &v
	return s
}

func (s *ListPolicyRequest) SetPolicyIdentifier(v string) *ListPolicyRequest {
	s.PolicyIdentifier = &v
	return s
}

func (s *ListPolicyRequest) SetPolicyName(v string) *ListPolicyRequest {
	s.PolicyName = &v
	return s
}

func (s *ListPolicyRequest) SetRegionId(v string) *ListPolicyRequest {
	s.RegionId = &v
	return s
}

func (s *ListPolicyRequest) SetSceneType(v int32) *ListPolicyRequest {
	s.SceneType = &v
	return s
}

func (s *ListPolicyRequest) SetSortBy(v string) *ListPolicyRequest {
	s.SortBy = &v
	return s
}

func (s *ListPolicyRequest) SetWorkspaceId(v int64) *ListPolicyRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ListPolicyRequest) Validate() error {
	return dara.Validate(s)
}
