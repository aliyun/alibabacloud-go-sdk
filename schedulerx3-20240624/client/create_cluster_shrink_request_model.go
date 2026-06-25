// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateClusterShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChargeType(v string) *CreateClusterShrinkRequest
	GetChargeType() *string
	SetClusterName(v string) *CreateClusterShrinkRequest
	GetClusterName() *string
	SetClusterSpec(v string) *CreateClusterShrinkRequest
	GetClusterSpec() *string
	SetClusterType(v int32) *CreateClusterShrinkRequest
	GetClusterType() *int32
	SetDuration(v int32) *CreateClusterShrinkRequest
	GetDuration() *int32
	SetEngineType(v string) *CreateClusterShrinkRequest
	GetEngineType() *string
	SetPricingCycle(v string) *CreateClusterShrinkRequest
	GetPricingCycle() *string
	SetSource(v string) *CreateClusterShrinkRequest
	GetSource() *string
	SetTag(v []*CreateClusterShrinkRequestTag) *CreateClusterShrinkRequest
	GetTag() []*CreateClusterShrinkRequestTag
	SetVSwitchesShrink(v string) *CreateClusterShrinkRequest
	GetVSwitchesShrink() *string
	SetVpcId(v string) *CreateClusterShrinkRequest
	GetVpcId() *string
}

type CreateClusterShrinkRequest struct {
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
	Tag []*CreateClusterShrinkRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The information about the vSwitches.
	VSwitchesShrink *string `json:"VSwitches,omitempty" xml:"VSwitches,omitempty"`
	// The ID of the VPC.
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s CreateClusterShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateClusterShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateClusterShrinkRequest) GetChargeType() *string {
	return s.ChargeType
}

func (s *CreateClusterShrinkRequest) GetClusterName() *string {
	return s.ClusterName
}

func (s *CreateClusterShrinkRequest) GetClusterSpec() *string {
	return s.ClusterSpec
}

func (s *CreateClusterShrinkRequest) GetClusterType() *int32 {
	return s.ClusterType
}

func (s *CreateClusterShrinkRequest) GetDuration() *int32 {
	return s.Duration
}

func (s *CreateClusterShrinkRequest) GetEngineType() *string {
	return s.EngineType
}

func (s *CreateClusterShrinkRequest) GetPricingCycle() *string {
	return s.PricingCycle
}

func (s *CreateClusterShrinkRequest) GetSource() *string {
	return s.Source
}

func (s *CreateClusterShrinkRequest) GetTag() []*CreateClusterShrinkRequestTag {
	return s.Tag
}

func (s *CreateClusterShrinkRequest) GetVSwitchesShrink() *string {
	return s.VSwitchesShrink
}

func (s *CreateClusterShrinkRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *CreateClusterShrinkRequest) SetChargeType(v string) *CreateClusterShrinkRequest {
	s.ChargeType = &v
	return s
}

func (s *CreateClusterShrinkRequest) SetClusterName(v string) *CreateClusterShrinkRequest {
	s.ClusterName = &v
	return s
}

func (s *CreateClusterShrinkRequest) SetClusterSpec(v string) *CreateClusterShrinkRequest {
	s.ClusterSpec = &v
	return s
}

func (s *CreateClusterShrinkRequest) SetClusterType(v int32) *CreateClusterShrinkRequest {
	s.ClusterType = &v
	return s
}

func (s *CreateClusterShrinkRequest) SetDuration(v int32) *CreateClusterShrinkRequest {
	s.Duration = &v
	return s
}

func (s *CreateClusterShrinkRequest) SetEngineType(v string) *CreateClusterShrinkRequest {
	s.EngineType = &v
	return s
}

func (s *CreateClusterShrinkRequest) SetPricingCycle(v string) *CreateClusterShrinkRequest {
	s.PricingCycle = &v
	return s
}

func (s *CreateClusterShrinkRequest) SetSource(v string) *CreateClusterShrinkRequest {
	s.Source = &v
	return s
}

func (s *CreateClusterShrinkRequest) SetTag(v []*CreateClusterShrinkRequestTag) *CreateClusterShrinkRequest {
	s.Tag = v
	return s
}

func (s *CreateClusterShrinkRequest) SetVSwitchesShrink(v string) *CreateClusterShrinkRequest {
	s.VSwitchesShrink = &v
	return s
}

func (s *CreateClusterShrinkRequest) SetVpcId(v string) *CreateClusterShrinkRequest {
	s.VpcId = &v
	return s
}

func (s *CreateClusterShrinkRequest) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateClusterShrinkRequestTag struct {
	// The tag key.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateClusterShrinkRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateClusterShrinkRequestTag) GoString() string {
	return s.String()
}

func (s *CreateClusterShrinkRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateClusterShrinkRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateClusterShrinkRequestTag) SetKey(v string) *CreateClusterShrinkRequestTag {
	s.Key = &v
	return s
}

func (s *CreateClusterShrinkRequestTag) SetValue(v string) *CreateClusterShrinkRequestTag {
	s.Value = &v
	return s
}

func (s *CreateClusterShrinkRequestTag) Validate() error {
	return dara.Validate(s)
}
