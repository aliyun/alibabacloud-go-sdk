// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDataLakeDatabaseRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCatalogName(v string) *CreateDataLakeDatabaseRequest
	GetCatalogName() *string
	SetDataRegion(v string) *CreateDataLakeDatabaseRequest
	GetDataRegion() *string
	SetDbName(v string) *CreateDataLakeDatabaseRequest
	GetDbName() *string
	SetDescription(v string) *CreateDataLakeDatabaseRequest
	GetDescription() *string
	SetLocation(v string) *CreateDataLakeDatabaseRequest
	GetLocation() *string
	SetParameters(v map[string]*string) *CreateDataLakeDatabaseRequest
	GetParameters() map[string]*string
	SetTid(v int64) *CreateDataLakeDatabaseRequest
	GetTid() *int64
	SetWorkspaceId(v int64) *CreateDataLakeDatabaseRequest
	GetWorkspaceId() *int64
}

type CreateDataLakeDatabaseRequest struct {
	// The name of the data catalog.
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
	// The storage path of the database. OSS, S3, and S3A are supported.
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

func (s CreateDataLakeDatabaseRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDataLakeDatabaseRequest) GoString() string {
	return s.String()
}

func (s *CreateDataLakeDatabaseRequest) GetCatalogName() *string {
	return s.CatalogName
}

func (s *CreateDataLakeDatabaseRequest) GetDataRegion() *string {
	return s.DataRegion
}

func (s *CreateDataLakeDatabaseRequest) GetDbName() *string {
	return s.DbName
}

func (s *CreateDataLakeDatabaseRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateDataLakeDatabaseRequest) GetLocation() *string {
	return s.Location
}

func (s *CreateDataLakeDatabaseRequest) GetParameters() map[string]*string {
	return s.Parameters
}

func (s *CreateDataLakeDatabaseRequest) GetTid() *int64 {
	return s.Tid
}

func (s *CreateDataLakeDatabaseRequest) GetWorkspaceId() *int64 {
	return s.WorkspaceId
}

func (s *CreateDataLakeDatabaseRequest) SetCatalogName(v string) *CreateDataLakeDatabaseRequest {
	s.CatalogName = &v
	return s
}

func (s *CreateDataLakeDatabaseRequest) SetDataRegion(v string) *CreateDataLakeDatabaseRequest {
	s.DataRegion = &v
	return s
}

func (s *CreateDataLakeDatabaseRequest) SetDbName(v string) *CreateDataLakeDatabaseRequest {
	s.DbName = &v
	return s
}

func (s *CreateDataLakeDatabaseRequest) SetDescription(v string) *CreateDataLakeDatabaseRequest {
	s.Description = &v
	return s
}

func (s *CreateDataLakeDatabaseRequest) SetLocation(v string) *CreateDataLakeDatabaseRequest {
	s.Location = &v
	return s
}

func (s *CreateDataLakeDatabaseRequest) SetParameters(v map[string]*string) *CreateDataLakeDatabaseRequest {
	s.Parameters = v
	return s
}

func (s *CreateDataLakeDatabaseRequest) SetTid(v int64) *CreateDataLakeDatabaseRequest {
	s.Tid = &v
	return s
}

func (s *CreateDataLakeDatabaseRequest) SetWorkspaceId(v int64) *CreateDataLakeDatabaseRequest {
	s.WorkspaceId = &v
	return s
}

func (s *CreateDataLakeDatabaseRequest) Validate() error {
	return dara.Validate(s)
}
