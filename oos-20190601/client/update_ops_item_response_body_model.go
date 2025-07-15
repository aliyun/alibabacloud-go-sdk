// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateOpsItemResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOpsItem(v *UpdateOpsItemResponseBodyOpsItem) *UpdateOpsItemResponseBody
	GetOpsItem() *UpdateOpsItemResponseBodyOpsItem
	SetRequestId(v string) *UpdateOpsItemResponseBody
	GetRequestId() *string
}

type UpdateOpsItemResponseBody struct {
	// The information about the O\\&M item.
	OpsItem *UpdateOpsItemResponseBodyOpsItem `json:"OpsItem,omitempty" xml:"OpsItem,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// C996DECB-3D2B-5321-B359-BE7031B6399E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateOpsItemResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateOpsItemResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateOpsItemResponseBody) GetOpsItem() *UpdateOpsItemResponseBodyOpsItem {
	return s.OpsItem
}

func (s *UpdateOpsItemResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateOpsItemResponseBody) SetOpsItem(v *UpdateOpsItemResponseBodyOpsItem) *UpdateOpsItemResponseBody {
	s.OpsItem = v
	return s
}

func (s *UpdateOpsItemResponseBody) SetRequestId(v string) *UpdateOpsItemResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateOpsItemResponseBody) Validate() error {
	return dara.Validate(s)
}

type UpdateOpsItemResponseBodyOpsItem struct {
	// The attributes of the O\\&M item.
	//
	// example:
	//
	// [{\\"Attribute\\": {\\"Weight\\": 100}, \\"RealServer\\": \\"uaejc8hnqzqz5valyh8dibolpvza48ik.yundunwaf5.com\\"}]
	Attributes *string `json:"Attributes,omitempty" xml:"Attributes,omitempty"`
	// The category.
	//
	// example:
	//
	// Security
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The time when the O\\&M item was created.
	//
	// example:
	//
	// 2023-03-16T07:04Z
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	// The user who created the patch baseline.
	//
	// example:
	//
	// root(130900000)
	CreatedBy *string `json:"CreatedBy,omitempty" xml:"CreatedBy,omitempty"`
	// The description.
	//
	// example:
	//
	// test-update
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The user who modified the O\\&M item.
	//
	// example:
	//
	// root(130900000)
	LastModifiedBy *string `json:"LastModifiedBy,omitempty" xml:"LastModifiedBy,omitempty"`
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
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmxsn4m4******
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The ARNs of the associated resources.
	Resources []*string `json:"Resources,omitempty" xml:"Resources,omitempty" type:"Repeated"`
	// The severity level.
	//
	// example:
	//
	// Medium
	Severity *string `json:"Severity,omitempty" xml:"Severity,omitempty"`
	// The solutions.
	Solutions []*string `json:"Solutions,omitempty" xml:"Solutions,omitempty" type:"Repeated"`
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
	// The time when the O\\&M item was updated.
	//
	// example:
	//
	// 2023-03-16T08:04Z
	UpdateDate *string `json:"UpdateDate,omitempty" xml:"UpdateDate,omitempty"`
}

func (s UpdateOpsItemResponseBodyOpsItem) String() string {
	return dara.Prettify(s)
}

func (s UpdateOpsItemResponseBodyOpsItem) GoString() string {
	return s.String()
}

func (s *UpdateOpsItemResponseBodyOpsItem) GetAttributes() *string {
	return s.Attributes
}

func (s *UpdateOpsItemResponseBodyOpsItem) GetCategory() *string {
	return s.Category
}

func (s *UpdateOpsItemResponseBodyOpsItem) GetCreateDate() *string {
	return s.CreateDate
}

func (s *UpdateOpsItemResponseBodyOpsItem) GetCreatedBy() *string {
	return s.CreatedBy
}

func (s *UpdateOpsItemResponseBodyOpsItem) GetDescription() *string {
	return s.Description
}

func (s *UpdateOpsItemResponseBodyOpsItem) GetLastModifiedBy() *string {
	return s.LastModifiedBy
}

func (s *UpdateOpsItemResponseBodyOpsItem) GetOpsItemId() *string {
	return s.OpsItemId
}

func (s *UpdateOpsItemResponseBodyOpsItem) GetPriority() *int32 {
	return s.Priority
}

func (s *UpdateOpsItemResponseBodyOpsItem) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *UpdateOpsItemResponseBodyOpsItem) GetResources() []*string {
	return s.Resources
}

func (s *UpdateOpsItemResponseBodyOpsItem) GetSeverity() *string {
	return s.Severity
}

func (s *UpdateOpsItemResponseBodyOpsItem) GetSolutions() []*string {
	return s.Solutions
}

func (s *UpdateOpsItemResponseBodyOpsItem) GetSource() *string {
	return s.Source
}

func (s *UpdateOpsItemResponseBodyOpsItem) GetStatus() *string {
	return s.Status
}

func (s *UpdateOpsItemResponseBodyOpsItem) GetTags() map[string]interface{} {
	return s.Tags
}

func (s *UpdateOpsItemResponseBodyOpsItem) GetTitle() *string {
	return s.Title
}

func (s *UpdateOpsItemResponseBodyOpsItem) GetUpdateDate() *string {
	return s.UpdateDate
}

func (s *UpdateOpsItemResponseBodyOpsItem) SetAttributes(v string) *UpdateOpsItemResponseBodyOpsItem {
	s.Attributes = &v
	return s
}

func (s *UpdateOpsItemResponseBodyOpsItem) SetCategory(v string) *UpdateOpsItemResponseBodyOpsItem {
	s.Category = &v
	return s
}

func (s *UpdateOpsItemResponseBodyOpsItem) SetCreateDate(v string) *UpdateOpsItemResponseBodyOpsItem {
	s.CreateDate = &v
	return s
}

func (s *UpdateOpsItemResponseBodyOpsItem) SetCreatedBy(v string) *UpdateOpsItemResponseBodyOpsItem {
	s.CreatedBy = &v
	return s
}

func (s *UpdateOpsItemResponseBodyOpsItem) SetDescription(v string) *UpdateOpsItemResponseBodyOpsItem {
	s.Description = &v
	return s
}

func (s *UpdateOpsItemResponseBodyOpsItem) SetLastModifiedBy(v string) *UpdateOpsItemResponseBodyOpsItem {
	s.LastModifiedBy = &v
	return s
}

func (s *UpdateOpsItemResponseBodyOpsItem) SetOpsItemId(v string) *UpdateOpsItemResponseBodyOpsItem {
	s.OpsItemId = &v
	return s
}

func (s *UpdateOpsItemResponseBodyOpsItem) SetPriority(v int32) *UpdateOpsItemResponseBodyOpsItem {
	s.Priority = &v
	return s
}

func (s *UpdateOpsItemResponseBodyOpsItem) SetResourceGroupId(v string) *UpdateOpsItemResponseBodyOpsItem {
	s.ResourceGroupId = &v
	return s
}

func (s *UpdateOpsItemResponseBodyOpsItem) SetResources(v []*string) *UpdateOpsItemResponseBodyOpsItem {
	s.Resources = v
	return s
}

func (s *UpdateOpsItemResponseBodyOpsItem) SetSeverity(v string) *UpdateOpsItemResponseBodyOpsItem {
	s.Severity = &v
	return s
}

func (s *UpdateOpsItemResponseBodyOpsItem) SetSolutions(v []*string) *UpdateOpsItemResponseBodyOpsItem {
	s.Solutions = v
	return s
}

func (s *UpdateOpsItemResponseBodyOpsItem) SetSource(v string) *UpdateOpsItemResponseBodyOpsItem {
	s.Source = &v
	return s
}

func (s *UpdateOpsItemResponseBodyOpsItem) SetStatus(v string) *UpdateOpsItemResponseBodyOpsItem {
	s.Status = &v
	return s
}

func (s *UpdateOpsItemResponseBodyOpsItem) SetTags(v map[string]interface{}) *UpdateOpsItemResponseBodyOpsItem {
	s.Tags = v
	return s
}

func (s *UpdateOpsItemResponseBodyOpsItem) SetTitle(v string) *UpdateOpsItemResponseBodyOpsItem {
	s.Title = &v
	return s
}

func (s *UpdateOpsItemResponseBodyOpsItem) SetUpdateDate(v string) *UpdateOpsItemResponseBodyOpsItem {
	s.UpdateDate = &v
	return s
}

func (s *UpdateOpsItemResponseBodyOpsItem) Validate() error {
	return dara.Validate(s)
}
