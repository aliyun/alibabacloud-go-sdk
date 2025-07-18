// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetStatementResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetStatementResultResponseBodyData) *GetStatementResultResponseBody
	GetData() *GetStatementResultResponseBodyData
	SetMessage(v string) *GetStatementResultResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetStatementResultResponseBody
	GetRequestId() *string
	SetStatus(v string) *GetStatementResultResponseBody
	GetStatus() *string
}

type GetStatementResultResponseBody struct {
	// The result of the asynchronous call.
	Data *GetStatementResultResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Return message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID.
	//
	// example:
	//
	// ABB39CC3-4488-4857-905D-2E4A051D0521
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// API execution status, with values as follows:
	//
	// - **false**: Execution failed.
	//
	// - **true**: Execution succeeded.
	//
	// example:
	//
	// success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetStatementResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetStatementResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetStatementResultResponseBody) GetData() *GetStatementResultResponseBodyData {
	return s.Data
}

func (s *GetStatementResultResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetStatementResultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetStatementResultResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetStatementResultResponseBody) SetData(v *GetStatementResultResponseBodyData) *GetStatementResultResponseBody {
	s.Data = v
	return s
}

func (s *GetStatementResultResponseBody) SetMessage(v string) *GetStatementResultResponseBody {
	s.Message = &v
	return s
}

func (s *GetStatementResultResponseBody) SetRequestId(v string) *GetStatementResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetStatementResultResponseBody) SetStatus(v string) *GetStatementResultResponseBody {
	s.Status = &v
	return s
}

func (s *GetStatementResultResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetStatementResultResponseBodyData struct {
	// List of column metadata.
	ColumnMetadata *GetStatementResultResponseBodyDataColumnMetadata `json:"ColumnMetadata,omitempty" xml:"ColumnMetadata,omitempty" type:"Struct"`
	// Multiple rows of data.
	Records *GetStatementResultResponseBodyDataRecords `json:"Records,omitempty" xml:"Records,omitempty" type:"Struct"`
	// Total number of entries.
	//
	// example:
	//
	// 10
	TotalNumRows *int64 `json:"TotalNumRows,omitempty" xml:"TotalNumRows,omitempty"`
}

func (s GetStatementResultResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetStatementResultResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetStatementResultResponseBodyData) GetColumnMetadata() *GetStatementResultResponseBodyDataColumnMetadata {
	return s.ColumnMetadata
}

func (s *GetStatementResultResponseBodyData) GetRecords() *GetStatementResultResponseBodyDataRecords {
	return s.Records
}

func (s *GetStatementResultResponseBodyData) GetTotalNumRows() *int64 {
	return s.TotalNumRows
}

func (s *GetStatementResultResponseBodyData) SetColumnMetadata(v *GetStatementResultResponseBodyDataColumnMetadata) *GetStatementResultResponseBodyData {
	s.ColumnMetadata = v
	return s
}

func (s *GetStatementResultResponseBodyData) SetRecords(v *GetStatementResultResponseBodyDataRecords) *GetStatementResultResponseBodyData {
	s.Records = v
	return s
}

func (s *GetStatementResultResponseBodyData) SetTotalNumRows(v int64) *GetStatementResultResponseBodyData {
	s.TotalNumRows = &v
	return s
}

func (s *GetStatementResultResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type GetStatementResultResponseBodyDataColumnMetadata struct {
	ColumnMetadata []*ColumnMetadata `json:"ColumnMetadata,omitempty" xml:"ColumnMetadata,omitempty" type:"Repeated"`
}

func (s GetStatementResultResponseBodyDataColumnMetadata) String() string {
	return dara.Prettify(s)
}

func (s GetStatementResultResponseBodyDataColumnMetadata) GoString() string {
	return s.String()
}

func (s *GetStatementResultResponseBodyDataColumnMetadata) GetColumnMetadata() []*ColumnMetadata {
	return s.ColumnMetadata
}

func (s *GetStatementResultResponseBodyDataColumnMetadata) SetColumnMetadata(v []*ColumnMetadata) *GetStatementResultResponseBodyDataColumnMetadata {
	s.ColumnMetadata = v
	return s
}

func (s *GetStatementResultResponseBodyDataColumnMetadata) Validate() error {
	return dara.Validate(s)
}

type GetStatementResultResponseBodyDataRecords struct {
	Records []*GetStatementResultResponseBodyDataRecordsRecords `json:"Records,omitempty" xml:"Records,omitempty" type:"Repeated"`
}

func (s GetStatementResultResponseBodyDataRecords) String() string {
	return dara.Prettify(s)
}

func (s GetStatementResultResponseBodyDataRecords) GoString() string {
	return s.String()
}

func (s *GetStatementResultResponseBodyDataRecords) GetRecords() []*GetStatementResultResponseBodyDataRecordsRecords {
	return s.Records
}

func (s *GetStatementResultResponseBodyDataRecords) SetRecords(v []*GetStatementResultResponseBodyDataRecordsRecords) *GetStatementResultResponseBodyDataRecords {
	s.Records = v
	return s
}

func (s *GetStatementResultResponseBodyDataRecords) Validate() error {
	return dara.Validate(s)
}

type GetStatementResultResponseBodyDataRecordsRecords struct {
	Record []*Field `json:"Record,omitempty" xml:"Record,omitempty" type:"Repeated"`
}

func (s GetStatementResultResponseBodyDataRecordsRecords) String() string {
	return dara.Prettify(s)
}

func (s GetStatementResultResponseBodyDataRecordsRecords) GoString() string {
	return s.String()
}

func (s *GetStatementResultResponseBodyDataRecordsRecords) GetRecord() []*Field {
	return s.Record
}

func (s *GetStatementResultResponseBodyDataRecordsRecords) SetRecord(v []*Field) *GetStatementResultResponseBodyDataRecordsRecords {
	s.Record = v
	return s
}

func (s *GetStatementResultResponseBodyDataRecordsRecords) Validate() error {
	return dara.Validate(s)
}
