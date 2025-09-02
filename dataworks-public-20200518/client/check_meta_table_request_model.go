// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckMetaTableRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *CheckMetaTableRequest
	GetClusterId() *string
	SetDataSourceType(v string) *CheckMetaTableRequest
	GetDataSourceType() *string
	SetDatabaseName(v string) *CheckMetaTableRequest
	GetDatabaseName() *string
	SetTableGuid(v string) *CheckMetaTableRequest
	GetTableGuid() *string
	SetTableName(v string) *CheckMetaTableRequest
	GetTableName() *string
}

type CheckMetaTableRequest struct {
	// The E-MapReduce (EMR) cluster ID.
	//
	// example:
	//
	// abc
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The type of the data source. Valid values: odps and emr.
	//
	// example:
	//
	// emr
	DataSourceType *string `json:"DataSourceType,omitempty" xml:"DataSourceType,omitempty"`
	// The name of the metadatabase of the EMR cluster.
	//
	// example:
	//
	// abc
	DatabaseName *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	// The GUID of the metatable.
	//
	// example:
	//
	// odps.engine_name.table_name
	TableGuid *string `json:"TableGuid,omitempty" xml:"TableGuid,omitempty"`
	// The name of the EMR metatable.
	//
	// example:
	//
	// abc
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s CheckMetaTableRequest) String() string {
	return dara.Prettify(s)
}

func (s CheckMetaTableRequest) GoString() string {
	return s.String()
}

func (s *CheckMetaTableRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *CheckMetaTableRequest) GetDataSourceType() *string {
	return s.DataSourceType
}

func (s *CheckMetaTableRequest) GetDatabaseName() *string {
	return s.DatabaseName
}

func (s *CheckMetaTableRequest) GetTableGuid() *string {
	return s.TableGuid
}

func (s *CheckMetaTableRequest) GetTableName() *string {
	return s.TableName
}

func (s *CheckMetaTableRequest) SetClusterId(v string) *CheckMetaTableRequest {
	s.ClusterId = &v
	return s
}

func (s *CheckMetaTableRequest) SetDataSourceType(v string) *CheckMetaTableRequest {
	s.DataSourceType = &v
	return s
}

func (s *CheckMetaTableRequest) SetDatabaseName(v string) *CheckMetaTableRequest {
	s.DatabaseName = &v
	return s
}

func (s *CheckMetaTableRequest) SetTableGuid(v string) *CheckMetaTableRequest {
	s.TableGuid = &v
	return s
}

func (s *CheckMetaTableRequest) SetTableName(v string) *CheckMetaTableRequest {
	s.TableName = &v
	return s
}

func (s *CheckMetaTableRequest) Validate() error {
	return dara.Validate(s)
}
