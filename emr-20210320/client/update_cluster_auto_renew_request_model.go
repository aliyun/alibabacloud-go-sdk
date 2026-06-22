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
	// The list of ECS instances for which to enable auto-renewal. This parameter takes effect only when RenewAllInstances is not set to true.
	AutoRenewInstances []*AutoRenewInstance `json:"AutoRenewInstances,omitempty" xml:"AutoRenewInstances,omitempty" type:"Repeated"`
	// Specifies whether to enable auto-renewal for the cluster. Valid values:
	//
	// - true: Enables auto-renewal.
	//
	// - false: Disables auto-renewal.
	//
	// Default value: false.
	//
	// example:
	//
	// true
	ClusterAutoRenew *bool `json:"ClusterAutoRenew,omitempty" xml:"ClusterAutoRenew,omitempty"`
	// The auto-renewal duration for the cluster. This parameter takes effect only when ClusterAutoRenew is set to true.
	//
	// If ClusterAutoRenewDurationUnit is set to Month, the valid values are 1, 2, 3, 4, 5, 6, 7, 8, 9, 12, 24, and 36. If ClusterAutoRenewDurationUnit is set to Year, the valid values are 1, 2, and 3.
	//
	// example:
	//
	// 1
	ClusterAutoRenewDuration *int32 `json:"ClusterAutoRenewDuration,omitempty" xml:"ClusterAutoRenewDuration,omitempty"`
	// The unit of the auto-renewal duration. Valid values:
	//
	// - Month
	//
	// - Year
	//
	// Default value: Month.
	//
	// example:
	//
	// Monthly
	ClusterAutoRenewDurationUnit *string `json:"ClusterAutoRenewDurationUnit,omitempty" xml:"ClusterAutoRenewDurationUnit,omitempty"`
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// c-d6661c71139a****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Specifies whether to enable auto-renewal for all ECS instances in the cluster. Valid values:
	//
	// - true: Enables auto-renewal for all ECS instances in the cluster. The default auto-renewal duration is one month.
	//
	// - false: Does not enable auto-renewal for all ECS instances in the cluster. You can specify the ECS instances for which to enable auto-renewal in the AutoRenewInstances parameter.
	//
	// Default value: false.
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
