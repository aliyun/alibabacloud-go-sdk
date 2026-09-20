// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResizeMultiZoneClusterNodeCountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetArbiterVSwitchId(v string) *ResizeMultiZoneClusterNodeCountRequest
	GetArbiterVSwitchId() *string
	SetClusterId(v string) *ResizeMultiZoneClusterNodeCountRequest
	GetClusterId() *string
	SetCoreNodeCount(v int32) *ResizeMultiZoneClusterNodeCountRequest
	GetCoreNodeCount() *int32
	SetLogNodeCount(v int32) *ResizeMultiZoneClusterNodeCountRequest
	GetLogNodeCount() *int32
	SetPrimaryCoreNodeCount(v int32) *ResizeMultiZoneClusterNodeCountRequest
	GetPrimaryCoreNodeCount() *int32
	SetPrimaryVSwitchId(v string) *ResizeMultiZoneClusterNodeCountRequest
	GetPrimaryVSwitchId() *string
	SetStandbyCoreNodeCount(v int32) *ResizeMultiZoneClusterNodeCountRequest
	GetStandbyCoreNodeCount() *int32
	SetStandbyVSwitchId(v string) *ResizeMultiZoneClusterNodeCountRequest
	GetStandbyVSwitchId() *string
}

type ResizeMultiZoneClusterNodeCountRequest struct {
	// The vSwitch ID of the arbitration node.
	//
	// example:
	//
	// vsw-hangxzhouxb*****
	ArbiterVSwitchId *string `json:"ArbiterVSwitchId,omitempty" xml:"ArbiterVSwitchId,omitempty"`
	// The ID of the multi-zone cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// ld-f5d8d6s4s2a1****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The number of core nodes. The minimum value is 4, and the increment must be a multiple of 2.
	//
	// example:
	//
	// 6
	CoreNodeCount *int32 `json:"CoreNodeCount,omitempty" xml:"CoreNodeCount,omitempty"`
	// The number of log nodes. The minimum value is 4, and the value must be a multiple of 4.
	//
	// example:
	//
	// 8
	LogNodeCount *int32 `json:"LogNodeCount,omitempty" xml:"LogNodeCount,omitempty"`
	// The number of core nodes in the primary zone instance. The minimum value is 4, and the increment must be a multiple of 2.
	//
	// example:
	//
	// 6
	PrimaryCoreNodeCount *int32 `json:"PrimaryCoreNodeCount,omitempty" xml:"PrimaryCoreNodeCount,omitempty"`
	// The vSwitch ID of the instance in the primary zone.
	//
	// example:
	//
	// vsw-hangxzhouxe*****
	PrimaryVSwitchId *string `json:"PrimaryVSwitchId,omitempty" xml:"PrimaryVSwitchId,omitempty"`
	// The number of core nodes in the secondary zone instance. The minimum value is 4, and the increment must be a multiple of 2.
	//
	// example:
	//
	// 6
	StandbyCoreNodeCount *int32 `json:"StandbyCoreNodeCount,omitempty" xml:"StandbyCoreNodeCount,omitempty"`
	// The vSwitch ID of the instance in the secondary zone.
	//
	// example:
	//
	// vsw-hangxzhouxf****
	StandbyVSwitchId *string `json:"StandbyVSwitchId,omitempty" xml:"StandbyVSwitchId,omitempty"`
}

func (s ResizeMultiZoneClusterNodeCountRequest) String() string {
	return dara.Prettify(s)
}

func (s ResizeMultiZoneClusterNodeCountRequest) GoString() string {
	return s.String()
}

func (s *ResizeMultiZoneClusterNodeCountRequest) GetArbiterVSwitchId() *string {
	return s.ArbiterVSwitchId
}

func (s *ResizeMultiZoneClusterNodeCountRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *ResizeMultiZoneClusterNodeCountRequest) GetCoreNodeCount() *int32 {
	return s.CoreNodeCount
}

func (s *ResizeMultiZoneClusterNodeCountRequest) GetLogNodeCount() *int32 {
	return s.LogNodeCount
}

func (s *ResizeMultiZoneClusterNodeCountRequest) GetPrimaryCoreNodeCount() *int32 {
	return s.PrimaryCoreNodeCount
}

func (s *ResizeMultiZoneClusterNodeCountRequest) GetPrimaryVSwitchId() *string {
	return s.PrimaryVSwitchId
}

func (s *ResizeMultiZoneClusterNodeCountRequest) GetStandbyCoreNodeCount() *int32 {
	return s.StandbyCoreNodeCount
}

func (s *ResizeMultiZoneClusterNodeCountRequest) GetStandbyVSwitchId() *string {
	return s.StandbyVSwitchId
}

func (s *ResizeMultiZoneClusterNodeCountRequest) SetArbiterVSwitchId(v string) *ResizeMultiZoneClusterNodeCountRequest {
	s.ArbiterVSwitchId = &v
	return s
}

func (s *ResizeMultiZoneClusterNodeCountRequest) SetClusterId(v string) *ResizeMultiZoneClusterNodeCountRequest {
	s.ClusterId = &v
	return s
}

func (s *ResizeMultiZoneClusterNodeCountRequest) SetCoreNodeCount(v int32) *ResizeMultiZoneClusterNodeCountRequest {
	s.CoreNodeCount = &v
	return s
}

func (s *ResizeMultiZoneClusterNodeCountRequest) SetLogNodeCount(v int32) *ResizeMultiZoneClusterNodeCountRequest {
	s.LogNodeCount = &v
	return s
}

func (s *ResizeMultiZoneClusterNodeCountRequest) SetPrimaryCoreNodeCount(v int32) *ResizeMultiZoneClusterNodeCountRequest {
	s.PrimaryCoreNodeCount = &v
	return s
}

func (s *ResizeMultiZoneClusterNodeCountRequest) SetPrimaryVSwitchId(v string) *ResizeMultiZoneClusterNodeCountRequest {
	s.PrimaryVSwitchId = &v
	return s
}

func (s *ResizeMultiZoneClusterNodeCountRequest) SetStandbyCoreNodeCount(v int32) *ResizeMultiZoneClusterNodeCountRequest {
	s.StandbyCoreNodeCount = &v
	return s
}

func (s *ResizeMultiZoneClusterNodeCountRequest) SetStandbyVSwitchId(v string) *ResizeMultiZoneClusterNodeCountRequest {
	s.StandbyVSwitchId = &v
	return s
}

func (s *ResizeMultiZoneClusterNodeCountRequest) Validate() error {
	return dara.Validate(s)
}
