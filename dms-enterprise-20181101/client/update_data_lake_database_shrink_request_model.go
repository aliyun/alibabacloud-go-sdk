// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDataLakeDatabaseShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCatalogName(v string) *UpdateDataLakeDatabaseShrinkRequest
	GetCatalogName() *string
	SetDataRegion(v string) *UpdateDataLakeDatabaseShrinkRequest
	GetDataRegion() *string
	SetDbName(v string) *UpdateDataLakeDatabaseShrinkRequest
	GetDbName() *string
	SetDescription(v string) *UpdateDataLakeDatabaseShrinkRequest
	GetDescription() *string
	SetLocation(v string) *UpdateDataLakeDatabaseShrinkRequest
	GetLocation() *string
	SetParametersShrink(v string) *UpdateDataLakeDatabaseShrinkRequest
	GetParametersShrink() *string
	SetTid(v int64) *UpdateDataLakeDatabaseShrinkRequest
	GetTid() *int64
	SetWorkspaceId(v int64) *UpdateDataLakeDatabaseShrinkRequest
	GetWorkspaceId() *int64
}

type UpdateDataLakeDatabaseShrinkRequest struct {
	// The catalog name.
	//
	// This parameter is required.
	//
	// example:
	//
	// hive
	CatalogName *string `json:"CatalogName,omitempty" xml:"CatalogName,omitempty"`
	// The region where the data lake resides.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	DataRegion *string `json:"DataRegion,omitempty" xml:"DataRegion,omitempty"`
	// The name of the database that you want to query.
	//
	// This parameter is required.
	//
	// example:
	//
	// default
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// The description of the database.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The storage path of the database. Supports the OSS, S3, and S3A protocols.
	//
	// This parameter is required.
	//
	// example:
	//
	// oss://path/to/database
	Location *string `json:"Location,omitempty" xml:"Location,omitempty"`
	// The key-value pairs of the database attributes.
	ParametersShrink *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// The ID of the tenant.
	//
	// > To view the ID of the tenant, go to the Data Management (DMS) console and move the pointer over the profile picture in the upper-right corner. For more information, see the [View information about the current tenant](https://help.aliyun.com/document_detail/181330.html) section of the "Manage DMS tenants" topic.
	//
	// example:
	//
	// 3***
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
	// The workspace ID.
	//
	// example:
	//
	// 12****
	WorkspaceId *int64 `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s UpdateDataLakeDatabaseShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataLakeDatabaseShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateDataLakeDatabaseShrinkRequest) GetCatalogName() *string {
	return s.CatalogName
}

func (s *UpdateDataLakeDatabaseShrinkRequest) GetDataRegion() *string {
	return s.DataRegion
}

func (s *UpdateDataLakeDatabaseShrinkRequest) GetDbName() *string {
	return s.DbName
}

func (s *UpdateDataLakeDatabaseShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateDataLakeDatabaseShrinkRequest) GetLocation() *string {
	return s.Location
}

func (s *UpdateDataLakeDatabaseShrinkRequest) GetParametersShrink() *string {
	return s.ParametersShrink
}

func (s *UpdateDataLakeDatabaseShrinkRequest) GetTid() *int64 {
	return s.Tid
}

func (s *UpdateDataLakeDatabaseShrinkRequest) GetWorkspaceId() *int64 {
	return s.WorkspaceId
}

func (s *UpdateDataLakeDatabaseShrinkRequest) SetCatalogName(v string) *UpdateDataLakeDatabaseShrinkRequest {
	s.CatalogName = &v
	return s
}

func (s *UpdateDataLakeDatabaseShrinkRequest) SetDataRegion(v string) *UpdateDataLakeDatabaseShrinkRequest {
	s.DataRegion = &v
	return s
}

func (s *UpdateDataLakeDatabaseShrinkRequest) SetDbName(v string) *UpdateDataLakeDatabaseShrinkRequest {
	s.DbName = &v
	return s
}

func (s *UpdateDataLakeDatabaseShrinkRequest) SetDescription(v string) *UpdateDataLakeDatabaseShrinkRequest {
	s.Description = &v
	return s
}

func (s *UpdateDataLakeDatabaseShrinkRequest) SetLocation(v string) *UpdateDataLakeDatabaseShrinkRequest {
	s.Location = &v
	return s
}

func (s *UpdateDataLakeDatabaseShrinkRequest) SetParametersShrink(v string) *UpdateDataLakeDatabaseShrinkRequest {
	s.ParametersShrink = &v
	return s
}

func (s *UpdateDataLakeDatabaseShrinkRequest) SetTid(v int64) *UpdateDataLakeDatabaseShrinkRequest {
	s.Tid = &v
	return s
}

func (s *UpdateDataLakeDatabaseShrinkRequest) SetWorkspaceId(v int64) *UpdateDataLakeDatabaseShrinkRequest {
	s.WorkspaceId = &v
	return s
}

func (s *UpdateDataLakeDatabaseShrinkRequest) Validate() error {
	return dara.Validate(s)
}
