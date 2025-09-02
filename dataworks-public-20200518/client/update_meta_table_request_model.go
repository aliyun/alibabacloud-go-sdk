// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMetaTableRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddedLabels(v string) *UpdateMetaTableRequest
	GetAddedLabels() *string
	SetCaption(v string) *UpdateMetaTableRequest
	GetCaption() *string
	SetCategoryId(v int64) *UpdateMetaTableRequest
	GetCategoryId() *int64
	SetEnvType(v int32) *UpdateMetaTableRequest
	GetEnvType() *int32
	SetNewOwnerId(v string) *UpdateMetaTableRequest
	GetNewOwnerId() *string
	SetProjectId(v int64) *UpdateMetaTableRequest
	GetProjectId() *int64
	SetRemovedLabels(v string) *UpdateMetaTableRequest
	GetRemovedLabels() *string
	SetSchema(v string) *UpdateMetaTableRequest
	GetSchema() *string
	SetTableGuid(v string) *UpdateMetaTableRequest
	GetTableGuid() *string
	SetTableName(v string) *UpdateMetaTableRequest
	GetTableName() *string
	SetVisibility(v int32) *UpdateMetaTableRequest
	GetVisibility() *int32
}

type UpdateMetaTableRequest struct {
	// The names of the labels that you want to add. Separate the labels with commas (,).
	//
	// example:
	//
	// a,b,c
	AddedLabels *string `json:"AddedLabels,omitempty" xml:"AddedLabels,omitempty"`
	// The display name of the table.
	//
	// example:
	//
	// test
	Caption *string `json:"Caption,omitempty" xml:"Caption,omitempty"`
	// The ID of the category that you want to associate.
	//
	// example:
	//
	// 101
	CategoryId *int64 `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	// The environment of the DataWorks workspace. Valid values: 0 and 1. The value 0 indicates the development environment. The value 1 indicates the production environment.
	//
	// example:
	//
	// 1
	EnvType *int32 `json:"EnvType,omitempty" xml:"EnvType,omitempty"`
	// The new owner ID. If you leave this parameter empty, the owner ID is not updated.
	//
	// example:
	//
	// 12345
	NewOwnerId *string `json:"NewOwnerId,omitempty" xml:"NewOwnerId,omitempty"`
	// The DataWorks workspace ID.
	//
	// example:
	//
	// 101
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The names of labels that you want to remove. Separate the labels with commas (,).
	//
	// example:
	//
	// a,b,c
	RemovedLabels *string `json:"RemovedLabels,omitempty" xml:"RemovedLabels,omitempty"`
	// The schema information about the table. You must configure this parameter if you enable the three-layer model of MaxCompute.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// default
	Schema *string `json:"Schema,omitempty" xml:"Schema,omitempty"`
	// The GUID of the table. Specify the GUID in the format of odps.{projectName}.{tableName}.
	//
	// example:
	//
	// odps.test.table1
	TableGuid *string `json:"TableGuid,omitempty" xml:"TableGuid,omitempty"`
	// The name of the table.
	//
	// example:
	//
	// table1
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	// The scope in which the table is visible. Valid values: 0, 1, and 2. The value 0 indicates that the table is invisible to all members. The value 1 indicates that the table is visible to all members. The value 2 indicates that the table is visible to workspace members.
	//
	// example:
	//
	// 1
	Visibility *int32 `json:"Visibility,omitempty" xml:"Visibility,omitempty"`
}

func (s UpdateMetaTableRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateMetaTableRequest) GoString() string {
	return s.String()
}

func (s *UpdateMetaTableRequest) GetAddedLabels() *string {
	return s.AddedLabels
}

func (s *UpdateMetaTableRequest) GetCaption() *string {
	return s.Caption
}

func (s *UpdateMetaTableRequest) GetCategoryId() *int64 {
	return s.CategoryId
}

func (s *UpdateMetaTableRequest) GetEnvType() *int32 {
	return s.EnvType
}

func (s *UpdateMetaTableRequest) GetNewOwnerId() *string {
	return s.NewOwnerId
}

func (s *UpdateMetaTableRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *UpdateMetaTableRequest) GetRemovedLabels() *string {
	return s.RemovedLabels
}

func (s *UpdateMetaTableRequest) GetSchema() *string {
	return s.Schema
}

func (s *UpdateMetaTableRequest) GetTableGuid() *string {
	return s.TableGuid
}

func (s *UpdateMetaTableRequest) GetTableName() *string {
	return s.TableName
}

func (s *UpdateMetaTableRequest) GetVisibility() *int32 {
	return s.Visibility
}

func (s *UpdateMetaTableRequest) SetAddedLabels(v string) *UpdateMetaTableRequest {
	s.AddedLabels = &v
	return s
}

func (s *UpdateMetaTableRequest) SetCaption(v string) *UpdateMetaTableRequest {
	s.Caption = &v
	return s
}

func (s *UpdateMetaTableRequest) SetCategoryId(v int64) *UpdateMetaTableRequest {
	s.CategoryId = &v
	return s
}

func (s *UpdateMetaTableRequest) SetEnvType(v int32) *UpdateMetaTableRequest {
	s.EnvType = &v
	return s
}

func (s *UpdateMetaTableRequest) SetNewOwnerId(v string) *UpdateMetaTableRequest {
	s.NewOwnerId = &v
	return s
}

func (s *UpdateMetaTableRequest) SetProjectId(v int64) *UpdateMetaTableRequest {
	s.ProjectId = &v
	return s
}

func (s *UpdateMetaTableRequest) SetRemovedLabels(v string) *UpdateMetaTableRequest {
	s.RemovedLabels = &v
	return s
}

func (s *UpdateMetaTableRequest) SetSchema(v string) *UpdateMetaTableRequest {
	s.Schema = &v
	return s
}

func (s *UpdateMetaTableRequest) SetTableGuid(v string) *UpdateMetaTableRequest {
	s.TableGuid = &v
	return s
}

func (s *UpdateMetaTableRequest) SetTableName(v string) *UpdateMetaTableRequest {
	s.TableName = &v
	return s
}

func (s *UpdateMetaTableRequest) SetVisibility(v int32) *UpdateMetaTableRequest {
	s.Visibility = &v
	return s
}

func (s *UpdateMetaTableRequest) Validate() error {
	return dara.Validate(s)
}
