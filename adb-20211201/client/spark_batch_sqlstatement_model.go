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
	// example:
	//
	// s20240122jsdx****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// SELECT 100
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 1723521767000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// table is not found
	Error *string `json:"Error,omitempty" xml:"Error,omitempty"`
	// example:
	//
	// JSON格式执行结果集
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
	// example:
	//
	// oss://yourbucket/result.json
	ResultUri *string `json:"ResultUri,omitempty" xml:"ResultUri,omitempty"`
	// example:
	//
	// 1723521767000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// RUNNING
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// example:
	//
	// amv-20240711-stmt
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
