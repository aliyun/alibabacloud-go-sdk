// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSparkBatchSQL interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *SparkBatchSQL
	GetDBClusterId() *string
	SetErrMessage(v string) *SparkBatchSQL
	GetErrMessage() *string
	SetQuery(v string) *SparkBatchSQL
	GetQuery() *string
	SetQueryEndTime(v int64) *SparkBatchSQL
	GetQueryEndTime() *int64
	SetQueryId(v string) *SparkBatchSQL
	GetQueryId() *string
	SetQueryStartTime(v int64) *SparkBatchSQL
	GetQueryStartTime() *int64
	SetQueryState(v string) *SparkBatchSQL
	GetQueryState() *string
	SetQuerySubmissionTime(v int64) *SparkBatchSQL
	GetQuerySubmissionTime() *int64
	SetResourceGroupName(v string) *SparkBatchSQL
	GetResourceGroupName() *string
	SetSchema(v string) *SparkBatchSQL
	GetSchema() *string
	SetStatements(v []*SparkBatchSQLStatement) *SparkBatchSQL
	GetStatements() []*SparkBatchSQLStatement
	SetUid(v int64) *SparkBatchSQL
	GetUid() *int64
}

type SparkBatchSQL struct {
	// example:
	//
	// amv-xxx
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// example:
	//
	// schema is not found
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// example:
	//
	// query
	Query *string `json:"Query,omitempty" xml:"Query,omitempty"`
	// example:
	//
	// 1723521767000
	QueryEndTime *int64 `json:"QueryEndTime,omitempty" xml:"QueryEndTime,omitempty"`
	// example:
	//
	// amv-202401-xx
	QueryId *string `json:"QueryId,omitempty" xml:"QueryId,omitempty"`
	// example:
	//
	// 1723521767000
	QueryStartTime *int64 `json:"QueryStartTime,omitempty" xml:"QueryStartTime,omitempty"`
	// example:
	//
	// RUNNING
	QueryState *string `json:"QueryState,omitempty" xml:"QueryState,omitempty"`
	// example:
	//
	// 1723521767000
	QuerySubmissionTime *int64 `json:"QuerySubmissionTime,omitempty" xml:"QuerySubmissionTime,omitempty"`
	// example:
	//
	// spark_rg
	ResourceGroupName *string `json:"ResourceGroupName,omitempty" xml:"ResourceGroupName,omitempty"`
	// example:
	//
	// default
	Schema     *string                   `json:"Schema,omitempty" xml:"Schema,omitempty"`
	Statements []*SparkBatchSQLStatement `json:"Statements,omitempty" xml:"Statements,omitempty" type:"Repeated"`
	// example:
	//
	// 12222222
	Uid *int64 `json:"Uid,omitempty" xml:"Uid,omitempty"`
}

func (s SparkBatchSQL) String() string {
	return dara.Prettify(s)
}

func (s SparkBatchSQL) GoString() string {
	return s.String()
}

func (s *SparkBatchSQL) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *SparkBatchSQL) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *SparkBatchSQL) GetQuery() *string {
	return s.Query
}

func (s *SparkBatchSQL) GetQueryEndTime() *int64 {
	return s.QueryEndTime
}

func (s *SparkBatchSQL) GetQueryId() *string {
	return s.QueryId
}

func (s *SparkBatchSQL) GetQueryStartTime() *int64 {
	return s.QueryStartTime
}

func (s *SparkBatchSQL) GetQueryState() *string {
	return s.QueryState
}

func (s *SparkBatchSQL) GetQuerySubmissionTime() *int64 {
	return s.QuerySubmissionTime
}

func (s *SparkBatchSQL) GetResourceGroupName() *string {
	return s.ResourceGroupName
}

func (s *SparkBatchSQL) GetSchema() *string {
	return s.Schema
}

func (s *SparkBatchSQL) GetStatements() []*SparkBatchSQLStatement {
	return s.Statements
}

func (s *SparkBatchSQL) GetUid() *int64 {
	return s.Uid
}

func (s *SparkBatchSQL) SetDBClusterId(v string) *SparkBatchSQL {
	s.DBClusterId = &v
	return s
}

func (s *SparkBatchSQL) SetErrMessage(v string) *SparkBatchSQL {
	s.ErrMessage = &v
	return s
}

func (s *SparkBatchSQL) SetQuery(v string) *SparkBatchSQL {
	s.Query = &v
	return s
}

func (s *SparkBatchSQL) SetQueryEndTime(v int64) *SparkBatchSQL {
	s.QueryEndTime = &v
	return s
}

func (s *SparkBatchSQL) SetQueryId(v string) *SparkBatchSQL {
	s.QueryId = &v
	return s
}

func (s *SparkBatchSQL) SetQueryStartTime(v int64) *SparkBatchSQL {
	s.QueryStartTime = &v
	return s
}

func (s *SparkBatchSQL) SetQueryState(v string) *SparkBatchSQL {
	s.QueryState = &v
	return s
}

func (s *SparkBatchSQL) SetQuerySubmissionTime(v int64) *SparkBatchSQL {
	s.QuerySubmissionTime = &v
	return s
}

func (s *SparkBatchSQL) SetResourceGroupName(v string) *SparkBatchSQL {
	s.ResourceGroupName = &v
	return s
}

func (s *SparkBatchSQL) SetSchema(v string) *SparkBatchSQL {
	s.Schema = &v
	return s
}

func (s *SparkBatchSQL) SetStatements(v []*SparkBatchSQLStatement) *SparkBatchSQL {
	s.Statements = v
	return s
}

func (s *SparkBatchSQL) SetUid(v int64) *SparkBatchSQL {
	s.Uid = &v
	return s
}

func (s *SparkBatchSQL) Validate() error {
	return dara.Validate(s)
}
