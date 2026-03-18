// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMmsTasksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ListMmsTasksResponseBodyData) *ListMmsTasksResponseBody
	GetData() *ListMmsTasksResponseBodyData
	SetRequestId(v string) *ListMmsTasksResponseBody
	GetRequestId() *string
}

type ListMmsTasksResponseBody struct {
	// The returned data.
	Data *ListMmsTasksResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 373A5CB2-8570-53BE-A98F-729B11D7A8B0
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListMmsTasksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListMmsTasksResponseBody) GoString() string {
	return s.String()
}

func (s *ListMmsTasksResponseBody) GetData() *ListMmsTasksResponseBodyData {
	return s.Data
}

func (s *ListMmsTasksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListMmsTasksResponseBody) SetData(v *ListMmsTasksResponseBodyData) *ListMmsTasksResponseBody {
	s.Data = v
	return s
}

func (s *ListMmsTasksResponseBody) SetRequestId(v string) *ListMmsTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMmsTasksResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListMmsTasksResponseBodyData struct {
	// The list of migration tasks.
	ObjectList []*ListMmsTasksResponseBodyDataObjectList `json:"objectList,omitempty" xml:"objectList,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNum *int32 `json:"pageNum,omitempty" xml:"pageNum,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The total number of records.
	//
	// example:
	//
	// 100
	Total *int32 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListMmsTasksResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListMmsTasksResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListMmsTasksResponseBodyData) GetObjectList() []*ListMmsTasksResponseBodyDataObjectList {
	return s.ObjectList
}

func (s *ListMmsTasksResponseBodyData) GetPageNum() *int32 {
	return s.PageNum
}

func (s *ListMmsTasksResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListMmsTasksResponseBodyData) GetTotal() *int32 {
	return s.Total
}

func (s *ListMmsTasksResponseBodyData) SetObjectList(v []*ListMmsTasksResponseBodyDataObjectList) *ListMmsTasksResponseBodyData {
	s.ObjectList = v
	return s
}

func (s *ListMmsTasksResponseBodyData) SetPageNum(v int32) *ListMmsTasksResponseBodyData {
	s.PageNum = &v
	return s
}

func (s *ListMmsTasksResponseBodyData) SetPageSize(v int32) *ListMmsTasksResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListMmsTasksResponseBodyData) SetTotal(v int32) *ListMmsTasksResponseBodyData {
	s.Total = &v
	return s
}

func (s *ListMmsTasksResponseBodyData) Validate() error {
	if s.ObjectList != nil {
		for _, item := range s.ObjectList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListMmsTasksResponseBodyDataObjectList struct {
	// The time when the task was created.
	//
	// example:
	//
	// 2024-10-25 04:21:01
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// The ID of the source database.
	//
	// example:
	//
	// 196
	DbId *int64 `json:"dbId,omitempty" xml:"dbId,omitempty"`
	// The destination MaxCompute project.
	//
	// example:
	//
	// mms_test
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
	// The ID of the migration task.
	//
	// example:
	//
	// 2323
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// The ID of the migration job.
	//
	// example:
	//
	// 87
	JobId *int64 `json:"jobId,omitempty" xml:"jobId,omitempty"`
	// The name of the migration job.
	//
	// example:
	//
	// test_odps_spark
	JobName *string `json:"jobName,omitempty" xml:"jobName,omitempty"`
	// The number of retries.
	//
	// example:
	//
	// 1
	RetriedTimes *int32 `json:"retriedTimes,omitempty" xml:"retriedTimes,omitempty"`
	// Indicates whether the task is running.
	//
	// example:
	//
	// true
	Running *bool `json:"running,omitempty" xml:"running,omitempty"`
	// The ID of the data source.
	//
	// example:
	//
	// 2000028
	SourceId *int64 `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	// The name of the data source.
	//
	// example:
	//
	// demo
	SourceName *string `json:"sourceName,omitempty" xml:"sourceName,omitempty"`
	// The name of the source database.
	//
	// example:
	//
	// db_1
	SrcDbName *string `json:"srcDbName,omitempty" xml:"srcDbName,omitempty"`
	// The name of the source schema. This parameter specifies the schema in a Layer 3 namespace.
	//
	// example:
	//
	// default
	SrcSchemaName *string `json:"srcSchemaName,omitempty" xml:"srcSchemaName,omitempty"`
	// The name of the source table.
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
	// The status of the migration task.
	//
	// example:
	//
	// DATA_DOING
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// Indicates whether the task is stopped.
	//
	// example:
	//
	// false
	Stopped *bool `json:"stopped,omitempty" xml:"stopped,omitempty"`
	// The ID of the source table.
	//
	// example:
	//
	// 23
	TableId *int64 `json:"tableId,omitempty" xml:"tableId,omitempty"`
	// The task type.
	//
	// example:
	//
	// BIGQUERY
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListMmsTasksResponseBodyDataObjectList) String() string {
	return dara.Prettify(s)
}

func (s ListMmsTasksResponseBodyDataObjectList) GoString() string {
	return s.String()
}

func (s *ListMmsTasksResponseBodyDataObjectList) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListMmsTasksResponseBodyDataObjectList) GetDbId() *int64 {
	return s.DbId
}

func (s *ListMmsTasksResponseBodyDataObjectList) GetDstDbName() *string {
	return s.DstDbName
}

func (s *ListMmsTasksResponseBodyDataObjectList) GetDstSchemaName() *string {
	return s.DstSchemaName
}

func (s *ListMmsTasksResponseBodyDataObjectList) GetDstTableName() *string {
	return s.DstTableName
}

func (s *ListMmsTasksResponseBodyDataObjectList) GetEndTime() *string {
	return s.EndTime
}

func (s *ListMmsTasksResponseBodyDataObjectList) GetId() *int64 {
	return s.Id
}

func (s *ListMmsTasksResponseBodyDataObjectList) GetJobId() *int64 {
	return s.JobId
}

func (s *ListMmsTasksResponseBodyDataObjectList) GetJobName() *string {
	return s.JobName
}

func (s *ListMmsTasksResponseBodyDataObjectList) GetRetriedTimes() *int32 {
	return s.RetriedTimes
}

func (s *ListMmsTasksResponseBodyDataObjectList) GetRunning() *bool {
	return s.Running
}

func (s *ListMmsTasksResponseBodyDataObjectList) GetSourceId() *int64 {
	return s.SourceId
}

func (s *ListMmsTasksResponseBodyDataObjectList) GetSourceName() *string {
	return s.SourceName
}

func (s *ListMmsTasksResponseBodyDataObjectList) GetSrcDbName() *string {
	return s.SrcDbName
}

func (s *ListMmsTasksResponseBodyDataObjectList) GetSrcSchemaName() *string {
	return s.SrcSchemaName
}

func (s *ListMmsTasksResponseBodyDataObjectList) GetSrcTableName() *string {
	return s.SrcTableName
}

func (s *ListMmsTasksResponseBodyDataObjectList) GetStartTime() *string {
	return s.StartTime
}

func (s *ListMmsTasksResponseBodyDataObjectList) GetStatus() *string {
	return s.Status
}

func (s *ListMmsTasksResponseBodyDataObjectList) GetStopped() *bool {
	return s.Stopped
}

func (s *ListMmsTasksResponseBodyDataObjectList) GetTableId() *int64 {
	return s.TableId
}

func (s *ListMmsTasksResponseBodyDataObjectList) GetType() *string {
	return s.Type
}

func (s *ListMmsTasksResponseBodyDataObjectList) SetCreateTime(v string) *ListMmsTasksResponseBodyDataObjectList {
	s.CreateTime = &v
	return s
}

func (s *ListMmsTasksResponseBodyDataObjectList) SetDbId(v int64) *ListMmsTasksResponseBodyDataObjectList {
	s.DbId = &v
	return s
}

func (s *ListMmsTasksResponseBodyDataObjectList) SetDstDbName(v string) *ListMmsTasksResponseBodyDataObjectList {
	s.DstDbName = &v
	return s
}

func (s *ListMmsTasksResponseBodyDataObjectList) SetDstSchemaName(v string) *ListMmsTasksResponseBodyDataObjectList {
	s.DstSchemaName = &v
	return s
}

func (s *ListMmsTasksResponseBodyDataObjectList) SetDstTableName(v string) *ListMmsTasksResponseBodyDataObjectList {
	s.DstTableName = &v
	return s
}

func (s *ListMmsTasksResponseBodyDataObjectList) SetEndTime(v string) *ListMmsTasksResponseBodyDataObjectList {
	s.EndTime = &v
	return s
}

func (s *ListMmsTasksResponseBodyDataObjectList) SetId(v int64) *ListMmsTasksResponseBodyDataObjectList {
	s.Id = &v
	return s
}

func (s *ListMmsTasksResponseBodyDataObjectList) SetJobId(v int64) *ListMmsTasksResponseBodyDataObjectList {
	s.JobId = &v
	return s
}

func (s *ListMmsTasksResponseBodyDataObjectList) SetJobName(v string) *ListMmsTasksResponseBodyDataObjectList {
	s.JobName = &v
	return s
}

func (s *ListMmsTasksResponseBodyDataObjectList) SetRetriedTimes(v int32) *ListMmsTasksResponseBodyDataObjectList {
	s.RetriedTimes = &v
	return s
}

func (s *ListMmsTasksResponseBodyDataObjectList) SetRunning(v bool) *ListMmsTasksResponseBodyDataObjectList {
	s.Running = &v
	return s
}

func (s *ListMmsTasksResponseBodyDataObjectList) SetSourceId(v int64) *ListMmsTasksResponseBodyDataObjectList {
	s.SourceId = &v
	return s
}

func (s *ListMmsTasksResponseBodyDataObjectList) SetSourceName(v string) *ListMmsTasksResponseBodyDataObjectList {
	s.SourceName = &v
	return s
}

func (s *ListMmsTasksResponseBodyDataObjectList) SetSrcDbName(v string) *ListMmsTasksResponseBodyDataObjectList {
	s.SrcDbName = &v
	return s
}

func (s *ListMmsTasksResponseBodyDataObjectList) SetSrcSchemaName(v string) *ListMmsTasksResponseBodyDataObjectList {
	s.SrcSchemaName = &v
	return s
}

func (s *ListMmsTasksResponseBodyDataObjectList) SetSrcTableName(v string) *ListMmsTasksResponseBodyDataObjectList {
	s.SrcTableName = &v
	return s
}

func (s *ListMmsTasksResponseBodyDataObjectList) SetStartTime(v string) *ListMmsTasksResponseBodyDataObjectList {
	s.StartTime = &v
	return s
}

func (s *ListMmsTasksResponseBodyDataObjectList) SetStatus(v string) *ListMmsTasksResponseBodyDataObjectList {
	s.Status = &v
	return s
}

func (s *ListMmsTasksResponseBodyDataObjectList) SetStopped(v bool) *ListMmsTasksResponseBodyDataObjectList {
	s.Stopped = &v
	return s
}

func (s *ListMmsTasksResponseBodyDataObjectList) SetTableId(v int64) *ListMmsTasksResponseBodyDataObjectList {
	s.TableId = &v
	return s
}

func (s *ListMmsTasksResponseBodyDataObjectList) SetType(v string) *ListMmsTasksResponseBodyDataObjectList {
	s.Type = &v
	return s
}

func (s *ListMmsTasksResponseBodyDataObjectList) Validate() error {
	return dara.Validate(s)
}
