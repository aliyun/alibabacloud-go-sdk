// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMmsTablesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDbName(v string) *UpdateMmsTablesRequest
	GetDbName() *string
	SetDstProjectName(v string) *UpdateMmsTablesRequest
	GetDstProjectName() *string
	SetDstSchemaName(v string) *UpdateMmsTablesRequest
	GetDstSchemaName() *string
	SetStatus(v string) *UpdateMmsTablesRequest
	GetStatus() *string
	SetTableNames(v []*string) *UpdateMmsTablesRequest
	GetTableNames() []*string
	SetTables(v []*int64) *UpdateMmsTablesRequest
	GetTables() []*int64
}

type UpdateMmsTablesRequest struct {
	// The name of the database that contains the source tables to be updated.
	//
	// example:
	//
	// default
	DbName *string `json:"dbName,omitempty" xml:"dbName,omitempty"`
	// The destination MaxCompute project.
	//
	// example:
	//
	// pj
	DstProjectName *string `json:"dstProjectName,omitempty" xml:"dstProjectName,omitempty"`
	// The destination MaxCompute schema.
	//
	// example:
	//
	// default
	DstSchemaName *string `json:"dstSchemaName,omitempty" xml:"dstSchemaName,omitempty"`
	// The migration status.
	//
	// example:
	//
	// INIT
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The list of source tables to be updated.
	TableNames []*string `json:"tableNames,omitempty" xml:"tableNames,omitempty" type:"Repeated"`
	// This parameter is deprecated.
	Tables []*int64 `json:"tables,omitempty" xml:"tables,omitempty" type:"Repeated"`
}

func (s UpdateMmsTablesRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateMmsTablesRequest) GoString() string {
	return s.String()
}

func (s *UpdateMmsTablesRequest) GetDbName() *string {
	return s.DbName
}

func (s *UpdateMmsTablesRequest) GetDstProjectName() *string {
	return s.DstProjectName
}

func (s *UpdateMmsTablesRequest) GetDstSchemaName() *string {
	return s.DstSchemaName
}

func (s *UpdateMmsTablesRequest) GetStatus() *string {
	return s.Status
}

func (s *UpdateMmsTablesRequest) GetTableNames() []*string {
	return s.TableNames
}

func (s *UpdateMmsTablesRequest) GetTables() []*int64 {
	return s.Tables
}

func (s *UpdateMmsTablesRequest) SetDbName(v string) *UpdateMmsTablesRequest {
	s.DbName = &v
	return s
}

func (s *UpdateMmsTablesRequest) SetDstProjectName(v string) *UpdateMmsTablesRequest {
	s.DstProjectName = &v
	return s
}

func (s *UpdateMmsTablesRequest) SetDstSchemaName(v string) *UpdateMmsTablesRequest {
	s.DstSchemaName = &v
	return s
}

func (s *UpdateMmsTablesRequest) SetStatus(v string) *UpdateMmsTablesRequest {
	s.Status = &v
	return s
}

func (s *UpdateMmsTablesRequest) SetTableNames(v []*string) *UpdateMmsTablesRequest {
	s.TableNames = v
	return s
}

func (s *UpdateMmsTablesRequest) SetTables(v []*int64) *UpdateMmsTablesRequest {
	s.Tables = v
	return s
}

func (s *UpdateMmsTablesRequest) Validate() error {
	return dara.Validate(s)
}
