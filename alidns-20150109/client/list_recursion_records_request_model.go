// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRecursionRecordsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnable(v string) *ListRecursionRecordsRequest
	GetEnable() *string
	SetMaxResults(v int32) *ListRecursionRecordsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListRecursionRecordsRequest
	GetNextToken() *string
	SetPageNumber(v int32) *ListRecursionRecordsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListRecursionRecordsRequest
	GetPageSize() *int32
	SetRemark(v string) *ListRecursionRecordsRequest
	GetRemark() *string
	SetRequestSource(v string) *ListRecursionRecordsRequest
	GetRequestSource() *string
	SetRr(v string) *ListRecursionRecordsRequest
	GetRr() *string
	SetTtl(v int32) *ListRecursionRecordsRequest
	GetTtl() *int32
	SetType(v string) *ListRecursionRecordsRequest
	GetType() *string
	SetWeight(v int32) *ListRecursionRecordsRequest
	GetWeight() *int32
	SetZoneId(v string) *ListRecursionRecordsRequest
	GetZoneId() *string
}

type ListRecursionRecordsRequest struct {
	// example:
	//
	// enable
	Enable *string `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// example:
	//
	// 100
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// 4698691
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
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
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// remark
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
	// 1
	Weight *int32 `json:"Weight,omitempty" xml:"Weight,omitempty"`
	// example:
	//
	// 17832322323
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s ListRecursionRecordsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListRecursionRecordsRequest) GoString() string {
	return s.String()
}

func (s *ListRecursionRecordsRequest) GetEnable() *string {
	return s.Enable
}

func (s *ListRecursionRecordsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListRecursionRecordsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListRecursionRecordsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListRecursionRecordsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListRecursionRecordsRequest) GetRemark() *string {
	return s.Remark
}

func (s *ListRecursionRecordsRequest) GetRequestSource() *string {
	return s.RequestSource
}

func (s *ListRecursionRecordsRequest) GetRr() *string {
	return s.Rr
}

func (s *ListRecursionRecordsRequest) GetTtl() *int32 {
	return s.Ttl
}

func (s *ListRecursionRecordsRequest) GetType() *string {
	return s.Type
}

func (s *ListRecursionRecordsRequest) GetWeight() *int32 {
	return s.Weight
}

func (s *ListRecursionRecordsRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *ListRecursionRecordsRequest) SetEnable(v string) *ListRecursionRecordsRequest {
	s.Enable = &v
	return s
}

func (s *ListRecursionRecordsRequest) SetMaxResults(v int32) *ListRecursionRecordsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListRecursionRecordsRequest) SetNextToken(v string) *ListRecursionRecordsRequest {
	s.NextToken = &v
	return s
}

func (s *ListRecursionRecordsRequest) SetPageNumber(v int32) *ListRecursionRecordsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListRecursionRecordsRequest) SetPageSize(v int32) *ListRecursionRecordsRequest {
	s.PageSize = &v
	return s
}

func (s *ListRecursionRecordsRequest) SetRemark(v string) *ListRecursionRecordsRequest {
	s.Remark = &v
	return s
}

func (s *ListRecursionRecordsRequest) SetRequestSource(v string) *ListRecursionRecordsRequest {
	s.RequestSource = &v
	return s
}

func (s *ListRecursionRecordsRequest) SetRr(v string) *ListRecursionRecordsRequest {
	s.Rr = &v
	return s
}

func (s *ListRecursionRecordsRequest) SetTtl(v int32) *ListRecursionRecordsRequest {
	s.Ttl = &v
	return s
}

func (s *ListRecursionRecordsRequest) SetType(v string) *ListRecursionRecordsRequest {
	s.Type = &v
	return s
}

func (s *ListRecursionRecordsRequest) SetWeight(v int32) *ListRecursionRecordsRequest {
	s.Weight = &v
	return s
}

func (s *ListRecursionRecordsRequest) SetZoneId(v string) *ListRecursionRecordsRequest {
	s.ZoneId = &v
	return s
}

func (s *ListRecursionRecordsRequest) Validate() error {
	return dara.Validate(s)
}
