// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSparkBatchSQLStatement interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *SparkBatchSQLStatement
	GetAppId() *string
	SetCode(v string) *SparkBatchSQLStatement
	GetCode() *string
	SetEndTime(v int64) *SparkBatchSQLStatement
	GetEndTime() *int64
	SetError(v string) *SparkBatchSQLStatement
	GetError() *string
	SetResult(v string) *SparkBatchSQLStatement
	GetResult() *string
	SetResultUri(v string) *SparkBatchSQLStatement
	GetResultUri() *string
	SetStartTime(v int64) *SparkBatchSQLStatement
	GetStartTime() *int64
	SetState(v string) *SparkBatchSQLStatement
	GetState() *string
	SetStatementId(v string) *SparkBatchSQLStatement
	GetStatementId() *string
}

type SparkBatchSQLStatement struct {
	// The ID of the Spark job that executes the SQL statement.
	//
	// example:
	//
	// s20240122jsdx****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The text of the SQL statement.
	//
	// example:
	//
	// SELECT 100
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The execution end timestamp of the SQL statement.
	//
	// example:
	//
	// 1723521767000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The error message for the SQL statement execution.
	//
	// example:
	//
	// table is not found
	Error *string `json:"Error,omitempty" xml:"Error,omitempty"`
	// The execution result of the SQL statement. The content is in the response format of the `DataFrame Show` statement in Spark.
	//
	// example:
	//
	// +---+-----+
	//
	// |age| name|
	//
	// +---+-----+
	//
	// | 14|  Tom|
	//
	// | 23|Alice|
	//
	// +---+-----+
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
	// The Object Storage Service (OSS) location where the execution results of the SQL statement are stored. If you do not specify a storage location, null is returned.
	//
	// example:
	//
	// oss://yourbucket/result.json
	ResultUri *string `json:"ResultUri,omitempty" xml:"ResultUri,omitempty"`
	// The execution start timestamp of the SQL statement.
	//
	// example:
	//
	// 1723521767000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The execution status of the SQL statement. Valid values:
	//
	// 	- FAILED
	//
	// 	- RUNNING
	//
	// 	- CANCELED
	//
	// 	- PENDING
	//
	// 	- FINISHED
	//
	// example:
	//
	// RUNNING
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The ID of the SQL statement.
	//
	// example:
	//
	// sq1723521767****-0001
	StatementId *string `json:"StatementId,omitempty" xml:"StatementId,omitempty"`
}

func (s SparkBatchSQLStatement) String() string {
	return dara.Prettify(s)
}

func (s SparkBatchSQLStatement) GoString() string {
	return s.String()
}

func (s *SparkBatchSQLStatement) GetAppId() *string {
	return s.AppId
}

func (s *SparkBatchSQLStatement) GetCode() *string {
	return s.Code
}

func (s *SparkBatchSQLStatement) GetEndTime() *int64 {
	return s.EndTime
}

func (s *SparkBatchSQLStatement) GetError() *string {
	return s.Error
}

func (s *SparkBatchSQLStatement) GetResult() *string {
	return s.Result
}

func (s *SparkBatchSQLStatement) GetResultUri() *string {
	return s.ResultUri
}

func (s *SparkBatchSQLStatement) GetStartTime() *int64 {
	return s.StartTime
}

func (s *SparkBatchSQLStatement) GetState() *string {
	return s.State
}

func (s *SparkBatchSQLStatement) GetStatementId() *string {
	return s.StatementId
}

func (s *SparkBatchSQLStatement) SetAppId(v string) *SparkBatchSQLStatement {
	s.AppId = &v
	return s
}

func (s *SparkBatchSQLStatement) SetCode(v string) *SparkBatchSQLStatement {
	s.Code = &v
	return s
}

func (s *SparkBatchSQLStatement) SetEndTime(v int64) *SparkBatchSQLStatement {
	s.EndTime = &v
	return s
}

func (s *SparkBatchSQLStatement) SetError(v string) *SparkBatchSQLStatement {
	s.Error = &v
	return s
}

func (s *SparkBatchSQLStatement) SetResult(v string) *SparkBatchSQLStatement {
	s.Result = &v
	return s
}

func (s *SparkBatchSQLStatement) SetResultUri(v string) *SparkBatchSQLStatement {
	s.ResultUri = &v
	return s
}

func (s *SparkBatchSQLStatement) SetStartTime(v int64) *SparkBatchSQLStatement {
	s.StartTime = &v
	return s
}

func (s *SparkBatchSQLStatement) SetState(v string) *SparkBatchSQLStatement {
	s.State = &v
	return s
}

func (s *SparkBatchSQLStatement) SetStatementId(v string) *SparkBatchSQLStatement {
	s.StatementId = &v
	return s
}

func (s *SparkBatchSQLStatement) Validate() error {
	return dara.Validate(s)
}
