// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListStreamingJobsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJobItems(v []*ListStreamingJobsResponseBodyJobItems) *ListStreamingJobsResponseBody
	GetJobItems() []*ListStreamingJobsResponseBodyJobItems
	SetPageNumber(v int32) *ListStreamingJobsResponseBody
	GetPageNumber() *int32
	SetPageRecordCount(v int32) *ListStreamingJobsResponseBody
	GetPageRecordCount() *int32
	SetRequestId(v string) *ListStreamingJobsResponseBody
	GetRequestId() *string
	SetTotalRecordCount(v int32) *ListStreamingJobsResponseBody
	GetTotalRecordCount() *int32
}

type ListStreamingJobsResponseBody struct {
	// The queried jobs.
	JobItems []*ListStreamingJobsResponseBodyJobItems `json:"JobItems,omitempty" xml:"JobItems,omitempty" type:"Repeated"`
	// Current page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// Number of records per page.
	//
	// example:
	//
	// 2
	PageRecordCount *int32 `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	// Request ID.
	//
	// example:
	//
	// B4CAF581-2AC7-41AD-8940-D56DF7AADF5B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Total number of records.
	//
	// example:
	//
	// 2
	TotalRecordCount *int32 `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
}

func (s ListStreamingJobsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListStreamingJobsResponseBody) GoString() string {
	return s.String()
}

func (s *ListStreamingJobsResponseBody) GetJobItems() []*ListStreamingJobsResponseBodyJobItems {
	return s.JobItems
}

func (s *ListStreamingJobsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListStreamingJobsResponseBody) GetPageRecordCount() *int32 {
	return s.PageRecordCount
}

func (s *ListStreamingJobsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListStreamingJobsResponseBody) GetTotalRecordCount() *int32 {
	return s.TotalRecordCount
}

func (s *ListStreamingJobsResponseBody) SetJobItems(v []*ListStreamingJobsResponseBodyJobItems) *ListStreamingJobsResponseBody {
	s.JobItems = v
	return s
}

func (s *ListStreamingJobsResponseBody) SetPageNumber(v int32) *ListStreamingJobsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListStreamingJobsResponseBody) SetPageRecordCount(v int32) *ListStreamingJobsResponseBody {
	s.PageRecordCount = &v
	return s
}

func (s *ListStreamingJobsResponseBody) SetRequestId(v string) *ListStreamingJobsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListStreamingJobsResponseBody) SetTotalRecordCount(v int32) *ListStreamingJobsResponseBody {
	s.TotalRecordCount = &v
	return s
}

func (s *ListStreamingJobsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListStreamingJobsResponseBodyJobItems struct {
	// The name of the database account.
	//
	// example:
	//
	// test-account
	Account *string `json:"Account,omitempty" xml:"Account,omitempty"`
	// The delivery guarantee setting.
	//
	// example:
	//
	// ATLEAST / EXACTLY
	Consistency *string `json:"Consistency,omitempty" xml:"Consistency,omitempty"`
	// The time when the job was created.
	//
	// example:
	//
	// 2019-09-08T16:00:00Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The data source ID.
	//
	// example:
	//
	// 58
	DataSourceId *string `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
	// The name of the data source.
	//
	// example:
	//
	// test-kafka
	DataSourceName *string `json:"DataSourceName,omitempty" xml:"DataSourceName,omitempty"`
	// The mapped fields in the destination table.
	DestColumns []*string `json:"DestColumns,omitempty" xml:"DestColumns,omitempty" type:"Repeated"`
	// The name of the destination database.
	//
	// example:
	//
	// dest-db
	DestDatabase *string `json:"DestDatabase,omitempty" xml:"DestDatabase,omitempty"`
	// The name of the destination namespace.
	//
	// example:
	//
	// dest-schema
	DestSchema *string `json:"DestSchema,omitempty" xml:"DestSchema,omitempty"`
	// The name of the destination table.
	//
	// example:
	//
	// dest-table
	DestTable *string `json:"DestTable,omitempty" xml:"DestTable,omitempty"`
	// The error message returned.
	//
	// This parameter is returned only when the return value of **Status*	- is **false**.
	//
	// example:
	//
	// ""
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The fallback offset for data consumption.
	//
	// 	- This parameter indicates the starting offset from which data consumption resumes when a consumer does not request a consumption offset or requests a consumption offset that is beyond the range of the offset information recorded in the current Kafka cluster. Valid values: EARLIEST and LATEST.
	//
	// example:
	//
	// EARLIEST /  LATEST
	FallbackOffset *string `json:"FallbackOffset,omitempty" xml:"FallbackOffset,omitempty"`
	// The description of the job.
	//
	// example:
	//
	// test job
	JobDescription *string `json:"JobDescription,omitempty" xml:"JobDescription,omitempty"`
	// The job ID.
	//
	// example:
	//
	// 1
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The name of the job.
	//
	// example:
	//
	// test-job
	JobName *string `json:"JobName,omitempty" xml:"JobName,omitempty"`
	// The update condition columns that are used to join the source data and the destination table. Typically, the columns are all the primary key columns of the destination table. If the values of all columns specified by this parameter in different rows are the same, the rows are considered duplicates.
	MatchColumns []*string `json:"MatchColumns,omitempty" xml:"MatchColumns,omitempty" type:"Repeated"`
	// The configuration mode. Valid values:
	//
	// 1.  basic: In basic mode, you must configure the configuration parameters.
	//
	// 2.  professional: In professional mode, you can submit a YAML configuration file.
	//
	// example:
	//
	// Basic / Professional
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// The time when the job was last modified.
	//
	// example:
	//
	// 2019-09-08T17:00:00Z
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The password of the database account.
	//
	// example:
	//
	// pwd123
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// The source fields.
	SrcColumns []*string `json:"SrcColumns,omitempty" xml:"SrcColumns,omitempty" type:"Repeated"`
	// The status of the job. Valid values:
	//
	// 	- Init
	//
	// 	- Running
	//
	// 	- Exception
	//
	// 	- Paused
	//
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The columns to be updated if a row of data meets the update condition. Typically, the columns are all non-primary key columns of the destination table. When the columns specified by the MatchColumns parameter are used as conditions to join the source data and the destination table, data in columns of the UpdateColumns type is updated if data is matched.
	UpdateColumns []*string `json:"UpdateColumns,omitempty" xml:"UpdateColumns,omitempty" type:"Repeated"`
	// The write mode.
	//
	// example:
	//
	// INSERT/UPDATE/MERGE
	WriteMode *string `json:"WriteMode,omitempty" xml:"WriteMode,omitempty"`
}

func (s ListStreamingJobsResponseBodyJobItems) String() string {
	return dara.Prettify(s)
}

func (s ListStreamingJobsResponseBodyJobItems) GoString() string {
	return s.String()
}

func (s *ListStreamingJobsResponseBodyJobItems) GetAccount() *string {
	return s.Account
}

func (s *ListStreamingJobsResponseBodyJobItems) GetConsistency() *string {
	return s.Consistency
}

func (s *ListStreamingJobsResponseBodyJobItems) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListStreamingJobsResponseBodyJobItems) GetDataSourceId() *string {
	return s.DataSourceId
}

func (s *ListStreamingJobsResponseBodyJobItems) GetDataSourceName() *string {
	return s.DataSourceName
}

func (s *ListStreamingJobsResponseBodyJobItems) GetDestColumns() []*string {
	return s.DestColumns
}

func (s *ListStreamingJobsResponseBodyJobItems) GetDestDatabase() *string {
	return s.DestDatabase
}

func (s *ListStreamingJobsResponseBodyJobItems) GetDestSchema() *string {
	return s.DestSchema
}

func (s *ListStreamingJobsResponseBodyJobItems) GetDestTable() *string {
	return s.DestTable
}

func (s *ListStreamingJobsResponseBodyJobItems) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListStreamingJobsResponseBodyJobItems) GetFallbackOffset() *string {
	return s.FallbackOffset
}

func (s *ListStreamingJobsResponseBodyJobItems) GetJobDescription() *string {
	return s.JobDescription
}

func (s *ListStreamingJobsResponseBodyJobItems) GetJobId() *string {
	return s.JobId
}

func (s *ListStreamingJobsResponseBodyJobItems) GetJobName() *string {
	return s.JobName
}

func (s *ListStreamingJobsResponseBodyJobItems) GetMatchColumns() []*string {
	return s.MatchColumns
}

func (s *ListStreamingJobsResponseBodyJobItems) GetMode() *string {
	return s.Mode
}

func (s *ListStreamingJobsResponseBodyJobItems) GetModifyTime() *string {
	return s.ModifyTime
}

func (s *ListStreamingJobsResponseBodyJobItems) GetPassword() *string {
	return s.Password
}

func (s *ListStreamingJobsResponseBodyJobItems) GetSrcColumns() []*string {
	return s.SrcColumns
}

func (s *ListStreamingJobsResponseBodyJobItems) GetStatus() *string {
	return s.Status
}

func (s *ListStreamingJobsResponseBodyJobItems) GetUpdateColumns() []*string {
	return s.UpdateColumns
}

func (s *ListStreamingJobsResponseBodyJobItems) GetWriteMode() *string {
	return s.WriteMode
}

func (s *ListStreamingJobsResponseBodyJobItems) SetAccount(v string) *ListStreamingJobsResponseBodyJobItems {
	s.Account = &v
	return s
}

func (s *ListStreamingJobsResponseBodyJobItems) SetConsistency(v string) *ListStreamingJobsResponseBodyJobItems {
	s.Consistency = &v
	return s
}

func (s *ListStreamingJobsResponseBodyJobItems) SetCreateTime(v string) *ListStreamingJobsResponseBodyJobItems {
	s.CreateTime = &v
	return s
}

func (s *ListStreamingJobsResponseBodyJobItems) SetDataSourceId(v string) *ListStreamingJobsResponseBodyJobItems {
	s.DataSourceId = &v
	return s
}

func (s *ListStreamingJobsResponseBodyJobItems) SetDataSourceName(v string) *ListStreamingJobsResponseBodyJobItems {
	s.DataSourceName = &v
	return s
}

func (s *ListStreamingJobsResponseBodyJobItems) SetDestColumns(v []*string) *ListStreamingJobsResponseBodyJobItems {
	s.DestColumns = v
	return s
}

func (s *ListStreamingJobsResponseBodyJobItems) SetDestDatabase(v string) *ListStreamingJobsResponseBodyJobItems {
	s.DestDatabase = &v
	return s
}

func (s *ListStreamingJobsResponseBodyJobItems) SetDestSchema(v string) *ListStreamingJobsResponseBodyJobItems {
	s.DestSchema = &v
	return s
}

func (s *ListStreamingJobsResponseBodyJobItems) SetDestTable(v string) *ListStreamingJobsResponseBodyJobItems {
	s.DestTable = &v
	return s
}

func (s *ListStreamingJobsResponseBodyJobItems) SetErrorMessage(v string) *ListStreamingJobsResponseBodyJobItems {
	s.ErrorMessage = &v
	return s
}

func (s *ListStreamingJobsResponseBodyJobItems) SetFallbackOffset(v string) *ListStreamingJobsResponseBodyJobItems {
	s.FallbackOffset = &v
	return s
}

func (s *ListStreamingJobsResponseBodyJobItems) SetJobDescription(v string) *ListStreamingJobsResponseBodyJobItems {
	s.JobDescription = &v
	return s
}

func (s *ListStreamingJobsResponseBodyJobItems) SetJobId(v string) *ListStreamingJobsResponseBodyJobItems {
	s.JobId = &v
	return s
}

func (s *ListStreamingJobsResponseBodyJobItems) SetJobName(v string) *ListStreamingJobsResponseBodyJobItems {
	s.JobName = &v
	return s
}

func (s *ListStreamingJobsResponseBodyJobItems) SetMatchColumns(v []*string) *ListStreamingJobsResponseBodyJobItems {
	s.MatchColumns = v
	return s
}

func (s *ListStreamingJobsResponseBodyJobItems) SetMode(v string) *ListStreamingJobsResponseBodyJobItems {
	s.Mode = &v
	return s
}

func (s *ListStreamingJobsResponseBodyJobItems) SetModifyTime(v string) *ListStreamingJobsResponseBodyJobItems {
	s.ModifyTime = &v
	return s
}

func (s *ListStreamingJobsResponseBodyJobItems) SetPassword(v string) *ListStreamingJobsResponseBodyJobItems {
	s.Password = &v
	return s
}

func (s *ListStreamingJobsResponseBodyJobItems) SetSrcColumns(v []*string) *ListStreamingJobsResponseBodyJobItems {
	s.SrcColumns = v
	return s
}

func (s *ListStreamingJobsResponseBodyJobItems) SetStatus(v string) *ListStreamingJobsResponseBodyJobItems {
	s.Status = &v
	return s
}

func (s *ListStreamingJobsResponseBodyJobItems) SetUpdateColumns(v []*string) *ListStreamingJobsResponseBodyJobItems {
	s.UpdateColumns = v
	return s
}

func (s *ListStreamingJobsResponseBodyJobItems) SetWriteMode(v string) *ListStreamingJobsResponseBodyJobItems {
	s.WriteMode = &v
	return s
}

func (s *ListStreamingJobsResponseBodyJobItems) Validate() error {
	return dara.Validate(s)
}
