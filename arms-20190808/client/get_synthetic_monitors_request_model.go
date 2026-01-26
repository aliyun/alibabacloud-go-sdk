// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSyntheticMonitorsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFilter(v *GetSyntheticMonitorsRequestFilter) *GetSyntheticMonitorsRequest
	GetFilter() *GetSyntheticMonitorsRequestFilter
	SetRegionId(v string) *GetSyntheticMonitorsRequest
	GetRegionId() *string
}

type GetSyntheticMonitorsRequest struct {
	// The query conditions.
	//
	// This parameter is required.
	Filter *GetSyntheticMonitorsRequestFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Struct"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetSyntheticMonitorsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSyntheticMonitorsRequest) GoString() string {
	return s.String()
}

func (s *GetSyntheticMonitorsRequest) GetFilter() *GetSyntheticMonitorsRequestFilter {
	return s.Filter
}

func (s *GetSyntheticMonitorsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetSyntheticMonitorsRequest) SetFilter(v *GetSyntheticMonitorsRequestFilter) *GetSyntheticMonitorsRequest {
	s.Filter = v
	return s
}

func (s *GetSyntheticMonitorsRequest) SetRegionId(v string) *GetSyntheticMonitorsRequest {
	s.RegionId = &v
	return s
}

func (s *GetSyntheticMonitorsRequest) Validate() error {
	if s.Filter != nil {
		if err := s.Filter.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetSyntheticMonitorsRequestFilter struct {
	// The type of the monitoring point. Valid values: 1: PC. 2: mobile device.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	MonitorCategory *int32 `json:"MonitorCategory,omitempty" xml:"MonitorCategory,omitempty"`
	// The network type. Valid values: 1: private network. 2: Internet.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	Network *int32 `json:"Network,omitempty" xml:"Network,omitempty"`
	// The type of the monitoring task. Valid values:
	//
	// 1: ICMP. 2: TCP. 3: DNS. 4: HTTP. 5: website speed. 6: file download.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	TaskType *int32 `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s GetSyntheticMonitorsRequestFilter) String() string {
	return dara.Prettify(s)
}

func (s GetSyntheticMonitorsRequestFilter) GoString() string {
	return s.String()
}

func (s *GetSyntheticMonitorsRequestFilter) GetMonitorCategory() *int32 {
	return s.MonitorCategory
}

func (s *GetSyntheticMonitorsRequestFilter) GetNetwork() *int32 {
	return s.Network
}

func (s *GetSyntheticMonitorsRequestFilter) GetTaskType() *int32 {
	return s.TaskType
}

func (s *GetSyntheticMonitorsRequestFilter) SetMonitorCategory(v int32) *GetSyntheticMonitorsRequestFilter {
	s.MonitorCategory = &v
	return s
}

func (s *GetSyntheticMonitorsRequestFilter) SetNetwork(v int32) *GetSyntheticMonitorsRequestFilter {
	s.Network = &v
	return s
}

func (s *GetSyntheticMonitorsRequestFilter) SetTaskType(v int32) *GetSyntheticMonitorsRequestFilter {
	s.TaskType = &v
	return s
}

func (s *GetSyntheticMonitorsRequestFilter) Validate() error {
	return dara.Validate(s)
}
