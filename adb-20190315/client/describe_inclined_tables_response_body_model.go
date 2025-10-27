// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInclinedTablesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDetectionItems(v []*DescribeInclinedTablesResponseBodyDetectionItems) *DescribeInclinedTablesResponseBody
	GetDetectionItems() []*DescribeInclinedTablesResponseBodyDetectionItems
	SetItems(v *DescribeInclinedTablesResponseBodyItems) *DescribeInclinedTablesResponseBody
	GetItems() *DescribeInclinedTablesResponseBodyItems
	SetPageNumber(v string) *DescribeInclinedTablesResponseBody
	GetPageNumber() *string
	SetPageSize(v string) *DescribeInclinedTablesResponseBody
	GetPageSize() *string
	SetRequestId(v string) *DescribeInclinedTablesResponseBody
	GetRequestId() *string
	SetTotalCount(v string) *DescribeInclinedTablesResponseBody
	GetTotalCount() *string
}

type DescribeInclinedTablesResponseBody struct {
	// The queried detection items and detection results.
	DetectionItems []*DescribeInclinedTablesResponseBodyDetectionItems `json:"DetectionItems,omitempty" xml:"DetectionItems,omitempty" type:"Repeated"`
	// The queried tables.
	Items *DescribeInclinedTablesResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 30
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1AD222E9-E606-4A42-BF6D-8A4442913CEF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 15
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeInclinedTablesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeInclinedTablesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInclinedTablesResponseBody) GetDetectionItems() []*DescribeInclinedTablesResponseBodyDetectionItems {
	return s.DetectionItems
}

func (s *DescribeInclinedTablesResponseBody) GetItems() *DescribeInclinedTablesResponseBodyItems {
	return s.Items
}

func (s *DescribeInclinedTablesResponseBody) GetPageNumber() *string {
	return s.PageNumber
}

func (s *DescribeInclinedTablesResponseBody) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribeInclinedTablesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeInclinedTablesResponseBody) GetTotalCount() *string {
	return s.TotalCount
}

func (s *DescribeInclinedTablesResponseBody) SetDetectionItems(v []*DescribeInclinedTablesResponseBodyDetectionItems) *DescribeInclinedTablesResponseBody {
	s.DetectionItems = v
	return s
}

func (s *DescribeInclinedTablesResponseBody) SetItems(v *DescribeInclinedTablesResponseBodyItems) *DescribeInclinedTablesResponseBody {
	s.Items = v
	return s
}

func (s *DescribeInclinedTablesResponseBody) SetPageNumber(v string) *DescribeInclinedTablesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeInclinedTablesResponseBody) SetPageSize(v string) *DescribeInclinedTablesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeInclinedTablesResponseBody) SetRequestId(v string) *DescribeInclinedTablesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInclinedTablesResponseBody) SetTotalCount(v string) *DescribeInclinedTablesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeInclinedTablesResponseBody) Validate() error {
	if s.DetectionItems != nil {
		for _, item := range s.DetectionItems {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Items != nil {
		if err := s.Items.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeInclinedTablesResponseBodyDetectionItems struct {
	// The message of the detection result.
	//
	// example:
	//
	// There are a total of 10 tables with an excessive number of primary keys.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The name of the detection item.
	//
	// example:
	//
	// Excessive primary key fields
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The severity level of the detection result. Valid values:
	//
	// - NORMAL
	//
	// - WARNING
	//
	// - CRITICAL
	//
	// example:
	//
	// WARNING
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeInclinedTablesResponseBodyDetectionItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeInclinedTablesResponseBodyDetectionItems) GoString() string {
	return s.String()
}

func (s *DescribeInclinedTablesResponseBodyDetectionItems) GetMessage() *string {
	return s.Message
}

func (s *DescribeInclinedTablesResponseBodyDetectionItems) GetName() *string {
	return s.Name
}

func (s *DescribeInclinedTablesResponseBodyDetectionItems) GetStatus() *string {
	return s.Status
}

func (s *DescribeInclinedTablesResponseBodyDetectionItems) SetMessage(v string) *DescribeInclinedTablesResponseBodyDetectionItems {
	s.Message = &v
	return s
}

func (s *DescribeInclinedTablesResponseBodyDetectionItems) SetName(v string) *DescribeInclinedTablesResponseBodyDetectionItems {
	s.Name = &v
	return s
}

func (s *DescribeInclinedTablesResponseBodyDetectionItems) SetStatus(v string) *DescribeInclinedTablesResponseBodyDetectionItems {
	s.Status = &v
	return s
}

func (s *DescribeInclinedTablesResponseBodyDetectionItems) Validate() error {
	return dara.Validate(s)
}

type DescribeInclinedTablesResponseBodyItems struct {
	// The queried table.
	Table []*DescribeInclinedTablesResponseBodyItemsTable `json:"Table,omitempty" xml:"Table,omitempty" type:"Repeated"`
}

func (s DescribeInclinedTablesResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeInclinedTablesResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeInclinedTablesResponseBodyItems) GetTable() []*DescribeInclinedTablesResponseBodyItemsTable {
	return s.Table
}

func (s *DescribeInclinedTablesResponseBodyItems) SetTable(v []*DescribeInclinedTablesResponseBodyItemsTable) *DescribeInclinedTablesResponseBodyItems {
	s.Table = v
	return s
}

func (s *DescribeInclinedTablesResponseBodyItems) Validate() error {
	if s.Table != nil {
		for _, item := range s.Table {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeInclinedTablesResponseBodyItemsTable struct {
	// Indicates whether data is skewed in the table.
	//
	// example:
	//
	// true
	IsIncline *string `json:"IsIncline,omitempty" xml:"IsIncline,omitempty"`
	// The name of the table.
	//
	// example:
	//
	// ff
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The number of rows in the table.
	//
	// example:
	//
	// 2565
	RowCount *int64 `json:"RowCount,omitempty" xml:"RowCount,omitempty"`
	// The name of the database.
	//
	// example:
	//
	// adm_analyze
	Schema *string `json:"Schema,omitempty" xml:"Schema,omitempty"`
	// The number of rows in the table.
	//
	// example:
	//
	// 2416527
	Size *string `json:"Size,omitempty" xml:"Size,omitempty"`
	// The percentage of the table size.
	//
	// example:
	//
	// 89
	SpaceRatio *float64 `json:"SpaceRatio,omitempty" xml:"SpaceRatio,omitempty"`
	// The total data size of the table.
	//
	// example:
	//
	// 65
	TotalSize *int64 `json:"TotalSize,omitempty" xml:"TotalSize,omitempty"`
	// The column type.
	//
	// example:
	//
	// executor
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeInclinedTablesResponseBodyItemsTable) String() string {
	return dara.Prettify(s)
}

func (s DescribeInclinedTablesResponseBodyItemsTable) GoString() string {
	return s.String()
}

func (s *DescribeInclinedTablesResponseBodyItemsTable) GetIsIncline() *string {
	return s.IsIncline
}

func (s *DescribeInclinedTablesResponseBodyItemsTable) GetName() *string {
	return s.Name
}

func (s *DescribeInclinedTablesResponseBodyItemsTable) GetRowCount() *int64 {
	return s.RowCount
}

func (s *DescribeInclinedTablesResponseBodyItemsTable) GetSchema() *string {
	return s.Schema
}

func (s *DescribeInclinedTablesResponseBodyItemsTable) GetSize() *string {
	return s.Size
}

func (s *DescribeInclinedTablesResponseBodyItemsTable) GetSpaceRatio() *float64 {
	return s.SpaceRatio
}

func (s *DescribeInclinedTablesResponseBodyItemsTable) GetTotalSize() *int64 {
	return s.TotalSize
}

func (s *DescribeInclinedTablesResponseBodyItemsTable) GetType() *string {
	return s.Type
}

func (s *DescribeInclinedTablesResponseBodyItemsTable) SetIsIncline(v string) *DescribeInclinedTablesResponseBodyItemsTable {
	s.IsIncline = &v
	return s
}

func (s *DescribeInclinedTablesResponseBodyItemsTable) SetName(v string) *DescribeInclinedTablesResponseBodyItemsTable {
	s.Name = &v
	return s
}

func (s *DescribeInclinedTablesResponseBodyItemsTable) SetRowCount(v int64) *DescribeInclinedTablesResponseBodyItemsTable {
	s.RowCount = &v
	return s
}

func (s *DescribeInclinedTablesResponseBodyItemsTable) SetSchema(v string) *DescribeInclinedTablesResponseBodyItemsTable {
	s.Schema = &v
	return s
}

func (s *DescribeInclinedTablesResponseBodyItemsTable) SetSize(v string) *DescribeInclinedTablesResponseBodyItemsTable {
	s.Size = &v
	return s
}

func (s *DescribeInclinedTablesResponseBodyItemsTable) SetSpaceRatio(v float64) *DescribeInclinedTablesResponseBodyItemsTable {
	s.SpaceRatio = &v
	return s
}

func (s *DescribeInclinedTablesResponseBodyItemsTable) SetTotalSize(v int64) *DescribeInclinedTablesResponseBodyItemsTable {
	s.TotalSize = &v
	return s
}

func (s *DescribeInclinedTablesResponseBodyItemsTable) SetType(v string) *DescribeInclinedTablesResponseBodyItemsTable {
	s.Type = &v
	return s
}

func (s *DescribeInclinedTablesResponseBodyItemsTable) Validate() error {
	return dara.Validate(s)
}
