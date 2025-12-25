// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClusterStage interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentStage(v int32) *ClusterStage
	GetCurrentStage() *int32
	SetMessage(v string) *ClusterStage
	GetMessage() *string
	SetStatus(v string) *ClusterStage
	GetStatus() *string
	SetTotalStageWithWeight(v []*StageWithWeight) *ClusterStage
	GetTotalStageWithWeight() []*StageWithWeight
}

type ClusterStage struct {
	CurrentStage         *int32             `json:"CurrentStage,omitempty" xml:"CurrentStage,omitempty"`
	Message              *string            `json:"Message,omitempty" xml:"Message,omitempty"`
	Status               *string            `json:"Status,omitempty" xml:"Status,omitempty"`
	TotalStageWithWeight []*StageWithWeight `json:"TotalStageWithWeight,omitempty" xml:"TotalStageWithWeight,omitempty" type:"Repeated"`
}

func (s ClusterStage) String() string {
	return dara.Prettify(s)
}

func (s ClusterStage) GoString() string {
	return s.String()
}

func (s *ClusterStage) GetCurrentStage() *int32 {
	return s.CurrentStage
}

func (s *ClusterStage) GetMessage() *string {
	return s.Message
}

func (s *ClusterStage) GetStatus() *string {
	return s.Status
}

func (s *ClusterStage) GetTotalStageWithWeight() []*StageWithWeight {
	return s.TotalStageWithWeight
}

func (s *ClusterStage) SetCurrentStage(v int32) *ClusterStage {
	s.CurrentStage = &v
	return s
}

func (s *ClusterStage) SetMessage(v string) *ClusterStage {
	s.Message = &v
	return s
}

func (s *ClusterStage) SetStatus(v string) *ClusterStage {
	s.Status = &v
	return s
}

func (s *ClusterStage) SetTotalStageWithWeight(v []*StageWithWeight) *ClusterStage {
	s.TotalStageWithWeight = v
	return s
}

func (s *ClusterStage) Validate() error {
	if s.TotalStageWithWeight != nil {
		for _, item := range s.TotalStageWithWeight {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
