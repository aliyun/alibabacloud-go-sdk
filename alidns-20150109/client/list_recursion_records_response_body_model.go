// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRecursionRecordsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListRecursionRecordsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListRecursionRecordsResponseBody
	GetNextToken() *string
	SetPageNumber(v int32) *ListRecursionRecordsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListRecursionRecordsResponseBody
	GetPageSize() *int32
	SetRecords(v *ListRecursionRecordsResponseBodyRecords) *ListRecursionRecordsResponseBody
	GetRecords() *ListRecursionRecordsResponseBodyRecords
	SetRequestId(v string) *ListRecursionRecordsResponseBody
	GetRequestId() *string
	SetTotalItems(v int32) *ListRecursionRecordsResponseBody
	GetTotalItems() *int32
	SetTotalPages(v int32) *ListRecursionRecordsResponseBody
	GetTotalPages() *int32
}

type ListRecursionRecordsResponseBody struct {
	// example:
	//
	// 500
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// 4698691
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 20
	PageSize *int32                                   `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Records  *ListRecursionRecordsResponseBodyRecords `json:"Records,omitempty" xml:"Records,omitempty" type:"Struct"`
	// example:
	//
	// 6856BCF6-11D6-4D7E-AC53-FD579933522B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 100
	TotalItems *int32 `json:"TotalItems,omitempty" xml:"TotalItems,omitempty"`
	// example:
	//
	// 5
	TotalPages *int32 `json:"TotalPages,omitempty" xml:"TotalPages,omitempty"`
}

func (s ListRecursionRecordsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListRecursionRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *ListRecursionRecordsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListRecursionRecordsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListRecursionRecordsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListRecursionRecordsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListRecursionRecordsResponseBody) GetRecords() *ListRecursionRecordsResponseBodyRecords {
	return s.Records
}

func (s *ListRecursionRecordsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListRecursionRecordsResponseBody) GetTotalItems() *int32 {
	return s.TotalItems
}

func (s *ListRecursionRecordsResponseBody) GetTotalPages() *int32 {
	return s.TotalPages
}

func (s *ListRecursionRecordsResponseBody) SetMaxResults(v int32) *ListRecursionRecordsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListRecursionRecordsResponseBody) SetNextToken(v string) *ListRecursionRecordsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListRecursionRecordsResponseBody) SetPageNumber(v int32) *ListRecursionRecordsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListRecursionRecordsResponseBody) SetPageSize(v int32) *ListRecursionRecordsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListRecursionRecordsResponseBody) SetRecords(v *ListRecursionRecordsResponseBodyRecords) *ListRecursionRecordsResponseBody {
	s.Records = v
	return s
}

func (s *ListRecursionRecordsResponseBody) SetRequestId(v string) *ListRecursionRecordsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRecursionRecordsResponseBody) SetTotalItems(v int32) *ListRecursionRecordsResponseBody {
	s.TotalItems = &v
	return s
}

func (s *ListRecursionRecordsResponseBody) SetTotalPages(v int32) *ListRecursionRecordsResponseBody {
	s.TotalPages = &v
	return s
}

func (s *ListRecursionRecordsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListRecursionRecordsResponseBodyRecords struct {
	Record []*ListRecursionRecordsResponseBodyRecordsRecord `json:"Record,omitempty" xml:"Record,omitempty" type:"Repeated"`
}

func (s ListRecursionRecordsResponseBodyRecords) String() string {
	return dara.Prettify(s)
}

func (s ListRecursionRecordsResponseBodyRecords) GoString() string {
	return s.String()
}

func (s *ListRecursionRecordsResponseBodyRecords) GetRecord() []*ListRecursionRecordsResponseBodyRecordsRecord {
	return s.Record
}

func (s *ListRecursionRecordsResponseBodyRecords) SetRecord(v []*ListRecursionRecordsResponseBodyRecordsRecord) *ListRecursionRecordsResponseBodyRecords {
	s.Record = v
	return s
}

func (s *ListRecursionRecordsResponseBodyRecords) Validate() error {
	return dara.Validate(s)
}

type ListRecursionRecordsResponseBodyRecordsRecord struct {
	// example:
	//
	// 2023-06-28T09:16Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 1721119172000
	CreateTimestamp *int64 `json:"CreateTimestamp,omitempty" xml:"CreateTimestamp,omitempty"`
	// example:
	//
	// 123121312
	Creator *int64 `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// example:
	//
	// custom
	CreatorSubType *string `json:"CreatorSubType,omitempty" xml:"CreatorSubType,omitempty"`
	// example:
	//
	// user
	CreatorType *string `json:"CreatorType,omitempty" xml:"CreatorType,omitempty"`
	// example:
	//
	// enable
	EnableStatus *string `json:"EnableStatus,omitempty" xml:"EnableStatus,omitempty"`
	// example:
	//
	// 10
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// example:
	//
	// 1897105298264316928
	RecordId *string `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
	// example:
	//
	// 1150
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// example:
	//
	// default
	RequestSource *string `json:"RequestSource,omitempty" xml:"RequestSource,omitempty"`
	// example:
	//
	// wap
	Rr *string `json:"Rr,omitempty" xml:"Rr,omitempty"`
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
	// 2024-10-22T09:54Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// example:
	//
	// 1620283051000
	UpdateTimestamp *int64 `json:"UpdateTimestamp,omitempty" xml:"UpdateTimestamp,omitempty"`
	// example:
	//
	// 120.38.70.54
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
	// example:
	//
	// 1
	Weight *int32 `json:"Weight,omitempty" xml:"Weight,omitempty"`
	// example:
	//
	// 169439170000011
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s ListRecursionRecordsResponseBodyRecordsRecord) String() string {
	return dara.Prettify(s)
}

func (s ListRecursionRecordsResponseBodyRecordsRecord) GoString() string {
	return s.String()
}

func (s *ListRecursionRecordsResponseBodyRecordsRecord) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListRecursionRecordsResponseBodyRecordsRecord) GetCreateTimestamp() *int64 {
	return s.CreateTimestamp
}

func (s *ListRecursionRecordsResponseBodyRecordsRecord) GetCreator() *int64 {
	return s.Creator
}

func (s *ListRecursionRecordsResponseBodyRecordsRecord) GetCreatorSubType() *string {
	return s.CreatorSubType
}

func (s *ListRecursionRecordsResponseBodyRecordsRecord) GetCreatorType() *string {
	return s.CreatorType
}

func (s *ListRecursionRecordsResponseBodyRecordsRecord) GetEnableStatus() *string {
	return s.EnableStatus
}

func (s *ListRecursionRecordsResponseBodyRecordsRecord) GetPriority() *int32 {
	return s.Priority
}

func (s *ListRecursionRecordsResponseBodyRecordsRecord) GetRecordId() *string {
	return s.RecordId
}

func (s *ListRecursionRecordsResponseBodyRecordsRecord) GetRemark() *string {
	return s.Remark
}

func (s *ListRecursionRecordsResponseBodyRecordsRecord) GetRequestSource() *string {
	return s.RequestSource
}

func (s *ListRecursionRecordsResponseBodyRecordsRecord) GetRr() *string {
	return s.Rr
}

func (s *ListRecursionRecordsResponseBodyRecordsRecord) GetTtl() *int32 {
	return s.Ttl
}

func (s *ListRecursionRecordsResponseBodyRecordsRecord) GetType() *string {
	return s.Type
}

func (s *ListRecursionRecordsResponseBodyRecordsRecord) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *ListRecursionRecordsResponseBodyRecordsRecord) GetUpdateTimestamp() *int64 {
	return s.UpdateTimestamp
}

func (s *ListRecursionRecordsResponseBodyRecordsRecord) GetValue() *string {
	return s.Value
}

func (s *ListRecursionRecordsResponseBodyRecordsRecord) GetWeight() *int32 {
	return s.Weight
}

func (s *ListRecursionRecordsResponseBodyRecordsRecord) GetZoneId() *string {
	return s.ZoneId
}

func (s *ListRecursionRecordsResponseBodyRecordsRecord) SetCreateTime(v string) *ListRecursionRecordsResponseBodyRecordsRecord {
	s.CreateTime = &v
	return s
}

func (s *ListRecursionRecordsResponseBodyRecordsRecord) SetCreateTimestamp(v int64) *ListRecursionRecordsResponseBodyRecordsRecord {
	s.CreateTimestamp = &v
	return s
}

func (s *ListRecursionRecordsResponseBodyRecordsRecord) SetCreator(v int64) *ListRecursionRecordsResponseBodyRecordsRecord {
	s.Creator = &v
	return s
}

func (s *ListRecursionRecordsResponseBodyRecordsRecord) SetCreatorSubType(v string) *ListRecursionRecordsResponseBodyRecordsRecord {
	s.CreatorSubType = &v
	return s
}

func (s *ListRecursionRecordsResponseBodyRecordsRecord) SetCreatorType(v string) *ListRecursionRecordsResponseBodyRecordsRecord {
	s.CreatorType = &v
	return s
}

func (s *ListRecursionRecordsResponseBodyRecordsRecord) SetEnableStatus(v string) *ListRecursionRecordsResponseBodyRecordsRecord {
	s.EnableStatus = &v
	return s
}

func (s *ListRecursionRecordsResponseBodyRecordsRecord) SetPriority(v int32) *ListRecursionRecordsResponseBodyRecordsRecord {
	s.Priority = &v
	return s
}

func (s *ListRecursionRecordsResponseBodyRecordsRecord) SetRecordId(v string) *ListRecursionRecordsResponseBodyRecordsRecord {
	s.RecordId = &v
	return s
}

func (s *ListRecursionRecordsResponseBodyRecordsRecord) SetRemark(v string) *ListRecursionRecordsResponseBodyRecordsRecord {
	s.Remark = &v
	return s
}

func (s *ListRecursionRecordsResponseBodyRecordsRecord) SetRequestSource(v string) *ListRecursionRecordsResponseBodyRecordsRecord {
	s.RequestSource = &v
	return s
}

func (s *ListRecursionRecordsResponseBodyRecordsRecord) SetRr(v string) *ListRecursionRecordsResponseBodyRecordsRecord {
	s.Rr = &v
	return s
}

func (s *ListRecursionRecordsResponseBodyRecordsRecord) SetTtl(v int32) *ListRecursionRecordsResponseBodyRecordsRecord {
	s.Ttl = &v
	return s
}

func (s *ListRecursionRecordsResponseBodyRecordsRecord) SetType(v string) *ListRecursionRecordsResponseBodyRecordsRecord {
	s.Type = &v
	return s
}

func (s *ListRecursionRecordsResponseBodyRecordsRecord) SetUpdateTime(v string) *ListRecursionRecordsResponseBodyRecordsRecord {
	s.UpdateTime = &v
	return s
}

func (s *ListRecursionRecordsResponseBodyRecordsRecord) SetUpdateTimestamp(v int64) *ListRecursionRecordsResponseBodyRecordsRecord {
	s.UpdateTimestamp = &v
	return s
}

func (s *ListRecursionRecordsResponseBodyRecordsRecord) SetValue(v string) *ListRecursionRecordsResponseBodyRecordsRecord {
	s.Value = &v
	return s
}

func (s *ListRecursionRecordsResponseBodyRecordsRecord) SetWeight(v int32) *ListRecursionRecordsResponseBodyRecordsRecord {
	s.Weight = &v
	return s
}

func (s *ListRecursionRecordsResponseBodyRecordsRecord) SetZoneId(v string) *ListRecursionRecordsResponseBodyRecordsRecord {
	s.ZoneId = &v
	return s
}

func (s *ListRecursionRecordsResponseBodyRecordsRecord) Validate() error {
	return dara.Validate(s)
}
