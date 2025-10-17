// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListResultExportJobHistoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListResultExportJobHistoryResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ListResultExportJobHistoryResponseBody
	GetHttpStatusCode() *int32
	SetItems(v []*ListResultExportJobHistoryResponseBodyItems) *ListResultExportJobHistoryResponseBody
	GetItems() []*ListResultExportJobHistoryResponseBodyItems
	SetMessage(v string) *ListResultExportJobHistoryResponseBody
	GetMessage() *string
	SetPageNumber(v int32) *ListResultExportJobHistoryResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListResultExportJobHistoryResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListResultExportJobHistoryResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListResultExportJobHistoryResponseBody
	GetSuccess() *bool
	SetTotalCount(v int64) *ListResultExportJobHistoryResponseBody
	GetTotalCount() *int64
}

type ListResultExportJobHistoryResponseBody struct {
	// The HTTP status code.
	//
	// example:
	//
	// InvalidInput
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response code. The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The queried result set export jobs.
	Items []*ListResultExportJobHistoryResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The returned message. Valid values:
	//
	// 	- If the request was successful, an **OK*	- message is returned.
	//
	// 	- If the request failed, an error message is returned.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1AD222E9-E606-4A42-BF6D-8A4442913CEF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 174
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListResultExportJobHistoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListResultExportJobHistoryResponseBody) GoString() string {
	return s.String()
}

func (s *ListResultExportJobHistoryResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListResultExportJobHistoryResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListResultExportJobHistoryResponseBody) GetItems() []*ListResultExportJobHistoryResponseBodyItems {
	return s.Items
}

func (s *ListResultExportJobHistoryResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListResultExportJobHistoryResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListResultExportJobHistoryResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListResultExportJobHistoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListResultExportJobHistoryResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListResultExportJobHistoryResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListResultExportJobHistoryResponseBody) SetCode(v string) *ListResultExportJobHistoryResponseBody {
	s.Code = &v
	return s
}

func (s *ListResultExportJobHistoryResponseBody) SetHttpStatusCode(v int32) *ListResultExportJobHistoryResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListResultExportJobHistoryResponseBody) SetItems(v []*ListResultExportJobHistoryResponseBodyItems) *ListResultExportJobHistoryResponseBody {
	s.Items = v
	return s
}

func (s *ListResultExportJobHistoryResponseBody) SetMessage(v string) *ListResultExportJobHistoryResponseBody {
	s.Message = &v
	return s
}

func (s *ListResultExportJobHistoryResponseBody) SetPageNumber(v int32) *ListResultExportJobHistoryResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListResultExportJobHistoryResponseBody) SetPageSize(v int32) *ListResultExportJobHistoryResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListResultExportJobHistoryResponseBody) SetRequestId(v string) *ListResultExportJobHistoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListResultExportJobHistoryResponseBody) SetSuccess(v bool) *ListResultExportJobHistoryResponseBody {
	s.Success = &v
	return s
}

func (s *ListResultExportJobHistoryResponseBody) SetTotalCount(v int64) *ListResultExportJobHistoryResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListResultExportJobHistoryResponseBody) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListResultExportJobHistoryResponseBodyItems struct {
	// The RAM user ID.
	//
	// example:
	//
	// 120010511678****
	AliUid *string `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	// The time when the result set export job was created. The time follows the ISO 8601 standard in the *yyyy-mm-ddThh:mm:ssZ	- format. The time is displayed in UTC.
	//
	// example:
	//
	// 2022-04-01T09:50:18Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the AnalyticDB for MySQL Data Lakehouse Edition (V3.0) cluster.
	//
	// example:
	//
	// amv-7xv5ty5m9o4v****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The name of the database account that is associated with the RAM user.
	//
	// example:
	//
	// ram_user
	DatabaseUser *string `json:"DatabaseUser,omitempty" xml:"DatabaseUser,omitempty"`
	// The end time of the result set export job. The time follows the ISO 8601 standard in the *yyyy-MM-ddTHH:mm:ssZ	- format. The time is displayed in UTC.
	//
	// >  The end time must be later than the start time.
	//
	// example:
	//
	// 2023-06-15T02:13:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The engine that is used to execute the result set export job. Only XIHE is returned.
	//
	// example:
	//
	// XIHE
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// The unique identifier of the result set export job.
	//
	// example:
	//
	// export_2024051319271219802100800401300****
	ExportJobId *string `json:"ExportJobId,omitempty" xml:"ExportJobId,omitempty"`
	// The complete URL of the path to store the exported result set.
	ExportPath *string `json:"ExportPath,omitempty" xml:"ExportPath,omitempty"`
	// The number of exported rows. This parameter is returned only when the request was successful.
	//
	// example:
	//
	// 10000
	ExportRows *string `json:"ExportRows,omitempty" xml:"ExportRows,omitempty"`
	// The type of the result set export job.
	//
	// example:
	//
	// -
	ExportType *string `json:"ExportType,omitempty" xml:"ExportType,omitempty"`
	// Indicates whether the result set export job has expired. Valid values:
	//
	// 	- **false**
	//
	// 	- **true**
	IsExpired *bool `json:"IsExpired,omitempty" xml:"IsExpired,omitempty"`
	// The returned message. This parameter is returned only when the request failed.
	//
	// example:
	//
	// Failed to execute SQL
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The query ID that can be used for diagnostics.
	//
	// >  You can call the [DescribeDiagnosisSQLInfo](https://help.aliyun.com/document_detail/612337.html) operation to query the execution information about a query.
	//
	// example:
	//
	// 202306121421111720161451770345339****
	ProcessId *string `json:"ProcessId,omitempty" xml:"ProcessId,omitempty"`
	// The progress of the result set export job. Unit: %. Valid values: 0 to 100.
	//
	// example:
	//
	// 30
	Progress *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// The name of the resource group that runs the result set export job.
	//
	// example:
	//
	// test
	ResourceGroup *string `json:"ResourceGroup,omitempty" xml:"ResourceGroup,omitempty"`
	// The name of the database.
	//
	// example:
	//
	// lake_db
	Schema *string `json:"Schema,omitempty" xml:"Schema,omitempty"`
	// The SQL statement that is used in the result set export job.
	//
	// example:
	//
	// SELECT 	- FROM `ADB_SampleData_TPCH`.`supplier` LIMIT 20
	Sql *string `json:"Sql,omitempty" xml:"Sql,omitempty"`
	// The start time of the result set export job. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2021-07-03T04:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The execution status of the result set export job. Valid values:
	//
	// 1.  **SUBMITTED**
	//
	// 2.  **RUNNING**
	//
	// 3.  **SUCCEEDED**
	//
	// 4.  **FAILED**
	//
	// example:
	//
	// RUNNING
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The amount of time consumed to export execution records. Unit: milliseconds.
	//
	// >  The value is the duration between the time when the result set export job starts and the time when the result set export job ends.
	//
	// example:
	//
	// 560
	TimeCost *int64 `json:"TimeCost,omitempty" xml:"TimeCost,omitempty"`
}

func (s ListResultExportJobHistoryResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s ListResultExportJobHistoryResponseBodyItems) GoString() string {
	return s.String()
}

func (s *ListResultExportJobHistoryResponseBodyItems) GetAliUid() *string {
	return s.AliUid
}

func (s *ListResultExportJobHistoryResponseBodyItems) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListResultExportJobHistoryResponseBodyItems) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ListResultExportJobHistoryResponseBodyItems) GetDatabaseUser() *string {
	return s.DatabaseUser
}

func (s *ListResultExportJobHistoryResponseBodyItems) GetEndTime() *string {
	return s.EndTime
}

func (s *ListResultExportJobHistoryResponseBodyItems) GetEngine() *string {
	return s.Engine
}

func (s *ListResultExportJobHistoryResponseBodyItems) GetExportJobId() *string {
	return s.ExportJobId
}

func (s *ListResultExportJobHistoryResponseBodyItems) GetExportPath() *string {
	return s.ExportPath
}

func (s *ListResultExportJobHistoryResponseBodyItems) GetExportRows() *string {
	return s.ExportRows
}

func (s *ListResultExportJobHistoryResponseBodyItems) GetExportType() *string {
	return s.ExportType
}

func (s *ListResultExportJobHistoryResponseBodyItems) GetIsExpired() *bool {
	return s.IsExpired
}

func (s *ListResultExportJobHistoryResponseBodyItems) GetMessage() *string {
	return s.Message
}

func (s *ListResultExportJobHistoryResponseBodyItems) GetProcessId() *string {
	return s.ProcessId
}

func (s *ListResultExportJobHistoryResponseBodyItems) GetProgress() *string {
	return s.Progress
}

func (s *ListResultExportJobHistoryResponseBodyItems) GetResourceGroup() *string {
	return s.ResourceGroup
}

func (s *ListResultExportJobHistoryResponseBodyItems) GetSchema() *string {
	return s.Schema
}

func (s *ListResultExportJobHistoryResponseBodyItems) GetSql() *string {
	return s.Sql
}

func (s *ListResultExportJobHistoryResponseBodyItems) GetStartTime() *string {
	return s.StartTime
}

func (s *ListResultExportJobHistoryResponseBodyItems) GetStatus() *string {
	return s.Status
}

func (s *ListResultExportJobHistoryResponseBodyItems) GetTimeCost() *int64 {
	return s.TimeCost
}

func (s *ListResultExportJobHistoryResponseBodyItems) SetAliUid(v string) *ListResultExportJobHistoryResponseBodyItems {
	s.AliUid = &v
	return s
}

func (s *ListResultExportJobHistoryResponseBodyItems) SetCreateTime(v string) *ListResultExportJobHistoryResponseBodyItems {
	s.CreateTime = &v
	return s
}

func (s *ListResultExportJobHistoryResponseBodyItems) SetDBClusterId(v string) *ListResultExportJobHistoryResponseBodyItems {
	s.DBClusterId = &v
	return s
}

func (s *ListResultExportJobHistoryResponseBodyItems) SetDatabaseUser(v string) *ListResultExportJobHistoryResponseBodyItems {
	s.DatabaseUser = &v
	return s
}

func (s *ListResultExportJobHistoryResponseBodyItems) SetEndTime(v string) *ListResultExportJobHistoryResponseBodyItems {
	s.EndTime = &v
	return s
}

func (s *ListResultExportJobHistoryResponseBodyItems) SetEngine(v string) *ListResultExportJobHistoryResponseBodyItems {
	s.Engine = &v
	return s
}

func (s *ListResultExportJobHistoryResponseBodyItems) SetExportJobId(v string) *ListResultExportJobHistoryResponseBodyItems {
	s.ExportJobId = &v
	return s
}

func (s *ListResultExportJobHistoryResponseBodyItems) SetExportPath(v string) *ListResultExportJobHistoryResponseBodyItems {
	s.ExportPath = &v
	return s
}

func (s *ListResultExportJobHistoryResponseBodyItems) SetExportRows(v string) *ListResultExportJobHistoryResponseBodyItems {
	s.ExportRows = &v
	return s
}

func (s *ListResultExportJobHistoryResponseBodyItems) SetExportType(v string) *ListResultExportJobHistoryResponseBodyItems {
	s.ExportType = &v
	return s
}

func (s *ListResultExportJobHistoryResponseBodyItems) SetIsExpired(v bool) *ListResultExportJobHistoryResponseBodyItems {
	s.IsExpired = &v
	return s
}

func (s *ListResultExportJobHistoryResponseBodyItems) SetMessage(v string) *ListResultExportJobHistoryResponseBodyItems {
	s.Message = &v
	return s
}

func (s *ListResultExportJobHistoryResponseBodyItems) SetProcessId(v string) *ListResultExportJobHistoryResponseBodyItems {
	s.ProcessId = &v
	return s
}

func (s *ListResultExportJobHistoryResponseBodyItems) SetProgress(v string) *ListResultExportJobHistoryResponseBodyItems {
	s.Progress = &v
	return s
}

func (s *ListResultExportJobHistoryResponseBodyItems) SetResourceGroup(v string) *ListResultExportJobHistoryResponseBodyItems {
	s.ResourceGroup = &v
	return s
}

func (s *ListResultExportJobHistoryResponseBodyItems) SetSchema(v string) *ListResultExportJobHistoryResponseBodyItems {
	s.Schema = &v
	return s
}

func (s *ListResultExportJobHistoryResponseBodyItems) SetSql(v string) *ListResultExportJobHistoryResponseBodyItems {
	s.Sql = &v
	return s
}

func (s *ListResultExportJobHistoryResponseBodyItems) SetStartTime(v string) *ListResultExportJobHistoryResponseBodyItems {
	s.StartTime = &v
	return s
}

func (s *ListResultExportJobHistoryResponseBodyItems) SetStatus(v string) *ListResultExportJobHistoryResponseBodyItems {
	s.Status = &v
	return s
}

func (s *ListResultExportJobHistoryResponseBodyItems) SetTimeCost(v int64) *ListResultExportJobHistoryResponseBodyItems {
	s.TimeCost = &v
	return s
}

func (s *ListResultExportJobHistoryResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
