// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRecursionRecordResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCreateTime(v string) *DescribeRecursionRecordResponseBody
	GetCreateTime() *string
	SetCreateTimestamp(v int64) *DescribeRecursionRecordResponseBody
	GetCreateTimestamp() *int64
	SetPriority(v int32) *DescribeRecursionRecordResponseBody
	GetPriority() *int32
	SetRecordId(v string) *DescribeRecursionRecordResponseBody
	GetRecordId() *string
	SetRemark(v string) *DescribeRecursionRecordResponseBody
	GetRemark() *string
	SetRequestId(v string) *DescribeRecursionRecordResponseBody
	GetRequestId() *string
	SetRr(v string) *DescribeRecursionRecordResponseBody
	GetRr() *string
	SetTtl(v int32) *DescribeRecursionRecordResponseBody
	GetTtl() *int32
	SetType(v string) *DescribeRecursionRecordResponseBody
	GetType() *string
	SetUpdateTime(v string) *DescribeRecursionRecordResponseBody
	GetUpdateTime() *string
	SetUpdateTimestamp(v int64) *DescribeRecursionRecordResponseBody
	GetUpdateTimestamp() *int64
	SetValue(v string) *DescribeRecursionRecordResponseBody
	GetValue() *string
	SetWeight(v int32) *DescribeRecursionRecordResponseBody
	GetWeight() *int32
	SetZoneId(v string) *DescribeRecursionRecordResponseBody
	GetZoneId() *string
}

type DescribeRecursionRecordResponseBody struct {
	// example:
	//
	// 2024-03-23T13:09Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 1527690629357
	CreateTimestamp *int64 `json:"CreateTimestamp,omitempty" xml:"CreateTimestamp,omitempty"`
	// example:
	//
	// 1
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// record ID
	//
	// example:
	//
	// 1903634198947606528
	RecordId *string `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
	// example:
	//
	// test
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// example:
	//
	// 6856BCF6-11D6-4D7E-AC53-FD579933522B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// test
	Rr *string `json:"Rr,omitempty" xml:"Rr,omitempty"`
	// ttl
	//
	// example:
	//
	// 60
	Ttl *int32 `json:"Ttl,omitempty" xml:"Ttl,omitempty"`
	// example:
	//
	// A
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// 2024-03-29T13:20Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// example:
	//
	// 1527690629357
	UpdateTimestamp *int64 `json:"UpdateTimestamp,omitempty" xml:"UpdateTimestamp,omitempty"`
	// example:
	//
	// 1.1.XX.XX
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
	// example:
	//
	// 2
	Weight *int32 `json:"Weight,omitempty" xml:"Weight,omitempty"`
	// zone ID
	//
	// example:
	//
	// 172635074000011
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeRecursionRecordResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRecursionRecordResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRecursionRecordResponseBody) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeRecursionRecordResponseBody) GetCreateTimestamp() *int64 {
	return s.CreateTimestamp
}

func (s *DescribeRecursionRecordResponseBody) GetPriority() *int32 {
	return s.Priority
}

func (s *DescribeRecursionRecordResponseBody) GetRecordId() *string {
	return s.RecordId
}

func (s *DescribeRecursionRecordResponseBody) GetRemark() *string {
	return s.Remark
}

func (s *DescribeRecursionRecordResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRecursionRecordResponseBody) GetRr() *string {
	return s.Rr
}

func (s *DescribeRecursionRecordResponseBody) GetTtl() *int32 {
	return s.Ttl
}

func (s *DescribeRecursionRecordResponseBody) GetType() *string {
	return s.Type
}

func (s *DescribeRecursionRecordResponseBody) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *DescribeRecursionRecordResponseBody) GetUpdateTimestamp() *int64 {
	return s.UpdateTimestamp
}

func (s *DescribeRecursionRecordResponseBody) GetValue() *string {
	return s.Value
}

func (s *DescribeRecursionRecordResponseBody) GetWeight() *int32 {
	return s.Weight
}

func (s *DescribeRecursionRecordResponseBody) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeRecursionRecordResponseBody) SetCreateTime(v string) *DescribeRecursionRecordResponseBody {
	s.CreateTime = &v
	return s
}

func (s *DescribeRecursionRecordResponseBody) SetCreateTimestamp(v int64) *DescribeRecursionRecordResponseBody {
	s.CreateTimestamp = &v
	return s
}

func (s *DescribeRecursionRecordResponseBody) SetPriority(v int32) *DescribeRecursionRecordResponseBody {
	s.Priority = &v
	return s
}

func (s *DescribeRecursionRecordResponseBody) SetRecordId(v string) *DescribeRecursionRecordResponseBody {
	s.RecordId = &v
	return s
}

func (s *DescribeRecursionRecordResponseBody) SetRemark(v string) *DescribeRecursionRecordResponseBody {
	s.Remark = &v
	return s
}

func (s *DescribeRecursionRecordResponseBody) SetRequestId(v string) *DescribeRecursionRecordResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRecursionRecordResponseBody) SetRr(v string) *DescribeRecursionRecordResponseBody {
	s.Rr = &v
	return s
}

func (s *DescribeRecursionRecordResponseBody) SetTtl(v int32) *DescribeRecursionRecordResponseBody {
	s.Ttl = &v
	return s
}

func (s *DescribeRecursionRecordResponseBody) SetType(v string) *DescribeRecursionRecordResponseBody {
	s.Type = &v
	return s
}

func (s *DescribeRecursionRecordResponseBody) SetUpdateTime(v string) *DescribeRecursionRecordResponseBody {
	s.UpdateTime = &v
	return s
}

func (s *DescribeRecursionRecordResponseBody) SetUpdateTimestamp(v int64) *DescribeRecursionRecordResponseBody {
	s.UpdateTimestamp = &v
	return s
}

func (s *DescribeRecursionRecordResponseBody) SetValue(v string) *DescribeRecursionRecordResponseBody {
	s.Value = &v
	return s
}

func (s *DescribeRecursionRecordResponseBody) SetWeight(v int32) *DescribeRecursionRecordResponseBody {
	s.Weight = &v
	return s
}

func (s *DescribeRecursionRecordResponseBody) SetZoneId(v string) *DescribeRecursionRecordResponseBody {
	s.ZoneId = &v
	return s
}

func (s *DescribeRecursionRecordResponseBody) Validate() error {
	return dara.Validate(s)
}
