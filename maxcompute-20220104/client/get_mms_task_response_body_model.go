// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMmsTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetMmsTaskResponseBodyData) *GetMmsTaskResponseBody
	GetData() *GetMmsTaskResponseBodyData
	SetRequestId(v string) *GetMmsTaskResponseBody
	GetRequestId() *string
}

type GetMmsTaskResponseBody struct {
	// The migration task object.
	Data *GetMmsTaskResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 73207140-0FD5-588A-B11A-3CE093924196
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetMmsTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetMmsTaskResponseBody) GoString() string {
	return s.String()
}

func (s *GetMmsTaskResponseBody) GetData() *GetMmsTaskResponseBodyData {
	return s.Data
}

func (s *GetMmsTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetMmsTaskResponseBody) SetData(v *GetMmsTaskResponseBodyData) *GetMmsTaskResponseBody {
	s.Data = v
	return s
}

func (s *GetMmsTaskResponseBody) SetRequestId(v string) *GetMmsTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMmsTaskResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetMmsTaskResponseBodyData struct {
	// The creation time of the task.
	//
	// example:
	//
	// 2024-10-25 04:21:01
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// The source database ID.
	//
	// example:
	//
	// 23
	DbId *int64 `json:"dbId,omitempty" xml:"dbId,omitempty"`
	// The destination MaxCompute project.
	//
	// example:
	//
	// mms_target
	DstDbName *string `json:"dstDbName,omitempty" xml:"dstDbName,omitempty"`
	// The destination MaxCompute schema.
	//
	// example:
	//
	// default
	DstSchemaName *string `json:"dstSchemaName,omitempty" xml:"dstSchemaName,omitempty"`
	// The destination MaxCompute table.
	//
	// example:
	//
	// table_1
	DstTableName *string `json:"dstTableName,omitempty" xml:"dstTableName,omitempty"`
	// The end time of the task.
	//
	// example:
	//
	// 2024-10-25 07:21:01
	EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// The migration task ID.
	//
	// example:
	//
	// 7680
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// The migration job ID.
	//
	// example:
	//
	// 87
	JobId *int64 `json:"jobId,omitempty" xml:"jobId,omitempty"`
	// The migration job name.
	//
	// example:
	//
	// test_odps_spark
	JobName *string `json:"jobName,omitempty" xml:"jobName,omitempty"`
	// The number of times the task has been retried.
	//
	// example:
	//
	// 1
	RetriedTimes *int32 `json:"retriedTimes,omitempty" xml:"retriedTimes,omitempty"`
	// Indicates if the task is running.
	//
	// example:
	//
	// true
	Running *bool `json:"running,omitempty" xml:"running,omitempty"`
	// The data source ID.
	//
	// example:
	//
	// 2000015
	SourceId *int64 `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	// The data source name.
	//
	// example:
	//
	// demo
	SourceName *string `json:"sourceName,omitempty" xml:"sourceName,omitempty"`
	// The source database name.
	//
	// example:
	//
	// mms_test
	SrcDbName *string `json:"srcDbName,omitempty" xml:"srcDbName,omitempty"`
	// The name of the source schema. This refers to the schema in a Layer 3 namespace.
	//
	// example:
	//
	// default
	SrcSchemaName *string `json:"srcSchemaName,omitempty" xml:"srcSchemaName,omitempty"`
	// The source table name.
	//
	// example:
	//
	// table_1
	SrcTableName *string `json:"srcTableName,omitempty" xml:"srcTableName,omitempty"`
	// The start time of the task.
	//
	// example:
	//
	// 2024-10-25 06:21:01
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// The migration task status.
	//
	// example:
	//
	// DATA_DOING
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// Indicates if the task is stopped.
	//
	// example:
	//
	// false
	Stopped *bool `json:"stopped,omitempty" xml:"stopped,omitempty"`
	// The source table ID.
	//
	// example:
	//
	// 2323
	TableId *int64 `json:"tableId,omitempty" xml:"tableId,omitempty"`
	// The task type.
	//
	// example:
	//
	// BIGQUERY
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetMmsTaskResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetMmsTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetMmsTaskResponseBodyData) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetMmsTaskResponseBodyData) GetDbId() *int64 {
	return s.DbId
}

func (s *GetMmsTaskResponseBodyData) GetDstDbName() *string {
	return s.DstDbName
}

func (s *GetMmsTaskResponseBodyData) GetDstSchemaName() *string {
	return s.DstSchemaName
}

func (s *GetMmsTaskResponseBodyData) GetDstTableName() *string {
	return s.DstTableName
}

func (s *GetMmsTaskResponseBodyData) GetEndTime() *string {
	return s.EndTime
}

func (s *GetMmsTaskResponseBodyData) GetId() *int64 {
	return s.Id
}

func (s *GetMmsTaskResponseBodyData) GetJobId() *int64 {
	return s.JobId
}

func (s *GetMmsTaskResponseBodyData) GetJobName() *string {
	return s.JobName
}

func (s *GetMmsTaskResponseBodyData) GetRetriedTimes() *int32 {
	return s.RetriedTimes
}

func (s *GetMmsTaskResponseBodyData) GetRunning() *bool {
	return s.Running
}

func (s *GetMmsTaskResponseBodyData) GetSourceId() *int64 {
	return s.SourceId
}

func (s *GetMmsTaskResponseBodyData) GetSourceName() *string {
	return s.SourceName
}

func (s *GetMmsTaskResponseBodyData) GetSrcDbName() *string {
	return s.SrcDbName
}

func (s *GetMmsTaskResponseBodyData) GetSrcSchemaName() *string {
	return s.SrcSchemaName
}

func (s *GetMmsTaskResponseBodyData) GetSrcTableName() *string {
	return s.SrcTableName
}

func (s *GetMmsTaskResponseBodyData) GetStartTime() *string {
	return s.StartTime
}

func (s *GetMmsTaskResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *GetMmsTaskResponseBodyData) GetStopped() *bool {
	return s.Stopped
}

func (s *GetMmsTaskResponseBodyData) GetTableId() *int64 {
	return s.TableId
}

func (s *GetMmsTaskResponseBodyData) GetType() *string {
	return s.Type
}

func (s *GetMmsTaskResponseBodyData) SetCreateTime(v string) *GetMmsTaskResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *GetMmsTaskResponseBodyData) SetDbId(v int64) *GetMmsTaskResponseBodyData {
	s.DbId = &v
	return s
}

func (s *GetMmsTaskResponseBodyData) SetDstDbName(v string) *GetMmsTaskResponseBodyData {
	s.DstDbName = &v
	return s
}

func (s *GetMmsTaskResponseBodyData) SetDstSchemaName(v string) *GetMmsTaskResponseBodyData {
	s.DstSchemaName = &v
	return s
}

func (s *GetMmsTaskResponseBodyData) SetDstTableName(v string) *GetMmsTaskResponseBodyData {
	s.DstTableName = &v
	return s
}

func (s *GetMmsTaskResponseBodyData) SetEndTime(v string) *GetMmsTaskResponseBodyData {
	s.EndTime = &v
	return s
}

func (s *GetMmsTaskResponseBodyData) SetId(v int64) *GetMmsTaskResponseBodyData {
	s.Id = &v
	return s
}

func (s *GetMmsTaskResponseBodyData) SetJobId(v int64) *GetMmsTaskResponseBodyData {
	s.JobId = &v
	return s
}

func (s *GetMmsTaskResponseBodyData) SetJobName(v string) *GetMmsTaskResponseBodyData {
	s.JobName = &v
	return s
}

func (s *GetMmsTaskResponseBodyData) SetRetriedTimes(v int32) *GetMmsTaskResponseBodyData {
	s.RetriedTimes = &v
	return s
}

func (s *GetMmsTaskResponseBodyData) SetRunning(v bool) *GetMmsTaskResponseBodyData {
	s.Running = &v
	return s
}

func (s *GetMmsTaskResponseBodyData) SetSourceId(v int64) *GetMmsTaskResponseBodyData {
	s.SourceId = &v
	return s
}

func (s *GetMmsTaskResponseBodyData) SetSourceName(v string) *GetMmsTaskResponseBodyData {
	s.SourceName = &v
	return s
}

func (s *GetMmsTaskResponseBodyData) SetSrcDbName(v string) *GetMmsTaskResponseBodyData {
	s.SrcDbName = &v
	return s
}

func (s *GetMmsTaskResponseBodyData) SetSrcSchemaName(v string) *GetMmsTaskResponseBodyData {
	s.SrcSchemaName = &v
	return s
}

func (s *GetMmsTaskResponseBodyData) SetSrcTableName(v string) *GetMmsTaskResponseBodyData {
	s.SrcTableName = &v
	return s
}

func (s *GetMmsTaskResponseBodyData) SetStartTime(v string) *GetMmsTaskResponseBodyData {
	s.StartTime = &v
	return s
}

func (s *GetMmsTaskResponseBodyData) SetStatus(v string) *GetMmsTaskResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetMmsTaskResponseBodyData) SetStopped(v bool) *GetMmsTaskResponseBodyData {
	s.Stopped = &v
	return s
}

func (s *GetMmsTaskResponseBodyData) SetTableId(v int64) *GetMmsTaskResponseBodyData {
	s.TableId = &v
	return s
}

func (s *GetMmsTaskResponseBodyData) SetType(v string) *GetMmsTaskResponseBodyData {
	s.Type = &v
	return s
}

func (s *GetMmsTaskResponseBodyData) Validate() error {
	return dara.Validate(s)
}
