// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetStandardStatisticsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *GetStandardStatisticsRequest
	GetOpTenantId() *int64
	SetStatisticsQuery(v *GetStandardStatisticsRequestStatisticsQuery) *GetStandardStatisticsRequest
	GetStatisticsQuery() *GetStandardStatisticsRequestStatisticsQuery
}

type GetStandardStatisticsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// This parameter is required.
	StatisticsQuery *GetStandardStatisticsRequestStatisticsQuery `json:"StatisticsQuery,omitempty" xml:"StatisticsQuery,omitempty" type:"Struct"`
}

func (s GetStandardStatisticsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetStandardStatisticsRequest) GoString() string {
	return s.String()
}

func (s *GetStandardStatisticsRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *GetStandardStatisticsRequest) GetStatisticsQuery() *GetStandardStatisticsRequestStatisticsQuery {
	return s.StatisticsQuery
}

func (s *GetStandardStatisticsRequest) SetOpTenantId(v int64) *GetStandardStatisticsRequest {
	s.OpTenantId = &v
	return s
}

func (s *GetStandardStatisticsRequest) SetStatisticsQuery(v *GetStandardStatisticsRequestStatisticsQuery) *GetStandardStatisticsRequest {
	s.StatisticsQuery = v
	return s
}

func (s *GetStandardStatisticsRequest) Validate() error {
	if s.StatisticsQuery != nil {
		if err := s.StatisticsQuery.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetStandardStatisticsRequestStatisticsQuery struct {
	CreateTimePeriod *GetStandardStatisticsRequestStatisticsQueryCreateTimePeriod `json:"CreateTimePeriod,omitempty" xml:"CreateTimePeriod,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// /dir1
	Directory         *string   `json:"Directory,omitempty" xml:"Directory,omitempty"`
	StandardStageList []*string `json:"StandardStageList,omitempty" xml:"StandardStageList,omitempty" type:"Repeated"`
}

func (s GetStandardStatisticsRequestStatisticsQuery) String() string {
	return dara.Prettify(s)
}

func (s GetStandardStatisticsRequestStatisticsQuery) GoString() string {
	return s.String()
}

func (s *GetStandardStatisticsRequestStatisticsQuery) GetCreateTimePeriod() *GetStandardStatisticsRequestStatisticsQueryCreateTimePeriod {
	return s.CreateTimePeriod
}

func (s *GetStandardStatisticsRequestStatisticsQuery) GetDirectory() *string {
	return s.Directory
}

func (s *GetStandardStatisticsRequestStatisticsQuery) GetStandardStageList() []*string {
	return s.StandardStageList
}

func (s *GetStandardStatisticsRequestStatisticsQuery) SetCreateTimePeriod(v *GetStandardStatisticsRequestStatisticsQueryCreateTimePeriod) *GetStandardStatisticsRequestStatisticsQuery {
	s.CreateTimePeriod = v
	return s
}

func (s *GetStandardStatisticsRequestStatisticsQuery) SetDirectory(v string) *GetStandardStatisticsRequestStatisticsQuery {
	s.Directory = &v
	return s
}

func (s *GetStandardStatisticsRequestStatisticsQuery) SetStandardStageList(v []*string) *GetStandardStatisticsRequestStatisticsQuery {
	s.StandardStageList = v
	return s
}

func (s *GetStandardStatisticsRequestStatisticsQuery) Validate() error {
	if s.CreateTimePeriod != nil {
		if err := s.CreateTimePeriod.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetStandardStatisticsRequestStatisticsQueryCreateTimePeriod struct {
	// example:
	//
	// 2025-06-30 00:00:00
	EndTime        *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	IncludeEndTime *bool   `json:"IncludeEndTime,omitempty" xml:"IncludeEndTime,omitempty"`
	// example:
	//
	// 2025-06-01 00:00:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s GetStandardStatisticsRequestStatisticsQueryCreateTimePeriod) String() string {
	return dara.Prettify(s)
}

func (s GetStandardStatisticsRequestStatisticsQueryCreateTimePeriod) GoString() string {
	return s.String()
}

func (s *GetStandardStatisticsRequestStatisticsQueryCreateTimePeriod) GetEndTime() *string {
	return s.EndTime
}

func (s *GetStandardStatisticsRequestStatisticsQueryCreateTimePeriod) GetIncludeEndTime() *bool {
	return s.IncludeEndTime
}

func (s *GetStandardStatisticsRequestStatisticsQueryCreateTimePeriod) GetStartTime() *string {
	return s.StartTime
}

func (s *GetStandardStatisticsRequestStatisticsQueryCreateTimePeriod) SetEndTime(v string) *GetStandardStatisticsRequestStatisticsQueryCreateTimePeriod {
	s.EndTime = &v
	return s
}

func (s *GetStandardStatisticsRequestStatisticsQueryCreateTimePeriod) SetIncludeEndTime(v bool) *GetStandardStatisticsRequestStatisticsQueryCreateTimePeriod {
	s.IncludeEndTime = &v
	return s
}

func (s *GetStandardStatisticsRequestStatisticsQueryCreateTimePeriod) SetStartTime(v string) *GetStandardStatisticsRequestStatisticsQueryCreateTimePeriod {
	s.StartTime = &v
	return s
}

func (s *GetStandardStatisticsRequestStatisticsQueryCreateTimePeriod) Validate() error {
	return dara.Validate(s)
}
