// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyApsSlsADBJobShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetColumnsShrink(v string) *ModifyApsSlsADBJobShrinkRequest
	GetColumnsShrink() *string
	SetDBClusterId(v string) *ModifyApsSlsADBJobShrinkRequest
	GetDBClusterId() *string
	SetDbName(v string) *ModifyApsSlsADBJobShrinkRequest
	GetDbName() *string
	SetDirtyDataProcessPattern(v string) *ModifyApsSlsADBJobShrinkRequest
	GetDirtyDataProcessPattern() *string
	SetExactlyOnce(v string) *ModifyApsSlsADBJobShrinkRequest
	GetExactlyOnce() *string
	SetPassword(v string) *ModifyApsSlsADBJobShrinkRequest
	GetPassword() *string
	SetRegionId(v string) *ModifyApsSlsADBJobShrinkRequest
	GetRegionId() *string
	SetStartingOffsets(v string) *ModifyApsSlsADBJobShrinkRequest
	GetStartingOffsets() *string
	SetTableName(v string) *ModifyApsSlsADBJobShrinkRequest
	GetTableName() *string
	SetUnixTimestampConvert(v string) *ModifyApsSlsADBJobShrinkRequest
	GetUnixTimestampConvert() *string
	SetUserName(v string) *ModifyApsSlsADBJobShrinkRequest
	GetUserName() *string
	SetWorkloadId(v string) *ModifyApsSlsADBJobShrinkRequest
	GetWorkloadId() *string
	SetWorkloadName(v string) *ModifyApsSlsADBJobShrinkRequest
	GetWorkloadName() *string
}

type ModifyApsSlsADBJobShrinkRequest struct {
	// The information about columns.
	//
	// example:
	//
	// -
	ColumnsShrink *string `json:"Columns,omitempty" xml:"Columns,omitempty"`
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

func (s ModifyApsSlsADBJobShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyApsSlsADBJobShrinkRequest) GoString() string {
	return s.String()
}

func (s *ModifyApsSlsADBJobShrinkRequest) GetColumnsShrink() *string {
	return s.ColumnsShrink
}

func (s *ModifyApsSlsADBJobShrinkRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ModifyApsSlsADBJobShrinkRequest) GetDbName() *string {
	return s.DbName
}

func (s *ModifyApsSlsADBJobShrinkRequest) GetDirtyDataProcessPattern() *string {
	return s.DirtyDataProcessPattern
}

func (s *ModifyApsSlsADBJobShrinkRequest) GetExactlyOnce() *string {
	return s.ExactlyOnce
}

func (s *ModifyApsSlsADBJobShrinkRequest) GetPassword() *string {
	return s.Password
}

func (s *ModifyApsSlsADBJobShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyApsSlsADBJobShrinkRequest) GetStartingOffsets() *string {
	return s.StartingOffsets
}

func (s *ModifyApsSlsADBJobShrinkRequest) GetTableName() *string {
	return s.TableName
}

func (s *ModifyApsSlsADBJobShrinkRequest) GetUnixTimestampConvert() *string {
	return s.UnixTimestampConvert
}

func (s *ModifyApsSlsADBJobShrinkRequest) GetUserName() *string {
	return s.UserName
}

func (s *ModifyApsSlsADBJobShrinkRequest) GetWorkloadId() *string {
	return s.WorkloadId
}

func (s *ModifyApsSlsADBJobShrinkRequest) GetWorkloadName() *string {
	return s.WorkloadName
}

func (s *ModifyApsSlsADBJobShrinkRequest) SetColumnsShrink(v string) *ModifyApsSlsADBJobShrinkRequest {
	s.ColumnsShrink = &v
	return s
}

func (s *ModifyApsSlsADBJobShrinkRequest) SetDBClusterId(v string) *ModifyApsSlsADBJobShrinkRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyApsSlsADBJobShrinkRequest) SetDbName(v string) *ModifyApsSlsADBJobShrinkRequest {
	s.DbName = &v
	return s
}

func (s *ModifyApsSlsADBJobShrinkRequest) SetDirtyDataProcessPattern(v string) *ModifyApsSlsADBJobShrinkRequest {
	s.DirtyDataProcessPattern = &v
	return s
}

func (s *ModifyApsSlsADBJobShrinkRequest) SetExactlyOnce(v string) *ModifyApsSlsADBJobShrinkRequest {
	s.ExactlyOnce = &v
	return s
}

func (s *ModifyApsSlsADBJobShrinkRequest) SetPassword(v string) *ModifyApsSlsADBJobShrinkRequest {
	s.Password = &v
	return s
}

func (s *ModifyApsSlsADBJobShrinkRequest) SetRegionId(v string) *ModifyApsSlsADBJobShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyApsSlsADBJobShrinkRequest) SetStartingOffsets(v string) *ModifyApsSlsADBJobShrinkRequest {
	s.StartingOffsets = &v
	return s
}

func (s *ModifyApsSlsADBJobShrinkRequest) SetTableName(v string) *ModifyApsSlsADBJobShrinkRequest {
	s.TableName = &v
	return s
}

func (s *ModifyApsSlsADBJobShrinkRequest) SetUnixTimestampConvert(v string) *ModifyApsSlsADBJobShrinkRequest {
	s.UnixTimestampConvert = &v
	return s
}

func (s *ModifyApsSlsADBJobShrinkRequest) SetUserName(v string) *ModifyApsSlsADBJobShrinkRequest {
	s.UserName = &v
	return s
}

func (s *ModifyApsSlsADBJobShrinkRequest) SetWorkloadId(v string) *ModifyApsSlsADBJobShrinkRequest {
	s.WorkloadId = &v
	return s
}

func (s *ModifyApsSlsADBJobShrinkRequest) SetWorkloadName(v string) *ModifyApsSlsADBJobShrinkRequest {
	s.WorkloadName = &v
	return s
}

func (s *ModifyApsSlsADBJobShrinkRequest) Validate() error {
	return dara.Validate(s)
}
