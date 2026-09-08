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
	CurrentPage *int32                            `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	Items       []*ListTemplatesResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	PageSize    *int32                            `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId   *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount  *int32                            `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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
	CurrentRiskLevel *int32  `json:"CurrentRiskLevel,omitempty" xml:"CurrentRiskLevel,omitempty"`
	Description      *string `json:"Description,omitempty" xml:"Description,omitempty"`
	GmtCreate        *int64  `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified      *int64  `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	Id               *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	MaxCategoryLevel *int32  `json:"MaxCategoryLevel,omitempty" xml:"MaxCategoryLevel,omitempty"`
	MaxRiskLevel     *int32  `json:"MaxRiskLevel,omitempty" xml:"MaxRiskLevel,omitempty"`
	Name             *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Status           *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	SupportEdit      *int32  `json:"SupportEdit,omitempty" xml:"SupportEdit,omitempty"`
	Type             *int32  `json:"Type,omitempty" xml:"Type,omitempty"`
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
