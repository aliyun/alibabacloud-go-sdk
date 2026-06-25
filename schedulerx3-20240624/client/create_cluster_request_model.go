// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateClusterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChargeType(v string) *CreateClusterRequest
	GetChargeType() *string
	SetClusterName(v string) *CreateClusterRequest
	GetClusterName() *string
	SetClusterSpec(v string) *CreateClusterRequest
	GetClusterSpec() *string
	SetClusterType(v int32) *CreateClusterRequest
	GetClusterType() *int32
	SetDuration(v int32) *CreateClusterRequest
	GetDuration() *int32
	SetEngineType(v string) *CreateClusterRequest
	GetEngineType() *string
	SetPricingCycle(v string) *CreateClusterRequest
	GetPricingCycle() *string
	SetSource(v string) *CreateClusterRequest
	GetSource() *string
	SetTag(v []*CreateClusterRequestTag) *CreateClusterRequest
	GetTag() []*CreateClusterRequestTag
	SetVSwitches(v []*CreateClusterRequestVSwitches) *CreateClusterRequest
	GetVSwitches() []*CreateClusterRequestVSwitches
	SetVpcId(v string) *CreateClusterRequest
	GetVpcId() *string
}

type CreateClusterRequest struct {
	// The billing type.
	//
	// example:
	//
	// POSTPAY
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The name of the cluster.
	//
	// This parameter is required.
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// The cluster specification. Valid values:
	//
	// - scx.dev.x1
	//
	// - scx.small.x1
	//
	// - scx.small.x2
	//
	// - scx.medium.x1
	//
	// - scx.medium.x2.
	ClusterSpec *string `json:"ClusterSpec,omitempty" xml:"ClusterSpec,omitempty"`
	// The cluster type.
	//
	// example:
	//
	// 1
	ClusterType *int32 `json:"ClusterType,omitempty" xml:"ClusterType,omitempty"`
	// The duration.
	//
	// example:
	//
	// 3
	Duration *int32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The engine type. Valid values: xxljob.
	EngineType *string `json:"EngineType,omitempty" xml:"EngineType,omitempty"`
	// The pricing cycle.
	//
	// example:
	//
	// Year
	PricingCycle *string `json:"PricingCycle,omitempty" xml:"PricingCycle,omitempty"`
	// The source.
	//
	// example:
	//
	// schedulerx
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The list of tags. A maximum of 20 tags are supported.
	Tag []*CreateClusterRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The information about the vSwitches.
	VSwitches []*CreateClusterRequestVSwitches `json:"VSwitches,omitempty" xml:"VSwitches,omitempty" type:"Repeated"`
	// The ID of the VPC.
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s CreateClusterRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateClusterRequest) GoString() string {
	return s.String()
}

func (s *CreateClusterRequest) GetChargeType() *string {
	return s.ChargeType
}

func (s *CreateClusterRequest) GetClusterName() *string {
	return s.ClusterName
}

func (s *CreateClusterRequest) GetClusterSpec() *string {
	return s.ClusterSpec
}

func (s *CreateClusterRequest) GetClusterType() *int32 {
	return s.ClusterType
}

func (s *CreateClusterRequest) GetDuration() *int32 {
	return s.Duration
}

func (s *CreateClusterRequest) GetEngineType() *string {
	return s.EngineType
}

func (s *CreateClusterRequest) GetPricingCycle() *string {
	return s.PricingCycle
}

func (s *CreateClusterRequest) GetSource() *string {
	return s.Source
}

func (s *CreateClusterRequest) GetTag() []*CreateClusterRequestTag {
	return s.Tag
}

func (s *CreateClusterRequest) GetVSwitches() []*CreateClusterRequestVSwitches {
	return s.VSwitches
}

func (s *CreateClusterRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *CreateClusterRequest) SetChargeType(v string) *CreateClusterRequest {
	s.ChargeType = &v
	return s
}

func (s *CreateClusterRequest) SetClusterName(v string) *CreateClusterRequest {
	s.ClusterName = &v
	return s
}

func (s *CreateClusterRequest) SetClusterSpec(v string) *CreateClusterRequest {
	s.ClusterSpec = &v
	return s
}

func (s *CreateClusterRequest) SetClusterType(v int32) *CreateClusterRequest {
	s.ClusterType = &v
	return s
}

func (s *CreateClusterRequest) SetDuration(v int32) *CreateClusterRequest {
	s.Duration = &v
	return s
}

func (s *CreateClusterRequest) SetEngineType(v string) *CreateClusterRequest {
	s.EngineType = &v
	return s
}

func (s *CreateClusterRequest) SetPricingCycle(v string) *CreateClusterRequest {
	s.PricingCycle = &v
	return s
}

func (s *CreateClusterRequest) SetSource(v string) *CreateClusterRequest {
	s.Source = &v
	return s
}

func (s *CreateClusterRequest) SetTag(v []*CreateClusterRequestTag) *CreateClusterRequest {
	s.Tag = v
	return s
}

func (s *CreateClusterRequest) SetVSwitches(v []*CreateClusterRequestVSwitches) *CreateClusterRequest {
	s.VSwitches = v
	return s
}

func (s *CreateClusterRequest) SetVpcId(v string) *CreateClusterRequest {
	s.VpcId = &v
	return s
}

func (s *CreateClusterRequest) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.VSwitches != nil {
		for _, item := range s.VSwitches {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateClusterRequestTag struct {
	// The tag key.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateClusterRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateClusterRequestTag) GoString() string {
	return s.String()
}

func (s *CreateClusterRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateClusterRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateClusterRequestTag) SetKey(v string) *CreateClusterRequestTag {
	s.Key = &v
	return s
}

func (s *CreateClusterRequestTag) SetValue(v string) *CreateClusterRequestTag {
	s.Value = &v
	return s
}

func (s *CreateClusterRequestTag) Validate() error {
	return dara.Validate(s)
}

type CreateClusterRequestVSwitches struct {
	// The ID of the vSwitch.
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The zone ID of the vSwitch.
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s CreateClusterRequestVSwitches) String() string {
	return dara.Prettify(s)
}

func (s CreateClusterRequestVSwitches) GoString() string {
	return s.String()
}

func (s *CreateClusterRequestVSwitches) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *CreateClusterRequestVSwitches) GetZoneId() *string {
	return s.ZoneId
}

func (s *CreateClusterRequestVSwitches) SetVSwitchId(v string) *CreateClusterRequestVSwitches {
	s.VSwitchId = &v
	return s
}

func (s *CreateClusterRequestVSwitches) SetZoneId(v string) *CreateClusterRequestVSwitches {
	s.ZoneId = &v
	return s
}

func (s *CreateClusterRequestVSwitches) Validate() error {
	return dara.Validate(s)
}
