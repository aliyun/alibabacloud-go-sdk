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
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetStatementResultResponseBodyData struct {
	ColumnMetadata *GetStatementResultResponseBodyDataColumnMetadata `json:"ColumnMetadata,omitempty" xml:"ColumnMetadata,omitempty" type:"Struct"`
	Records        *GetStatementResultResponseBodyDataRecords        `json:"Records,omitempty" xml:"Records,omitempty" type:"Struct"`
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
	if s.ColumnMetadata != nil {
		if err := s.ColumnMetadata.Validate(); err != nil {
			return err
		}
	}
	if s.Records != nil {
		if err := s.Records.Validate(); err != nil {
			return err
		}
	}
	return nil
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
	if s.ColumnMetadata != nil {
		for _, item := range s.ColumnMetadata {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
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
	if s.Records != nil {
		for _, item := range s.Records {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
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
	if s.Record != nil {
		for _, item := range s.Record {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
