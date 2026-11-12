// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeStreamingJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccount(v string) *DescribeStreamingJobResponseBody
	GetAccount() *string
	SetConsistency(v string) *DescribeStreamingJobResponseBody
	GetConsistency() *string
	SetCreateTime(v string) *DescribeStreamingJobResponseBody
	GetCreateTime() *string
	SetDataSourceId(v string) *DescribeStreamingJobResponseBody
	GetDataSourceId() *string
	SetDataSourceName(v string) *DescribeStreamingJobResponseBody
	GetDataSourceName() *string
	SetDestColumns(v []*string) *DescribeStreamingJobResponseBody
	GetDestColumns() []*string
	SetDestDatabase(v string) *DescribeStreamingJobResponseBody
	GetDestDatabase() *string
	SetDestSchema(v string) *DescribeStreamingJobResponseBody
	GetDestSchema() *string
	SetDestTable(v string) *DescribeStreamingJobResponseBody
	GetDestTable() *string
	SetErrorLimitCount(v int32) *DescribeStreamingJobResponseBody
	GetErrorLimitCount() *int32
	SetErrorMessage(v string) *DescribeStreamingJobResponseBody
	GetErrorMessage() *string
	SetFallbackOffset(v string) *DescribeStreamingJobResponseBody
	GetFallbackOffset() *string
	SetGroupName(v string) *DescribeStreamingJobResponseBody
	GetGroupName() *string
	SetJobConfig(v string) *DescribeStreamingJobResponseBody
	GetJobConfig() *string
	SetJobDescription(v string) *DescribeStreamingJobResponseBody
	GetJobDescription() *string
	SetJobId(v string) *DescribeStreamingJobResponseBody
	GetJobId() *string
	SetJobName(v string) *DescribeStreamingJobResponseBody
	GetJobName() *string
	SetMatchColumns(v []*string) *DescribeStreamingJobResponseBody
	GetMatchColumns() []*string
	SetMode(v string) *DescribeStreamingJobResponseBody
	GetMode() *string
	SetModifyTime(v string) *DescribeStreamingJobResponseBody
	GetModifyTime() *string
	SetPassword(v string) *DescribeStreamingJobResponseBody
	GetPassword() *string
	SetRequestId(v string) *DescribeStreamingJobResponseBody
	GetRequestId() *string
	SetSrcColumns(v []*string) *DescribeStreamingJobResponseBody
	GetSrcColumns() []*string
	SetStatus(v string) *DescribeStreamingJobResponseBody
	GetStatus() *string
	SetUpdateColumns(v []*string) *DescribeStreamingJobResponseBody
	GetUpdateColumns() []*string
	SetWriteMode(v string) *DescribeStreamingJobResponseBody
	GetWriteMode() *string
}

type DescribeStreamingJobResponseBody struct {
	// The destination database account.
	//
	// example:
	//
	// test-account
	Account *string `json:"Account,omitempty" xml:"Account,omitempty"`
	// The delivery guarantee.
	//
	// example:
	//
	// ATLEAST / EXACTLY
	Consistency *string `json:"Consistency,omitempty" xml:"Consistency,omitempty"`
	// The creation time.
	//
	// Use the UTC time format: yyyy-MM-ddTHH:mm:ssZ
	//
	// example:
	//
	// 2019-09-08T16:00:00Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The data source ID.
	//
	// example:
	//
	// 2
	DataSourceId *string `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
	// The data source name.
	//
	// example:
	//
	// test_kafka
	DataSourceName *string `json:"DataSourceName,omitempty" xml:"DataSourceName,omitempty"`
	// The list of mapped columns in the destination table.
	DestColumns []*string `json:"DestColumns,omitempty" xml:"DestColumns,omitempty" type:"Repeated"`
	// The destination database name.
	//
	// example:
	//
	// dest-db
	DestDatabase *string `json:"DestDatabase,omitempty" xml:"DestDatabase,omitempty"`
	// The destination namespace.
	//
	// example:
	//
	// dest-schema
	DestSchema *string `json:"DestSchema,omitempty" xml:"DestSchema,omitempty"`
	// The destination table name.
	//
	// example:
	//
	// dest-table
	DestTable *string `json:"DestTable,omitempty" xml:"DestTable,omitempty"`
	// The maximum number of allowed error rows. If data in Kafka does not match the destination AnalyticDB for PostgreSQL table, a write failure occurs. If the number of error rows exceeds this value, the job fails.
	//
	// example:
	//
	// 5
	ErrorLimitCount *int32 `json:"ErrorLimitCount,omitempty" xml:"ErrorLimitCount,omitempty"`
	// The service status message. For example, if an exception occurs, the cause of the exception is displayed. This parameter is empty when the service is in the Running state.
	//
	// example:
	//
	// ""
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The fallback offset.
	//
	// - The FallbackOffset parameter defines the behavior when a consumer does not request a specific offset to consume, or when the requested offset exceeds the offset range recorded in the current Kafka cluster. You can choose to start consuming from the earliest or latest offset.
	//
	// example:
	//
	// EARLIEST /  LATEST
	FallbackOffset *string `json:"FallbackOffset,omitempty" xml:"FallbackOffset,omitempty"`
	// The Kafka group name.
	//
	// example:
	//
	// test_group
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The job configuration file.
	//
	// example:
	//
	// DATABASE: adbpgss_test
	//
	// USER: adbpgss_test
	//
	// PASSWORD: adbpgssTest
	//
	// HOST: gp-xxx-master.gpdb.rds-aliyun-pre.rds.aliyuncs.com
	//
	// PORT: 5432
	//
	// KAFKA:
	//
	//   INPUT:
	//
	//     SOURCE:
	//
	//       BROKERS: broker1:9092,broker2:9092,broker3:9092
	//
	//       TOPIC: testtopic
	//
	//       FALLBACK_OFFSET: earliest
	//
	//     KEY:
	//
	//       COLUMNS:
	//
	//       - NAME: customer_id
	//
	//         TYPE: int
	//
	//       FORMAT: delimited
	//
	//       DELIMITED_OPTION:
	//
	//         DELIMITER: \\"|\\"
	//
	//     VALUE:
	//
	//       COLUMNS:
	//
	//       - TYPE: integer
	//
	//         NAME: l_orderkey
	//
	//       - TYPE: integer
	//
	//         NAME: l_partkey
	//
	//       - TYPE: integer
	//
	//         NAME: l_suppkey
	//
	//       - TYPE: integer
	//
	//         NAME: l_linenumber
	//
	//       - TYPE: decimal
	//
	//         NAME: l_quantity
	//
	//       - TYPE: decimal
	//
	//         NAME: l_extendedprice
	//
	//       - TYPE: decimal
	//
	//         NAME: l_discount
	//
	//       - TYPE: decimal
	//
	//         NAME: l_tax
	//
	//       - TYPE: char
	//
	//         NAME: l_returnflag
	//
	//       - TYPE: char
	//
	//         NAME: l_linestatus
	//
	//       - TYPE: date
	//
	//         NAME: l_shipdate
	//
	//       - TYPE: date
	//
	//         NAME: l_commitdate
	//
	//       - TYPE: date
	//
	//         NAME: l_receiptdate
	//
	//       - TYPE: text
	//
	//         NAME: l_shipinstruct
	//
	//       - TYPE: text
	//
	//         NAME: l_shipmode
	//
	//       - TYPE: text
	//
	//         NAME: l_comment
	//
	//       FORMAT: delimited
	//
	//       DELIMITED_OPTION:
	//
	//         DELIMITER: \\"|\\"
	//
	//     ERROR_LIMIT: 10
	//
	//   OUTPUT:
	//
	//     SCHEMA: adbpgss_test
	//
	//     TABLE: write_with_insert_plaintext
	//
	//     MODE: MERGE
	//
	//     MATCH_COLUMNS:
	//
	//     - l_orderkey
	//
	//     - l_partkey
	//
	//     - l_suppkey
	//
	//     UPDATE_COLUMNS:
	//
	//     - l_linenumber
	//
	//     - l_quantity
	//
	//     - l_extendedprice
	//
	//     - l_discount
	//
	//     - l_tax
	//
	//     - l_returnflag
	//
	//     - l_linestatus
	//
	//     - l_shipdate
	//
	//     - l_commitdate
	//
	//     - l_receiptdate
	//
	//     - l_shipinstruct
	//
	//     - l_shipmode
	//
	//     - l_comment
	//
	//     MAPPING:
	//
	//     - EXPRESSION: l_orderkey
	//
	//       NAME: l_orderkey
	//
	//     - EXPRESSION: l_partkey
	//
	//       NAME: l_partkey
	//
	//     - EXPRESSION: l_suppkey
	//
	//       NAME: l_suppkey
	//
	//     - EXPRESSION: l_linenumber
	//
	//       NAME: l_linenumber
	//
	//     - EXPRESSION: l_quantity
	//
	//       NAME: l_quantity
	//
	//     - EXPRESSION: l_extendedprice
	//
	//       NAME: l_extendedprice
	//
	//     - EXPRESSION: l_discount
	//
	//       NAME: l_discount
	//
	//     - EXPRESSION: l_tax
	//
	//       NAME: l_tax
	//
	//     - EXPRESSION: l_returnflag
	//
	//       NAME: l_returnflag
	//
	//     - EXPRESSION: l_linestatus
	//
	//       NAME: l_linestatus
	//
	//     - EXPRESSION: l_shipdate
	//
	//       NAME: l_shipdate
	//
	//     - EXPRESSION: l_commitdate
	//
	//       NAME: l_commitdate
	//
	//     - EXPRESSION: l_receiptdate
	//
	//       NAME: l_receiptdate
	//
	//     - EXPRESSION: l_shipinstruct
	//
	//       NAME: l_shipinstruct
	//
	//     - EXPRESSION: l_shipmode
	//
	//       NAME: l_shipmode
	//
	//     - EXPRESSION: l_comment
	//
	//       NAME: l_comment
	//
	//   COMMIT:
	//
	//     MAX_ROW: 1000
	//
	//     MINIMAL_INTERVAL: 1000
	//
	//     CONSISTENCY: ATLEAST
	//
	//   POLL:
	//
	//     BATCHSIZE: 1000
	//
	//     TIMEOUT: 1000
	//
	//   PROPERTIES:
	//
	//     group.id: testgroup
	JobConfig *string `json:"JobConfig,omitempty" xml:"JobConfig,omitempty"`
	// The job description.
	//
	// example:
	//
	// test_job
	JobDescription *string `json:"JobDescription,omitempty" xml:"JobDescription,omitempty"`
	// The job ID.
	//
	// example:
	//
	// 1
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The job name.
	//
	// example:
	//
	// test-job
	JobName *string `json:"JobName,omitempty" xml:"JobName,omitempty"`
	// The match columns. Typically, these are all primary key columns of the destination table. If the values of all columns in this configuration are the same, the two rows of data are considered duplicate.
	MatchColumns []*string `json:"MatchColumns,omitempty" xml:"MatchColumns,omitempty" type:"Repeated"`
	// The configuration mode.
	//
	// 1. basic: requires specific configuration fields.
	//
	// 2. professional: supports submitting YAML files.
	//
	// example:
	//
	// basic/professional
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// The last modification time.
	//
	// Use the UTC time format: yyyy-MM-ddTHH:mm:ssZ
	//
	// example:
	//
	// 2019-09-08T17:00:00Z
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The destination database password.
	//
	// example:
	//
	// pwd123
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// The request ID.
	//
	// example:
	//
	// B4CAF581-2AC7-41AD-8940-D56DF7AADF5B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of source columns.
	SrcColumns []*string `json:"SrcColumns,omitempty" xml:"SrcColumns,omitempty" type:"Repeated"`
	// The service status. Valid values:
	//
	// - Init: initializing.
	//
	// - Running: running.
	//
	// - Exception: an exception occurred.
	//
	// - Paused: paused.
	//
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The update columns. Typically, these are all non-primary key columns of the destination table. When duplicate data is identified based on the match columns, the values of the update columns are updated so that the new data overwrites the old data.
	UpdateColumns []*string `json:"UpdateColumns,omitempty" xml:"UpdateColumns,omitempty" type:"Repeated"`
	// The write mode.
	//
	// example:
	//
	// INSERT/UPDATE/MERGE
	WriteMode *string `json:"WriteMode,omitempty" xml:"WriteMode,omitempty"`
}

func (s DescribeStreamingJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeStreamingJobResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeStreamingJobResponseBody) GetAccount() *string {
	return s.Account
}

func (s *DescribeStreamingJobResponseBody) GetConsistency() *string {
	return s.Consistency
}

func (s *DescribeStreamingJobResponseBody) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeStreamingJobResponseBody) GetDataSourceId() *string {
	return s.DataSourceId
}

func (s *DescribeStreamingJobResponseBody) GetDataSourceName() *string {
	return s.DataSourceName
}

func (s *DescribeStreamingJobResponseBody) GetDestColumns() []*string {
	return s.DestColumns
}

func (s *DescribeStreamingJobResponseBody) GetDestDatabase() *string {
	return s.DestDatabase
}

func (s *DescribeStreamingJobResponseBody) GetDestSchema() *string {
	return s.DestSchema
}

func (s *DescribeStreamingJobResponseBody) GetDestTable() *string {
	return s.DestTable
}

func (s *DescribeStreamingJobResponseBody) GetErrorLimitCount() *int32 {
	return s.ErrorLimitCount
}

func (s *DescribeStreamingJobResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DescribeStreamingJobResponseBody) GetFallbackOffset() *string {
	return s.FallbackOffset
}

func (s *DescribeStreamingJobResponseBody) GetGroupName() *string {
	return s.GroupName
}

func (s *DescribeStreamingJobResponseBody) GetJobConfig() *string {
	return s.JobConfig
}

func (s *DescribeStreamingJobResponseBody) GetJobDescription() *string {
	return s.JobDescription
}

func (s *DescribeStreamingJobResponseBody) GetJobId() *string {
	return s.JobId
}

func (s *DescribeStreamingJobResponseBody) GetJobName() *string {
	return s.JobName
}

func (s *DescribeStreamingJobResponseBody) GetMatchColumns() []*string {
	return s.MatchColumns
}

func (s *DescribeStreamingJobResponseBody) GetMode() *string {
	return s.Mode
}

func (s *DescribeStreamingJobResponseBody) GetModifyTime() *string {
	return s.ModifyTime
}

func (s *DescribeStreamingJobResponseBody) GetPassword() *string {
	return s.Password
}

func (s *DescribeStreamingJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeStreamingJobResponseBody) GetSrcColumns() []*string {
	return s.SrcColumns
}

func (s *DescribeStreamingJobResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DescribeStreamingJobResponseBody) GetUpdateColumns() []*string {
	return s.UpdateColumns
}

func (s *DescribeStreamingJobResponseBody) GetWriteMode() *string {
	return s.WriteMode
}

func (s *DescribeStreamingJobResponseBody) SetAccount(v string) *DescribeStreamingJobResponseBody {
	s.Account = &v
	return s
}

func (s *DescribeStreamingJobResponseBody) SetConsistency(v string) *DescribeStreamingJobResponseBody {
	s.Consistency = &v
	return s
}

func (s *DescribeStreamingJobResponseBody) SetCreateTime(v string) *DescribeStreamingJobResponseBody {
	s.CreateTime = &v
	return s
}

func (s *DescribeStreamingJobResponseBody) SetDataSourceId(v string) *DescribeStreamingJobResponseBody {
	s.DataSourceId = &v
	return s
}

func (s *DescribeStreamingJobResponseBody) SetDataSourceName(v string) *DescribeStreamingJobResponseBody {
	s.DataSourceName = &v
	return s
}

func (s *DescribeStreamingJobResponseBody) SetDestColumns(v []*string) *DescribeStreamingJobResponseBody {
	s.DestColumns = v
	return s
}

func (s *DescribeStreamingJobResponseBody) SetDestDatabase(v string) *DescribeStreamingJobResponseBody {
	s.DestDatabase = &v
	return s
}

func (s *DescribeStreamingJobResponseBody) SetDestSchema(v string) *DescribeStreamingJobResponseBody {
	s.DestSchema = &v
	return s
}

func (s *DescribeStreamingJobResponseBody) SetDestTable(v string) *DescribeStreamingJobResponseBody {
	s.DestTable = &v
	return s
}

func (s *DescribeStreamingJobResponseBody) SetErrorLimitCount(v int32) *DescribeStreamingJobResponseBody {
	s.ErrorLimitCount = &v
	return s
}

func (s *DescribeStreamingJobResponseBody) SetErrorMessage(v string) *DescribeStreamingJobResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeStreamingJobResponseBody) SetFallbackOffset(v string) *DescribeStreamingJobResponseBody {
	s.FallbackOffset = &v
	return s
}

func (s *DescribeStreamingJobResponseBody) SetGroupName(v string) *DescribeStreamingJobResponseBody {
	s.GroupName = &v
	return s
}

func (s *DescribeStreamingJobResponseBody) SetJobConfig(v string) *DescribeStreamingJobResponseBody {
	s.JobConfig = &v
	return s
}

func (s *DescribeStreamingJobResponseBody) SetJobDescription(v string) *DescribeStreamingJobResponseBody {
	s.JobDescription = &v
	return s
}

func (s *DescribeStreamingJobResponseBody) SetJobId(v string) *DescribeStreamingJobResponseBody {
	s.JobId = &v
	return s
}

func (s *DescribeStreamingJobResponseBody) SetJobName(v string) *DescribeStreamingJobResponseBody {
	s.JobName = &v
	return s
}

func (s *DescribeStreamingJobResponseBody) SetMatchColumns(v []*string) *DescribeStreamingJobResponseBody {
	s.MatchColumns = v
	return s
}

func (s *DescribeStreamingJobResponseBody) SetMode(v string) *DescribeStreamingJobResponseBody {
	s.Mode = &v
	return s
}

func (s *DescribeStreamingJobResponseBody) SetModifyTime(v string) *DescribeStreamingJobResponseBody {
	s.ModifyTime = &v
	return s
}

func (s *DescribeStreamingJobResponseBody) SetPassword(v string) *DescribeStreamingJobResponseBody {
	s.Password = &v
	return s
}

func (s *DescribeStreamingJobResponseBody) SetRequestId(v string) *DescribeStreamingJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeStreamingJobResponseBody) SetSrcColumns(v []*string) *DescribeStreamingJobResponseBody {
	s.SrcColumns = v
	return s
}

func (s *DescribeStreamingJobResponseBody) SetStatus(v string) *DescribeStreamingJobResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeStreamingJobResponseBody) SetUpdateColumns(v []*string) *DescribeStreamingJobResponseBody {
	s.UpdateColumns = v
	return s
}

func (s *DescribeStreamingJobResponseBody) SetWriteMode(v string) *DescribeStreamingJobResponseBody {
	s.WriteMode = &v
	return s
}

func (s *DescribeStreamingJobResponseBody) Validate() error {
	return dara.Validate(s)
}
