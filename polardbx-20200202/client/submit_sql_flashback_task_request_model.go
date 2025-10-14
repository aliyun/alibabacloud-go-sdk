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
	SetEndTime(v string) *SubmitSqlFlashbackTaskRequest
	GetEndTime() *string
	SetPolardbxInstanceId(v string) *SubmitSqlFlashbackTaskRequest
	GetPolardbxInstanceId() *string
	SetRecallRestoreType(v string) *SubmitSqlFlashbackTaskRequest
	GetRecallRestoreType() *string
	SetRecallType(v string) *SubmitSqlFlashbackTaskRequest
	GetRecallType() *string
	SetRegionId(v string) *SubmitSqlFlashbackTaskRequest
	GetRegionId() *string
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
	// This parameter is required.
	//
	// example:
	//
	// test
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2019-09-10 23:23:23
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pxc-*********
	PolardbxInstanceId *string `json:"PolardbxInstanceId,omitempty" xml:"PolardbxInstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0
	RecallRestoreType *string `json:"RecallRestoreType,omitempty" xml:"RecallRestoreType,omitempty"`
	// if can be null:
	// true
	//
	// example:
	//
	// 0
	RecallType *string `json:"RecallType,omitempty" xml:"RecallType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// if can be null:
	// true
	//
	// example:
	//
	// 1111
	SqlPk *string `json:"SqlPk,omitempty" xml:"SqlPk,omitempty"`
	// if can be null:
	// true
	//
	// example:
	//
	// INSERT,UPDATE
	SqlType *string `json:"SqlType,omitempty" xml:"SqlType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2019-09-10 20:23:23
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// if can be null:
	// true
	//
	// example:
	//
	// test
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	// if can be null:
	// true
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

func (s *SubmitSqlFlashbackTaskRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *SubmitSqlFlashbackTaskRequest) GetPolardbxInstanceId() *string {
	return s.PolardbxInstanceId
}

func (s *SubmitSqlFlashbackTaskRequest) GetRecallRestoreType() *string {
	return s.RecallRestoreType
}

func (s *SubmitSqlFlashbackTaskRequest) GetRecallType() *string {
	return s.RecallType
}

func (s *SubmitSqlFlashbackTaskRequest) GetRegionId() *string {
	return s.RegionId
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

func (s *SubmitSqlFlashbackTaskRequest) SetEndTime(v string) *SubmitSqlFlashbackTaskRequest {
	s.EndTime = &v
	return s
}

func (s *SubmitSqlFlashbackTaskRequest) SetPolardbxInstanceId(v string) *SubmitSqlFlashbackTaskRequest {
	s.PolardbxInstanceId = &v
	return s
}

func (s *SubmitSqlFlashbackTaskRequest) SetRecallRestoreType(v string) *SubmitSqlFlashbackTaskRequest {
	s.RecallRestoreType = &v
	return s
}

func (s *SubmitSqlFlashbackTaskRequest) SetRecallType(v string) *SubmitSqlFlashbackTaskRequest {
	s.RecallType = &v
	return s
}

func (s *SubmitSqlFlashbackTaskRequest) SetRegionId(v string) *SubmitSqlFlashbackTaskRequest {
	s.RegionId = &v
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
