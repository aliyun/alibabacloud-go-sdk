// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRecognizeTableResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *RecognizeTableResponseBodyData) *RecognizeTableResponseBody
	GetData() *RecognizeTableResponseBodyData
	SetRequestId(v string) *RecognizeTableResponseBody
	GetRequestId() *string
}

type RecognizeTableResponseBody struct {
	Data *RecognizeTableResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// CBC36BE6-2A18-5256-82BD-8B5477E5D058
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeTableResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RecognizeTableResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeTableResponseBody) GetData() *RecognizeTableResponseBodyData {
	return s.Data
}

func (s *RecognizeTableResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RecognizeTableResponseBody) SetData(v *RecognizeTableResponseBodyData) *RecognizeTableResponseBody {
	s.Data = v
	return s
}

func (s *RecognizeTableResponseBody) SetRequestId(v string) *RecognizeTableResponseBody {
	s.RequestId = &v
	return s
}

func (s *RecognizeTableResponseBody) Validate() error {
	return dara.Validate(s)
}

type RecognizeTableResponseBodyData struct {
	// example:
	//
	// UEsDBBQAAAAIAAAAIQBukMk4WAIAA****
	FileContent *string                                 `json:"FileContent,omitempty" xml:"FileContent,omitempty"`
	Tables      []*RecognizeTableResponseBodyDataTables `json:"Tables,omitempty" xml:"Tables,omitempty" type:"Repeated"`
}

func (s RecognizeTableResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s RecognizeTableResponseBodyData) GoString() string {
	return s.String()
}

func (s *RecognizeTableResponseBodyData) GetFileContent() *string {
	return s.FileContent
}

func (s *RecognizeTableResponseBodyData) GetTables() []*RecognizeTableResponseBodyDataTables {
	return s.Tables
}

func (s *RecognizeTableResponseBodyData) SetFileContent(v string) *RecognizeTableResponseBodyData {
	s.FileContent = &v
	return s
}

func (s *RecognizeTableResponseBodyData) SetTables(v []*RecognizeTableResponseBodyDataTables) *RecognizeTableResponseBodyData {
	s.Tables = v
	return s
}

func (s *RecognizeTableResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type RecognizeTableResponseBodyDataTables struct {
	Head      []*string                                        `json:"Head,omitempty" xml:"Head,omitempty" type:"Repeated"`
	TableRows []*RecognizeTableResponseBodyDataTablesTableRows `json:"TableRows,omitempty" xml:"TableRows,omitempty" type:"Repeated"`
	Tail      []*string                                        `json:"Tail,omitempty" xml:"Tail,omitempty" type:"Repeated"`
}

func (s RecognizeTableResponseBodyDataTables) String() string {
	return dara.Prettify(s)
}

func (s RecognizeTableResponseBodyDataTables) GoString() string {
	return s.String()
}

func (s *RecognizeTableResponseBodyDataTables) GetHead() []*string {
	return s.Head
}

func (s *RecognizeTableResponseBodyDataTables) GetTableRows() []*RecognizeTableResponseBodyDataTablesTableRows {
	return s.TableRows
}

func (s *RecognizeTableResponseBodyDataTables) GetTail() []*string {
	return s.Tail
}

func (s *RecognizeTableResponseBodyDataTables) SetHead(v []*string) *RecognizeTableResponseBodyDataTables {
	s.Head = v
	return s
}

func (s *RecognizeTableResponseBodyDataTables) SetTableRows(v []*RecognizeTableResponseBodyDataTablesTableRows) *RecognizeTableResponseBodyDataTables {
	s.TableRows = v
	return s
}

func (s *RecognizeTableResponseBodyDataTables) SetTail(v []*string) *RecognizeTableResponseBodyDataTables {
	s.Tail = v
	return s
}

func (s *RecognizeTableResponseBodyDataTables) Validate() error {
	return dara.Validate(s)
}

type RecognizeTableResponseBodyDataTablesTableRows struct {
	TableColumns []*RecognizeTableResponseBodyDataTablesTableRowsTableColumns `json:"TableColumns,omitempty" xml:"TableColumns,omitempty" type:"Repeated"`
}

func (s RecognizeTableResponseBodyDataTablesTableRows) String() string {
	return dara.Prettify(s)
}

func (s RecognizeTableResponseBodyDataTablesTableRows) GoString() string {
	return s.String()
}

func (s *RecognizeTableResponseBodyDataTablesTableRows) GetTableColumns() []*RecognizeTableResponseBodyDataTablesTableRowsTableColumns {
	return s.TableColumns
}

func (s *RecognizeTableResponseBodyDataTablesTableRows) SetTableColumns(v []*RecognizeTableResponseBodyDataTablesTableRowsTableColumns) *RecognizeTableResponseBodyDataTablesTableRows {
	s.TableColumns = v
	return s
}

func (s *RecognizeTableResponseBodyDataTablesTableRows) Validate() error {
	return dara.Validate(s)
}

type RecognizeTableResponseBodyDataTablesTableRowsTableColumns struct {
	// example:
	//
	// 4
	EndColumn *int32 `json:"EndColumn,omitempty" xml:"EndColumn,omitempty"`
	// example:
	//
	// 1
	EndRow *int32 `json:"EndRow,omitempty" xml:"EndRow,omitempty"`
	// example:
	//
	// 0
	Height *int32 `json:"Height,omitempty" xml:"Height,omitempty"`
	// example:
	//
	// 1
	StartColumn *int32 `json:"StartColumn,omitempty" xml:"StartColumn,omitempty"`
	// example:
	//
	// 0
	StartRow *int32    `json:"StartRow,omitempty" xml:"StartRow,omitempty"`
	Texts    []*string `json:"Texts,omitempty" xml:"Texts,omitempty" type:"Repeated"`
	// example:
	//
	// 0
	Width *int32 `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s RecognizeTableResponseBodyDataTablesTableRowsTableColumns) String() string {
	return dara.Prettify(s)
}

func (s RecognizeTableResponseBodyDataTablesTableRowsTableColumns) GoString() string {
	return s.String()
}

func (s *RecognizeTableResponseBodyDataTablesTableRowsTableColumns) GetEndColumn() *int32 {
	return s.EndColumn
}

func (s *RecognizeTableResponseBodyDataTablesTableRowsTableColumns) GetEndRow() *int32 {
	return s.EndRow
}

func (s *RecognizeTableResponseBodyDataTablesTableRowsTableColumns) GetHeight() *int32 {
	return s.Height
}

func (s *RecognizeTableResponseBodyDataTablesTableRowsTableColumns) GetStartColumn() *int32 {
	return s.StartColumn
}

func (s *RecognizeTableResponseBodyDataTablesTableRowsTableColumns) GetStartRow() *int32 {
	return s.StartRow
}

func (s *RecognizeTableResponseBodyDataTablesTableRowsTableColumns) GetTexts() []*string {
	return s.Texts
}

func (s *RecognizeTableResponseBodyDataTablesTableRowsTableColumns) GetWidth() *int32 {
	return s.Width
}

func (s *RecognizeTableResponseBodyDataTablesTableRowsTableColumns) SetEndColumn(v int32) *RecognizeTableResponseBodyDataTablesTableRowsTableColumns {
	s.EndColumn = &v
	return s
}

func (s *RecognizeTableResponseBodyDataTablesTableRowsTableColumns) SetEndRow(v int32) *RecognizeTableResponseBodyDataTablesTableRowsTableColumns {
	s.EndRow = &v
	return s
}

func (s *RecognizeTableResponseBodyDataTablesTableRowsTableColumns) SetHeight(v int32) *RecognizeTableResponseBodyDataTablesTableRowsTableColumns {
	s.Height = &v
	return s
}

func (s *RecognizeTableResponseBodyDataTablesTableRowsTableColumns) SetStartColumn(v int32) *RecognizeTableResponseBodyDataTablesTableRowsTableColumns {
	s.StartColumn = &v
	return s
}

func (s *RecognizeTableResponseBodyDataTablesTableRowsTableColumns) SetStartRow(v int32) *RecognizeTableResponseBodyDataTablesTableRowsTableColumns {
	s.StartRow = &v
	return s
}

func (s *RecognizeTableResponseBodyDataTablesTableRowsTableColumns) SetTexts(v []*string) *RecognizeTableResponseBodyDataTablesTableRowsTableColumns {
	s.Texts = v
	return s
}

func (s *RecognizeTableResponseBodyDataTablesTableRowsTableColumns) SetWidth(v int32) *RecognizeTableResponseBodyDataTablesTableRowsTableColumns {
	s.Width = &v
	return s
}

func (s *RecognizeTableResponseBodyDataTablesTableRowsTableColumns) Validate() error {
	return dara.Validate(s)
}
