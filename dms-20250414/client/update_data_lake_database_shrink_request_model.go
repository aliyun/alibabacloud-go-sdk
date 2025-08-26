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
	// 3****
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
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
