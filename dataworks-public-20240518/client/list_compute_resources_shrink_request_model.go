// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListComputeResourcesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnvType(v string) *ListComputeResourcesShrinkRequest
	GetEnvType() *string
	SetName(v string) *ListComputeResourcesShrinkRequest
	GetName() *string
	SetOrder(v string) *ListComputeResourcesShrinkRequest
	GetOrder() *string
	SetPageNumber(v int32) *ListComputeResourcesShrinkRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListComputeResourcesShrinkRequest
	GetPageSize() *int32
	SetProjectId(v int64) *ListComputeResourcesShrinkRequest
	GetProjectId() *int64
	SetSortBy(v string) *ListComputeResourcesShrinkRequest
	GetSortBy() *string
	SetTypesShrink(v string) *ListComputeResourcesShrinkRequest
	GetTypesShrink() *string
}

type ListComputeResourcesShrinkRequest struct {
	EnvType    *string `json:"EnvType,omitempty" xml:"EnvType,omitempty"`
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Order      *string `json:"Order,omitempty" xml:"Order,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	ProjectId   *int64  `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	SortBy      *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	TypesShrink *string `json:"Types,omitempty" xml:"Types,omitempty"`
}

func (s ListComputeResourcesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListComputeResourcesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListComputeResourcesShrinkRequest) GetEnvType() *string {
	return s.EnvType
}

func (s *ListComputeResourcesShrinkRequest) GetName() *string {
	return s.Name
}

func (s *ListComputeResourcesShrinkRequest) GetOrder() *string {
	return s.Order
}

func (s *ListComputeResourcesShrinkRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListComputeResourcesShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListComputeResourcesShrinkRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListComputeResourcesShrinkRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListComputeResourcesShrinkRequest) GetTypesShrink() *string {
	return s.TypesShrink
}

func (s *ListComputeResourcesShrinkRequest) SetEnvType(v string) *ListComputeResourcesShrinkRequest {
	s.EnvType = &v
	return s
}

func (s *ListComputeResourcesShrinkRequest) SetName(v string) *ListComputeResourcesShrinkRequest {
	s.Name = &v
	return s
}

func (s *ListComputeResourcesShrinkRequest) SetOrder(v string) *ListComputeResourcesShrinkRequest {
	s.Order = &v
	return s
}

func (s *ListComputeResourcesShrinkRequest) SetPageNumber(v int32) *ListComputeResourcesShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListComputeResourcesShrinkRequest) SetPageSize(v int32) *ListComputeResourcesShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListComputeResourcesShrinkRequest) SetProjectId(v int64) *ListComputeResourcesShrinkRequest {
	s.ProjectId = &v
	return s
}

func (s *ListComputeResourcesShrinkRequest) SetSortBy(v string) *ListComputeResourcesShrinkRequest {
	s.SortBy = &v
	return s
}

func (s *ListComputeResourcesShrinkRequest) SetTypesShrink(v string) *ListComputeResourcesShrinkRequest {
	s.TypesShrink = &v
	return s
}

func (s *ListComputeResourcesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
