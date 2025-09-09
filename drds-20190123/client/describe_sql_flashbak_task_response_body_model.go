// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSqlFlashbakTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeSqlFlashbakTaskResponseBody
	GetRequestId() *string
	SetSqlFlashbackTasks(v *DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasks) *DescribeSqlFlashbakTaskResponseBody
	GetSqlFlashbackTasks() *DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasks
	SetSuccess(v bool) *DescribeSqlFlashbakTaskResponseBody
	GetSuccess() *bool
}

type DescribeSqlFlashbakTaskResponseBody struct {
	// Indicates the ID of the request.
	//
	// example:
	//
	// 5D64DE59-44A1-E541-E0CB-B7E5C4305162
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates the information about flashback tasks.
	SqlFlashbackTasks *DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasks `json:"SqlFlashbackTasks,omitempty" xml:"SqlFlashbackTasks,omitempty" type:"Struct"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeSqlFlashbakTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSqlFlashbakTaskResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSqlFlashbakTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSqlFlashbakTaskResponseBody) GetSqlFlashbackTasks() *DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasks {
	return s.SqlFlashbackTasks
}

func (s *DescribeSqlFlashbakTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeSqlFlashbakTaskResponseBody) SetRequestId(v string) *DescribeSqlFlashbakTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSqlFlashbakTaskResponseBody) SetSqlFlashbackTasks(v *DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasks) *DescribeSqlFlashbakTaskResponseBody {
	s.SqlFlashbackTasks = v
	return s
}

func (s *DescribeSqlFlashbakTaskResponseBody) SetSuccess(v bool) *DescribeSqlFlashbakTaskResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeSqlFlashbakTaskResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasks struct {
	SqlFlashbackTask []*DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasksSqlFlashbackTask `json:"SqlFlashbackTask,omitempty" xml:"SqlFlashbackTask,omitempty" type:"Repeated"`
}

func (s DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasks) String() string {
	return dara.Prettify(s)
}

func (s DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasks) GoString() string {
	return s.String()
}

func (s *DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasks) GetSqlFlashbackTask() []*DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasksSqlFlashbackTask {
	return s.SqlFlashbackTask
}

func (s *DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasks) SetSqlFlashbackTask(v []*DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasksSqlFlashbackTask) *DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasks {
	s.SqlFlashbackTask = v
	return s
}

func (s *DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasks) Validate() error {
	return dara.Validate(s)
}

type DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasksSqlFlashbackTask struct {
	// Indicates the name of the database on which a flashback task is performed.
	//
	// example:
	//
	// drds_flashback
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// Indicates the download URL.
	//
	// example:
	//
	// http://...
	DownloadUrl *string `json:"DownloadUrl,omitempty" xml:"DownloadUrl,omitempty"`
	// Indicates the time when the download URL expires.
	//
	// example:
	//
	// 1569216270000
	ExpireTime *int64 `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// Indicates the point in time when the instance was created.
	//
	// example:
	//
	// 1568611408000
	GmtCreate *int64 `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// Indicates the point in time when the flashback task is performed.
	//
	// example:
	//
	// 1568611469000
	GmtModified *int64 `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// Indicates the ID of the primary key that corresponds to a table used in the flashback task.
	//
	// example:
	//
	// 238
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// Indicates the ID of the instance.
	//
	// example:
	//
	// drdshbga76p6****
	InstId *string `json:"InstId,omitempty" xml:"InstId,omitempty"`
	// Indicates the progress of the reverse call.
	//
	// example:
	//
	// 10
	RecallProgress *int32 `json:"RecallProgress,omitempty" xml:"RecallProgress,omitempty"`
	// Indicates the type of the flashback task. Valid values:
	//
	// 	- **1**: image restoration
	//
	// 	- **2**: reverse restoration
	//
	// example:
	//
	// 1
	RecallRestoreType *int32 `json:"RecallRestoreType,omitempty" xml:"RecallRestoreType,omitempty"`
	// Indicates the status of the data recall task.
	//
	// example:
	//
	// 1
	RecallStatus *int32 `json:"RecallStatus,omitempty" xml:"RecallStatus,omitempty"`
	// Indicates the type of the reverse call. Valid values:
	//
	// 	- **0**: exact search
	//
	// 	- **1**: fuzzy search
	//
	// example:
	//
	// 1
	RecallType *int32 `json:"RecallType,omitempty" xml:"RecallType,omitempty"`
	// Indicates the start time of the reverse call.
	//
	// example:
	//
	// 1568509597000
	SearchEndTime *int64 `json:"SearchEndTime,omitempty" xml:"SearchEndTime,omitempty"`
	// Indicates the end time of the reverse call.
	//
	// example:
	//
	// 1568609597000
	SearchStartTime *int64 `json:"SearchStartTime,omitempty" xml:"SearchStartTime,omitempty"`
	// Indicates the number of data rows that are flashed back.
	//
	// example:
	//
	// 0
	SqlCounter *int64 `json:"SqlCounter,omitempty" xml:"SqlCounter,omitempty"`
	// Indicates the primary key specified in the SQL statements.
	//
	// example:
	//
	// id
	SqlPk *string `json:"SqlPk,omitempty" xml:"SqlPk,omitempty"`
	// Indicates the types of the SQL statements.
	//
	// example:
	//
	// Insert,Update,Delete
	SqlType *string `json:"SqlType,omitempty" xml:"SqlType,omitempty"`
	// Indicates the name of the table that contains the data that are flashed back.
	//
	// example:
	//
	// drds_params
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	// Indicates the ID of the trace of the SQL query.
	//
	// example:
	//
	// trace
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasksSqlFlashbackTask) String() string {
	return dara.Prettify(s)
}

func (s DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasksSqlFlashbackTask) GoString() string {
	return s.String()
}

func (s *DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasksSqlFlashbackTask) GetDbName() *string {
	return s.DbName
}

func (s *DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasksSqlFlashbackTask) GetDownloadUrl() *string {
	return s.DownloadUrl
}

func (s *DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasksSqlFlashbackTask) GetExpireTime() *int64 {
	return s.ExpireTime
}

func (s *DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasksSqlFlashbackTask) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasksSqlFlashbackTask) GetGmtModified() *int64 {
	return s.GmtModified
}

func (s *DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasksSqlFlashbackTask) GetId() *int64 {
	return s.Id
}

func (s *DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasksSqlFlashbackTask) GetInstId() *string {
	return s.InstId
}

func (s *DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasksSqlFlashbackTask) GetRecallProgress() *int32 {
	return s.RecallProgress
}

func (s *DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasksSqlFlashbackTask) GetRecallRestoreType() *int32 {
	return s.RecallRestoreType
}

func (s *DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasksSqlFlashbackTask) GetRecallStatus() *int32 {
	return s.RecallStatus
}

func (s *DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasksSqlFlashbackTask) GetRecallType() *int32 {
	return s.RecallType
}

func (s *DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasksSqlFlashbackTask) GetSearchEndTime() *int64 {
	return s.SearchEndTime
}

func (s *DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasksSqlFlashbackTask) GetSearchStartTime() *int64 {
	return s.SearchStartTime
}

func (s *DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasksSqlFlashbackTask) GetSqlCounter() *int64 {
	return s.SqlCounter
}

func (s *DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasksSqlFlashbackTask) GetSqlPk() *string {
	return s.SqlPk
}

func (s *DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasksSqlFlashbackTask) GetSqlType() *string {
	return s.SqlType
}

func (s *DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasksSqlFlashbackTask) GetTableName() *string {
	return s.TableName
}

func (s *DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasksSqlFlashbackTask) GetTraceId() *string {
	return s.TraceId
}

func (s *DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasksSqlFlashbackTask) SetDbName(v string) *DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasksSqlFlashbackTask {
	s.DbName = &v
	return s
}

func (s *DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasksSqlFlashbackTask) SetDownloadUrl(v string) *DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasksSqlFlashbackTask {
	s.DownloadUrl = &v
	return s
}

func (s *DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasksSqlFlashbackTask) SetExpireTime(v int64) *DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasksSqlFlashbackTask {
	s.ExpireTime = &v
	return s
}

func (s *DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasksSqlFlashbackTask) SetGmtCreate(v int64) *DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasksSqlFlashbackTask {
	s.GmtCreate = &v
	return s
}

func (s *DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasksSqlFlashbackTask) SetGmtModified(v int64) *DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasksSqlFlashbackTask {
	s.GmtModified = &v
	return s
}

func (s *DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasksSqlFlashbackTask) SetId(v int64) *DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasksSqlFlashbackTask {
	s.Id = &v
	return s
}

func (s *DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasksSqlFlashbackTask) SetInstId(v string) *DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasksSqlFlashbackTask {
	s.InstId = &v
	return s
}

func (s *DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasksSqlFlashbackTask) SetRecallProgress(v int32) *DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasksSqlFlashbackTask {
	s.RecallProgress = &v
	return s
}

func (s *DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasksSqlFlashbackTask) SetRecallRestoreType(v int32) *DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasksSqlFlashbackTask {
	s.RecallRestoreType = &v
	return s
}

func (s *DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasksSqlFlashbackTask) SetRecallStatus(v int32) *DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasksSqlFlashbackTask {
	s.RecallStatus = &v
	return s
}

func (s *DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasksSqlFlashbackTask) SetRecallType(v int32) *DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasksSqlFlashbackTask {
	s.RecallType = &v
	return s
}

func (s *DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasksSqlFlashbackTask) SetSearchEndTime(v int64) *DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasksSqlFlashbackTask {
	s.SearchEndTime = &v
	return s
}

func (s *DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasksSqlFlashbackTask) SetSearchStartTime(v int64) *DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasksSqlFlashbackTask {
	s.SearchStartTime = &v
	return s
}

func (s *DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasksSqlFlashbackTask) SetSqlCounter(v int64) *DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasksSqlFlashbackTask {
	s.SqlCounter = &v
	return s
}

func (s *DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasksSqlFlashbackTask) SetSqlPk(v string) *DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasksSqlFlashbackTask {
	s.SqlPk = &v
	return s
}

func (s *DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasksSqlFlashbackTask) SetSqlType(v string) *DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasksSqlFlashbackTask {
	s.SqlType = &v
	return s
}

func (s *DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasksSqlFlashbackTask) SetTableName(v string) *DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasksSqlFlashbackTask {
	s.TableName = &v
	return s
}

func (s *DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasksSqlFlashbackTask) SetTraceId(v string) *DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasksSqlFlashbackTask {
	s.TraceId = &v
	return s
}

func (s *DescribeSqlFlashbakTaskResponseBodySqlFlashbackTasksSqlFlashbackTask) Validate() error {
	return dara.Validate(s)
}
