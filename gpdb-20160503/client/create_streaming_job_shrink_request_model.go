// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateStreamingJobShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccount(v string) *CreateStreamingJobShrinkRequest
	GetAccount() *string
	SetConsistency(v string) *CreateStreamingJobShrinkRequest
	GetConsistency() *string
	SetDBInstanceId(v string) *CreateStreamingJobShrinkRequest
	GetDBInstanceId() *string
	SetDataSourceId(v string) *CreateStreamingJobShrinkRequest
	GetDataSourceId() *string
	SetDestColumnsShrink(v string) *CreateStreamingJobShrinkRequest
	GetDestColumnsShrink() *string
	SetDestDatabase(v string) *CreateStreamingJobShrinkRequest
	GetDestDatabase() *string
	SetDestSchema(v string) *CreateStreamingJobShrinkRequest
	GetDestSchema() *string
	SetDestTable(v string) *CreateStreamingJobShrinkRequest
	GetDestTable() *string
	SetErrorLimitCount(v int64) *CreateStreamingJobShrinkRequest
	GetErrorLimitCount() *int64
	SetFallbackOffset(v string) *CreateStreamingJobShrinkRequest
	GetFallbackOffset() *string
	SetGroupName(v string) *CreateStreamingJobShrinkRequest
	GetGroupName() *string
	SetJobConfig(v string) *CreateStreamingJobShrinkRequest
	GetJobConfig() *string
	SetJobDescription(v string) *CreateStreamingJobShrinkRequest
	GetJobDescription() *string
	SetJobName(v string) *CreateStreamingJobShrinkRequest
	GetJobName() *string
	SetMatchColumnsShrink(v string) *CreateStreamingJobShrinkRequest
	GetMatchColumnsShrink() *string
	SetMode(v string) *CreateStreamingJobShrinkRequest
	GetMode() *string
	SetPassword(v string) *CreateStreamingJobShrinkRequest
	GetPassword() *string
	SetRegionId(v string) *CreateStreamingJobShrinkRequest
	GetRegionId() *string
	SetSrcColumnsShrink(v string) *CreateStreamingJobShrinkRequest
	GetSrcColumnsShrink() *string
	SetTryRun(v bool) *CreateStreamingJobShrinkRequest
	GetTryRun() *bool
	SetUpdateColumnsShrink(v string) *CreateStreamingJobShrinkRequest
	GetUpdateColumnsShrink() *string
	SetWriteMode(v string) *CreateStreamingJobShrinkRequest
	GetWriteMode() *string
}

type CreateStreamingJobShrinkRequest struct {
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
	DestColumnsShrink *string `json:"DestColumns,omitempty" xml:"DestColumns,omitempty"`
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
	MatchColumnsShrink *string `json:"MatchColumns,omitempty" xml:"MatchColumns,omitempty"`
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
	SrcColumnsShrink *string `json:"SrcColumns,omitempty" xml:"SrcColumns,omitempty"`
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
	UpdateColumnsShrink *string `json:"UpdateColumns,omitempty" xml:"UpdateColumns,omitempty"`
	// The write mode.
	//
	// example:
	//
	// INSERT/UPDATE/MERGE
	WriteMode *string `json:"WriteMode,omitempty" xml:"WriteMode,omitempty"`
}

func (s CreateStreamingJobShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateStreamingJobShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateStreamingJobShrinkRequest) GetAccount() *string {
	return s.Account
}

func (s *CreateStreamingJobShrinkRequest) GetConsistency() *string {
	return s.Consistency
}

func (s *CreateStreamingJobShrinkRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *CreateStreamingJobShrinkRequest) GetDataSourceId() *string {
	return s.DataSourceId
}

func (s *CreateStreamingJobShrinkRequest) GetDestColumnsShrink() *string {
	return s.DestColumnsShrink
}

func (s *CreateStreamingJobShrinkRequest) GetDestDatabase() *string {
	return s.DestDatabase
}

func (s *CreateStreamingJobShrinkRequest) GetDestSchema() *string {
	return s.DestSchema
}

func (s *CreateStreamingJobShrinkRequest) GetDestTable() *string {
	return s.DestTable
}

func (s *CreateStreamingJobShrinkRequest) GetErrorLimitCount() *int64 {
	return s.ErrorLimitCount
}

func (s *CreateStreamingJobShrinkRequest) GetFallbackOffset() *string {
	return s.FallbackOffset
}

func (s *CreateStreamingJobShrinkRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *CreateStreamingJobShrinkRequest) GetJobConfig() *string {
	return s.JobConfig
}

func (s *CreateStreamingJobShrinkRequest) GetJobDescription() *string {
	return s.JobDescription
}

func (s *CreateStreamingJobShrinkRequest) GetJobName() *string {
	return s.JobName
}

func (s *CreateStreamingJobShrinkRequest) GetMatchColumnsShrink() *string {
	return s.MatchColumnsShrink
}

func (s *CreateStreamingJobShrinkRequest) GetMode() *string {
	return s.Mode
}

func (s *CreateStreamingJobShrinkRequest) GetPassword() *string {
	return s.Password
}

func (s *CreateStreamingJobShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateStreamingJobShrinkRequest) GetSrcColumnsShrink() *string {
	return s.SrcColumnsShrink
}

func (s *CreateStreamingJobShrinkRequest) GetTryRun() *bool {
	return s.TryRun
}

func (s *CreateStreamingJobShrinkRequest) GetUpdateColumnsShrink() *string {
	return s.UpdateColumnsShrink
}

func (s *CreateStreamingJobShrinkRequest) GetWriteMode() *string {
	return s.WriteMode
}

func (s *CreateStreamingJobShrinkRequest) SetAccount(v string) *CreateStreamingJobShrinkRequest {
	s.Account = &v
	return s
}

func (s *CreateStreamingJobShrinkRequest) SetConsistency(v string) *CreateStreamingJobShrinkRequest {
	s.Consistency = &v
	return s
}

func (s *CreateStreamingJobShrinkRequest) SetDBInstanceId(v string) *CreateStreamingJobShrinkRequest {
	s.DBInstanceId = &v
	return s
}

func (s *CreateStreamingJobShrinkRequest) SetDataSourceId(v string) *CreateStreamingJobShrinkRequest {
	s.DataSourceId = &v
	return s
}

func (s *CreateStreamingJobShrinkRequest) SetDestColumnsShrink(v string) *CreateStreamingJobShrinkRequest {
	s.DestColumnsShrink = &v
	return s
}

func (s *CreateStreamingJobShrinkRequest) SetDestDatabase(v string) *CreateStreamingJobShrinkRequest {
	s.DestDatabase = &v
	return s
}

func (s *CreateStreamingJobShrinkRequest) SetDestSchema(v string) *CreateStreamingJobShrinkRequest {
	s.DestSchema = &v
	return s
}

func (s *CreateStreamingJobShrinkRequest) SetDestTable(v string) *CreateStreamingJobShrinkRequest {
	s.DestTable = &v
	return s
}

func (s *CreateStreamingJobShrinkRequest) SetErrorLimitCount(v int64) *CreateStreamingJobShrinkRequest {
	s.ErrorLimitCount = &v
	return s
}

func (s *CreateStreamingJobShrinkRequest) SetFallbackOffset(v string) *CreateStreamingJobShrinkRequest {
	s.FallbackOffset = &v
	return s
}

func (s *CreateStreamingJobShrinkRequest) SetGroupName(v string) *CreateStreamingJobShrinkRequest {
	s.GroupName = &v
	return s
}

func (s *CreateStreamingJobShrinkRequest) SetJobConfig(v string) *CreateStreamingJobShrinkRequest {
	s.JobConfig = &v
	return s
}

func (s *CreateStreamingJobShrinkRequest) SetJobDescription(v string) *CreateStreamingJobShrinkRequest {
	s.JobDescription = &v
	return s
}

func (s *CreateStreamingJobShrinkRequest) SetJobName(v string) *CreateStreamingJobShrinkRequest {
	s.JobName = &v
	return s
}

func (s *CreateStreamingJobShrinkRequest) SetMatchColumnsShrink(v string) *CreateStreamingJobShrinkRequest {
	s.MatchColumnsShrink = &v
	return s
}

func (s *CreateStreamingJobShrinkRequest) SetMode(v string) *CreateStreamingJobShrinkRequest {
	s.Mode = &v
	return s
}

func (s *CreateStreamingJobShrinkRequest) SetPassword(v string) *CreateStreamingJobShrinkRequest {
	s.Password = &v
	return s
}

func (s *CreateStreamingJobShrinkRequest) SetRegionId(v string) *CreateStreamingJobShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *CreateStreamingJobShrinkRequest) SetSrcColumnsShrink(v string) *CreateStreamingJobShrinkRequest {
	s.SrcColumnsShrink = &v
	return s
}

func (s *CreateStreamingJobShrinkRequest) SetTryRun(v bool) *CreateStreamingJobShrinkRequest {
	s.TryRun = &v
	return s
}

func (s *CreateStreamingJobShrinkRequest) SetUpdateColumnsShrink(v string) *CreateStreamingJobShrinkRequest {
	s.UpdateColumnsShrink = &v
	return s
}

func (s *CreateStreamingJobShrinkRequest) SetWriteMode(v string) *CreateStreamingJobShrinkRequest {
	s.WriteMode = &v
	return s
}

func (s *CreateStreamingJobShrinkRequest) Validate() error {
	return dara.Validate(s)
}
