// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTableDDLRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *GetTableDDLRequest
	GetDBClusterId() *string
	SetRegionId(v string) *GetTableDDLRequest
	GetRegionId() *string
	SetSchemaName(v string) *GetTableDDLRequest
	GetSchemaName() *string
	SetTableName(v string) *GetTableDDLRequest
	GetTableName() *string
}

type GetTableDDLRequest struct {
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// am-bp1ub9grke1****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the database.
	//
	// example:
	//
	// adb_demo
	SchemaName *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
	// The name of the table.
	//
	// example:
	//
	// test
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s GetTableDDLRequest) String() string {
	return dara.Prettify(s)
}

func (s GetTableDDLRequest) GoString() string {
	return s.String()
}

func (s *GetTableDDLRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *GetTableDDLRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetTableDDLRequest) GetSchemaName() *string {
	return s.SchemaName
}

func (s *GetTableDDLRequest) GetTableName() *string {
	return s.TableName
}

func (s *GetTableDDLRequest) SetDBClusterId(v string) *GetTableDDLRequest {
	s.DBClusterId = &v
	return s
}

func (s *GetTableDDLRequest) SetRegionId(v string) *GetTableDDLRequest {
	s.RegionId = &v
	return s
}

func (s *GetTableDDLRequest) SetSchemaName(v string) *GetTableDDLRequest {
	s.SchemaName = &v
	return s
}

func (s *GetTableDDLRequest) SetTableName(v string) *GetTableDDLRequest {
	s.TableName = &v
	return s
}

func (s *GetTableDDLRequest) Validate() error {
	return dara.Validate(s)
}
