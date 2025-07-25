// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSubDomainRecordsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDomainRecords(v *DescribeSubDomainRecordsResponseBodyDomainRecords) *DescribeSubDomainRecordsResponseBody
	GetDomainRecords() *DescribeSubDomainRecordsResponseBodyDomainRecords
	SetPageNumber(v int64) *DescribeSubDomainRecordsResponseBody
	GetPageNumber() *int64
	SetPageSize(v int64) *DescribeSubDomainRecordsResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *DescribeSubDomainRecordsResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *DescribeSubDomainRecordsResponseBody
	GetTotalCount() *int64
}

type DescribeSubDomainRecordsResponseBody struct {
	// The returned Domain Name System (DNS) records.
	DomainRecords *DescribeSubDomainRecordsResponseBodyDomainRecords `json:"DomainRecords,omitempty" xml:"DomainRecords,omitempty" type:"Struct"`
	// The page number. Pages start from page **1**. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 2
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

func (s DescribeSubDomainRecordsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSubDomainRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSubDomainRecordsResponseBody) GetDomainRecords() *DescribeSubDomainRecordsResponseBodyDomainRecords {
	return s.DomainRecords
}

func (s *DescribeSubDomainRecordsResponseBody) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeSubDomainRecordsResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeSubDomainRecordsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSubDomainRecordsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeSubDomainRecordsResponseBody) SetDomainRecords(v *DescribeSubDomainRecordsResponseBodyDomainRecords) *DescribeSubDomainRecordsResponseBody {
	s.DomainRecords = v
	return s
}

func (s *DescribeSubDomainRecordsResponseBody) SetPageNumber(v int64) *DescribeSubDomainRecordsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeSubDomainRecordsResponseBody) SetPageSize(v int64) *DescribeSubDomainRecordsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeSubDomainRecordsResponseBody) SetRequestId(v string) *DescribeSubDomainRecordsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSubDomainRecordsResponseBody) SetTotalCount(v int64) *DescribeSubDomainRecordsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeSubDomainRecordsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeSubDomainRecordsResponseBodyDomainRecords struct {
	Record []*DescribeSubDomainRecordsResponseBodyDomainRecordsRecord `json:"Record,omitempty" xml:"Record,omitempty" type:"Repeated"`
}

func (s DescribeSubDomainRecordsResponseBodyDomainRecords) String() string {
	return dara.Prettify(s)
}

func (s DescribeSubDomainRecordsResponseBodyDomainRecords) GoString() string {
	return s.String()
}

func (s *DescribeSubDomainRecordsResponseBodyDomainRecords) GetRecord() []*DescribeSubDomainRecordsResponseBodyDomainRecordsRecord {
	return s.Record
}

func (s *DescribeSubDomainRecordsResponseBodyDomainRecords) SetRecord(v []*DescribeSubDomainRecordsResponseBodyDomainRecordsRecord) *DescribeSubDomainRecordsResponseBodyDomainRecords {
	s.Record = v
	return s
}

func (s *DescribeSubDomainRecordsResponseBodyDomainRecords) Validate() error {
	return dara.Validate(s)
}

type DescribeSubDomainRecordsResponseBodyDomainRecordsRecord struct {
	// The domain name.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The DNS resolution line.
	//
	// example:
	//
	// default
	Line *string `json:"Line,omitempty" xml:"Line,omitempty"`
	// The lock status of the DNS record.
	//
	// example:
	//
	// false
	Locked *bool `json:"Locked,omitempty" xml:"Locked,omitempty"`
	// The priority of the mail exchanger (MX) record.
	//
	// example:
	//
	// 2
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
	// The time-to-live (TTL) of the DNS record.
	//
	// example:
	//
	// 600
	TTL *int64 `json:"TTL,omitempty" xml:"TTL,omitempty"`
	// The type of the DNS record.
	//
	// example:
	//
	// CNAME
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The record value.
	//
	// example:
	//
	// example.net
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
	// The weight of the DNS record.
	//
	// example:
	//
	// 10
	Weight *int32 `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s DescribeSubDomainRecordsResponseBodyDomainRecordsRecord) String() string {
	return dara.Prettify(s)
}

func (s DescribeSubDomainRecordsResponseBodyDomainRecordsRecord) GoString() string {
	return s.String()
}

func (s *DescribeSubDomainRecordsResponseBodyDomainRecordsRecord) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeSubDomainRecordsResponseBodyDomainRecordsRecord) GetLine() *string {
	return s.Line
}

func (s *DescribeSubDomainRecordsResponseBodyDomainRecordsRecord) GetLocked() *bool {
	return s.Locked
}

func (s *DescribeSubDomainRecordsResponseBodyDomainRecordsRecord) GetPriority() *int64 {
	return s.Priority
}

func (s *DescribeSubDomainRecordsResponseBodyDomainRecordsRecord) GetRR() *string {
	return s.RR
}

func (s *DescribeSubDomainRecordsResponseBodyDomainRecordsRecord) GetRecordId() *string {
	return s.RecordId
}

func (s *DescribeSubDomainRecordsResponseBodyDomainRecordsRecord) GetRemark() *string {
	return s.Remark
}

func (s *DescribeSubDomainRecordsResponseBodyDomainRecordsRecord) GetStatus() *string {
	return s.Status
}

func (s *DescribeSubDomainRecordsResponseBodyDomainRecordsRecord) GetTTL() *int64 {
	return s.TTL
}

func (s *DescribeSubDomainRecordsResponseBodyDomainRecordsRecord) GetType() *string {
	return s.Type
}

func (s *DescribeSubDomainRecordsResponseBodyDomainRecordsRecord) GetValue() *string {
	return s.Value
}

func (s *DescribeSubDomainRecordsResponseBodyDomainRecordsRecord) GetWeight() *int32 {
	return s.Weight
}

func (s *DescribeSubDomainRecordsResponseBodyDomainRecordsRecord) SetDomainName(v string) *DescribeSubDomainRecordsResponseBodyDomainRecordsRecord {
	s.DomainName = &v
	return s
}

func (s *DescribeSubDomainRecordsResponseBodyDomainRecordsRecord) SetLine(v string) *DescribeSubDomainRecordsResponseBodyDomainRecordsRecord {
	s.Line = &v
	return s
}

func (s *DescribeSubDomainRecordsResponseBodyDomainRecordsRecord) SetLocked(v bool) *DescribeSubDomainRecordsResponseBodyDomainRecordsRecord {
	s.Locked = &v
	return s
}

func (s *DescribeSubDomainRecordsResponseBodyDomainRecordsRecord) SetPriority(v int64) *DescribeSubDomainRecordsResponseBodyDomainRecordsRecord {
	s.Priority = &v
	return s
}

func (s *DescribeSubDomainRecordsResponseBodyDomainRecordsRecord) SetRR(v string) *DescribeSubDomainRecordsResponseBodyDomainRecordsRecord {
	s.RR = &v
	return s
}

func (s *DescribeSubDomainRecordsResponseBodyDomainRecordsRecord) SetRecordId(v string) *DescribeSubDomainRecordsResponseBodyDomainRecordsRecord {
	s.RecordId = &v
	return s
}

func (s *DescribeSubDomainRecordsResponseBodyDomainRecordsRecord) SetRemark(v string) *DescribeSubDomainRecordsResponseBodyDomainRecordsRecord {
	s.Remark = &v
	return s
}

func (s *DescribeSubDomainRecordsResponseBodyDomainRecordsRecord) SetStatus(v string) *DescribeSubDomainRecordsResponseBodyDomainRecordsRecord {
	s.Status = &v
	return s
}

func (s *DescribeSubDomainRecordsResponseBodyDomainRecordsRecord) SetTTL(v int64) *DescribeSubDomainRecordsResponseBodyDomainRecordsRecord {
	s.TTL = &v
	return s
}

func (s *DescribeSubDomainRecordsResponseBodyDomainRecordsRecord) SetType(v string) *DescribeSubDomainRecordsResponseBodyDomainRecordsRecord {
	s.Type = &v
	return s
}

func (s *DescribeSubDomainRecordsResponseBodyDomainRecordsRecord) SetValue(v string) *DescribeSubDomainRecordsResponseBodyDomainRecordsRecord {
	s.Value = &v
	return s
}

func (s *DescribeSubDomainRecordsResponseBodyDomainRecordsRecord) SetWeight(v int32) *DescribeSubDomainRecordsResponseBodyDomainRecordsRecord {
	s.Weight = &v
	return s
}

func (s *DescribeSubDomainRecordsResponseBodyDomainRecordsRecord) Validate() error {
	return dara.Validate(s)
}
