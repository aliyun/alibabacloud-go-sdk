// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSheetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetColumnCount(v int64) *GetSheetResponseBody
	GetColumnCount() *int64
	SetId(v string) *GetSheetResponseBody
	GetId() *string
	SetLastNonEmptyColumn(v int64) *GetSheetResponseBody
	GetLastNonEmptyColumn() *int64
	SetLastNonEmptyRow(v int64) *GetSheetResponseBody
	GetLastNonEmptyRow() *int64
	SetName(v string) *GetSheetResponseBody
	GetName() *string
	SetRequestId(v string) *GetSheetResponseBody
	GetRequestId() *string
	SetRowCount(v int64) *GetSheetResponseBody
	GetRowCount() *int64
	SetVisibility(v string) *GetSheetResponseBody
	GetVisibility() *string
}

type GetSheetResponseBody struct {
	// example:
	//
	// 20
	ColumnCount *int64 `json:"columnCount,omitempty" xml:"columnCount,omitempty"`
	// example:
	//
	// stxxxx
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// 1
	LastNonEmptyColumn *int64 `json:"lastNonEmptyColumn,omitempty" xml:"lastNonEmptyColumn,omitempty"`
	// example:
	//
	// 2
	LastNonEmptyRow *int64 `json:"lastNonEmptyRow,omitempty" xml:"lastNonEmptyRow,omitempty"`
	// example:
	//
	// Sheet1
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// requestId
	//
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 10
	RowCount *int64 `json:"rowCount,omitempty" xml:"rowCount,omitempty"`
	// example:
	//
	// visible
	Visibility *string `json:"visibility,omitempty" xml:"visibility,omitempty"`
}

func (s GetSheetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSheetResponseBody) GoString() string {
	return s.String()
}

func (s *GetSheetResponseBody) GetColumnCount() *int64 {
	return s.ColumnCount
}

func (s *GetSheetResponseBody) GetId() *string {
	return s.Id
}

func (s *GetSheetResponseBody) GetLastNonEmptyColumn() *int64 {
	return s.LastNonEmptyColumn
}

func (s *GetSheetResponseBody) GetLastNonEmptyRow() *int64 {
	return s.LastNonEmptyRow
}

func (s *GetSheetResponseBody) GetName() *string {
	return s.Name
}

func (s *GetSheetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSheetResponseBody) GetRowCount() *int64 {
	return s.RowCount
}

func (s *GetSheetResponseBody) GetVisibility() *string {
	return s.Visibility
}

func (s *GetSheetResponseBody) SetColumnCount(v int64) *GetSheetResponseBody {
	s.ColumnCount = &v
	return s
}

func (s *GetSheetResponseBody) SetId(v string) *GetSheetResponseBody {
	s.Id = &v
	return s
}

func (s *GetSheetResponseBody) SetLastNonEmptyColumn(v int64) *GetSheetResponseBody {
	s.LastNonEmptyColumn = &v
	return s
}

func (s *GetSheetResponseBody) SetLastNonEmptyRow(v int64) *GetSheetResponseBody {
	s.LastNonEmptyRow = &v
	return s
}

func (s *GetSheetResponseBody) SetName(v string) *GetSheetResponseBody {
	s.Name = &v
	return s
}

func (s *GetSheetResponseBody) SetRequestId(v string) *GetSheetResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSheetResponseBody) SetRowCount(v int64) *GetSheetResponseBody {
	s.RowCount = &v
	return s
}

func (s *GetSheetResponseBody) SetVisibility(v string) *GetSheetResponseBody {
	s.Visibility = &v
	return s
}

func (s *GetSheetResponseBody) Validate() error {
	return dara.Validate(s)
}
