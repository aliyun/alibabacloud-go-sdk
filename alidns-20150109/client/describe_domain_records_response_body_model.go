// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainRecordsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDomainRecords(v *DescribeDomainRecordsResponseBodyDomainRecords) *DescribeDomainRecordsResponseBody
	GetDomainRecords() *DescribeDomainRecordsResponseBodyDomainRecords
	SetPageNumber(v int64) *DescribeDomainRecordsResponseBody
	GetPageNumber() *int64
	SetPageSize(v int64) *DescribeDomainRecordsResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *DescribeDomainRecordsResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *DescribeDomainRecordsResponseBody
	GetTotalCount() *int64
}

type DescribeDomainRecordsResponseBody struct {
	// The returned Domain Name System (DNS) records.
	DomainRecords *DescribeDomainRecordsResponseBodyDomainRecords `json:"DomainRecords,omitempty" xml:"DomainRecords,omitempty" type:"Struct"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 536E9CAD-DB30-4647-AC87-AA5CC38C5382
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 2
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDomainRecordsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainRecordsResponseBody) GetDomainRecords() *DescribeDomainRecordsResponseBodyDomainRecords {
	return s.DomainRecords
}

func (s *DescribeDomainRecordsResponseBody) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeDomainRecordsResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeDomainRecordsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDomainRecordsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeDomainRecordsResponseBody) SetDomainRecords(v *DescribeDomainRecordsResponseBodyDomainRecords) *DescribeDomainRecordsResponseBody {
	s.DomainRecords = v
	return s
}

func (s *DescribeDomainRecordsResponseBody) SetPageNumber(v int64) *DescribeDomainRecordsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeDomainRecordsResponseBody) SetPageSize(v int64) *DescribeDomainRecordsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeDomainRecordsResponseBody) SetRequestId(v string) *DescribeDomainRecordsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainRecordsResponseBody) SetTotalCount(v int64) *DescribeDomainRecordsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeDomainRecordsResponseBody) Validate() error {
	if s.DomainRecords != nil {
		if err := s.DomainRecords.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDomainRecordsResponseBodyDomainRecords struct {
	Record []*DescribeDomainRecordsResponseBodyDomainRecordsRecord `json:"Record,omitempty" xml:"Record,omitempty" type:"Repeated"`
}

func (s DescribeDomainRecordsResponseBodyDomainRecords) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainRecordsResponseBodyDomainRecords) GoString() string {
	return s.String()
}

func (s *DescribeDomainRecordsResponseBodyDomainRecords) GetRecord() []*DescribeDomainRecordsResponseBodyDomainRecordsRecord {
	return s.Record
}

func (s *DescribeDomainRecordsResponseBodyDomainRecords) SetRecord(v []*DescribeDomainRecordsResponseBodyDomainRecordsRecord) *DescribeDomainRecordsResponseBodyDomainRecords {
	s.Record = v
	return s
}

func (s *DescribeDomainRecordsResponseBodyDomainRecords) Validate() error {
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

type DescribeDomainRecordsResponseBodyDomainRecordsRecord struct {
	// The time when the DNS record was created. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since 00:00:00 UTC on January 1, 1970.
	//
	// example:
	//
	// 1666501957000
	CreateTimestamp *int64 `json:"CreateTimestamp,omitempty" xml:"CreateTimestamp,omitempty"`
	// The domain name.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	LbaStatus  *bool   `json:"LbaStatus,omitempty" xml:"LbaStatus,omitempty"`
	// The resolution line.
	//
	// example:
	//
	// default
	Line *string `json:"Line,omitempty" xml:"Line,omitempty"`
	// Indicates whether the DNS record is locked.
	//
	// example:
	//
	// false
	Locked *bool `json:"Locked,omitempty" xml:"Locked,omitempty"`
	// The priority of the mail exchanger (MX) record.
	//
	// example:
	//
	// 5
	Priority *int64 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The hostname.
	//
	// example:
	//
	// www
	RR *string `json:"RR,omitempty" xml:"RR,omitempty"`
	// The ID of the DNS record.
	//
	// example:
	//
	// 9999985
	RecordId *string `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
	// The description of the DNS record.
	//
	// example:
	//
	// test
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The status of the DNS record.
	//
	// example:
	//
	// Enable
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The time-to-live (TTL) of the cached DNS record. Unit: seconds.
	//
	// example:
	//
	// 600
	TTL *int64 `json:"TTL,omitempty" xml:"TTL,omitempty"`
	// The type of the DNS record.
	//
	// example:
	//
	// MX
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The time when the DNS record was updated. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since 00:00:00 UTC on January 1, 1970.
	//
	// example:
	//
	// 1676872961000
	UpdateTimestamp *int64 `json:"UpdateTimestamp,omitempty" xml:"UpdateTimestamp,omitempty"`
	// The record value.
	//
	// example:
	//
	// mail1.hichina.com
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
	// The weight of the DNS record.
	//
	// example:
	//
	// 2
	Weight *int32 `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s DescribeDomainRecordsResponseBodyDomainRecordsRecord) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainRecordsResponseBodyDomainRecordsRecord) GoString() string {
	return s.String()
}

func (s *DescribeDomainRecordsResponseBodyDomainRecordsRecord) GetCreateTimestamp() *int64 {
	return s.CreateTimestamp
}

func (s *DescribeDomainRecordsResponseBodyDomainRecordsRecord) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDomainRecordsResponseBodyDomainRecordsRecord) GetLbaStatus() *bool {
	return s.LbaStatus
}

func (s *DescribeDomainRecordsResponseBodyDomainRecordsRecord) GetLine() *string {
	return s.Line
}

func (s *DescribeDomainRecordsResponseBodyDomainRecordsRecord) GetLocked() *bool {
	return s.Locked
}

func (s *DescribeDomainRecordsResponseBodyDomainRecordsRecord) GetPriority() *int64 {
	return s.Priority
}

func (s *DescribeDomainRecordsResponseBodyDomainRecordsRecord) GetRR() *string {
	return s.RR
}

func (s *DescribeDomainRecordsResponseBodyDomainRecordsRecord) GetRecordId() *string {
	return s.RecordId
}

func (s *DescribeDomainRecordsResponseBodyDomainRecordsRecord) GetRemark() *string {
	return s.Remark
}

func (s *DescribeDomainRecordsResponseBodyDomainRecordsRecord) GetStatus() *string {
	return s.Status
}

func (s *DescribeDomainRecordsResponseBodyDomainRecordsRecord) GetTTL() *int64 {
	return s.TTL
}

func (s *DescribeDomainRecordsResponseBodyDomainRecordsRecord) GetType() *string {
	return s.Type
}

func (s *DescribeDomainRecordsResponseBodyDomainRecordsRecord) GetUpdateTimestamp() *int64 {
	return s.UpdateTimestamp
}

func (s *DescribeDomainRecordsResponseBodyDomainRecordsRecord) GetValue() *string {
	return s.Value
}

func (s *DescribeDomainRecordsResponseBodyDomainRecordsRecord) GetWeight() *int32 {
	return s.Weight
}

func (s *DescribeDomainRecordsResponseBodyDomainRecordsRecord) SetCreateTimestamp(v int64) *DescribeDomainRecordsResponseBodyDomainRecordsRecord {
	s.CreateTimestamp = &v
	return s
}

func (s *DescribeDomainRecordsResponseBodyDomainRecordsRecord) SetDomainName(v string) *DescribeDomainRecordsResponseBodyDomainRecordsRecord {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainRecordsResponseBodyDomainRecordsRecord) SetLbaStatus(v bool) *DescribeDomainRecordsResponseBodyDomainRecordsRecord {
	s.LbaStatus = &v
	return s
}

func (s *DescribeDomainRecordsResponseBodyDomainRecordsRecord) SetLine(v string) *DescribeDomainRecordsResponseBodyDomainRecordsRecord {
	s.Line = &v
	return s
}

func (s *DescribeDomainRecordsResponseBodyDomainRecordsRecord) SetLocked(v bool) *DescribeDomainRecordsResponseBodyDomainRecordsRecord {
	s.Locked = &v
	return s
}

func (s *DescribeDomainRecordsResponseBodyDomainRecordsRecord) SetPriority(v int64) *DescribeDomainRecordsResponseBodyDomainRecordsRecord {
	s.Priority = &v
	return s
}

func (s *DescribeDomainRecordsResponseBodyDomainRecordsRecord) SetRR(v string) *DescribeDomainRecordsResponseBodyDomainRecordsRecord {
	s.RR = &v
	return s
}

func (s *DescribeDomainRecordsResponseBodyDomainRecordsRecord) SetRecordId(v string) *DescribeDomainRecordsResponseBodyDomainRecordsRecord {
	s.RecordId = &v
	return s
}

func (s *DescribeDomainRecordsResponseBodyDomainRecordsRecord) SetRemark(v string) *DescribeDomainRecordsResponseBodyDomainRecordsRecord {
	s.Remark = &v
	return s
}

func (s *DescribeDomainRecordsResponseBodyDomainRecordsRecord) SetStatus(v string) *DescribeDomainRecordsResponseBodyDomainRecordsRecord {
	s.Status = &v
	return s
}

func (s *DescribeDomainRecordsResponseBodyDomainRecordsRecord) SetTTL(v int64) *DescribeDomainRecordsResponseBodyDomainRecordsRecord {
	s.TTL = &v
	return s
}

func (s *DescribeDomainRecordsResponseBodyDomainRecordsRecord) SetType(v string) *DescribeDomainRecordsResponseBodyDomainRecordsRecord {
	s.Type = &v
	return s
}

func (s *DescribeDomainRecordsResponseBodyDomainRecordsRecord) SetUpdateTimestamp(v int64) *DescribeDomainRecordsResponseBodyDomainRecordsRecord {
	s.UpdateTimestamp = &v
	return s
}

func (s *DescribeDomainRecordsResponseBodyDomainRecordsRecord) SetValue(v string) *DescribeDomainRecordsResponseBodyDomainRecordsRecord {
	s.Value = &v
	return s
}

func (s *DescribeDomainRecordsResponseBodyDomainRecordsRecord) SetWeight(v int32) *DescribeDomainRecordsResponseBodyDomainRecordsRecord {
	s.Weight = &v
	return s
}

func (s *DescribeDomainRecordsResponseBodyDomainRecordsRecord) Validate() error {
	return dara.Validate(s)
}
