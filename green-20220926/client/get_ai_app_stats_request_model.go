// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAiAppStatsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetByMonth(v bool) *GetAiAppStatsRequest
	GetByMonth() *bool
	SetEndDate(v string) *GetAiAppStatsRequest
	GetEndDate() *string
	SetQuery(v string) *GetAiAppStatsRequest
	GetQuery() *string
	SetRegionId(v string) *GetAiAppStatsRequest
	GetRegionId() *string
	SetStartDate(v string) *GetAiAppStatsRequest
	GetStartDate() *string
	SetType(v string) *GetAiAppStatsRequest
	GetType() *string
}

type GetAiAppStatsRequest struct {
	// Specifies whether to aggregate by month. Default value: false.
	//
	// example:
	//
	// false
	ByMonth *bool `json:"ByMonth,omitempty" xml:"ByMonth,omitempty"`
	// The query end date.
	//
	// example:
	//
	// 2026-01-02 00:00:00
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// The query condition.
	//
	// example:
	//
	// {}
	Query *string `json:"Query,omitempty" xml:"Query,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The query start date.
	//
	// example:
	//
	// 2026-01-01 00:00:00
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	// The type.
	//
	// example:
	//
	// sensitive_data
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetAiAppStatsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAiAppStatsRequest) GoString() string {
	return s.String()
}

func (s *GetAiAppStatsRequest) GetByMonth() *bool {
	return s.ByMonth
}

func (s *GetAiAppStatsRequest) GetEndDate() *string {
	return s.EndDate
}

func (s *GetAiAppStatsRequest) GetQuery() *string {
	return s.Query
}

func (s *GetAiAppStatsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetAiAppStatsRequest) GetStartDate() *string {
	return s.StartDate
}

func (s *GetAiAppStatsRequest) GetType() *string {
	return s.Type
}

func (s *GetAiAppStatsRequest) SetByMonth(v bool) *GetAiAppStatsRequest {
	s.ByMonth = &v
	return s
}

func (s *GetAiAppStatsRequest) SetEndDate(v string) *GetAiAppStatsRequest {
	s.EndDate = &v
	return s
}

func (s *GetAiAppStatsRequest) SetQuery(v string) *GetAiAppStatsRequest {
	s.Query = &v
	return s
}

func (s *GetAiAppStatsRequest) SetRegionId(v string) *GetAiAppStatsRequest {
	s.RegionId = &v
	return s
}

func (s *GetAiAppStatsRequest) SetStartDate(v string) *GetAiAppStatsRequest {
	s.StartDate = &v
	return s
}

func (s *GetAiAppStatsRequest) SetType(v string) *GetAiAppStatsRequest {
	s.Type = &v
	return s
}

func (s *GetAiAppStatsRequest) Validate() error {
	return dara.Validate(s)
}
