// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOpsItemShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCategory(v string) *CreateOpsItemShrinkRequest
	GetCategory() *string
	SetClientToken(v string) *CreateOpsItemShrinkRequest
	GetClientToken() *string
	SetDedupString(v string) *CreateOpsItemShrinkRequest
	GetDedupString() *string
	SetDescription(v string) *CreateOpsItemShrinkRequest
	GetDescription() *string
	SetPriority(v int32) *CreateOpsItemShrinkRequest
	GetPriority() *int32
	SetRegionId(v string) *CreateOpsItemShrinkRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreateOpsItemShrinkRequest
	GetResourceGroupId() *string
	SetResources(v string) *CreateOpsItemShrinkRequest
	GetResources() *string
	SetSeverity(v string) *CreateOpsItemShrinkRequest
	GetSeverity() *string
	SetSolutions(v string) *CreateOpsItemShrinkRequest
	GetSolutions() *string
	SetSource(v string) *CreateOpsItemShrinkRequest
	GetSource() *string
	SetTagsShrink(v string) *CreateOpsItemShrinkRequest
	GetTagsShrink() *string
	SetTitle(v string) *CreateOpsItemShrinkRequest
	GetTitle() *string
}

type CreateOpsItemShrinkRequest struct {
	// The category.
	//
	// Valid values:
	//
	// 	- Availability
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- Performance
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- Security
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- Cost
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- Recovery
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// This parameter is required.
	//
	// example:
	//
	// Security
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// example:
	//
	// 123e56767-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The string to be deduplicated.
	//
	// example:
	//
	// ecs_instance_Sys
	DedupString *string `json:"DedupString,omitempty" xml:"DedupString,omitempty"`
	// The description of the operation.
	//
	// example:
	//
	// OpsItem
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The priority. Valid values: 1 to 5. 1 indicates the highest priority.
	//
	// example:
	//
	// 4
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
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
	// rg-acfmxsn4m4******
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The Alibaba Cloud Resource Names (ARNs) of the associated resources.
	//
	// example:
	//
	// [\\"acs:oos:cn-hangzhou:1563457855438522:application/test-33333/applicationgroup/default-cn-hangzhou-1\\"]
	Resources *string `json:"Resources,omitempty" xml:"Resources,omitempty"`
	// The severity level.
	//
	// Valid values:
	//
	// 	- High
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- Low
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- Medium
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- Critical
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// This parameter is required.
	//
	// example:
	//
	// Medium
	Severity *string `json:"Severity,omitempty" xml:"Severity,omitempty"`
	// The solutions.
	//
	// example:
	//
	// [{\\n \\\\"priority\\\\":3,\\n \\\\"type\\\\":\\\\"SingleAZEcs\\\\",\\n \\\\"url\\\\":\\\\"http://ecs.consle.aliyuncs.com\\\\",\\n \\\\"description\\\\":\\\\"Create Elastic Compute Service (ECS) instances in different zones and import them to an application group.\\\\"\\n}]
	Solutions *string `json:"Solutions,omitempty" xml:"Solutions,omitempty"`
	// The source business.
	//
	// This parameter is required.
	//
	// example:
	//
	// /aliyun/ecs
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The tags.
	//
	// example:
	//
	// {
	//
	//       "k1": "v1",
	//
	//       "k2": "v2"
	//
	// }
	TagsShrink *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The title of the O\\&M item.
	//
	// This parameter is required.
	//
	// example:
	//
	// ECS reboot is scheduled
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s CreateOpsItemShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateOpsItemShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateOpsItemShrinkRequest) GetCategory() *string {
	return s.Category
}

func (s *CreateOpsItemShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateOpsItemShrinkRequest) GetDedupString() *string {
	return s.DedupString
}

func (s *CreateOpsItemShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateOpsItemShrinkRequest) GetPriority() *int32 {
	return s.Priority
}

func (s *CreateOpsItemShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateOpsItemShrinkRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateOpsItemShrinkRequest) GetResources() *string {
	return s.Resources
}

func (s *CreateOpsItemShrinkRequest) GetSeverity() *string {
	return s.Severity
}

func (s *CreateOpsItemShrinkRequest) GetSolutions() *string {
	return s.Solutions
}

func (s *CreateOpsItemShrinkRequest) GetSource() *string {
	return s.Source
}

func (s *CreateOpsItemShrinkRequest) GetTagsShrink() *string {
	return s.TagsShrink
}

func (s *CreateOpsItemShrinkRequest) GetTitle() *string {
	return s.Title
}

func (s *CreateOpsItemShrinkRequest) SetCategory(v string) *CreateOpsItemShrinkRequest {
	s.Category = &v
	return s
}

func (s *CreateOpsItemShrinkRequest) SetClientToken(v string) *CreateOpsItemShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateOpsItemShrinkRequest) SetDedupString(v string) *CreateOpsItemShrinkRequest {
	s.DedupString = &v
	return s
}

func (s *CreateOpsItemShrinkRequest) SetDescription(v string) *CreateOpsItemShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreateOpsItemShrinkRequest) SetPriority(v int32) *CreateOpsItemShrinkRequest {
	s.Priority = &v
	return s
}

func (s *CreateOpsItemShrinkRequest) SetRegionId(v string) *CreateOpsItemShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *CreateOpsItemShrinkRequest) SetResourceGroupId(v string) *CreateOpsItemShrinkRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateOpsItemShrinkRequest) SetResources(v string) *CreateOpsItemShrinkRequest {
	s.Resources = &v
	return s
}

func (s *CreateOpsItemShrinkRequest) SetSeverity(v string) *CreateOpsItemShrinkRequest {
	s.Severity = &v
	return s
}

func (s *CreateOpsItemShrinkRequest) SetSolutions(v string) *CreateOpsItemShrinkRequest {
	s.Solutions = &v
	return s
}

func (s *CreateOpsItemShrinkRequest) SetSource(v string) *CreateOpsItemShrinkRequest {
	s.Source = &v
	return s
}

func (s *CreateOpsItemShrinkRequest) SetTagsShrink(v string) *CreateOpsItemShrinkRequest {
	s.TagsShrink = &v
	return s
}

func (s *CreateOpsItemShrinkRequest) SetTitle(v string) *CreateOpsItemShrinkRequest {
	s.Title = &v
	return s
}

func (s *CreateOpsItemShrinkRequest) Validate() error {
	return dara.Validate(s)
}
