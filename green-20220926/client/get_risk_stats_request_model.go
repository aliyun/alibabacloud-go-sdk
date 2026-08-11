// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRiskStatsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClassify(v string) *GetRiskStatsRequest
	GetClassify() *string
	SetEndTime(v string) *GetRiskStatsRequest
	GetEndTime() *string
	SetRegionId(v string) *GetRiskStatsRequest
	GetRegionId() *string
	SetStartTime(v string) *GetRiskStatsRequest
	GetStartTime() *string
	SetType(v string) *GetRiskStatsRequest
	GetType() *string
}

type GetRiskStatsRequest struct {
	// The classification.
	//
	// example:
	//
	// guard-scene
	Classify *string `json:"Classify,omitempty" xml:"Classify,omitempty"`
	// The end time.
	//
	// example:
	//
	// 2026-01-02 00:00:00
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The start time.
	//
	// example:
	//
	// 2026-01-01 00:00:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The type.
	//
	// This parameter is required.
	//
	// example:
	//
	// RealTime
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetRiskStatsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetRiskStatsRequest) GoString() string {
	return s.String()
}

func (s *GetRiskStatsRequest) GetClassify() *string {
	return s.Classify
}

func (s *GetRiskStatsRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *GetRiskStatsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetRiskStatsRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *GetRiskStatsRequest) GetType() *string {
	return s.Type
}

func (s *GetRiskStatsRequest) SetClassify(v string) *GetRiskStatsRequest {
	s.Classify = &v
	return s
}

func (s *GetRiskStatsRequest) SetEndTime(v string) *GetRiskStatsRequest {
	s.EndTime = &v
	return s
}

func (s *GetRiskStatsRequest) SetRegionId(v string) *GetRiskStatsRequest {
	s.RegionId = &v
	return s
}

func (s *GetRiskStatsRequest) SetStartTime(v string) *GetRiskStatsRequest {
	s.StartTime = &v
	return s
}

func (s *GetRiskStatsRequest) SetType(v string) *GetRiskStatsRequest {
	s.Type = &v
	return s
}

func (s *GetRiskStatsRequest) Validate() error {
	return dara.Validate(s)
}
