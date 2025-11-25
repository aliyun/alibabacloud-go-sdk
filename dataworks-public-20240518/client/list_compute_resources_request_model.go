// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListComputeResourcesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnvType(v string) *ListComputeResourcesRequest
	GetEnvType() *string
	SetName(v string) *ListComputeResourcesRequest
	GetName() *string
	SetOrder(v string) *ListComputeResourcesRequest
	GetOrder() *string
	SetPageNumber(v int32) *ListComputeResourcesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListComputeResourcesRequest
	GetPageSize() *int32
	SetProjectId(v int64) *ListComputeResourcesRequest
	GetProjectId() *int64
	SetSortBy(v string) *ListComputeResourcesRequest
	GetSortBy() *string
	SetTypes(v []*string) *ListComputeResourcesRequest
	GetTypes() []*string
}

type ListComputeResourcesRequest struct {
	EnvType    *string `json:"EnvType,omitempty" xml:"EnvType,omitempty"`
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Order      *string `json:"Order,omitempty" xml:"Order,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	ProjectId *int64    `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	SortBy    *string   `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	Types     []*string `json:"Types,omitempty" xml:"Types,omitempty" type:"Repeated"`
}

func (s ListComputeResourcesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListComputeResourcesRequest) GoString() string {
	return s.String()
}

func (s *ListComputeResourcesRequest) GetEnvType() *string {
	return s.EnvType
}

func (s *ListComputeResourcesRequest) GetName() *string {
	return s.Name
}

func (s *ListComputeResourcesRequest) GetOrder() *string {
	return s.Order
}

func (s *ListComputeResourcesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListComputeResourcesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListComputeResourcesRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListComputeResourcesRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListComputeResourcesRequest) GetTypes() []*string {
	return s.Types
}

func (s *ListComputeResourcesRequest) SetEnvType(v string) *ListComputeResourcesRequest {
	s.EnvType = &v
	return s
}

func (s *ListComputeResourcesRequest) SetName(v string) *ListComputeResourcesRequest {
	s.Name = &v
	return s
}

func (s *ListComputeResourcesRequest) SetOrder(v string) *ListComputeResourcesRequest {
	s.Order = &v
	return s
}

func (s *ListComputeResourcesRequest) SetPageNumber(v int32) *ListComputeResourcesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListComputeResourcesRequest) SetPageSize(v int32) *ListComputeResourcesRequest {
	s.PageSize = &v
	return s
}

func (s *ListComputeResourcesRequest) SetProjectId(v int64) *ListComputeResourcesRequest {
	s.ProjectId = &v
	return s
}

func (s *ListComputeResourcesRequest) SetSortBy(v string) *ListComputeResourcesRequest {
	s.SortBy = &v
	return s
}

func (s *ListComputeResourcesRequest) SetTypes(v []*string) *ListComputeResourcesRequest {
	s.Types = v
	return s
}

func (s *ListComputeResourcesRequest) Validate() error {
	return dara.Validate(s)
}
