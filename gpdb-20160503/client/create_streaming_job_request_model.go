// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateStreamingJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccount(v string) *CreateStreamingJobRequest
	GetAccount() *string
	SetConsistency(v string) *CreateStreamingJobRequest
	GetConsistency() *string
	SetDBInstanceId(v string) *CreateStreamingJobRequest
	GetDBInstanceId() *string
	SetDataSourceId(v string) *CreateStreamingJobRequest
	GetDataSourceId() *string
	SetDestColumns(v []*string) *CreateStreamingJobRequest
	GetDestColumns() []*string
	SetDestDatabase(v string) *CreateStreamingJobRequest
	GetDestDatabase() *string
	SetDestSchema(v string) *CreateStreamingJobRequest
	GetDestSchema() *string
	SetDestTable(v string) *CreateStreamingJobRequest
	GetDestTable() *string
	SetErrorLimitCount(v int64) *CreateStreamingJobRequest
	GetErrorLimitCount() *int64
	SetFallbackOffset(v string) *CreateStreamingJobRequest
	GetFallbackOffset() *string
	SetGroupName(v string) *CreateStreamingJobRequest
	GetGroupName() *string
	SetJobConfig(v string) *CreateStreamingJobRequest
	GetJobConfig() *string
	SetJobDescription(v string) *CreateStreamingJobRequest
	GetJobDescription() *string
	SetJobName(v string) *CreateStreamingJobRequest
	GetJobName() *string
	SetMatchColumns(v []*string) *CreateStreamingJobRequest
	GetMatchColumns() []*string
	SetMode(v string) *CreateStreamingJobRequest
	GetMode() *string
	SetPassword(v string) *CreateStreamingJobRequest
	GetPassword() *string
	SetRegionId(v string) *CreateStreamingJobRequest
	GetRegionId() *string
	SetSrcColumns(v []*string) *CreateStreamingJobRequest
	GetSrcColumns() []*string
	SetTryRun(v bool) *CreateStreamingJobRequest
	GetTryRun() *bool
	SetUpdateColumns(v []*string) *CreateStreamingJobRequest
	GetUpdateColumns() []*string
	SetWriteMode(v string) *CreateStreamingJobRequest
	GetWriteMode() *string
}

type CreateStreamingJobRequest struct {
	// The account for the destination database.
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
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// gp-bp10g78o9807yv9h3
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The data source ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	DataSourceId *string `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
	// The list of mapped fields in the destination table.
	DestColumns []*string `json:"DestColumns,omitempty" xml:"DestColumns,omitempty" type:"Repeated"`
	// The name of the destination database.
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
	// The name of the destination table.
	//
	// example:
	//
	// dest-table
	DestTable *string `json:"DestTable,omitempty" xml:"DestTable,omitempty"`
	// The maximum number of allowed error rows. When data in Kafka does not match the destination AnalyticDB for PostgreSQL table, a write failure occurs. If the number of error rows exceeds this value, the task fails.
	//
	// example:
	//
	// 5
	ErrorLimitCount *int64 `json:"ErrorLimitCount,omitempty" xml:"ErrorLimitCount,omitempty"`
	// The fallback offset.
	//
	// - The FallbackOffset parameter defines the behavior when a consumer does not request a specific offset to consume, or when the requested offset exceeds the offset range recorded in the current Kafka cluster. You can choose to start consuming from the earliest or latest offset.
	//
	// example:
	//
	// EARLIEST /  LATEST
	FallbackOffset *string `json:"FallbackOffset,omitempty" xml:"FallbackOffset,omitempty"`
	// The name of the Kafka group.
	//
	// example:
	//
	// group_name。
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The task configuration file. This parameter is required in professional mode.
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
	// The task description.
	//
	// example:
	//
	// test-job
	JobDescription *string `json:"JobDescription,omitempty" xml:"JobDescription,omitempty"`
	// The task name.
	//
	// This parameter is required.
	//
	// example:
	//
	// test-job
	JobName *string `json:"JobName,omitempty" xml:"JobName,omitempty"`
	// The match columns. Typically, these are all the primary key columns of the destination table. If the values of all columns in this configuration are the same, the two rows of data are considered duplicate.
	MatchColumns []*string `json:"MatchColumns,omitempty" xml:"MatchColumns,omitempty" type:"Repeated"`
	// The configuration mode.
	//
	// 1. The basic mode requires you to specify certain configuration fields.
	//
	// 2. The professional mode supports submitting a YAML file.
	//
	// example:
	//
	// basic / professional
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// The password for the destination database.
	//
	// example:
	//
	// pwd123
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// The region ID.
	//
	// > You can call the [DescribeRegions](https://help.aliyun.com/document_detail/86912.html) operation to query available region IDs.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The list of source data fields.
	SrcColumns []*string `json:"SrcColumns,omitempty" xml:"SrcColumns,omitempty" type:"Repeated"`
	// Specifies whether to test the real-time task. Valid values:
	//
	// - true
	//
	// - false
	//
	// Default value: false.
	//
	// example:
	//
	// true
	TryRun *bool `json:"TryRun,omitempty" xml:"TryRun,omitempty"`
	// The update columns. Typically, these are all the non-primary key columns of the destination table. When duplicate data is identified based on MatchColumns, updating the values of the UpdateColumns overwrites the old data with the new data.
	UpdateColumns []*string `json:"UpdateColumns,omitempty" xml:"UpdateColumns,omitempty" type:"Repeated"`
	// The write mode.
	//
	// example:
	//
	// INSERT/UPDATE/MERGE
	WriteMode *string `json:"WriteMode,omitempty" xml:"WriteMode,omitempty"`
}

func (s CreateStreamingJobRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateStreamingJobRequest) GoString() string {
	return s.String()
}

func (s *CreateStreamingJobRequest) GetAccount() *string {
	return s.Account
}

func (s *CreateStreamingJobRequest) GetConsistency() *string {
	return s.Consistency
}

func (s *CreateStreamingJobRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *CreateStreamingJobRequest) GetDataSourceId() *string {
	return s.DataSourceId
}

func (s *CreateStreamingJobRequest) GetDestColumns() []*string {
	return s.DestColumns
}

func (s *CreateStreamingJobRequest) GetDestDatabase() *string {
	return s.DestDatabase
}

func (s *CreateStreamingJobRequest) GetDestSchema() *string {
	return s.DestSchema
}

func (s *CreateStreamingJobRequest) GetDestTable() *string {
	return s.DestTable
}

func (s *CreateStreamingJobRequest) GetErrorLimitCount() *int64 {
	return s.ErrorLimitCount
}

func (s *CreateStreamingJobRequest) GetFallbackOffset() *string {
	return s.FallbackOffset
}

func (s *CreateStreamingJobRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *CreateStreamingJobRequest) GetJobConfig() *string {
	return s.JobConfig
}

func (s *CreateStreamingJobRequest) GetJobDescription() *string {
	return s.JobDescription
}

func (s *CreateStreamingJobRequest) GetJobName() *string {
	return s.JobName
}

func (s *CreateStreamingJobRequest) GetMatchColumns() []*string {
	return s.MatchColumns
}

func (s *CreateStreamingJobRequest) GetMode() *string {
	return s.Mode
}

func (s *CreateStreamingJobRequest) GetPassword() *string {
	return s.Password
}

func (s *CreateStreamingJobRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateStreamingJobRequest) GetSrcColumns() []*string {
	return s.SrcColumns
}

func (s *CreateStreamingJobRequest) GetTryRun() *bool {
	return s.TryRun
}

func (s *CreateStreamingJobRequest) GetUpdateColumns() []*string {
	return s.UpdateColumns
}

func (s *CreateStreamingJobRequest) GetWriteMode() *string {
	return s.WriteMode
}

func (s *CreateStreamingJobRequest) SetAccount(v string) *CreateStreamingJobRequest {
	s.Account = &v
	return s
}

func (s *CreateStreamingJobRequest) SetConsistency(v string) *CreateStreamingJobRequest {
	s.Consistency = &v
	return s
}

func (s *CreateStreamingJobRequest) SetDBInstanceId(v string) *CreateStreamingJobRequest {
	s.DBInstanceId = &v
	return s
}

func (s *CreateStreamingJobRequest) SetDataSourceId(v string) *CreateStreamingJobRequest {
	s.DataSourceId = &v
	return s
}

func (s *CreateStreamingJobRequest) SetDestColumns(v []*string) *CreateStreamingJobRequest {
	s.DestColumns = v
	return s
}

func (s *CreateStreamingJobRequest) SetDestDatabase(v string) *CreateStreamingJobRequest {
	s.DestDatabase = &v
	return s
}

func (s *CreateStreamingJobRequest) SetDestSchema(v string) *CreateStreamingJobRequest {
	s.DestSchema = &v
	return s
}

func (s *CreateStreamingJobRequest) SetDestTable(v string) *CreateStreamingJobRequest {
	s.DestTable = &v
	return s
}

func (s *CreateStreamingJobRequest) SetErrorLimitCount(v int64) *CreateStreamingJobRequest {
	s.ErrorLimitCount = &v
	return s
}

func (s *CreateStreamingJobRequest) SetFallbackOffset(v string) *CreateStreamingJobRequest {
	s.FallbackOffset = &v
	return s
}

func (s *CreateStreamingJobRequest) SetGroupName(v string) *CreateStreamingJobRequest {
	s.GroupName = &v
	return s
}

func (s *CreateStreamingJobRequest) SetJobConfig(v string) *CreateStreamingJobRequest {
	s.JobConfig = &v
	return s
}

func (s *CreateStreamingJobRequest) SetJobDescription(v string) *CreateStreamingJobRequest {
	s.JobDescription = &v
	return s
}

func (s *CreateStreamingJobRequest) SetJobName(v string) *CreateStreamingJobRequest {
	s.JobName = &v
	return s
}

func (s *CreateStreamingJobRequest) SetMatchColumns(v []*string) *CreateStreamingJobRequest {
	s.MatchColumns = v
	return s
}

func (s *CreateStreamingJobRequest) SetMode(v string) *CreateStreamingJobRequest {
	s.Mode = &v
	return s
}

func (s *CreateStreamingJobRequest) SetPassword(v string) *CreateStreamingJobRequest {
	s.Password = &v
	return s
}

func (s *CreateStreamingJobRequest) SetRegionId(v string) *CreateStreamingJobRequest {
	s.RegionId = &v
	return s
}

func (s *CreateStreamingJobRequest) SetSrcColumns(v []*string) *CreateStreamingJobRequest {
	s.SrcColumns = v
	return s
}

func (s *CreateStreamingJobRequest) SetTryRun(v bool) *CreateStreamingJobRequest {
	s.TryRun = &v
	return s
}

func (s *CreateStreamingJobRequest) SetUpdateColumns(v []*string) *CreateStreamingJobRequest {
	s.UpdateColumns = v
	return s
}

func (s *CreateStreamingJobRequest) SetWriteMode(v string) *CreateStreamingJobRequest {
	s.WriteMode = &v
	return s
}

func (s *CreateStreamingJobRequest) Validate() error {
	return dara.Validate(s)
}
