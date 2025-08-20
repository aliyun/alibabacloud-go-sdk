// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAccountPrivilegeObjectsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*DescribeAccountPrivilegeObjectsResponseBodyData) *DescribeAccountPrivilegeObjectsResponseBody
	GetData() []*DescribeAccountPrivilegeObjectsResponseBodyData
	SetPageNumber(v int64) *DescribeAccountPrivilegeObjectsResponseBody
	GetPageNumber() *int64
	SetPageSize(v int64) *DescribeAccountPrivilegeObjectsResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *DescribeAccountPrivilegeObjectsResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *DescribeAccountPrivilegeObjectsResponseBody
	GetTotalCount() *int64
}

type DescribeAccountPrivilegeObjectsResponseBody struct {
	// The permissions.
	Data []*DescribeAccountPrivilegeObjectsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 34B2AD29-682F-1C14-B3AA-9EF1A96084B8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 23
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeAccountPrivilegeObjectsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAccountPrivilegeObjectsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAccountPrivilegeObjectsResponseBody) GetData() []*DescribeAccountPrivilegeObjectsResponseBodyData {
	return s.Data
}

func (s *DescribeAccountPrivilegeObjectsResponseBody) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeAccountPrivilegeObjectsResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeAccountPrivilegeObjectsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAccountPrivilegeObjectsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeAccountPrivilegeObjectsResponseBody) SetData(v []*DescribeAccountPrivilegeObjectsResponseBodyData) *DescribeAccountPrivilegeObjectsResponseBody {
	s.Data = v
	return s
}

func (s *DescribeAccountPrivilegeObjectsResponseBody) SetPageNumber(v int64) *DescribeAccountPrivilegeObjectsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeAccountPrivilegeObjectsResponseBody) SetPageSize(v int64) *DescribeAccountPrivilegeObjectsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeAccountPrivilegeObjectsResponseBody) SetRequestId(v string) *DescribeAccountPrivilegeObjectsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAccountPrivilegeObjectsResponseBody) SetTotalCount(v int64) *DescribeAccountPrivilegeObjectsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeAccountPrivilegeObjectsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeAccountPrivilegeObjectsResponseBodyData struct {
	// The name of the column. This parameter is returned when PrivilegeType is set to Column.
	//
	// example:
	//
	// column1
	Column *string `json:"Column,omitempty" xml:"Column,omitempty"`
	// The name of the database. This parameter is returned when PrivilegeType is set to Database, Table, or Column.
	//
	// example:
	//
	// tdb1
	Database *string `json:"Database,omitempty" xml:"Database,omitempty"`
	// The description that is specified when you create a table or column. This parameter is returned only when PrivilegeType is set to Database or Table, indicating the database description or table description.
	//
	// example:
	//
	// a test db
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the table. This parameter is returned when PrivilegeType is set to Table or Column.
	//
	// example:
	//
	// table1
	Table *string `json:"Table,omitempty" xml:"Table,omitempty"`
}

func (s DescribeAccountPrivilegeObjectsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeAccountPrivilegeObjectsResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeAccountPrivilegeObjectsResponseBodyData) GetColumn() *string {
	return s.Column
}

func (s *DescribeAccountPrivilegeObjectsResponseBodyData) GetDatabase() *string {
	return s.Database
}

func (s *DescribeAccountPrivilegeObjectsResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *DescribeAccountPrivilegeObjectsResponseBodyData) GetTable() *string {
	return s.Table
}

func (s *DescribeAccountPrivilegeObjectsResponseBodyData) SetColumn(v string) *DescribeAccountPrivilegeObjectsResponseBodyData {
	s.Column = &v
	return s
}

func (s *DescribeAccountPrivilegeObjectsResponseBodyData) SetDatabase(v string) *DescribeAccountPrivilegeObjectsResponseBodyData {
	s.Database = &v
	return s
}

func (s *DescribeAccountPrivilegeObjectsResponseBodyData) SetDescription(v string) *DescribeAccountPrivilegeObjectsResponseBodyData {
	s.Description = &v
	return s
}

func (s *DescribeAccountPrivilegeObjectsResponseBodyData) SetTable(v string) *DescribeAccountPrivilegeObjectsResponseBodyData {
	s.Table = &v
	return s
}

func (s *DescribeAccountPrivilegeObjectsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
