// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataAgentTaskModelUsageMetricsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBeginTime(v string) *GetDataAgentTaskModelUsageMetricsRequest
	GetBeginTime() *string
	SetDMSUnit(v string) *GetDataAgentTaskModelUsageMetricsRequest
	GetDMSUnit() *string
	SetEndTime(v string) *GetDataAgentTaskModelUsageMetricsRequest
	GetEndTime() *string
	SetInstanceIds(v []*string) *GetDataAgentTaskModelUsageMetricsRequest
	GetInstanceIds() []*string
	SetPayLevel(v string) *GetDataAgentTaskModelUsageMetricsRequest
	GetPayLevel() *string
	SetRegionId(v string) *GetDataAgentTaskModelUsageMetricsRequest
	GetRegionId() *string
}

type GetDataAgentTaskModelUsageMetricsRequest struct {
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
	EndTime     *string   `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	PayLevel    *string   `json:"PayLevel,omitempty" xml:"PayLevel,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetDataAgentTaskModelUsageMetricsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDataAgentTaskModelUsageMetricsRequest) GoString() string {
	return s.String()
}

func (s *GetDataAgentTaskModelUsageMetricsRequest) GetBeginTime() *string {
	return s.BeginTime
}

func (s *GetDataAgentTaskModelUsageMetricsRequest) GetDMSUnit() *string {
	return s.DMSUnit
}

func (s *GetDataAgentTaskModelUsageMetricsRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *GetDataAgentTaskModelUsageMetricsRequest) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *GetDataAgentTaskModelUsageMetricsRequest) GetPayLevel() *string {
	return s.PayLevel
}

func (s *GetDataAgentTaskModelUsageMetricsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetDataAgentTaskModelUsageMetricsRequest) SetBeginTime(v string) *GetDataAgentTaskModelUsageMetricsRequest {
	s.BeginTime = &v
	return s
}

func (s *GetDataAgentTaskModelUsageMetricsRequest) SetDMSUnit(v string) *GetDataAgentTaskModelUsageMetricsRequest {
	s.DMSUnit = &v
	return s
}

func (s *GetDataAgentTaskModelUsageMetricsRequest) SetEndTime(v string) *GetDataAgentTaskModelUsageMetricsRequest {
	s.EndTime = &v
	return s
}

func (s *GetDataAgentTaskModelUsageMetricsRequest) SetInstanceIds(v []*string) *GetDataAgentTaskModelUsageMetricsRequest {
	s.InstanceIds = v
	return s
}

func (s *GetDataAgentTaskModelUsageMetricsRequest) SetPayLevel(v string) *GetDataAgentTaskModelUsageMetricsRequest {
	s.PayLevel = &v
	return s
}

func (s *GetDataAgentTaskModelUsageMetricsRequest) SetRegionId(v string) *GetDataAgentTaskModelUsageMetricsRequest {
	s.RegionId = &v
	return s
}

func (s *GetDataAgentTaskModelUsageMetricsRequest) Validate() error {
	return dara.Validate(s)
}
