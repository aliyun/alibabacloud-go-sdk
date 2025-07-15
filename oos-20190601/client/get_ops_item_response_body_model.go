// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOpsItemResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOpsItem(v *GetOpsItemResponseBodyOpsItem) *GetOpsItemResponseBody
	GetOpsItem() *GetOpsItemResponseBodyOpsItem
	SetRequestId(v string) *GetOpsItemResponseBody
	GetRequestId() *string
}

type GetOpsItemResponseBody struct {
	// The information about the O\\&M item.
	OpsItem *GetOpsItemResponseBodyOpsItem `json:"OpsItem,omitempty" xml:"OpsItem,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 8BED4C16-BD30-5E27-94D4-7EBCCECF70C1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetOpsItemResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetOpsItemResponseBody) GoString() string {
	return s.String()
}

func (s *GetOpsItemResponseBody) GetOpsItem() *GetOpsItemResponseBodyOpsItem {
	return s.OpsItem
}

func (s *GetOpsItemResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetOpsItemResponseBody) SetOpsItem(v *GetOpsItemResponseBodyOpsItem) *GetOpsItemResponseBody {
	s.OpsItem = v
	return s
}

func (s *GetOpsItemResponseBody) SetRequestId(v string) *GetOpsItemResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOpsItemResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetOpsItemResponseBodyOpsItem struct {
	// The information about the attributes of the O\\&M item.
	Attributes map[string]interface{} `json:"Attributes,omitempty" xml:"Attributes,omitempty"`
	// The category of the O\\&M item.
	//
	// example:
	//
	// Security
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The user who created the O\\&M item.
	//
	// example:
	//
	// root(130900000)
	CreateBy *string `json:"CreateBy,omitempty" xml:"CreateBy,omitempty"`
	// The time when the O\\&M item was created.
	//
	// example:
	//
	// 2023-04-10T06:15Z
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	// The description.
	//
	// example:
	//
	// test-update
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The user who last modified the O\\&M item.
	//
	// example:
	//
	// modifiedBy
	LastModifiedBy *string `json:"LastModifiedBy,omitempty" xml:"LastModifiedBy,omitempty"`
	// The O\\&M item ID.
	//
	// example:
	//
	// oi-d52b08695e2b46ae8413
	OpsItemId *string `json:"OpsItemId,omitempty" xml:"OpsItemId,omitempty"`
	// The priority of the O\\&M item.
	//
	// example:
	//
	// 1
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-aekzxkofnlxtn2i
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The Alibaba Cloud Resource Names (ARNs) of the associated resources.
	Resources []*string `json:"Resources,omitempty" xml:"Resources,omitempty" type:"Repeated"`
	// The severity level of the O\\&M item.
	//
	// example:
	//
	// Medium
	Severity *string `json:"Severity,omitempty" xml:"Severity,omitempty"`
	// The solutions to the O\\&M item.
	Solutions []map[string]interface{} `json:"Solutions,omitempty" xml:"Solutions,omitempty" type:"Repeated"`
	// The source business of the O\\&M item.
	//
	// example:
	//
	// /aliyun/appManager
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The status of the O\\&M item.
	//
	// example:
	//
	// Open
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tags attached to the O\\&M item.
	//
	// example:
	//
	// {"K1":"V1"}
	Tags map[string]interface{} `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The title of the O\\&M item.
	//
	// example:
	//
	// test
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// The time when the O\\&M item was updated.
	//
	// example:
	//
	// 2023-04-10T06:15Z
	UpdateDate *string `json:"UpdateDate,omitempty" xml:"UpdateDate,omitempty"`
}

func (s GetOpsItemResponseBodyOpsItem) String() string {
	return dara.Prettify(s)
}

func (s GetOpsItemResponseBodyOpsItem) GoString() string {
	return s.String()
}

func (s *GetOpsItemResponseBodyOpsItem) GetAttributes() map[string]interface{} {
	return s.Attributes
}

func (s *GetOpsItemResponseBodyOpsItem) GetCategory() *string {
	return s.Category
}

func (s *GetOpsItemResponseBodyOpsItem) GetCreateBy() *string {
	return s.CreateBy
}

func (s *GetOpsItemResponseBodyOpsItem) GetCreateDate() *string {
	return s.CreateDate
}

func (s *GetOpsItemResponseBodyOpsItem) GetDescription() *string {
	return s.Description
}

func (s *GetOpsItemResponseBodyOpsItem) GetLastModifiedBy() *string {
	return s.LastModifiedBy
}

func (s *GetOpsItemResponseBodyOpsItem) GetOpsItemId() *string {
	return s.OpsItemId
}

func (s *GetOpsItemResponseBodyOpsItem) GetPriority() *int32 {
	return s.Priority
}

func (s *GetOpsItemResponseBodyOpsItem) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *GetOpsItemResponseBodyOpsItem) GetResources() []*string {
	return s.Resources
}

func (s *GetOpsItemResponseBodyOpsItem) GetSeverity() *string {
	return s.Severity
}

func (s *GetOpsItemResponseBodyOpsItem) GetSolutions() []map[string]interface{} {
	return s.Solutions
}

func (s *GetOpsItemResponseBodyOpsItem) GetSource() *string {
	return s.Source
}

func (s *GetOpsItemResponseBodyOpsItem) GetStatus() *string {
	return s.Status
}

func (s *GetOpsItemResponseBodyOpsItem) GetTags() map[string]interface{} {
	return s.Tags
}

func (s *GetOpsItemResponseBodyOpsItem) GetTitle() *string {
	return s.Title
}

func (s *GetOpsItemResponseBodyOpsItem) GetUpdateDate() *string {
	return s.UpdateDate
}

func (s *GetOpsItemResponseBodyOpsItem) SetAttributes(v map[string]interface{}) *GetOpsItemResponseBodyOpsItem {
	s.Attributes = v
	return s
}

func (s *GetOpsItemResponseBodyOpsItem) SetCategory(v string) *GetOpsItemResponseBodyOpsItem {
	s.Category = &v
	return s
}

func (s *GetOpsItemResponseBodyOpsItem) SetCreateBy(v string) *GetOpsItemResponseBodyOpsItem {
	s.CreateBy = &v
	return s
}

func (s *GetOpsItemResponseBodyOpsItem) SetCreateDate(v string) *GetOpsItemResponseBodyOpsItem {
	s.CreateDate = &v
	return s
}

func (s *GetOpsItemResponseBodyOpsItem) SetDescription(v string) *GetOpsItemResponseBodyOpsItem {
	s.Description = &v
	return s
}

func (s *GetOpsItemResponseBodyOpsItem) SetLastModifiedBy(v string) *GetOpsItemResponseBodyOpsItem {
	s.LastModifiedBy = &v
	return s
}

func (s *GetOpsItemResponseBodyOpsItem) SetOpsItemId(v string) *GetOpsItemResponseBodyOpsItem {
	s.OpsItemId = &v
	return s
}

func (s *GetOpsItemResponseBodyOpsItem) SetPriority(v int32) *GetOpsItemResponseBodyOpsItem {
	s.Priority = &v
	return s
}

func (s *GetOpsItemResponseBodyOpsItem) SetResourceGroupId(v string) *GetOpsItemResponseBodyOpsItem {
	s.ResourceGroupId = &v
	return s
}

func (s *GetOpsItemResponseBodyOpsItem) SetResources(v []*string) *GetOpsItemResponseBodyOpsItem {
	s.Resources = v
	return s
}

func (s *GetOpsItemResponseBodyOpsItem) SetSeverity(v string) *GetOpsItemResponseBodyOpsItem {
	s.Severity = &v
	return s
}

func (s *GetOpsItemResponseBodyOpsItem) SetSolutions(v []map[string]interface{}) *GetOpsItemResponseBodyOpsItem {
	s.Solutions = v
	return s
}

func (s *GetOpsItemResponseBodyOpsItem) SetSource(v string) *GetOpsItemResponseBodyOpsItem {
	s.Source = &v
	return s
}

func (s *GetOpsItemResponseBodyOpsItem) SetStatus(v string) *GetOpsItemResponseBodyOpsItem {
	s.Status = &v
	return s
}

func (s *GetOpsItemResponseBodyOpsItem) SetTags(v map[string]interface{}) *GetOpsItemResponseBodyOpsItem {
	s.Tags = v
	return s
}

func (s *GetOpsItemResponseBodyOpsItem) SetTitle(v string) *GetOpsItemResponseBodyOpsItem {
	s.Title = &v
	return s
}

func (s *GetOpsItemResponseBodyOpsItem) SetUpdateDate(v string) *GetOpsItemResponseBodyOpsItem {
	s.UpdateDate = &v
	return s
}

func (s *GetOpsItemResponseBodyOpsItem) Validate() error {
	return dara.Validate(s)
}
