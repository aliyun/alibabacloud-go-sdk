// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateOpsItemRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCategory(v string) *UpdateOpsItemRequest
	GetCategory() *string
	SetClientToken(v string) *UpdateOpsItemRequest
	GetClientToken() *string
	SetDedupString(v string) *UpdateOpsItemRequest
	GetDedupString() *string
	SetDescription(v string) *UpdateOpsItemRequest
	GetDescription() *string
	SetOpsItemId(v string) *UpdateOpsItemRequest
	GetOpsItemId() *string
	SetPriority(v int32) *UpdateOpsItemRequest
	GetPriority() *int32
	SetRegionId(v string) *UpdateOpsItemRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *UpdateOpsItemRequest
	GetResourceGroupId() *string
	SetResources(v string) *UpdateOpsItemRequest
	GetResources() *string
	SetSeverity(v string) *UpdateOpsItemRequest
	GetSeverity() *string
	SetSolutions(v string) *UpdateOpsItemRequest
	GetSolutions() *string
	SetSource(v string) *UpdateOpsItemRequest
	GetSource() *string
	SetStatus(v string) *UpdateOpsItemRequest
	GetStatus() *string
	SetTags(v map[string]interface{}) *UpdateOpsItemRequest
	GetTags() map[string]interface{}
	SetTitle(v string) *UpdateOpsItemRequest
	GetTitle() *string
}

type UpdateOpsItemRequest struct {
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
	Tags map[string]interface{} `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The title of the O\\&M item.
	//
	// example:
	//
	// Test
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s UpdateOpsItemRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateOpsItemRequest) GoString() string {
	return s.String()
}

func (s *UpdateOpsItemRequest) GetCategory() *string {
	return s.Category
}

func (s *UpdateOpsItemRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateOpsItemRequest) GetDedupString() *string {
	return s.DedupString
}

func (s *UpdateOpsItemRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateOpsItemRequest) GetOpsItemId() *string {
	return s.OpsItemId
}

func (s *UpdateOpsItemRequest) GetPriority() *int32 {
	return s.Priority
}

func (s *UpdateOpsItemRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateOpsItemRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *UpdateOpsItemRequest) GetResources() *string {
	return s.Resources
}

func (s *UpdateOpsItemRequest) GetSeverity() *string {
	return s.Severity
}

func (s *UpdateOpsItemRequest) GetSolutions() *string {
	return s.Solutions
}

func (s *UpdateOpsItemRequest) GetSource() *string {
	return s.Source
}

func (s *UpdateOpsItemRequest) GetStatus() *string {
	return s.Status
}

func (s *UpdateOpsItemRequest) GetTags() map[string]interface{} {
	return s.Tags
}

func (s *UpdateOpsItemRequest) GetTitle() *string {
	return s.Title
}

func (s *UpdateOpsItemRequest) SetCategory(v string) *UpdateOpsItemRequest {
	s.Category = &v
	return s
}

func (s *UpdateOpsItemRequest) SetClientToken(v string) *UpdateOpsItemRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateOpsItemRequest) SetDedupString(v string) *UpdateOpsItemRequest {
	s.DedupString = &v
	return s
}

func (s *UpdateOpsItemRequest) SetDescription(v string) *UpdateOpsItemRequest {
	s.Description = &v
	return s
}

func (s *UpdateOpsItemRequest) SetOpsItemId(v string) *UpdateOpsItemRequest {
	s.OpsItemId = &v
	return s
}

func (s *UpdateOpsItemRequest) SetPriority(v int32) *UpdateOpsItemRequest {
	s.Priority = &v
	return s
}

func (s *UpdateOpsItemRequest) SetRegionId(v string) *UpdateOpsItemRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateOpsItemRequest) SetResourceGroupId(v string) *UpdateOpsItemRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *UpdateOpsItemRequest) SetResources(v string) *UpdateOpsItemRequest {
	s.Resources = &v
	return s
}

func (s *UpdateOpsItemRequest) SetSeverity(v string) *UpdateOpsItemRequest {
	s.Severity = &v
	return s
}

func (s *UpdateOpsItemRequest) SetSolutions(v string) *UpdateOpsItemRequest {
	s.Solutions = &v
	return s
}

func (s *UpdateOpsItemRequest) SetSource(v string) *UpdateOpsItemRequest {
	s.Source = &v
	return s
}

func (s *UpdateOpsItemRequest) SetStatus(v string) *UpdateOpsItemRequest {
	s.Status = &v
	return s
}

func (s *UpdateOpsItemRequest) SetTags(v map[string]interface{}) *UpdateOpsItemRequest {
	s.Tags = v
	return s
}

func (s *UpdateOpsItemRequest) SetTitle(v string) *UpdateOpsItemRequest {
	s.Title = &v
	return s
}

func (s *UpdateOpsItemRequest) Validate() error {
	return dara.Validate(s)
}
