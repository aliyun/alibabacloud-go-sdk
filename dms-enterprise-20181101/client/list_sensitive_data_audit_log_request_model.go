// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSensitiveDataAuditLogRequest interface {
	dara.Model
	String() string
	GoString() string
	SetColumnName(v string) *ListSensitiveDataAuditLogRequest
	GetColumnName() *string
	SetDbName(v string) *ListSensitiveDataAuditLogRequest
	GetDbName() *string
	SetEndTime(v string) *ListSensitiveDataAuditLogRequest
	GetEndTime() *string
	SetModuleName(v string) *ListSensitiveDataAuditLogRequest
	GetModuleName() *string
	SetOpUserName(v string) *ListSensitiveDataAuditLogRequest
	GetOpUserName() *string
	SetPageNumber(v int32) *ListSensitiveDataAuditLogRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListSensitiveDataAuditLogRequest
	GetPageSize() *int32
	SetStartTime(v string) *ListSensitiveDataAuditLogRequest
	GetStartTime() *string
	SetTableName(v string) *ListSensitiveDataAuditLogRequest
	GetTableName() *string
	SetTid(v int64) *ListSensitiveDataAuditLogRequest
	GetTid() *int64
}

type ListSensitiveDataAuditLogRequest struct {
	// The name of the column that contains sensitive data.
	//
	// example:
	//
	// ExampleColumnName
	ColumnName *string `json:"ColumnName,omitempty" xml:"ColumnName,omitempty"`
	// The name of the database that stores the sensitive data.
	//
	// example:
	//
	// ExampleDbName
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// The end of the time range for which you want to query the audit logs for sensitive information. Specify the time in the yyyy-MM-DD HH:mm:ss format.
	//
	// example:
	//
	// 2022-11-18 11:00:00
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The function module whose audit logs you want to query for sensitive data. If you do not specify this parameter, all audit logs are queried. Valid values:
	//
	// 	- **SQL_CONSOLE**: data query
	//
	// 	- **SQL_CONSOLE_EXPORT**: query result export
	//
	// 	- **DATA_CHANGE**: data change
	//
	// 	- **DATA_EXPORT**: data export
	//
	// example:
	//
	// SQL_CONSOLE
	ModuleName *string `json:"ModuleName,omitempty" xml:"ModuleName,omitempty"`
	// The username of the requester.
	//
	// example:
	//
	// ExampleOpUserName
	OpUserName *string `json:"OpUserName,omitempty" xml:"OpUserName,omitempty"`
	// The number of the page to return.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Example: 100
	//
	// example:
	//
	// 100
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The beginning of the time range for which you want to query the audit logs for sensitive information. Specify the time in the yyyy-MM-DD HH:mm:ss format.
	//
	// example:
	//
	// 2022-11-18 10:00:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The name of the table that stores the sensitive data.
	//
	// example:
	//
	// ExampleTableName
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	// The ID of the tenant.
	//
	// example:
	//
	// 3***
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s ListSensitiveDataAuditLogRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSensitiveDataAuditLogRequest) GoString() string {
	return s.String()
}

func (s *ListSensitiveDataAuditLogRequest) GetColumnName() *string {
	return s.ColumnName
}

func (s *ListSensitiveDataAuditLogRequest) GetDbName() *string {
	return s.DbName
}

func (s *ListSensitiveDataAuditLogRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *ListSensitiveDataAuditLogRequest) GetModuleName() *string {
	return s.ModuleName
}

func (s *ListSensitiveDataAuditLogRequest) GetOpUserName() *string {
	return s.OpUserName
}

func (s *ListSensitiveDataAuditLogRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListSensitiveDataAuditLogRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListSensitiveDataAuditLogRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *ListSensitiveDataAuditLogRequest) GetTableName() *string {
	return s.TableName
}

func (s *ListSensitiveDataAuditLogRequest) GetTid() *int64 {
	return s.Tid
}

func (s *ListSensitiveDataAuditLogRequest) SetColumnName(v string) *ListSensitiveDataAuditLogRequest {
	s.ColumnName = &v
	return s
}

func (s *ListSensitiveDataAuditLogRequest) SetDbName(v string) *ListSensitiveDataAuditLogRequest {
	s.DbName = &v
	return s
}

func (s *ListSensitiveDataAuditLogRequest) SetEndTime(v string) *ListSensitiveDataAuditLogRequest {
	s.EndTime = &v
	return s
}

func (s *ListSensitiveDataAuditLogRequest) SetModuleName(v string) *ListSensitiveDataAuditLogRequest {
	s.ModuleName = &v
	return s
}

func (s *ListSensitiveDataAuditLogRequest) SetOpUserName(v string) *ListSensitiveDataAuditLogRequest {
	s.OpUserName = &v
	return s
}

func (s *ListSensitiveDataAuditLogRequest) SetPageNumber(v int32) *ListSensitiveDataAuditLogRequest {
	s.PageNumber = &v
	return s
}

func (s *ListSensitiveDataAuditLogRequest) SetPageSize(v int32) *ListSensitiveDataAuditLogRequest {
	s.PageSize = &v
	return s
}

func (s *ListSensitiveDataAuditLogRequest) SetStartTime(v string) *ListSensitiveDataAuditLogRequest {
	s.StartTime = &v
	return s
}

func (s *ListSensitiveDataAuditLogRequest) SetTableName(v string) *ListSensitiveDataAuditLogRequest {
	s.TableName = &v
	return s
}

func (s *ListSensitiveDataAuditLogRequest) SetTid(v int64) *ListSensitiveDataAuditLogRequest {
	s.Tid = &v
	return s
}

func (s *ListSensitiveDataAuditLogRequest) Validate() error {
	return dara.Validate(s)
}
