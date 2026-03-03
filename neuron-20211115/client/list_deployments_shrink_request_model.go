// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDeploymentsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExcludeStatusShrink(v string) *ListDeploymentsShrinkRequest
	GetExcludeStatusShrink() *string
	SetOrderBy(v string) *ListDeploymentsShrinkRequest
	GetOrderBy() *string
	SetOrderDirection(v string) *ListDeploymentsShrinkRequest
	GetOrderDirection() *string
	SetPageNumber(v int32) *ListDeploymentsShrinkRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListDeploymentsShrinkRequest
	GetPageSize() *int32
	SetServiceGroupId(v int64) *ListDeploymentsShrinkRequest
	GetServiceGroupId() *int64
	SetStatusShrink(v string) *ListDeploymentsShrinkRequest
	GetStatusShrink() *string
}

type ListDeploymentsShrinkRequest struct {
	ExcludeStatusShrink *string `json:"excludeStatus,omitempty" xml:"excludeStatus,omitempty"`
	// example:
	//
	// gmtCreated
	OrderBy *string `json:"orderBy,omitempty" xml:"orderBy,omitempty"`
	// example:
	//
	// desc
	OrderDirection *string `json:"orderDirection,omitempty" xml:"orderDirection,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	ServiceGroupId *int64  `json:"serviceGroupId,omitempty" xml:"serviceGroupId,omitempty"`
	StatusShrink   *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s ListDeploymentsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDeploymentsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListDeploymentsShrinkRequest) GetExcludeStatusShrink() *string {
	return s.ExcludeStatusShrink
}

func (s *ListDeploymentsShrinkRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *ListDeploymentsShrinkRequest) GetOrderDirection() *string {
	return s.OrderDirection
}

func (s *ListDeploymentsShrinkRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListDeploymentsShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDeploymentsShrinkRequest) GetServiceGroupId() *int64 {
	return s.ServiceGroupId
}

func (s *ListDeploymentsShrinkRequest) GetStatusShrink() *string {
	return s.StatusShrink
}

func (s *ListDeploymentsShrinkRequest) SetExcludeStatusShrink(v string) *ListDeploymentsShrinkRequest {
	s.ExcludeStatusShrink = &v
	return s
}

func (s *ListDeploymentsShrinkRequest) SetOrderBy(v string) *ListDeploymentsShrinkRequest {
	s.OrderBy = &v
	return s
}

func (s *ListDeploymentsShrinkRequest) SetOrderDirection(v string) *ListDeploymentsShrinkRequest {
	s.OrderDirection = &v
	return s
}

func (s *ListDeploymentsShrinkRequest) SetPageNumber(v int32) *ListDeploymentsShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListDeploymentsShrinkRequest) SetPageSize(v int32) *ListDeploymentsShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListDeploymentsShrinkRequest) SetServiceGroupId(v int64) *ListDeploymentsShrinkRequest {
	s.ServiceGroupId = &v
	return s
}

func (s *ListDeploymentsShrinkRequest) SetStatusShrink(v string) *ListDeploymentsShrinkRequest {
	s.StatusShrink = &v
	return s
}

func (s *ListDeploymentsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
