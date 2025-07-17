// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstanceLoginAuditLogRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *ListInstanceLoginAuditLogRequest
	GetEndTime() *string
	SetOpUserName(v string) *ListInstanceLoginAuditLogRequest
	GetOpUserName() *string
	SetPageNumber(v int32) *ListInstanceLoginAuditLogRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListInstanceLoginAuditLogRequest
	GetPageSize() *int32
	SetSearchName(v string) *ListInstanceLoginAuditLogRequest
	GetSearchName() *string
	SetStartTime(v string) *ListInstanceLoginAuditLogRequest
	GetStartTime() *string
	SetTid(v int64) *ListInstanceLoginAuditLogRequest
	GetTid() *int64
}

type ListInstanceLoginAuditLogRequest struct {
	// The end of the time range to query.
	//
	// >  The end time supports fuzzy match. Specify the time in the YYYY-MM-DD hh:mm:ss format. We recommend that you use the StartTime and EndTime parameters to specify a time range that does not exceed one day. This way, the returned entries can be displayed by page to increase query efficiency.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2021-11-18 18:00:00
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The alias of the user.
	//
	// example:
	//
	// test_OpUserName
	OpUserName *string `json:"OpUserName,omitempty" xml:"OpUserName,omitempty"`
	// The number of the page to return.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Maximum value: 100.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The name of the database or instance whose logon records you want to query.
	//
	// >  If SQL statements are executed at the instance level, you can set this parameter to an instance name. If SQL statements are executed at the database level, you can set this parameter to a database name.
	//
	// example:
	//
	// test_SearchName
	SearchName *string `json:"SearchName,omitempty" xml:"SearchName,omitempty"`
	// The beginning of the time range to query.
	//
	// >  The start time supports fuzzy match. Specify the time in the YYYY-MM-DD hh:mm:ss format.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2021-11-18 11:00:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The ID of the tenant. You can call the [GetUserActiveTenant](https://help.aliyun.com/document_detail/198073.html) operation to query the tenant ID.
	//
	// example:
	//
	// 3***
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s ListInstanceLoginAuditLogRequest) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceLoginAuditLogRequest) GoString() string {
	return s.String()
}

func (s *ListInstanceLoginAuditLogRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *ListInstanceLoginAuditLogRequest) GetOpUserName() *string {
	return s.OpUserName
}

func (s *ListInstanceLoginAuditLogRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListInstanceLoginAuditLogRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListInstanceLoginAuditLogRequest) GetSearchName() *string {
	return s.SearchName
}

func (s *ListInstanceLoginAuditLogRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *ListInstanceLoginAuditLogRequest) GetTid() *int64 {
	return s.Tid
}

func (s *ListInstanceLoginAuditLogRequest) SetEndTime(v string) *ListInstanceLoginAuditLogRequest {
	s.EndTime = &v
	return s
}

func (s *ListInstanceLoginAuditLogRequest) SetOpUserName(v string) *ListInstanceLoginAuditLogRequest {
	s.OpUserName = &v
	return s
}

func (s *ListInstanceLoginAuditLogRequest) SetPageNumber(v int32) *ListInstanceLoginAuditLogRequest {
	s.PageNumber = &v
	return s
}

func (s *ListInstanceLoginAuditLogRequest) SetPageSize(v int32) *ListInstanceLoginAuditLogRequest {
	s.PageSize = &v
	return s
}

func (s *ListInstanceLoginAuditLogRequest) SetSearchName(v string) *ListInstanceLoginAuditLogRequest {
	s.SearchName = &v
	return s
}

func (s *ListInstanceLoginAuditLogRequest) SetStartTime(v string) *ListInstanceLoginAuditLogRequest {
	s.StartTime = &v
	return s
}

func (s *ListInstanceLoginAuditLogRequest) SetTid(v int64) *ListInstanceLoginAuditLogRequest {
	s.Tid = &v
	return s
}

func (s *ListInstanceLoginAuditLogRequest) Validate() error {
	return dara.Validate(s)
}
