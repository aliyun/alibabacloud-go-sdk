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
	// A list of custom header fields.
	//
	// > If you set XffStatus to 1, WAF uses the first IP address from the specified header field as the client IP address to prevent XFF forgery. If you specify multiple header fields, WAF tries to obtain the client IP address from the header fields in sequence. If WAF fails to obtain the client IP address from the specified header fields, it uses the first IP address in the X-Forwarded-For header field.
	CustomHeadersShrink *string `json:"CustomHeaders,omitempty" xml:"CustomHeaders,omitempty"`
	// The description of the protected object.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The detailed parameters of the protected object. This parameter is a string that consists of a JSON struct.
	//
	// > The parameters vary based on the values of **Product*	- and **Pattern**. For more information, see the "**Description of the Detail parameter**" section.
	//
	// This parameter is required.
	//
	// example:
	//
	// {\\"domain\\": \\"zhhclb4test096-05111.test.com\\"}
	Detail *string `json:"Detail,omitempty" xml:"Detail,omitempty"`
	// The ID of the WAF instance.
	//
	// > Call [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) to query the ID of the WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_v3prepaid_public_cn-4xl*******
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the account to which the protected object belongs. This parameter is used in multi-account scenarios. By default, the protected object belongs to the WAF administrator account.
	//
	// example:
	//
	// 123221XXX
	OwnerUserId *string `json:"OwnerUserId,omitempty" xml:"OwnerUserId,omitempty"`
	// The protection mode of the protected object. Valid values:
	//
	// - **domain**: domain name.
	//
	// - **multi_service**: hybrid cloud deployment.
	//
	// This parameter is required.
	//
	// example:
	//
	// domain
	Pattern *string `json:"Pattern,omitempty" xml:"Pattern,omitempty"`
	// The name of the Alibaba Cloud service. Valid values:
	//
	// - **alb**: Application Load Balancer (ALB).
	//
	// - **ecs**: Elastic Compute Service (ECS).
	//
	// - **clb4**: Layer 4 Classic Load Balancer (CLB).
	//
	// - **clb7**: Layer 7 CLB.
	//
	// - **nlb**: Network Load Balancer (NLB).
	//
	// - **waf**: Web Application Firewall (WAF).
	//
	// This parameter is required.
	//
	// example:
	//
	// ecs
	Product *string `json:"Product,omitempty" xml:"Product,omitempty"`
	// The region where the WAF instance is deployed. Valid values:
	//
	// - **cn-hangzhou**: the Chinese mainland.
	//
	// - **ap-southeast-1**: outside the Chinese mainland.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the protected object.
	//
	// > - Only protected objects of hybrid cloud deployments support custom names.
	//
	// example:
	//
	// abctest.com
	Resource *string `json:"Resource,omitempty" xml:"Resource,omitempty"`
	// The name of the protection group to which you want to add the protected object. This parameter is optional.
	//
	// example:
	//
	// testGroup
	ResourceGroup *string `json:"ResourceGroup,omitempty" xml:"ResourceGroup,omitempty"`
	// The ID of the Alibaba Cloud resource group.
	//
	// example:
	//
	// rg-acfm***q
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
	// The source of the protected object. Only the following value is supported:
	//
	// - **custom**: a custom object.
	//
	// This parameter is required.
	//
	// example:
	//
	// custom
	ResourceOrigin *string `json:"ResourceOrigin,omitempty" xml:"ResourceOrigin,omitempty"`
	// A list of tags. You can add up to 20 tags.
	Tag []*CreateDefenseResourceShrinkRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// Specifies whether to enable the X-Forwarded-For (XFF) proxy. Valid values:
	//
	// - **0**: disabled. This is the default value.
	//
	// - **1**: enabled.
	//
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

type CreateDefenseResourceShrinkRequestTag struct {
	// The tag key.
	//
	// example:
	//
	// demoTagKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
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
