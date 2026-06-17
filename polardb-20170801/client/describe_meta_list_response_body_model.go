// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMetaListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeMetaListResponseBody
	GetDBClusterId() *string
	SetItems(v []*DescribeMetaListResponseBodyItems) *DescribeMetaListResponseBody
	GetItems() []*DescribeMetaListResponseBodyItems
	SetPageNumber(v string) *DescribeMetaListResponseBody
	GetPageNumber() *string
	SetPageSize(v string) *DescribeMetaListResponseBody
	GetPageSize() *string
	SetRequestId(v string) *DescribeMetaListResponseBody
	GetRequestId() *string
	SetTotalPageCount(v string) *DescribeMetaListResponseBody
	GetTotalPageCount() *string
	SetTotalRecordCount(v string) *DescribeMetaListResponseBody
	GetTotalRecordCount() *string
}

type DescribeMetaListResponseBody struct {
	// The cluster ID.
	//
	// example:
	//
	// pc-bp1s826a1up******
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The details of recoverable databases and tables.
	Items []*DescribeMetaListResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of records on the current page.
	//
	// example:
	//
	// 30
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// AA815DE7-B576-4B22-B33C-3FB31A******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of pages.
	//
	// example:
	//
	// 1
	TotalPageCount *string `json:"TotalPageCount,omitempty" xml:"TotalPageCount,omitempty"`
	// The total number of records.
	//
	// example:
	//
	// 2
	TotalRecordCount *string `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
}

func (s DescribeMetaListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeMetaListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMetaListResponseBody) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeMetaListResponseBody) GetItems() []*DescribeMetaListResponseBodyItems {
	return s.Items
}

func (s *DescribeMetaListResponseBody) GetPageNumber() *string {
	return s.PageNumber
}

func (s *DescribeMetaListResponseBody) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribeMetaListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeMetaListResponseBody) GetTotalPageCount() *string {
	return s.TotalPageCount
}

func (s *DescribeMetaListResponseBody) GetTotalRecordCount() *string {
	return s.TotalRecordCount
}

func (s *DescribeMetaListResponseBody) SetDBClusterId(v string) *DescribeMetaListResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *DescribeMetaListResponseBody) SetItems(v []*DescribeMetaListResponseBodyItems) *DescribeMetaListResponseBody {
	s.Items = v
	return s
}

func (s *DescribeMetaListResponseBody) SetPageNumber(v string) *DescribeMetaListResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeMetaListResponseBody) SetPageSize(v string) *DescribeMetaListResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeMetaListResponseBody) SetRequestId(v string) *DescribeMetaListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeMetaListResponseBody) SetTotalPageCount(v string) *DescribeMetaListResponseBody {
	s.TotalPageCount = &v
	return s
}

func (s *DescribeMetaListResponseBody) SetTotalRecordCount(v string) *DescribeMetaListResponseBody {
	s.TotalRecordCount = &v
	return s
}

func (s *DescribeMetaListResponseBody) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeMetaListResponseBodyItems struct {
	// The name of the database that can be recovered.
	//
	// example:
	//
	// test_db
	Database *string `json:"Database,omitempty" xml:"Database,omitempty"`
	// The size of the database or table, in bytes.
	Size []*int64 `json:"Size,omitempty" xml:"Size,omitempty" type:"Repeated"`
	// The names of the tables that can be recovered.
	Tables []*string `json:"Tables,omitempty" xml:"Tables,omitempty" type:"Repeated"`
}

func (s DescribeMetaListResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeMetaListResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeMetaListResponseBodyItems) GetDatabase() *string {
	return s.Database
}

func (s *DescribeMetaListResponseBodyItems) GetSize() []*int64 {
	return s.Size
}

func (s *DescribeMetaListResponseBodyItems) GetTables() []*string {
	return s.Tables
}

func (s *DescribeMetaListResponseBodyItems) SetDatabase(v string) *DescribeMetaListResponseBodyItems {
	s.Database = &v
	return s
}

func (s *DescribeMetaListResponseBodyItems) SetSize(v []*int64) *DescribeMetaListResponseBodyItems {
	s.Size = v
	return s
}

func (s *DescribeMetaListResponseBodyItems) SetTables(v []*string) *DescribeMetaListResponseBodyItems {
	s.Tables = v
	return s
}

func (s *DescribeMetaListResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
