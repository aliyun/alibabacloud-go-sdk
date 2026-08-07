// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCategoryTemplateListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *DescribeCategoryTemplateListResponseBody
	GetCurrentPage() *int32
	SetItems(v []*DescribeCategoryTemplateListResponseBodyItems) *DescribeCategoryTemplateListResponseBody
	GetItems() []*DescribeCategoryTemplateListResponseBodyItems
	SetPageSize(v int32) *DescribeCategoryTemplateListResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeCategoryTemplateListResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeCategoryTemplateListResponseBody
	GetTotalCount() *int32
}

type DescribeCategoryTemplateListResponseBody struct {
	// The page number in a paged query. Default value: 1.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The list of industry templates.
	Items []*DescribeCategoryTemplateListResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The number of entries per page in a paging query. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The unique ID that Alibaba Cloud generates for the request.
	//
	// example:
	//
	// 8491DBFD-48C0-4E11-B6FC-6F38921244A9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 12
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeCategoryTemplateListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCategoryTemplateListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCategoryTemplateListResponseBody) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeCategoryTemplateListResponseBody) GetItems() []*DescribeCategoryTemplateListResponseBodyItems {
	return s.Items
}

func (s *DescribeCategoryTemplateListResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeCategoryTemplateListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCategoryTemplateListResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeCategoryTemplateListResponseBody) SetCurrentPage(v int32) *DescribeCategoryTemplateListResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *DescribeCategoryTemplateListResponseBody) SetItems(v []*DescribeCategoryTemplateListResponseBodyItems) *DescribeCategoryTemplateListResponseBody {
	s.Items = v
	return s
}

func (s *DescribeCategoryTemplateListResponseBody) SetPageSize(v int32) *DescribeCategoryTemplateListResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeCategoryTemplateListResponseBody) SetRequestId(v string) *DescribeCategoryTemplateListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCategoryTemplateListResponseBody) SetTotalCount(v int32) *DescribeCategoryTemplateListResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeCategoryTemplateListResponseBody) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeCategoryTemplateListResponseBodyItems struct {
	// The ID of the current risk level.
	//
	// example:
	//
	// 5
	CurrentRiskLevel *int32 `json:"CurrentRiskLevel,omitempty" xml:"CurrentRiskLevel,omitempty"`
	// The description of the industry template.
	//
	// example:
	//
	// description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The time when the industry template was created.
	//
	// example:
	//
	// 1582992000000
	GmtCreate *int64 `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The time when the industry template was last modified.
	//
	// example:
	//
	// 1545277010000
	GmtModified *int64 `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The unique ID of the industry template.
	//
	// example:
	//
	// 1
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The maximum category level.
	//
	// example:
	//
	// 2
	MaxCategoryLevel *int32 `json:"MaxCategoryLevel,omitempty" xml:"MaxCategoryLevel,omitempty"`
	// The ID of the maximum risk level.
	//
	// example:
	//
	// 5
	MaxRiskLevel *int32 `json:"MaxRiskLevel,omitempty" xml:"MaxRiskLevel,omitempty"`
	// The name of the industry template.
	//
	// example:
	//
	// built-in template
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The status of the industry template. Valid values:
	//
	// - **0**: Disabled.
	//
	// - **1**: Enabled. The template is the current primary template of the user.
	//
	// - **2**: Active. Both enabled and active templates can be applied to detection tasks.
	//
	// - **3**: General detection model template.
	//
	// > Templates in the enabled or active state can be used as the industry template ID parameter for [DescribeDataObjects](https://help.aliyun.com/document_detail/2399253.html).
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// Indicates whether the industry template supports editing. Valid values:
	//
	// - **0**: Not supported.
	//
	// - **1**: Supported.
	//
	// example:
	//
	// 0
	SupportEdit *int32 `json:"SupportEdit,omitempty" xml:"SupportEdit,omitempty"`
	// The type of the industry template.
	//
	// example:
	//
	// 6
	Type *int32 `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeCategoryTemplateListResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeCategoryTemplateListResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeCategoryTemplateListResponseBodyItems) GetCurrentRiskLevel() *int32 {
	return s.CurrentRiskLevel
}

func (s *DescribeCategoryTemplateListResponseBodyItems) GetDescription() *string {
	return s.Description
}

func (s *DescribeCategoryTemplateListResponseBodyItems) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *DescribeCategoryTemplateListResponseBodyItems) GetGmtModified() *int64 {
	return s.GmtModified
}

func (s *DescribeCategoryTemplateListResponseBodyItems) GetId() *int64 {
	return s.Id
}

func (s *DescribeCategoryTemplateListResponseBodyItems) GetMaxCategoryLevel() *int32 {
	return s.MaxCategoryLevel
}

func (s *DescribeCategoryTemplateListResponseBodyItems) GetMaxRiskLevel() *int32 {
	return s.MaxRiskLevel
}

func (s *DescribeCategoryTemplateListResponseBodyItems) GetName() *string {
	return s.Name
}

func (s *DescribeCategoryTemplateListResponseBodyItems) GetStatus() *int32 {
	return s.Status
}

func (s *DescribeCategoryTemplateListResponseBodyItems) GetSupportEdit() *int32 {
	return s.SupportEdit
}

func (s *DescribeCategoryTemplateListResponseBodyItems) GetType() *int32 {
	return s.Type
}

func (s *DescribeCategoryTemplateListResponseBodyItems) SetCurrentRiskLevel(v int32) *DescribeCategoryTemplateListResponseBodyItems {
	s.CurrentRiskLevel = &v
	return s
}

func (s *DescribeCategoryTemplateListResponseBodyItems) SetDescription(v string) *DescribeCategoryTemplateListResponseBodyItems {
	s.Description = &v
	return s
}

func (s *DescribeCategoryTemplateListResponseBodyItems) SetGmtCreate(v int64) *DescribeCategoryTemplateListResponseBodyItems {
	s.GmtCreate = &v
	return s
}

func (s *DescribeCategoryTemplateListResponseBodyItems) SetGmtModified(v int64) *DescribeCategoryTemplateListResponseBodyItems {
	s.GmtModified = &v
	return s
}

func (s *DescribeCategoryTemplateListResponseBodyItems) SetId(v int64) *DescribeCategoryTemplateListResponseBodyItems {
	s.Id = &v
	return s
}

func (s *DescribeCategoryTemplateListResponseBodyItems) SetMaxCategoryLevel(v int32) *DescribeCategoryTemplateListResponseBodyItems {
	s.MaxCategoryLevel = &v
	return s
}

func (s *DescribeCategoryTemplateListResponseBodyItems) SetMaxRiskLevel(v int32) *DescribeCategoryTemplateListResponseBodyItems {
	s.MaxRiskLevel = &v
	return s
}

func (s *DescribeCategoryTemplateListResponseBodyItems) SetName(v string) *DescribeCategoryTemplateListResponseBodyItems {
	s.Name = &v
	return s
}

func (s *DescribeCategoryTemplateListResponseBodyItems) SetStatus(v int32) *DescribeCategoryTemplateListResponseBodyItems {
	s.Status = &v
	return s
}

func (s *DescribeCategoryTemplateListResponseBodyItems) SetSupportEdit(v int32) *DescribeCategoryTemplateListResponseBodyItems {
	s.SupportEdit = &v
	return s
}

func (s *DescribeCategoryTemplateListResponseBodyItems) SetType(v int32) *DescribeCategoryTemplateListResponseBodyItems {
	s.Type = &v
	return s
}

func (s *DescribeCategoryTemplateListResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
