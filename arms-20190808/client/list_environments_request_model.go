// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEnvironmentsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddonName(v string) *ListEnvironmentsRequest
	GetAddonName() *string
	SetBindResourceId(v string) *ListEnvironmentsRequest
	GetBindResourceId() *string
	SetEnvironmentType(v string) *ListEnvironmentsRequest
	GetEnvironmentType() *string
	SetFeePackage(v string) *ListEnvironmentsRequest
	GetFeePackage() *string
	SetFilterRegionIds(v string) *ListEnvironmentsRequest
	GetFilterRegionIds() *string
	SetRegionId(v string) *ListEnvironmentsRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *ListEnvironmentsRequest
	GetResourceGroupId() *string
	SetTag(v []*ListEnvironmentsRequestTag) *ListEnvironmentsRequest
	GetTag() []*ListEnvironmentsRequestTag
}

type ListEnvironmentsRequest struct {
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
	Tag []*ListEnvironmentsRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListEnvironmentsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListEnvironmentsRequest) GoString() string {
	return s.String()
}

func (s *ListEnvironmentsRequest) GetAddonName() *string {
	return s.AddonName
}

func (s *ListEnvironmentsRequest) GetBindResourceId() *string {
	return s.BindResourceId
}

func (s *ListEnvironmentsRequest) GetEnvironmentType() *string {
	return s.EnvironmentType
}

func (s *ListEnvironmentsRequest) GetFeePackage() *string {
	return s.FeePackage
}

func (s *ListEnvironmentsRequest) GetFilterRegionIds() *string {
	return s.FilterRegionIds
}

func (s *ListEnvironmentsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListEnvironmentsRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListEnvironmentsRequest) GetTag() []*ListEnvironmentsRequestTag {
	return s.Tag
}

func (s *ListEnvironmentsRequest) SetAddonName(v string) *ListEnvironmentsRequest {
	s.AddonName = &v
	return s
}

func (s *ListEnvironmentsRequest) SetBindResourceId(v string) *ListEnvironmentsRequest {
	s.BindResourceId = &v
	return s
}

func (s *ListEnvironmentsRequest) SetEnvironmentType(v string) *ListEnvironmentsRequest {
	s.EnvironmentType = &v
	return s
}

func (s *ListEnvironmentsRequest) SetFeePackage(v string) *ListEnvironmentsRequest {
	s.FeePackage = &v
	return s
}

func (s *ListEnvironmentsRequest) SetFilterRegionIds(v string) *ListEnvironmentsRequest {
	s.FilterRegionIds = &v
	return s
}

func (s *ListEnvironmentsRequest) SetRegionId(v string) *ListEnvironmentsRequest {
	s.RegionId = &v
	return s
}

func (s *ListEnvironmentsRequest) SetResourceGroupId(v string) *ListEnvironmentsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListEnvironmentsRequest) SetTag(v []*ListEnvironmentsRequestTag) *ListEnvironmentsRequest {
	s.Tag = v
	return s
}

func (s *ListEnvironmentsRequest) Validate() error {
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

type ListEnvironmentsRequestTag struct {
	// The tag key.
	//
	// example:
	//
	// env
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// gfn_web_outbound_add
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListEnvironmentsRequestTag) String() string {
	return dara.Prettify(s)
}

func (s ListEnvironmentsRequestTag) GoString() string {
	return s.String()
}

func (s *ListEnvironmentsRequestTag) GetKey() *string {
	return s.Key
}

func (s *ListEnvironmentsRequestTag) GetValue() *string {
	return s.Value
}

func (s *ListEnvironmentsRequestTag) SetKey(v string) *ListEnvironmentsRequestTag {
	s.Key = &v
	return s
}

func (s *ListEnvironmentsRequestTag) SetValue(v string) *ListEnvironmentsRequestTag {
	s.Value = &v
	return s
}

func (s *ListEnvironmentsRequestTag) Validate() error {
	return dara.Validate(s)
}
