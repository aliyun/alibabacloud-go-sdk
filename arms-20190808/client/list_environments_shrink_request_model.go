// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEnvironmentsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddonName(v string) *ListEnvironmentsShrinkRequest
	GetAddonName() *string
	SetBindResourceId(v string) *ListEnvironmentsShrinkRequest
	GetBindResourceId() *string
	SetEnvironmentType(v string) *ListEnvironmentsShrinkRequest
	GetEnvironmentType() *string
	SetFeePackage(v string) *ListEnvironmentsShrinkRequest
	GetFeePackage() *string
	SetFilterRegionIds(v string) *ListEnvironmentsShrinkRequest
	GetFilterRegionIds() *string
	SetRegionId(v string) *ListEnvironmentsShrinkRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *ListEnvironmentsShrinkRequest
	GetResourceGroupId() *string
	SetTagShrink(v string) *ListEnvironmentsShrinkRequest
	GetTagShrink() *string
}

type ListEnvironmentsShrinkRequest struct {
	// The add-on name. You must specify at least one of the AddonName and EnvironmentType parameters.
	//
	// example:
	//
	// trace-java
	AddonName *string `json:"AddonName,omitempty" xml:"AddonName,omitempty"`
	// The ID of the resource.
	//
	// example:
	//
	// cff30f0d67d7542dfb05bd114b4b1d7af
	BindResourceId *string `json:"BindResourceId,omitempty" xml:"BindResourceId,omitempty"`
	// The environment type. You must specify at least one of the AddonName and EnvironmentType parameters.
	//
	// Valid values:
	//
	// 	- CS
	//
	//     <!-- -->
	//
	//     :
	//
	//     <!-- -->
	//
	//     Container Service for Kubernetes (ACK)
	//
	//     <!-- -->
	//
	// 	- ECS
	//
	//     <!-- -->
	//
	//     :
	//
	//     <!-- -->
	//
	//     Elastic Compute Service (ECS)
	//
	//     <!-- -->
	//
	// 	- Cloud
	//
	//     <!-- -->
	//
	//     :
	//
	//     <!-- -->
	//
	//     cloud service
	//
	//     <!-- -->
	//
	// example:
	//
	// CS
	EnvironmentType *string `json:"EnvironmentType,omitempty" xml:"EnvironmentType,omitempty"`
	// The payable resource plan.
	//
	// 	- If the EnvironmentType parameter is set to CS, set the value to CS_Basic or CS_Pro. Default value: CS_Basic.
	//
	// 	- Otherwise, leave the parameter empty.
	//
	// Valid values:
	//
	// 	- CS_Pro: Container Monitoring Pro
	//
	// 	- CS_Basic: Container Monitoring Basic
	//
	// example:
	//
	// CS_Pro
	FeePackage *string `json:"FeePackage,omitempty" xml:"FeePackage,omitempty"`
	// The region IDs to be queried.
	//
	// example:
	//
	// cn-beijing,cn-hangzhou
	FilterRegionIds *string `json:"FilterRegionIds,omitempty" xml:"FilterRegionIds,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-aekzfurdatohtka
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The tags.
	TagShrink *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
}

func (s ListEnvironmentsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListEnvironmentsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListEnvironmentsShrinkRequest) GetAddonName() *string {
	return s.AddonName
}

func (s *ListEnvironmentsShrinkRequest) GetBindResourceId() *string {
	return s.BindResourceId
}

func (s *ListEnvironmentsShrinkRequest) GetEnvironmentType() *string {
	return s.EnvironmentType
}

func (s *ListEnvironmentsShrinkRequest) GetFeePackage() *string {
	return s.FeePackage
}

func (s *ListEnvironmentsShrinkRequest) GetFilterRegionIds() *string {
	return s.FilterRegionIds
}

func (s *ListEnvironmentsShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListEnvironmentsShrinkRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListEnvironmentsShrinkRequest) GetTagShrink() *string {
	return s.TagShrink
}

func (s *ListEnvironmentsShrinkRequest) SetAddonName(v string) *ListEnvironmentsShrinkRequest {
	s.AddonName = &v
	return s
}

func (s *ListEnvironmentsShrinkRequest) SetBindResourceId(v string) *ListEnvironmentsShrinkRequest {
	s.BindResourceId = &v
	return s
}

func (s *ListEnvironmentsShrinkRequest) SetEnvironmentType(v string) *ListEnvironmentsShrinkRequest {
	s.EnvironmentType = &v
	return s
}

func (s *ListEnvironmentsShrinkRequest) SetFeePackage(v string) *ListEnvironmentsShrinkRequest {
	s.FeePackage = &v
	return s
}

func (s *ListEnvironmentsShrinkRequest) SetFilterRegionIds(v string) *ListEnvironmentsShrinkRequest {
	s.FilterRegionIds = &v
	return s
}

func (s *ListEnvironmentsShrinkRequest) SetRegionId(v string) *ListEnvironmentsShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *ListEnvironmentsShrinkRequest) SetResourceGroupId(v string) *ListEnvironmentsShrinkRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListEnvironmentsShrinkRequest) SetTagShrink(v string) *ListEnvironmentsShrinkRequest {
	s.TagShrink = &v
	return s
}

func (s *ListEnvironmentsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
