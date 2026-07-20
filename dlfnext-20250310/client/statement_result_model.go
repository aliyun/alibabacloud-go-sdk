// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStatementResult interface {
	dara.Model
	String() string
	GoString() string
	SetDownloadUrl(v string) *StatementResult
	GetDownloadUrl() *string
	SetError(v string) *StatementResult
	GetError() *string
	SetErrorCode(v string) *StatementResult
	GetErrorCode() *string
	SetExecutionTime(v int64) *StatementResult
	GetExecutionTime() *int64
	SetIndex(v int32) *StatementResult
	GetIndex() *int32
	SetRowCount(v int32) *StatementResult
	GetRowCount() *int32
	SetSchema(v []*StatementResultSchema) *StatementResult
	GetSchema() []*StatementResultSchema
	SetSql(v string) *StatementResult
	GetSql() *string
	SetStatus(v string) *StatementResult
	GetStatus() *string
}

type StatementResult struct {
	// The presigned URL of the Arrow IPC file. This parameter is returned when a result set exists. The URL is valid for 1 hour and contains full data. The value is null for an empty result set (rowCount == 0).
	//
	// example:
	//
	// https://xxx.oss-cn-hangzhou.aliyuncs.com/xxxx
	DownloadUrl *string `json:"downloadUrl,omitempty" xml:"downloadUrl,omitempty"`
	// The error message. This parameter is returned only when the status is FAILED.
	//
	// example:
	//
	// SQL_ERROR
	Error *string `json:"error,omitempty" xml:"error,omitempty"`
	// The error code. This parameter is returned only when the status is FAILED.
	//
	// example:
	//
	// This feature is not implemented: xxx
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// The execution duration of the statement, in milliseconds.
	//
	// example:
	//
	// 100
	ExecutionTime *int64 `json:"executionTime,omitempty" xml:"executionTime,omitempty"`
	// The statement sequence number (0-based).
	//
	// example:
	//
	// 0
	Index *int32 `json:"index,omitempty" xml:"index,omitempty"`
	// The total number of rows in the result. The value is 0 for statements that do not return a result set.
	//
	// example:
	//
	// 1000
	RowCount *int32 `json:"rowCount,omitempty" xml:"rowCount,omitempty"`
	// The result column information. This parameter is returned when a result set exists.
	Schema []*StatementResultSchema `json:"schema,omitempty" xml:"schema,omitempty" type:"Repeated"`
	// The SQL text of the statement.
	//
	// example:
	//
	// select 	- from table_name;
	Sql *string `json:"sql,omitempty" xml:"sql,omitempty"`
	// The status of the statement. Valid values: COMPLETED and FAILED.
	//
	// example:
	//
	// COMPLETED
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s StatementResult) String() string {
	return dara.Prettify(s)
}

func (s StatementResult) GoString() string {
	return s.String()
}

func (s *StatementResult) GetDownloadUrl() *string {
	return s.DownloadUrl
}

func (s *StatementResult) GetError() *string {
	return s.Error
}

func (s *StatementResult) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *StatementResult) GetExecutionTime() *int64 {
	return s.ExecutionTime
}

func (s *StatementResult) GetIndex() *int32 {
	return s.Index
}

func (s *StatementResult) GetRowCount() *int32 {
	return s.RowCount
}

func (s *StatementResult) GetSchema() []*StatementResultSchema {
	return s.Schema
}

func (s *StatementResult) GetSql() *string {
	return s.Sql
}

func (s *StatementResult) GetStatus() *string {
	return s.Status
}

func (s *StatementResult) SetDownloadUrl(v string) *StatementResult {
	s.DownloadUrl = &v
	return s
}

func (s *StatementResult) SetError(v string) *StatementResult {
	s.Error = &v
	return s
}

func (s *StatementResult) SetErrorCode(v string) *StatementResult {
	s.ErrorCode = &v
	return s
}

func (s *StatementResult) SetExecutionTime(v int64) *StatementResult {
	s.ExecutionTime = &v
	return s
}

func (s *StatementResult) SetIndex(v int32) *StatementResult {
	s.Index = &v
	return s
}

func (s *StatementResult) SetRowCount(v int32) *StatementResult {
	s.RowCount = &v
	return s
}

func (s *StatementResult) SetSchema(v []*StatementResultSchema) *StatementResult {
	s.Schema = v
	return s
}

func (s *StatementResult) SetSql(v string) *StatementResult {
	s.Sql = &v
	return s
}

func (s *StatementResult) SetStatus(v string) *StatementResult {
	s.Status = &v
	return s
}

func (s *StatementResult) Validate() error {
	if s.Schema != nil {
		for _, item := range s.Schema {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type StatementResultSchema struct {
	// The column name.
	//
	// example:
	//
	// name
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The data type.
	//
	// example:
	//
	// string
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s StatementResultSchema) String() string {
	return dara.Prettify(s)
}

func (s StatementResultSchema) GoString() string {
	return s.String()
}

func (s *StatementResultSchema) GetName() *string {
	return s.Name
}

func (s *StatementResultSchema) GetType() *string {
	return s.Type
}

func (s *StatementResultSchema) SetName(v string) *StatementResultSchema {
	s.Name = &v
	return s
}

func (s *StatementResultSchema) SetType(v string) *StatementResultSchema {
	s.Type = &v
	return s
}

func (s *StatementResultSchema) Validate() error {
	return dara.Validate(s)
}
