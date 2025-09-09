// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRecursionZonesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListRecursionZonesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListRecursionZonesRequest
	GetNextToken() *string
	SetPageNumber(v int32) *ListRecursionZonesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListRecursionZonesRequest
	GetPageSize() *int32
	SetRemark(v string) *ListRecursionZonesRequest
	GetRemark() *string
	SetZoneName(v string) *ListRecursionZonesRequest
	GetZoneName() *string
}

type ListRecursionZonesRequest struct {
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
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// test
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// example:
	//
	// lisheng999.com
	ZoneName *string `json:"ZoneName,omitempty" xml:"ZoneName,omitempty"`
}

func (s ListRecursionZonesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListRecursionZonesRequest) GoString() string {
	return s.String()
}

func (s *ListRecursionZonesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListRecursionZonesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListRecursionZonesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListRecursionZonesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListRecursionZonesRequest) GetRemark() *string {
	return s.Remark
}

func (s *ListRecursionZonesRequest) GetZoneName() *string {
	return s.ZoneName
}

func (s *ListRecursionZonesRequest) SetMaxResults(v int32) *ListRecursionZonesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListRecursionZonesRequest) SetNextToken(v string) *ListRecursionZonesRequest {
	s.NextToken = &v
	return s
}

func (s *ListRecursionZonesRequest) SetPageNumber(v int32) *ListRecursionZonesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListRecursionZonesRequest) SetPageSize(v int32) *ListRecursionZonesRequest {
	s.PageSize = &v
	return s
}

func (s *ListRecursionZonesRequest) SetRemark(v string) *ListRecursionZonesRequest {
	s.Remark = &v
	return s
}

func (s *ListRecursionZonesRequest) SetZoneName(v string) *ListRecursionZonesRequest {
	s.ZoneName = &v
	return s
}

func (s *ListRecursionZonesRequest) Validate() error {
	return dara.Validate(s)
}
