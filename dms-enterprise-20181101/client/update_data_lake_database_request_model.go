// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDataLakeDatabaseRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCatalogName(v string) *UpdateDataLakeDatabaseRequest
	GetCatalogName() *string
	SetDataRegion(v string) *UpdateDataLakeDatabaseRequest
	GetDataRegion() *string
	SetDbName(v string) *UpdateDataLakeDatabaseRequest
	GetDbName() *string
	SetDescription(v string) *UpdateDataLakeDatabaseRequest
	GetDescription() *string
	SetLocation(v string) *UpdateDataLakeDatabaseRequest
	GetLocation() *string
	SetParameters(v map[string]*string) *UpdateDataLakeDatabaseRequest
	GetParameters() map[string]*string
	SetTid(v int64) *UpdateDataLakeDatabaseRequest
	GetTid() *int64
	SetWorkspaceId(v int64) *UpdateDataLakeDatabaseRequest
	GetWorkspaceId() *int64
}

type UpdateDataLakeDatabaseRequest struct {
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
	Parameters map[string]*string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
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

func (s UpdateDataLakeDatabaseRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataLakeDatabaseRequest) GoString() string {
	return s.String()
}

func (s *UpdateDataLakeDatabaseRequest) GetCatalogName() *string {
	return s.CatalogName
}

func (s *UpdateDataLakeDatabaseRequest) GetDataRegion() *string {
	return s.DataRegion
}

func (s *UpdateDataLakeDatabaseRequest) GetDbName() *string {
	return s.DbName
}

func (s *UpdateDataLakeDatabaseRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateDataLakeDatabaseRequest) GetLocation() *string {
	return s.Location
}

func (s *UpdateDataLakeDatabaseRequest) GetParameters() map[string]*string {
	return s.Parameters
}

func (s *UpdateDataLakeDatabaseRequest) GetTid() *int64 {
	return s.Tid
}

func (s *UpdateDataLakeDatabaseRequest) GetWorkspaceId() *int64 {
	return s.WorkspaceId
}

func (s *UpdateDataLakeDatabaseRequest) SetCatalogName(v string) *UpdateDataLakeDatabaseRequest {
	s.CatalogName = &v
	return s
}

func (s *UpdateDataLakeDatabaseRequest) SetDataRegion(v string) *UpdateDataLakeDatabaseRequest {
	s.DataRegion = &v
	return s
}

func (s *UpdateDataLakeDatabaseRequest) SetDbName(v string) *UpdateDataLakeDatabaseRequest {
	s.DbName = &v
	return s
}

func (s *UpdateDataLakeDatabaseRequest) SetDescription(v string) *UpdateDataLakeDatabaseRequest {
	s.Description = &v
	return s
}

func (s *UpdateDataLakeDatabaseRequest) SetLocation(v string) *UpdateDataLakeDatabaseRequest {
	s.Location = &v
	return s
}

func (s *UpdateDataLakeDatabaseRequest) SetParameters(v map[string]*string) *UpdateDataLakeDatabaseRequest {
	s.Parameters = v
	return s
}

func (s *UpdateDataLakeDatabaseRequest) SetTid(v int64) *UpdateDataLakeDatabaseRequest {
	s.Tid = &v
	return s
}

func (s *UpdateDataLakeDatabaseRequest) SetWorkspaceId(v int64) *UpdateDataLakeDatabaseRequest {
	s.WorkspaceId = &v
	return s
}

func (s *UpdateDataLakeDatabaseRequest) Validate() error {
	return dara.Validate(s)
}
