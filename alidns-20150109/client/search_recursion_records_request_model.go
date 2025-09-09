// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchRecursionRecordsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDirection(v string) *SearchRecursionRecordsRequest
	GetDirection() *string
	SetEnableStatus(v string) *SearchRecursionRecordsRequest
	GetEnableStatus() *string
	SetMaxResults(v int32) *SearchRecursionRecordsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *SearchRecursionRecordsRequest
	GetNextToken() *string
	SetOrderBy(v string) *SearchRecursionRecordsRequest
	GetOrderBy() *string
	SetPageNumber(v int32) *SearchRecursionRecordsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *SearchRecursionRecordsRequest
	GetPageSize() *int32
	SetRemark(v string) *SearchRecursionRecordsRequest
	GetRemark() *string
	SetRequestSource(v string) *SearchRecursionRecordsRequest
	GetRequestSource() *string
	SetRr(v string) *SearchRecursionRecordsRequest
	GetRr() *string
	SetTtl(v int32) *SearchRecursionRecordsRequest
	GetTtl() *int32
	SetType(v string) *SearchRecursionRecordsRequest
	GetType() *string
	SetValue(v string) *SearchRecursionRecordsRequest
	GetValue() *string
	SetWeight(v int32) *SearchRecursionRecordsRequest
	GetWeight() *int32
	SetZoneId(v string) *SearchRecursionRecordsRequest
	GetZoneId() *string
}

type SearchRecursionRecordsRequest struct {
	// example:
	//
	// asc
	Direction *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	// example:
	//
	// enable
	EnableStatus *string `json:"EnableStatus,omitempty" xml:"EnableStatus,omitempty"`
	// example:
	//
	// 100
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// 4698691
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// rr
	OrderBy *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// test
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// example:
	//
	// default
	RequestSource *string `json:"RequestSource,omitempty" xml:"RequestSource,omitempty"`
	// example:
	//
	// www
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
	// 1.1.XX.XX
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
	// example:
	//
	// 2
	Weight *int32 `json:"Weight,omitempty" xml:"Weight,omitempty"`
	// example:
	//
	// 169438909000011
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s SearchRecursionRecordsRequest) String() string {
	return dara.Prettify(s)
}

func (s SearchRecursionRecordsRequest) GoString() string {
	return s.String()
}

func (s *SearchRecursionRecordsRequest) GetDirection() *string {
	return s.Direction
}

func (s *SearchRecursionRecordsRequest) GetEnableStatus() *string {
	return s.EnableStatus
}

func (s *SearchRecursionRecordsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *SearchRecursionRecordsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *SearchRecursionRecordsRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *SearchRecursionRecordsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *SearchRecursionRecordsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *SearchRecursionRecordsRequest) GetRemark() *string {
	return s.Remark
}

func (s *SearchRecursionRecordsRequest) GetRequestSource() *string {
	return s.RequestSource
}

func (s *SearchRecursionRecordsRequest) GetRr() *string {
	return s.Rr
}

func (s *SearchRecursionRecordsRequest) GetTtl() *int32 {
	return s.Ttl
}

func (s *SearchRecursionRecordsRequest) GetType() *string {
	return s.Type
}

func (s *SearchRecursionRecordsRequest) GetValue() *string {
	return s.Value
}

func (s *SearchRecursionRecordsRequest) GetWeight() *int32 {
	return s.Weight
}

func (s *SearchRecursionRecordsRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *SearchRecursionRecordsRequest) SetDirection(v string) *SearchRecursionRecordsRequest {
	s.Direction = &v
	return s
}

func (s *SearchRecursionRecordsRequest) SetEnableStatus(v string) *SearchRecursionRecordsRequest {
	s.EnableStatus = &v
	return s
}

func (s *SearchRecursionRecordsRequest) SetMaxResults(v int32) *SearchRecursionRecordsRequest {
	s.MaxResults = &v
	return s
}

func (s *SearchRecursionRecordsRequest) SetNextToken(v string) *SearchRecursionRecordsRequest {
	s.NextToken = &v
	return s
}

func (s *SearchRecursionRecordsRequest) SetOrderBy(v string) *SearchRecursionRecordsRequest {
	s.OrderBy = &v
	return s
}

func (s *SearchRecursionRecordsRequest) SetPageNumber(v int32) *SearchRecursionRecordsRequest {
	s.PageNumber = &v
	return s
}

func (s *SearchRecursionRecordsRequest) SetPageSize(v int32) *SearchRecursionRecordsRequest {
	s.PageSize = &v
	return s
}

func (s *SearchRecursionRecordsRequest) SetRemark(v string) *SearchRecursionRecordsRequest {
	s.Remark = &v
	return s
}

func (s *SearchRecursionRecordsRequest) SetRequestSource(v string) *SearchRecursionRecordsRequest {
	s.RequestSource = &v
	return s
}

func (s *SearchRecursionRecordsRequest) SetRr(v string) *SearchRecursionRecordsRequest {
	s.Rr = &v
	return s
}

func (s *SearchRecursionRecordsRequest) SetTtl(v int32) *SearchRecursionRecordsRequest {
	s.Ttl = &v
	return s
}

func (s *SearchRecursionRecordsRequest) SetType(v string) *SearchRecursionRecordsRequest {
	s.Type = &v
	return s
}

func (s *SearchRecursionRecordsRequest) SetValue(v string) *SearchRecursionRecordsRequest {
	s.Value = &v
	return s
}

func (s *SearchRecursionRecordsRequest) SetWeight(v int32) *SearchRecursionRecordsRequest {
	s.Weight = &v
	return s
}

func (s *SearchRecursionRecordsRequest) SetZoneId(v string) *SearchRecursionRecordsRequest {
	s.ZoneId = &v
	return s
}

func (s *SearchRecursionRecordsRequest) Validate() error {
	return dara.Validate(s)
}
