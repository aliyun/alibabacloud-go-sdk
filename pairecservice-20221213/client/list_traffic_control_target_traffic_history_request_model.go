// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTrafficControlTargetTrafficHistoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *ListTrafficControlTargetTrafficHistoryRequest
	GetEndTime() *string
	SetEnvironment(v string) *ListTrafficControlTargetTrafficHistoryRequest
	GetEnvironment() *string
	SetExperimentGroupId(v string) *ListTrafficControlTargetTrafficHistoryRequest
	GetExperimentGroupId() *string
	SetExperimentId(v string) *ListTrafficControlTargetTrafficHistoryRequest
	GetExperimentId() *string
	SetInstanceId(v string) *ListTrafficControlTargetTrafficHistoryRequest
	GetInstanceId() *string
	SetItemId(v string) *ListTrafficControlTargetTrafficHistoryRequest
	GetItemId() *string
	SetStartTime(v string) *ListTrafficControlTargetTrafficHistoryRequest
	GetStartTime() *string
	SetThreshold(v string) *ListTrafficControlTargetTrafficHistoryRequest
	GetThreshold() *string
}

type ListTrafficControlTargetTrafficHistoryRequest struct {
	EndTime           *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Environment       *string `json:"Environment,omitempty" xml:"Environment,omitempty"`
	ExperimentGroupId *string `json:"ExperimentGroupId,omitempty" xml:"ExperimentGroupId,omitempty"`
	ExperimentId      *string `json:"ExperimentId,omitempty" xml:"ExperimentId,omitempty"`
	InstanceId        *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ItemId            *string `json:"ItemId,omitempty" xml:"ItemId,omitempty"`
	StartTime         *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Threshold         *string `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
}

func (s ListTrafficControlTargetTrafficHistoryRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTrafficControlTargetTrafficHistoryRequest) GoString() string {
	return s.String()
}

func (s *ListTrafficControlTargetTrafficHistoryRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *ListTrafficControlTargetTrafficHistoryRequest) GetEnvironment() *string {
	return s.Environment
}

func (s *ListTrafficControlTargetTrafficHistoryRequest) GetExperimentGroupId() *string {
	return s.ExperimentGroupId
}

func (s *ListTrafficControlTargetTrafficHistoryRequest) GetExperimentId() *string {
	return s.ExperimentId
}

func (s *ListTrafficControlTargetTrafficHistoryRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListTrafficControlTargetTrafficHistoryRequest) GetItemId() *string {
	return s.ItemId
}

func (s *ListTrafficControlTargetTrafficHistoryRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *ListTrafficControlTargetTrafficHistoryRequest) GetThreshold() *string {
	return s.Threshold
}

func (s *ListTrafficControlTargetTrafficHistoryRequest) SetEndTime(v string) *ListTrafficControlTargetTrafficHistoryRequest {
	s.EndTime = &v
	return s
}

func (s *ListTrafficControlTargetTrafficHistoryRequest) SetEnvironment(v string) *ListTrafficControlTargetTrafficHistoryRequest {
	s.Environment = &v
	return s
}

func (s *ListTrafficControlTargetTrafficHistoryRequest) SetExperimentGroupId(v string) *ListTrafficControlTargetTrafficHistoryRequest {
	s.ExperimentGroupId = &v
	return s
}

func (s *ListTrafficControlTargetTrafficHistoryRequest) SetExperimentId(v string) *ListTrafficControlTargetTrafficHistoryRequest {
	s.ExperimentId = &v
	return s
}

func (s *ListTrafficControlTargetTrafficHistoryRequest) SetInstanceId(v string) *ListTrafficControlTargetTrafficHistoryRequest {
	s.InstanceId = &v
	return s
}

func (s *ListTrafficControlTargetTrafficHistoryRequest) SetItemId(v string) *ListTrafficControlTargetTrafficHistoryRequest {
	s.ItemId = &v
	return s
}

func (s *ListTrafficControlTargetTrafficHistoryRequest) SetStartTime(v string) *ListTrafficControlTargetTrafficHistoryRequest {
	s.StartTime = &v
	return s
}

func (s *ListTrafficControlTargetTrafficHistoryRequest) SetThreshold(v string) *ListTrafficControlTargetTrafficHistoryRequest {
	s.Threshold = &v
	return s
}

func (s *ListTrafficControlTargetTrafficHistoryRequest) Validate() error {
	return dara.Validate(s)
}
