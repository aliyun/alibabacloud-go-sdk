// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeZoneRecordsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *DescribeZoneRecordsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeZoneRecordsResponseBody
	GetPageSize() *int32
	SetRecords(v *DescribeZoneRecordsResponseBodyRecords) *DescribeZoneRecordsResponseBody
	GetRecords() *DescribeZoneRecordsResponseBodyRecords
	SetRequestId(v string) *DescribeZoneRecordsResponseBody
	GetRequestId() *string
	SetTotalItems(v int32) *DescribeZoneRecordsResponseBody
	GetTotalItems() *int32
	SetTotalPages(v int32) *DescribeZoneRecordsResponseBody
	GetTotalPages() *int32
}

type DescribeZoneRecordsResponseBody struct {
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 100
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The DNS records.
	Records *DescribeZoneRecordsResponseBodyRecords `json:"Records,omitempty" xml:"Records,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 7B07FBC3-3A53-4939-A3C6-2BDFE407BAB2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 100
	TotalItems *int32 `json:"TotalItems,omitempty" xml:"TotalItems,omitempty"`
	// The total number of pages returned.
	//
	// example:
	//
	// 100
	TotalPages *int32 `json:"TotalPages,omitempty" xml:"TotalPages,omitempty"`
}

func (s DescribeZoneRecordsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeZoneRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeZoneRecordsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeZoneRecordsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeZoneRecordsResponseBody) GetRecords() *DescribeZoneRecordsResponseBodyRecords {
	return s.Records
}

func (s *DescribeZoneRecordsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeZoneRecordsResponseBody) GetTotalItems() *int32 {
	return s.TotalItems
}

func (s *DescribeZoneRecordsResponseBody) GetTotalPages() *int32 {
	return s.TotalPages
}

func (s *DescribeZoneRecordsResponseBody) SetPageNumber(v int32) *DescribeZoneRecordsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeZoneRecordsResponseBody) SetPageSize(v int32) *DescribeZoneRecordsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeZoneRecordsResponseBody) SetRecords(v *DescribeZoneRecordsResponseBodyRecords) *DescribeZoneRecordsResponseBody {
	s.Records = v
	return s
}

func (s *DescribeZoneRecordsResponseBody) SetRequestId(v string) *DescribeZoneRecordsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeZoneRecordsResponseBody) SetTotalItems(v int32) *DescribeZoneRecordsResponseBody {
	s.TotalItems = &v
	return s
}

func (s *DescribeZoneRecordsResponseBody) SetTotalPages(v int32) *DescribeZoneRecordsResponseBody {
	s.TotalPages = &v
	return s
}

func (s *DescribeZoneRecordsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeZoneRecordsResponseBodyRecords struct {
	Record []*DescribeZoneRecordsResponseBodyRecordsRecord `json:"Record,omitempty" xml:"Record,omitempty" type:"Repeated"`
}

func (s DescribeZoneRecordsResponseBodyRecords) String() string {
	return dara.Prettify(s)
}

func (s DescribeZoneRecordsResponseBodyRecords) GoString() string {
	return s.String()
}

func (s *DescribeZoneRecordsResponseBodyRecords) GetRecord() []*DescribeZoneRecordsResponseBodyRecordsRecord {
	return s.Record
}

func (s *DescribeZoneRecordsResponseBodyRecords) SetRecord(v []*DescribeZoneRecordsResponseBodyRecordsRecord) *DescribeZoneRecordsResponseBodyRecords {
	s.Record = v
	return s
}

func (s *DescribeZoneRecordsResponseBodyRecords) Validate() error {
	return dara.Validate(s)
}

type DescribeZoneRecordsResponseBodyRecordsRecord struct {
	// The time when the DNS record was created. The time follows the ISO 8601 standard in the YYYY-MM-DDThh:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2019-03-14T03:47Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The time when the DNS record was created. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1672740294000
	CreateTimestamp *int64 `json:"CreateTimestamp,omitempty" xml:"CreateTimestamp,omitempty"`
	// The resolution line.
	//
	// example:
	//
	// default
	Line *string `json:"Line,omitempty" xml:"Line,omitempty"`
	// The priority of the mail exchanger (MX) record.
	//
	// example:
	//
	// 10
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The ID of the DNS record.
	//
	// example:
	//
	// 246959****
	RecordId *int64 `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
	// The description of the DNS record.
	//
	// example:
	//
	// test
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The hostname.
	//
	// example:
	//
	// www
	Rr *string `json:"Rr,omitempty" xml:"Rr,omitempty"`
	// The state of the DNS record. Valid values:
	//
	// 	- ENABLE: The DNS record is enabled.
	//
	// 	- DISABLE: The DNS record is disabled.
	//
	// example:
	//
	// ENABLE
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The time to live (TTL) period.
	//
	// example:
	//
	// 60
	Ttl *int32 `json:"Ttl,omitempty" xml:"Ttl,omitempty"`
	// The type of the DNS record. Valid values:
	//
	// 	- **A**: An A record points a domain name to an IPv4 address.
	//
	// 	- **AAAA**: An AAAA record points a domain name to an IPv6 address.
	//
	// 	- **CNAME**: A canonical name (CNAME) record points a domain name to another domain name.
	//
	// 	- **TXT**: A text (TXT) record usually serves as a Sender Policy Framework (SPF) record to prevent email spam. The record value of the TXT record can be up to 255 characters in length.
	//
	// 	- **MX**: A mail exchanger (MX) record points a domain name to a mail server address.
	//
	// 	- **PTR**: A pointer (PTR) points an IP address to a domain name.
	//
	// 	- **SRV**: A service (SRV) record specifies a server that hosts a specific service.
	//
	// example:
	//
	// A
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The time when the DNS record was updated. The time follows the ISO 8601 standard in the YYYY-MM-DDThh:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2023-05-08T02:31Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// The time when the DNS record was updated. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1654777678000
	UpdateTimestamp *int64 `json:"UpdateTimestamp,omitempty" xml:"UpdateTimestamp,omitempty"`
	// The record value.
	//
	// example:
	//
	// 127.0.XX.XX
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
	// The weight value of the address. You can set a different weight value for each address. This way, addresses are returned based on the weight values for DNS requests. A weight value must be an integer that ranges from 1 to 100.
	//
	// example:
	//
	// 1
	Weight *int32 `json:"Weight,omitempty" xml:"Weight,omitempty"`
	// The zone ID.
	//
	// example:
	//
	// a49f55537f3b0b1e6e43add0bf5f****
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeZoneRecordsResponseBodyRecordsRecord) String() string {
	return dara.Prettify(s)
}

func (s DescribeZoneRecordsResponseBodyRecordsRecord) GoString() string {
	return s.String()
}

func (s *DescribeZoneRecordsResponseBodyRecordsRecord) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeZoneRecordsResponseBodyRecordsRecord) GetCreateTimestamp() *int64 {
	return s.CreateTimestamp
}

func (s *DescribeZoneRecordsResponseBodyRecordsRecord) GetLine() *string {
	return s.Line
}

func (s *DescribeZoneRecordsResponseBodyRecordsRecord) GetPriority() *int32 {
	return s.Priority
}

func (s *DescribeZoneRecordsResponseBodyRecordsRecord) GetRecordId() *int64 {
	return s.RecordId
}

func (s *DescribeZoneRecordsResponseBodyRecordsRecord) GetRemark() *string {
	return s.Remark
}

func (s *DescribeZoneRecordsResponseBodyRecordsRecord) GetRr() *string {
	return s.Rr
}

func (s *DescribeZoneRecordsResponseBodyRecordsRecord) GetStatus() *string {
	return s.Status
}

func (s *DescribeZoneRecordsResponseBodyRecordsRecord) GetTtl() *int32 {
	return s.Ttl
}

func (s *DescribeZoneRecordsResponseBodyRecordsRecord) GetType() *string {
	return s.Type
}

func (s *DescribeZoneRecordsResponseBodyRecordsRecord) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *DescribeZoneRecordsResponseBodyRecordsRecord) GetUpdateTimestamp() *int64 {
	return s.UpdateTimestamp
}

func (s *DescribeZoneRecordsResponseBodyRecordsRecord) GetValue() *string {
	return s.Value
}

func (s *DescribeZoneRecordsResponseBodyRecordsRecord) GetWeight() *int32 {
	return s.Weight
}

func (s *DescribeZoneRecordsResponseBodyRecordsRecord) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeZoneRecordsResponseBodyRecordsRecord) SetCreateTime(v string) *DescribeZoneRecordsResponseBodyRecordsRecord {
	s.CreateTime = &v
	return s
}

func (s *DescribeZoneRecordsResponseBodyRecordsRecord) SetCreateTimestamp(v int64) *DescribeZoneRecordsResponseBodyRecordsRecord {
	s.CreateTimestamp = &v
	return s
}

func (s *DescribeZoneRecordsResponseBodyRecordsRecord) SetLine(v string) *DescribeZoneRecordsResponseBodyRecordsRecord {
	s.Line = &v
	return s
}

func (s *DescribeZoneRecordsResponseBodyRecordsRecord) SetPriority(v int32) *DescribeZoneRecordsResponseBodyRecordsRecord {
	s.Priority = &v
	return s
}

func (s *DescribeZoneRecordsResponseBodyRecordsRecord) SetRecordId(v int64) *DescribeZoneRecordsResponseBodyRecordsRecord {
	s.RecordId = &v
	return s
}

func (s *DescribeZoneRecordsResponseBodyRecordsRecord) SetRemark(v string) *DescribeZoneRecordsResponseBodyRecordsRecord {
	s.Remark = &v
	return s
}

func (s *DescribeZoneRecordsResponseBodyRecordsRecord) SetRr(v string) *DescribeZoneRecordsResponseBodyRecordsRecord {
	s.Rr = &v
	return s
}

func (s *DescribeZoneRecordsResponseBodyRecordsRecord) SetStatus(v string) *DescribeZoneRecordsResponseBodyRecordsRecord {
	s.Status = &v
	return s
}

func (s *DescribeZoneRecordsResponseBodyRecordsRecord) SetTtl(v int32) *DescribeZoneRecordsResponseBodyRecordsRecord {
	s.Ttl = &v
	return s
}

func (s *DescribeZoneRecordsResponseBodyRecordsRecord) SetType(v string) *DescribeZoneRecordsResponseBodyRecordsRecord {
	s.Type = &v
	return s
}

func (s *DescribeZoneRecordsResponseBodyRecordsRecord) SetUpdateTime(v string) *DescribeZoneRecordsResponseBodyRecordsRecord {
	s.UpdateTime = &v
	return s
}

func (s *DescribeZoneRecordsResponseBodyRecordsRecord) SetUpdateTimestamp(v int64) *DescribeZoneRecordsResponseBodyRecordsRecord {
	s.UpdateTimestamp = &v
	return s
}

func (s *DescribeZoneRecordsResponseBodyRecordsRecord) SetValue(v string) *DescribeZoneRecordsResponseBodyRecordsRecord {
	s.Value = &v
	return s
}

func (s *DescribeZoneRecordsResponseBodyRecordsRecord) SetWeight(v int32) *DescribeZoneRecordsResponseBodyRecordsRecord {
	s.Weight = &v
	return s
}

func (s *DescribeZoneRecordsResponseBodyRecordsRecord) SetZoneId(v string) *DescribeZoneRecordsResponseBodyRecordsRecord {
	s.ZoneId = &v
	return s
}

func (s *DescribeZoneRecordsResponseBodyRecordsRecord) Validate() error {
	return dara.Validate(s)
}
