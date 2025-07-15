// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateOpsItemShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCategory(v string) *UpdateOpsItemShrinkRequest
	GetCategory() *string
	SetClientToken(v string) *UpdateOpsItemShrinkRequest
	GetClientToken() *string
	SetDedupString(v string) *UpdateOpsItemShrinkRequest
	GetDedupString() *string
	SetDescription(v string) *UpdateOpsItemShrinkRequest
	GetDescription() *string
	SetOpsItemId(v string) *UpdateOpsItemShrinkRequest
	GetOpsItemId() *string
	SetPriority(v int32) *UpdateOpsItemShrinkRequest
	GetPriority() *int32
	SetRegionId(v string) *UpdateOpsItemShrinkRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *UpdateOpsItemShrinkRequest
	GetResourceGroupId() *string
	SetResources(v string) *UpdateOpsItemShrinkRequest
	GetResources() *string
	SetSeverity(v string) *UpdateOpsItemShrinkRequest
	GetSeverity() *string
	SetSolutions(v string) *UpdateOpsItemShrinkRequest
	GetSolutions() *string
	SetSource(v string) *UpdateOpsItemShrinkRequest
	GetSource() *string
	SetStatus(v string) *UpdateOpsItemShrinkRequest
	GetStatus() *string
	SetTagsShrink(v string) *UpdateOpsItemShrinkRequest
	GetTagsShrink() *string
	SetTitle(v string) *UpdateOpsItemShrinkRequest
	GetTitle() *string
}

type UpdateOpsItemShrinkRequest struct {
	// The category.
	//
	// example:
	//
	// Security
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// example:
	//
	// DASKJJLKADS-AHKLJHJSAKL-AJK
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The string to be deduplicated.
	//
	// example:
	//
	// ecs_instance_SystemMaintenance.Reboot
	DedupString *string `json:"DedupString,omitempty" xml:"DedupString,omitempty"`
	// The description of the O\\&M item.
	//
	// example:
	//
	// test-update
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the O\\&M item.
	//
	// example:
	//
	// oi-e2264dcf040c472598e9
	OpsItemId *string `json:"OpsItemId,omitempty" xml:"OpsItemId,omitempty"`
	// The priority.
	//
	// example:
	//
	// 2
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
	// The Alibaba Resource Names (ARNs) of the associated resources.
	//
	// example:
	//
	// [\\"arn:acs:ecs:cn-heyuan:1139354755361920:instance/i-f8z928h7aqotd3o65032\\"]
	Resources *string `json:"Resources,omitempty" xml:"Resources,omitempty"`
	// The severity level.
	//
	// example:
	//
	// Medium
	Severity *string `json:"Severity,omitempty" xml:"Severity,omitempty"`
	// The solutions.
	//
	// example:
	//
	// [{\\n \\\\"priority\\\\":3,\\n \\\\"type\\\\":\\\\"url\\\\",\\n \\\\"url\\\\":\\\\"https://example.com\\\\",\\n \\\\"description\\\\":\\\\"Specify a cross-zone high availability cluster. \\\\"\\n}]
	Solutions *string `json:"Solutions,omitempty" xml:"Solutions,omitempty"`
	// The source business.
	//
	// example:
	//
	// /aliyun/ecs
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The status.
	//
	// example:
	//
	// Open
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
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
	// example:
	//
	// Test
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s UpdateOpsItemShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateOpsItemShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateOpsItemShrinkRequest) GetCategory() *string {
	return s.Category
}

func (s *UpdateOpsItemShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateOpsItemShrinkRequest) GetDedupString() *string {
	return s.DedupString
}

func (s *UpdateOpsItemShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateOpsItemShrinkRequest) GetOpsItemId() *string {
	return s.OpsItemId
}

func (s *UpdateOpsItemShrinkRequest) GetPriority() *int32 {
	return s.Priority
}

func (s *UpdateOpsItemShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateOpsItemShrinkRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *UpdateOpsItemShrinkRequest) GetResources() *string {
	return s.Resources
}

func (s *UpdateOpsItemShrinkRequest) GetSeverity() *string {
	return s.Severity
}

func (s *UpdateOpsItemShrinkRequest) GetSolutions() *string {
	return s.Solutions
}

func (s *UpdateOpsItemShrinkRequest) GetSource() *string {
	return s.Source
}

func (s *UpdateOpsItemShrinkRequest) GetStatus() *string {
	return s.Status
}

func (s *UpdateOpsItemShrinkRequest) GetTagsShrink() *string {
	return s.TagsShrink
}

func (s *UpdateOpsItemShrinkRequest) GetTitle() *string {
	return s.Title
}

func (s *UpdateOpsItemShrinkRequest) SetCategory(v string) *UpdateOpsItemShrinkRequest {
	s.Category = &v
	return s
}

func (s *UpdateOpsItemShrinkRequest) SetClientToken(v string) *UpdateOpsItemShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateOpsItemShrinkRequest) SetDedupString(v string) *UpdateOpsItemShrinkRequest {
	s.DedupString = &v
	return s
}

func (s *UpdateOpsItemShrinkRequest) SetDescription(v string) *UpdateOpsItemShrinkRequest {
	s.Description = &v
	return s
}

func (s *UpdateOpsItemShrinkRequest) SetOpsItemId(v string) *UpdateOpsItemShrinkRequest {
	s.OpsItemId = &v
	return s
}

func (s *UpdateOpsItemShrinkRequest) SetPriority(v int32) *UpdateOpsItemShrinkRequest {
	s.Priority = &v
	return s
}

func (s *UpdateOpsItemShrinkRequest) SetRegionId(v string) *UpdateOpsItemShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateOpsItemShrinkRequest) SetResourceGroupId(v string) *UpdateOpsItemShrinkRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *UpdateOpsItemShrinkRequest) SetResources(v string) *UpdateOpsItemShrinkRequest {
	s.Resources = &v
	return s
}

func (s *UpdateOpsItemShrinkRequest) SetSeverity(v string) *UpdateOpsItemShrinkRequest {
	s.Severity = &v
	return s
}

func (s *UpdateOpsItemShrinkRequest) SetSolutions(v string) *UpdateOpsItemShrinkRequest {
	s.Solutions = &v
	return s
}

func (s *UpdateOpsItemShrinkRequest) SetSource(v string) *UpdateOpsItemShrinkRequest {
	s.Source = &v
	return s
}

func (s *UpdateOpsItemShrinkRequest) SetStatus(v string) *UpdateOpsItemShrinkRequest {
	s.Status = &v
	return s
}

func (s *UpdateOpsItemShrinkRequest) SetTagsShrink(v string) *UpdateOpsItemShrinkRequest {
	s.TagsShrink = &v
	return s
}

func (s *UpdateOpsItemShrinkRequest) SetTitle(v string) *UpdateOpsItemShrinkRequest {
	s.Title = &v
	return s
}

func (s *UpdateOpsItemShrinkRequest) Validate() error {
	return dara.Validate(s)
}
