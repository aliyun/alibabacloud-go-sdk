// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMigrationsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMigrationType(v string) *ListMigrationsRequest
	GetMigrationType() *string
	SetOwner(v string) *ListMigrationsRequest
	GetOwner() *string
	SetPageNumber(v int32) *ListMigrationsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListMigrationsRequest
	GetPageSize() *int32
	SetProjectId(v int64) *ListMigrationsRequest
	GetProjectId() *int64
}

type ListMigrationsRequest struct {
	// The migration task type. Valid values: IMPORT and EXPORT.
	//
	// This parameter is required.
	//
	// if can be null:
	// false
	//
	// example:
	//
	// IMPORT
	MigrationType *string `json:"MigrationType,omitempty" xml:"MigrationType,omitempty"`
	// The owner ID.
	//
	// example:
	//
	// 193379****
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The page number. Default value: 1. Maximum value: 100.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 10. Maximum value: 100.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The DataWorks workspace ID. You can log on to the [DataWorks console](https://workbench.data.aliyun.com/console) and go to the Workspace page to obtain the workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12345
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s ListMigrationsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListMigrationsRequest) GoString() string {
	return s.String()
}

func (s *ListMigrationsRequest) GetMigrationType() *string {
	return s.MigrationType
}

func (s *ListMigrationsRequest) GetOwner() *string {
	return s.Owner
}

func (s *ListMigrationsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListMigrationsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListMigrationsRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListMigrationsRequest) SetMigrationType(v string) *ListMigrationsRequest {
	s.MigrationType = &v
	return s
}

func (s *ListMigrationsRequest) SetOwner(v string) *ListMigrationsRequest {
	s.Owner = &v
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

func (s *ListMigrationsRequest) SetProjectId(v int64) *ListMigrationsRequest {
	s.ProjectId = &v
	return s
}

func (s *ListMigrationsRequest) Validate() error {
	return dara.Validate(s)
}
