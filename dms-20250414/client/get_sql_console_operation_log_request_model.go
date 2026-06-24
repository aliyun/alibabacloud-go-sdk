// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSqlConsoleOperationLogRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *GetSqlConsoleOperationLogRequest
	GetEndTime() *string
	SetInstanceId(v int64) *GetSqlConsoleOperationLogRequest
	GetInstanceId() *int64
	SetPageNumber(v int32) *GetSqlConsoleOperationLogRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *GetSqlConsoleOperationLogRequest
	GetPageSize() *int32
	SetSchema(v string) *GetSqlConsoleOperationLogRequest
	GetSchema() *string
	SetSqlType(v string) *GetSqlConsoleOperationLogRequest
	GetSqlType() *string
	SetStartTime(v string) *GetSqlConsoleOperationLogRequest
	GetStartTime() *string
	SetUsername(v string) *GetSqlConsoleOperationLogRequest
	GetUsername() *string
}

type GetSqlConsoleOperationLogRequest struct {
	// The end time of the logs.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2026-06-01 00:00:00
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// 1
	InstanceId *int64 `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The page number. Pages start from page 1.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Maximum value: 100.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The database schema.
	//
	// example:
	//
	// mysql
	Schema *string `json:"Schema,omitempty" xml:"Schema,omitempty"`
	// The type of the SQL statement.
	//
	// example:
	//
	// SELECT
	SqlType *string `json:"SqlType,omitempty" xml:"SqlType,omitempty"`
	// The start time of the logs.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2026-01-01 00:00:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The username.
	//
	// example:
	//
	// user
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s GetSqlConsoleOperationLogRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSqlConsoleOperationLogRequest) GoString() string {
	return s.String()
}

func (s *GetSqlConsoleOperationLogRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *GetSqlConsoleOperationLogRequest) GetInstanceId() *int64 {
	return s.InstanceId
}

func (s *GetSqlConsoleOperationLogRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *GetSqlConsoleOperationLogRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetSqlConsoleOperationLogRequest) GetSchema() *string {
	return s.Schema
}

func (s *GetSqlConsoleOperationLogRequest) GetSqlType() *string {
	return s.SqlType
}

func (s *GetSqlConsoleOperationLogRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *GetSqlConsoleOperationLogRequest) GetUsername() *string {
	return s.Username
}

func (s *GetSqlConsoleOperationLogRequest) SetEndTime(v string) *GetSqlConsoleOperationLogRequest {
	s.EndTime = &v
	return s
}

func (s *GetSqlConsoleOperationLogRequest) SetInstanceId(v int64) *GetSqlConsoleOperationLogRequest {
	s.InstanceId = &v
	return s
}

func (s *GetSqlConsoleOperationLogRequest) SetPageNumber(v int32) *GetSqlConsoleOperationLogRequest {
	s.PageNumber = &v
	return s
}

func (s *GetSqlConsoleOperationLogRequest) SetPageSize(v int32) *GetSqlConsoleOperationLogRequest {
	s.PageSize = &v
	return s
}

func (s *GetSqlConsoleOperationLogRequest) SetSchema(v string) *GetSqlConsoleOperationLogRequest {
	s.Schema = &v
	return s
}

func (s *GetSqlConsoleOperationLogRequest) SetSqlType(v string) *GetSqlConsoleOperationLogRequest {
	s.SqlType = &v
	return s
}

func (s *GetSqlConsoleOperationLogRequest) SetStartTime(v string) *GetSqlConsoleOperationLogRequest {
	s.StartTime = &v
	return s
}

func (s *GetSqlConsoleOperationLogRequest) SetUsername(v string) *GetSqlConsoleOperationLogRequest {
	s.Username = &v
	return s
}

func (s *GetSqlConsoleOperationLogRequest) Validate() error {
	return dara.Validate(s)
}
