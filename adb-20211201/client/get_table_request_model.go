// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTableRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *GetTableRequest
	GetDBClusterId() *string
	SetDbName(v string) *GetTableRequest
	GetDbName() *string
	SetRegionId(v string) *GetTableRequest
	GetRegionId() *string
	SetTableName(v string) *GetTableRequest
	GetTableName() *string
}

type GetTableRequest struct {
	// The ID of the cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// amv-*******
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The name of the database.
	//
	// example:
	//
	// dbName
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// The ID of the region in which the cluster resides.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the table.
	//
	// example:
	//
	// tableName
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s GetTableRequest) String() string {
	return dara.Prettify(s)
}

func (s GetTableRequest) GoString() string {
	return s.String()
}

func (s *GetTableRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *GetTableRequest) GetDbName() *string {
	return s.DbName
}

func (s *GetTableRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetTableRequest) GetTableName() *string {
	return s.TableName
}

func (s *GetTableRequest) SetDBClusterId(v string) *GetTableRequest {
	s.DBClusterId = &v
	return s
}

func (s *GetTableRequest) SetDbName(v string) *GetTableRequest {
	s.DbName = &v
	return s
}

func (s *GetTableRequest) SetRegionId(v string) *GetTableRequest {
	s.RegionId = &v
	return s
}

func (s *GetTableRequest) SetTableName(v string) *GetTableRequest {
	s.TableName = &v
	return s
}

func (s *GetTableRequest) Validate() error {
	return dara.Validate(s)
}
