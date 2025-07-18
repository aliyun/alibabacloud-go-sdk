// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyStreamingJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccount(v string) *ModifyStreamingJobRequest
	GetAccount() *string
	SetConsistency(v string) *ModifyStreamingJobRequest
	GetConsistency() *string
	SetDBInstanceId(v string) *ModifyStreamingJobRequest
	GetDBInstanceId() *string
	SetDestColumns(v []*string) *ModifyStreamingJobRequest
	GetDestColumns() []*string
	SetDestDatabase(v string) *ModifyStreamingJobRequest
	GetDestDatabase() *string
	SetDestSchema(v string) *ModifyStreamingJobRequest
	GetDestSchema() *string
	SetDestTable(v string) *ModifyStreamingJobRequest
	GetDestTable() *string
	SetErrorLimitCount(v int64) *ModifyStreamingJobRequest
	GetErrorLimitCount() *int64
	SetFallbackOffset(v string) *ModifyStreamingJobRequest
	GetFallbackOffset() *string
	SetGroupName(v string) *ModifyStreamingJobRequest
	GetGroupName() *string
	SetJobConfig(v string) *ModifyStreamingJobRequest
	GetJobConfig() *string
	SetJobDescription(v string) *ModifyStreamingJobRequest
	GetJobDescription() *string
	SetJobId(v int64) *ModifyStreamingJobRequest
	GetJobId() *int64
	SetMatchColumns(v []*string) *ModifyStreamingJobRequest
	GetMatchColumns() []*string
	SetPassword(v string) *ModifyStreamingJobRequest
	GetPassword() *string
	SetRegionId(v string) *ModifyStreamingJobRequest
	GetRegionId() *string
	SetSrcColumns(v []*string) *ModifyStreamingJobRequest
	GetSrcColumns() []*string
	SetTryRun(v bool) *ModifyStreamingJobRequest
	GetTryRun() *bool
	SetUpdateColumns(v []*string) *ModifyStreamingJobRequest
	GetUpdateColumns() []*string
	SetWriteMode(v string) *ModifyStreamingJobRequest
	GetWriteMode() *string
}

type ModifyStreamingJobRequest struct {
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
	DestColumns []*string `json:"DestColumns,omitempty" xml:"DestColumns,omitempty" type:"Repeated"`
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
	MatchColumns []*string `json:"MatchColumns,omitempty" xml:"MatchColumns,omitempty" type:"Repeated"`
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
	SrcColumns []*string `json:"SrcColumns,omitempty" xml:"SrcColumns,omitempty" type:"Repeated"`
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
	UpdateColumns []*string `json:"UpdateColumns,omitempty" xml:"UpdateColumns,omitempty" type:"Repeated"`
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

func (s ModifyStreamingJobRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyStreamingJobRequest) GoString() string {
	return s.String()
}

func (s *ModifyStreamingJobRequest) GetAccount() *string {
	return s.Account
}

func (s *ModifyStreamingJobRequest) GetConsistency() *string {
	return s.Consistency
}

func (s *ModifyStreamingJobRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ModifyStreamingJobRequest) GetDestColumns() []*string {
	return s.DestColumns
}

func (s *ModifyStreamingJobRequest) GetDestDatabase() *string {
	return s.DestDatabase
}

func (s *ModifyStreamingJobRequest) GetDestSchema() *string {
	return s.DestSchema
}

func (s *ModifyStreamingJobRequest) GetDestTable() *string {
	return s.DestTable
}

func (s *ModifyStreamingJobRequest) GetErrorLimitCount() *int64 {
	return s.ErrorLimitCount
}

func (s *ModifyStreamingJobRequest) GetFallbackOffset() *string {
	return s.FallbackOffset
}

func (s *ModifyStreamingJobRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *ModifyStreamingJobRequest) GetJobConfig() *string {
	return s.JobConfig
}

func (s *ModifyStreamingJobRequest) GetJobDescription() *string {
	return s.JobDescription
}

func (s *ModifyStreamingJobRequest) GetJobId() *int64 {
	return s.JobId
}

func (s *ModifyStreamingJobRequest) GetMatchColumns() []*string {
	return s.MatchColumns
}

func (s *ModifyStreamingJobRequest) GetPassword() *string {
	return s.Password
}

func (s *ModifyStreamingJobRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyStreamingJobRequest) GetSrcColumns() []*string {
	return s.SrcColumns
}

func (s *ModifyStreamingJobRequest) GetTryRun() *bool {
	return s.TryRun
}

func (s *ModifyStreamingJobRequest) GetUpdateColumns() []*string {
	return s.UpdateColumns
}

func (s *ModifyStreamingJobRequest) GetWriteMode() *string {
	return s.WriteMode
}

func (s *ModifyStreamingJobRequest) SetAccount(v string) *ModifyStreamingJobRequest {
	s.Account = &v
	return s
}

func (s *ModifyStreamingJobRequest) SetConsistency(v string) *ModifyStreamingJobRequest {
	s.Consistency = &v
	return s
}

func (s *ModifyStreamingJobRequest) SetDBInstanceId(v string) *ModifyStreamingJobRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyStreamingJobRequest) SetDestColumns(v []*string) *ModifyStreamingJobRequest {
	s.DestColumns = v
	return s
}

func (s *ModifyStreamingJobRequest) SetDestDatabase(v string) *ModifyStreamingJobRequest {
	s.DestDatabase = &v
	return s
}

func (s *ModifyStreamingJobRequest) SetDestSchema(v string) *ModifyStreamingJobRequest {
	s.DestSchema = &v
	return s
}

func (s *ModifyStreamingJobRequest) SetDestTable(v string) *ModifyStreamingJobRequest {
	s.DestTable = &v
	return s
}

func (s *ModifyStreamingJobRequest) SetErrorLimitCount(v int64) *ModifyStreamingJobRequest {
	s.ErrorLimitCount = &v
	return s
}

func (s *ModifyStreamingJobRequest) SetFallbackOffset(v string) *ModifyStreamingJobRequest {
	s.FallbackOffset = &v
	return s
}

func (s *ModifyStreamingJobRequest) SetGroupName(v string) *ModifyStreamingJobRequest {
	s.GroupName = &v
	return s
}

func (s *ModifyStreamingJobRequest) SetJobConfig(v string) *ModifyStreamingJobRequest {
	s.JobConfig = &v
	return s
}

func (s *ModifyStreamingJobRequest) SetJobDescription(v string) *ModifyStreamingJobRequest {
	s.JobDescription = &v
	return s
}

func (s *ModifyStreamingJobRequest) SetJobId(v int64) *ModifyStreamingJobRequest {
	s.JobId = &v
	return s
}

func (s *ModifyStreamingJobRequest) SetMatchColumns(v []*string) *ModifyStreamingJobRequest {
	s.MatchColumns = v
	return s
}

func (s *ModifyStreamingJobRequest) SetPassword(v string) *ModifyStreamingJobRequest {
	s.Password = &v
	return s
}

func (s *ModifyStreamingJobRequest) SetRegionId(v string) *ModifyStreamingJobRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyStreamingJobRequest) SetSrcColumns(v []*string) *ModifyStreamingJobRequest {
	s.SrcColumns = v
	return s
}

func (s *ModifyStreamingJobRequest) SetTryRun(v bool) *ModifyStreamingJobRequest {
	s.TryRun = &v
	return s
}

func (s *ModifyStreamingJobRequest) SetUpdateColumns(v []*string) *ModifyStreamingJobRequest {
	s.UpdateColumns = v
	return s
}

func (s *ModifyStreamingJobRequest) SetWriteMode(v string) *ModifyStreamingJobRequest {
	s.WriteMode = &v
	return s
}

func (s *ModifyStreamingJobRequest) Validate() error {
	return dara.Validate(s)
}
