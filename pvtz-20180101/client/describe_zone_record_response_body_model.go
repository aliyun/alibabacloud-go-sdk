// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeZoneRecordResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCreateTime(v string) *DescribeZoneRecordResponseBody
	GetCreateTime() *string
	SetCreateTimestamp(v int64) *DescribeZoneRecordResponseBody
	GetCreateTimestamp() *int64
	SetLine(v string) *DescribeZoneRecordResponseBody
	GetLine() *string
	SetPriority(v int32) *DescribeZoneRecordResponseBody
	GetPriority() *int32
	SetRecordId(v int64) *DescribeZoneRecordResponseBody
	GetRecordId() *int64
	SetRemark(v string) *DescribeZoneRecordResponseBody
	GetRemark() *string
	SetRequestId(v string) *DescribeZoneRecordResponseBody
	GetRequestId() *string
	SetRr(v string) *DescribeZoneRecordResponseBody
	GetRr() *string
	SetStatus(v string) *DescribeZoneRecordResponseBody
	GetStatus() *string
	SetTtl(v int32) *DescribeZoneRecordResponseBody
	GetTtl() *int32
	SetType(v string) *DescribeZoneRecordResponseBody
	GetType() *string
	SetUpdateTime(v string) *DescribeZoneRecordResponseBody
	GetUpdateTime() *string
	SetUpdateTimestamp(v int64) *DescribeZoneRecordResponseBody
	GetUpdateTimestamp() *int64
	SetValue(v string) *DescribeZoneRecordResponseBody
	GetValue() *string
	SetWeight(v int32) *DescribeZoneRecordResponseBody
	GetWeight() *int32
	SetZoneId(v string) *DescribeZoneRecordResponseBody
	GetZoneId() *string
}

type DescribeZoneRecordResponseBody struct {
	// The time when the DNS record was created. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2018-01-23T03:15Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The time when the DNS record was created. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1516775741000
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
	// 5
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The ID of the DNS record.
	//
	// example:
	//
	// 5808
	RecordId *int64 `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
	// The description of the DNS record.
	//
	// example:
	//
	// test record
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0B7AD377-7E86-44A8-B9A8-53E8666E72FE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The hostname.
	//
	// example:
	//
	// www
	Rr *string `json:"Rr,omitempty" xml:"Rr,omitempty"`
	// The state of the DNS record. Valid values:
	//
	// 	- **ENABLE**: The DNS record is enabled.
	//
	// 	- **DISABLE**: The DNS record is disabled.
	//
	// example:
	//
	// ENABLE
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The time to live (TTL) of the DNS record.
	//
	// example:
	//
	// 60
	Ttl *int32 `json:"Ttl,omitempty" xml:"Ttl,omitempty"`
	// The type of the DNS record.
	//
	// example:
	//
	// A
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The time when the DNS record was updated. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2018-01-24T06:35Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// The time when the DNS record was updated. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1516775741000
	UpdateTimestamp *int64 `json:"UpdateTimestamp,omitempty" xml:"UpdateTimestamp,omitempty"`
	// The record value.
	//
	// example:
	//
	// 127.0.0.1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
	// The weight value of the DNS record.
	//
	// example:
	//
	// 1
	Weight *int32 `json:"Weight,omitempty" xml:"Weight,omitempty"`
	// The zone ID.
	//
	// example:
	//
	// CAgICA1OA_58
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeZoneRecordResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeZoneRecordResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeZoneRecordResponseBody) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeZoneRecordResponseBody) GetCreateTimestamp() *int64 {
	return s.CreateTimestamp
}

func (s *DescribeZoneRecordResponseBody) GetLine() *string {
	return s.Line
}

func (s *DescribeZoneRecordResponseBody) GetPriority() *int32 {
	return s.Priority
}

func (s *DescribeZoneRecordResponseBody) GetRecordId() *int64 {
	return s.RecordId
}

func (s *DescribeZoneRecordResponseBody) GetRemark() *string {
	return s.Remark
}

func (s *DescribeZoneRecordResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeZoneRecordResponseBody) GetRr() *string {
	return s.Rr
}

func (s *DescribeZoneRecordResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DescribeZoneRecordResponseBody) GetTtl() *int32 {
	return s.Ttl
}

func (s *DescribeZoneRecordResponseBody) GetType() *string {
	return s.Type
}

func (s *DescribeZoneRecordResponseBody) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *DescribeZoneRecordResponseBody) GetUpdateTimestamp() *int64 {
	return s.UpdateTimestamp
}

func (s *DescribeZoneRecordResponseBody) GetValue() *string {
	return s.Value
}

func (s *DescribeZoneRecordResponseBody) GetWeight() *int32 {
	return s.Weight
}

func (s *DescribeZoneRecordResponseBody) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeZoneRecordResponseBody) SetCreateTime(v string) *DescribeZoneRecordResponseBody {
	s.CreateTime = &v
	return s
}

func (s *DescribeZoneRecordResponseBody) SetCreateTimestamp(v int64) *DescribeZoneRecordResponseBody {
	s.CreateTimestamp = &v
	return s
}

func (s *DescribeZoneRecordResponseBody) SetLine(v string) *DescribeZoneRecordResponseBody {
	s.Line = &v
	return s
}

func (s *DescribeZoneRecordResponseBody) SetPriority(v int32) *DescribeZoneRecordResponseBody {
	s.Priority = &v
	return s
}

func (s *DescribeZoneRecordResponseBody) SetRecordId(v int64) *DescribeZoneRecordResponseBody {
	s.RecordId = &v
	return s
}

func (s *DescribeZoneRecordResponseBody) SetRemark(v string) *DescribeZoneRecordResponseBody {
	s.Remark = &v
	return s
}

func (s *DescribeZoneRecordResponseBody) SetRequestId(v string) *DescribeZoneRecordResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeZoneRecordResponseBody) SetRr(v string) *DescribeZoneRecordResponseBody {
	s.Rr = &v
	return s
}

func (s *DescribeZoneRecordResponseBody) SetStatus(v string) *DescribeZoneRecordResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeZoneRecordResponseBody) SetTtl(v int32) *DescribeZoneRecordResponseBody {
	s.Ttl = &v
	return s
}

func (s *DescribeZoneRecordResponseBody) SetType(v string) *DescribeZoneRecordResponseBody {
	s.Type = &v
	return s
}

func (s *DescribeZoneRecordResponseBody) SetUpdateTime(v string) *DescribeZoneRecordResponseBody {
	s.UpdateTime = &v
	return s
}

func (s *DescribeZoneRecordResponseBody) SetUpdateTimestamp(v int64) *DescribeZoneRecordResponseBody {
	s.UpdateTimestamp = &v
	return s
}

func (s *DescribeZoneRecordResponseBody) SetValue(v string) *DescribeZoneRecordResponseBody {
	s.Value = &v
	return s
}

func (s *DescribeZoneRecordResponseBody) SetWeight(v int32) *DescribeZoneRecordResponseBody {
	s.Weight = &v
	return s
}

func (s *DescribeZoneRecordResponseBody) SetZoneId(v string) *DescribeZoneRecordResponseBody {
	s.ZoneId = &v
	return s
}

func (s *DescribeZoneRecordResponseBody) Validate() error {
	return dara.Validate(s)
}
