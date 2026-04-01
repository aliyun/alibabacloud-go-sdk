// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMetaListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceName(v string) *DescribeMetaListResponseBody
	GetDBInstanceName() *string
	SetItems(v *DescribeMetaListResponseBodyItems) *DescribeMetaListResponseBody
	GetItems() *DescribeMetaListResponseBodyItems
	SetPageNumber(v int32) *DescribeMetaListResponseBody
	GetPageNumber() *int32
	SetPageRecordCount(v int32) *DescribeMetaListResponseBody
	GetPageRecordCount() *int32
	SetRequestId(v string) *DescribeMetaListResponseBody
	GetRequestId() *string
	SetTotalPageCount(v int32) *DescribeMetaListResponseBody
	GetTotalPageCount() *int32
	SetTotalRecordCount(v int32) *DescribeMetaListResponseBody
	GetTotalRecordCount() *int32
}

type DescribeMetaListResponseBody struct {
	// The instance name.
	//
	// example:
	//
	// rm-uf6wjk5xxxxxxx
	DBInstanceName *string                            `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	Items          *DescribeMetaListResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 1
	PageRecordCount *int32 `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 60F9A12A-16B8-4728-B099-4CA38D32C31C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of pages returned.
	//
	// example:
	//
	// 1
	TotalPageCount *int32 `json:"TotalPageCount,omitempty" xml:"TotalPageCount,omitempty"`
	// The total number of returned entries.
	//
	// example:
	//
	// 1
	TotalRecordCount *int32 `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
}

func (s DescribeMetaListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeMetaListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMetaListResponseBody) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *DescribeMetaListResponseBody) GetItems() *DescribeMetaListResponseBodyItems {
	return s.Items
}

func (s *DescribeMetaListResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeMetaListResponseBody) GetPageRecordCount() *int32 {
	return s.PageRecordCount
}

func (s *DescribeMetaListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeMetaListResponseBody) GetTotalPageCount() *int32 {
	return s.TotalPageCount
}

func (s *DescribeMetaListResponseBody) GetTotalRecordCount() *int32 {
	return s.TotalRecordCount
}

func (s *DescribeMetaListResponseBody) SetDBInstanceName(v string) *DescribeMetaListResponseBody {
	s.DBInstanceName = &v
	return s
}

func (s *DescribeMetaListResponseBody) SetItems(v *DescribeMetaListResponseBodyItems) *DescribeMetaListResponseBody {
	s.Items = v
	return s
}

func (s *DescribeMetaListResponseBody) SetPageNumber(v int32) *DescribeMetaListResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeMetaListResponseBody) SetPageRecordCount(v int32) *DescribeMetaListResponseBody {
	s.PageRecordCount = &v
	return s
}

func (s *DescribeMetaListResponseBody) SetRequestId(v string) *DescribeMetaListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeMetaListResponseBody) SetTotalPageCount(v int32) *DescribeMetaListResponseBody {
	s.TotalPageCount = &v
	return s
}

func (s *DescribeMetaListResponseBody) SetTotalRecordCount(v int32) *DescribeMetaListResponseBody {
	s.TotalRecordCount = &v
	return s
}

func (s *DescribeMetaListResponseBody) Validate() error {
	if s.Items != nil {
		if err := s.Items.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeMetaListResponseBodyItems struct {
	Meta []*DescribeMetaListResponseBodyItemsMeta `json:"Meta,omitempty" xml:"Meta,omitempty" type:"Repeated"`
}

func (s DescribeMetaListResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeMetaListResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeMetaListResponseBodyItems) GetMeta() []*DescribeMetaListResponseBodyItemsMeta {
	return s.Meta
}

func (s *DescribeMetaListResponseBodyItems) SetMeta(v []*DescribeMetaListResponseBodyItemsMeta) *DescribeMetaListResponseBodyItems {
	s.Meta = v
	return s
}

func (s *DescribeMetaListResponseBodyItems) Validate() error {
	if s.Meta != nil {
		for _, item := range s.Meta {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeMetaListResponseBodyItemsMeta struct {
	Database *string `json:"Database,omitempty" xml:"Database,omitempty"`
	Size     *string `json:"Size,omitempty" xml:"Size,omitempty"`
	Tables   *string `json:"Tables,omitempty" xml:"Tables,omitempty"`
}

func (s DescribeMetaListResponseBodyItemsMeta) String() string {
	return dara.Prettify(s)
}

func (s DescribeMetaListResponseBodyItemsMeta) GoString() string {
	return s.String()
}

func (s *DescribeMetaListResponseBodyItemsMeta) GetDatabase() *string {
	return s.Database
}

func (s *DescribeMetaListResponseBodyItemsMeta) GetSize() *string {
	return s.Size
}

func (s *DescribeMetaListResponseBodyItemsMeta) GetTables() *string {
	return s.Tables
}

func (s *DescribeMetaListResponseBodyItemsMeta) SetDatabase(v string) *DescribeMetaListResponseBodyItemsMeta {
	s.Database = &v
	return s
}

func (s *DescribeMetaListResponseBodyItemsMeta) SetSize(v string) *DescribeMetaListResponseBodyItemsMeta {
	s.Size = &v
	return s
}

func (s *DescribeMetaListResponseBodyItemsMeta) SetTables(v string) *DescribeMetaListResponseBodyItemsMeta {
	s.Tables = &v
	return s
}

func (s *DescribeMetaListResponseBodyItemsMeta) Validate() error {
	return dara.Validate(s)
}
