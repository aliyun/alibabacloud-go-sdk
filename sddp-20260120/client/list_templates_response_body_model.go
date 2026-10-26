// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTemplatesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *ListTemplatesResponseBody
	GetCurrentPage() *int32
	SetItems(v []*ListTemplatesResponseBodyItems) *ListTemplatesResponseBody
	GetItems() []*ListTemplatesResponseBodyItems
	SetPageSize(v int32) *ListTemplatesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListTemplatesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListTemplatesResponseBody
	GetTotalCount() *int32
}

type ListTemplatesResponseBody struct {
	// The page number of the returned data.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The list of records returned for the current query.
	Items []*ListTemplatesResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The number of records per page for the current query.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID, which can be used to trace the current API call.
	//
	// example:
	//
	// 7C6D8E9F-1234-5678-ABCD-0123456789AB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of records that match the current query conditions.
	//
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListTemplatesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *ListTemplatesResponseBody) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListTemplatesResponseBody) GetItems() []*ListTemplatesResponseBodyItems {
	return s.Items
}

func (s *ListTemplatesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListTemplatesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTemplatesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListTemplatesResponseBody) SetCurrentPage(v int32) *ListTemplatesResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *ListTemplatesResponseBody) SetItems(v []*ListTemplatesResponseBodyItems) *ListTemplatesResponseBody {
	s.Items = v
	return s
}

func (s *ListTemplatesResponseBody) SetPageSize(v int32) *ListTemplatesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListTemplatesResponseBody) SetRequestId(v string) *ListTemplatesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTemplatesResponseBody) SetTotalCount(v int32) *ListTemplatesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListTemplatesResponseBody) Validate() error {
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

type ListTemplatesResponseBodyItems struct {
	// The current sensitivity level of the template.
	//
	// example:
	//
	// 3
	CurrentRiskLevel *int32 `json:"CurrentRiskLevel,omitempty" xml:"CurrentRiskLevel,omitempty"`
	// The description of the classification and grading template.
	//
	// example:
	//
	// Used to detect common sensitive data
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The creation time of the template. The value is a UNIX timestamp in milliseconds since 1970-01-01 00:00:00 UTC.
	//
	// example:
	//
	// 1788566400000
	GmtCreate *int64 `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The last modification time of the template. The value is a UNIX timestamp in milliseconds since 1970-01-01 00:00:00 UTC.
	//
	// example:
	//
	// 1788566400000
	GmtModified *int64 `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The identifier of the classification and grading template.
	//
	// example:
	//
	// 1
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The maximum classification level supported by the template.
	//
	// example:
	//
	// 3
	MaxCategoryLevel *int32 `json:"MaxCategoryLevel,omitempty" xml:"MaxCategoryLevel,omitempty"`
	// The highest sensitivity level supported by the template.
	//
	// example:
	//
	// 5
	MaxRiskLevel *int32 `json:"MaxRiskLevel,omitempty" xml:"MaxRiskLevel,omitempty"`
	// The name of the classification and grading template.
	//
	// example:
	//
	// General classification and grading template
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The status of the template. Valid values: 0: Disabled. 1: Enabled primary template. 2: Active template. 3: General detection model.
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// Specifies whether the template can be edited. Valid values: 0: Not supported. 1: Supported.
	//
	// example:
	//
	// 1
	SupportEdit *int32 `json:"SupportEdit,omitempty" xml:"SupportEdit,omitempty"`
	// The category of the template. Valid values: 1: Alibaba Cloud template. 2: Finance template. 3: Energy template. 4: Personal information security template. 5: Travel template. 65: General template. 66: Custom template. 67: Template from other cloud services.
	//
	// example:
	//
	// 65
	Type *int32 `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListTemplatesResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s ListTemplatesResponseBodyItems) GoString() string {
	return s.String()
}

func (s *ListTemplatesResponseBodyItems) GetCurrentRiskLevel() *int32 {
	return s.CurrentRiskLevel
}

func (s *ListTemplatesResponseBodyItems) GetDescription() *string {
	return s.Description
}

func (s *ListTemplatesResponseBodyItems) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *ListTemplatesResponseBodyItems) GetGmtModified() *int64 {
	return s.GmtModified
}

func (s *ListTemplatesResponseBodyItems) GetId() *int64 {
	return s.Id
}

func (s *ListTemplatesResponseBodyItems) GetMaxCategoryLevel() *int32 {
	return s.MaxCategoryLevel
}

func (s *ListTemplatesResponseBodyItems) GetMaxRiskLevel() *int32 {
	return s.MaxRiskLevel
}

func (s *ListTemplatesResponseBodyItems) GetName() *string {
	return s.Name
}

func (s *ListTemplatesResponseBodyItems) GetStatus() *int32 {
	return s.Status
}

func (s *ListTemplatesResponseBodyItems) GetSupportEdit() *int32 {
	return s.SupportEdit
}

func (s *ListTemplatesResponseBodyItems) GetType() *int32 {
	return s.Type
}

func (s *ListTemplatesResponseBodyItems) SetCurrentRiskLevel(v int32) *ListTemplatesResponseBodyItems {
	s.CurrentRiskLevel = &v
	return s
}

func (s *ListTemplatesResponseBodyItems) SetDescription(v string) *ListTemplatesResponseBodyItems {
	s.Description = &v
	return s
}

func (s *ListTemplatesResponseBodyItems) SetGmtCreate(v int64) *ListTemplatesResponseBodyItems {
	s.GmtCreate = &v
	return s
}

func (s *ListTemplatesResponseBodyItems) SetGmtModified(v int64) *ListTemplatesResponseBodyItems {
	s.GmtModified = &v
	return s
}

func (s *ListTemplatesResponseBodyItems) SetId(v int64) *ListTemplatesResponseBodyItems {
	s.Id = &v
	return s
}

func (s *ListTemplatesResponseBodyItems) SetMaxCategoryLevel(v int32) *ListTemplatesResponseBodyItems {
	s.MaxCategoryLevel = &v
	return s
}

func (s *ListTemplatesResponseBodyItems) SetMaxRiskLevel(v int32) *ListTemplatesResponseBodyItems {
	s.MaxRiskLevel = &v
	return s
}

func (s *ListTemplatesResponseBodyItems) SetName(v string) *ListTemplatesResponseBodyItems {
	s.Name = &v
	return s
}

func (s *ListTemplatesResponseBodyItems) SetStatus(v int32) *ListTemplatesResponseBodyItems {
	s.Status = &v
	return s
}

func (s *ListTemplatesResponseBodyItems) SetSupportEdit(v int32) *ListTemplatesResponseBodyItems {
	s.SupportEdit = &v
	return s
}

func (s *ListTemplatesResponseBodyItems) SetType(v int32) *ListTemplatesResponseBodyItems {
	s.Type = &v
	return s
}

func (s *ListTemplatesResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
