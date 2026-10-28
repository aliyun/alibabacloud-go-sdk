// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryColumnarLogRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceName(v string) *QueryColumnarLogRequest
	GetDBInstanceName() *string
	SetMaxResultRows(v int64) *QueryColumnarLogRequest
	GetMaxResultRows() *int64
	SetRegionId(v string) *QueryColumnarLogRequest
	GetRegionId() *string
	SetSQL(v string) *QueryColumnarLogRequest
	GetSQL() *string
}

type QueryColumnarLogRequest struct {
	// The ID of the PolarDB-X instance for which you want to query column store audit logs.
	//
	// This parameter is required.
	//
	// example:
	//
	// pxc-********
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// The maximum number of result rows to return for this request. Valid values: 1 to 1000. Default value: 100. The actual number of returned rows is also subject to the top-level LIMIT clause in the SQL statement and the current service policy.
	//
	// example:
	//
	// 1000
	MaxResultRows *int64 `json:"MaxResultRows,omitempty" xml:"MaxResultRows,omitempty"`
	// The region ID of the request. The value must match the region where the SQLQuery service is deployed.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The read-only query statement to execute. Only a single MySQL SELECT statement is supported. Query tables must use the fully qualified form polardbx_sls.<table_name>, and only the following tables are allowed: polardbx_sls.polardbx_log (kernel logs), polardbx_sls.slow_detail_log (slow query logs), polardbx_sls.slow_trans_log (slow transaction logs), and polardbx_sls.sql_audit_log (audit logs). Multi-statement queries, write operations, locks, user variables, dynamic placeholders, and reserved hints are not supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// SELECT 	- FROM polardbx_sls.sql_audit_log LIMIT 100
	SQL *string `json:"SQL,omitempty" xml:"SQL,omitempty"`
}

func (s QueryColumnarLogRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryColumnarLogRequest) GoString() string {
	return s.String()
}

func (s *QueryColumnarLogRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *QueryColumnarLogRequest) GetMaxResultRows() *int64 {
	return s.MaxResultRows
}

func (s *QueryColumnarLogRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *QueryColumnarLogRequest) GetSQL() *string {
	return s.SQL
}

func (s *QueryColumnarLogRequest) SetDBInstanceName(v string) *QueryColumnarLogRequest {
	s.DBInstanceName = &v
	return s
}

func (s *QueryColumnarLogRequest) SetMaxResultRows(v int64) *QueryColumnarLogRequest {
	s.MaxResultRows = &v
	return s
}

func (s *QueryColumnarLogRequest) SetRegionId(v string) *QueryColumnarLogRequest {
	s.RegionId = &v
	return s
}

func (s *QueryColumnarLogRequest) SetSQL(v string) *QueryColumnarLogRequest {
	s.SQL = &v
	return s
}

func (s *QueryColumnarLogRequest) Validate() error {
	return dara.Validate(s)
}
