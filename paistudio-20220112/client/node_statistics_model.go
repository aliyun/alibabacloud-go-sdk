// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iNodeStatistics interface {
	dara.Model
	String() string
	GoString() string
	SetActualMinResources(v *StatisticsResources) *NodeStatistics
	GetActualMinResources() *StatisticsResources
	SetGPUDetails(v []*NodeStatisticsGPUDetails) *NodeStatistics
	GetGPUDetails() []*NodeStatisticsGPUDetails
	SetHyperNodeDetails(v []*NodeStatisticsHyperNodeDetails) *NodeStatistics
	GetHyperNodeDetails() []*NodeStatisticsHyperNodeDetails
	SetIdleResources(v *StatisticsResources) *NodeStatistics
	GetIdleResources() *StatisticsResources
	SetSchedulableResources(v *StatisticsResources) *NodeStatistics
	GetSchedulableResources() *StatisticsResources
	SetSystemReservedResources(v *StatisticsResources) *NodeStatistics
	GetSystemReservedResources() *StatisticsResources
}

type NodeStatistics struct {
	ActualMinResources      *StatisticsResources              `json:"ActualMinResources,omitempty" xml:"ActualMinResources,omitempty"`
	GPUDetails              []*NodeStatisticsGPUDetails       `json:"GPUDetails,omitempty" xml:"GPUDetails,omitempty" type:"Repeated"`
	HyperNodeDetails        []*NodeStatisticsHyperNodeDetails `json:"HyperNodeDetails,omitempty" xml:"HyperNodeDetails,omitempty" type:"Repeated"`
	IdleResources           *StatisticsResources              `json:"IdleResources,omitempty" xml:"IdleResources,omitempty"`
	SchedulableResources    *StatisticsResources              `json:"SchedulableResources,omitempty" xml:"SchedulableResources,omitempty"`
	SystemReservedResources *StatisticsResources              `json:"SystemReservedResources,omitempty" xml:"SystemReservedResources,omitempty"`
}

func (s NodeStatistics) String() string {
	return dara.Prettify(s)
}

func (s NodeStatistics) GoString() string {
	return s.String()
}

func (s *NodeStatistics) GetActualMinResources() *StatisticsResources {
	return s.ActualMinResources
}

func (s *NodeStatistics) GetGPUDetails() []*NodeStatisticsGPUDetails {
	return s.GPUDetails
}

func (s *NodeStatistics) GetHyperNodeDetails() []*NodeStatisticsHyperNodeDetails {
	return s.HyperNodeDetails
}

func (s *NodeStatistics) GetIdleResources() *StatisticsResources {
	return s.IdleResources
}

func (s *NodeStatistics) GetSchedulableResources() *StatisticsResources {
	return s.SchedulableResources
}

func (s *NodeStatistics) GetSystemReservedResources() *StatisticsResources {
	return s.SystemReservedResources
}

func (s *NodeStatistics) SetActualMinResources(v *StatisticsResources) *NodeStatistics {
	s.ActualMinResources = v
	return s
}

func (s *NodeStatistics) SetGPUDetails(v []*NodeStatisticsGPUDetails) *NodeStatistics {
	s.GPUDetails = v
	return s
}

func (s *NodeStatistics) SetHyperNodeDetails(v []*NodeStatisticsHyperNodeDetails) *NodeStatistics {
	s.HyperNodeDetails = v
	return s
}

func (s *NodeStatistics) SetIdleResources(v *StatisticsResources) *NodeStatistics {
	s.IdleResources = v
	return s
}

func (s *NodeStatistics) SetSchedulableResources(v *StatisticsResources) *NodeStatistics {
	s.SchedulableResources = v
	return s
}

func (s *NodeStatistics) SetSystemReservedResources(v *StatisticsResources) *NodeStatistics {
	s.SystemReservedResources = v
	return s
}

func (s *NodeStatistics) Validate() error {
	if s.ActualMinResources != nil {
		if err := s.ActualMinResources.Validate(); err != nil {
			return err
		}
	}
	if s.GPUDetails != nil {
		for _, item := range s.GPUDetails {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.HyperNodeDetails != nil {
		for _, item := range s.HyperNodeDetails {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.IdleResources != nil {
		if err := s.IdleResources.Validate(); err != nil {
			return err
		}
	}
	if s.SchedulableResources != nil {
		if err := s.SchedulableResources.Validate(); err != nil {
			return err
		}
	}
	if s.SystemReservedResources != nil {
		if err := s.SystemReservedResources.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type NodeStatisticsGPUDetails struct {
	Count   *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	IdleNum *int64 `json:"IdleNum,omitempty" xml:"IdleNum,omitempty"`
}

func (s NodeStatisticsGPUDetails) String() string {
	return dara.Prettify(s)
}

func (s NodeStatisticsGPUDetails) GoString() string {
	return s.String()
}

func (s *NodeStatisticsGPUDetails) GetCount() *int64 {
	return s.Count
}

func (s *NodeStatisticsGPUDetails) GetIdleNum() *int64 {
	return s.IdleNum
}

func (s *NodeStatisticsGPUDetails) SetCount(v int64) *NodeStatisticsGPUDetails {
	s.Count = &v
	return s
}

func (s *NodeStatisticsGPUDetails) SetIdleNum(v int64) *NodeStatisticsGPUDetails {
	s.IdleNum = &v
	return s
}

func (s *NodeStatisticsGPUDetails) Validate() error {
	return dara.Validate(s)
}

type NodeStatisticsHyperNodeDetails struct {
	Count   *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	IdleNum *int64 `json:"IdleNum,omitempty" xml:"IdleNum,omitempty"`
}

func (s NodeStatisticsHyperNodeDetails) String() string {
	return dara.Prettify(s)
}

func (s NodeStatisticsHyperNodeDetails) GoString() string {
	return s.String()
}

func (s *NodeStatisticsHyperNodeDetails) GetCount() *int64 {
	return s.Count
}

func (s *NodeStatisticsHyperNodeDetails) GetIdleNum() *int64 {
	return s.IdleNum
}

func (s *NodeStatisticsHyperNodeDetails) SetCount(v int64) *NodeStatisticsHyperNodeDetails {
	s.Count = &v
	return s
}

func (s *NodeStatisticsHyperNodeDetails) SetIdleNum(v int64) *NodeStatisticsHyperNodeDetails {
	s.IdleNum = &v
	return s
}

func (s *NodeStatisticsHyperNodeDetails) Validate() error {
	return dara.Validate(s)
}
