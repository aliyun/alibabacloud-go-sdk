// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDataLakeDatabaseShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCatalogName(v string) *CreateDataLakeDatabaseShrinkRequest
	GetCatalogName() *string
	SetDataRegion(v string) *CreateDataLakeDatabaseShrinkRequest
	GetDataRegion() *string
	SetDbName(v string) *CreateDataLakeDatabaseShrinkRequest
	GetDbName() *string
	SetDescription(v string) *CreateDataLakeDatabaseShrinkRequest
	GetDescription() *string
	SetLocation(v string) *CreateDataLakeDatabaseShrinkRequest
	GetLocation() *string
	SetParametersShrink(v string) *CreateDataLakeDatabaseShrinkRequest
	GetParametersShrink() *string
	SetTid(v int64) *CreateDataLakeDatabaseShrinkRequest
	GetTid() *int64
	SetWorkspaceId(v int64) *CreateDataLakeDatabaseShrinkRequest
	GetWorkspaceId() *int64
}

type CreateDataLakeDatabaseShrinkRequest struct {
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
	Location         *string `json:"Location,omitempty" xml:"Location,omitempty"`
	ParametersShrink *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// example:
	//
	// 3***
	Tid         *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
	WorkspaceId *int64 `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s CreateDataLakeDatabaseShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDataLakeDatabaseShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateDataLakeDatabaseShrinkRequest) GetCatalogName() *string {
	return s.CatalogName
}

func (s *CreateDataLakeDatabaseShrinkRequest) GetDataRegion() *string {
	return s.DataRegion
}

func (s *CreateDataLakeDatabaseShrinkRequest) GetDbName() *string {
	return s.DbName
}

func (s *CreateDataLakeDatabaseShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateDataLakeDatabaseShrinkRequest) GetLocation() *string {
	return s.Location
}

func (s *CreateDataLakeDatabaseShrinkRequest) GetParametersShrink() *string {
	return s.ParametersShrink
}

func (s *CreateDataLakeDatabaseShrinkRequest) GetTid() *int64 {
	return s.Tid
}

func (s *CreateDataLakeDatabaseShrinkRequest) GetWorkspaceId() *int64 {
	return s.WorkspaceId
}

func (s *CreateDataLakeDatabaseShrinkRequest) SetCatalogName(v string) *CreateDataLakeDatabaseShrinkRequest {
	s.CatalogName = &v
	return s
}

func (s *CreateDataLakeDatabaseShrinkRequest) SetDataRegion(v string) *CreateDataLakeDatabaseShrinkRequest {
	s.DataRegion = &v
	return s
}

func (s *CreateDataLakeDatabaseShrinkRequest) SetDbName(v string) *CreateDataLakeDatabaseShrinkRequest {
	s.DbName = &v
	return s
}

func (s *CreateDataLakeDatabaseShrinkRequest) SetDescription(v string) *CreateDataLakeDatabaseShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreateDataLakeDatabaseShrinkRequest) SetLocation(v string) *CreateDataLakeDatabaseShrinkRequest {
	s.Location = &v
	return s
}

func (s *CreateDataLakeDatabaseShrinkRequest) SetParametersShrink(v string) *CreateDataLakeDatabaseShrinkRequest {
	s.ParametersShrink = &v
	return s
}

func (s *CreateDataLakeDatabaseShrinkRequest) SetTid(v int64) *CreateDataLakeDatabaseShrinkRequest {
	s.Tid = &v
	return s
}

func (s *CreateDataLakeDatabaseShrinkRequest) SetWorkspaceId(v int64) *CreateDataLakeDatabaseShrinkRequest {
	s.WorkspaceId = &v
	return s
}

func (s *CreateDataLakeDatabaseShrinkRequest) Validate() error {
	return dara.Validate(s)
}
