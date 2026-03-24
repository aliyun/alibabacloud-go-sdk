// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCloudResourceShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCloudResourceId(v string) *ModifyCloudResourceShrinkRequest
	GetCloudResourceId() *string
	SetInstanceId(v string) *ModifyCloudResourceShrinkRequest
	GetInstanceId() *string
	SetListenShrink(v string) *ModifyCloudResourceShrinkRequest
	GetListenShrink() *string
	SetRedirectShrink(v string) *ModifyCloudResourceShrinkRequest
	GetRedirectShrink() *string
	SetRegionId(v string) *ModifyCloudResourceShrinkRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *ModifyCloudResourceShrinkRequest
	GetResourceManagerResourceGroupId() *string
}

type ModifyCloudResourceShrinkRequest struct {
	// The ID of the cloud resource that is added to WAF.
	//
	// > Call [CreateCloudResource](https://help.aliyun.com/document_detail/2839876.html) to add a cloud resource. The resource ID is included in the response.
	//
	// example:
	//
	// lb-***-80-clb7
	CloudResourceId *string `json:"CloudResourceId,omitempty" xml:"CloudResourceId,omitempty"`
	// The ID of the WAF instance.
	//
	// > Call [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) to query the ID of the WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_v3prepaid_public_cn-***
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The listener configuration.
	//
	// This parameter is required.
	ListenShrink *string `json:"Listen,omitempty" xml:"Listen,omitempty"`
	// The forwarding configuration.
	RedirectShrink *string `json:"Redirect,omitempty" xml:"Redirect,omitempty"`
	// The region of the WAF instance. Valid values:
	//
	// - **cn-hangzhou**: the Chinese mainland.
	//
	// - **ap-southeast-1**: outside the Chinese mainland.
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
}

func (s ModifyCloudResourceShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyCloudResourceShrinkRequest) GoString() string {
	return s.String()
}

func (s *ModifyCloudResourceShrinkRequest) GetCloudResourceId() *string {
	return s.CloudResourceId
}

func (s *ModifyCloudResourceShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyCloudResourceShrinkRequest) GetListenShrink() *string {
	return s.ListenShrink
}

func (s *ModifyCloudResourceShrinkRequest) GetRedirectShrink() *string {
	return s.RedirectShrink
}

func (s *ModifyCloudResourceShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyCloudResourceShrinkRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *ModifyCloudResourceShrinkRequest) SetCloudResourceId(v string) *ModifyCloudResourceShrinkRequest {
	s.CloudResourceId = &v
	return s
}

func (s *ModifyCloudResourceShrinkRequest) SetInstanceId(v string) *ModifyCloudResourceShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyCloudResourceShrinkRequest) SetListenShrink(v string) *ModifyCloudResourceShrinkRequest {
	s.ListenShrink = &v
	return s
}

func (s *ModifyCloudResourceShrinkRequest) SetRedirectShrink(v string) *ModifyCloudResourceShrinkRequest {
	s.RedirectShrink = &v
	return s
}

func (s *ModifyCloudResourceShrinkRequest) SetRegionId(v string) *ModifyCloudResourceShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyCloudResourceShrinkRequest) SetResourceManagerResourceGroupId(v string) *ModifyCloudResourceShrinkRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *ModifyCloudResourceShrinkRequest) Validate() error {
	return dara.Validate(s)
}
