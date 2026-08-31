// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryColumnarLogSSERequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceName(v string) *QueryColumnarLogSSERequest
	GetDBInstanceName() *string
	SetMaxResultRows(v int64) *QueryColumnarLogSSERequest
	GetMaxResultRows() *int64
	SetRegionId(v string) *QueryColumnarLogSSERequest
	GetRegionId() *string
	SetSQL(v string) *QueryColumnarLogSSERequest
	GetSQL() *string
}

type QueryColumnarLogSSERequest struct {
	// The ID of the PolarDB-X instance whose column store audit logs you want to query.
	//
	// This parameter is required.
	//
	// example:
	//
	// pxc-********
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// The maximum number of result rows to return for this request. Valid values: 1 to 10000. If this parameter is not specified, no additional row limit is imposed on the SQL submitted by the caller. The server-side SSE upper limit of 10000 rows and the top-level LIMIT clause in the SQL statement still apply.
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
	// The read-only query statement to execute. Only a single MySQL SELECT statement is supported, and it must access the fully qualified polardbx_sls table. Multi-statement queries, write operations, locks, user variables, dynamic placeholders, and reserved hints are not supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// select 	- from device where name = \\"108001022203365239\\"
	SQL *string `json:"SQL,omitempty" xml:"SQL,omitempty"`
}

func (s QueryColumnarLogSSERequest) String() string {
	return dara.Prettify(s)
}

func (s QueryColumnarLogSSERequest) GoString() string {
	return s.String()
}

func (s *QueryColumnarLogSSERequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *QueryColumnarLogSSERequest) GetMaxResultRows() *int64 {
	return s.MaxResultRows
}

func (s *QueryColumnarLogSSERequest) GetRegionId() *string {
	return s.RegionId
}

func (s *QueryColumnarLogSSERequest) GetSQL() *string {
	return s.SQL
}

func (s *QueryColumnarLogSSERequest) SetDBInstanceName(v string) *QueryColumnarLogSSERequest {
	s.DBInstanceName = &v
	return s
}

func (s *QueryColumnarLogSSERequest) SetMaxResultRows(v int64) *QueryColumnarLogSSERequest {
	s.MaxResultRows = &v
	return s
}

func (s *QueryColumnarLogSSERequest) SetRegionId(v string) *QueryColumnarLogSSERequest {
	s.RegionId = &v
	return s
}

func (s *QueryColumnarLogSSERequest) SetSQL(v string) *QueryColumnarLogSSERequest {
	s.SQL = &v
	return s
}

func (s *QueryColumnarLogSSERequest) Validate() error {
	return dara.Validate(s)
}
