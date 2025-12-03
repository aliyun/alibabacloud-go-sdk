// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBackupTablesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBackupRecords(v *DescribeBackupTablesResponseBodyBackupRecords) *DescribeBackupTablesResponseBody
	GetBackupRecords() *DescribeBackupTablesResponseBodyBackupRecords
	SetPageNumber(v int32) *DescribeBackupTablesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeBackupTablesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeBackupTablesResponseBody
	GetRequestId() *string
	SetTables(v *DescribeBackupTablesResponseBodyTables) *DescribeBackupTablesResponseBody
	GetTables() *DescribeBackupTablesResponseBodyTables
	SetTotal(v int64) *DescribeBackupTablesResponseBody
	GetTotal() *int64
}

type DescribeBackupTablesResponseBody struct {
	BackupRecords *DescribeBackupTablesResponseBodyBackupRecords `json:"BackupRecords,omitempty" xml:"BackupRecords,omitempty" type:"Struct"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 01262E9C-B0CC-4663-82FA-D50173649F92
	RequestId *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Tables    *DescribeBackupTablesResponseBodyTables `json:"Tables,omitempty" xml:"Tables,omitempty" type:"Struct"`
	// example:
	//
	// 1
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeBackupTablesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupTablesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBackupTablesResponseBody) GetBackupRecords() *DescribeBackupTablesResponseBodyBackupRecords {
	return s.BackupRecords
}

func (s *DescribeBackupTablesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeBackupTablesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeBackupTablesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeBackupTablesResponseBody) GetTables() *DescribeBackupTablesResponseBodyTables {
	return s.Tables
}

func (s *DescribeBackupTablesResponseBody) GetTotal() *int64 {
	return s.Total
}

func (s *DescribeBackupTablesResponseBody) SetBackupRecords(v *DescribeBackupTablesResponseBodyBackupRecords) *DescribeBackupTablesResponseBody {
	s.BackupRecords = v
	return s
}

func (s *DescribeBackupTablesResponseBody) SetPageNumber(v int32) *DescribeBackupTablesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeBackupTablesResponseBody) SetPageSize(v int32) *DescribeBackupTablesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeBackupTablesResponseBody) SetRequestId(v string) *DescribeBackupTablesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBackupTablesResponseBody) SetTables(v *DescribeBackupTablesResponseBodyTables) *DescribeBackupTablesResponseBody {
	s.Tables = v
	return s
}

func (s *DescribeBackupTablesResponseBody) SetTotal(v int64) *DescribeBackupTablesResponseBody {
	s.Total = &v
	return s
}

func (s *DescribeBackupTablesResponseBody) Validate() error {
	if s.BackupRecords != nil {
		if err := s.BackupRecords.Validate(); err != nil {
			return err
		}
	}
	if s.Tables != nil {
		if err := s.Tables.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeBackupTablesResponseBodyBackupRecords struct {
	BackupRecord []*DescribeBackupTablesResponseBodyBackupRecordsBackupRecord `json:"BackupRecord,omitempty" xml:"BackupRecord,omitempty" type:"Repeated"`
}

func (s DescribeBackupTablesResponseBodyBackupRecords) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupTablesResponseBodyBackupRecords) GoString() string {
	return s.String()
}

func (s *DescribeBackupTablesResponseBodyBackupRecords) GetBackupRecord() []*DescribeBackupTablesResponseBodyBackupRecordsBackupRecord {
	return s.BackupRecord
}

func (s *DescribeBackupTablesResponseBodyBackupRecords) SetBackupRecord(v []*DescribeBackupTablesResponseBodyBackupRecordsBackupRecord) *DescribeBackupTablesResponseBodyBackupRecords {
	s.BackupRecord = v
	return s
}

func (s *DescribeBackupTablesResponseBodyBackupRecords) Validate() error {
	if s.BackupRecord != nil {
		for _, item := range s.BackupRecord {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeBackupTablesResponseBodyBackupRecordsBackupRecord struct {
	// example:
	//
	// 1.2 kB
	DataSize *string `json:"DataSize,omitempty" xml:"DataSize,omitempty"`
	// example:
	//
	// 2020-11-02T18:00:05Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// null
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 14/14
	Process *string `json:"Process,omitempty" xml:"Process,omitempty"`
	// example:
	//
	// 0.00 MB/s
	Speed *string `json:"Speed,omitempty" xml:"Speed,omitempty"`
	// example:
	//
	// 2020-11-02T18:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// SUCCEEDED
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// example:
	//
	// default:test1
	Table *string `json:"Table,omitempty" xml:"Table,omitempty"`
}

func (s DescribeBackupTablesResponseBodyBackupRecordsBackupRecord) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupTablesResponseBodyBackupRecordsBackupRecord) GoString() string {
	return s.String()
}

func (s *DescribeBackupTablesResponseBodyBackupRecordsBackupRecord) GetDataSize() *string {
	return s.DataSize
}

func (s *DescribeBackupTablesResponseBodyBackupRecordsBackupRecord) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeBackupTablesResponseBodyBackupRecordsBackupRecord) GetMessage() *string {
	return s.Message
}

func (s *DescribeBackupTablesResponseBodyBackupRecordsBackupRecord) GetProcess() *string {
	return s.Process
}

func (s *DescribeBackupTablesResponseBodyBackupRecordsBackupRecord) GetSpeed() *string {
	return s.Speed
}

func (s *DescribeBackupTablesResponseBodyBackupRecordsBackupRecord) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeBackupTablesResponseBodyBackupRecordsBackupRecord) GetState() *string {
	return s.State
}

func (s *DescribeBackupTablesResponseBodyBackupRecordsBackupRecord) GetTable() *string {
	return s.Table
}

func (s *DescribeBackupTablesResponseBodyBackupRecordsBackupRecord) SetDataSize(v string) *DescribeBackupTablesResponseBodyBackupRecordsBackupRecord {
	s.DataSize = &v
	return s
}

func (s *DescribeBackupTablesResponseBodyBackupRecordsBackupRecord) SetEndTime(v string) *DescribeBackupTablesResponseBodyBackupRecordsBackupRecord {
	s.EndTime = &v
	return s
}

func (s *DescribeBackupTablesResponseBodyBackupRecordsBackupRecord) SetMessage(v string) *DescribeBackupTablesResponseBodyBackupRecordsBackupRecord {
	s.Message = &v
	return s
}

func (s *DescribeBackupTablesResponseBodyBackupRecordsBackupRecord) SetProcess(v string) *DescribeBackupTablesResponseBodyBackupRecordsBackupRecord {
	s.Process = &v
	return s
}

func (s *DescribeBackupTablesResponseBodyBackupRecordsBackupRecord) SetSpeed(v string) *DescribeBackupTablesResponseBodyBackupRecordsBackupRecord {
	s.Speed = &v
	return s
}

func (s *DescribeBackupTablesResponseBodyBackupRecordsBackupRecord) SetStartTime(v string) *DescribeBackupTablesResponseBodyBackupRecordsBackupRecord {
	s.StartTime = &v
	return s
}

func (s *DescribeBackupTablesResponseBodyBackupRecordsBackupRecord) SetState(v string) *DescribeBackupTablesResponseBodyBackupRecordsBackupRecord {
	s.State = &v
	return s
}

func (s *DescribeBackupTablesResponseBodyBackupRecordsBackupRecord) SetTable(v string) *DescribeBackupTablesResponseBodyBackupRecordsBackupRecord {
	s.Table = &v
	return s
}

func (s *DescribeBackupTablesResponseBodyBackupRecordsBackupRecord) Validate() error {
	return dara.Validate(s)
}

type DescribeBackupTablesResponseBodyTables struct {
	Table []*string `json:"Table,omitempty" xml:"Table,omitempty" type:"Repeated"`
}

func (s DescribeBackupTablesResponseBodyTables) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupTablesResponseBodyTables) GoString() string {
	return s.String()
}

func (s *DescribeBackupTablesResponseBodyTables) GetTable() []*string {
	return s.Table
}

func (s *DescribeBackupTablesResponseBodyTables) SetTable(v []*string) *DescribeBackupTablesResponseBodyTables {
	s.Table = v
	return s
}

func (s *DescribeBackupTablesResponseBodyTables) Validate() error {
	return dara.Validate(s)
}
