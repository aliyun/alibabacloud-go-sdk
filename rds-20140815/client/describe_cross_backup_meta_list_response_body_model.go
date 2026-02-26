// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCrossBackupMetaListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceName(v string) *DescribeCrossBackupMetaListResponseBody
	GetDBInstanceName() *string
	SetItems(v *DescribeCrossBackupMetaListResponseBodyItems) *DescribeCrossBackupMetaListResponseBody
	GetItems() *DescribeCrossBackupMetaListResponseBodyItems
	SetPageNumber(v int32) *DescribeCrossBackupMetaListResponseBody
	GetPageNumber() *int32
	SetPageRecordCount(v int32) *DescribeCrossBackupMetaListResponseBody
	GetPageRecordCount() *int32
	SetRequestId(v string) *DescribeCrossBackupMetaListResponseBody
	GetRequestId() *string
	SetTotalPageCount(v int32) *DescribeCrossBackupMetaListResponseBody
	GetTotalPageCount() *int32
	SetTotalRecordCount(v int32) *DescribeCrossBackupMetaListResponseBody
	GetTotalRecordCount() *int32
}

type DescribeCrossBackupMetaListResponseBody struct {
	// The instance to which the cross-region backup file belongs.
	//
	// example:
	//
	// rm-uf6wjk5xxxxxxx
	DBInstanceName *string                                       `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	Items          *DescribeCrossBackupMetaListResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
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

func (s DescribeCrossBackupMetaListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCrossBackupMetaListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCrossBackupMetaListResponseBody) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *DescribeCrossBackupMetaListResponseBody) GetItems() *DescribeCrossBackupMetaListResponseBodyItems {
	return s.Items
}

func (s *DescribeCrossBackupMetaListResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeCrossBackupMetaListResponseBody) GetPageRecordCount() *int32 {
	return s.PageRecordCount
}

func (s *DescribeCrossBackupMetaListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCrossBackupMetaListResponseBody) GetTotalPageCount() *int32 {
	return s.TotalPageCount
}

func (s *DescribeCrossBackupMetaListResponseBody) GetTotalRecordCount() *int32 {
	return s.TotalRecordCount
}

func (s *DescribeCrossBackupMetaListResponseBody) SetDBInstanceName(v string) *DescribeCrossBackupMetaListResponseBody {
	s.DBInstanceName = &v
	return s
}

func (s *DescribeCrossBackupMetaListResponseBody) SetItems(v *DescribeCrossBackupMetaListResponseBodyItems) *DescribeCrossBackupMetaListResponseBody {
	s.Items = v
	return s
}

func (s *DescribeCrossBackupMetaListResponseBody) SetPageNumber(v int32) *DescribeCrossBackupMetaListResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeCrossBackupMetaListResponseBody) SetPageRecordCount(v int32) *DescribeCrossBackupMetaListResponseBody {
	s.PageRecordCount = &v
	return s
}

func (s *DescribeCrossBackupMetaListResponseBody) SetRequestId(v string) *DescribeCrossBackupMetaListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCrossBackupMetaListResponseBody) SetTotalPageCount(v int32) *DescribeCrossBackupMetaListResponseBody {
	s.TotalPageCount = &v
	return s
}

func (s *DescribeCrossBackupMetaListResponseBody) SetTotalRecordCount(v int32) *DescribeCrossBackupMetaListResponseBody {
	s.TotalRecordCount = &v
	return s
}

func (s *DescribeCrossBackupMetaListResponseBody) Validate() error {
	if s.Items != nil {
		if err := s.Items.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeCrossBackupMetaListResponseBodyItems struct {
	Meta []*DescribeCrossBackupMetaListResponseBodyItemsMeta `json:"Meta,omitempty" xml:"Meta,omitempty" type:"Repeated"`
}

func (s DescribeCrossBackupMetaListResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeCrossBackupMetaListResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeCrossBackupMetaListResponseBodyItems) GetMeta() []*DescribeCrossBackupMetaListResponseBodyItemsMeta {
	return s.Meta
}

func (s *DescribeCrossBackupMetaListResponseBodyItems) SetMeta(v []*DescribeCrossBackupMetaListResponseBodyItemsMeta) *DescribeCrossBackupMetaListResponseBodyItems {
	s.Meta = v
	return s
}

func (s *DescribeCrossBackupMetaListResponseBodyItems) Validate() error {
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

type DescribeCrossBackupMetaListResponseBodyItemsMeta struct {
	Database *string `json:"Database,omitempty" xml:"Database,omitempty"`
	Size     *string `json:"Size,omitempty" xml:"Size,omitempty"`
	Tables   *string `json:"Tables,omitempty" xml:"Tables,omitempty"`
}

func (s DescribeCrossBackupMetaListResponseBodyItemsMeta) String() string {
	return dara.Prettify(s)
}

func (s DescribeCrossBackupMetaListResponseBodyItemsMeta) GoString() string {
	return s.String()
}

func (s *DescribeCrossBackupMetaListResponseBodyItemsMeta) GetDatabase() *string {
	return s.Database
}

func (s *DescribeCrossBackupMetaListResponseBodyItemsMeta) GetSize() *string {
	return s.Size
}

func (s *DescribeCrossBackupMetaListResponseBodyItemsMeta) GetTables() *string {
	return s.Tables
}

func (s *DescribeCrossBackupMetaListResponseBodyItemsMeta) SetDatabase(v string) *DescribeCrossBackupMetaListResponseBodyItemsMeta {
	s.Database = &v
	return s
}

func (s *DescribeCrossBackupMetaListResponseBodyItemsMeta) SetSize(v string) *DescribeCrossBackupMetaListResponseBodyItemsMeta {
	s.Size = &v
	return s
}

func (s *DescribeCrossBackupMetaListResponseBodyItemsMeta) SetTables(v string) *DescribeCrossBackupMetaListResponseBodyItemsMeta {
	s.Tables = &v
	return s
}

func (s *DescribeCrossBackupMetaListResponseBodyItemsMeta) Validate() error {
	return dara.Validate(s)
}
