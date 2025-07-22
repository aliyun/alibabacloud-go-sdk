// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFullRequestOriginStatByInstanceIdRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAsc(v bool) *GetFullRequestOriginStatByInstanceIdRequest
	GetAsc() *bool
	SetEnd(v int64) *GetFullRequestOriginStatByInstanceIdRequest
	GetEnd() *int64
	SetInstanceId(v string) *GetFullRequestOriginStatByInstanceIdRequest
	GetInstanceId() *string
	SetNodeId(v string) *GetFullRequestOriginStatByInstanceIdRequest
	GetNodeId() *string
	SetOrderBy(v string) *GetFullRequestOriginStatByInstanceIdRequest
	GetOrderBy() *string
	SetPageNo(v int32) *GetFullRequestOriginStatByInstanceIdRequest
	GetPageNo() *int32
	SetPageSize(v int32) *GetFullRequestOriginStatByInstanceIdRequest
	GetPageSize() *int32
	SetRole(v string) *GetFullRequestOriginStatByInstanceIdRequest
	GetRole() *string
	SetSqlType(v string) *GetFullRequestOriginStatByInstanceIdRequest
	GetSqlType() *string
	SetStart(v int64) *GetFullRequestOriginStatByInstanceIdRequest
	GetStart() *int64
	SetUserId(v string) *GetFullRequestOriginStatByInstanceIdRequest
	GetUserId() *string
}

type GetFullRequestOriginStatByInstanceIdRequest struct {
	// Specifies whether to sort the results in ascending order. By default, the results are not sorted in ascending order.
	//
	// example:
	//
	// Disabled
	Asc *bool `json:"Asc,omitempty" xml:"Asc,omitempty"`
	// The end of the time range to query. Set this parameter to a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// >  The end time must be later than the start time. The interval between the start time and the end time cannot exceed 24 hours.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1644803409000
	End *int64 `json:"End,omitempty" xml:"End,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-2ze8g2am97624****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The node ID.
	//
	// >  This parameter must be specified if the database instance is a PolarDB for MySQL cluster.
	//
	// example:
	//
	// pi-bp12v7243x012****
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The field by which the results to be returned are sorted. Default value: **count**. Valid values:
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
	// The role of the PolarDB-X 2.0 node. Valid values:
	//
	// 	- **polarx_cn**: compute node.
	//
	// 	- **polarx_en**: data node.
	//
	// example:
	//
	// polarx_cn
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// The type of the SQL statement. Valid values: **SELECT**, **INSERT**, **UPDATE**, **DELETE**, **MERGE**, **ALTER**, **CREATEINDEX**, **DROPINDEX**, **CREATE**, **DROP**, **SET**, **DESC**, **REPLACE**, **CALL**, **BEGIN**, **DESCRIBE**, **ROLLBACK**, **FLUSH**, **USE**, **SHOW**, **START**, **COMMIT**, and **RENAME**.
	//
	// >  If the database instance is an ApsaraDB RDS for MySQL instance, a PolarDB for MySQL instance, or a PolarDB-X 2.0 instance, statistics can be collected based on the SQL statement type.
	//
	// example:
	//
	// SELECT
	SqlType *string `json:"SqlType,omitempty" xml:"SqlType,omitempty"`
	// The beginning of the time range to query. Set this parameter to a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// >  The start time must be within the storage duration of the SQL Explorer of the database instance, and can be up to 90 days earlier than the current time.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1644716649000
	Start *int64 `json:"Start,omitempty" xml:"Start,omitempty"`
	// The ID of the Alibaba Cloud account that is used to create the database instance.
	//
	// >  This parameter is optional. The system can automatically obtain the account ID based on the value of InstanceId when you call this operation.
	//
	// example:
	//
	// 196278346919****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s GetFullRequestOriginStatByInstanceIdRequest) String() string {
	return dara.Prettify(s)
}

func (s GetFullRequestOriginStatByInstanceIdRequest) GoString() string {
	return s.String()
}

func (s *GetFullRequestOriginStatByInstanceIdRequest) GetAsc() *bool {
	return s.Asc
}

func (s *GetFullRequestOriginStatByInstanceIdRequest) GetEnd() *int64 {
	return s.End
}

func (s *GetFullRequestOriginStatByInstanceIdRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetFullRequestOriginStatByInstanceIdRequest) GetNodeId() *string {
	return s.NodeId
}

func (s *GetFullRequestOriginStatByInstanceIdRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *GetFullRequestOriginStatByInstanceIdRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *GetFullRequestOriginStatByInstanceIdRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetFullRequestOriginStatByInstanceIdRequest) GetRole() *string {
	return s.Role
}

func (s *GetFullRequestOriginStatByInstanceIdRequest) GetSqlType() *string {
	return s.SqlType
}

func (s *GetFullRequestOriginStatByInstanceIdRequest) GetStart() *int64 {
	return s.Start
}

func (s *GetFullRequestOriginStatByInstanceIdRequest) GetUserId() *string {
	return s.UserId
}

func (s *GetFullRequestOriginStatByInstanceIdRequest) SetAsc(v bool) *GetFullRequestOriginStatByInstanceIdRequest {
	s.Asc = &v
	return s
}

func (s *GetFullRequestOriginStatByInstanceIdRequest) SetEnd(v int64) *GetFullRequestOriginStatByInstanceIdRequest {
	s.End = &v
	return s
}

func (s *GetFullRequestOriginStatByInstanceIdRequest) SetInstanceId(v string) *GetFullRequestOriginStatByInstanceIdRequest {
	s.InstanceId = &v
	return s
}

func (s *GetFullRequestOriginStatByInstanceIdRequest) SetNodeId(v string) *GetFullRequestOriginStatByInstanceIdRequest {
	s.NodeId = &v
	return s
}

func (s *GetFullRequestOriginStatByInstanceIdRequest) SetOrderBy(v string) *GetFullRequestOriginStatByInstanceIdRequest {
	s.OrderBy = &v
	return s
}

func (s *GetFullRequestOriginStatByInstanceIdRequest) SetPageNo(v int32) *GetFullRequestOriginStatByInstanceIdRequest {
	s.PageNo = &v
	return s
}

func (s *GetFullRequestOriginStatByInstanceIdRequest) SetPageSize(v int32) *GetFullRequestOriginStatByInstanceIdRequest {
	s.PageSize = &v
	return s
}

func (s *GetFullRequestOriginStatByInstanceIdRequest) SetRole(v string) *GetFullRequestOriginStatByInstanceIdRequest {
	s.Role = &v
	return s
}

func (s *GetFullRequestOriginStatByInstanceIdRequest) SetSqlType(v string) *GetFullRequestOriginStatByInstanceIdRequest {
	s.SqlType = &v
	return s
}

func (s *GetFullRequestOriginStatByInstanceIdRequest) SetStart(v int64) *GetFullRequestOriginStatByInstanceIdRequest {
	s.Start = &v
	return s
}

func (s *GetFullRequestOriginStatByInstanceIdRequest) SetUserId(v string) *GetFullRequestOriginStatByInstanceIdRequest {
	s.UserId = &v
	return s
}

func (s *GetFullRequestOriginStatByInstanceIdRequest) Validate() error {
	return dara.Validate(s)
}
