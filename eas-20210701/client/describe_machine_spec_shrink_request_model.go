// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMachineSpecShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChargeType(v string) *DescribeMachineSpecShrinkRequest
	GetChargeType() *string
	SetInstanceTypesShrink(v string) *DescribeMachineSpecShrinkRequest
	GetInstanceTypesShrink() *string
	SetResourceType(v string) *DescribeMachineSpecShrinkRequest
	GetResourceType() *string
}

type DescribeMachineSpecShrinkRequest struct {
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// Deprecated
	//
	// This parameter is deprecated.
	InstanceTypesShrink *string `json:"InstanceTypes,omitempty" xml:"InstanceTypes,omitempty"`
	ResourceType        *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s DescribeMachineSpecShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeMachineSpecShrinkRequest) GoString() string {
	return s.String()
}

func (s *DescribeMachineSpecShrinkRequest) GetChargeType() *string {
	return s.ChargeType
}

func (s *DescribeMachineSpecShrinkRequest) GetInstanceTypesShrink() *string {
	return s.InstanceTypesShrink
}

func (s *DescribeMachineSpecShrinkRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *DescribeMachineSpecShrinkRequest) SetChargeType(v string) *DescribeMachineSpecShrinkRequest {
	s.ChargeType = &v
	return s
}

func (s *DescribeMachineSpecShrinkRequest) SetInstanceTypesShrink(v string) *DescribeMachineSpecShrinkRequest {
	s.InstanceTypesShrink = &v
	return s
}

func (s *DescribeMachineSpecShrinkRequest) SetResourceType(v string) *DescribeMachineSpecShrinkRequest {
	s.ResourceType = &v
	return s
}

func (s *DescribeMachineSpecShrinkRequest) Validate() error {
	return dara.Validate(s)
}
