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
	DownloadUrl   *string                  `json:"downloadUrl,omitempty" xml:"downloadUrl,omitempty"`
	Error         *string                  `json:"error,omitempty" xml:"error,omitempty"`
	ErrorCode     *string                  `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ExecutionTime *int64                   `json:"executionTime,omitempty" xml:"executionTime,omitempty"`
	Index         *int32                   `json:"index,omitempty" xml:"index,omitempty"`
	RowCount      *int32                   `json:"rowCount,omitempty" xml:"rowCount,omitempty"`
	Schema        []*StatementResultSchema `json:"schema,omitempty" xml:"schema,omitempty" type:"Repeated"`
	Sql           *string                  `json:"sql,omitempty" xml:"sql,omitempty"`
	Status        *string                  `json:"status,omitempty" xml:"status,omitempty"`
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
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
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
