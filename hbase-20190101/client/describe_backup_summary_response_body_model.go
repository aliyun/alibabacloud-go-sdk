// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBackupSummaryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFull(v *DescribeBackupSummaryResponseBodyFull) *DescribeBackupSummaryResponseBody
	GetFull() *DescribeBackupSummaryResponseBodyFull
	SetIncr(v *DescribeBackupSummaryResponseBodyIncr) *DescribeBackupSummaryResponseBody
	GetIncr() *DescribeBackupSummaryResponseBodyIncr
	SetRequestId(v string) *DescribeBackupSummaryResponseBody
	GetRequestId() *string
}

type DescribeBackupSummaryResponseBody struct {
	Full *DescribeBackupSummaryResponseBodyFull `json:"Full,omitempty" xml:"Full,omitempty" type:"Struct"`
	Incr *DescribeBackupSummaryResponseBodyIncr `json:"Incr,omitempty" xml:"Incr,omitempty" type:"Struct"`
	// example:
	//
	// 168793CB-7B31-43E7-ADAB-FE3E8D584D6E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeBackupSummaryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupSummaryResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBackupSummaryResponseBody) GetFull() *DescribeBackupSummaryResponseBodyFull {
	return s.Full
}

func (s *DescribeBackupSummaryResponseBody) GetIncr() *DescribeBackupSummaryResponseBodyIncr {
	return s.Incr
}

func (s *DescribeBackupSummaryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeBackupSummaryResponseBody) SetFull(v *DescribeBackupSummaryResponseBodyFull) *DescribeBackupSummaryResponseBody {
	s.Full = v
	return s
}

func (s *DescribeBackupSummaryResponseBody) SetIncr(v *DescribeBackupSummaryResponseBodyIncr) *DescribeBackupSummaryResponseBody {
	s.Incr = v
	return s
}

func (s *DescribeBackupSummaryResponseBody) SetRequestId(v string) *DescribeBackupSummaryResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBackupSummaryResponseBody) Validate() error {
	if s.Full != nil {
		if err := s.Full.Validate(); err != nil {
			return err
		}
	}
	if s.Incr != nil {
		if err := s.Incr.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeBackupSummaryResponseBodyFull struct {
	// example:
	//
	// false
	HasMore *string `json:"HasMore,omitempty" xml:"HasMore,omitempty"`
	// example:
	//
	// 2020-11-09T18:00:00Z
	NextFullBackupDate *string `json:"NextFullBackupDate,omitempty" xml:"NextFullBackupDate,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32                                        `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Records  *DescribeBackupSummaryResponseBodyFullRecords `json:"Records,omitempty" xml:"Records,omitempty" type:"Struct"`
	// example:
	//
	// 2
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeBackupSummaryResponseBodyFull) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupSummaryResponseBodyFull) GoString() string {
	return s.String()
}

func (s *DescribeBackupSummaryResponseBodyFull) GetHasMore() *string {
	return s.HasMore
}

func (s *DescribeBackupSummaryResponseBodyFull) GetNextFullBackupDate() *string {
	return s.NextFullBackupDate
}

func (s *DescribeBackupSummaryResponseBodyFull) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeBackupSummaryResponseBodyFull) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeBackupSummaryResponseBodyFull) GetRecords() *DescribeBackupSummaryResponseBodyFullRecords {
	return s.Records
}

func (s *DescribeBackupSummaryResponseBodyFull) GetTotal() *int32 {
	return s.Total
}

func (s *DescribeBackupSummaryResponseBodyFull) SetHasMore(v string) *DescribeBackupSummaryResponseBodyFull {
	s.HasMore = &v
	return s
}

func (s *DescribeBackupSummaryResponseBodyFull) SetNextFullBackupDate(v string) *DescribeBackupSummaryResponseBodyFull {
	s.NextFullBackupDate = &v
	return s
}

func (s *DescribeBackupSummaryResponseBodyFull) SetPageNumber(v int32) *DescribeBackupSummaryResponseBodyFull {
	s.PageNumber = &v
	return s
}

func (s *DescribeBackupSummaryResponseBodyFull) SetPageSize(v int32) *DescribeBackupSummaryResponseBodyFull {
	s.PageSize = &v
	return s
}

func (s *DescribeBackupSummaryResponseBodyFull) SetRecords(v *DescribeBackupSummaryResponseBodyFullRecords) *DescribeBackupSummaryResponseBodyFull {
	s.Records = v
	return s
}

func (s *DescribeBackupSummaryResponseBodyFull) SetTotal(v int32) *DescribeBackupSummaryResponseBodyFull {
	s.Total = &v
	return s
}

func (s *DescribeBackupSummaryResponseBodyFull) Validate() error {
	if s.Records != nil {
		if err := s.Records.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeBackupSummaryResponseBodyFullRecords struct {
	Record []*DescribeBackupSummaryResponseBodyFullRecordsRecord `json:"Record,omitempty" xml:"Record,omitempty" type:"Repeated"`
}

func (s DescribeBackupSummaryResponseBodyFullRecords) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupSummaryResponseBodyFullRecords) GoString() string {
	return s.String()
}

func (s *DescribeBackupSummaryResponseBodyFullRecords) GetRecord() []*DescribeBackupSummaryResponseBodyFullRecordsRecord {
	return s.Record
}

func (s *DescribeBackupSummaryResponseBodyFullRecords) SetRecord(v []*DescribeBackupSummaryResponseBodyFullRecordsRecord) *DescribeBackupSummaryResponseBodyFullRecords {
	s.Record = v
	return s
}

func (s *DescribeBackupSummaryResponseBodyFullRecords) Validate() error {
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

type DescribeBackupSummaryResponseBodyFullRecordsRecord struct {
	// example:
	//
	// 2020-11-02T18:00:00Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 1.2 kB
	DataSize *string `json:"DataSize,omitempty" xml:"DataSize,omitempty"`
	// example:
	//
	// 2020-11-02T18:02:04Z
	FinishTime *string `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	// example:
	//
	// 1/1
	Process *string `json:"Process,omitempty" xml:"Process,omitempty"`
	// example:
	//
	// 20201103020000
	RecordId *string `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
	// example:
	//
	// 0.00 MB/s
	Speed *string `json:"Speed,omitempty" xml:"Speed,omitempty"`
	// example:
	//
	// SUCCESS
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeBackupSummaryResponseBodyFullRecordsRecord) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupSummaryResponseBodyFullRecordsRecord) GoString() string {
	return s.String()
}

func (s *DescribeBackupSummaryResponseBodyFullRecordsRecord) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeBackupSummaryResponseBodyFullRecordsRecord) GetDataSize() *string {
	return s.DataSize
}

func (s *DescribeBackupSummaryResponseBodyFullRecordsRecord) GetFinishTime() *string {
	return s.FinishTime
}

func (s *DescribeBackupSummaryResponseBodyFullRecordsRecord) GetProcess() *string {
	return s.Process
}

func (s *DescribeBackupSummaryResponseBodyFullRecordsRecord) GetRecordId() *string {
	return s.RecordId
}

func (s *DescribeBackupSummaryResponseBodyFullRecordsRecord) GetSpeed() *string {
	return s.Speed
}

func (s *DescribeBackupSummaryResponseBodyFullRecordsRecord) GetStatus() *string {
	return s.Status
}

func (s *DescribeBackupSummaryResponseBodyFullRecordsRecord) SetCreateTime(v string) *DescribeBackupSummaryResponseBodyFullRecordsRecord {
	s.CreateTime = &v
	return s
}

func (s *DescribeBackupSummaryResponseBodyFullRecordsRecord) SetDataSize(v string) *DescribeBackupSummaryResponseBodyFullRecordsRecord {
	s.DataSize = &v
	return s
}

func (s *DescribeBackupSummaryResponseBodyFullRecordsRecord) SetFinishTime(v string) *DescribeBackupSummaryResponseBodyFullRecordsRecord {
	s.FinishTime = &v
	return s
}

func (s *DescribeBackupSummaryResponseBodyFullRecordsRecord) SetProcess(v string) *DescribeBackupSummaryResponseBodyFullRecordsRecord {
	s.Process = &v
	return s
}

func (s *DescribeBackupSummaryResponseBodyFullRecordsRecord) SetRecordId(v string) *DescribeBackupSummaryResponseBodyFullRecordsRecord {
	s.RecordId = &v
	return s
}

func (s *DescribeBackupSummaryResponseBodyFullRecordsRecord) SetSpeed(v string) *DescribeBackupSummaryResponseBodyFullRecordsRecord {
	s.Speed = &v
	return s
}

func (s *DescribeBackupSummaryResponseBodyFullRecordsRecord) SetStatus(v string) *DescribeBackupSummaryResponseBodyFullRecordsRecord {
	s.Status = &v
	return s
}

func (s *DescribeBackupSummaryResponseBodyFullRecordsRecord) Validate() error {
	return dara.Validate(s)
}

type DescribeBackupSummaryResponseBodyIncr struct {
	// example:
	//
	// 266 B
	BackupLogSize *string `json:"BackupLogSize,omitempty" xml:"BackupLogSize,omitempty"`
	// example:
	//
	// 2020-11-05T01:20:31Z
	Pos *string `json:"Pos,omitempty" xml:"Pos,omitempty"`
	// example:
	//
	// 0
	QueueLogNum *string `json:"QueueLogNum,omitempty" xml:"QueueLogNum,omitempty"`
	// example:
	//
	// 2
	RunningLogNum *string `json:"RunningLogNum,omitempty" xml:"RunningLogNum,omitempty"`
	// example:
	//
	// 0.00 MB/s
	Speed *string `json:"Speed,omitempty" xml:"Speed,omitempty"`
	// example:
	//
	// RUNNING
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeBackupSummaryResponseBodyIncr) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupSummaryResponseBodyIncr) GoString() string {
	return s.String()
}

func (s *DescribeBackupSummaryResponseBodyIncr) GetBackupLogSize() *string {
	return s.BackupLogSize
}

func (s *DescribeBackupSummaryResponseBodyIncr) GetPos() *string {
	return s.Pos
}

func (s *DescribeBackupSummaryResponseBodyIncr) GetQueueLogNum() *string {
	return s.QueueLogNum
}

func (s *DescribeBackupSummaryResponseBodyIncr) GetRunningLogNum() *string {
	return s.RunningLogNum
}

func (s *DescribeBackupSummaryResponseBodyIncr) GetSpeed() *string {
	return s.Speed
}

func (s *DescribeBackupSummaryResponseBodyIncr) GetStatus() *string {
	return s.Status
}

func (s *DescribeBackupSummaryResponseBodyIncr) SetBackupLogSize(v string) *DescribeBackupSummaryResponseBodyIncr {
	s.BackupLogSize = &v
	return s
}

func (s *DescribeBackupSummaryResponseBodyIncr) SetPos(v string) *DescribeBackupSummaryResponseBodyIncr {
	s.Pos = &v
	return s
}

func (s *DescribeBackupSummaryResponseBodyIncr) SetQueueLogNum(v string) *DescribeBackupSummaryResponseBodyIncr {
	s.QueueLogNum = &v
	return s
}

func (s *DescribeBackupSummaryResponseBodyIncr) SetRunningLogNum(v string) *DescribeBackupSummaryResponseBodyIncr {
	s.RunningLogNum = &v
	return s
}

func (s *DescribeBackupSummaryResponseBodyIncr) SetSpeed(v string) *DescribeBackupSummaryResponseBodyIncr {
	s.Speed = &v
	return s
}

func (s *DescribeBackupSummaryResponseBodyIncr) SetStatus(v string) *DescribeBackupSummaryResponseBodyIncr {
	s.Status = &v
	return s
}

func (s *DescribeBackupSummaryResponseBodyIncr) Validate() error {
	return dara.Validate(s)
}
