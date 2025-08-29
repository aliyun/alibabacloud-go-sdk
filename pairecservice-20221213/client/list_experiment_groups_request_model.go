// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListExperimentGroupsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListExperimentGroupsRequest
	GetInstanceId() *string
	SetLayerId(v string) *ListExperimentGroupsRequest
	GetLayerId() *string
	SetRegionId(v string) *ListExperimentGroupsRequest
	GetRegionId() *string
	SetStatus(v string) *ListExperimentGroupsRequest
	GetStatus() *string
	SetTimeRangeEnd(v string) *ListExperimentGroupsRequest
	GetTimeRangeEnd() *string
	SetTimeRangeStart(v string) *ListExperimentGroupsRequest
	GetTimeRangeStart() *string
}

type ListExperimentGroupsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pairec-cn-abcdefg1234
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 3
	LayerId  *string `json:"LayerId,omitempty" xml:"LayerId,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// Online
	Status         *string `json:"Status,omitempty" xml:"Status,omitempty"`
	TimeRangeEnd   *string `json:"TimeRangeEnd,omitempty" xml:"TimeRangeEnd,omitempty"`
	TimeRangeStart *string `json:"TimeRangeStart,omitempty" xml:"TimeRangeStart,omitempty"`
}

func (s ListExperimentGroupsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListExperimentGroupsRequest) GoString() string {
	return s.String()
}

func (s *ListExperimentGroupsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListExperimentGroupsRequest) GetLayerId() *string {
	return s.LayerId
}

func (s *ListExperimentGroupsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListExperimentGroupsRequest) GetStatus() *string {
	return s.Status
}

func (s *ListExperimentGroupsRequest) GetTimeRangeEnd() *string {
	return s.TimeRangeEnd
}

func (s *ListExperimentGroupsRequest) GetTimeRangeStart() *string {
	return s.TimeRangeStart
}

func (s *ListExperimentGroupsRequest) SetInstanceId(v string) *ListExperimentGroupsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListExperimentGroupsRequest) SetLayerId(v string) *ListExperimentGroupsRequest {
	s.LayerId = &v
	return s
}

func (s *ListExperimentGroupsRequest) SetRegionId(v string) *ListExperimentGroupsRequest {
	s.RegionId = &v
	return s
}

func (s *ListExperimentGroupsRequest) SetStatus(v string) *ListExperimentGroupsRequest {
	s.Status = &v
	return s
}

func (s *ListExperimentGroupsRequest) SetTimeRangeEnd(v string) *ListExperimentGroupsRequest {
	s.TimeRangeEnd = &v
	return s
}

func (s *ListExperimentGroupsRequest) SetTimeRangeStart(v string) *ListExperimentGroupsRequest {
	s.TimeRangeStart = &v
	return s
}

func (s *ListExperimentGroupsRequest) Validate() error {
	return dara.Validate(s)
}
