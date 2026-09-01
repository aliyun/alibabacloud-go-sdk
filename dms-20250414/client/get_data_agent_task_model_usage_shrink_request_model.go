// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataAgentTaskModelUsageShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBeginTime(v int64) *GetDataAgentTaskModelUsageShrinkRequest
	GetBeginTime() *int64
	SetDMSUnit(v string) *GetDataAgentTaskModelUsageShrinkRequest
	GetDMSUnit() *string
	SetEndTime(v int64) *GetDataAgentTaskModelUsageShrinkRequest
	GetEndTime() *int64
	SetInstanceIdsShrink(v string) *GetDataAgentTaskModelUsageShrinkRequest
	GetInstanceIdsShrink() *string
	SetPayLevel(v string) *GetDataAgentTaskModelUsageShrinkRequest
	GetPayLevel() *string
	SetRegionId(v string) *GetDataAgentTaskModelUsageShrinkRequest
	GetRegionId() *string
}

type GetDataAgentTaskModelUsageShrinkRequest struct {
	// The start time of the query time range. The value is a UNIX timestamp in seconds. The recommended interval length is no longer than one month.
	//
	// example:
	//
	// 1735660800
	BeginTime *int64 `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	// The current Data Management unit.
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
	EndTime           *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	InstanceIdsShrink *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
	PayLevel          *string `json:"PayLevel,omitempty" xml:"PayLevel,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetDataAgentTaskModelUsageShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDataAgentTaskModelUsageShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetDataAgentTaskModelUsageShrinkRequest) GetBeginTime() *int64 {
	return s.BeginTime
}

func (s *GetDataAgentTaskModelUsageShrinkRequest) GetDMSUnit() *string {
	return s.DMSUnit
}

func (s *GetDataAgentTaskModelUsageShrinkRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *GetDataAgentTaskModelUsageShrinkRequest) GetInstanceIdsShrink() *string {
	return s.InstanceIdsShrink
}

func (s *GetDataAgentTaskModelUsageShrinkRequest) GetPayLevel() *string {
	return s.PayLevel
}

func (s *GetDataAgentTaskModelUsageShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetDataAgentTaskModelUsageShrinkRequest) SetBeginTime(v int64) *GetDataAgentTaskModelUsageShrinkRequest {
	s.BeginTime = &v
	return s
}

func (s *GetDataAgentTaskModelUsageShrinkRequest) SetDMSUnit(v string) *GetDataAgentTaskModelUsageShrinkRequest {
	s.DMSUnit = &v
	return s
}

func (s *GetDataAgentTaskModelUsageShrinkRequest) SetEndTime(v int64) *GetDataAgentTaskModelUsageShrinkRequest {
	s.EndTime = &v
	return s
}

func (s *GetDataAgentTaskModelUsageShrinkRequest) SetInstanceIdsShrink(v string) *GetDataAgentTaskModelUsageShrinkRequest {
	s.InstanceIdsShrink = &v
	return s
}

func (s *GetDataAgentTaskModelUsageShrinkRequest) SetPayLevel(v string) *GetDataAgentTaskModelUsageShrinkRequest {
	s.PayLevel = &v
	return s
}

func (s *GetDataAgentTaskModelUsageShrinkRequest) SetRegionId(v string) *GetDataAgentTaskModelUsageShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *GetDataAgentTaskModelUsageShrinkRequest) Validate() error {
	return dara.Validate(s)
}
