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
	// This parameter is required.
	//
	// example:
	//
	// hive
	CatalogName *string `json:"CatalogName,omitempty" xml:"CatalogName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	DataRegion *string `json:"DataRegion,omitempty" xml:"DataRegion,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// default
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// oss://path/to/database
	Location   *string            `json:"Location,omitempty" xml:"Location,omitempty"`
	Parameters map[string]*string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// example:
	//
	// 3***
	Tid         *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
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
