// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMigrationsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFilter(v string) *ListMigrationsRequest
	GetFilter() *string
	SetMigrationType(v string) *ListMigrationsRequest
	GetMigrationType() *string
	SetPageNumber(v int32) *ListMigrationsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListMigrationsRequest
	GetPageSize() *int32
	SetResourceGroupId(v string) *ListMigrationsRequest
	GetResourceGroupId() *string
}

type ListMigrationsRequest struct {
	// example:
	//
	// xx
	Filter *string `json:"filter,omitempty" xml:"filter,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// RESOURCE_IMPORT
	MigrationType *string `json:"migrationType,omitempty" xml:"migrationType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize        *int32  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	ResourceGroupId *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
}

func (s ListMigrationsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListMigrationsRequest) GoString() string {
	return s.String()
}

func (s *ListMigrationsRequest) GetFilter() *string {
	return s.Filter
}

func (s *ListMigrationsRequest) GetMigrationType() *string {
	return s.MigrationType
}

func (s *ListMigrationsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListMigrationsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListMigrationsRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListMigrationsRequest) SetFilter(v string) *ListMigrationsRequest {
	s.Filter = &v
	return s
}

func (s *ListMigrationsRequest) SetMigrationType(v string) *ListMigrationsRequest {
	s.MigrationType = &v
	return s
}

func (s *ListMigrationsRequest) SetPageNumber(v int32) *ListMigrationsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListMigrationsRequest) SetPageSize(v int32) *ListMigrationsRequest {
	s.PageSize = &v
	return s
}

func (s *ListMigrationsRequest) SetResourceGroupId(v string) *ListMigrationsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListMigrationsRequest) Validate() error {
	return dara.Validate(s)
}
