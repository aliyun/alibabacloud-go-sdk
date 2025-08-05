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
	Data *GetMmsTaskResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
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
	return dara.Validate(s)
}

type GetMmsTaskResponseBodyData struct {
	// example:
	//
	// 2024-10-25 04:21:01
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// 23
	DbId *int64 `json:"dbId,omitempty" xml:"dbId,omitempty"`
	// example:
	//
	// mms_target
	DstDbName *string `json:"dstDbName,omitempty" xml:"dstDbName,omitempty"`
	// example:
	//
	// default
	DstSchemaName *string `json:"dstSchemaName,omitempty" xml:"dstSchemaName,omitempty"`
	// example:
	//
	// table_1
	DstTableName *string `json:"dstTableName,omitempty" xml:"dstTableName,omitempty"`
	// example:
	//
	// 2024-10-25 07:21:01
	EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// example:
	//
	// 7680
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// 87
	JobId *int64 `json:"jobId,omitempty" xml:"jobId,omitempty"`
	// example:
	//
	// test_odps_spark
	JobName *string `json:"jobName,omitempty" xml:"jobName,omitempty"`
	// example:
	//
	// 1
	RetriedTimes *int32 `json:"retriedTimes,omitempty" xml:"retriedTimes,omitempty"`
	// example:
	//
	// true
	Running *bool `json:"running,omitempty" xml:"running,omitempty"`
	// example:
	//
	// 2000015
	SourceId *int64 `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	// example:
	//
	// demo
	SourceName *string `json:"sourceName,omitempty" xml:"sourceName,omitempty"`
	// example:
	//
	// mms_test
	SrcDbName *string `json:"srcDbName,omitempty" xml:"srcDbName,omitempty"`
	// example:
	//
	// default
	SrcSchemaName *string `json:"srcSchemaName,omitempty" xml:"srcSchemaName,omitempty"`
	// example:
	//
	// table_1
	SrcTableName *string `json:"srcTableName,omitempty" xml:"srcTableName,omitempty"`
	// example:
	//
	// 2024-10-25 06:21:01
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// example:
	//
	// DATA_DOING
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// false
	Stopped *bool `json:"stopped,omitempty" xml:"stopped,omitempty"`
	// example:
	//
	// 2323
	TableId *int64 `json:"tableId,omitempty" xml:"tableId,omitempty"`
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
