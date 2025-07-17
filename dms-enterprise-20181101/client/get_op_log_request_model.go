// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOpLogRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDatabaseName(v string) *GetOpLogRequest
	GetDatabaseName() *string
	SetEndTime(v string) *GetOpLogRequest
	GetEndTime() *string
	SetModule(v string) *GetOpLogRequest
	GetModule() *string
	SetPageNumber(v int32) *GetOpLogRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *GetOpLogRequest
	GetPageSize() *int32
	SetStartTime(v string) *GetOpLogRequest
	GetStartTime() *string
	SetTid(v int64) *GetOpLogRequest
	GetTid() *int64
	SetUserNick(v string) *GetOpLogRequest
	GetUserNick() *string
}

type GetOpLogRequest struct {
	// DatabaseName.
	//
	// example:
	//
	// dmstest@rm-bp1qb97d4b****.mysql.rds.aliyuncs.com:3306[poc_dev]
	DatabaseName *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	// The end of the time range to query. Specify the time in the yyyy-MM-DD HH:mm:ss format.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2022-03-29 10:00:00
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The functional module for which you want to query operation logs. If you do not specify this parameter, operation logs for all functional modules are returned. Valid values:
	//
	// 	- **PERMISSION**: permissions
	//
	// 	- **OWNER**: data owner
	//
	// 	- **SQL_CONSOLE**: data query
	//
	// 	- **SQL_CONSOLE_EXPORT**: query result export
	//
	// 	- **DATA_CHANGE**: data change
	//
	// 	- **DATA_EXPORT**: data export
	//
	// 	- **SQL_REVIEW**: SQL review
	//
	// 	- **DT_SYNC**: database and table synchronization
	//
	// 	- **DT_DETAIL**: database and table details
	//
	// 	- **DB_TASK**: task management
	//
	// 	- **INSTANCE_MANAGE**: instance management
	//
	// 	- **USER_MANAGE**: user management
	//
	// 	- **SECURITY_RULE**: security rules
	//
	// 	- **CONFIG_MANAGE**: configuration management
	//
	// 	- **RESOURCE_AUTH**: resource authorization
	//
	// 	- **ACCESS_WHITE_IP**: access IP address whitelist
	//
	// 	- **NDDL**: schema design
	//
	// 	- **DSQL_CONSOLE**: cross-database data query
	//
	// 	- **DSQL_CONSOLE_EXPORT**: cross-database query result export
	//
	// 	- **DATA_TRACT**: data tracking
	//
	// 	- **DATA_QUALITY**: data quality
	//
	// 	- **DATALINK_MANAGE*	- :DBLink management
	//
	// 	- **DATASEC_MANAGE**: sensitive data management
	//
	// 	- **SELL**: sales
	//
	// example:
	//
	// SECURITY_RULE
	Module *string `json:"Module,omitempty" xml:"Module,omitempty"`
	// The number of the page to return. Pages start from page 1.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Valid values:
	//
	// 	- 30
	//
	// 	- 50
	//
	// 	- 100
	//
	// This parameter is required.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The beginning of the time range to query. Specify the time in the yyyy-MM-DD HH:mm:ss format.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2022-03-23 10:00:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The ID of the tenant. You can call the [GetUserActiveTenant](https://help.aliyun.com/document_detail/198073.html) or [ListUserTenants](https://help.aliyun.com/document_detail/198074.html) operation to query the tenant ID.
	//
	// example:
	//
	// 3***
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
	// UserNick.
	//
	// example:
	//
	// test_name
	UserNick *string `json:"UserNick,omitempty" xml:"UserNick,omitempty"`
}

func (s GetOpLogRequest) String() string {
	return dara.Prettify(s)
}

func (s GetOpLogRequest) GoString() string {
	return s.String()
}

func (s *GetOpLogRequest) GetDatabaseName() *string {
	return s.DatabaseName
}

func (s *GetOpLogRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *GetOpLogRequest) GetModule() *string {
	return s.Module
}

func (s *GetOpLogRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *GetOpLogRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetOpLogRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *GetOpLogRequest) GetTid() *int64 {
	return s.Tid
}

func (s *GetOpLogRequest) GetUserNick() *string {
	return s.UserNick
}

func (s *GetOpLogRequest) SetDatabaseName(v string) *GetOpLogRequest {
	s.DatabaseName = &v
	return s
}

func (s *GetOpLogRequest) SetEndTime(v string) *GetOpLogRequest {
	s.EndTime = &v
	return s
}

func (s *GetOpLogRequest) SetModule(v string) *GetOpLogRequest {
	s.Module = &v
	return s
}

func (s *GetOpLogRequest) SetPageNumber(v int32) *GetOpLogRequest {
	s.PageNumber = &v
	return s
}

func (s *GetOpLogRequest) SetPageSize(v int32) *GetOpLogRequest {
	s.PageSize = &v
	return s
}

func (s *GetOpLogRequest) SetStartTime(v string) *GetOpLogRequest {
	s.StartTime = &v
	return s
}

func (s *GetOpLogRequest) SetTid(v int64) *GetOpLogRequest {
	s.Tid = &v
	return s
}

func (s *GetOpLogRequest) SetUserNick(v string) *GetOpLogRequest {
	s.UserNick = &v
	return s
}

func (s *GetOpLogRequest) Validate() error {
	return dara.Validate(s)
}
