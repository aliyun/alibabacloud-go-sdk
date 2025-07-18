// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyStreamingJobShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccount(v string) *ModifyStreamingJobShrinkRequest
	GetAccount() *string
	SetConsistency(v string) *ModifyStreamingJobShrinkRequest
	GetConsistency() *string
	SetDBInstanceId(v string) *ModifyStreamingJobShrinkRequest
	GetDBInstanceId() *string
	SetDestColumnsShrink(v string) *ModifyStreamingJobShrinkRequest
	GetDestColumnsShrink() *string
	SetDestDatabase(v string) *ModifyStreamingJobShrinkRequest
	GetDestDatabase() *string
	SetDestSchema(v string) *ModifyStreamingJobShrinkRequest
	GetDestSchema() *string
	SetDestTable(v string) *ModifyStreamingJobShrinkRequest
	GetDestTable() *string
	SetErrorLimitCount(v int64) *ModifyStreamingJobShrinkRequest
	GetErrorLimitCount() *int64
	SetFallbackOffset(v string) *ModifyStreamingJobShrinkRequest
	GetFallbackOffset() *string
	SetGroupName(v string) *ModifyStreamingJobShrinkRequest
	GetGroupName() *string
	SetJobConfig(v string) *ModifyStreamingJobShrinkRequest
	GetJobConfig() *string
	SetJobDescription(v string) *ModifyStreamingJobShrinkRequest
	GetJobDescription() *string
	SetJobId(v int64) *ModifyStreamingJobShrinkRequest
	GetJobId() *int64
	SetMatchColumnsShrink(v string) *ModifyStreamingJobShrinkRequest
	GetMatchColumnsShrink() *string
	SetPassword(v string) *ModifyStreamingJobShrinkRequest
	GetPassword() *string
	SetRegionId(v string) *ModifyStreamingJobShrinkRequest
	GetRegionId() *string
	SetSrcColumnsShrink(v string) *ModifyStreamingJobShrinkRequest
	GetSrcColumnsShrink() *string
	SetTryRun(v bool) *ModifyStreamingJobShrinkRequest
	GetTryRun() *bool
	SetUpdateColumnsShrink(v string) *ModifyStreamingJobShrinkRequest
	GetUpdateColumnsShrink() *string
	SetWriteMode(v string) *ModifyStreamingJobShrinkRequest
	GetWriteMode() *string
}

type ModifyStreamingJobShrinkRequest struct {
	// Account name.
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
	// Instance ID
	//
	// This parameter is required.
	//
	// example:
	//
	// ModifyStreamingJob
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// Target data table mapping field list.
	DestColumnsShrink *string `json:"DestColumns,omitempty" xml:"DestColumns,omitempty"`
	// Target database name.
	//
	// example:
	//
	// dest-db
	DestDatabase *string `json:"DestDatabase,omitempty" xml:"DestDatabase,omitempty"`
	// Target schema.
	//
	// example:
	//
	// dest-schema
	DestSchema *string `json:"DestSchema,omitempty" xml:"DestSchema,omitempty"`
	// Target table name.
	//
	// example:
	//
	// dest-table
	DestTable *string `json:"DestTable,omitempty" xml:"DestTable,omitempty"`
	// When the data in Kafka does not match the ADBPG target table, it will cause a write failure. This value is the number of error rows allowed; exceeding this will cause the task to fail.
	//
	// example:
	//
	// 5
	ErrorLimitCount *int64 `json:"ErrorLimitCount,omitempty" xml:"ErrorLimitCount,omitempty"`
	// The fallback offset for data consumption.
	//
	// 	- This parameter specifies the starting offset from which data consumption resumes when a consumer does not request a consumption offset or requests a consumption offset that is beyond the range of the offset information recorded in the current Kafka cluster. You can choose to start data consumption from the earliest or latest offset.
	//
	// Valid values:
	//
	// 	- EARLIEST
	//
	// 	- LATEST
	//
	// example:
	//
	// EARLIEST /  LATEST
	FallbackOffset *string `json:"FallbackOffset,omitempty" xml:"FallbackOffset,omitempty"`
	// Kafka group name
	//
	// example:
	//
	// group_name
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// Job configuration file, required for professional mode.
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
	// Job description.
	//
	// example:
	//
	// test-job
	JobDescription *string `json:"JobDescription,omitempty" xml:"JobDescription,omitempty"`
	// Job ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	JobId *int64 `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// Match columns, usually all primary key columns of the target table. If all column values in this configuration are the same, the two rows of data are considered duplicates.
	MatchColumnsShrink *string `json:"MatchColumns,omitempty" xml:"MatchColumns,omitempty"`
	// Password.
	//
	// example:
	//
	// pwd123
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// Region ID.
	//
	// > You can call the [DescribeRegions](https://help.aliyun.com/document_detail/86912.html) API to view available region IDs.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Source data field list.
	SrcColumnsShrink *string `json:"SrcColumns,omitempty" xml:"SrcColumns,omitempty"`
	// Specifies whether to test the real-time job. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// Default value: false.
	//
	// example:
	//
	// true
	TryRun *bool `json:"TryRun,omitempty" xml:"TryRun,omitempty"`
	// Update columns, usually all non-primary key columns of the target table. When data duplication is determined through MatchColumns, updating the UpdateColumns column values will result in new data overwriting old data.
	UpdateColumnsShrink *string `json:"UpdateColumns,omitempty" xml:"UpdateColumns,omitempty"`
	// The write mode.
	//
	// Valid values:
	//
	// 	- MERGE
	//
	// 	- INSERT
	//
	// 	- UPDATE
	//
	// example:
	//
	// INSERT/UPDATE/MERGE
	WriteMode *string `json:"WriteMode,omitempty" xml:"WriteMode,omitempty"`
}

func (s ModifyStreamingJobShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyStreamingJobShrinkRequest) GoString() string {
	return s.String()
}

func (s *ModifyStreamingJobShrinkRequest) GetAccount() *string {
	return s.Account
}

func (s *ModifyStreamingJobShrinkRequest) GetConsistency() *string {
	return s.Consistency
}

func (s *ModifyStreamingJobShrinkRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ModifyStreamingJobShrinkRequest) GetDestColumnsShrink() *string {
	return s.DestColumnsShrink
}

func (s *ModifyStreamingJobShrinkRequest) GetDestDatabase() *string {
	return s.DestDatabase
}

func (s *ModifyStreamingJobShrinkRequest) GetDestSchema() *string {
	return s.DestSchema
}

func (s *ModifyStreamingJobShrinkRequest) GetDestTable() *string {
	return s.DestTable
}

func (s *ModifyStreamingJobShrinkRequest) GetErrorLimitCount() *int64 {
	return s.ErrorLimitCount
}

func (s *ModifyStreamingJobShrinkRequest) GetFallbackOffset() *string {
	return s.FallbackOffset
}

func (s *ModifyStreamingJobShrinkRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *ModifyStreamingJobShrinkRequest) GetJobConfig() *string {
	return s.JobConfig
}

func (s *ModifyStreamingJobShrinkRequest) GetJobDescription() *string {
	return s.JobDescription
}

func (s *ModifyStreamingJobShrinkRequest) GetJobId() *int64 {
	return s.JobId
}

func (s *ModifyStreamingJobShrinkRequest) GetMatchColumnsShrink() *string {
	return s.MatchColumnsShrink
}

func (s *ModifyStreamingJobShrinkRequest) GetPassword() *string {
	return s.Password
}

func (s *ModifyStreamingJobShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyStreamingJobShrinkRequest) GetSrcColumnsShrink() *string {
	return s.SrcColumnsShrink
}

func (s *ModifyStreamingJobShrinkRequest) GetTryRun() *bool {
	return s.TryRun
}

func (s *ModifyStreamingJobShrinkRequest) GetUpdateColumnsShrink() *string {
	return s.UpdateColumnsShrink
}

func (s *ModifyStreamingJobShrinkRequest) GetWriteMode() *string {
	return s.WriteMode
}

func (s *ModifyStreamingJobShrinkRequest) SetAccount(v string) *ModifyStreamingJobShrinkRequest {
	s.Account = &v
	return s
}

func (s *ModifyStreamingJobShrinkRequest) SetConsistency(v string) *ModifyStreamingJobShrinkRequest {
	s.Consistency = &v
	return s
}

func (s *ModifyStreamingJobShrinkRequest) SetDBInstanceId(v string) *ModifyStreamingJobShrinkRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyStreamingJobShrinkRequest) SetDestColumnsShrink(v string) *ModifyStreamingJobShrinkRequest {
	s.DestColumnsShrink = &v
	return s
}

func (s *ModifyStreamingJobShrinkRequest) SetDestDatabase(v string) *ModifyStreamingJobShrinkRequest {
	s.DestDatabase = &v
	return s
}

func (s *ModifyStreamingJobShrinkRequest) SetDestSchema(v string) *ModifyStreamingJobShrinkRequest {
	s.DestSchema = &v
	return s
}

func (s *ModifyStreamingJobShrinkRequest) SetDestTable(v string) *ModifyStreamingJobShrinkRequest {
	s.DestTable = &v
	return s
}

func (s *ModifyStreamingJobShrinkRequest) SetErrorLimitCount(v int64) *ModifyStreamingJobShrinkRequest {
	s.ErrorLimitCount = &v
	return s
}

func (s *ModifyStreamingJobShrinkRequest) SetFallbackOffset(v string) *ModifyStreamingJobShrinkRequest {
	s.FallbackOffset = &v
	return s
}

func (s *ModifyStreamingJobShrinkRequest) SetGroupName(v string) *ModifyStreamingJobShrinkRequest {
	s.GroupName = &v
	return s
}

func (s *ModifyStreamingJobShrinkRequest) SetJobConfig(v string) *ModifyStreamingJobShrinkRequest {
	s.JobConfig = &v
	return s
}

func (s *ModifyStreamingJobShrinkRequest) SetJobDescription(v string) *ModifyStreamingJobShrinkRequest {
	s.JobDescription = &v
	return s
}

func (s *ModifyStreamingJobShrinkRequest) SetJobId(v int64) *ModifyStreamingJobShrinkRequest {
	s.JobId = &v
	return s
}

func (s *ModifyStreamingJobShrinkRequest) SetMatchColumnsShrink(v string) *ModifyStreamingJobShrinkRequest {
	s.MatchColumnsShrink = &v
	return s
}

func (s *ModifyStreamingJobShrinkRequest) SetPassword(v string) *ModifyStreamingJobShrinkRequest {
	s.Password = &v
	return s
}

func (s *ModifyStreamingJobShrinkRequest) SetRegionId(v string) *ModifyStreamingJobShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyStreamingJobShrinkRequest) SetSrcColumnsShrink(v string) *ModifyStreamingJobShrinkRequest {
	s.SrcColumnsShrink = &v
	return s
}

func (s *ModifyStreamingJobShrinkRequest) SetTryRun(v bool) *ModifyStreamingJobShrinkRequest {
	s.TryRun = &v
	return s
}

func (s *ModifyStreamingJobShrinkRequest) SetUpdateColumnsShrink(v string) *ModifyStreamingJobShrinkRequest {
	s.UpdateColumnsShrink = &v
	return s
}

func (s *ModifyStreamingJobShrinkRequest) SetWriteMode(v string) *ModifyStreamingJobShrinkRequest {
	s.WriteMode = &v
	return s
}

func (s *ModifyStreamingJobShrinkRequest) Validate() error {
	return dara.Validate(s)
}
