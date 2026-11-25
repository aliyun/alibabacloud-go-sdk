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
	// The list of specified header fields.
	//
	// > When XffStatus is set to 1, the first IP address in the specified header field is used as the originating IP address of the client to prevent XFF spoofing. If multiple headers are specified, the system attempts to obtain the originating IP address in sequence. If the first header does not contain an IP address, the system checks the second header, and so on. If no IP address is found in any of the specified headers, the first IP address in the X-Forwarded-For header is used. When XffStatus is set to 1, it indicates that the first IP address is obtained.
	CustomHeadersShrink *string `json:"CustomHeaders,omitempty" xml:"CustomHeaders,omitempty"`
	// The description of the protected object.
	//
	// example:
	//
	// ResourceTest
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The detailed parameter information of the protected object, which is converted into a JSON string constructed from a series of parameters.
	//
	// > The specific parameters involved vary depending on the specified **cloud service*	- (**Product**) and **protection mode*	- (**Pattern**). For more information, see **Parameter description of Detail for protected objects**.
	//
	// 	Notice: When **Product*	- is set to **ecs**, **clb4**, **clb7**, or **nlb**, domain names that are connected to regions in the Chinese mainland must have an ICP filing.</notice>
	//
	// This parameter is required.
	//
	// example:
	//
	// {\\"domain\\": \\"zhhclb4test096-05111.test.com\\"}
	Detail *string `json:"Detail,omitempty" xml:"Detail,omitempty"`
	// The ID of the WAF instance.
	//
	// > You can call the [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) operation to query the ID of the current WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_v3prepaid_public_cn-4xl*******
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the account to which the protected object belongs in a multi-account scenario. By default, the protected object belongs to the WAF administrator account.
	//
	// example:
	//
	// 123221XXX
	OwnerUserId *string `json:"OwnerUserId,omitempty" xml:"OwnerUserId,omitempty"`
	// The protection mode of the protected object. Valid values:
	//
	// - **domain**: domain name-based protection.
	//
	// - **multi_service**: hybrid cloud service-based protection.
	//
	// > Currently, only the following combinations are supported: When **Product*	- is set to **alb**, **ecs**, **clb4**, **clb7**, or **nlb**, **Pattern*	- must be set to **domain**. When **Product*	- is set to **waf**, **Pattern*	- must be set to **multi_service**.
	//
	// This parameter is required.
	//
	// example:
	//
	// domain
	Pattern *string `json:"Pattern,omitempty" xml:"Pattern,omitempty"`
	// The name of the cloud service. Valid values:
	//
	// - **alb**: ALB.
	//
	// - **ecs**: ECS.
	//
	// - **clb4**: CLB Layer 4.
	//
	// - **clb7**: CLB Layer 7.
	//
	// - **nlb**: NLB.
	//
	// - **waf**: WAF.
	//
	// This parameter is required.
	//
	// example:
	//
	// ecs
	Product *string `json:"Product,omitempty" xml:"Product,omitempty"`
	// The region where the WAF instance resides. Valid values:
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
	// >
	//
	// > - Only protected objects in hybrid cloud service mode support custom protected object names.
	//
	// example:
	//
	// abctest.com
	Resource *string `json:"Resource,omitempty" xml:"Resource,omitempty"`
	// The name of the protection group to which the protected object is added. This parameter is optional.
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
	// The source of the protected object. Valid values:
	//
	// - **custom**: user-defined.
	//
	// This parameter is required.
	//
	// example:
	//
	// custom
	ResourceOrigin *string `json:"ResourceOrigin,omitempty" xml:"ResourceOrigin,omitempty"`
	// The list of tags. You can specify up to 20 tags.
	Tag []*CreateDefenseResourceShrinkRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// Specifies whether to enable the X-Forwarded-For (XFF) proxy for the protected object. Valid values:
	//
	// - **0**: disabled (default).
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
