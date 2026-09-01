// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataAgentTaskModelUsageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBeginTime(v int64) *GetDataAgentTaskModelUsageRequest
	GetBeginTime() *int64
	SetDMSUnit(v string) *GetDataAgentTaskModelUsageRequest
	GetDMSUnit() *string
	SetEndTime(v int64) *GetDataAgentTaskModelUsageRequest
	GetEndTime() *int64
	SetInstanceIds(v []*string) *GetDataAgentTaskModelUsageRequest
	GetInstanceIds() []*string
	SetPayLevel(v string) *GetDataAgentTaskModelUsageRequest
	GetPayLevel() *string
	SetRegionId(v string) *GetDataAgentTaskModelUsageRequest
	GetRegionId() *string
}

type GetDataAgentTaskModelUsageRequest struct {
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
	EndTime     *int64    `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	PayLevel    *string   `json:"PayLevel,omitempty" xml:"PayLevel,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetDataAgentTaskModelUsageRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDataAgentTaskModelUsageRequest) GoString() string {
	return s.String()
}

func (s *GetDataAgentTaskModelUsageRequest) GetBeginTime() *int64 {
	return s.BeginTime
}

func (s *GetDataAgentTaskModelUsageRequest) GetDMSUnit() *string {
	return s.DMSUnit
}

func (s *GetDataAgentTaskModelUsageRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *GetDataAgentTaskModelUsageRequest) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *GetDataAgentTaskModelUsageRequest) GetPayLevel() *string {
	return s.PayLevel
}

func (s *GetDataAgentTaskModelUsageRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetDataAgentTaskModelUsageRequest) SetBeginTime(v int64) *GetDataAgentTaskModelUsageRequest {
	s.BeginTime = &v
	return s
}

func (s *GetDataAgentTaskModelUsageRequest) SetDMSUnit(v string) *GetDataAgentTaskModelUsageRequest {
	s.DMSUnit = &v
	return s
}

func (s *GetDataAgentTaskModelUsageRequest) SetEndTime(v int64) *GetDataAgentTaskModelUsageRequest {
	s.EndTime = &v
	return s
}

func (s *GetDataAgentTaskModelUsageRequest) SetInstanceIds(v []*string) *GetDataAgentTaskModelUsageRequest {
	s.InstanceIds = v
	return s
}

func (s *GetDataAgentTaskModelUsageRequest) SetPayLevel(v string) *GetDataAgentTaskModelUsageRequest {
	s.PayLevel = &v
	return s
}

func (s *GetDataAgentTaskModelUsageRequest) SetRegionId(v string) *GetDataAgentTaskModelUsageRequest {
	s.RegionId = &v
	return s
}

func (s *GetDataAgentTaskModelUsageRequest) Validate() error {
	return dara.Validate(s)
}
