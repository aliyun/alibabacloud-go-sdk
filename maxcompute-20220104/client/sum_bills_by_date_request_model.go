// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSumBillsByDateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndDate(v int64) *SumBillsByDateRequest
	GetEndDate() *int64
	SetProjectNames(v []*string) *SumBillsByDateRequest
	GetProjectNames() []*string
	SetStartDate(v int64) *SumBillsByDateRequest
	GetStartDate() *int64
	SetStatsType(v string) *SumBillsByDateRequest
	GetStatsType() *string
	SetTopN(v int32) *SumBillsByDateRequest
	GetTopN() *int32
}

type SumBillsByDateRequest struct {
	// example:
	//
	// 1718590596556
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
	// 8
	TopN *int32 `json:"topN,omitempty" xml:"topN,omitempty"`
}

func (s SumBillsByDateRequest) String() string {
	return dara.Prettify(s)
}

func (s SumBillsByDateRequest) GoString() string {
	return s.String()
}

func (s *SumBillsByDateRequest) GetEndDate() *int64 {
	return s.EndDate
}

func (s *SumBillsByDateRequest) GetProjectNames() []*string {
	return s.ProjectNames
}

func (s *SumBillsByDateRequest) GetStartDate() *int64 {
	return s.StartDate
}

func (s *SumBillsByDateRequest) GetStatsType() *string {
	return s.StatsType
}

func (s *SumBillsByDateRequest) GetTopN() *int32 {
	return s.TopN
}

func (s *SumBillsByDateRequest) SetEndDate(v int64) *SumBillsByDateRequest {
	s.EndDate = &v
	return s
}

func (s *SumBillsByDateRequest) SetProjectNames(v []*string) *SumBillsByDateRequest {
	s.ProjectNames = v
	return s
}

func (s *SumBillsByDateRequest) SetStartDate(v int64) *SumBillsByDateRequest {
	s.StartDate = &v
	return s
}

func (s *SumBillsByDateRequest) SetStatsType(v string) *SumBillsByDateRequest {
	s.StatsType = &v
	return s
}

func (s *SumBillsByDateRequest) SetTopN(v int32) *SumBillsByDateRequest {
	s.TopN = &v
	return s
}

func (s *SumBillsByDateRequest) Validate() error {
	return dara.Validate(s)
}
