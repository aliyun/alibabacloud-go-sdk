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
	// The list of specified header fields.
	//
	// > When XffStatus is set to 1, the first IP address in the specified header field is used as the originating IP address of the client to prevent XFF spoofing. If multiple headers are specified, the system attempts to obtain the originating IP address in sequence. If the first header does not contain an IP address, the system checks the second header, and so on. If no IP address is found in any of the specified headers, the first IP address in the X-Forwarded-For header is used. When XffStatus is set to 1, it indicates that the first IP address is obtained.
	CustomHeaders []*string `json:"CustomHeaders,omitempty" xml:"CustomHeaders,omitempty" type:"Repeated"`
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
	Tag []*CreateDefenseResourceRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
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
