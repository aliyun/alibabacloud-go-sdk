// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDefenseResourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustomHeaders(v []*string) *CreateDefenseResourceRequest
	GetCustomHeaders() []*string
	SetDescription(v string) *CreateDefenseResourceRequest
	GetDescription() *string
	SetDetail(v string) *CreateDefenseResourceRequest
	GetDetail() *string
	SetInstanceId(v string) *CreateDefenseResourceRequest
	GetInstanceId() *string
	SetOwnerUserId(v string) *CreateDefenseResourceRequest
	GetOwnerUserId() *string
	SetPattern(v string) *CreateDefenseResourceRequest
	GetPattern() *string
	SetProduct(v string) *CreateDefenseResourceRequest
	GetProduct() *string
	SetRegionId(v string) *CreateDefenseResourceRequest
	GetRegionId() *string
	SetResource(v string) *CreateDefenseResourceRequest
	GetResource() *string
	SetResourceGroup(v string) *CreateDefenseResourceRequest
	GetResourceGroup() *string
	SetResourceManagerResourceGroupId(v string) *CreateDefenseResourceRequest
	GetResourceManagerResourceGroupId() *string
	SetResourceOrigin(v string) *CreateDefenseResourceRequest
	GetResourceOrigin() *string
	SetTag(v []*CreateDefenseResourceRequestTag) *CreateDefenseResourceRequest
	GetTag() []*CreateDefenseResourceRequestTag
	SetXffStatus(v int32) *CreateDefenseResourceRequest
	GetXffStatus() *int32
}

type CreateDefenseResourceRequest struct {
	CustomHeaders []*string `json:"CustomHeaders,omitempty" xml:"CustomHeaders,omitempty" type:"Repeated"`
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
	ResourceOrigin *string                            `json:"ResourceOrigin,omitempty" xml:"ResourceOrigin,omitempty"`
	Tag            []*CreateDefenseResourceRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// example:
	//
	// 0
	XffStatus *int32 `json:"XffStatus,omitempty" xml:"XffStatus,omitempty"`
}

func (s CreateDefenseResourceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDefenseResourceRequest) GoString() string {
	return s.String()
}

func (s *CreateDefenseResourceRequest) GetCustomHeaders() []*string {
	return s.CustomHeaders
}

func (s *CreateDefenseResourceRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateDefenseResourceRequest) GetDetail() *string {
	return s.Detail
}

func (s *CreateDefenseResourceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateDefenseResourceRequest) GetOwnerUserId() *string {
	return s.OwnerUserId
}

func (s *CreateDefenseResourceRequest) GetPattern() *string {
	return s.Pattern
}

func (s *CreateDefenseResourceRequest) GetProduct() *string {
	return s.Product
}

func (s *CreateDefenseResourceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateDefenseResourceRequest) GetResource() *string {
	return s.Resource
}

func (s *CreateDefenseResourceRequest) GetResourceGroup() *string {
	return s.ResourceGroup
}

func (s *CreateDefenseResourceRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *CreateDefenseResourceRequest) GetResourceOrigin() *string {
	return s.ResourceOrigin
}

func (s *CreateDefenseResourceRequest) GetTag() []*CreateDefenseResourceRequestTag {
	return s.Tag
}

func (s *CreateDefenseResourceRequest) GetXffStatus() *int32 {
	return s.XffStatus
}

func (s *CreateDefenseResourceRequest) SetCustomHeaders(v []*string) *CreateDefenseResourceRequest {
	s.CustomHeaders = v
	return s
}

func (s *CreateDefenseResourceRequest) SetDescription(v string) *CreateDefenseResourceRequest {
	s.Description = &v
	return s
}

func (s *CreateDefenseResourceRequest) SetDetail(v string) *CreateDefenseResourceRequest {
	s.Detail = &v
	return s
}

func (s *CreateDefenseResourceRequest) SetInstanceId(v string) *CreateDefenseResourceRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateDefenseResourceRequest) SetOwnerUserId(v string) *CreateDefenseResourceRequest {
	s.OwnerUserId = &v
	return s
}

func (s *CreateDefenseResourceRequest) SetPattern(v string) *CreateDefenseResourceRequest {
	s.Pattern = &v
	return s
}

func (s *CreateDefenseResourceRequest) SetProduct(v string) *CreateDefenseResourceRequest {
	s.Product = &v
	return s
}

func (s *CreateDefenseResourceRequest) SetRegionId(v string) *CreateDefenseResourceRequest {
	s.RegionId = &v
	return s
}

func (s *CreateDefenseResourceRequest) SetResource(v string) *CreateDefenseResourceRequest {
	s.Resource = &v
	return s
}

func (s *CreateDefenseResourceRequest) SetResourceGroup(v string) *CreateDefenseResourceRequest {
	s.ResourceGroup = &v
	return s
}

func (s *CreateDefenseResourceRequest) SetResourceManagerResourceGroupId(v string) *CreateDefenseResourceRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *CreateDefenseResourceRequest) SetResourceOrigin(v string) *CreateDefenseResourceRequest {
	s.ResourceOrigin = &v
	return s
}

func (s *CreateDefenseResourceRequest) SetTag(v []*CreateDefenseResourceRequestTag) *CreateDefenseResourceRequest {
	s.Tag = v
	return s
}

func (s *CreateDefenseResourceRequest) SetXffStatus(v int32) *CreateDefenseResourceRequest {
	s.XffStatus = &v
	return s
}

func (s *CreateDefenseResourceRequest) Validate() error {
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

type CreateDefenseResourceRequestTag struct {
	// example:
	//
	// demoTagKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// TagValue1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateDefenseResourceRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateDefenseResourceRequestTag) GoString() string {
	return s.String()
}

func (s *CreateDefenseResourceRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateDefenseResourceRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateDefenseResourceRequestTag) SetKey(v string) *CreateDefenseResourceRequestTag {
	s.Key = &v
	return s
}

func (s *CreateDefenseResourceRequestTag) SetValue(v string) *CreateDefenseResourceRequestTag {
	s.Value = &v
	return s
}

func (s *CreateDefenseResourceRequestTag) Validate() error {
	return dara.Validate(s)
}
