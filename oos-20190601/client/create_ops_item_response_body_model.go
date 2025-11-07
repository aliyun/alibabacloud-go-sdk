// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOpsItemResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOpsItem(v *CreateOpsItemResponseBodyOpsItem) *CreateOpsItemResponseBody
	GetOpsItem() *CreateOpsItemResponseBodyOpsItem
	SetRequestId(v string) *CreateOpsItemResponseBody
	GetRequestId() *string
}

type CreateOpsItemResponseBody struct {
	// The information about the O\\&M item.
	OpsItem *CreateOpsItemResponseBodyOpsItem `json:"OpsItem,omitempty" xml:"OpsItem,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// DA4F08D4-DA54-5407-84B9-108C71D75B17
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateOpsItemResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateOpsItemResponseBody) GoString() string {
	return s.String()
}

func (s *CreateOpsItemResponseBody) GetOpsItem() *CreateOpsItemResponseBodyOpsItem {
	return s.OpsItem
}

func (s *CreateOpsItemResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateOpsItemResponseBody) SetOpsItem(v *CreateOpsItemResponseBodyOpsItem) *CreateOpsItemResponseBody {
	s.OpsItem = v
	return s
}

func (s *CreateOpsItemResponseBody) SetRequestId(v string) *CreateOpsItemResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateOpsItemResponseBody) Validate() error {
	if s.OpsItem != nil {
		if err := s.OpsItem.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateOpsItemResponseBodyOpsItem struct {
	// The attributes of the O\\&M item.
	//
	// example:
	//
	// {\\"regionId\\":\\"cn-zhangjiakou\\",\\"appId\\":\\"992BKO1X75GRQ4E8\\",\\"instanceId\\":\\"i-8vbcymxagqsyyyjppbr4\\",\\"instance_name\\":\\"cdae\\"}
	Attributes *string `json:"Attributes,omitempty" xml:"Attributes,omitempty"`
	// The category of the O\\&M item.
	//
	// example:
	//
	// Security
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The time when the O\\&M item was created.
	//
	// example:
	//
	// 2023-03-24T03:55Z
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	// The user who created the O\\&M item.
	//
	// example:
	//
	// root(130900000)
	CreatedBy *string `json:"CreatedBy,omitempty" xml:"CreatedBy,omitempty"`
	// The description of the O\\&M item.
	//
	// example:
	//
	// OpsItem
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The user who last modified the O\\&M item.
	//
	// example:
	//
	// root(130900000)
	LastModifiedBy *string `json:"LastModifiedBy,omitempty" xml:"LastModifiedBy,omitempty"`
	// The ID of the O\\&M item.
	//
	// example:
	//
	// oi-dba9c6eec9254a4d87c1
	OpsItemId *string `json:"OpsItemId,omitempty" xml:"OpsItemId,omitempty"`
	// The priority of the O\\&M item.
	//
	// example:
	//
	// 2
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmxsn4m4******
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The ARNs of the associated resources.
	//
	// example:
	//
	// [\\"acs:oos:cn-hangzhou:1563457855438522:application/dingTest/applicationgroup/fltest\\"]
	Resources *string `json:"Resources,omitempty" xml:"Resources,omitempty"`
	// The severity level of the O\\&M item.
	//
	// example:
	//
	// Medium
	Severity *string `json:"Severity,omitempty" xml:"Severity,omitempty"`
	// The solutions.
	//
	// example:
	//
	// [{\\n \\\\"priority\\\\":3,\\n \\\\"type\\\\":\\\\"url\\\\",\\n \\\\"url\\\\":\\\\"https://example..com\\\\",\\n \\\\"description\\\\":\\\\"Specify a cross-zone high availability cluster. \\\\"\\n}]
	Solutions *string `json:"Solutions,omitempty" xml:"Solutions,omitempty"`
	// The source business of the O\\&M item.
	//
	// example:
	//
	// /aliyun/ecs
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The state of the O\\&M item.
	//
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tags of the O\\&M item.
	//
	// example:
	//
	// {"k1": "v1"}
	Tags map[string]interface{} `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The title of the O\\&M item.
	//
	// example:
	//
	// ECS reboot is scheduled
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// The time when the O\\&M item was updated.
	//
	// example:
	//
	// 2023-03-24T03:55Z
	UpdateDate *string `json:"UpdateDate,omitempty" xml:"UpdateDate,omitempty"`
}

func (s CreateOpsItemResponseBodyOpsItem) String() string {
	return dara.Prettify(s)
}

func (s CreateOpsItemResponseBodyOpsItem) GoString() string {
	return s.String()
}

func (s *CreateOpsItemResponseBodyOpsItem) GetAttributes() *string {
	return s.Attributes
}

func (s *CreateOpsItemResponseBodyOpsItem) GetCategory() *string {
	return s.Category
}

func (s *CreateOpsItemResponseBodyOpsItem) GetCreateDate() *string {
	return s.CreateDate
}

func (s *CreateOpsItemResponseBodyOpsItem) GetCreatedBy() *string {
	return s.CreatedBy
}

func (s *CreateOpsItemResponseBodyOpsItem) GetDescription() *string {
	return s.Description
}

func (s *CreateOpsItemResponseBodyOpsItem) GetLastModifiedBy() *string {
	return s.LastModifiedBy
}

func (s *CreateOpsItemResponseBodyOpsItem) GetOpsItemId() *string {
	return s.OpsItemId
}

func (s *CreateOpsItemResponseBodyOpsItem) GetPriority() *int32 {
	return s.Priority
}

func (s *CreateOpsItemResponseBodyOpsItem) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateOpsItemResponseBodyOpsItem) GetResources() *string {
	return s.Resources
}

func (s *CreateOpsItemResponseBodyOpsItem) GetSeverity() *string {
	return s.Severity
}

func (s *CreateOpsItemResponseBodyOpsItem) GetSolutions() *string {
	return s.Solutions
}

func (s *CreateOpsItemResponseBodyOpsItem) GetSource() *string {
	return s.Source
}

func (s *CreateOpsItemResponseBodyOpsItem) GetStatus() *string {
	return s.Status
}

func (s *CreateOpsItemResponseBodyOpsItem) GetTags() map[string]interface{} {
	return s.Tags
}

func (s *CreateOpsItemResponseBodyOpsItem) GetTitle() *string {
	return s.Title
}

func (s *CreateOpsItemResponseBodyOpsItem) GetUpdateDate() *string {
	return s.UpdateDate
}

func (s *CreateOpsItemResponseBodyOpsItem) SetAttributes(v string) *CreateOpsItemResponseBodyOpsItem {
	s.Attributes = &v
	return s
}

func (s *CreateOpsItemResponseBodyOpsItem) SetCategory(v string) *CreateOpsItemResponseBodyOpsItem {
	s.Category = &v
	return s
}

func (s *CreateOpsItemResponseBodyOpsItem) SetCreateDate(v string) *CreateOpsItemResponseBodyOpsItem {
	s.CreateDate = &v
	return s
}

func (s *CreateOpsItemResponseBodyOpsItem) SetCreatedBy(v string) *CreateOpsItemResponseBodyOpsItem {
	s.CreatedBy = &v
	return s
}

func (s *CreateOpsItemResponseBodyOpsItem) SetDescription(v string) *CreateOpsItemResponseBodyOpsItem {
	s.Description = &v
	return s
}

func (s *CreateOpsItemResponseBodyOpsItem) SetLastModifiedBy(v string) *CreateOpsItemResponseBodyOpsItem {
	s.LastModifiedBy = &v
	return s
}

func (s *CreateOpsItemResponseBodyOpsItem) SetOpsItemId(v string) *CreateOpsItemResponseBodyOpsItem {
	s.OpsItemId = &v
	return s
}

func (s *CreateOpsItemResponseBodyOpsItem) SetPriority(v int32) *CreateOpsItemResponseBodyOpsItem {
	s.Priority = &v
	return s
}

func (s *CreateOpsItemResponseBodyOpsItem) SetResourceGroupId(v string) *CreateOpsItemResponseBodyOpsItem {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateOpsItemResponseBodyOpsItem) SetResources(v string) *CreateOpsItemResponseBodyOpsItem {
	s.Resources = &v
	return s
}

func (s *CreateOpsItemResponseBodyOpsItem) SetSeverity(v string) *CreateOpsItemResponseBodyOpsItem {
	s.Severity = &v
	return s
}

func (s *CreateOpsItemResponseBodyOpsItem) SetSolutions(v string) *CreateOpsItemResponseBodyOpsItem {
	s.Solutions = &v
	return s
}

func (s *CreateOpsItemResponseBodyOpsItem) SetSource(v string) *CreateOpsItemResponseBodyOpsItem {
	s.Source = &v
	return s
}

func (s *CreateOpsItemResponseBodyOpsItem) SetStatus(v string) *CreateOpsItemResponseBodyOpsItem {
	s.Status = &v
	return s
}

func (s *CreateOpsItemResponseBodyOpsItem) SetTags(v map[string]interface{}) *CreateOpsItemResponseBodyOpsItem {
	s.Tags = v
	return s
}

func (s *CreateOpsItemResponseBodyOpsItem) SetTitle(v string) *CreateOpsItemResponseBodyOpsItem {
	s.Title = &v
	return s
}

func (s *CreateOpsItemResponseBodyOpsItem) SetUpdateDate(v string) *CreateOpsItemResponseBodyOpsItem {
	s.UpdateDate = &v
	return s
}

func (s *CreateOpsItemResponseBodyOpsItem) Validate() error {
	return dara.Validate(s)
}
