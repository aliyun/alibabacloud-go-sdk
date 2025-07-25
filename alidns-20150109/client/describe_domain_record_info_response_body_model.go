// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainRecordInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDomainId(v string) *DescribeDomainRecordInfoResponseBody
	GetDomainId() *string
	SetDomainName(v string) *DescribeDomainRecordInfoResponseBody
	GetDomainName() *string
	SetGroupId(v string) *DescribeDomainRecordInfoResponseBody
	GetGroupId() *string
	SetGroupName(v string) *DescribeDomainRecordInfoResponseBody
	GetGroupName() *string
	SetLine(v string) *DescribeDomainRecordInfoResponseBody
	GetLine() *string
	SetLocked(v bool) *DescribeDomainRecordInfoResponseBody
	GetLocked() *bool
	SetPriority(v int64) *DescribeDomainRecordInfoResponseBody
	GetPriority() *int64
	SetPunyCode(v string) *DescribeDomainRecordInfoResponseBody
	GetPunyCode() *string
	SetRR(v string) *DescribeDomainRecordInfoResponseBody
	GetRR() *string
	SetRecordId(v string) *DescribeDomainRecordInfoResponseBody
	GetRecordId() *string
	SetRemark(v string) *DescribeDomainRecordInfoResponseBody
	GetRemark() *string
	SetRequestId(v string) *DescribeDomainRecordInfoResponseBody
	GetRequestId() *string
	SetStatus(v string) *DescribeDomainRecordInfoResponseBody
	GetStatus() *string
	SetTTL(v int64) *DescribeDomainRecordInfoResponseBody
	GetTTL() *int64
	SetType(v string) *DescribeDomainRecordInfoResponseBody
	GetType() *string
	SetValue(v string) *DescribeDomainRecordInfoResponseBody
	GetValue() *string
}

type DescribeDomainRecordInfoResponseBody struct {
	// The ID of the domain name.
	//
	// example:
	//
	// 00efd71a-770e-4255-b54e-6fe5659baffe
	DomainId *string `json:"DomainId,omitempty" xml:"DomainId,omitempty"`
	// The domain name.
	//
	// example:
	//
	// dns-example.top
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The ID of the domain name group.
	//
	// example:
	//
	// 2223
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The name of the domain name group.
	//
	// example:
	//
	// MyGroup
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The DNS resolution line.
	//
	// example:
	//
	// default
	Line *string `json:"Line,omitempty" xml:"Line,omitempty"`
	// The lock state of the DNS record. Valid values: **true and false**.
	//
	// example:
	//
	// true
	Locked *bool `json:"Locked,omitempty" xml:"Locked,omitempty"`
	// The priority of the mail exchanger (MX) record.
	//
	// example:
	//
	// 5
	Priority *int64 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The Punycode for the domain name. This parameter is returned only for Chinese domain names.
	//
	// example:
	//
	// xn--fsq270a.com
	PunyCode *string `json:"PunyCode,omitempty" xml:"PunyCode,omitempty"`
	// The hostname.
	//
	// example:
	//
	// @
	RR *string `json:"RR,omitempty" xml:"RR,omitempty"`
	// The ID of the DNS record.
	//
	// example:
	//
	// 9999985
	RecordId *string `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
	// The description of your DNS record.
	//
	// example:
	//
	// test
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 536E9CAD-DB30-4647-AC87-AA5CC38C5382
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The status of the DNS record. Valid values:
	//
	// Enable: enabled
	//
	// Disable: disabled
	//
	// example:
	//
	// Enable
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The time for which the DNS record is cached in a local DNS system.
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
	// The record value.
	//
	// example:
	//
	// 1.1.XX.XX
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDomainRecordInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainRecordInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainRecordInfoResponseBody) GetDomainId() *string {
	return s.DomainId
}

func (s *DescribeDomainRecordInfoResponseBody) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDomainRecordInfoResponseBody) GetGroupId() *string {
	return s.GroupId
}

func (s *DescribeDomainRecordInfoResponseBody) GetGroupName() *string {
	return s.GroupName
}

func (s *DescribeDomainRecordInfoResponseBody) GetLine() *string {
	return s.Line
}

func (s *DescribeDomainRecordInfoResponseBody) GetLocked() *bool {
	return s.Locked
}

func (s *DescribeDomainRecordInfoResponseBody) GetPriority() *int64 {
	return s.Priority
}

func (s *DescribeDomainRecordInfoResponseBody) GetPunyCode() *string {
	return s.PunyCode
}

func (s *DescribeDomainRecordInfoResponseBody) GetRR() *string {
	return s.RR
}

func (s *DescribeDomainRecordInfoResponseBody) GetRecordId() *string {
	return s.RecordId
}

func (s *DescribeDomainRecordInfoResponseBody) GetRemark() *string {
	return s.Remark
}

func (s *DescribeDomainRecordInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDomainRecordInfoResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DescribeDomainRecordInfoResponseBody) GetTTL() *int64 {
	return s.TTL
}

func (s *DescribeDomainRecordInfoResponseBody) GetType() *string {
	return s.Type
}

func (s *DescribeDomainRecordInfoResponseBody) GetValue() *string {
	return s.Value
}

func (s *DescribeDomainRecordInfoResponseBody) SetDomainId(v string) *DescribeDomainRecordInfoResponseBody {
	s.DomainId = &v
	return s
}

func (s *DescribeDomainRecordInfoResponseBody) SetDomainName(v string) *DescribeDomainRecordInfoResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainRecordInfoResponseBody) SetGroupId(v string) *DescribeDomainRecordInfoResponseBody {
	s.GroupId = &v
	return s
}

func (s *DescribeDomainRecordInfoResponseBody) SetGroupName(v string) *DescribeDomainRecordInfoResponseBody {
	s.GroupName = &v
	return s
}

func (s *DescribeDomainRecordInfoResponseBody) SetLine(v string) *DescribeDomainRecordInfoResponseBody {
	s.Line = &v
	return s
}

func (s *DescribeDomainRecordInfoResponseBody) SetLocked(v bool) *DescribeDomainRecordInfoResponseBody {
	s.Locked = &v
	return s
}

func (s *DescribeDomainRecordInfoResponseBody) SetPriority(v int64) *DescribeDomainRecordInfoResponseBody {
	s.Priority = &v
	return s
}

func (s *DescribeDomainRecordInfoResponseBody) SetPunyCode(v string) *DescribeDomainRecordInfoResponseBody {
	s.PunyCode = &v
	return s
}

func (s *DescribeDomainRecordInfoResponseBody) SetRR(v string) *DescribeDomainRecordInfoResponseBody {
	s.RR = &v
	return s
}

func (s *DescribeDomainRecordInfoResponseBody) SetRecordId(v string) *DescribeDomainRecordInfoResponseBody {
	s.RecordId = &v
	return s
}

func (s *DescribeDomainRecordInfoResponseBody) SetRemark(v string) *DescribeDomainRecordInfoResponseBody {
	s.Remark = &v
	return s
}

func (s *DescribeDomainRecordInfoResponseBody) SetRequestId(v string) *DescribeDomainRecordInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainRecordInfoResponseBody) SetStatus(v string) *DescribeDomainRecordInfoResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeDomainRecordInfoResponseBody) SetTTL(v int64) *DescribeDomainRecordInfoResponseBody {
	s.TTL = &v
	return s
}

func (s *DescribeDomainRecordInfoResponseBody) SetType(v string) *DescribeDomainRecordInfoResponseBody {
	s.Type = &v
	return s
}

func (s *DescribeDomainRecordInfoResponseBody) SetValue(v string) *DescribeDomainRecordInfoResponseBody {
	s.Value = &v
	return s
}

func (s *DescribeDomainRecordInfoResponseBody) Validate() error {
	return dara.Validate(s)
}
