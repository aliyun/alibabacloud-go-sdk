// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDomainShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessType(v string) *CreateDomainShrinkRequest
	GetAccessType() *string
	SetDomain(v string) *CreateDomainShrinkRequest
	GetDomain() *string
	SetInstanceId(v string) *CreateDomainShrinkRequest
	GetInstanceId() *string
	SetListenShrink(v string) *CreateDomainShrinkRequest
	GetListenShrink() *string
	SetRedirectShrink(v string) *CreateDomainShrinkRequest
	GetRedirectShrink() *string
	SetRegionId(v string) *CreateDomainShrinkRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *CreateDomainShrinkRequest
	GetResourceManagerResourceGroupId() *string
	SetTag(v []*CreateDomainShrinkRequestTag) *CreateDomainShrinkRequest
	GetTag() []*CreateDomainShrinkRequestTag
}

type CreateDomainShrinkRequest struct {
	// The mode in which you want to add the domain name to WAF. Valid values:
	//
	// 	- **share:*	- adds the domain name to WAF in CNAME record mode. This is the default value.
	//
	// 	- **hybrid_cloud_cname:*	- adds the domain name to WAF in hybrid cloud reverse proxy mode.
	//
	// example:
	//
	// share
	AccessType *string `json:"AccessType,omitempty" xml:"AccessType,omitempty"`
	// The domain name that you want to add to WAF.
	//
	// This parameter is required.
	//
	// example:
	//
	// www.aliyundoc.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The ID of the Web Application Firewall (WAF) instance.
	//
	// > You can call the [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) operation to obtain the ID of the WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_cdnsdf3****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The listener configurations.
	//
	// This parameter is required.
	ListenShrink *string `json:"Listen,omitempty" xml:"Listen,omitempty"`
	// The forwarding configurations.
	//
	// This parameter is required.
	RedirectShrink *string `json:"Redirect,omitempty" xml:"Redirect,omitempty"`
	// The region where the WAF instance resides. Valid values:
	//
	// 	- **cn-hangzhou**: the Chinese mainland
	//
	// 	- **ap-southeast-1**: outside the Chinese mainland
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfm***q
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
	// The tags. You can specify up to 20 tags.
	Tag []*CreateDomainShrinkRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s CreateDomainShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDomainShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateDomainShrinkRequest) GetAccessType() *string {
	return s.AccessType
}

func (s *CreateDomainShrinkRequest) GetDomain() *string {
	return s.Domain
}

func (s *CreateDomainShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateDomainShrinkRequest) GetListenShrink() *string {
	return s.ListenShrink
}

func (s *CreateDomainShrinkRequest) GetRedirectShrink() *string {
	return s.RedirectShrink
}

func (s *CreateDomainShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateDomainShrinkRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *CreateDomainShrinkRequest) GetTag() []*CreateDomainShrinkRequestTag {
	return s.Tag
}

func (s *CreateDomainShrinkRequest) SetAccessType(v string) *CreateDomainShrinkRequest {
	s.AccessType = &v
	return s
}

func (s *CreateDomainShrinkRequest) SetDomain(v string) *CreateDomainShrinkRequest {
	s.Domain = &v
	return s
}

func (s *CreateDomainShrinkRequest) SetInstanceId(v string) *CreateDomainShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateDomainShrinkRequest) SetListenShrink(v string) *CreateDomainShrinkRequest {
	s.ListenShrink = &v
	return s
}

func (s *CreateDomainShrinkRequest) SetRedirectShrink(v string) *CreateDomainShrinkRequest {
	s.RedirectShrink = &v
	return s
}

func (s *CreateDomainShrinkRequest) SetRegionId(v string) *CreateDomainShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *CreateDomainShrinkRequest) SetResourceManagerResourceGroupId(v string) *CreateDomainShrinkRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *CreateDomainShrinkRequest) SetTag(v []*CreateDomainShrinkRequestTag) *CreateDomainShrinkRequest {
	s.Tag = v
	return s
}

func (s *CreateDomainShrinkRequest) Validate() error {
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

type CreateDomainShrinkRequestTag struct {
	// The key of the tag.
	//
	// example:
	//
	// Tagkey1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag.
	//
	// example:
	//
	// TagValue1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateDomainShrinkRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateDomainShrinkRequestTag) GoString() string {
	return s.String()
}

func (s *CreateDomainShrinkRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateDomainShrinkRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateDomainShrinkRequestTag) SetKey(v string) *CreateDomainShrinkRequestTag {
	s.Key = &v
	return s
}

func (s *CreateDomainShrinkRequestTag) SetValue(v string) *CreateDomainShrinkRequestTag {
	s.Value = &v
	return s
}

func (s *CreateDomainShrinkRequestTag) Validate() error {
	return dara.Validate(s)
}
