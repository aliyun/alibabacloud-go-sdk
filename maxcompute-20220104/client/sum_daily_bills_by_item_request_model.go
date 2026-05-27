// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSumDailyBillsByItemRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndDate(v int64) *SumDailyBillsByItemRequest
	GetEndDate() *int64
	SetPageNumber(v int64) *SumDailyBillsByItemRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *SumDailyBillsByItemRequest
	GetPageSize() *int64
	SetProjectNames(v []*string) *SumDailyBillsByItemRequest
	GetProjectNames() []*string
	SetStartDate(v int64) *SumDailyBillsByItemRequest
	GetStartDate() *int64
	SetStatsType(v string) *SumDailyBillsByItemRequest
	GetStatsType() *string
	SetTypes(v []*string) *SumDailyBillsByItemRequest
	GetTypes() []*string
}

type SumDailyBillsByItemRequest struct {
	// example:
	//
	// 1718590596556
	EndDate *int64 `json:"endDate,omitempty" xml:"endDate,omitempty"`
	// example:
	//
	// 1
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize     *int64    `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	ProjectNames []*string `json:"projectNames,omitempty" xml:"projectNames,omitempty" type:"Repeated"`
	// example:
	//
	// 1715393576201
	StartDate *int64 `json:"startDate,omitempty" xml:"startDate,omitempty"`
	// example:
	//
	// PROJECT
	StatsType *string   `json:"statsType,omitempty" xml:"statsType,omitempty"`
	Types     []*string `json:"types,omitempty" xml:"types,omitempty" type:"Repeated"`
}

func (s SumDailyBillsByItemRequest) String() string {
	return dara.Prettify(s)
}

func (s SumDailyBillsByItemRequest) GoString() string {
	return s.String()
}

func (s *SumDailyBillsByItemRequest) GetEndDate() *int64 {
	return s.EndDate
}

func (s *SumDailyBillsByItemRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *SumDailyBillsByItemRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *SumDailyBillsByItemRequest) GetProjectNames() []*string {
	return s.ProjectNames
}

func (s *SumDailyBillsByItemRequest) GetStartDate() *int64 {
	return s.StartDate
}

func (s *SumDailyBillsByItemRequest) GetStatsType() *string {
	return s.StatsType
}

func (s *SumDailyBillsByItemRequest) GetTypes() []*string {
	return s.Types
}

func (s *SumDailyBillsByItemRequest) SetEndDate(v int64) *SumDailyBillsByItemRequest {
	s.EndDate = &v
	return s
}

func (s *SumDailyBillsByItemRequest) SetPageNumber(v int64) *SumDailyBillsByItemRequest {
	s.PageNumber = &v
	return s
}

func (s *SumDailyBillsByItemRequest) SetPageSize(v int64) *SumDailyBillsByItemRequest {
	s.PageSize = &v
	return s
}

func (s *SumDailyBillsByItemRequest) SetProjectNames(v []*string) *SumDailyBillsByItemRequest {
	s.ProjectNames = v
	return s
}

func (s *SumDailyBillsByItemRequest) SetStartDate(v int64) *SumDailyBillsByItemRequest {
	s.StartDate = &v
	return s
}

func (s *SumDailyBillsByItemRequest) SetStatsType(v string) *SumDailyBillsByItemRequest {
	s.StatsType = &v
	return s
}

func (s *SumDailyBillsByItemRequest) SetTypes(v []*string) *SumDailyBillsByItemRequest {
	s.Types = v
	return s
}

func (s *SumDailyBillsByItemRequest) Validate() error {
	return dara.Validate(s)
}
