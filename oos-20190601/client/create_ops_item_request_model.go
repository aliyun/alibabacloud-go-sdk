// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOpsItemRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCategory(v string) *CreateOpsItemRequest
	GetCategory() *string
	SetClientToken(v string) *CreateOpsItemRequest
	GetClientToken() *string
	SetDedupString(v string) *CreateOpsItemRequest
	GetDedupString() *string
	SetDescription(v string) *CreateOpsItemRequest
	GetDescription() *string
	SetPriority(v int32) *CreateOpsItemRequest
	GetPriority() *int32
	SetRegionId(v string) *CreateOpsItemRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreateOpsItemRequest
	GetResourceGroupId() *string
	SetResources(v string) *CreateOpsItemRequest
	GetResources() *string
	SetSeverity(v string) *CreateOpsItemRequest
	GetSeverity() *string
	SetSolutions(v string) *CreateOpsItemRequest
	GetSolutions() *string
	SetSource(v string) *CreateOpsItemRequest
	GetSource() *string
	SetTags(v map[string]interface{}) *CreateOpsItemRequest
	GetTags() map[string]interface{}
	SetTitle(v string) *CreateOpsItemRequest
	GetTitle() *string
}

type CreateOpsItemRequest struct {
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
	Tags map[string]interface{} `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The title of the O\\&M item.
	//
	// This parameter is required.
	//
	// example:
	//
	// ECS reboot is scheduled
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s CreateOpsItemRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateOpsItemRequest) GoString() string {
	return s.String()
}

func (s *CreateOpsItemRequest) GetCategory() *string {
	return s.Category
}

func (s *CreateOpsItemRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateOpsItemRequest) GetDedupString() *string {
	return s.DedupString
}

func (s *CreateOpsItemRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateOpsItemRequest) GetPriority() *int32 {
	return s.Priority
}

func (s *CreateOpsItemRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateOpsItemRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateOpsItemRequest) GetResources() *string {
	return s.Resources
}

func (s *CreateOpsItemRequest) GetSeverity() *string {
	return s.Severity
}

func (s *CreateOpsItemRequest) GetSolutions() *string {
	return s.Solutions
}

func (s *CreateOpsItemRequest) GetSource() *string {
	return s.Source
}

func (s *CreateOpsItemRequest) GetTags() map[string]interface{} {
	return s.Tags
}

func (s *CreateOpsItemRequest) GetTitle() *string {
	return s.Title
}

func (s *CreateOpsItemRequest) SetCategory(v string) *CreateOpsItemRequest {
	s.Category = &v
	return s
}

func (s *CreateOpsItemRequest) SetClientToken(v string) *CreateOpsItemRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateOpsItemRequest) SetDedupString(v string) *CreateOpsItemRequest {
	s.DedupString = &v
	return s
}

func (s *CreateOpsItemRequest) SetDescription(v string) *CreateOpsItemRequest {
	s.Description = &v
	return s
}

func (s *CreateOpsItemRequest) SetPriority(v int32) *CreateOpsItemRequest {
	s.Priority = &v
	return s
}

func (s *CreateOpsItemRequest) SetRegionId(v string) *CreateOpsItemRequest {
	s.RegionId = &v
	return s
}

func (s *CreateOpsItemRequest) SetResourceGroupId(v string) *CreateOpsItemRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateOpsItemRequest) SetResources(v string) *CreateOpsItemRequest {
	s.Resources = &v
	return s
}

func (s *CreateOpsItemRequest) SetSeverity(v string) *CreateOpsItemRequest {
	s.Severity = &v
	return s
}

func (s *CreateOpsItemRequest) SetSolutions(v string) *CreateOpsItemRequest {
	s.Solutions = &v
	return s
}

func (s *CreateOpsItemRequest) SetSource(v string) *CreateOpsItemRequest {
	s.Source = &v
	return s
}

func (s *CreateOpsItemRequest) SetTags(v map[string]interface{}) *CreateOpsItemRequest {
	s.Tags = v
	return s
}

func (s *CreateOpsItemRequest) SetTitle(v string) *CreateOpsItemRequest {
	s.Title = &v
	return s
}

func (s *CreateOpsItemRequest) Validate() error {
	return dara.Validate(s)
}
