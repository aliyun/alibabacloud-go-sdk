// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSumBillsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndDate(v int64) *SumBillsRequest
	GetEndDate() *int64
	SetProjectNames(v []*string) *SumBillsRequest
	GetProjectNames() []*string
	SetStartDate(v int64) *SumBillsRequest
	GetStartDate() *int64
	SetStatsType(v string) *SumBillsRequest
	GetStatsType() *string
	SetTopN(v int32) *SumBillsRequest
	GetTopN() *int32
}

type SumBillsRequest struct {
	// example:
	//
	// 1776232895313
	EndDate      *int64    `json:"endDate,omitempty" xml:"endDate,omitempty"`
	ProjectNames []*string `json:"projectNames,omitempty" xml:"projectNames,omitempty" type:"Repeated"`
	// example:
	//
	// 1715393576201
	StartDate *int64 `json:"startDate,omitempty" xml:"startDate,omitempty"`
	// example:
	//
	// PROJECT
	StatsType *string `json:"statsType,omitempty" xml:"statsType,omitempty"`
	// example:
	//
	// 5
	TopN *int32 `json:"topN,omitempty" xml:"topN,omitempty"`
}

func (s SumBillsRequest) String() string {
	return dara.Prettify(s)
}

func (s SumBillsRequest) GoString() string {
	return s.String()
}

func (s *SumBillsRequest) GetEndDate() *int64 {
	return s.EndDate
}

func (s *SumBillsRequest) GetProjectNames() []*string {
	return s.ProjectNames
}

func (s *SumBillsRequest) GetStartDate() *int64 {
	return s.StartDate
}

func (s *SumBillsRequest) GetStatsType() *string {
	return s.StatsType
}

func (s *SumBillsRequest) GetTopN() *int32 {
	return s.TopN
}

func (s *SumBillsRequest) SetEndDate(v int64) *SumBillsRequest {
	s.EndDate = &v
	return s
}

func (s *SumBillsRequest) SetProjectNames(v []*string) *SumBillsRequest {
	s.ProjectNames = v
	return s
}

func (s *SumBillsRequest) SetStartDate(v int64) *SumBillsRequest {
	s.StartDate = &v
	return s
}

func (s *SumBillsRequest) SetStatsType(v string) *SumBillsRequest {
	s.StatsType = &v
	return s
}

func (s *SumBillsRequest) SetTopN(v int32) *SumBillsRequest {
	s.TopN = &v
	return s
}

func (s *SumBillsRequest) Validate() error {
	return dara.Validate(s)
}
