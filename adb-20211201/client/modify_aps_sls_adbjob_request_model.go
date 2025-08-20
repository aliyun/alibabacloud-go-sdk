// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyApsSlsADBJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetColumns(v []*ModifyApsSlsADBJobRequestColumns) *ModifyApsSlsADBJobRequest
	GetColumns() []*ModifyApsSlsADBJobRequestColumns
	SetDBClusterId(v string) *ModifyApsSlsADBJobRequest
	GetDBClusterId() *string
	SetDbName(v string) *ModifyApsSlsADBJobRequest
	GetDbName() *string
	SetDirtyDataProcessPattern(v string) *ModifyApsSlsADBJobRequest
	GetDirtyDataProcessPattern() *string
	SetExactlyOnce(v string) *ModifyApsSlsADBJobRequest
	GetExactlyOnce() *string
	SetPassword(v string) *ModifyApsSlsADBJobRequest
	GetPassword() *string
	SetRegionId(v string) *ModifyApsSlsADBJobRequest
	GetRegionId() *string
	SetStartingOffsets(v string) *ModifyApsSlsADBJobRequest
	GetStartingOffsets() *string
	SetTableName(v string) *ModifyApsSlsADBJobRequest
	GetTableName() *string
	SetUnixTimestampConvert(v string) *ModifyApsSlsADBJobRequest
	GetUnixTimestampConvert() *string
	SetUserName(v string) *ModifyApsSlsADBJobRequest
	GetUserName() *string
	SetWorkloadId(v string) *ModifyApsSlsADBJobRequest
	GetWorkloadId() *string
	SetWorkloadName(v string) *ModifyApsSlsADBJobRequest
	GetWorkloadName() *string
}

type ModifyApsSlsADBJobRequest struct {
	// The information about columns.
	//
	// example:
	//
	// -
	Columns []*ModifyApsSlsADBJobRequestColumns `json:"Columns,omitempty" xml:"Columns,omitempty" type:"Repeated"`
	// The cluster ID.
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
	// The dirty data processing mode.
	//
	// example:
	//
	// STOP
	DirtyDataProcessPattern *string `json:"DirtyDataProcessPattern,omitempty" xml:"DirtyDataProcessPattern,omitempty"`
	// Specifies whether to enable the consistency check.
	//
	// example:
	//
	// true
	ExactlyOnce *string `json:"ExactlyOnce,omitempty" xml:"ExactlyOnce,omitempty"`
	// The password of the database account.
	//
	// example:
	//
	// test_123
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The start offset.
	//
	// example:
	//
	// end_cursor
	StartingOffsets *string `json:"StartingOffsets,omitempty" xml:"StartingOffsets,omitempty"`
	// The name of the table.
	//
	// example:
	//
	// test
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	// The timestamp conversion.
	//
	// example:
	//
	// -
	UnixTimestampConvert *string `json:"UnixTimestampConvert,omitempty" xml:"UnixTimestampConvert,omitempty"`
	// The name of the database account.
	//
	// example:
	//
	// user-name
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	// The job ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// aps-******
	WorkloadId *string `json:"WorkloadId,omitempty" xml:"WorkloadId,omitempty"`
	// The name of the workload.
	//
	// example:
	//
	// test-name
	WorkloadName *string `json:"WorkloadName,omitempty" xml:"WorkloadName,omitempty"`
}

func (s ModifyApsSlsADBJobRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyApsSlsADBJobRequest) GoString() string {
	return s.String()
}

func (s *ModifyApsSlsADBJobRequest) GetColumns() []*ModifyApsSlsADBJobRequestColumns {
	return s.Columns
}

func (s *ModifyApsSlsADBJobRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ModifyApsSlsADBJobRequest) GetDbName() *string {
	return s.DbName
}

func (s *ModifyApsSlsADBJobRequest) GetDirtyDataProcessPattern() *string {
	return s.DirtyDataProcessPattern
}

func (s *ModifyApsSlsADBJobRequest) GetExactlyOnce() *string {
	return s.ExactlyOnce
}

func (s *ModifyApsSlsADBJobRequest) GetPassword() *string {
	return s.Password
}

func (s *ModifyApsSlsADBJobRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyApsSlsADBJobRequest) GetStartingOffsets() *string {
	return s.StartingOffsets
}

func (s *ModifyApsSlsADBJobRequest) GetTableName() *string {
	return s.TableName
}

func (s *ModifyApsSlsADBJobRequest) GetUnixTimestampConvert() *string {
	return s.UnixTimestampConvert
}

func (s *ModifyApsSlsADBJobRequest) GetUserName() *string {
	return s.UserName
}

func (s *ModifyApsSlsADBJobRequest) GetWorkloadId() *string {
	return s.WorkloadId
}

func (s *ModifyApsSlsADBJobRequest) GetWorkloadName() *string {
	return s.WorkloadName
}

func (s *ModifyApsSlsADBJobRequest) SetColumns(v []*ModifyApsSlsADBJobRequestColumns) *ModifyApsSlsADBJobRequest {
	s.Columns = v
	return s
}

func (s *ModifyApsSlsADBJobRequest) SetDBClusterId(v string) *ModifyApsSlsADBJobRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyApsSlsADBJobRequest) SetDbName(v string) *ModifyApsSlsADBJobRequest {
	s.DbName = &v
	return s
}

func (s *ModifyApsSlsADBJobRequest) SetDirtyDataProcessPattern(v string) *ModifyApsSlsADBJobRequest {
	s.DirtyDataProcessPattern = &v
	return s
}

func (s *ModifyApsSlsADBJobRequest) SetExactlyOnce(v string) *ModifyApsSlsADBJobRequest {
	s.ExactlyOnce = &v
	return s
}

func (s *ModifyApsSlsADBJobRequest) SetPassword(v string) *ModifyApsSlsADBJobRequest {
	s.Password = &v
	return s
}

func (s *ModifyApsSlsADBJobRequest) SetRegionId(v string) *ModifyApsSlsADBJobRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyApsSlsADBJobRequest) SetStartingOffsets(v string) *ModifyApsSlsADBJobRequest {
	s.StartingOffsets = &v
	return s
}

func (s *ModifyApsSlsADBJobRequest) SetTableName(v string) *ModifyApsSlsADBJobRequest {
	s.TableName = &v
	return s
}

func (s *ModifyApsSlsADBJobRequest) SetUnixTimestampConvert(v string) *ModifyApsSlsADBJobRequest {
	s.UnixTimestampConvert = &v
	return s
}

func (s *ModifyApsSlsADBJobRequest) SetUserName(v string) *ModifyApsSlsADBJobRequest {
	s.UserName = &v
	return s
}

func (s *ModifyApsSlsADBJobRequest) SetWorkloadId(v string) *ModifyApsSlsADBJobRequest {
	s.WorkloadId = &v
	return s
}

func (s *ModifyApsSlsADBJobRequest) SetWorkloadName(v string) *ModifyApsSlsADBJobRequest {
	s.WorkloadName = &v
	return s
}

func (s *ModifyApsSlsADBJobRequest) Validate() error {
	return dara.Validate(s)
}

type ModifyApsSlsADBJobRequestColumns struct {
	// The name of the mapping.
	//
	// example:
	//
	// map_name
	MapName *string `json:"MapName,omitempty" xml:"MapName,omitempty"`
	// The type of the mapping.
	//
	// example:
	//
	// bigint
	MapType *string `json:"MapType,omitempty" xml:"MapType,omitempty"`
	// The name of the column.
	//
	// example:
	//
	// test_name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The data type of the column.
	//
	// example:
	//
	// bigint
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ModifyApsSlsADBJobRequestColumns) String() string {
	return dara.Prettify(s)
}

func (s ModifyApsSlsADBJobRequestColumns) GoString() string {
	return s.String()
}

func (s *ModifyApsSlsADBJobRequestColumns) GetMapName() *string {
	return s.MapName
}

func (s *ModifyApsSlsADBJobRequestColumns) GetMapType() *string {
	return s.MapType
}

func (s *ModifyApsSlsADBJobRequestColumns) GetName() *string {
	return s.Name
}

func (s *ModifyApsSlsADBJobRequestColumns) GetType() *string {
	return s.Type
}

func (s *ModifyApsSlsADBJobRequestColumns) SetMapName(v string) *ModifyApsSlsADBJobRequestColumns {
	s.MapName = &v
	return s
}

func (s *ModifyApsSlsADBJobRequestColumns) SetMapType(v string) *ModifyApsSlsADBJobRequestColumns {
	s.MapType = &v
	return s
}

func (s *ModifyApsSlsADBJobRequestColumns) SetName(v string) *ModifyApsSlsADBJobRequestColumns {
	s.Name = &v
	return s
}

func (s *ModifyApsSlsADBJobRequestColumns) SetType(v string) *ModifyApsSlsADBJobRequestColumns {
	s.Type = &v
	return s
}

func (s *ModifyApsSlsADBJobRequestColumns) Validate() error {
	return dara.Validate(s)
}
