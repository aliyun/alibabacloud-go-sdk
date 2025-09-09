// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitSqlFlashbackTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDbName(v string) *SubmitSqlFlashbackTaskRequest
	GetDbName() *string
	SetDrdsInstanceId(v string) *SubmitSqlFlashbackTaskRequest
	GetDrdsInstanceId() *string
	SetEndTime(v string) *SubmitSqlFlashbackTaskRequest
	GetEndTime() *string
	SetRecallRestoreType(v int32) *SubmitSqlFlashbackTaskRequest
	GetRecallRestoreType() *int32
	SetRecallType(v int32) *SubmitSqlFlashbackTaskRequest
	GetRecallType() *int32
	SetSqlPk(v string) *SubmitSqlFlashbackTaskRequest
	GetSqlPk() *string
	SetSqlType(v string) *SubmitSqlFlashbackTaskRequest
	GetSqlType() *string
	SetStartTime(v string) *SubmitSqlFlashbackTaskRequest
	GetStartTime() *string
	SetTableName(v string) *SubmitSqlFlashbackTaskRequest
	GetTableName() *string
	SetTraceId(v string) *SubmitSqlFlashbackTaskRequest
	GetTraceId() *string
}

type SubmitSqlFlashbackTaskRequest struct {
	// The name of the DRDS database.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// The ID of a DRDS instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// drds*********
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	// The time when the SQL flashback task ends.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2019-09-10 23:23:23
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The restoration type. Valid values:
	//
	// 	- 1: Image restoration
	//
	// 	- 0: reverse recovery
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	RecallRestoreType *int32 `json:"RecallRestoreType,omitempty" xml:"RecallRestoreType,omitempty"`
	// Exact match or fuzzy match. Valid values:
	//
	// 	- 0: the exact match.
	//
	// 	- 1: the fuzzy match.
	//
	// example:
	//
	// 0
	RecallType *int32 `json:"RecallType,omitempty" xml:"RecallType,omitempty"`
	// The primary key of flashback SQL.
	//
	// example:
	//
	// 11111
	SqlPk *string `json:"SqlPk,omitempty" xml:"SqlPk,omitempty"`
	// The type of the SQL statement. Valid values: INSERT, UPDATE, and DELETE. Separate multiple types with commas (,).
	//
	// example:
	//
	// INSERT,UPDATE
	SqlType *string `json:"SqlType,omitempty" xml:"SqlType,omitempty"`
	// The start time of the flashback SQL statement.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2019-09-10 20:23:23
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The name of the table where the flashback SQL operation was performed.
	//
	// example:
	//
	// test
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	// The Trace ID of the flashback SQL.
	//
	// example:
	//
	// ase*****
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s SubmitSqlFlashbackTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitSqlFlashbackTaskRequest) GoString() string {
	return s.String()
}

func (s *SubmitSqlFlashbackTaskRequest) GetDbName() *string {
	return s.DbName
}

func (s *SubmitSqlFlashbackTaskRequest) GetDrdsInstanceId() *string {
	return s.DrdsInstanceId
}

func (s *SubmitSqlFlashbackTaskRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *SubmitSqlFlashbackTaskRequest) GetRecallRestoreType() *int32 {
	return s.RecallRestoreType
}

func (s *SubmitSqlFlashbackTaskRequest) GetRecallType() *int32 {
	return s.RecallType
}

func (s *SubmitSqlFlashbackTaskRequest) GetSqlPk() *string {
	return s.SqlPk
}

func (s *SubmitSqlFlashbackTaskRequest) GetSqlType() *string {
	return s.SqlType
}

func (s *SubmitSqlFlashbackTaskRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *SubmitSqlFlashbackTaskRequest) GetTableName() *string {
	return s.TableName
}

func (s *SubmitSqlFlashbackTaskRequest) GetTraceId() *string {
	return s.TraceId
}

func (s *SubmitSqlFlashbackTaskRequest) SetDbName(v string) *SubmitSqlFlashbackTaskRequest {
	s.DbName = &v
	return s
}

func (s *SubmitSqlFlashbackTaskRequest) SetDrdsInstanceId(v string) *SubmitSqlFlashbackTaskRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *SubmitSqlFlashbackTaskRequest) SetEndTime(v string) *SubmitSqlFlashbackTaskRequest {
	s.EndTime = &v
	return s
}

func (s *SubmitSqlFlashbackTaskRequest) SetRecallRestoreType(v int32) *SubmitSqlFlashbackTaskRequest {
	s.RecallRestoreType = &v
	return s
}

func (s *SubmitSqlFlashbackTaskRequest) SetRecallType(v int32) *SubmitSqlFlashbackTaskRequest {
	s.RecallType = &v
	return s
}

func (s *SubmitSqlFlashbackTaskRequest) SetSqlPk(v string) *SubmitSqlFlashbackTaskRequest {
	s.SqlPk = &v
	return s
}

func (s *SubmitSqlFlashbackTaskRequest) SetSqlType(v string) *SubmitSqlFlashbackTaskRequest {
	s.SqlType = &v
	return s
}

func (s *SubmitSqlFlashbackTaskRequest) SetStartTime(v string) *SubmitSqlFlashbackTaskRequest {
	s.StartTime = &v
	return s
}

func (s *SubmitSqlFlashbackTaskRequest) SetTableName(v string) *SubmitSqlFlashbackTaskRequest {
	s.TableName = &v
	return s
}

func (s *SubmitSqlFlashbackTaskRequest) SetTraceId(v string) *SubmitSqlFlashbackTaskRequest {
	s.TraceId = &v
	return s
}

func (s *SubmitSqlFlashbackTaskRequest) Validate() error {
	return dara.Validate(s)
}
