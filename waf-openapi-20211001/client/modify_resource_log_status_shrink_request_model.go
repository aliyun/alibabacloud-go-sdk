// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyResourceLogStatusShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ModifyResourceLogStatusShrinkRequest
	GetInstanceId() *string
	SetRegionId(v string) *ModifyResourceLogStatusShrinkRequest
	GetRegionId() *string
	SetResource(v string) *ModifyResourceLogStatusShrinkRequest
	GetResource() *string
	SetResourceManagerResourceGroupId(v string) *ModifyResourceLogStatusShrinkRequest
	GetResourceManagerResourceGroupId() *string
	SetStatus(v bool) *ModifyResourceLogStatusShrinkRequest
	GetStatus() *bool
	SetTraceConfigShrink(v string) *ModifyResourceLogStatusShrinkRequest
	GetTraceConfigShrink() *string
	SetTraceStatus(v bool) *ModifyResourceLogStatusShrinkRequest
	GetTraceStatus() *bool
}

type ModifyResourceLogStatusShrinkRequest struct {
	// The ID of the WAF instance.
	//
	// > Call the [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) operation to query the ID of the WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf-cn-zz11zcl****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region of the WAF instance. Valid values:
	//
	// - **cn-hangzhou**: The Chinese mainland.
	//
	// - **ap-southeast-1**: Outside the Chinese mainland.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The protected object on which you want to manage the log collection feature.
	//
	// This parameter is required.
	//
	// example:
	//
	// alb-wewbb23dfsetetcic1242-0****
	Resource *string `json:"Resource,omitempty" xml:"Resource,omitempty"`
	// The ID of the Alibaba Cloud resource group.
	//
	// example:
	//
	// rg-acfm***q
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
	// Specifies whether to enable the log collection feature for the protected object. Valid values:
	//
	// - **true**: Enables the feature.
	//
	// - **false**: Disables the feature.
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	Status            *bool   `json:"Status,omitempty" xml:"Status,omitempty"`
	TraceConfigShrink *string `json:"TraceConfig,omitempty" xml:"TraceConfig,omitempty"`
	TraceStatus       *bool   `json:"TraceStatus,omitempty" xml:"TraceStatus,omitempty"`
}

func (s ModifyResourceLogStatusShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyResourceLogStatusShrinkRequest) GoString() string {
	return s.String()
}

func (s *ModifyResourceLogStatusShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyResourceLogStatusShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyResourceLogStatusShrinkRequest) GetResource() *string {
	return s.Resource
}

func (s *ModifyResourceLogStatusShrinkRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *ModifyResourceLogStatusShrinkRequest) GetStatus() *bool {
	return s.Status
}

func (s *ModifyResourceLogStatusShrinkRequest) GetTraceConfigShrink() *string {
	return s.TraceConfigShrink
}

func (s *ModifyResourceLogStatusShrinkRequest) GetTraceStatus() *bool {
	return s.TraceStatus
}

func (s *ModifyResourceLogStatusShrinkRequest) SetInstanceId(v string) *ModifyResourceLogStatusShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyResourceLogStatusShrinkRequest) SetRegionId(v string) *ModifyResourceLogStatusShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyResourceLogStatusShrinkRequest) SetResource(v string) *ModifyResourceLogStatusShrinkRequest {
	s.Resource = &v
	return s
}

func (s *ModifyResourceLogStatusShrinkRequest) SetResourceManagerResourceGroupId(v string) *ModifyResourceLogStatusShrinkRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *ModifyResourceLogStatusShrinkRequest) SetStatus(v bool) *ModifyResourceLogStatusShrinkRequest {
	s.Status = &v
	return s
}

func (s *ModifyResourceLogStatusShrinkRequest) SetTraceConfigShrink(v string) *ModifyResourceLogStatusShrinkRequest {
	s.TraceConfigShrink = &v
	return s
}

func (s *ModifyResourceLogStatusShrinkRequest) SetTraceStatus(v bool) *ModifyResourceLogStatusShrinkRequest {
	s.TraceStatus = &v
	return s
}

func (s *ModifyResourceLogStatusShrinkRequest) Validate() error {
	return dara.Validate(s)
}
