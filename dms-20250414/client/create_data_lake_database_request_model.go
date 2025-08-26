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
	Location   *string            `json:"Location,omitempty" xml:"Location,omitempty"`
	Parameters map[string]*string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// example:
	//
	// 3****
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
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
