// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateClusterAutoRenewRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoRenewInstances(v []*AutoRenewInstance) *UpdateClusterAutoRenewRequest
	GetAutoRenewInstances() []*AutoRenewInstance
	SetClusterAutoRenew(v bool) *UpdateClusterAutoRenewRequest
	GetClusterAutoRenew() *bool
	SetClusterAutoRenewDuration(v int32) *UpdateClusterAutoRenewRequest
	GetClusterAutoRenewDuration() *int32
	SetClusterAutoRenewDurationUnit(v string) *UpdateClusterAutoRenewRequest
	GetClusterAutoRenewDurationUnit() *string
	SetClusterId(v string) *UpdateClusterAutoRenewRequest
	GetClusterId() *string
	SetRegionId(v string) *UpdateClusterAutoRenewRequest
	GetRegionId() *string
	SetRenewAllInstances(v bool) *UpdateClusterAutoRenewRequest
	GetRenewAllInstances() *bool
}

type UpdateClusterAutoRenewRequest struct {
	// 自动续费ECS实例列表。
	AutoRenewInstances []*AutoRenewInstance `json:"AutoRenewInstances,omitempty" xml:"AutoRenewInstances,omitempty" type:"Repeated"`
	// 集群是否自动续费。
	//
	// example:
	//
	// true
	ClusterAutoRenew *bool `json:"ClusterAutoRenew,omitempty" xml:"ClusterAutoRenew,omitempty"`
	// 集群自动续费时长。
	//
	// example:
	//
	// 1
	ClusterAutoRenewDuration *int32 `json:"ClusterAutoRenewDuration,omitempty" xml:"ClusterAutoRenewDuration,omitempty"`
	// 集群续费时长单位。
	//
	// example:
	//
	// Monthly
	ClusterAutoRenewDurationUnit *string `json:"ClusterAutoRenewDurationUnit,omitempty" xml:"ClusterAutoRenewDurationUnit,omitempty"`
	// 集群ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// c-d6661c71139a****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// 区域ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// 续费所有ECS实例。
	RenewAllInstances *bool `json:"RenewAllInstances,omitempty" xml:"RenewAllInstances,omitempty"`
}

func (s UpdateClusterAutoRenewRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateClusterAutoRenewRequest) GoString() string {
	return s.String()
}

func (s *UpdateClusterAutoRenewRequest) GetAutoRenewInstances() []*AutoRenewInstance {
	return s.AutoRenewInstances
}

func (s *UpdateClusterAutoRenewRequest) GetClusterAutoRenew() *bool {
	return s.ClusterAutoRenew
}

func (s *UpdateClusterAutoRenewRequest) GetClusterAutoRenewDuration() *int32 {
	return s.ClusterAutoRenewDuration
}

func (s *UpdateClusterAutoRenewRequest) GetClusterAutoRenewDurationUnit() *string {
	return s.ClusterAutoRenewDurationUnit
}

func (s *UpdateClusterAutoRenewRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *UpdateClusterAutoRenewRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateClusterAutoRenewRequest) GetRenewAllInstances() *bool {
	return s.RenewAllInstances
}

func (s *UpdateClusterAutoRenewRequest) SetAutoRenewInstances(v []*AutoRenewInstance) *UpdateClusterAutoRenewRequest {
	s.AutoRenewInstances = v
	return s
}

func (s *UpdateClusterAutoRenewRequest) SetClusterAutoRenew(v bool) *UpdateClusterAutoRenewRequest {
	s.ClusterAutoRenew = &v
	return s
}

func (s *UpdateClusterAutoRenewRequest) SetClusterAutoRenewDuration(v int32) *UpdateClusterAutoRenewRequest {
	s.ClusterAutoRenewDuration = &v
	return s
}

func (s *UpdateClusterAutoRenewRequest) SetClusterAutoRenewDurationUnit(v string) *UpdateClusterAutoRenewRequest {
	s.ClusterAutoRenewDurationUnit = &v
	return s
}

func (s *UpdateClusterAutoRenewRequest) SetClusterId(v string) *UpdateClusterAutoRenewRequest {
	s.ClusterId = &v
	return s
}

func (s *UpdateClusterAutoRenewRequest) SetRegionId(v string) *UpdateClusterAutoRenewRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateClusterAutoRenewRequest) SetRenewAllInstances(v bool) *UpdateClusterAutoRenewRequest {
	s.RenewAllInstances = &v
	return s
}

func (s *UpdateClusterAutoRenewRequest) Validate() error {
	if s.AutoRenewInstances != nil {
		for _, item := range s.AutoRenewInstances {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
