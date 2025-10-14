// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchRecursionRecordsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *SearchRecursionRecordsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *SearchRecursionRecordsResponseBody
	GetNextToken() *string
	SetPageNumber(v int32) *SearchRecursionRecordsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *SearchRecursionRecordsResponseBody
	GetPageSize() *int32
	SetRecords(v *SearchRecursionRecordsResponseBodyRecords) *SearchRecursionRecordsResponseBody
	GetRecords() *SearchRecursionRecordsResponseBodyRecords
	SetRequestId(v string) *SearchRecursionRecordsResponseBody
	GetRequestId() *string
	SetTotalItems(v int32) *SearchRecursionRecordsResponseBody
	GetTotalItems() *int32
	SetTotalPages(v int32) *SearchRecursionRecordsResponseBody
	GetTotalPages() *int32
}

type SearchRecursionRecordsResponseBody struct {
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
	// 1
	PageSize *int32                                     `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Records  *SearchRecursionRecordsResponseBodyRecords `json:"Records,omitempty" xml:"Records,omitempty" type:"Struct"`
	// example:
	//
	// 536E9CAD-DB30-4647-AC87-AA5CC38C5382
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 100
	TotalItems *int32 `json:"TotalItems,omitempty" xml:"TotalItems,omitempty"`
	// example:
	//
	// 1
	TotalPages *int32 `json:"TotalPages,omitempty" xml:"TotalPages,omitempty"`
}

func (s SearchRecursionRecordsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SearchRecursionRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *SearchRecursionRecordsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *SearchRecursionRecordsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *SearchRecursionRecordsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *SearchRecursionRecordsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *SearchRecursionRecordsResponseBody) GetRecords() *SearchRecursionRecordsResponseBodyRecords {
	return s.Records
}

func (s *SearchRecursionRecordsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SearchRecursionRecordsResponseBody) GetTotalItems() *int32 {
	return s.TotalItems
}

func (s *SearchRecursionRecordsResponseBody) GetTotalPages() *int32 {
	return s.TotalPages
}

func (s *SearchRecursionRecordsResponseBody) SetMaxResults(v int32) *SearchRecursionRecordsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *SearchRecursionRecordsResponseBody) SetNextToken(v string) *SearchRecursionRecordsResponseBody {
	s.NextToken = &v
	return s
}

func (s *SearchRecursionRecordsResponseBody) SetPageNumber(v int32) *SearchRecursionRecordsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *SearchRecursionRecordsResponseBody) SetPageSize(v int32) *SearchRecursionRecordsResponseBody {
	s.PageSize = &v
	return s
}

func (s *SearchRecursionRecordsResponseBody) SetRecords(v *SearchRecursionRecordsResponseBodyRecords) *SearchRecursionRecordsResponseBody {
	s.Records = v
	return s
}

func (s *SearchRecursionRecordsResponseBody) SetRequestId(v string) *SearchRecursionRecordsResponseBody {
	s.RequestId = &v
	return s
}

func (s *SearchRecursionRecordsResponseBody) SetTotalItems(v int32) *SearchRecursionRecordsResponseBody {
	s.TotalItems = &v
	return s
}

func (s *SearchRecursionRecordsResponseBody) SetTotalPages(v int32) *SearchRecursionRecordsResponseBody {
	s.TotalPages = &v
	return s
}

func (s *SearchRecursionRecordsResponseBody) Validate() error {
	if s.Records != nil {
		if err := s.Records.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SearchRecursionRecordsResponseBodyRecords struct {
	Record []*SearchRecursionRecordsResponseBodyRecordsRecord `json:"Record,omitempty" xml:"Record,omitempty" type:"Repeated"`
}

func (s SearchRecursionRecordsResponseBodyRecords) String() string {
	return dara.Prettify(s)
}

func (s SearchRecursionRecordsResponseBodyRecords) GoString() string {
	return s.String()
}

func (s *SearchRecursionRecordsResponseBodyRecords) GetRecord() []*SearchRecursionRecordsResponseBodyRecordsRecord {
	return s.Record
}

func (s *SearchRecursionRecordsResponseBodyRecords) SetRecord(v []*SearchRecursionRecordsResponseBodyRecordsRecord) *SearchRecursionRecordsResponseBodyRecords {
	s.Record = v
	return s
}

func (s *SearchRecursionRecordsResponseBodyRecords) Validate() error {
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

type SearchRecursionRecordsResponseBodyRecordsRecord struct {
	// example:
	//
	// 2025-06-16T02:18Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 1615182315000
	CreateTimestamp *int64 `json:"CreateTimestamp,omitempty" xml:"CreateTimestamp,omitempty"`
	// example:
	//
	// 1521021611066866
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
	// 1
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// record ID
	//
	// example:
	//
	// 1927184083990351872
	RecordId *string `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
	// example:
	//
	// vxwk
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
	// 2022-10-28T13:07Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// example:
	//
	// 1722107570000
	UpdateTimestamp *int64 `json:"UpdateTimestamp,omitempty" xml:"UpdateTimestamp,omitempty"`
	// example:
	//
	// 14.19.85.71
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
	// example:
	//
	// 4
	Weight *string `json:"Weight,omitempty" xml:"Weight,omitempty"`
	// zone ID
	//
	// example:
	//
	// 169783221000012
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s SearchRecursionRecordsResponseBodyRecordsRecord) String() string {
	return dara.Prettify(s)
}

func (s SearchRecursionRecordsResponseBodyRecordsRecord) GoString() string {
	return s.String()
}

func (s *SearchRecursionRecordsResponseBodyRecordsRecord) GetCreateTime() *string {
	return s.CreateTime
}

func (s *SearchRecursionRecordsResponseBodyRecordsRecord) GetCreateTimestamp() *int64 {
	return s.CreateTimestamp
}

func (s *SearchRecursionRecordsResponseBodyRecordsRecord) GetCreator() *int64 {
	return s.Creator
}

func (s *SearchRecursionRecordsResponseBodyRecordsRecord) GetCreatorSubType() *string {
	return s.CreatorSubType
}

func (s *SearchRecursionRecordsResponseBodyRecordsRecord) GetCreatorType() *string {
	return s.CreatorType
}

func (s *SearchRecursionRecordsResponseBodyRecordsRecord) GetEnableStatus() *string {
	return s.EnableStatus
}

func (s *SearchRecursionRecordsResponseBodyRecordsRecord) GetPriority() *int32 {
	return s.Priority
}

func (s *SearchRecursionRecordsResponseBodyRecordsRecord) GetRecordId() *string {
	return s.RecordId
}

func (s *SearchRecursionRecordsResponseBodyRecordsRecord) GetRemark() *string {
	return s.Remark
}

func (s *SearchRecursionRecordsResponseBodyRecordsRecord) GetRequestSource() *string {
	return s.RequestSource
}

func (s *SearchRecursionRecordsResponseBodyRecordsRecord) GetRr() *string {
	return s.Rr
}

func (s *SearchRecursionRecordsResponseBodyRecordsRecord) GetTtl() *int32 {
	return s.Ttl
}

func (s *SearchRecursionRecordsResponseBodyRecordsRecord) GetType() *string {
	return s.Type
}

func (s *SearchRecursionRecordsResponseBodyRecordsRecord) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *SearchRecursionRecordsResponseBodyRecordsRecord) GetUpdateTimestamp() *int64 {
	return s.UpdateTimestamp
}

func (s *SearchRecursionRecordsResponseBodyRecordsRecord) GetValue() *string {
	return s.Value
}

func (s *SearchRecursionRecordsResponseBodyRecordsRecord) GetWeight() *string {
	return s.Weight
}

func (s *SearchRecursionRecordsResponseBodyRecordsRecord) GetZoneId() *string {
	return s.ZoneId
}

func (s *SearchRecursionRecordsResponseBodyRecordsRecord) SetCreateTime(v string) *SearchRecursionRecordsResponseBodyRecordsRecord {
	s.CreateTime = &v
	return s
}

func (s *SearchRecursionRecordsResponseBodyRecordsRecord) SetCreateTimestamp(v int64) *SearchRecursionRecordsResponseBodyRecordsRecord {
	s.CreateTimestamp = &v
	return s
}

func (s *SearchRecursionRecordsResponseBodyRecordsRecord) SetCreator(v int64) *SearchRecursionRecordsResponseBodyRecordsRecord {
	s.Creator = &v
	return s
}

func (s *SearchRecursionRecordsResponseBodyRecordsRecord) SetCreatorSubType(v string) *SearchRecursionRecordsResponseBodyRecordsRecord {
	s.CreatorSubType = &v
	return s
}

func (s *SearchRecursionRecordsResponseBodyRecordsRecord) SetCreatorType(v string) *SearchRecursionRecordsResponseBodyRecordsRecord {
	s.CreatorType = &v
	return s
}

func (s *SearchRecursionRecordsResponseBodyRecordsRecord) SetEnableStatus(v string) *SearchRecursionRecordsResponseBodyRecordsRecord {
	s.EnableStatus = &v
	return s
}

func (s *SearchRecursionRecordsResponseBodyRecordsRecord) SetPriority(v int32) *SearchRecursionRecordsResponseBodyRecordsRecord {
	s.Priority = &v
	return s
}

func (s *SearchRecursionRecordsResponseBodyRecordsRecord) SetRecordId(v string) *SearchRecursionRecordsResponseBodyRecordsRecord {
	s.RecordId = &v
	return s
}

func (s *SearchRecursionRecordsResponseBodyRecordsRecord) SetRemark(v string) *SearchRecursionRecordsResponseBodyRecordsRecord {
	s.Remark = &v
	return s
}

func (s *SearchRecursionRecordsResponseBodyRecordsRecord) SetRequestSource(v string) *SearchRecursionRecordsResponseBodyRecordsRecord {
	s.RequestSource = &v
	return s
}

func (s *SearchRecursionRecordsResponseBodyRecordsRecord) SetRr(v string) *SearchRecursionRecordsResponseBodyRecordsRecord {
	s.Rr = &v
	return s
}

func (s *SearchRecursionRecordsResponseBodyRecordsRecord) SetTtl(v int32) *SearchRecursionRecordsResponseBodyRecordsRecord {
	s.Ttl = &v
	return s
}

func (s *SearchRecursionRecordsResponseBodyRecordsRecord) SetType(v string) *SearchRecursionRecordsResponseBodyRecordsRecord {
	s.Type = &v
	return s
}

func (s *SearchRecursionRecordsResponseBodyRecordsRecord) SetUpdateTime(v string) *SearchRecursionRecordsResponseBodyRecordsRecord {
	s.UpdateTime = &v
	return s
}

func (s *SearchRecursionRecordsResponseBodyRecordsRecord) SetUpdateTimestamp(v int64) *SearchRecursionRecordsResponseBodyRecordsRecord {
	s.UpdateTimestamp = &v
	return s
}

func (s *SearchRecursionRecordsResponseBodyRecordsRecord) SetValue(v string) *SearchRecursionRecordsResponseBodyRecordsRecord {
	s.Value = &v
	return s
}

func (s *SearchRecursionRecordsResponseBodyRecordsRecord) SetWeight(v string) *SearchRecursionRecordsResponseBodyRecordsRecord {
	s.Weight = &v
	return s
}

func (s *SearchRecursionRecordsResponseBodyRecordsRecord) SetZoneId(v string) *SearchRecursionRecordsResponseBodyRecordsRecord {
	s.ZoneId = &v
	return s
}

func (s *SearchRecursionRecordsResponseBodyRecordsRecord) Validate() error {
	return dara.Validate(s)
}
