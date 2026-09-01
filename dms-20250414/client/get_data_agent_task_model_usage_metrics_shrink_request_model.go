// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataAgentTaskModelUsageMetricsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBeginTime(v string) *GetDataAgentTaskModelUsageMetricsShrinkRequest
	GetBeginTime() *string
	SetDMSUnit(v string) *GetDataAgentTaskModelUsageMetricsShrinkRequest
	GetDMSUnit() *string
	SetEndTime(v string) *GetDataAgentTaskModelUsageMetricsShrinkRequest
	GetEndTime() *string
	SetInstanceIdsShrink(v string) *GetDataAgentTaskModelUsageMetricsShrinkRequest
	GetInstanceIdsShrink() *string
	SetPayLevel(v string) *GetDataAgentTaskModelUsageMetricsShrinkRequest
	GetPayLevel() *string
	SetRegionId(v string) *GetDataAgentTaskModelUsageMetricsShrinkRequest
	GetRegionId() *string
}

type GetDataAgentTaskModelUsageMetricsShrinkRequest struct {
	// The start time of the query time range. The value is a UNIX timestamp in seconds. The recommended interval length is no longer than one month.
	//
	// example:
	//
	// 1735660800
	BeginTime *string `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	// The current DMS unit.
	//
	// example:
	//
	// cn-hangzhou
	DMSUnit *string `json:"DMSUnit,omitempty" xml:"DMSUnit,omitempty"`
	// The end time of the query time range. The value is a UNIX timestamp in seconds. The recommended interval length is no longer than one month.
	//
	// example:
	//
	// 1735747200
	EndTime           *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	InstanceIdsShrink *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
	PayLevel          *string `json:"PayLevel,omitempty" xml:"PayLevel,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetDataAgentTaskModelUsageMetricsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDataAgentTaskModelUsageMetricsShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetDataAgentTaskModelUsageMetricsShrinkRequest) GetBeginTime() *string {
	return s.BeginTime
}

func (s *GetDataAgentTaskModelUsageMetricsShrinkRequest) GetDMSUnit() *string {
	return s.DMSUnit
}

func (s *GetDataAgentTaskModelUsageMetricsShrinkRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *GetDataAgentTaskModelUsageMetricsShrinkRequest) GetInstanceIdsShrink() *string {
	return s.InstanceIdsShrink
}

func (s *GetDataAgentTaskModelUsageMetricsShrinkRequest) GetPayLevel() *string {
	return s.PayLevel
}

func (s *GetDataAgentTaskModelUsageMetricsShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetDataAgentTaskModelUsageMetricsShrinkRequest) SetBeginTime(v string) *GetDataAgentTaskModelUsageMetricsShrinkRequest {
	s.BeginTime = &v
	return s
}

func (s *GetDataAgentTaskModelUsageMetricsShrinkRequest) SetDMSUnit(v string) *GetDataAgentTaskModelUsageMetricsShrinkRequest {
	s.DMSUnit = &v
	return s
}

func (s *GetDataAgentTaskModelUsageMetricsShrinkRequest) SetEndTime(v string) *GetDataAgentTaskModelUsageMetricsShrinkRequest {
	s.EndTime = &v
	return s
}

func (s *GetDataAgentTaskModelUsageMetricsShrinkRequest) SetInstanceIdsShrink(v string) *GetDataAgentTaskModelUsageMetricsShrinkRequest {
	s.InstanceIdsShrink = &v
	return s
}

func (s *GetDataAgentTaskModelUsageMetricsShrinkRequest) SetPayLevel(v string) *GetDataAgentTaskModelUsageMetricsShrinkRequest {
	s.PayLevel = &v
	return s
}

func (s *GetDataAgentTaskModelUsageMetricsShrinkRequest) SetRegionId(v string) *GetDataAgentTaskModelUsageMetricsShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *GetDataAgentTaskModelUsageMetricsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
