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
	// Specifies whether to sort the results in ascending order. By default, the results are not sorted in ascending order.
	//
	// example:
	//
	// Disabled
	Asc *bool `json:"Asc,omitempty" xml:"Asc,omitempty"`
	// The name of the database.
	//
	// example:
	//
	// dbtest01
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// The end of the time range to query. Set this parameter to a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// >  The end time must be later than the start time. The interval cannot exceed one day.
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
	// The keywords that are used for query.
	//
	// example:
	//
	// dbtest01
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// The node ID.
	//
	// >  You must specify the node ID if your database instance is a PolarDB for MySQL cluster.
	//
	// example:
	//
	// pi-bp12v7243x012****
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The field by which to sort the returned entries. Default value: **count**. Valid values:
	//
	// 	- **count**: the number of executions.
	//
	// 	- **avgRt**: the average execution duration.
	//
	// 	- **rtRate**: the execution duration percentage.
	//
	// 	- **rowsExamined**: the total number of scanned rows.
	//
	// 	- **avgRowsExamined**: the average number of scanned rows.
	//
	// 	- **avgRowsReturned**: the average number of returned rows.
	//
	// example:
	//
	// count
	OrderBy *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	// The IP address of the client that executes the SQL statement.
	//
	// >  This parameter is optional. If this parameter is specified, the full request statistics of the specified IP address are collected. If this parameter is left empty, the full request statistics of the entire database instance are collected.
	//
	// example:
	//
	// 172.26.XX.XXX
	OriginHost *string `json:"OriginHost,omitempty" xml:"OriginHost,omitempty"`
	// The page number. Pages start from page 1. Default value: 1.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page. Default value: 20.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The role of the node in the PolarDB-X 2.0 instance. Valid values:
	//
	// 	- **polarx_cn**: compute node.
	//
	// 	- **polarx_dn**: data node.
	//
	// example:
	//
	// polarx_cn
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// The SQL ID.
	//
	// >  If this parameter is specified, the full request statistics of the specified SQL query are collected. If this parameter is left empty, the full request statistics of the entire database instance are collected.
	//
	// example:
	//
	// d71f82be1eef72bd105128204d2e****
	SqlId *string `json:"SqlId,omitempty" xml:"SqlId,omitempty"`
	// The type of the SQL statement. Valid values: **SELECT**, **INSERT**, **UPDATE**, **DELETE**, **LOGIN**, **LOGOUT**, **MERGE**, **ALTER**, **CREATEINDEX**, **DROPINDEX**, **CREATE**, **DROP**, **SET**, **DESC**, **REPLACE**, **CALL**, **BEGIN**, **DESCRIBE**, **ROLLBACK**, **FLUSH**, **USE**, **SHOW**, **START**, **COMMIT**, and **RENAME**.
	//
	// >  If your database instance is an ApsaraDB RDS for MySQL instance, a PolarDB for MySQL cluster, or a PolarDB-X 2.0 instance, the statistics can be collected based on the SQL statement type.
	//
	// example:
	//
	// SELECT
	SqlType *string `json:"SqlType,omitempty" xml:"SqlType,omitempty"`
	// The beginning of the time range to query. Set this parameter to a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// >  The start time can be up to 90 days earlier than the current time.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1645581813000
	Start *int64 `json:"Start,omitempty" xml:"Start,omitempty"`
	// The ID of the Alibaba Cloud account that was used to create the database instance.
	//
	// >  This parameter is optional. The system can automatically obtain the Alibaba Cloud account ID based on the value of InstanceId when you call the GetFullRequestOriginStatByInstanceId operation.
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
