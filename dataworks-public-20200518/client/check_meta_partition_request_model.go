// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckMetaPartitionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *CheckMetaPartitionRequest
	GetClusterId() *string
	SetDataSourceType(v string) *CheckMetaPartitionRequest
	GetDataSourceType() *string
	SetDatabaseName(v string) *CheckMetaPartitionRequest
	GetDatabaseName() *string
	SetPartition(v string) *CheckMetaPartitionRequest
	GetPartition() *string
	SetTableGuid(v string) *CheckMetaPartitionRequest
	GetTableGuid() *string
	SetTableName(v string) *CheckMetaPartitionRequest
	GetTableName() *string
}

type CheckMetaPartitionRequest struct {
	// This parameter is deprecated.
	//
	// example:
	//
	// abc
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The type of the data source. Set the value to odps.
	//
	// example:
	//
	// emr
	DataSourceType *string `json:"DataSourceType,omitempty" xml:"DataSourceType,omitempty"`
	// The name of the metadatabase.
	//
	// example:
	//
	// abc
	DatabaseName *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	// The name of the partition in the MaxCompute metatable.
	//
	// This parameter is required.
	//
	// example:
	//
	// ds=202005
	Partition *string `json:"Partition,omitempty" xml:"Partition,omitempty"`
	// The GUID of the MaxCompute metatable.
	//
	// example:
	//
	// odps.engine_name.table_name
	TableGuid *string `json:"TableGuid,omitempty" xml:"TableGuid,omitempty"`
	// The name of the metatable.
	//
	// example:
	//
	// abc
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s CheckMetaPartitionRequest) String() string {
	return dara.Prettify(s)
}

func (s CheckMetaPartitionRequest) GoString() string {
	return s.String()
}

func (s *CheckMetaPartitionRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *CheckMetaPartitionRequest) GetDataSourceType() *string {
	return s.DataSourceType
}

func (s *CheckMetaPartitionRequest) GetDatabaseName() *string {
	return s.DatabaseName
}

func (s *CheckMetaPartitionRequest) GetPartition() *string {
	return s.Partition
}

func (s *CheckMetaPartitionRequest) GetTableGuid() *string {
	return s.TableGuid
}

func (s *CheckMetaPartitionRequest) GetTableName() *string {
	return s.TableName
}

func (s *CheckMetaPartitionRequest) SetClusterId(v string) *CheckMetaPartitionRequest {
	s.ClusterId = &v
	return s
}

func (s *CheckMetaPartitionRequest) SetDataSourceType(v string) *CheckMetaPartitionRequest {
	s.DataSourceType = &v
	return s
}

func (s *CheckMetaPartitionRequest) SetDatabaseName(v string) *CheckMetaPartitionRequest {
	s.DatabaseName = &v
	return s
}

func (s *CheckMetaPartitionRequest) SetPartition(v string) *CheckMetaPartitionRequest {
	s.Partition = &v
	return s
}

func (s *CheckMetaPartitionRequest) SetTableGuid(v string) *CheckMetaPartitionRequest {
	s.TableGuid = &v
	return s
}

func (s *CheckMetaPartitionRequest) SetTableName(v string) *CheckMetaPartitionRequest {
	s.TableName = &v
	return s
}

func (s *CheckMetaPartitionRequest) Validate() error {
	return dara.Validate(s)
}
