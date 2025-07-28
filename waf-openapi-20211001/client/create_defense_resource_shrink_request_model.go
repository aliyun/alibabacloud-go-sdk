// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDefenseResourceShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustomHeadersShrink(v string) *CreateDefenseResourceShrinkRequest
	GetCustomHeadersShrink() *string
	SetDescription(v string) *CreateDefenseResourceShrinkRequest
	GetDescription() *string
	SetDetail(v string) *CreateDefenseResourceShrinkRequest
	GetDetail() *string
	SetInstanceId(v string) *CreateDefenseResourceShrinkRequest
	GetInstanceId() *string
	SetOwnerUserId(v string) *CreateDefenseResourceShrinkRequest
	GetOwnerUserId() *string
	SetPattern(v string) *CreateDefenseResourceShrinkRequest
	GetPattern() *string
	SetProduct(v string) *CreateDefenseResourceShrinkRequest
	GetProduct() *string
	SetRegionId(v string) *CreateDefenseResourceShrinkRequest
	GetRegionId() *string
	SetResource(v string) *CreateDefenseResourceShrinkRequest
	GetResource() *string
	SetResourceGroup(v string) *CreateDefenseResourceShrinkRequest
	GetResourceGroup() *string
	SetResourceManagerResourceGroupId(v string) *CreateDefenseResourceShrinkRequest
	GetResourceManagerResourceGroupId() *string
	SetResourceOrigin(v string) *CreateDefenseResourceShrinkRequest
	GetResourceOrigin() *string
	SetTag(v []*CreateDefenseResourceShrinkRequestTag) *CreateDefenseResourceShrinkRequest
	GetTag() []*CreateDefenseResourceShrinkRequestTag
	SetXffStatus(v int32) *CreateDefenseResourceShrinkRequest
	GetXffStatus() *int32
}

type CreateDefenseResourceShrinkRequest struct {
	CustomHeadersShrink *string `json:"CustomHeaders,omitempty" xml:"CustomHeaders,omitempty"`
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// {\\"domain\\": \\"zhhclb4test096-05111.test.com\\"}
	Detail *string `json:"Detail,omitempty" xml:"Detail,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// waf_v3prepaid_public_cn-4xl*******
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 123221XXX
	OwnerUserId *string `json:"OwnerUserId,omitempty" xml:"OwnerUserId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// domain
	Pattern *string `json:"Pattern,omitempty" xml:"Pattern,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ecs
	Product *string `json:"Product,omitempty" xml:"Product,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// abctest.com
	Resource      *string `json:"Resource,omitempty" xml:"Resource,omitempty"`
	ResourceGroup *string `json:"ResourceGroup,omitempty" xml:"ResourceGroup,omitempty"`
	// example:
	//
	// rg-acfm***q
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// custom
	ResourceOrigin *string                                  `json:"ResourceOrigin,omitempty" xml:"ResourceOrigin,omitempty"`
	Tag            []*CreateDefenseResourceShrinkRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// example:
	//
	// 0
	XffStatus *int32 `json:"XffStatus,omitempty" xml:"XffStatus,omitempty"`
}

func (s CreateDefenseResourceShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDefenseResourceShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateDefenseResourceShrinkRequest) GetCustomHeadersShrink() *string {
	return s.CustomHeadersShrink
}

func (s *CreateDefenseResourceShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateDefenseResourceShrinkRequest) GetDetail() *string {
	return s.Detail
}

func (s *CreateDefenseResourceShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateDefenseResourceShrinkRequest) GetOwnerUserId() *string {
	return s.OwnerUserId
}

func (s *CreateDefenseResourceShrinkRequest) GetPattern() *string {
	return s.Pattern
}

func (s *CreateDefenseResourceShrinkRequest) GetProduct() *string {
	return s.Product
}

func (s *CreateDefenseResourceShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateDefenseResourceShrinkRequest) GetResource() *string {
	return s.Resource
}

func (s *CreateDefenseResourceShrinkRequest) GetResourceGroup() *string {
	return s.ResourceGroup
}

func (s *CreateDefenseResourceShrinkRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *CreateDefenseResourceShrinkRequest) GetResourceOrigin() *string {
	return s.ResourceOrigin
}

func (s *CreateDefenseResourceShrinkRequest) GetTag() []*CreateDefenseResourceShrinkRequestTag {
	return s.Tag
}

func (s *CreateDefenseResourceShrinkRequest) GetXffStatus() *int32 {
	return s.XffStatus
}

func (s *CreateDefenseResourceShrinkRequest) SetCustomHeadersShrink(v string) *CreateDefenseResourceShrinkRequest {
	s.CustomHeadersShrink = &v
	return s
}

func (s *CreateDefenseResourceShrinkRequest) SetDescription(v string) *CreateDefenseResourceShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreateDefenseResourceShrinkRequest) SetDetail(v string) *CreateDefenseResourceShrinkRequest {
	s.Detail = &v
	return s
}

func (s *CreateDefenseResourceShrinkRequest) SetInstanceId(v string) *CreateDefenseResourceShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateDefenseResourceShrinkRequest) SetOwnerUserId(v string) *CreateDefenseResourceShrinkRequest {
	s.OwnerUserId = &v
	return s
}

func (s *CreateDefenseResourceShrinkRequest) SetPattern(v string) *CreateDefenseResourceShrinkRequest {
	s.Pattern = &v
	return s
}

func (s *CreateDefenseResourceShrinkRequest) SetProduct(v string) *CreateDefenseResourceShrinkRequest {
	s.Product = &v
	return s
}

func (s *CreateDefenseResourceShrinkRequest) SetRegionId(v string) *CreateDefenseResourceShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *CreateDefenseResourceShrinkRequest) SetResource(v string) *CreateDefenseResourceShrinkRequest {
	s.Resource = &v
	return s
}

func (s *CreateDefenseResourceShrinkRequest) SetResourceGroup(v string) *CreateDefenseResourceShrinkRequest {
	s.ResourceGroup = &v
	return s
}

func (s *CreateDefenseResourceShrinkRequest) SetResourceManagerResourceGroupId(v string) *CreateDefenseResourceShrinkRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *CreateDefenseResourceShrinkRequest) SetResourceOrigin(v string) *CreateDefenseResourceShrinkRequest {
	s.ResourceOrigin = &v
	return s
}

func (s *CreateDefenseResourceShrinkRequest) SetTag(v []*CreateDefenseResourceShrinkRequestTag) *CreateDefenseResourceShrinkRequest {
	s.Tag = v
	return s
}

func (s *CreateDefenseResourceShrinkRequest) SetXffStatus(v int32) *CreateDefenseResourceShrinkRequest {
	s.XffStatus = &v
	return s
}

func (s *CreateDefenseResourceShrinkRequest) Validate() error {
	return dara.Validate(s)
}

type CreateDefenseResourceShrinkRequestTag struct {
	// example:
	//
	// demoTagKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// TagValue1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateDefenseResourceShrinkRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateDefenseResourceShrinkRequestTag) GoString() string {
	return s.String()
}

func (s *CreateDefenseResourceShrinkRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateDefenseResourceShrinkRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateDefenseResourceShrinkRequestTag) SetKey(v string) *CreateDefenseResourceShrinkRequestTag {
	s.Key = &v
	return s
}

func (s *CreateDefenseResourceShrinkRequestTag) SetValue(v string) *CreateDefenseResourceShrinkRequestTag {
	s.Value = &v
	return s
}

func (s *CreateDefenseResourceShrinkRequestTag) Validate() error {
	return dara.Validate(s)
}
