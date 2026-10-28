// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFullRequestStatResultByInstanceIdRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAsc(v bool) *GetFullRequestStatResultByInstanceIdRequest
	GetAsc() *bool
	SetDbName(v string) *GetFullRequestStatResultByInstanceIdRequest
	GetDbName() *string
	SetEnd(v int64) *GetFullRequestStatResultByInstanceIdRequest
	GetEnd() *int64
	SetInstanceId(v string) *GetFullRequestStatResultByInstanceIdRequest
	GetInstanceId() *string
	SetKeyword(v string) *GetFullRequestStatResultByInstanceIdRequest
	GetKeyword() *string
	SetNodeId(v string) *GetFullRequestStatResultByInstanceIdRequest
	GetNodeId() *string
	SetOrderBy(v string) *GetFullRequestStatResultByInstanceIdRequest
	GetOrderBy() *string
	SetOriginHost(v string) *GetFullRequestStatResultByInstanceIdRequest
	GetOriginHost() *string
	SetPageNo(v int32) *GetFullRequestStatResultByInstanceIdRequest
	GetPageNo() *int32
	SetPageSize(v int32) *GetFullRequestStatResultByInstanceIdRequest
	GetPageSize() *int32
	SetRole(v string) *GetFullRequestStatResultByInstanceIdRequest
	GetRole() *string
	SetSqlId(v string) *GetFullRequestStatResultByInstanceIdRequest
	GetSqlId() *string
	SetSqlType(v string) *GetFullRequestStatResultByInstanceIdRequest
	GetSqlType() *string
	SetStart(v int64) *GetFullRequestStatResultByInstanceIdRequest
	GetStart() *int64
	SetUserId(v string) *GetFullRequestStatResultByInstanceIdRequest
	GetUserId() *string
}

type GetFullRequestStatResultByInstanceIdRequest struct {
	// Specifies whether to sort the results in ascending order. Default value: false.
	//
	// example:
	//
	// False
	Asc *bool `json:"Asc,omitempty" xml:"Asc,omitempty"`
	// The name of the database.
	//
	// example:
	//
	// dbtest01
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// The end time of the query. The time is a UNIX timestamp in milliseconds.
	//
	// > The end time must be later than the start time, and the interval between the start time and the end time cannot exceed one day.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1645668213000
	End *int64 `json:"End,omitempty" xml:"End,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-2ze8g2am97624****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The query keyword.
	//
	// example:
	//
	// dbtest01
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// The node ID.
	//
	// > For PolarDB for MySQL instances, the node ID is required.
	//
	// example:
	//
	// pi-bp12v7243x012****
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The field by which to sort the results. Default value: **count**.
	//
	// - **count**: the number of executions.
	//
	// - **avgRt**: the average execution time.
	//
	// - **rtRate**: the execution time proportion.
	//
	// - **rowsExamined**: the total number of scanned rows.
	//
	// - **avgRowsExamined**: the average number of scanned rows.
	//
	// - **avgRowsReturned**: the average number of returned rows.
	//
	// example:
	//
	// count
	OrderBy *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	// The IP address of the client that executed the SQL statement.
	//
	// > This parameter is optional. If you specify a client IP address, the full request data of the specified client IP address is aggregated. If you leave this parameter empty, the full request data of the entire database instance is aggregated.
	//
	// example:
	//
	// 172.26.XX.XXX
	OriginHost *string `json:"OriginHost,omitempty" xml:"OriginHost,omitempty"`
	// The page number for the paged query. The paging starts from 1. Default value: 1.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The data volume per page for paging in a paged query. Default value: 20.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The node information of the PolarDB-X 2.0 database instance. Valid values:
	//
	// - **polarx_cn**: compute node.
	//
	// - **polarx_dn**: data node.
	//
	// example:
	//
	// polarx_cn
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// The SQL ID.
	//
	// > If you specify an SQL ID, the full request data of the specified SQL ID is aggregated. If you leave this parameter empty, the full request data of the entire database instance is aggregated.
	//
	// example:
	//
	// d71f82be1eef72bd105128204d2e****
	SqlId *string `json:"SqlId,omitempty" xml:"SqlId,omitempty"`
	// The type of the SQL statement. Valid values: **SELECT**, **INSERT**, **UPDATE**, **DELETE**, **LOGIN**, **LOGOUT**, **MERGE**, **ALTER**, **CREATEINDEX**, **DROPINDEX**, **CREATE**, **DROP**, **SET**, **DESC**, **REPLACE**, **CALL**, **BEGIN**, **DESCRIBE**, **ROLLBACK**, **FLUSH**, **USE**, **SHOW**, **START**, **COMMIT**, and **RENAME**.
	//
	// > If the database instance is an ApsaraDB RDS for MySQL, PolarDB for MySQL, or PolarDB-X 2.0 instance, you can aggregate data by SQL statement type.
	//
	// example:
	//
	// SELECT
	SqlType *string `json:"SqlType,omitempty" xml:"SqlType,omitempty"`
	// The start time of the query. The format is a UNIX timestamp, in milliseconds.
	//
	// > The start time must be within the SQL Explorer storage duration for the database instance and cannot be more than 90 days earlier than the current time.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1645581813000
	Start *int64 `json:"Start,omitempty" xml:"Start,omitempty"`
	// The instance ID, which is the ID of the Alibaba Cloud account that created the database instance.
	//
	// > This parameter is optional. After you enter the ID of the target database instance, the system automatically obtains the instance ID when you call the API operation.
	//
	// example:
	//
	// 196278346919****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s GetFullRequestStatResultByInstanceIdRequest) String() string {
	return dara.Prettify(s)
}

func (s GetFullRequestStatResultByInstanceIdRequest) GoString() string {
	return s.String()
}

func (s *GetFullRequestStatResultByInstanceIdRequest) GetAsc() *bool {
	return s.Asc
}

func (s *GetFullRequestStatResultByInstanceIdRequest) GetDbName() *string {
	return s.DbName
}

func (s *GetFullRequestStatResultByInstanceIdRequest) GetEnd() *int64 {
	return s.End
}

func (s *GetFullRequestStatResultByInstanceIdRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetFullRequestStatResultByInstanceIdRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *GetFullRequestStatResultByInstanceIdRequest) GetNodeId() *string {
	return s.NodeId
}

func (s *GetFullRequestStatResultByInstanceIdRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *GetFullRequestStatResultByInstanceIdRequest) GetOriginHost() *string {
	return s.OriginHost
}

func (s *GetFullRequestStatResultByInstanceIdRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *GetFullRequestStatResultByInstanceIdRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetFullRequestStatResultByInstanceIdRequest) GetRole() *string {
	return s.Role
}

func (s *GetFullRequestStatResultByInstanceIdRequest) GetSqlId() *string {
	return s.SqlId
}

func (s *GetFullRequestStatResultByInstanceIdRequest) GetSqlType() *string {
	return s.SqlType
}

func (s *GetFullRequestStatResultByInstanceIdRequest) GetStart() *int64 {
	return s.Start
}

func (s *GetFullRequestStatResultByInstanceIdRequest) GetUserId() *string {
	return s.UserId
}

func (s *GetFullRequestStatResultByInstanceIdRequest) SetAsc(v bool) *GetFullRequestStatResultByInstanceIdRequest {
	s.Asc = &v
	return s
}

func (s *GetFullRequestStatResultByInstanceIdRequest) SetDbName(v string) *GetFullRequestStatResultByInstanceIdRequest {
	s.DbName = &v
	return s
}

func (s *GetFullRequestStatResultByInstanceIdRequest) SetEnd(v int64) *GetFullRequestStatResultByInstanceIdRequest {
	s.End = &v
	return s
}

func (s *GetFullRequestStatResultByInstanceIdRequest) SetInstanceId(v string) *GetFullRequestStatResultByInstanceIdRequest {
	s.InstanceId = &v
	return s
}

func (s *GetFullRequestStatResultByInstanceIdRequest) SetKeyword(v string) *GetFullRequestStatResultByInstanceIdRequest {
	s.Keyword = &v
	return s
}

func (s *GetFullRequestStatResultByInstanceIdRequest) SetNodeId(v string) *GetFullRequestStatResultByInstanceIdRequest {
	s.NodeId = &v
	return s
}

func (s *GetFullRequestStatResultByInstanceIdRequest) SetOrderBy(v string) *GetFullRequestStatResultByInstanceIdRequest {
	s.OrderBy = &v
	return s
}

func (s *GetFullRequestStatResultByInstanceIdRequest) SetOriginHost(v string) *GetFullRequestStatResultByInstanceIdRequest {
	s.OriginHost = &v
	return s
}

func (s *GetFullRequestStatResultByInstanceIdRequest) SetPageNo(v int32) *GetFullRequestStatResultByInstanceIdRequest {
	s.PageNo = &v
	return s
}

func (s *GetFullRequestStatResultByInstanceIdRequest) SetPageSize(v int32) *GetFullRequestStatResultByInstanceIdRequest {
	s.PageSize = &v
	return s
}

func (s *GetFullRequestStatResultByInstanceIdRequest) SetRole(v string) *GetFullRequestStatResultByInstanceIdRequest {
	s.Role = &v
	return s
}

func (s *GetFullRequestStatResultByInstanceIdRequest) SetSqlId(v string) *GetFullRequestStatResultByInstanceIdRequest {
	s.SqlId = &v
	return s
}

func (s *GetFullRequestStatResultByInstanceIdRequest) SetSqlType(v string) *GetFullRequestStatResultByInstanceIdRequest {
	s.SqlType = &v
	return s
}

func (s *GetFullRequestStatResultByInstanceIdRequest) SetStart(v int64) *GetFullRequestStatResultByInstanceIdRequest {
	s.Start = &v
	return s
}

func (s *GetFullRequestStatResultByInstanceIdRequest) SetUserId(v string) *GetFullRequestStatResultByInstanceIdRequest {
	s.UserId = &v
	return s
}

func (s *GetFullRequestStatResultByInstanceIdRequest) Validate() error {
	return dara.Validate(s)
}
